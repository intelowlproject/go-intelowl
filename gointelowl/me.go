package gointelowl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

type Owner struct {
	Username string    `json:"username"`
	FullName string    `json:"full_name"`
	Joined   time.Time `json:"joined"`
}

type Organization struct {
	MembersCount int        `json:"members_count"`
	Owner        Owner      `json:"owner"`
	IsUserOwner  bool       `json:"is_user_owner,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	Name         string     `json:"name"`
}

type OrganizationParams struct {
	Name string `json:"name"`
}

type MemberParams struct {
	Username string `json:"username"`
}

type Invite struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"`
}

type Invitation struct {
	Invite
	Organization Organization `json:"organization"`
}

type InvitationParams struct {
	Organization OrganizationParams `json:"organization"`
	Status       string             `json:"status"`
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

/*
* Desc: Get Organization details!
* Endpoint: GET /api/me/organization
 */
func (userService *UserService) Organization(ctx context.Context) (*Organization, error) {
	requestUrl := fmt.Sprintf("%s/api/me/organization", userService.client.options.Url)
	contentType := "application/json"
	method := "GET"
	request, err := userService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return nil, err
	}

	org := Organization{}
	successResp, err := userService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	if unmarshalError := json.Unmarshal(successResp.Data, &org); unmarshalError != nil {
		return nil, unmarshalError
	}
	return &org, nil
}

/*
* Desc: Create a super cool organization!
* Endpoint: POST /api/me/organization
 */
func (userService *UserService) CreateOrganization(ctx context.Context, organizationParams *OrganizationParams) (*Organization, error) {
	requestUrl := fmt.Sprintf("%s/api/me/organization", userService.client.options.Url)
	// Getting the relevant JSON data
	orgJson, err := json.Marshal(organizationParams)
	if err != nil {
		return nil, err
	}
	contentType := "application/json"
	method := "POST"
	body := bytes.NewBuffer(orgJson)
	request, err := userService.client.buildRequest(ctx, method, contentType, body, requestUrl)
	if err != nil {
		return nil, err
	}

	org := Organization{}
	successResp, err := userService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	if unmarshalError := json.Unmarshal(successResp.Data, &org); unmarshalError != nil {
		return nil, unmarshalError
	}
	return &org, nil
}

/*
* Desc: Invite someone to your super cool Organization!
* Endpoint: POST /api/me/organization/invite
 */
func (userService *UserService) InviteToOrganization(ctx context.Context, memberParams *MemberParams) (*Invite, error) {
	requestUrl := fmt.Sprintf("%s/api/me/organization/invite", userService.client.options.Url)
	// Getting the relevant JSON data
	memberJson, err := json.Marshal(memberParams)
	if err != nil {
		return nil, err
	}
	contentType := "application/json"
	method := "POST"
	body := bytes.NewBuffer(memberJson)
	request, err := userService.client.buildRequest(ctx, method, contentType, body, requestUrl)
	if err != nil {
		return nil, err
	}

	invite := Invite{}
	successResp, err := userService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	if unmarshalError := json.Unmarshal(successResp.Data, &invite); unmarshalError != nil {
		return nil, unmarshalError
	}
	return &invite, nil
}

/*
* Desc: Remove someone from your super cool Organization! (you had your reasons)
* Endpoint: POST /api/me/organization/remove_member
 */
func (userService *UserService) RemoveMemberFromOrganization(ctx context.Context, memberParams *MemberParams) (bool, error) {
	requestUrl := fmt.Sprintf("%s/api/me/organization/remove_member", userService.client.options.Url)
	// Getting the relevant JSON data
	memberJson, err := json.Marshal(memberParams)
	if err != nil {
		return false, err
	}
	contentType := "application/json"
	method := "POST"
	body := bytes.NewBuffer(memberJson)
	request, err := userService.client.buildRequest(ctx, method, contentType, body, requestUrl)
	if err != nil {
		return false, err
	}

	successResp, err := userService.client.newRequest(ctx, request)
	if err != nil {
		return false, err
	}

	if successResp.StatusCode == http.StatusNoContent {
		return true, nil
	}
	return false, nil
}
