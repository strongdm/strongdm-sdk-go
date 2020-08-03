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
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/okta/okta-sdk-golang/okta"
	"github.com/okta/okta-sdk-golang/okta/query"
	"github.com/pkg/errors"
	sdm "github.com/strongdm/strongdm-sdk-go"
	"gopkg.in/yaml.v2"
)

var verbose = flag.Bool("json", false, "dump a JSON report for debugging")
var plan = flag.Bool("plan", false, "do not apply changes just plan and output the result")

func init() {
	flag.Parse()
}

type syncReport struct {
	Start         time.Time  `json:"start"`
	Complete      time.Time  `json:"complete"`
	OktaUserCount int        `json:"oktaUsersCount"`
	OktaUsers     []oktaUser `json:"oktaUsers"`

	SDMUserCount int       `json:"sdmUsersCount"`
	SDMUsers     []userRow `json:"sdmUsers"`

	BothUserCount int `json:"bothUsersCount"`

	SDMResourcesCount int           `json:"sdmResourcesCount"`
	SDMResources      []entitlement `json:"sdmResources"`

	PermissionsGranted int              `json:"permissionsGranted"`
	PermissionsRevoked int              `json:"permissionsRevoked"`
	Grants             []entitlement    `json:"grants"`
	Revocations        []permissionsRow `json:"revocations"`

	Matchers *MatcherConfig `json:"matchers"`
}

func (rpt *syncReport) String() string {
	if !*verbose {
		return rpt.short()
	}

	out, err := json.MarshalIndent(rpt, "", "\t")
	if err != nil {
		return fmt.Sprintf("error building JSON report: %s\n\n%s", err, rpt.short())
	}
	return string(out)
}

func (rpt *syncReport) short() string {
	return fmt.Sprintf("%d Okta users, %d strongDM users, %d overlapping users, %d grants, %d revocations\n",
		rpt.OktaUserCount, rpt.SDMUserCount, rpt.BothUserCount,
		rpt.PermissionsGranted, rpt.PermissionsRevoked)
}

func main() {
	ctx := context.Background()

	if os.Getenv("SDM_API_ACCESS_KEY") == "" ||
		os.Getenv("SDM_API_SECRET_KEY") == "" ||
		os.Getenv("OKTA_CLIENT_TOKEN") == "" ||
		os.Getenv("OKTA_CLIENT_ORGURL") == "" {
		fmt.Println("SDM_API_ACCESS_KEY, SDM_API_SECRET_KEY, OKTA_CLIENT_TOKEN, and OKTA_CLIENT_ORGURL must be set")
		os.Exit(1)
		return
	}

	client, err := sdm.New(os.Getenv("SDM_API_ACCESS_KEY"), os.Getenv("SDM_API_SECRET_KEY"))
	if err != nil {
		fmt.Println("failed to initialize strongDM client: ", err)
		os.Exit(1)
		return
	}

	var rpt syncReport
	rpt.Start = time.Now()

	matchers, err := loadMatchers()
	if err != nil {
		fmt.Printf("error loading Matchers users: %v\n", err)
		os.Exit(1)
		return
	}
	rpt.Matchers = matchers

	oktaUsers, err := loadOktaUsers(ctx)
	if err != nil {
		fmt.Printf("error loading Okta users: %v\n", err)
		os.Exit(1)
		return
	}
	rpt.OktaUsers = oktaUsers
	rpt.OktaUserCount = len(oktaUsers)

	permissions, err := loadAccountGrants(ctx, client)
	if err != nil {
		fmt.Printf("error loading permissions: %v\n", err)
		os.Exit(1)
		return
	}

	users, err := loadAccounts(ctx, client)
	if err != nil {
		fmt.Printf("error loading users: %v\n", err)
		os.Exit(1)
		return
	}
	rpt.SDMUsers = users
	rpt.SDMUserCount = len(users)

	resources, err := loadResources(ctx, client)
	if err != nil {
		fmt.Printf("error loading datasources: %v\n", err)
		os.Exit(1)
		return
	}
	rpt.SDMResources = resources
	rpt.SDMResourcesCount = len(resources)

	// determine set of datasources and servers they should have access to by group
	entitlements, err := matchEntitlements(ctx, client, matchers)
	if err != nil {
		fmt.Printf("error matching entitlements: %v\n", err)
		os.Exit(1)
		return
	}

	// intended entitlements per okta group matching for each user defined in strongDM
	matchingByUser := make(map[userRow]entitlementList)
	bothCount := make(map[string]bool)
	for _, sdmUser := range users {
		uniq := make(map[entitlement]bool)
		for _, oktaUser := range oktaUsers {
			if strings.ToLower(sdmUser.Email) == strings.ToLower(oktaUser.Login) {
				bothCount[sdmUser.Email] = true
				for _, g := range oktaUser.Groups {
					for _, e := range entitlements[g] {
						uniq[e] = true
					}
				}
			}
		}
		for e := range uniq {
			matchingByUser[sdmUser] = append(matchingByUser[sdmUser], e)
		}
	}
	rpt.BothUserCount = len(bothCount)

	// determine implied grants and revocations
	toGrant := []permissionsRow{}
	toRevoke := []permissionsRow{}
	for u, entitlements := range matchingByUser {
		// are there any entitlements not permitted? grant.
		for _, e := range entitlements {
			found := false
			for _, p := range permissions {
				if p.UserID == u.ID && p.DatasourceID == e.DatasourceID {
					found = true
				}
			}
			if !found {
				if !*plan {
					toGrant = append(toGrant, permissionsRow{UserID: u.ID, DatasourceID: e.DatasourceID})
					rpt.PermissionsGranted++
				} else {
					fmt.Printf("Plan: grant %v to user %v\n", e.DatasourceID, u.ID)
				}
			}
		}
	}

	rpt.Grants = []entitlement{}
	for _, g := range toGrant {
		rpt.Grants = append(rpt.Grants, entitlement{DatasourceID: g.DatasourceID})
	}

	// are there any permissions not entitled? revoke.
	for _, p := range permissions {
		found := false
		for u, entitlements := range matchingByUser {
			// skip when permissions inherited from role
			if u.Role != "" {
				continue
			}

			if p.UserID == u.ID {
				for _, e := range entitlements {
					if p.UserID == u.ID && e.DatasourceID == p.DatasourceID {
						found = true
					}
				}
			}
		}
		if !found {
			if !*plan {
				toRevoke = append(toRevoke, p)
				rpt.PermissionsRevoked++
			} else {
				fmt.Printf("Plan: revoke %s from user %s\n", p.DatasourceID, p.UserID)
			}
		}
	}

	rpt.Revocations = toRevoke

	if !*plan {
		for _, grant := range toGrant {
			_, err := client.AccountGrants().Create(ctx, &sdm.AccountGrant{
				AccountID:  grant.UserID,
				ResourceID: grant.DatasourceID,
			})
			var alreadyExistsErr *sdm.AlreadyExistsError
			if err != nil && !errors.As(err, &alreadyExistsErr) {
				fmt.Println("error granting", err)
			}
		}
		for _, grant := range toRevoke {
			_, err := client.AccountGrants().Delete(ctx, grant.ID)
			var notFoundError *sdm.NotFoundError
			if err != nil && !errors.As(err, &notFoundError) {
				fmt.Println("error revoking", err)
			}
		}
	}

	rpt.Complete = time.Now()
	fmt.Println(rpt.String())
}

type oktaUser struct {
	Login     string   `json:"login"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Groups    []string `json:"groups"`
}

type userList []userRow

type userRow struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Role      string `json:"roleName"`
}

type permissionsList []permissionsRow

type permissionsRow struct {
	ID           string `json:"-"`
	UserID       string `json:"userID"`
	DatasourceID string `json:"datasourceID"`
}

type entitlementList []entitlement

type entitlement struct {
	DatasourceID string `json:"id"`
	Name         string `json:"name"`
}

func loadOktaUsers(ctx context.Context) ([]oktaUser, error) {
	client, err := okta.NewClient(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "invalid Okta configuration")
	}
	search := query.NewQueryParams(query.WithSearch("profile.department eq \"Engineering\" and (status eq \"ACTIVE\")"))

	apiUsers, _, err := client.User.ListUsers(search)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve okta users")
	}

	var users []oktaUser
	for _, u := range apiUsers {
		login := (*u.Profile)["login"].(string)

		groups, _, err := client.User.ListUserGroups(u.Id, nil)
		if err != nil {
			return nil, errors.Wrap(err, "unable to retrieve okta user groups")
		}

		var groupNames []string
		for _, g := range groups {
			groupNames = append(groupNames, g.Profile.Name)
		}

		var u oktaUser
		u.Login = login
		u.Groups = groupNames
		users = append(users, u)
	}
	return users, nil
}

func loadAccountGrants(ctx context.Context, client *sdm.Client) ([]permissionsRow, error) {
	grants, err := client.AccountGrants().List(ctx, "")
	if err != nil {
		return nil, err
	}
	var result permissionsList
	for grants.Next() {
		grant := grants.Value()
		result = append(result, permissionsRow{
			ID:           grant.ID,
			UserID:       grant.AccountID,
			DatasourceID: grant.ResourceID,
		})
	}
	if grants.Err() != nil {
		return nil, grants.Err()
	}
	return result, nil
}

func loadAccounts(ctx context.Context, client *sdm.Client) ([]userRow, error) {
	accountAttachments, err := client.AccountAttachments().List(ctx, "")
	if err != nil {
		return nil, err
	}
	roles := map[string]string{}
	for accountAttachments.Next() {
		attachment := accountAttachments.Value()
		roles[attachment.AccountID] = attachment.RoleID
	}
	if accountAttachments.Err() != nil {
		return nil, accountAttachments.Err()
	}

	accounts, err := client.Accounts().List(ctx, "type:user")
	if err != nil {
		return nil, err
	}
	var result userList
	for accounts.Next() {
		user := accounts.Value().(*sdm.User)
		result = append(result, userRow{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      roles[user.ID],
		})
	}
	if accounts.Err() != nil {
		return nil, accounts.Err()
	}
	return result, nil
}

func loadResources(ctx context.Context, client *sdm.Client) ([]entitlement, error) {
	// limit grant/revoke to datasources and servers only, allowing websites
	// to be granted manually for the time being
	var resources entitlementList
	resp, err := client.Resources().List(ctx, "category:datasource")
	if err != nil {
		return nil, err
	}
	for resp.Next() {
		resource := resp.Value()
		resources = append(resources, entitlement{
			DatasourceID: resource.GetID(),
			Name:         resource.GetName(),
		})
	}
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	resp, err = client.Resources().List(ctx, "category:server")
	if err != nil {
		return nil, err
	}
	for resp.Next() {
		resource := resp.Value()
		resources = append(resources, entitlement{
			DatasourceID: resource.GetID(),
			Name:         resource.GetName(),
		})
	}
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	return resources, nil
}

type MatcherConfig struct {
	Groups []struct {
		Name      string   `yaml:"name"`
		Resources []string `yaml:"resources"`
	} `yaml:"groups"`
}

func loadMatchers() (*MatcherConfig, error) {
	body, err := ioutil.ReadFile("matchers.yml")
	if err != nil {
		return nil, errors.Wrap(err, "unable to read from matchers configuration file")
	}

	var m MatcherConfig
	err = yaml.UnmarshalStrict(body, &m)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling matcher configuration")
	}

	return &m, err
}

// matchEntitlements creates lists of concrete datasources and servers by group name
func matchEntitlements(ctx context.Context, client *sdm.Client, matchers *MatcherConfig) (map[string]entitlementList, error) {
	result := make(map[string]entitlementList)
	for _, matcher := range matchers.Groups {
		uniq := make(map[entitlement]bool)
		for _, expression := range matcher.Resources {
			resources, err := client.Resources().List(ctx, expression)
			if err != nil {
				return nil, err
			}
			for resources.Next() {
				rs := resources.Value()
				uniq[entitlement{DatasourceID: rs.GetID()}] = true
			}
			if resources.Err() != nil {
				return nil, err
			}
		}

		for u := range uniq {
			result[matcher.Name] = append(result[matcher.Name], u)
		}
	}
	return result, nil
}
