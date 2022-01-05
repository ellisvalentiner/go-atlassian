package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type FieldConfigurationService struct {
	client *Client
	Item   *FieldConfigurationItemService
}

// Gets Returns a paginated list of all field configurations.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-all-field-configurations
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-field-configurations/#api-rest-api-2-fieldconfiguration-get
func (f *FieldConfigurationService) Gets(ctx context.Context, ids []int, isDefault bool, startAt, maxResults int) (
	result *models.FieldConfigurationPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if isDefault {
		params.Add("isDefault", "true")
	}

	for _, id := range ids {
		params.Add("id", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/2/fieldconfiguration?%v", params.Encode())

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Create creates a field configuration. The field configuration is created with the same field properties as the
// default configuration, with all the fields being optional.
// This operation can only create configurations for use in company-managed (classic) projects.
// EXPERIMENTAL
func (f *FieldConfigurationService) Create(ctx context.Context, name, description string) (result *models.FieldConfigurationScheme,
	response *ResponseScheme, err error) {

	if name == "" {
		return nil, nil, models.ErrNoFieldConfigurationNameError
	}

	payload := struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}
	payloadAsReader, _ := transformStructToReader(&payload)
	endpoint := "rest/api/2/fieldconfiguration"

	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update updates a field configuration. The name and the description provided in the request override the existing values.
// This operation can only update configurations used in company-managed (classic) projects.
// EXPERIMENTAL
func (f *FieldConfigurationService) Update(ctx context.Context, fieldConfigurationID int, name, description string) (
	response *ResponseScheme, err error) {

	if fieldConfigurationID == 0 {
		return nil, models.ErrNoFieldConfigurationIDError
	}

	if name == "" {
		return nil, models.ErrNoFieldConfigurationNameError
	}

	payload := struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	endpoint := fmt.Sprintf("rest/api/2/fieldconfiguration/%v", fieldConfigurationID)

	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Delete deletes a field configuration.
// This operation can only delete configurations used in company-managed (classic) projects.
// EXPERIMENTAL
func (f *FieldConfigurationService) Delete(ctx context.Context, fieldConfigurationID int) (response *ResponseScheme, err error) {

	if fieldConfigurationID == 0 {
		return nil, models.ErrNoFieldConfigurationIDError
	}

	endpoint := fmt.Sprintf("rest/api/2/fieldconfiguration/%v", fieldConfigurationID)

	request, err := f.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Schemes Returns a paginated list of field configuration schemes.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-all-field-configuration-schemes
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-field-configurations/#api-rest-api-2-fieldconfigurationscheme-get
func (f *FieldConfigurationService) Schemes(ctx context.Context, IDs []int, startAt, maxResults int) (
	result *models.FieldConfigurationSchemePageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range IDs {
		params.Add("id", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/2/fieldconfigurationscheme?%v", params.Encode())

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// IssueTypeItems Returns a paginated list of field configuration issue type items.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-field-configuration-issue-type-items
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-field-configurations/#api-rest-api-2-fieldconfigurationscheme-mapping-get
func (f *FieldConfigurationService) IssueTypeItems(ctx context.Context, fieldConfigIDs []int, startAt, maxResults int) (
	result *models.FieldConfigurationIssueTypeItemPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range fieldConfigIDs {
		params.Add("fieldConfigurationSchemeId", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/2/fieldconfigurationscheme/mapping?%v", params.Encode())

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// SchemesByProject Returns a paginated list of field configuration schemes and, for each scheme, a list of the projects that use it.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-field-configuration-schemes-for-projects
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-field-configurations/#api-rest-api-2-fieldconfigurationscheme-project-get
func (f *FieldConfigurationService) SchemesByProject(ctx context.Context, projectIDs []int, startAt, maxResults int) (
	result *models.FieldConfigurationSchemeProjectPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, projectID := range projectIDs {
		params.Add("projectId", strconv.Itoa(projectID))
	}

	var endpoint = fmt.Sprintf("rest/api/2/fieldconfigurationscheme/project?%v", params.Encode())

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Assign assigns a field configuration scheme to a project.
// If the field configuration scheme ID is null, the operation assigns the default field configuration scheme.
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-field-configurations/#api-rest-api-2-fieldconfigurationscheme-project-put
func (f *FieldConfigurationService) Assign(ctx context.Context, fieldConfigurationSchemeID, projectID string) (response *ResponseScheme, err error) {

	if len(projectID) == 0 {
		return nil, models.ErrNoProjectIDError
	}

	payload := struct {
		SchemeID  string `json:"fieldConfigurationSchemeId,omitempty"`
		ProjectID string `json:"projectId,omitempty"`
	}{
		SchemeID:  fieldConfigurationSchemeID,
		ProjectID: projectID,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	var endpoint = "rest/api/2/fieldconfigurationscheme/project"

	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, nil)
	if err != nil {
		return
	}

	return
}
