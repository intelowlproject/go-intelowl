package gointelowl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TagParams struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type Tag struct {
	client *IntelOwlClient
	ID     uint64 `json:"id"`
	Label  string `json:"label"`
	Color  string `json:"color"`
}

// Stubs for implementing the error interface
type TagExistError struct {
	Label []string `json:"label"`
}

type TagAttributeRequiredError struct {
	Label []string `json:"label,omitempty"`
	Color []string `json:"color,omitempty"`
}

// * getting on all tags
func (tag *Tag) List(ctx context.Context) (*[]Tag, error) {
	requestUrl := fmt.Sprintf("%s/api/tags", tag.client.options.Url)
	fmt.Println(requestUrl)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	tagResponse := []Tag{}
	if err := tag.client.makeRequest(ctx, request, &tagResponse); err != nil {
		return nil, err
	}
	return &tagResponse, nil
}

// * Getting a tag through its ID
func (tag *Tag) Get(ctx context.Context, tagId uint64) (*Tag, error) {
	requestUrl := fmt.Sprintf("%s/api/tags/%d", tag.client.options.Url, tagId)
	fmt.Println(requestUrl)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	tagResponse := Tag{}
	if err := tag.client.makeRequest(ctx, request, &tagResponse); err != nil {
		return nil, err
	}
	return &tagResponse, nil
}

//* Creating a Tag
func (tag *Tag) Create(ctx context.Context, tagParams *TagParams) (*Tag, error) {
	requestUrl := fmt.Sprintf("%s/api/tags", tag.client.options.Url)
	fmt.Println("Url: " + requestUrl)
	tagJson, err := json.Marshal(tagParams)
	fmt.Println(string(tagJson))
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(tagJson))
	if err != nil {
		return nil, err
	}
	createdTag := Tag{}
	if err := tag.client.makeRequest(ctx, request, &createdTag); err != nil {
		return nil, err
	}
	return &createdTag, nil
}

//* Updating a tag
func (tag *Tag) Update(ctx context.Context, tagId uint64, tagParams *TagParams) (*Tag, error) {
	requestUrl := fmt.Sprintf("%s/api/tags/%d", tag.client.options.Url, tagId)
	// printing the request
	fmt.Println("Url: " + requestUrl)

	// Getting the relevant JSON data
	tagJson, err := json.Marshal(tagParams)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("PUT", requestUrl, bytes.NewBuffer(tagJson))
	if err != nil {
		return nil, err
	}
	updatedTag := Tag{}
	if err := tag.client.makeRequest(ctx, request, &updatedTag); err != nil {
		return nil, err
	}
	return &updatedTag, nil
}

//* Deleting a tag
func (tag *Tag) Delete(ctx context.Context, tagId uint64) (bool, error) {
	requestUrl := fmt.Sprintf("%s/api/tags/%d", tag.client.options.Url, tagId)
	// printing the request
	fmt.Println("Url: " + requestUrl)
	request, err := http.NewRequest("DELETE", requestUrl, nil)
	if err != nil {
		return false, err
	}
	var data interface{}
	if err := tag.client.makeRequest(ctx, request, &data); err != nil {
		return false, err
	}
	if data == nil {
		fmt.Println("FUCK ME DADDY")
		return true, nil
	}
	return false, nil
}

//* Pretty printing the tag
func (tag *Tag) Display() error {
	data, err := json.MarshalIndent(tag, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
