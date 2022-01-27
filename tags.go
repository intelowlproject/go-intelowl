package gointelowl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Tag struct {
	id    int    `json:"id"`
	label string `json:"label"`
	color string `json:"color"`
}

type TagList []Tag

func (c *IntelOwlClient) GetAllTags(ctx context.Context) (TagList, error) {
	var tags_json TagList

	// Make API request to /api/tags
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/tags", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	responded, err := c.sendRequest(req)

	if err != nil {
		return nil, err
	}

	tags_bytes := []byte(responded.Data)
	json.Unmarshal(tags_bytes, &tags_json)

	return tags_json, nil
}
