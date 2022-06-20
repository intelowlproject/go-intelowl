package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/intelowlproject/gointelowl/gointelowl"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestClientTagList(t *testing.T) {
	expected := []gointelowl.Tag{
		{
			ID:    1,
			Label: "test-label1",
			Color: "#fffff",
		},
	}
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tagJson, _ := json.Marshal(expected)
		w.Write(tagJson)
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
	tagList, err := client.TagService.List(ctx)
	if err != nil {
		t.Errorf("Error listing tags: %v", err)
	}
	fmt.Println(tagList)
	if reflect.TypeOf((*tagList)) != reflect.TypeOf([]gointelowl.Tag{}) {
		t.Errorf("Expected []Tag{}")
	}

}
