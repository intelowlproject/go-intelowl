package gointelowl

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Tag optional parameters to easily create and update a tag
// Tag optional parameters to easily create and update a tag
type TagParams struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

// Tag model to easily marshal and unmarshal.
// Tag model to easily marshal and unmarshal.
type Tag struct {
	ID    uint64 `json:"id"`
	Label string `json:"label"`
	Color string `json:"color"`
}

// Service Object to easily access the tag endpoints
type TagService struct {
	client *IntelOwlClient
}

// * helper functions!
func checkTagID(id uint64) error {
	if id > 0 {
		return nil
	}
	return errors.New("Tag ID cannot be 0")
}

// Desc: Getting the list of tags
//
//	Endpoint: GET "/api/tags"
func (tagService *TagService) List(ctx context.Context) (*[]Tag, error) {
	requestUrl := fmt.Sprintf(BASE_TAG_URL, tagService.client.options.Url)
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

// Desc: Get a tag through it ID!
//
//	Endpoint: GET "/api/tags/{id}"
func (tagService *TagService) Get(ctx context.Context, tagId uint64) (*Tag, error) {
	if err := checkTagID(tagId); err != nil {
		return nil, err
	}
	requestUrl := fmt.Sprintf(SPECIFIC_TAG_URL, tagService.client.options.Url, tagId)
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

// Desc: Create a Tag!
//
//	Endpoint: POST "/api/tags/"
func (tagService *TagService) Create(ctx context.Context, tagParams *TagParams) (*Tag, error) {
	requestUrl := fmt.Sprintf(BASE_TAG_URL, tagService.client.options.Url)
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

// Desc: Update a tag through it ID!
//
//	Endpoint: PUT "/api/tags/{id}"
func (tagService *TagService) Update(ctx context.Context, tagId uint64, tagParams *TagParams) (*Tag, error) {
	requestUrl := fmt.Sprintf(SPECIFIC_TAG_URL, tagService.client.options.Url, tagId)
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

// Desc: Delete a tag through it ID!
//
//	Endpoint: DELETE "/api/tags/{id}"
func (tagService *TagService) Delete(ctx context.Context, tagId uint64) (bool, error) {
	if err := checkTagID(tagId); err != nil {
		return false, err
	}
	requestUrl := fmt.Sprintf(SPECIFIC_TAG_URL, tagService.client.options.Url, tagId)
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
