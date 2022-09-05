package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/intelowlproject/go-intelowl/constants"
	"github.com/intelowlproject/go-intelowl/gointelowl"
)

func TestUserServiceAccess(t *testing.T) {
	userStringJson := `{"user":{"username":"hussain","first_name":"h","last_name":"k","full_name":"h k","email":"mshk9991@gmail.com"},"access":{"total_submissions":38,"month_submissions":28}}`
	userResponse := &gointelowl.User{}
	if unmarshalError := json.Unmarshal([]byte(userStringJson), &userResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}
	// table test case
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      nil,
		Data:       userStringJson,
		StatusCode: http.StatusOK,
		Want:       userResponse,
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			apiHandler.Handle(constants.USER_DETAILS_URL, serverHandler(t, testCase, "GET"))
			ctx := context.Background()
			gottenUserResponse, err := client.UserService.Access(ctx)
			if err != nil {
				testError(t, testCase, err)
			} else {
				testWantData(t, testCase.Want, gottenUserResponse)
			}
		})
	}
}

func TestUserServiceOrganization(t *testing.T) {
	orgRespJsonStr := `{"members_count":1,"owner":{"username":"hussain","full_name":"h k","joined":"2022-07-23T09:11:08.674294Z"},"is_user_owner":true,"created_at":"2022-07-23T09:11:08.580533Z","name":"StrawHats"}`
	orgResponse := &gointelowl.Organization{}
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
			client, apiHandler, closeServer := setup()
			defer closeServer()
			apiHandler.Handle(constants.ORGANIZATION_URL, serverHandler(t, testCase, "GET"))
			ctx := context.Background()
			gottenUserResponse, err := client.UserService.Organization(ctx)
			if err != nil {
				testError(t, testCase, err)
			} else {
				testWantData(t, testCase.Want, gottenUserResponse)
			}
		})
	}
}

func TestUserServiceCreateOrganization(t *testing.T) {
	orgRespJsonStr := `{"members_count":1,"owner":{"username":"notHussain","full_name":"noy Hussain","joined":"2022-07-24T17:34:55.032629Z"},"is_user_owner":true,"created_at":"2022-07-24T17:34:54.971735Z","name":"TestOrganization"}`
	orgResponse := &gointelowl.Organization{}
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
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			apiHandler.Handle(constants.ORGANIZATION_URL, serverHandler(t, testCase, "POST"))
			params, ok := testCase.Input.(gointelowl.OrganizationParams)
			if ok {
				gottenOrgResponse, err := client.UserService.CreateOrganization(ctx, &params)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, gottenOrgResponse)
				}
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
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			apiHandler.Handle(constants.REMOVE_MEMBER_FROM_ORGANIZATION_URL, serverHandler(t, testCase, "POST"))
			params, ok := testCase.Input.(gointelowl.MemberParams)
			if ok {
				left, err := client.UserService.RemoveMemberFromOrganization(ctx, &params)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, left)
				}
			}
		})
	}
}

func TestUserServiceInviteToOrganization(t *testing.T) {
	inviteJsonStr := `{"id":12,"created_at":"2022-07-24T18:43:42.299318Z","status":"pending"}`
	inviteResponse := &gointelowl.Invite{}
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
			}
		})
	}
}
