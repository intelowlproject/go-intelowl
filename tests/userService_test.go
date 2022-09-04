package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

<<<<<<< HEAD
=======
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
>>>>>>> 257680317d89984491a227ed5cbef1defea3dec9
	"github.com/intelowlproject/go-intelowl/gointelowl"
)

func TestUserServiceAccess(t *testing.T) {
	userStringJson := `{"user":{"username":"hussain","first_name":"h","last_name":"k","full_name":"h k","email":"mshk9991@gmail.com"},"access":{"total_submissions":38,"month_submissions":28}}`
<<<<<<< HEAD
	userResponse := &gointelowl.User{}
	if unmarshalError := json.Unmarshal([]byte(userStringJson), &userResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}
	// table test case
=======
	userResponse := gointelowl.User{}
	if unmarshalError := json.Unmarshal([]byte(userStringJson), &userResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}
	// *table test case
>>>>>>> 257680317d89984491a227ed5cbef1defea3dec9
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      nil,
		Data:       userStringJson,
		StatusCode: http.StatusOK,
		Want:       userResponse,
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
<<<<<<< HEAD
			client, apiHandler, closeServer := setup()
			defer closeServer()
			apiHandler.Handle("/api/me/access", serverHandler(t, testCase, "GET"))
			ctx := context.Background()
			gottenUserResponse, err := client.UserService.Access(ctx)
			if err != nil {
				testError(t, testCase, err)
			} else {
				testWantData(t, testCase.Want, gottenUserResponse)
			}
=======
			testServer := NewTestServer(&testCase)
			defer testServer.Close()
			client := NewTestIntelOwlClient(testServer.URL)
			ctx := context.Background()
			gottenUserResponse, err := client.UserService.Access(ctx)
			if testCase.StatusCode < http.StatusOK || testCase.StatusCode >= http.StatusBadRequest {
				diff := cmp.Diff(testCase.Want, err, cmpopts.IgnoreFields(gointelowl.IntelOwlError{}, "Response"))
				if diff != "" {
					t.Fatalf(diff)
				}
			} else {
				diff := cmp.Diff(testCase.Want, (*gottenUserResponse))
				if diff != "" {
					t.Fatalf(diff)
				}
			}

>>>>>>> 257680317d89984491a227ed5cbef1defea3dec9
		})
	}
}

func TestUserServiceOrganization(t *testing.T) {
	orgRespJsonStr := `{"members_count":1,"owner":{"username":"hussain","full_name":"h k","joined":"2022-07-23T09:11:08.674294Z"},"is_user_owner":true,"created_at":"2022-07-23T09:11:08.580533Z","name":"StrawHats"}`
<<<<<<< HEAD
	orgResponse := &gointelowl.Organization{}
=======
	orgResponse := gointelowl.Organization{}
>>>>>>> 257680317d89984491a227ed5cbef1defea3dec9
	if unmarshalError := json.Unmarshal([]byte(orgRespJsonStr), &orgResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}
	// *table test case
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      nil,
		Data:       orgRespJsonStr,
		StatusCode: http.StatusOK,
		Want:       orgResponse,
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
<<<<<<< HEAD
			client, apiHandler, closeServer := setup()
			defer closeServer()
			apiHandler.Handle("/api/me/organization", serverHandler(t, testCase, "GET"))
			ctx := context.Background()
			gottenUserResponse, err := client.UserService.Organization(ctx)
			if err != nil {
				testError(t, testCase, err)
			} else {
				testWantData(t, testCase.Want, gottenUserResponse)
=======
			testServer := NewTestServer(&testCase)
			defer testServer.Close()
			client := NewTestIntelOwlClient(testServer.URL)
			ctx := context.Background()
			gottenOrgResponse, err := client.UserService.Organization(ctx)
			if testCase.StatusCode < http.StatusOK || testCase.StatusCode >= http.StatusBadRequest {
				diff := cmp.Diff(testCase.Want, err, cmpopts.IgnoreFields(gointelowl.IntelOwlError{}, "Response"))
				if diff != "" {
					t.Fatalf(diff)
				}
			} else {
				diff := cmp.Diff(testCase.Want, (*gottenOrgResponse))
				if diff != "" {
					t.Fatalf(diff)
				}
>>>>>>> 257680317d89984491a227ed5cbef1defea3dec9
			}
		})
	}
}

func TestUserServiceCreateOrganization(t *testing.T) {
	orgRespJsonStr := `{"members_count":1,"owner":{"username":"notHussain","full_name":"noy Hussain","joined":"2022-07-24T17:34:55.032629Z"},"is_user_owner":true,"created_at":"2022-07-24T17:34:54.971735Z","name":"TestOrganization"}`
<<<<<<< HEAD
	orgResponse := &gointelowl.Organization{}
=======
	orgResponse := gointelowl.Organization{}
>>>>>>> 257680317d89984491a227ed5cbef1defea3dec9
	if unmarshalError := json.Unmarshal([]byte(orgRespJsonStr), &orgResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}
	orgParams := gointelowl.OrganizationParams{
		Name: "TestOrganization",
	}
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      orgParams,
		Data:       orgRespJsonStr,
		StatusCode: http.StatusOK,
		Want:       orgResponse,
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
<<<<<<< HEAD
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			apiHandler.Handle("/api/me/organization", serverHandler(t, testCase, "POST"))
			params, ok := testCase.Input.(gointelowl.OrganizationParams)
			if ok {
				gottenOrgResponse, err := client.UserService.CreateOrganization(ctx, &params)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, gottenOrgResponse)
				}
=======
			testServer := NewTestServer(&testCase)
			defer testServer.Close()
			client := NewTestIntelOwlClient(testServer.URL)
			ctx := context.Background()
			params, ok := testCase.Input.(gointelowl.OrganizationParams)
			if ok {
				gottenOrgResponse, err := client.UserService.CreateOrganization(ctx, &params)
				if testCase.StatusCode < http.StatusOK || testCase.StatusCode >= http.StatusBadRequest {
					diff := cmp.Diff(testCase.Want, err, cmpopts.IgnoreFields(gointelowl.IntelOwlError{}, "Response"))
					if diff != "" {
						t.Fatalf(diff)
					}
				} else {
					diff := cmp.Diff(testCase.Want, (*gottenOrgResponse))
					if diff != "" {
						t.Fatalf(diff)
					}
				}
			} else {
				t.Fatalf("Casting failed!")
>>>>>>> 257680317d89984491a227ed5cbef1defea3dec9
			}
		})
	}
}

func TestUserServiceRemoveMemberFromOrganization(t *testing.T) {
	memberParams := gointelowl.MemberParams{
		Username: "TestUser",
	}
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      memberParams,
		Data:       ``,
		StatusCode: http.StatusNoContent,
		Want:       true,
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
<<<<<<< HEAD
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			apiHandler.Handle("/api/me/organization/remove_member", serverHandler(t, testCase, "POST"))
			params, ok := testCase.Input.(gointelowl.MemberParams)
			if ok {
				left, err := client.UserService.RemoveMemberFromOrganization(ctx, &params)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, left)
				}
=======
			testServer := NewTestServer(&testCase)
			defer testServer.Close()
			client := NewTestIntelOwlClient(testServer.URL)
			ctx := context.Background()
			params, ok := testCase.Input.(gointelowl.MemberParams)
			if ok {
				left, err := client.UserService.RemoveMemberFromOrganization(ctx, &params)
				if testCase.StatusCode < http.StatusOK || testCase.StatusCode >= http.StatusBadRequest {
					diff := cmp.Diff(testCase.Want, err, cmpopts.IgnoreFields(gointelowl.IntelOwlError{}, "Response"))
					if diff != "" {
						t.Fatalf(diff)
					}
				} else {
					diff := cmp.Diff(testCase.Want, left)
					if diff != "" {
						t.Fatalf(diff)
					}
				}
			} else {
				t.Fatalf("Casting failed!")
>>>>>>> 257680317d89984491a227ed5cbef1defea3dec9
			}
		})
	}
}

func TestUserServiceInviteToOrganization(t *testing.T) {
	inviteJsonStr := `{"id":12,"created_at":"2022-07-24T18:43:42.299318Z","status":"pending"}`
<<<<<<< HEAD
	inviteResponse := &gointelowl.Invite{}
=======
	inviteResponse := gointelowl.Invite{}
>>>>>>> 257680317d89984491a227ed5cbef1defea3dec9
	if unmarshalError := json.Unmarshal([]byte(inviteJsonStr), &inviteResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}
	memberParams := gointelowl.MemberParams{
		Username: "TestUser",
	}
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      memberParams,
		Data:       inviteJsonStr,
		StatusCode: http.StatusCreated,
		Want:       inviteResponse,
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
<<<<<<< HEAD
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			apiHandler.Handle("/api/me/organization/invite", serverHandler(t, testCase, "POST"))
			params, ok := testCase.Input.(gointelowl.MemberParams)
			if ok {
				gottenInviteResponse, err := client.UserService.InviteToOrganization(ctx, &params)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, gottenInviteResponse)
				}
=======
			testServer := NewTestServer(&testCase)
			defer testServer.Close()
			client := NewTestIntelOwlClient(testServer.URL)
			ctx := context.Background()
			params, ok := testCase.Input.(gointelowl.MemberParams)
			if ok {
				gottenInviteResponse, err := client.UserService.InviteToOrganization(ctx, &params)
				if testCase.StatusCode < http.StatusOK || testCase.StatusCode >= http.StatusBadRequest {
					diff := cmp.Diff(testCase.Want, err, cmpopts.IgnoreFields(gointelowl.IntelOwlError{}, "Response"))
					if diff != "" {
						t.Fatalf(diff)
					}
				} else {
					diff := cmp.Diff(testCase.Want, (*gottenInviteResponse))
					if diff != "" {
						t.Fatalf(diff)
					}
				}
			} else {
				t.Fatalf("Casting failed!")
>>>>>>> 257680317d89984491a227ed5cbef1defea3dec9
			}
		})
	}
}
