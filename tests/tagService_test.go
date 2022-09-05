package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/intelowlproject/go-intelowl/constants"
	"github.com/intelowlproject/go-intelowl/gointelowl"
)

// * Test for TagService.List method!
func TestTagServiceList(t *testing.T) {

	testCases := make(map[string]TestData)

	testCases["simple"] = TestData{
		Input: nil,
		Data: `[
			{"id": 1,"label": "TEST1","color": "#1c71d8"},
			{"id": 2,"label": "TEST2","color": "#1c71d7"},
			{"id": 3,"label": "TEST3","color": "#1c71d6"}
		]`,
		StatusCode: http.StatusOK,
		Want: []gointelowl.Tag{
			{
				ID:    1,
				Label: "TEST1",
				Color: "#1c71d8",
			},
			{
				ID:    2,
				Label: "TEST2",
				Color: "#1c71d7",
			},
			{
				ID:    3,
				Label: "TEST3",
				Color: "#1c71d6",
			},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			apiHandler.Handle(constants.BASE_TAG_URL, serverHandler(t, testCase, "GET"))
			ctx := context.Background()
			gottenTagList, err := client.TagService.List(ctx)
			if err != nil {
				testError(t, testCase, err)
			} else {
				testWantData(t, testCase.Want, (*gottenTagList))
			}
		})
	}
}

func TestTagServiceGet(t *testing.T) {
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      1,
		StatusCode: http.StatusOK,
		Data:       `{"id": 1,"label": "TEST","color": "#1c71d8"}`,
		Want: &gointelowl.Tag{
			ID:    1,
			Label: "TEST",
			Color: "#1c71d8",
		},
	}
	testCases["cantFind"] = TestData{
		Input:      9000,
		StatusCode: http.StatusNotFound,
		Data:       `{"detail": "Not found."}`,
		Want: &gointelowl.IntelOwlError{
			StatusCode: http.StatusNotFound,
			Message:    `{"detail": "Not found."}`,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			id, ok := testCase.Input.(int)
			if ok {
				tagId := uint64(id)
				testUrl := fmt.Sprintf(constants.SPECIFIC_TAG_URL, tagId)
				apiHandler.Handle(testUrl, serverHandler(t, testCase, "GET"))
				gottenTag, err := client.TagService.Get(ctx, tagId)
				// Helper test to check error
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, gottenTag)
				}
			} else {
				t.Fatalf("Casting failed!")
				t.Fatalf("You didn't pass the correct Input datatype")
			}
		})
	}
}

func TestTagServiceCreate(t *testing.T) {
	// *table test case
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input: gointelowl.TagParams{
			Label: "TEST TAG",
			Color: "#fffff",
		},
		Data:       `{"id": 1,"label": "TEST TAG","color": "#fffff"}`,
		StatusCode: http.StatusOK,
		Want: &gointelowl.Tag{
			ID:    1,
			Label: "TEST TAG",
			Color: "#fffff",
		},
	}
	testCases["duplicate"] = TestData{
		Input: gointelowl.TagParams{
			Label: "TEST TAG",
			Color: "#fffff",
		},
		Data:       `{"label":["tag with this label already exists."]}`,
		StatusCode: http.StatusBadRequest,
		Want: &gointelowl.IntelOwlError{
			StatusCode: http.StatusBadRequest,
			Message:    `{"label":["tag with this label already exists."]}`,
		},
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			apiHandler.Handle(constants.BASE_TAG_URL, serverHandler(t, testCase, "POST"))
			tagParams, ok := testCase.Input.(gointelowl.TagParams)
			if ok {
				gottenTag, err := client.TagService.Create(ctx, &tagParams)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, gottenTag)
				}
			} else {
				t.Fatalf("Casting failed!")
			}
		})
	}
}

func TestTagServiceUpdate(t *testing.T) {
	// *table test case
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input: gointelowl.Tag{
			ID:    1,
			Label: "UPDATED TEST TAG",
			Color: "#f4",
		},
		Data:       `{"id": 1,"label": "UPDATED TEST TAG","color": "#f4"}`,
		StatusCode: http.StatusOK,
		Want: &gointelowl.Tag{
			ID:    1,
			Label: "UPDATED TEST TAG",
			Color: "#f4",
		},
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			tag, ok := testCase.Input.(gointelowl.Tag)
			if ok {
				testUrl := fmt.Sprintf(constants.SPECIFIC_TAG_URL, tag.ID)
				apiHandler.Handle(testUrl, serverHandler(t, testCase, "PUT"))
				gottenTag, err := client.TagService.Update(ctx, tag.ID, &gointelowl.TagParams{
					Label: tag.Label,
					Color: tag.Color,
				})
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, gottenTag)
				}
			}
		})
	}
}

func TestTagServiceDelete(t *testing.T) {
	// *table test case
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      1,
		Data:       "",
		StatusCode: http.StatusNoContent,
		Want:       true,
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			id, ok := testCase.Input.(int)
			if ok {
				tagId := uint64(id)
				testUrl := fmt.Sprintf(constants.SPECIFIC_TAG_URL, tagId)
				apiHandler.Handle(testUrl, serverHandler(t, testCase, "DELETE"))
				isDeleted, err := client.TagService.Delete(ctx, tagId)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, isDeleted)
				}
			}
		})
	}
}
