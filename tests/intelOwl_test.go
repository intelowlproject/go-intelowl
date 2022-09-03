package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/intelowlproject/go-intelowl/gointelowl"
	"github.com/sirupsen/logrus"
)

// var (
// 	TestServer *httptest.Server
// 	ApiHandler *http.ServeMux
// 	TestClient gointelowl.IntelOwlClient
// )

// * Test Data Struct used for every struct obj
type TestData struct {
	Input      interface{}
	Data       string
	StatusCode int
	Want       interface{}
}

func setup() (testClient gointelowl.IntelOwlClient, apiHandler *http.ServeMux, closeServer func()) {

	apiHandler = http.NewServeMux()

	testServer := httptest.NewServer(apiHandler)

	testClient = NewTestIntelOwlClient(testServer.URL)

	return testClient, apiHandler, testServer.Close

}

func testError(t *testing.T, testData TestData, err error) {
	t.Helper()
	if testData.StatusCode < http.StatusOK || testData.StatusCode >= http.StatusBadRequest {
		diff := cmp.Diff(testData.Want, err, cmpopts.IgnoreFields(gointelowl.IntelOwlError{}, "Response"))
		if diff != "" {
			t.Fatalf(diff)
		}
	}
}

func testWantData(t *testing.T, want interface{}, data interface{}) {
	t.Helper()
	diff := cmp.Diff(want, data)
	if diff != "" {
		t.Fatalf(diff)
	}
}

func serverHandler(testData TestData) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if testData.StatusCode > 0 {
			w.WriteHeader(testData.StatusCode)
		}
		if len(testData.Data) > 0 {
			_, err := w.Write([]byte(testData.Data))
			if err != nil {
				//* writing an empty object to signifiy could not convert data!
				fmt.Fprintf(w, "{}")
			}
		}
	}

	return http.HandlerFunc(handler)
}

func NewTestServer(testData *TestData) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if testData.StatusCode > 0 {
			w.WriteHeader(testData.StatusCode)
		}
		if len(testData.Data) > 0 {
			_, err := w.Write([]byte(testData.Data))
			if err != nil {
				//* writing an empty object to signifiy could not convert data!
				fmt.Fprintf(w, "{}")
			}
		}
	}))
}

func NewTestIntelOwlClient(url string) gointelowl.IntelOwlClient {
	return gointelowl.NewIntelOwlClient(
		&gointelowl.IntelOwlClientOptions{
			Url:         url,
			Token:       "test-token",
			Certificate: "",
		},
		nil,
		&gointelowl.LoggerParams{
			File:      nil,
			Formatter: nil,
			Level:     logrus.DebugLevel,
		},
	)
}
