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

/*
* Desc: Getting all tags
* Endpoint: GET "/api/tags"
 */
func (tagService *TagService) List(ctx context.Context) (*[]Tag, error) {
	requestUrl := fmt.Sprintf("%s/api/tags", tagService.client.options.Url)
	request, err := http.NewRequest("GET", requestUrl, nil)
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

/*
* Desc: Getting a tag through it ID!
* Endpoint: GET "/api/tags/{id}"
 */
func (tagService *TagService) Get(ctx context.Context, tagId uint64) (*Tag, error) {
	if err := checkTagID(tagId); err != nil {
		return nil, err
	}
	requestUrl := fmt.Sprintf("%s/api/tags/%d", tagService.client.options.Url, tagId)
	request, err := http.NewRequest("GET", requestUrl, nil)
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

/*
* Desc: Creating a Tag!
* Endpoint: POST "/api/tags/"
 */
func (tagService *TagService) Create(ctx context.Context, tagParams *TagParams) (*Tag, error) {
	requestUrl := fmt.Sprintf("%s/api/tags", tagService.client.options.Url)
	tagJson, err := json.Marshal(tagParams)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(tagJson))
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

/*
* Desc: Updating a tag through it ID!
* Endpoint: PUT "/api/tags/{id}"
 */
func (tagService *TagService) Update(ctx context.Context, tagId uint64, tagParams *TagParams) (*Tag, error) {
	requestUrl := fmt.Sprintf("%s/api/tags/%d", tagService.client.options.Url, tagId)
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

/*
* Desc: Deleting a tag through it ID!
* Endpoint: DELETE "/api/tags/{id}"
 */
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
	successResp, err := tagService.client.newRequest(ctx, request)
	if err != nil {
		return false, err
	}
	if successResp.StatusCode == http.StatusNoContent {
		return true, nil
	}
	return false, nil
}

//* Pretty printing the tag
func (tag *Tag) Display() {
	display := fmt.Sprintf("ID:%d\nLabel:%s\nColor:%s", tag.ID, tag.Label, tag.Color)
	fmt.Println(display)
}
