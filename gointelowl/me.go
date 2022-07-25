package gointelowl

import (
	"context"
	"encoding/json"
	"fmt"
)

type Details struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
}

type AccessDetails struct {
	TotalSubmissions int `json:"total_submissions"`
	MonthSubmissions int `json:"month_submissions"`
}

type User struct {
	User   Details       `json:"user"`
	Access AccessDetails `json:"access"`
}

type UserService struct {
	client *IntelOwlClient
}

/*
* Desc: Get User Details
* Endpoint: GET /api/me/access
 */
func (userService *UserService) Access(ctx context.Context) (*User, error) {
	requestUrl := fmt.Sprintf("%s/api/me/access", userService.client.options.Url)
	contentType := "application/json"
	method := "GET"
	request, err := userService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return nil, err
	}
	user := User{}
	successResp, err := userService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	if unmarshalError := json.Unmarshal(successResp.Data, &user); unmarshalError != nil {
		return nil, unmarshalError
	}
	return &user, nil
}
