package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/intelowlproject/go-intelowl/gointelowl"
)

func TestUserServiceAccess(t *testing.T) {
	userStringJson := `{"user":{"username":"hussain","first_name":"h","last_name":"k","full_name":"h k","email":"mshk9991@gmail.com"},"access":{"total_submissions":38,"month_submissions":28}}`
	userResponse := gointelowl.User{}
	if unmarshalError := json.Unmarshal([]byte(userStringJson), &userResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}
	// *table test case
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      nil,
		Data:       userStringJson,
		StatusCode: http.StatusOK,
		Want:       userResponse,
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
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

		})
	}
}
