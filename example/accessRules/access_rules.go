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
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	sdm "github.com/strongdm/strongdm-sdk-go/v2"
)

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

	// create a redis
	redisID := createExampleResource(ctx, client)

	// create role with initial access rule
	role := &sdm.Role{
		Name: "accessRulesTestRole",
		AccessRules: sdm.AccessRules{
			sdm.AccessRule{
				IDs: []string{redisID},
			},
		},
	}
	roleResp, err := client.Roles().Create(ctx, role)
	if err != nil {
		log.Fatalf("failed to create role: %v", err)
	}
	role = roleResp.Role

	// update access rules
	role.AccessRules = sdm.AccessRules{
		sdm.AccessRule{
			Tags: sdm.Tags{
				"env": "staging",
			},
		},
		sdm.AccessRule{
			Type: "postgres",
		},
	}
	_, err = client.Roles().Update(ctx, role)
	if err != nil {
		log.Fatalf("failed to update role: %v", err)
	}

	// The RoleGrants API has been deprecated in favor of Access Rules.
	// When using Access Rules the best practice is to grant Resources access based on Type and Tags.
	// If it is _necessary_ to grant access to specific Resources in the same way as RoleGrants did,
	// you can use Resource IDs directly in Access Rules as shown in the following examples.

	err = createRoleGrantViaAccessRulesExample(ctx, client)
	if err != nil {
		log.Fatalf("error in createRoleGrantViaAccessRulesExample: %v", err)
	}
	err = deleteRoleGrantViaAccessRulesExample(ctx, client)
	if err != nil {
		log.Fatalf("error in deleteRoleGrantViaAccessRulesExample: %v", err)
	}
	err = listRoleGrantsViaAccessRulesExample(ctx, client)
	if err != nil {
		log.Fatalf("error in listRoleGrantsViaAccessRulesExample: %v", err)
	}
}

// createExampleRole create a role with empty access rules and return the ID
func createExampleRole(ctx context.Context, client *sdm.Client, ar sdm.AccessRules) string {
	role := &sdm.Role{
		Name:        "exampleRole-" + fmt.Sprint(rand.Int()),
		AccessRules: ar,
	}
	roleResp, err := client.Roles().Create(ctx, role)
	if err != nil {
		log.Fatalf("error creating role: %v", err)
	}
	return roleResp.Role.ID
}

// createExampleResource create a sample resource and return the ID
func createExampleResource(ctx context.Context, client *sdm.Client) string {
	redis := &sdm.Redis{
		Name:         "exampleResource-" + fmt.Sprint(rand.Int()),
		Hostname:     "example.com",
		Port:         6379,
		PortOverride: int32(rand.Intn(20000) + 3000),
	}
	resp, err := client.Resources().Create(ctx, redis)
	if err != nil {
		log.Fatalf("error creating resource: %v", err)
	}
	return resp.Resource.GetID()
}

func createRoleGrantViaAccessRulesExample(ctx context.Context, client *sdm.Client) error {
	// create example resources
	resourceID1 := createExampleResource(ctx, client)
	resourceID2 := createExampleResource(ctx, client)
	roleID := createExampleRole(ctx, client, sdm.AccessRules{
		sdm.AccessRule{
			IDs: []string{resourceID1},
		},
	})

	// get the role
	getResp, err := client.Roles().Get(ctx, roleID)
	if err != nil {
		return fmt.Errorf("error getting role: %v", err)
	}
	role := getResp.Role

	// append the ID to an existing static access rule
	if len(role.AccessRules) != 1 || len(role.AccessRules[0].IDs) == 0 {
		return fmt.Errorf("unexpected access rules in role")
	}
	role.AccessRules[0].IDs = append(role.AccessRules[0].IDs, resourceID2)

	// update the role
	_, err = client.Roles().Update(ctx, role)
	if err != nil {
		return fmt.Errorf("error updating role: %v", err)
	}
	return nil
}

func deleteRoleGrantViaAccessRulesExample(ctx context.Context, client *sdm.Client) error {
	// create example resources
	resourceID1 := createExampleResource(ctx, client)
	resourceID2 := createExampleResource(ctx, client)
	roleID := createExampleRole(ctx, client, sdm.AccessRules{
		sdm.AccessRule{
			IDs: []string{resourceID1, resourceID2},
		},
	})

	// get the role
	getResp, err := client.Roles().Get(ctx, roleID)
	if err != nil {
		return fmt.Errorf("error getting role: %v", err)
	}
	role := getResp.Role

	if len(role.AccessRules) != 1 || len(role.AccessRules[0].IDs) == 0 {
		return fmt.Errorf("unexpected access rules in role")
	}

	// remove the ID of the second Resource
	for i, r := range role.AccessRules[0].IDs {
		if r == resourceID2 {
			role.AccessRules[0].IDs = removeElement(role.AccessRules[0].IDs, i)
			break
		}
	}

	// update the role
	_, err = client.Roles().Update(ctx, role)
	if err != nil {
		return fmt.Errorf("error updating role: %v", err)
	}
	return nil
}

func listRoleGrantsViaAccessRulesExample(ctx context.Context, client *sdm.Client) error {
	// create example resources
	resourceID := createExampleResource(ctx, client)
	roleID := createExampleRole(ctx, client, sdm.AccessRules{
		sdm.AccessRule{
			IDs: []string{resourceID},
		},
	})

	// get the role
	getResp, err := client.Roles().Get(ctx, roleID)
	if err != nil {
		return fmt.Errorf("error getting role: %v", err)
	}
	role := getResp.Role

	// role.AccessRules contains each AccessRule associated with the Role
	for _, resourceID := range role.AccessRules[0].IDs {
		fmt.Println(resourceID)
	}

	return nil
}

func removeElement(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
