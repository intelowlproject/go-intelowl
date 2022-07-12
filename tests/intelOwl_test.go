package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/intelowlproject/go-intelowl/gointelowl"
)

// * Test Data Struct used for every struct obj
type TestData struct {
	Input      interface{}
	Data       string
	StatusCode int
	Want       interface{}
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
	)
}
