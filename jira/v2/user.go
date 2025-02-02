package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type UserService struct {
	client *Client
	Search *UserSearchService
}

// Get returns a user
// Docs: https://docs.go-atlassian.io/jira-software-cloud/users#get-user
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-users/#api-rest-api-2-user-get
func (u *UserService) Get(ctx context.Context, accountID string, expand []string) (result *models.UserScheme,
	response *ResponseScheme, err error) {

	if len(accountID) == 0 {
		return nil, nil, models.ErrNoAccountIDError
	}

	params := url.Values{}

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	if len(accountID) != 0 {
		params.Add("accountId", accountID)
	}

	var endpoint strings.Builder
	endpoint.WriteString("/rest/api/2/user")

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Create creates a user. This resource is retained for legacy compatibility.
// As soon as a more suitable alternative is available this resource will be deprecated.
// The option is provided to set or generate a password for the user.
// When using the option to generate a password, by omitting password from the request, include "notification": "true" to ensure the user is
// sent an email advising them that their account is created.
// This email includes a link for the user to set their password. If the notification isn't sent for a generated password,
// the user will need to be sent a reset password request from Jira.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/users#create-user
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-users/#api-rest-api-2-user-post
func (u *UserService) Create(ctx context.Context, payload *models.UserPayloadScheme) (result *models.UserScheme, response *ResponseScheme,
	err error) {

	var endpoint = "rest/api/2/user"

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := u.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = u.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes a user.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/users#delete-user
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-users/#api-rest-api-2-user-delete
func (u *UserService) Delete(ctx context.Context, accountID string) (response *ResponseScheme, err error) {

	if len(accountID) == 0 {
		return nil, models.ErrNoAccountIDError
	}

	params := url.Values{}
	params.Add("accountId", accountID)
	var endpoint = fmt.Sprintf("rest/api/2/user?%v", params.Encode())

	request, err := u.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = u.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Find returns a paginated list of the users specified by one or more account IDs.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/users#bulk-get-users
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-users/#api-rest-api-2-user-bulk-get
func (u *UserService) Find(ctx context.Context, accountIDs []string, startAt, maxResults int) (result *models.UserSearchPageScheme,
	response *ResponseScheme, err error) {

	if len(accountIDs) == 0 {
		return nil, nil, models.ErrNoAccountSliceError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, accountID := range accountIDs {
		params.Add("accountId", accountID)
	}

	var endpoint = fmt.Sprintf("rest/api/2/user/bulk?%v", params.Encode())

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Groups returns the groups to which a user belongs.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/users#get-user-groups
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-users/#api-rest-api-2-user-groups-get
func (u *UserService) Groups(ctx context.Context, accountID string) (result []*models.UserGroupScheme, response *ResponseScheme,
	err error) {

	if len(accountID) == 0 {
		return nil, nil, models.ErrNoAccountIDError
	}

	params := url.Values{}
	params.Add("accountId", accountID)

	var endpoint = fmt.Sprintf("rest/api/2/user/groups?%v", params.Encode())

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Gets returns a list of all (active and inactive) users.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/users#get-all-users
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-users/#api-rest-api-2-users-search-get
func (u *UserService) Gets(ctx context.Context, startAt, maxResults int) (result []*models.UserScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("rest/api/2/users/search?%v", params.Encode())

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
