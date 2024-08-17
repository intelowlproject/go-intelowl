package gointelowl

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/intelowlproject/go-intelowl/constants"
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

// Access retrieves user details
//
//	Endpoint: GET /api/me/access
//
// IntelOwl REST API docs: https://intelowlproject.github.io/docs/IntelOwl/api_docs/#tag/me/operation/me_access_retrieve
func (userService *UserService) Access(ctx context.Context) (*User, error) {
	requestUrl := userService.client.options.Url + constants.USER_DETAILS_URL
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

// Organization returns the organization's details.
//
//	Endpoint: GET /api/me/organization
//
// IntelOwl REST API docs: https://intelowlproject.github.io/docs/IntelOwl/api_docs/#tag/me/operation/me_organization_list
func (userService *UserService) Organization(ctx context.Context) (*Organization, error) {
	requestUrl := userService.client.options.Url + constants.ORGANIZATION_URL
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

// CreateOrganization allows you to create a super cool organization!
//
//	Endpoint: POST /api/me/organization
//
// IntelOwl REST API docs: https://intelowlproject.github.io/docs/IntelOwl/api_docs/#tag/me/operation/me_organization_create
func (userService *UserService) CreateOrganization(ctx context.Context, organizationParams *OrganizationParams) (*Organization, error) {
	requestUrl := userService.client.options.Url + constants.ORGANIZATION_URL
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

// InviteToOrganization allows you to invite someone to your super cool organization!
// This is only accessible to the organization's owner.
//
//	Endpoint: POST /api/me/organization/invite
//
// IntelOwl REST API docs: https://intelowlproject.github.io/docs/IntelOwl/api_docs/#tag/me/operation/me_organization_invite_create
func (userService *UserService) InviteToOrganization(ctx context.Context, memberParams *MemberParams) (*Invite, error) {
	requestUrl := userService.client.options.Url + constants.INVITE_TO_ORGANIZATION_URL
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

// RemoveMemberFromOrganization lets you remove someone from your super cool organization! (you had your reasons)
// This is only accessible to the organization's owner.
//
//	Endpoint: POST /api/me/organization/remove_member
//
// IntelOwl REST API docs: https://intelowlproject.github.io/docs/IntelOwl/api_docs/#tag/me/operation/me_organization_create
func (userService *UserService) RemoveMemberFromOrganization(ctx context.Context, memberParams *MemberParams) (bool, error) {
	requestUrl := userService.client.options.Url + constants.REMOVE_MEMBER_FROM_ORGANIZATION_URL
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
