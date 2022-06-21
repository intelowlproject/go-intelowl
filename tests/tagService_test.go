package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/intelowlproject/go-intelowl/gointelowl"
)

// * Test for TagService.List method!
func TestTagServiceList(t *testing.T) {
	// * table test case
	testCases := map[string]struct {
		input string
		want  []gointelowl.Tag
	}{
		"simple": {
			input: `[
				{"id": 1,"label": "TEST1","color": "#1c71d8"},
				{"id": 2,"label": "TEST2","color": "#1c71d7"},
				{"id": 3,"label": "TEST3","color": "#1c71d6"}
			]`,
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
		},
		"empty": {
			input: `[]`,
			want:  []gointelowl.Tag{},
		},
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(testCase.input))
			}))
			defer testServer.Close()
			client := gointelowl.MakeNewIntelOwlClient(
				&gointelowl.IntelOwlClientOptions{
					Url:         testServer.URL,
					Token:       "test-token",
					Certificate: "",
				},
				nil,
			)
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
	testCases := map[string]struct {
		input      uint64
		data       string
		statusCode int
		want       interface{}
	}{
		"simple": {
			input:      1,
			data:       `{"id": 1,"label": "TEST","color": "#1c71d8"}`,
			statusCode: 200,
			want: &gointelowl.Tag{
				ID:    1,
				Label: "TEST",
				Color: "#1c71d8",
			},
		},
		"cantFind": {
			input:      9000,
			data:       `{"detail": "Not found."}`,
			statusCode: 404,
			want: &gointelowl.IntelOwlError{
				StatusCode: 404,
				Data:       []byte(`{"detail": "Not found."}`),
			},
		},
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(testCase.statusCode)
				w.Write([]byte(testCase.data))
			}))
			defer testServer.Close()
			client := gointelowl.MakeNewIntelOwlClient(
				&gointelowl.IntelOwlClientOptions{
					Url:         testServer.URL,
					Token:       "test-token",
					Certificate: "",
				},
				nil,
			)
			ctx := context.Background()
			gottenTag, err := client.TagService.Get(ctx, testCase.input)
			if testCase.statusCode < http.StatusOK || testCase.statusCode >= http.StatusBadRequest {
				diff := cmp.Diff(testCase.want, err)
				if diff != "" {
					t.Fatalf(diff)
				}
			} else {
				diff := cmp.Diff(testCase.want, gottenTag)
				if diff != "" {
					t.Fatalf(diff)
				}
			}
		})
	}
}

func TestTagServiceCreate(t *testing.T) {
	// *table test case
	testCases := map[string]struct {
		input      *gointelowl.TagParams
		data       string
		statusCode int
		want       interface{}
	}{
		"simple": {
			input: &gointelowl.TagParams{
				Label: "TEST TAG",
				Color: "#fffff",
			},
			data:       `{"id": 1,"label": "TEST TAG","color": "#fffff"}`,
			statusCode: 200,
			want: &gointelowl.Tag{
				ID:    1,
				Label: "TEST TAG",
				Color: "#fffff",
			},
		},
		"duplicate": {
			input: &gointelowl.TagParams{
				Label: "TEST TAG",
				Color: "#fffff",
			},
			data:       `{"label":["tag with this label already exists."]}`,
			statusCode: 400,
			want: &gointelowl.IntelOwlError{
				StatusCode: 400,
				Data:       []byte(`{"label":["tag with this label already exists."]}`),
			},
		},
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(testCase.statusCode)
				w.Write([]byte(testCase.data))
			}))
			defer testServer.Close()
			client := gointelowl.MakeNewIntelOwlClient(
				&gointelowl.IntelOwlClientOptions{
					Url:         testServer.URL,
					Token:       "test-token",
					Certificate: "",
				},
				nil,
			)
			ctx := context.Background()
			gottenTag, err := client.TagService.Create(ctx, testCase.input)
			if testCase.statusCode < http.StatusOK || testCase.statusCode >= http.StatusBadRequest {
				diff := cmp.Diff(testCase.want, err)
				if diff != "" {
					t.Fatalf(diff)
				}
			} else {
				diff := cmp.Diff(testCase.want, gottenTag)
				if diff != "" {
					t.Fatalf(diff)
				}
			}
		})
	}
}

func TestTagServiceUpdate(t *testing.T) {
	// *table test case
	testCases := map[string]struct {
		input      gointelowl.Tag
		data       string
		statusCode int
		want       interface{}
	}{
		"simple": {
			input: gointelowl.Tag{
				ID:    1,
				Label: "UPDATED TEST TAG",
				Color: "#f4",
			},
			data:       `{"id": 1,"label": "UPDATED TEST TAG","color": "#f4"}`,
			statusCode: 200,
			want: &gointelowl.Tag{
				ID:    1,
				Label: "UPDATED TEST TAG",
				Color: "#f4",
			},
		},
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(testCase.statusCode)
				w.Write([]byte(testCase.data))
			}))
			defer testServer.Close()
			client := gointelowl.MakeNewIntelOwlClient(
				&gointelowl.IntelOwlClientOptions{
					Url:         testServer.URL,
					Token:       "test-token",
					Certificate: "",
				},
				nil,
			)
			ctx := context.Background()
			gottenTag, err := client.TagService.Update(ctx, testCase.input.ID, &gointelowl.TagParams{
				Label: testCase.input.Label,
				Color: testCase.input.Color,
			})
			if testCase.statusCode < http.StatusOK || testCase.statusCode >= http.StatusBadRequest {
				diff := cmp.Diff(testCase.want, err)
				if diff != "" {
					t.Fatalf(diff)
				}
			} else {
				diff := cmp.Diff(testCase.want, gottenTag)
				if diff != "" {
					t.Fatalf(diff)
				}
			}
		})
	}
}

func TestTagServiceDelete(t *testing.T) {
	// *table test case
	testCases := map[string]struct {
		input      uint64
		data       string
		statusCode int
		want       interface{}
	}{
		"simple": {
			input:      1,
			statusCode: 204,
			want:       true,
		},
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(testCase.statusCode)
			}))
			defer testServer.Close()
			client := gointelowl.MakeNewIntelOwlClient(
				&gointelowl.IntelOwlClientOptions{
					Url:         testServer.URL,
					Token:       "test-token",
					Certificate: "",
				},
				nil,
			)
			ctx := context.Background()
			isDeleted, err := client.TagService.Delete(ctx, testCase.input)
			if testCase.statusCode < http.StatusOK || testCase.statusCode >= http.StatusBadRequest {
				diff := cmp.Diff(testCase.want, err)
				if diff != "" {
					t.Fatalf(diff)
				}
			} else {
				diff := cmp.Diff(testCase.want, isDeleted)
				if diff != "" {
					t.Fatalf(diff)
				}
			}
		})
	}
}

// func TestTagServiceGet(t *testing.T) {
// 	testServer := httptest.NewServer(http.HandlerFunc(func(w htsstp.ResponseWriter, r *http.Request) {
// 		// expected := `{
// 		// 	"id": 1,
// 		// 	"label": "TEST",
// 		// 	"color": "#1c71d8"
// 		//   }`
// 		expected := `{"id": 1,"label": "TEST","color": "#1c71d8"}`
// 		// tagJson, _ := json.Marshal(expected)
// 		// fmt.Println(tagJson)
// 		w.Write([]byte(expected))
// 	}))
// 	defer testServer.Close()
// 	client := gointelowl.MakeNewIntelOwlClient(
// 		&gointelowl.IntelOwlClientOptions{
// 			Url:         testServer.URL,
// 			Token:       "test-token",
// 			Certificate: "",
// 		},
// 		nil,
// 	)

// 	ctx := context.Background()
// 	tag, err := client.TagService.Get(ctx, 1)
// 	if err != nil {
// 		t.Errorf("Error in getting tag: %v", err)
// 	}
// 	tag.Display()
// 	if reflect.TypeOf((*tag)) != reflect.TypeOf(gointelowl.Tag{}) {
// 		t.Errorf("Expected Tag")
// 	}
// }
