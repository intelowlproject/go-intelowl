package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/intelowlproject/go-intelowl/gointelowl"
)

type TestData struct {
	StatusCode int
	Data       string
}

func MakeNewTestServer(testData *TestData) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if testData.StatusCode > 0 {
			w.WriteHeader(testData.StatusCode)
		}
		if len(testData.Data) > 0 {
			_, err := w.Write([]byte(testData.Data))
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
		}
	}))
}

func MakeNewTestIntelOwlClient(url string) gointelowl.IntelOwlClient {
	return gointelowl.MakeNewIntelOwlClient(
		&gointelowl.IntelOwlClientOptions{
			Url:         url,
			Token:       "test-token",
			Certificate: "",
		},
		nil,
	)
}
