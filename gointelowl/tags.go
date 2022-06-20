package gointelowl

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type TagParams struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type Tag struct {
	ID    uint64 `json:"id"`
	Label string `json:"label"`
	Color string `json:"color"`
}

// * Service Object!
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

// * getting on all tags
func (tagService *TagService) List(ctx context.Context) (*[]Tag, error) {
	requestUrl := fmt.Sprintf("%s/api/tags", tagService.client.options.Url)
	fmt.Println(requestUrl)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	successResp, err := tagService.client.makeRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	var tagList []Tag
	json.Unmarshal(successResp.Data, &tagList)
	return &tagList, nil
}

// * Getting a tag through its ID
func (tagService *TagService) Get(ctx context.Context, tagId uint64) (*Tag, error) {
	if err := checkTagID(tagId); err != nil {
		return nil, err
	}
	requestUrl := fmt.Sprintf("%s/api/tags/%d", tagService.client.options.Url, tagId)
	fmt.Println(requestUrl)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	var tagResponse Tag
	successResp, err := tagService.client.makeRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(successResp.Data, &tagResponse)
	return &tagResponse, nil
}

// //* Creating a Tag
func (tagService *TagService) Create(ctx context.Context, tagParams *TagParams) (*Tag, error) {
	requestUrl := fmt.Sprintf("%s/api/tags", tagService.client.options.Url)
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
	var createdTag Tag
	successResp, err := tagService.client.makeRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(successResp.Data, &createdTag)
	return &createdTag, nil
}

//* Updating a tag
func (tagService *TagService) Update(ctx context.Context, tagId uint64, tagParams *TagParams) (*Tag, error) {
	requestUrl := fmt.Sprintf("%s/api/tags/%d", tagService.client.options.Url, tagId)
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
	var updatedTag Tag
	successResp, err := tagService.client.makeRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(successResp.Data, &updatedTag)
	return &updatedTag, nil
}

//* Deleting a tag
func (tagService *TagService) Delete(ctx context.Context, tagId uint64) (bool, error) {
	if err := checkTagID(tagId); err != nil {
		return false, err
	}
	requestUrl := fmt.Sprintf("%s/api/tags/%d", tagService.client.options.Url, tagId)
	// printing the request
	fmt.Println("Url: " + requestUrl)
	request, err := http.NewRequest("DELETE", requestUrl, nil)
	if err != nil {
		return false, err
	}
	successResp, err := tagService.client.makeRequest(ctx, request)
	if err != nil {
		return false, err
	}
	if successResp.StatusCode == 204 {
		return true, nil
	}
	return false, nil
}

// Pretty printing the tag
func (tag *Tag) Display() {
	display := fmt.Sprintf("ID:%d\nLabel:%s\nColor:%s", tag.ID, tag.Label, tag.Color)
	fmt.Println(display)
}
