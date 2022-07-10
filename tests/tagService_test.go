package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/intelowlproject/go-intelowl/gointelowl"
)

type TagTest struct {
	input      interface{}
	data       string
	statusCode int
	want       interface{}
}

// * Test for TagService.List method!
func TestTagServiceList(t *testing.T) {
	// * table test case
	testCases := make(map[string]TagTest)
	testCases["simple"] = TagTest{
		input: nil,
		data: `[
			{"id": 1,"label": "TEST1","color": "#1c71d8"},
			{"id": 2,"label": "TEST2","color": "#1c71d7"},
			{"id": 3,"label": "TEST3","color": "#1c71d6"}
		]`,
		statusCode: http.StatusOK,
		want: []gointelowl.Tag{
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
		//* Subtest
		t.Run(name, func(t *testing.T) {
			testData := TestData{
				StatusCode: testCase.statusCode,
				Data:       testCase.data,
			}
			testServer := MakeNewTestServer(&testData)
			defer testServer.Close()
			client := MakeNewTestIntelOwlClient(testServer.URL)
			ctx := context.Background()
			gottenTagList, err := client.TagService.List(ctx)
			if err != nil {
				t.Fatalf("Error listing tags: %v", err)
			}
			diff := cmp.Diff(testCase.want, (*gottenTagList))
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestTagServiceGet(t *testing.T) {
	// *table test case
	testCases := make(map[string]TagTest)
	testCases["simple"] = TagTest{
		input:      1,
		data:       `{"id": 1,"label": "TEST","color": "#1c71d8"}`,
		statusCode: http.StatusOK,
		want: &gointelowl.Tag{
			ID:    1,
			Label: "TEST",
			Color: "#1c71d8",
		},
	}
	testCases["cantFind"] = TagTest{
		input:      9000,
		data:       `{"detail": "Not found."}`,
		statusCode: http.StatusNotFound,
		want: &gointelowl.IntelOwlError{
			StatusCode: http.StatusNotFound,
			Data:       []byte(`{"detail": "Not found."}`),
		},
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			testData := TestData{
				StatusCode: testCase.statusCode,
				Data:       testCase.data,
			}
			testServer := MakeNewTestServer(&testData)
			defer testServer.Close()
			client := MakeNewTestIntelOwlClient(testServer.URL)
			ctx := context.Background()
			id, ok := testCase.input.(int)
			if ok {
				tagId := uint64(id)
				gottenTag, err := client.TagService.Get(ctx, tagId)
				if testCase.statusCode < http.StatusOK || testCase.statusCode >= http.StatusBadRequest {
					diff := cmp.Diff(testCase.want, err, cmpopts.IgnoreFields(gointelowl.IntelOwlError{}, "Response"))
					if diff != "" {
						t.Fatalf(diff)
					}
				} else {
					diff := cmp.Diff(testCase.want, gottenTag)
					if diff != "" {
						t.Fatalf(diff)
					}
				}
			} else {
				t.Fatalf("Casting failed!")
			}
		})
	}
}

func TestTagServiceCreate(t *testing.T) {
	// *table test case
	testCases := make(map[string]TagTest)
	testCases["simple"] = TagTest{
		input: gointelowl.TagParams{
			Label: "TEST TAG",
			Color: "#fffff",
		},
		data:       `{"id": 1,"label": "TEST TAG","color": "#fffff"}`,
		statusCode: http.StatusOK,
		want: &gointelowl.Tag{
			ID:    1,
			Label: "TEST TAG",
			Color: "#fffff",
		},
	}
	testCases["duplicate"] = TagTest{
		input: gointelowl.TagParams{
			Label: "TEST TAG",
			Color: "#fffff",
		},
		data:       `{"label":["tag with this label already exists."]}`,
		statusCode: http.StatusBadRequest,
		want: &gointelowl.IntelOwlError{
			StatusCode: http.StatusBadRequest,
			Data:       []byte(`{"label":["tag with this label already exists."]}`),
		},
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			testData := &TestData{
				StatusCode: testCase.statusCode,
				Data:       testCase.data,
			}
			testServer := MakeNewTestServer(testData)
			defer testServer.Close()
			client := MakeNewTestIntelOwlClient(testServer.URL)
			ctx := context.Background()
			tagParams, ok := testCase.input.(gointelowl.TagParams)
			if ok {
				gottenTag, err := client.TagService.Create(ctx, &tagParams)
				if testCase.statusCode < http.StatusOK || testCase.statusCode >= http.StatusBadRequest {
					diff := cmp.Diff(testCase.want, err, cmpopts.IgnoreFields(gointelowl.IntelOwlError{}, "Response"))
					if diff != "" {
						t.Fatalf(diff)
					}
				} else {
					diff := cmp.Diff(testCase.want, gottenTag)
					if diff != "" {
						t.Fatalf(diff)
					}
				}
			} else {
				t.Fatalf("Casting failed!")
			}
		})
	}
}

func TestTagServiceUpdate(t *testing.T) {
	// *table test case
	testCases := make(map[string]TagTest)
	testCases["simple"] = TagTest{
		input: gointelowl.Tag{
			ID:    1,
			Label: "UPDATED TEST TAG",
			Color: "#f4",
		},
		data:       `{"id": 1,"label": "UPDATED TEST TAG","color": "#f4"}`,
		statusCode: http.StatusOK,
		want: &gointelowl.Tag{
			ID:    1,
			Label: "UPDATED TEST TAG",
			Color: "#f4",
		},
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			testData := &TestData{
				StatusCode: testCase.statusCode,
				Data:       testCase.data,
			}
			testServer := MakeNewTestServer(testData)
			defer testServer.Close()
			client := MakeNewTestIntelOwlClient(testServer.URL)
			ctx := context.Background()
			tag, ok := testCase.input.(gointelowl.Tag)
			if ok {
				gottenTag, err := client.TagService.Update(ctx, tag.ID, &gointelowl.TagParams{
					Label: tag.Label,
					Color: tag.Color,
				})
				if testCase.statusCode < http.StatusOK || testCase.statusCode >= http.StatusBadRequest {
					diff := cmp.Diff(testCase.want, err, cmpopts.IgnoreFields(gointelowl.IntelOwlError{}, "Response"))
					if diff != "" {
						t.Fatalf(diff)
					}
				} else {
					diff := cmp.Diff(testCase.want, gottenTag)
					if diff != "" {
						t.Fatalf(diff)
					}
				}
			}
		})
	}
}

func TestTagServiceDelete(t *testing.T) {
	// *table test case
	testCases := make(map[string]TagTest)
	testCases["simple"] = TagTest{
		input:      1,
		data:       "",
		statusCode: http.StatusNoContent,
		want:       true,
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			testData := &TestData{
				StatusCode: testCase.statusCode,
				Data:       testCase.data,
			}
			testServer := MakeNewTestServer(testData)
			defer testServer.Close()
			client := MakeNewTestIntelOwlClient(testServer.URL)
			ctx := context.Background()
			id, ok := testCase.input.(int)
			if ok {
				tagId := uint64(id)
				isDeleted, err := client.TagService.Delete(ctx, tagId)
				if testCase.statusCode < http.StatusOK || testCase.statusCode >= http.StatusBadRequest {
					diff := cmp.Diff(testCase.want, err, cmpopts.IgnoreFields(gointelowl.IntelOwlError{}, "Response"))
					if diff != "" {
						t.Fatalf(diff)
					}
				} else {
					diff := cmp.Diff(testCase.want, isDeleted)
					if diff != "" {
						t.Fatalf(diff)
					}
				}
			}
		})
	}
}
