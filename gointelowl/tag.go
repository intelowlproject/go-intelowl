package gointelowl

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/intelowlproject/go-intelowl/constants"
)

// TagParams represents the fields needed for creating and updating tags
type TagParams struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

// Tag represents a tag in an IntelOwl job.
type Tag struct {
	ID    uint64 `json:"id"`
	Label string `json:"label"`
	Color string `json:"color"`
}

// TagService handles communication with tag related methods of IntelOwl API.
//
// IntelOwl REST API tag docs: https://intelowlproject.github.io/docs/IntelOwl/api_docs/#tag/tags
type TagService struct {
	client *IntelOwlClient
}

// checkTagID is used to check if a tag	ID is valid (id should be greater than zero).
func checkTagID(id uint64) error {
	if id > 0 {
		return nil
	}
	return errors.New("Tag ID cannot be 0")
}

// List fetches all the working tags in IntelOwl.
//
//	Endpoint: GET "/api/tags"
//
// IntelOwl REST API docs: https://intelowlproject.github.io/docs/IntelOwl/api_docs/#tag/tags/operation/tags_list
func (tagService *TagService) List(ctx context.Context) (*[]Tag, error) {
	requestUrl := tagService.client.options.Url + constants.BASE_TAG_URL
	contentType := "application/json"
	method := "GET"
	request, err := tagService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return nil, err
	}
	successResp, err := tagService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	var tagList []Tag
	marashalError := json.Unmarshal(successResp.Data, &tagList)
	if marashalError != nil {
		return nil, marashalError
	}

	return &tagList, nil
}

// Get fetches a specific tag through its tag ID.
//
//	Endpoint: GET "/api/tags/{id}"
//
// IntelOwl REST API docs: https://intelowlproject.github.io/docs/IntelOwl/api_docs/#tag/tags/operation/tags_retrieve
func (tagService *TagService) Get(ctx context.Context, tagId uint64) (*Tag, error) {
	if err := checkTagID(tagId); err != nil {
		return nil, err
	}
	route := tagService.client.options.Url + constants.SPECIFIC_TAG_URL
	requestUrl := fmt.Sprintf(route, tagId)
	contentType := "application/json"
	method := "GET"
	request, err := tagService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return nil, err
	}
	var tagResponse Tag
	successResp, err := tagService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	unmarshalError := json.Unmarshal(successResp.Data, &tagResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return &tagResponse, nil
}

// Create lets you easily create a new tag by passing TagParams.
//
//	Endpoint: POST "/api/tags/"
//
// IntelOwl REST API docs: https://intelowlproject.github.io/docs/IntelOwl/api_docs/#tag/tags/operation/tags_create
func (tagService *TagService) Create(ctx context.Context, tagParams *TagParams) (*Tag, error) {
	requestUrl := tagService.client.options.Url + constants.BASE_TAG_URL
	tagJson, err := json.Marshal(tagParams)
	if err != nil {
		return nil, err
	}
	contentType := "application/json"
	method := "POST"
	body := bytes.NewBuffer(tagJson)
	request, err := tagService.client.buildRequest(ctx, method, contentType, body, requestUrl)
	if err != nil {
		return nil, err
	}
	var createdTag Tag
	successResp, err := tagService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	unmarshalError := json.Unmarshal(successResp.Data, &createdTag)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return &createdTag, nil
}

// Update lets you edit a tag throght its tag ID.
//
//	Endpoint: PUT "/api/tags/{id}"
//
// IntelOwl REST API docs: https://intelowlproject.github.io/docs/IntelOwl/api_docs/#tag/tags/operation/tags_update
func (tagService *TagService) Update(ctx context.Context, tagId uint64, tagParams *TagParams) (*Tag, error) {
	route := tagService.client.options.Url + constants.SPECIFIC_TAG_URL
	requestUrl := fmt.Sprintf(route, tagId)
	// Getting the relevant JSON data
	tagJson, err := json.Marshal(tagParams)
	if err != nil {
		return nil, err
	}
	contentType := "application/json"
	method := "PUT"
	body := bytes.NewBuffer(tagJson)
	request, err := tagService.client.buildRequest(ctx, method, contentType, body, requestUrl)
	if err != nil {
		return nil, err
	}
	var updatedTag Tag
	successResp, err := tagService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	unmarshalError := json.Unmarshal(successResp.Data, &updatedTag)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &updatedTag, nil
}

// Delete removes the given tag from your IntelOwl instance.
//
//	Endpoint: DELETE "/api/tags/{id}"
//
// IntelOwl REST API docs: https://intelowlproject.github.io/docs/IntelOwl/api_docs/#tag/tags/operation/tags_destroy
func (tagService *TagService) Delete(ctx context.Context, tagId uint64) (bool, error) {
	if err := checkTagID(tagId); err != nil {
		return false, err
	}
	route := tagService.client.options.Url + constants.SPECIFIC_TAG_URL
	requestUrl := fmt.Sprintf(route, tagId)
	contentType := "application/json"
	method := "DELETE"
	request, err := tagService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return false, err
	}
	successResp, err := tagService.client.newRequest(ctx, request)
	if err != nil {
		return false, err
	}
	if successResp.StatusCode == http.StatusNoContent {
		return true, nil
	}
	return false, nil
}
