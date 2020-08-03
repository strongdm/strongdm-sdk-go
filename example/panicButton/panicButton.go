// Copyright 2020 StrongDM Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	sdm "github.com/strongdm/strongdm-sdk-go"
)

type State struct {
	Attachments []*sdm.AccountAttachment `json:"attachments,omitempty"`
	Grants      []*sdm.AccountGrant      `json:"grants,omitempty"`
}

func main() {
	log.SetFlags(0)
	accessKey := os.Getenv("SDM_API_ACCESS_KEY")
	secretKey := os.Getenv("SDM_API_SECRET_KEY")
	if accessKey == "" || secretKey == "" {
		log.Fatal("SDM_API_ACCESS_KEY and SDM_API_SECRET_KEY must be provided")
	}

	client, err := sdm.New(accessKey, secretKey)
	if err != nil {
		log.Fatal("failed to create strongDM client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if len(os.Args) == 2 && os.Args[1] == "revert" {
		f, err := os.Open("state.json")
		if err != nil {
			log.Fatal("no state file found:", err)
		}

		var state State
		err = json.NewDecoder(f).Decode(&state)
		if err != nil {
			log.Fatal("malformed state file:", err)
		}

		accountResp, err := client.Accounts().List(ctx, "")
		if err != nil {
			log.Fatal("failed to query accounts:", err)
		}

		reinstated := 0
		recreatedAttachments := 0
		recreatedGrants := 0

		for accountResp.Next() {
			account := accountResp.Value()
			if !account.IsSuspended() {
				continue
			}
			reinstated++
			switch v := account.(type) {
			case *sdm.User:
				v.Suspended = false
			case *sdm.Service:
				v.Suspended = false
			}
			_, err := client.Accounts().Update(ctx, account)
			if err != nil {
				log.Fatal("failed to reinstate account:", err)
			}
		}
		if err := accountResp.Err(); err != nil {
			log.Fatal("account list failed:", err)
		}
		for _, attachment := range state.Attachments {
			_, err = client.AccountAttachments().Create(ctx, attachment)
			var alreadyExists *sdm.AlreadyExistsError
			if errors.As(err, &alreadyExists) {
				continue
			} else if err != nil {
				log.Println("skipping creation of attachment due to error:", err)
				continue
			}
			recreatedAttachments++
		}
		for _, grant := range state.Grants {
			_, err = client.AccountGrants().Create(ctx, grant)
			var alreadyExists *sdm.AlreadyExistsError
			if errors.As(err, &alreadyExists) {
				continue
			} else if err != nil {
				log.Println("skipping creation of grant due to error:", err)
				continue
			}
			recreatedGrants++
		}
		fmt.Printf("reinstated %d users\n", reinstated)
		fmt.Printf("recreated %d account attachments\n", recreatedAttachments)
		fmt.Printf("recreated %d account grants\n", recreatedGrants)
		return
	}

	adminEmail := ""
	if len(os.Args) == 2 {
		adminEmail = os.Args[1]
	} else {
		fmt.Println("please provide an admin email to preserve")
		os.Exit(1)
	}

	adminUserID := ""
	accountResp, err := client.Accounts().List(ctx, "email:?", adminEmail)
	if err != nil {
		log.Fatal("failed to list accounts:", err)
	}
	for accountResp.Next() {
		adminUserID = accountResp.Value().GetID()
	}
	if err := accountResp.Err(); err != nil {
		log.Fatal("account list failed:", err)
	}

	accountAttachments, err := client.AccountAttachments().List(ctx, "")
	if err != nil {
		log.Fatal("failed to list account attachments:", err)
	}
	accountGrants, err := client.AccountGrants().List(ctx, "")
	if err != nil {
		log.Fatal("failed to list account grants:", err)
	}

	state := &State{}

	for accountAttachments.Next() {
		attachment := accountAttachments.Value()
		if attachment.AccountID != adminUserID {
			state.Attachments = append(state.Attachments, attachment)
		}
	}
	if err := accountAttachments.Err(); err != nil {
		log.Fatal("account attachment list failed:", err)
	}

	for accountGrants.Next() {
		grant := accountGrants.Value()
		if grant.AccountID != adminUserID {
			state.Grants = append(state.Grants, grant)
		}
	}
	if err := accountGrants.Err(); err != nil {
		log.Fatal("account grant list failed:", err)
	}

	fmt.Printf("storing %d account attachments in state\n", len(state.Attachments))
	fmt.Printf("storing %d account grants in state\n", len(state.Grants))

	f, err := os.Create("state.json")
	if err != nil {
		log.Fatal("failed to create state file:", err)
	}

	err = json.NewEncoder(f).Encode(state)
	if err != nil {
		log.Fatal("failed to encode state file:", err)
	}

	suspendedCount := 0

	accountResp, err = client.Accounts().List(ctx, "")
	if err != nil {
		log.Fatal("failed to list accounts:", err)
	}
	for accountResp.Next() {
		account := accountResp.Value()
		switch v := account.(type) {
		case *sdm.User:
			if v.Email == adminEmail {
				continue
			}
			v.Suspended = true
		case *sdm.Service:
			v.Suspended = true
		}
		_, err := client.Accounts().Update(ctx, account)
		if err != nil {
			log.Printf("skipping user %s on account of error %v", account.GetID(), err)
			continue
		}
		suspendedCount++
	}

	fmt.Printf("suspended %d users\n", suspendedCount)
}
