/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.39.0-748eb4ca-20210917-165907
 */

// Package iampolicymanagementv1 : Operations and models for the IamPolicyManagementV1 service
package iampolicymanagementv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/go-openapi/strfmt"
)

// IamPolicyManagementV1 : IAM Policy Management API
//
// API Version: 1.0.1
type IamPolicyManagementV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://iam.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "iam_policy_management"

// IamPolicyManagementV1Options : Service options
type IamPolicyManagementV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewIamPolicyManagementV1UsingExternalConfig : constructs an instance of IamPolicyManagementV1 with passed in options and external configuration.
func NewIamPolicyManagementV1UsingExternalConfig(options *IamPolicyManagementV1Options) (iamPolicyManagement *IamPolicyManagementV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	iamPolicyManagement, err = NewIamPolicyManagementV1(options)
	if err != nil {
		return
	}

	err = iamPolicyManagement.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = iamPolicyManagement.Service.SetServiceURL(options.URL)
	}
	return
}

// NewIamPolicyManagementV1 : constructs an instance of IamPolicyManagementV1 with passed in options.
func NewIamPolicyManagementV1(options *IamPolicyManagementV1Options) (service *IamPolicyManagementV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &IamPolicyManagementV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "iamPolicyManagement" suitable for processing requests.
func (iamPolicyManagement *IamPolicyManagementV1) Clone() *IamPolicyManagementV1 {
	if core.IsNil(iamPolicyManagement) {
		return nil
	}
	clone := *iamPolicyManagement
	clone.Service = iamPolicyManagement.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (iamPolicyManagement *IamPolicyManagementV1) SetServiceURL(url string) error {
	return iamPolicyManagement.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (iamPolicyManagement *IamPolicyManagementV1) GetServiceURL() string {
	return iamPolicyManagement.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (iamPolicyManagement *IamPolicyManagementV1) SetDefaultHeaders(headers http.Header) {
	iamPolicyManagement.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (iamPolicyManagement *IamPolicyManagementV1) SetEnableGzipCompression(enableGzip bool) {
	iamPolicyManagement.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (iamPolicyManagement *IamPolicyManagementV1) GetEnableGzipCompression() bool {
	return iamPolicyManagement.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (iamPolicyManagement *IamPolicyManagementV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	iamPolicyManagement.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (iamPolicyManagement *IamPolicyManagementV1) DisableRetries() {
	iamPolicyManagement.Service.DisableRetries()
}

// ListPolicies : Get policies by attributes
// Get policies and filter by attributes. While managing policies, you may want to retrieve policies in the account and
// filter by attribute values. This can be done through query parameters. Currently, only the following attributes are
// supported: account_id, iam_id, access_group_id, type, service_type, sort, format and state. account_id is a required
// query parameter. Only policies that have the specified attributes and that the caller has read access to are
// returned. If the caller does not have read access to any policies an empty array is returned.
func (iamPolicyManagement *IamPolicyManagementV1) ListPolicies(listPoliciesOptions *ListPoliciesOptions) (result *PolicyList, response *core.DetailedResponse, err error) {
	return iamPolicyManagement.ListPoliciesWithContext(context.Background(), listPoliciesOptions)
}

// ListPoliciesWithContext is an alternate form of the ListPolicies method which supports a Context parameter
func (iamPolicyManagement *IamPolicyManagementV1) ListPoliciesWithContext(ctx context.Context, listPoliciesOptions *ListPoliciesOptions) (result *PolicyList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listPoliciesOptions, "listPoliciesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listPoliciesOptions, "listPoliciesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamPolicyManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamPolicyManagement.Service.Options.URL, `/v1/policies`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPoliciesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "ListPolicies")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPoliciesOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*listPoliciesOptions.AcceptLanguage))
	}

	builder.AddQuery("account_id", fmt.Sprint(*listPoliciesOptions.AccountID))
	if listPoliciesOptions.IamID != nil {
		builder.AddQuery("iam_id", fmt.Sprint(*listPoliciesOptions.IamID))
	}
	if listPoliciesOptions.AccessGroupID != nil {
		builder.AddQuery("access_group_id", fmt.Sprint(*listPoliciesOptions.AccessGroupID))
	}
	if listPoliciesOptions.Type != nil {
		builder.AddQuery("type", fmt.Sprint(*listPoliciesOptions.Type))
	}
	if listPoliciesOptions.ServiceType != nil {
		builder.AddQuery("service_type", fmt.Sprint(*listPoliciesOptions.ServiceType))
	}
	if listPoliciesOptions.TagName != nil {
		builder.AddQuery("tag_name", fmt.Sprint(*listPoliciesOptions.TagName))
	}
	if listPoliciesOptions.TagValue != nil {
		builder.AddQuery("tag_value", fmt.Sprint(*listPoliciesOptions.TagValue))
	}
	if listPoliciesOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listPoliciesOptions.Sort))
	}
	if listPoliciesOptions.Format != nil {
		builder.AddQuery("format", fmt.Sprint(*listPoliciesOptions.Format))
	}
	if listPoliciesOptions.State != nil {
		builder.AddQuery("state", fmt.Sprint(*listPoliciesOptions.State))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPolicyList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreatePolicy : Create a policy
// Creates a policy to grant access between a subject and a resource. There are two types of policies: **access** and
// **authorization**. A policy administrator might want to create an access policy which grants access to a user,
// service-id, or an access group. They might also want to create an authorization policy and setup access between
// services.
//
// ### Access
//
// To create an access policy, use **`"type": "access"`** in the body. The possible subject attributes are **`iam_id`**
// and **`access_group_id`**. Use the **`iam_id`** subject attribute for assigning access for a user or service-id. Use
// the **`access_group_id`** subject attribute for assigning access for an access group. The roles must be a subset of a
// service's or the platform's supported roles. The resource attributes must be a subset of a service's or the
// platform's supported attributes. The policy resource must include either the **`serviceType`**, **`serviceName`**,
// or **`resourceGroupId`** attribute and the **`accountId`** attribute.` If the subject is a locked service-id, the
// request will fail.
//
// ### Authorization
//
// Authorization policies are supported by services on a case by case basis. Refer to service documentation to verify
// their support of authorization policies. To create an authorization policy, use **`"type": "authorization"`** in the
// body. The subject attributes must match the supported authorization subjects of the resource. Multiple subject
// attributes might be provided. The following attributes are supported:
//   serviceName, serviceInstance, region, resourceType, resource, accountId The policy roles must be a subset of the
// supported authorization roles supported by the target service. The user must also have the same level of access or
// greater to the target resource in order to grant the role. The resource attributes must be a subset of a service's or
// the platform's supported attributes. Both the policy subject and the policy resource must include the
// **`serviceName`** and **`accountId`** attributes.
//
// ### Attribute Operators
//
// Currently, only the `stringEquals` and the `stringMatch` operators are available. Resource attributes may support one
// or both operators.  For more information, see [how to assign access by using wildcards
// policies](https://cloud.ibm.com/docs/account?topic=account-wildcard).
//
// ### Attribute Validations
//
// Policy attribute values must be between 1 and 1,000 characters in length. If location related attributes like
// geography, country, metro, region, satellite, and locationvalues are supported by the service, they are validated
// against Global Catalog locations.
func (iamPolicyManagement *IamPolicyManagementV1) CreatePolicy(createPolicyOptions *CreatePolicyOptions) (result *Policy, response *core.DetailedResponse, err error) {
	return iamPolicyManagement.CreatePolicyWithContext(context.Background(), createPolicyOptions)
}

// CreatePolicyWithContext is an alternate form of the CreatePolicy method which supports a Context parameter
func (iamPolicyManagement *IamPolicyManagementV1) CreatePolicyWithContext(ctx context.Context, createPolicyOptions *CreatePolicyOptions) (result *Policy, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPolicyOptions, "createPolicyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPolicyOptions, "createPolicyOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamPolicyManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamPolicyManagement.Service.Options.URL, `/v1/policies`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPolicyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "CreatePolicy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createPolicyOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*createPolicyOptions.AcceptLanguage))
	}

	body := make(map[string]interface{})
	if createPolicyOptions.Type != nil {
		body["type"] = createPolicyOptions.Type
	}
	if createPolicyOptions.Subjects != nil {
		body["subjects"] = createPolicyOptions.Subjects
	}
	if createPolicyOptions.Roles != nil {
		body["roles"] = createPolicyOptions.Roles
	}
	if createPolicyOptions.Resources != nil {
		body["resources"] = createPolicyOptions.Resources
	}
	if createPolicyOptions.Description != nil {
		body["description"] = createPolicyOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPolicy)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdatePolicy : Update a policy
// Update a policy to grant access between a subject and a resource. A policy administrator might want to update an
// existing policy. The policy type cannot be changed (You cannot change an access policy to an authorization policy).
//
// ### Access
//
// To update an access policy, use **`"type": "access"`** in the body. The possible subject attributes are **`iam_id`**
// and **`access_group_id`**. Use the **`iam_id`** subject attribute for assigning access for a user or service-id. Use
// the **`access_group_id`** subject attribute for assigning access for an access group. The roles must be a subset of a
// service's or the platform's supported roles. The resource attributes must be a subset of a service's or the
// platform's supported attributes. The policy resource must include either the **`serviceType`**, **`serviceName`**,
// or **`resourceGroupId`** attribute and the **`accountId`** attribute.` If the subject is a locked service-id, the
// request will fail.
//
// ### Authorization
//
// To update an authorization policy, use **`"type": "authorization"`** in the body. The subject attributes must match
// the supported authorization subjects of the resource. Multiple subject attributes might be provided. The following
// attributes are supported:
//   serviceName, serviceInstance, region, resourceType, resource, accountId The policy roles must be a subset of the
// supported authorization roles supported by the target service. The user must also have the same level of access or
// greater to the target resource in order to grant the role. The resource attributes must be a subset of a service's or
// the platform's supported attributes. Both the policy subject and the policy resource must include the
// **`serviceName`** and **`accountId`** attributes.
//
// ### Attribute Operators
//
// Currently, only the `stringEquals` and the `stringMatch` operators are available. Resource attributes might support
// one or both operators.  For more information, see [how to assign access by using wildcards
// policies](https://cloud.ibm.com/docs/account?topic=account-wildcard).
//
// ### Attribute Validations
//
// Policy attribute values must be between 1 and 1,000 characters in length. If location related attributes like
// geography, country, metro, region, satellite, and locationvalues are supported by the service, they are validated
// against Global Catalog locations.
func (iamPolicyManagement *IamPolicyManagementV1) UpdatePolicy(updatePolicyOptions *UpdatePolicyOptions) (result *Policy, response *core.DetailedResponse, err error) {
	return iamPolicyManagement.UpdatePolicyWithContext(context.Background(), updatePolicyOptions)
}

// UpdatePolicyWithContext is an alternate form of the UpdatePolicy method which supports a Context parameter
func (iamPolicyManagement *IamPolicyManagementV1) UpdatePolicyWithContext(ctx context.Context, updatePolicyOptions *UpdatePolicyOptions) (result *Policy, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updatePolicyOptions, "updatePolicyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updatePolicyOptions, "updatePolicyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"policy_id": *updatePolicyOptions.PolicyID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamPolicyManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamPolicyManagement.Service.Options.URL, `/v1/policies/{policy_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updatePolicyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "UpdatePolicy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updatePolicyOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updatePolicyOptions.IfMatch))
	}

	body := make(map[string]interface{})
	if updatePolicyOptions.Type != nil {
		body["type"] = updatePolicyOptions.Type
	}
	if updatePolicyOptions.Subjects != nil {
		body["subjects"] = updatePolicyOptions.Subjects
	}
	if updatePolicyOptions.Roles != nil {
		body["roles"] = updatePolicyOptions.Roles
	}
	if updatePolicyOptions.Resources != nil {
		body["resources"] = updatePolicyOptions.Resources
	}
	if updatePolicyOptions.Description != nil {
		body["description"] = updatePolicyOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPolicy)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetPolicy : Retrieve a policy by ID
// Retrieve a policy by providing a policy ID.
func (iamPolicyManagement *IamPolicyManagementV1) GetPolicy(getPolicyOptions *GetPolicyOptions) (result *Policy, response *core.DetailedResponse, err error) {
	return iamPolicyManagement.GetPolicyWithContext(context.Background(), getPolicyOptions)
}

// GetPolicyWithContext is an alternate form of the GetPolicy method which supports a Context parameter
func (iamPolicyManagement *IamPolicyManagementV1) GetPolicyWithContext(ctx context.Context, getPolicyOptions *GetPolicyOptions) (result *Policy, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPolicyOptions, "getPolicyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPolicyOptions, "getPolicyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"policy_id": *getPolicyOptions.PolicyID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamPolicyManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamPolicyManagement.Service.Options.URL, `/v1/policies/{policy_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPolicyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "GetPolicy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPolicy)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeletePolicy : Delete a policy by ID
// Delete a policy by providing a policy ID. A policy cannot be deleted if the subject ID contains a locked service ID.
// If the subject of the policy is a locked service-id, the request will fail.
func (iamPolicyManagement *IamPolicyManagementV1) DeletePolicy(deletePolicyOptions *DeletePolicyOptions) (response *core.DetailedResponse, err error) {
	return iamPolicyManagement.DeletePolicyWithContext(context.Background(), deletePolicyOptions)
}

// DeletePolicyWithContext is an alternate form of the DeletePolicy method which supports a Context parameter
func (iamPolicyManagement *IamPolicyManagementV1) DeletePolicyWithContext(ctx context.Context, deletePolicyOptions *DeletePolicyOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePolicyOptions, "deletePolicyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deletePolicyOptions, "deletePolicyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"policy_id": *deletePolicyOptions.PolicyID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamPolicyManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamPolicyManagement.Service.Options.URL, `/v1/policies/{policy_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deletePolicyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "DeletePolicy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamPolicyManagement.Service.Request(request, nil)

	return
}

// PatchPolicy : Restore a deleted policy by ID
// Restore a policy that has recently been deleted. A policy administrator might want to restore a deleted policy. To
// restore a policy, use **`"state": "active"`** in the body.
func (iamPolicyManagement *IamPolicyManagementV1) PatchPolicy(patchPolicyOptions *PatchPolicyOptions) (result *Policy, response *core.DetailedResponse, err error) {
	return iamPolicyManagement.PatchPolicyWithContext(context.Background(), patchPolicyOptions)
}

// PatchPolicyWithContext is an alternate form of the PatchPolicy method which supports a Context parameter
func (iamPolicyManagement *IamPolicyManagementV1) PatchPolicyWithContext(ctx context.Context, patchPolicyOptions *PatchPolicyOptions) (result *Policy, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(patchPolicyOptions, "patchPolicyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(patchPolicyOptions, "patchPolicyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"policy_id": *patchPolicyOptions.PolicyID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamPolicyManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamPolicyManagement.Service.Options.URL, `/v1/policies/{policy_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range patchPolicyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "PatchPolicy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if patchPolicyOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*patchPolicyOptions.IfMatch))
	}

	body := make(map[string]interface{})
	if patchPolicyOptions.State != nil {
		body["state"] = patchPolicyOptions.State
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPolicy)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListRoles : Get roles by filters
// Get roles based on the filters. While managing roles, you may want to retrieve roles and filter by usages. This can
// be done through query parameters. Currently, we only support the following attributes: account_id, service_name,
// source_service_name and policy_type. Only roles that match the filter and that the caller has read access to are
// returned. If the caller does not have read access to any roles an empty array is returned.
func (iamPolicyManagement *IamPolicyManagementV1) ListRoles(listRolesOptions *ListRolesOptions) (result *RoleList, response *core.DetailedResponse, err error) {
	return iamPolicyManagement.ListRolesWithContext(context.Background(), listRolesOptions)
}

// ListRolesWithContext is an alternate form of the ListRoles method which supports a Context parameter
func (iamPolicyManagement *IamPolicyManagementV1) ListRolesWithContext(ctx context.Context, listRolesOptions *ListRolesOptions) (result *RoleList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listRolesOptions, "listRolesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamPolicyManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamPolicyManagement.Service.Options.URL, `/v2/roles`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listRolesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "ListRoles")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listRolesOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*listRolesOptions.AcceptLanguage))
	}

	if listRolesOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*listRolesOptions.AccountID))
	}
	if listRolesOptions.ServiceName != nil {
		builder.AddQuery("service_name", fmt.Sprint(*listRolesOptions.ServiceName))
	}
	if listRolesOptions.SourceServiceName != nil {
		builder.AddQuery("source_service_name", fmt.Sprint(*listRolesOptions.SourceServiceName))
	}
	if listRolesOptions.PolicyType != nil {
		builder.AddQuery("policy_type", fmt.Sprint(*listRolesOptions.PolicyType))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRoleList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateRole : Create a role
// Creates a custom role for a specific service within the account. An account owner or a user assigned the
// Administrator role on the Role management service can create a custom role. Any number of actions for a single
// service can be mapped to the new role, but there must be at least one service-defined action to successfully create
// the new role.
func (iamPolicyManagement *IamPolicyManagementV1) CreateRole(createRoleOptions *CreateRoleOptions) (result *CustomRole, response *core.DetailedResponse, err error) {
	return iamPolicyManagement.CreateRoleWithContext(context.Background(), createRoleOptions)
}

// CreateRoleWithContext is an alternate form of the CreateRole method which supports a Context parameter
func (iamPolicyManagement *IamPolicyManagementV1) CreateRoleWithContext(ctx context.Context, createRoleOptions *CreateRoleOptions) (result *CustomRole, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createRoleOptions, "createRoleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createRoleOptions, "createRoleOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamPolicyManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamPolicyManagement.Service.Options.URL, `/v2/roles`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createRoleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "CreateRole")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createRoleOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*createRoleOptions.AcceptLanguage))
	}

	body := make(map[string]interface{})
	if createRoleOptions.DisplayName != nil {
		body["display_name"] = createRoleOptions.DisplayName
	}
	if createRoleOptions.Actions != nil {
		body["actions"] = createRoleOptions.Actions
	}
	if createRoleOptions.Name != nil {
		body["name"] = createRoleOptions.Name
	}
	if createRoleOptions.AccountID != nil {
		body["account_id"] = createRoleOptions.AccountID
	}
	if createRoleOptions.ServiceName != nil {
		body["service_name"] = createRoleOptions.ServiceName
	}
	if createRoleOptions.Description != nil {
		body["description"] = createRoleOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomRole)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateRole : Update a role
// Update a custom role. A role administrator might want to update an existing role by updating the display name,
// description, or the actions that are mapped to the role. The name, account_id, and service_name can't be changed.
func (iamPolicyManagement *IamPolicyManagementV1) UpdateRole(updateRoleOptions *UpdateRoleOptions) (result *CustomRole, response *core.DetailedResponse, err error) {
	return iamPolicyManagement.UpdateRoleWithContext(context.Background(), updateRoleOptions)
}

// UpdateRoleWithContext is an alternate form of the UpdateRole method which supports a Context parameter
func (iamPolicyManagement *IamPolicyManagementV1) UpdateRoleWithContext(ctx context.Context, updateRoleOptions *UpdateRoleOptions) (result *CustomRole, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateRoleOptions, "updateRoleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateRoleOptions, "updateRoleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"role_id": *updateRoleOptions.RoleID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamPolicyManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamPolicyManagement.Service.Options.URL, `/v2/roles/{role_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateRoleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "UpdateRole")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateRoleOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateRoleOptions.IfMatch))
	}

	body := make(map[string]interface{})
	if updateRoleOptions.DisplayName != nil {
		body["display_name"] = updateRoleOptions.DisplayName
	}
	if updateRoleOptions.Description != nil {
		body["description"] = updateRoleOptions.Description
	}
	if updateRoleOptions.Actions != nil {
		body["actions"] = updateRoleOptions.Actions
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomRole)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetRole : Retrieve a role by ID
// Retrieve a role by providing a role ID.
func (iamPolicyManagement *IamPolicyManagementV1) GetRole(getRoleOptions *GetRoleOptions) (result *CustomRole, response *core.DetailedResponse, err error) {
	return iamPolicyManagement.GetRoleWithContext(context.Background(), getRoleOptions)
}

// GetRoleWithContext is an alternate form of the GetRole method which supports a Context parameter
func (iamPolicyManagement *IamPolicyManagementV1) GetRoleWithContext(ctx context.Context, getRoleOptions *GetRoleOptions) (result *CustomRole, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getRoleOptions, "getRoleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getRoleOptions, "getRoleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"role_id": *getRoleOptions.RoleID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamPolicyManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamPolicyManagement.Service.Options.URL, `/v2/roles/{role_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRoleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "GetRole")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomRole)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteRole : Delete a role by ID
// Delete a role by providing a role ID.
func (iamPolicyManagement *IamPolicyManagementV1) DeleteRole(deleteRoleOptions *DeleteRoleOptions) (response *core.DetailedResponse, err error) {
	return iamPolicyManagement.DeleteRoleWithContext(context.Background(), deleteRoleOptions)
}

// DeleteRoleWithContext is an alternate form of the DeleteRole method which supports a Context parameter
func (iamPolicyManagement *IamPolicyManagementV1) DeleteRoleWithContext(ctx context.Context, deleteRoleOptions *DeleteRoleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteRoleOptions, "deleteRoleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteRoleOptions, "deleteRoleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"role_id": *deleteRoleOptions.RoleID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamPolicyManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamPolicyManagement.Service.Options.URL, `/v2/roles/{role_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteRoleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "DeleteRole")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamPolicyManagement.Service.Request(request, nil)

	return
}

// CreatePolicyOptions : The CreatePolicy options.
type CreatePolicyOptions struct {
	// The policy type; either 'access' or 'authorization'.
	Type *string `json:"type" validate:"required"`

	// The subjects associated with a policy.
	Subjects []PolicySubject `json:"subjects" validate:"required"`

	// A set of role cloud resource names (CRNs) granted by the policy.
	Roles []PolicyRole `json:"roles" validate:"required"`

	// The resources associated with a policy.
	Resources []PolicyResource `json:"resources" validate:"required"`

	// Customer-defined description.
	Description *string `json:"description,omitempty"`

	// Language code for translations
	// * `default` - English
	// * `de` -  German (Standard)
	// * `en` - English
	// * `es` - Spanish (Spain)
	// * `fr` - French (Standard)
	// * `it` - Italian (Standard)
	// * `ja` - Japanese
	// * `ko` - Korean
	// * `pt-br` - Portuguese (Brazil)
	// * `zh-cn` - Chinese (Simplified, PRC)
	// * `zh-tw` - (Chinese, Taiwan).
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreatePolicyOptions : Instantiate CreatePolicyOptions
func (*IamPolicyManagementV1) NewCreatePolicyOptions(typeVar string, subjects []PolicySubject, roles []PolicyRole, resources []PolicyResource) *CreatePolicyOptions {
	return &CreatePolicyOptions{
		Type: core.StringPtr(typeVar),
		Subjects: subjects,
		Roles: roles,
		Resources: resources,
	}
}

// SetType : Allow user to set Type
func (_options *CreatePolicyOptions) SetType(typeVar string) *CreatePolicyOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetSubjects : Allow user to set Subjects
func (_options *CreatePolicyOptions) SetSubjects(subjects []PolicySubject) *CreatePolicyOptions {
	_options.Subjects = subjects
	return _options
}

// SetRoles : Allow user to set Roles
func (_options *CreatePolicyOptions) SetRoles(roles []PolicyRole) *CreatePolicyOptions {
	_options.Roles = roles
	return _options
}

// SetResources : Allow user to set Resources
func (_options *CreatePolicyOptions) SetResources(resources []PolicyResource) *CreatePolicyOptions {
	_options.Resources = resources
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreatePolicyOptions) SetDescription(description string) *CreatePolicyOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *CreatePolicyOptions) SetAcceptLanguage(acceptLanguage string) *CreatePolicyOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePolicyOptions) SetHeaders(param map[string]string) *CreatePolicyOptions {
	options.Headers = param
	return options
}

// CreateRoleOptions : The CreateRole options.
type CreateRoleOptions struct {
	// The display name of the role that is shown in the console.
	DisplayName *string `json:"display_name" validate:"required"`

	// The actions of the role. Please refer to [IAM roles and
	// actions](https://cloud.ibm.com/docs/account?topic=account-iam-service-roles-actions).
	Actions []string `json:"actions" validate:"required"`

	// The name of the role that is used in the CRN. Can only be alphanumeric and has to be capitalized.
	Name *string `json:"name" validate:"required"`

	// The account GUID.
	AccountID *string `json:"account_id" validate:"required"`

	// The service name.
	ServiceName *string `json:"service_name" validate:"required"`

	// The description of the role.
	Description *string `json:"description,omitempty"`

	// Language code for translations
	// * `default` - English
	// * `de` -  German (Standard)
	// * `en` - English
	// * `es` - Spanish (Spain)
	// * `fr` - French (Standard)
	// * `it` - Italian (Standard)
	// * `ja` - Japanese
	// * `ko` - Korean
	// * `pt-br` - Portuguese (Brazil)
	// * `zh-cn` - Chinese (Simplified, PRC)
	// * `zh-tw` - (Chinese, Taiwan).
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateRoleOptions : Instantiate CreateRoleOptions
func (*IamPolicyManagementV1) NewCreateRoleOptions(displayName string, actions []string, name string, accountID string, serviceName string) *CreateRoleOptions {
	return &CreateRoleOptions{
		DisplayName: core.StringPtr(displayName),
		Actions: actions,
		Name: core.StringPtr(name),
		AccountID: core.StringPtr(accountID),
		ServiceName: core.StringPtr(serviceName),
	}
}

// SetDisplayName : Allow user to set DisplayName
func (_options *CreateRoleOptions) SetDisplayName(displayName string) *CreateRoleOptions {
	_options.DisplayName = core.StringPtr(displayName)
	return _options
}

// SetActions : Allow user to set Actions
func (_options *CreateRoleOptions) SetActions(actions []string) *CreateRoleOptions {
	_options.Actions = actions
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateRoleOptions) SetName(name string) *CreateRoleOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *CreateRoleOptions) SetAccountID(accountID string) *CreateRoleOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetServiceName : Allow user to set ServiceName
func (_options *CreateRoleOptions) SetServiceName(serviceName string) *CreateRoleOptions {
	_options.ServiceName = core.StringPtr(serviceName)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateRoleOptions) SetDescription(description string) *CreateRoleOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *CreateRoleOptions) SetAcceptLanguage(acceptLanguage string) *CreateRoleOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateRoleOptions) SetHeaders(param map[string]string) *CreateRoleOptions {
	options.Headers = param
	return options
}

// DeletePolicyOptions : The DeletePolicy options.
type DeletePolicyOptions struct {
	// The policy ID.
	PolicyID *string `json:"policy_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeletePolicyOptions : Instantiate DeletePolicyOptions
func (*IamPolicyManagementV1) NewDeletePolicyOptions(policyID string) *DeletePolicyOptions {
	return &DeletePolicyOptions{
		PolicyID: core.StringPtr(policyID),
	}
}

// SetPolicyID : Allow user to set PolicyID
func (_options *DeletePolicyOptions) SetPolicyID(policyID string) *DeletePolicyOptions {
	_options.PolicyID = core.StringPtr(policyID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePolicyOptions) SetHeaders(param map[string]string) *DeletePolicyOptions {
	options.Headers = param
	return options
}

// DeleteRoleOptions : The DeleteRole options.
type DeleteRoleOptions struct {
	// The role ID.
	RoleID *string `json:"role_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteRoleOptions : Instantiate DeleteRoleOptions
func (*IamPolicyManagementV1) NewDeleteRoleOptions(roleID string) *DeleteRoleOptions {
	return &DeleteRoleOptions{
		RoleID: core.StringPtr(roleID),
	}
}

// SetRoleID : Allow user to set RoleID
func (_options *DeleteRoleOptions) SetRoleID(roleID string) *DeleteRoleOptions {
	_options.RoleID = core.StringPtr(roleID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteRoleOptions) SetHeaders(param map[string]string) *DeleteRoleOptions {
	options.Headers = param
	return options
}

// GetPolicyOptions : The GetPolicy options.
type GetPolicyOptions struct {
	// The policy ID.
	PolicyID *string `json:"policy_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPolicyOptions : Instantiate GetPolicyOptions
func (*IamPolicyManagementV1) NewGetPolicyOptions(policyID string) *GetPolicyOptions {
	return &GetPolicyOptions{
		PolicyID: core.StringPtr(policyID),
	}
}

// SetPolicyID : Allow user to set PolicyID
func (_options *GetPolicyOptions) SetPolicyID(policyID string) *GetPolicyOptions {
	_options.PolicyID = core.StringPtr(policyID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPolicyOptions) SetHeaders(param map[string]string) *GetPolicyOptions {
	options.Headers = param
	return options
}

// GetRoleOptions : The GetRole options.
type GetRoleOptions struct {
	// The role ID.
	RoleID *string `json:"role_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRoleOptions : Instantiate GetRoleOptions
func (*IamPolicyManagementV1) NewGetRoleOptions(roleID string) *GetRoleOptions {
	return &GetRoleOptions{
		RoleID: core.StringPtr(roleID),
	}
}

// SetRoleID : Allow user to set RoleID
func (_options *GetRoleOptions) SetRoleID(roleID string) *GetRoleOptions {
	_options.RoleID = core.StringPtr(roleID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetRoleOptions) SetHeaders(param map[string]string) *GetRoleOptions {
	options.Headers = param
	return options
}

// ListPoliciesOptions : The ListPolicies options.
type ListPoliciesOptions struct {
	// The account GUID in which the policies belong to.
	AccountID *string `json:"account_id" validate:"required"`

	// Language code for translations
	// * `default` - English
	// * `de` -  German (Standard)
	// * `en` - English
	// * `es` - Spanish (Spain)
	// * `fr` - French (Standard)
	// * `it` - Italian (Standard)
	// * `ja` - Japanese
	// * `ko` - Korean
	// * `pt-br` - Portuguese (Brazil)
	// * `zh-cn` - Chinese (Simplified, PRC)
	// * `zh-tw` - (Chinese, Taiwan).
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Optional IAM ID used to identify the subject.
	IamID *string `json:"iam_id,omitempty"`

	// Optional access group id.
	AccessGroupID *string `json:"access_group_id,omitempty"`

	// Optional type of policy.
	Type *string `json:"type,omitempty"`

	// Optional type of service.
	ServiceType *string `json:"service_type,omitempty"`

	// Optional name of the access management tag in the policy.
	TagName *string `json:"tag_name,omitempty"`

	// Optional value of the access management tag in the policy.
	TagValue *string `json:"tag_value,omitempty"`

	// Optional top level policy field to sort results. Ascending sort is default. Descending sort available by prepending
	// '-' to field. Example '-last_modified_at'.
	Sort *string `json:"sort,omitempty"`

	// Include additional data per policy returned
	// * `include_last_permit` - returns details of when the policy last granted a permit decision and the number of times
	// it has done so
	// * `display` - returns the list of all actions included in each of the policy roles.
	Format *string `json:"format,omitempty"`

	// The state of the policy.
	// * `active` - returns active policies
	// * `deleted` - returns non-active policies.
	State *string `json:"state,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListPoliciesOptions.Type property.
// Optional type of policy.
const (
	ListPoliciesOptionsTypeAccessConst = "access"
	ListPoliciesOptionsTypeAuthorizationConst = "authorization"
)

// Constants associated with the ListPoliciesOptions.ServiceType property.
// Optional type of service.
const (
	ListPoliciesOptionsServiceTypePlatformServiceConst = "platform_service"
	ListPoliciesOptionsServiceTypeServiceConst = "service"
)

// Constants associated with the ListPoliciesOptions.Sort property.
// Optional top level policy field to sort results. Ascending sort is default. Descending sort available by prepending
// '-' to field. Example '-last_modified_at'.
const (
	ListPoliciesOptionsSortCreatedAtConst = "created_at"
	ListPoliciesOptionsSortCreatedByIDConst = "created_by_id"
	ListPoliciesOptionsSortHrefConst = "href"
	ListPoliciesOptionsSortIDConst = "id"
	ListPoliciesOptionsSortLastModifiedAtConst = "last_modified_at"
	ListPoliciesOptionsSortLastModifiedByIDConst = "last_modified_by_id"
	ListPoliciesOptionsSortStateConst = "state"
	ListPoliciesOptionsSortTypeConst = "type"
)

// Constants associated with the ListPoliciesOptions.Format property.
// Include additional data per policy returned
// * `include_last_permit` - returns details of when the policy last granted a permit decision and the number of times
// it has done so
// * `display` - returns the list of all actions included in each of the policy roles.
const (
	ListPoliciesOptionsFormatDisplayConst = "display"
	ListPoliciesOptionsFormatIncludeLastPermitConst = "include_last_permit"
)

// Constants associated with the ListPoliciesOptions.State property.
// The state of the policy.
// * `active` - returns active policies
// * `deleted` - returns non-active policies.
const (
	ListPoliciesOptionsStateActiveConst = "active"
	ListPoliciesOptionsStateDeletedConst = "deleted"
)

// NewListPoliciesOptions : Instantiate ListPoliciesOptions
func (*IamPolicyManagementV1) NewListPoliciesOptions(accountID string) *ListPoliciesOptions {
	return &ListPoliciesOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *ListPoliciesOptions) SetAccountID(accountID string) *ListPoliciesOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *ListPoliciesOptions) SetAcceptLanguage(acceptLanguage string) *ListPoliciesOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetIamID : Allow user to set IamID
func (_options *ListPoliciesOptions) SetIamID(iamID string) *ListPoliciesOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetAccessGroupID : Allow user to set AccessGroupID
func (_options *ListPoliciesOptions) SetAccessGroupID(accessGroupID string) *ListPoliciesOptions {
	_options.AccessGroupID = core.StringPtr(accessGroupID)
	return _options
}

// SetType : Allow user to set Type
func (_options *ListPoliciesOptions) SetType(typeVar string) *ListPoliciesOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetServiceType : Allow user to set ServiceType
func (_options *ListPoliciesOptions) SetServiceType(serviceType string) *ListPoliciesOptions {
	_options.ServiceType = core.StringPtr(serviceType)
	return _options
}

// SetTagName : Allow user to set TagName
func (_options *ListPoliciesOptions) SetTagName(tagName string) *ListPoliciesOptions {
	_options.TagName = core.StringPtr(tagName)
	return _options
}

// SetTagValue : Allow user to set TagValue
func (_options *ListPoliciesOptions) SetTagValue(tagValue string) *ListPoliciesOptions {
	_options.TagValue = core.StringPtr(tagValue)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListPoliciesOptions) SetSort(sort string) *ListPoliciesOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetFormat : Allow user to set Format
func (_options *ListPoliciesOptions) SetFormat(format string) *ListPoliciesOptions {
	_options.Format = core.StringPtr(format)
	return _options
}

// SetState : Allow user to set State
func (_options *ListPoliciesOptions) SetState(state string) *ListPoliciesOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListPoliciesOptions) SetHeaders(param map[string]string) *ListPoliciesOptions {
	options.Headers = param
	return options
}

// ListRolesOptions : The ListRoles options.
type ListRolesOptions struct {
	// Language code for translations
	// * `default` - English
	// * `de` -  German (Standard)
	// * `en` - English
	// * `es` - Spanish (Spain)
	// * `fr` - French (Standard)
	// * `it` - Italian (Standard)
	// * `ja` - Japanese
	// * `ko` - Korean
	// * `pt-br` - Portuguese (Brazil)
	// * `zh-cn` - Chinese (Simplified, PRC)
	// * `zh-tw` - (Chinese, Taiwan).
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Optional account GUID in which the roles belong to.
	AccountID *string `json:"account_id,omitempty"`

	// Optional name of IAM enabled service.
	ServiceName *string `json:"service_name,omitempty"`

	// Optional name of source IAM enabled service.
	SourceServiceName *string `json:"source_service_name,omitempty"`

	// Optional Policy Type.
	PolicyType *string `json:"policy_type,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListRolesOptions : Instantiate ListRolesOptions
func (*IamPolicyManagementV1) NewListRolesOptions() *ListRolesOptions {
	return &ListRolesOptions{}
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *ListRolesOptions) SetAcceptLanguage(acceptLanguage string) *ListRolesOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *ListRolesOptions) SetAccountID(accountID string) *ListRolesOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetServiceName : Allow user to set ServiceName
func (_options *ListRolesOptions) SetServiceName(serviceName string) *ListRolesOptions {
	_options.ServiceName = core.StringPtr(serviceName)
	return _options
}

// SetSourceServiceName : Allow user to set SourceServiceName
func (_options *ListRolesOptions) SetSourceServiceName(sourceServiceName string) *ListRolesOptions {
	_options.SourceServiceName = core.StringPtr(sourceServiceName)
	return _options
}

// SetPolicyType : Allow user to set PolicyType
func (_options *ListRolesOptions) SetPolicyType(policyType string) *ListRolesOptions {
	_options.PolicyType = core.StringPtr(policyType)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListRolesOptions) SetHeaders(param map[string]string) *ListRolesOptions {
	options.Headers = param
	return options
}

// PatchPolicyOptions : The PatchPolicy options.
type PatchPolicyOptions struct {
	// The policy ID.
	PolicyID *string `json:"policy_id" validate:"required,ne="`

	// The revision number for updating a policy and must match the ETag value of the existing policy. The Etag can be
	// retrieved using the GET /v1/policies/{policy_id} API and looking at the ETag response header.
	IfMatch *string `json:"If-Match" validate:"required"`

	// The policy state.
	State *string `json:"state,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the PatchPolicyOptions.State property.
// The policy state.
const (
	PatchPolicyOptionsStateActiveConst = "active"
	PatchPolicyOptionsStateDeletedConst = "deleted"
)

// NewPatchPolicyOptions : Instantiate PatchPolicyOptions
func (*IamPolicyManagementV1) NewPatchPolicyOptions(policyID string, ifMatch string) *PatchPolicyOptions {
	return &PatchPolicyOptions{
		PolicyID: core.StringPtr(policyID),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetPolicyID : Allow user to set PolicyID
func (_options *PatchPolicyOptions) SetPolicyID(policyID string) *PatchPolicyOptions {
	_options.PolicyID = core.StringPtr(policyID)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *PatchPolicyOptions) SetIfMatch(ifMatch string) *PatchPolicyOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetState : Allow user to set State
func (_options *PatchPolicyOptions) SetState(state string) *PatchPolicyOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PatchPolicyOptions) SetHeaders(param map[string]string) *PatchPolicyOptions {
	options.Headers = param
	return options
}

// UpdatePolicyOptions : The UpdatePolicy options.
type UpdatePolicyOptions struct {
	// The policy ID.
	PolicyID *string `json:"policy_id" validate:"required,ne="`

	// The revision number for updating a policy and must match the ETag value of the existing policy. The Etag can be
	// retrieved using the GET /v1/policies/{policy_id} API and looking at the ETag response header.
	IfMatch *string `json:"If-Match" validate:"required"`

	// The policy type; either 'access' or 'authorization'.
	Type *string `json:"type" validate:"required"`

	// The subjects associated with a policy.
	Subjects []PolicySubject `json:"subjects" validate:"required"`

	// A set of role cloud resource names (CRNs) granted by the policy.
	Roles []PolicyRole `json:"roles" validate:"required"`

	// The resources associated with a policy.
	Resources []PolicyResource `json:"resources" validate:"required"`

	// Customer-defined description.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdatePolicyOptions : Instantiate UpdatePolicyOptions
func (*IamPolicyManagementV1) NewUpdatePolicyOptions(policyID string, ifMatch string, typeVar string, subjects []PolicySubject, roles []PolicyRole, resources []PolicyResource) *UpdatePolicyOptions {
	return &UpdatePolicyOptions{
		PolicyID: core.StringPtr(policyID),
		IfMatch: core.StringPtr(ifMatch),
		Type: core.StringPtr(typeVar),
		Subjects: subjects,
		Roles: roles,
		Resources: resources,
	}
}

// SetPolicyID : Allow user to set PolicyID
func (_options *UpdatePolicyOptions) SetPolicyID(policyID string) *UpdatePolicyOptions {
	_options.PolicyID = core.StringPtr(policyID)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *UpdatePolicyOptions) SetIfMatch(ifMatch string) *UpdatePolicyOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetType : Allow user to set Type
func (_options *UpdatePolicyOptions) SetType(typeVar string) *UpdatePolicyOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetSubjects : Allow user to set Subjects
func (_options *UpdatePolicyOptions) SetSubjects(subjects []PolicySubject) *UpdatePolicyOptions {
	_options.Subjects = subjects
	return _options
}

// SetRoles : Allow user to set Roles
func (_options *UpdatePolicyOptions) SetRoles(roles []PolicyRole) *UpdatePolicyOptions {
	_options.Roles = roles
	return _options
}

// SetResources : Allow user to set Resources
func (_options *UpdatePolicyOptions) SetResources(resources []PolicyResource) *UpdatePolicyOptions {
	_options.Resources = resources
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdatePolicyOptions) SetDescription(description string) *UpdatePolicyOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdatePolicyOptions) SetHeaders(param map[string]string) *UpdatePolicyOptions {
	options.Headers = param
	return options
}

// UpdateRoleOptions : The UpdateRole options.
type UpdateRoleOptions struct {
	// The role ID.
	RoleID *string `json:"role_id" validate:"required,ne="`

	// The revision number for updating a role and must match the ETag value of the existing role. The Etag can be
	// retrieved using the GET /v2/roles/{role_id} API and looking at the ETag response header.
	IfMatch *string `json:"If-Match" validate:"required"`

	// The display name of the role that is shown in the console.
	DisplayName *string `json:"display_name,omitempty"`

	// The description of the role.
	Description *string `json:"description,omitempty"`

	// The actions of the role. Please refer to [IAM roles and
	// actions](https://cloud.ibm.com/docs/account?topic=account-iam-service-roles-actions).
	Actions []string `json:"actions,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateRoleOptions : Instantiate UpdateRoleOptions
func (*IamPolicyManagementV1) NewUpdateRoleOptions(roleID string, ifMatch string) *UpdateRoleOptions {
	return &UpdateRoleOptions{
		RoleID: core.StringPtr(roleID),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetRoleID : Allow user to set RoleID
func (_options *UpdateRoleOptions) SetRoleID(roleID string) *UpdateRoleOptions {
	_options.RoleID = core.StringPtr(roleID)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *UpdateRoleOptions) SetIfMatch(ifMatch string) *UpdateRoleOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetDisplayName : Allow user to set DisplayName
func (_options *UpdateRoleOptions) SetDisplayName(displayName string) *UpdateRoleOptions {
	_options.DisplayName = core.StringPtr(displayName)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateRoleOptions) SetDescription(description string) *UpdateRoleOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetActions : Allow user to set Actions
func (_options *UpdateRoleOptions) SetActions(actions []string) *UpdateRoleOptions {
	_options.Actions = actions
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateRoleOptions) SetHeaders(param map[string]string) *UpdateRoleOptions {
	options.Headers = param
	return options
}

// CustomRole : An additional set of properties associated with a role.
type CustomRole struct {
	// The role ID. Composed of hexadecimal characters.
	ID *string `json:"id,omitempty"`

	// The display name of the role that is shown in the console.
	DisplayName *string `json:"display_name,omitempty"`

	// The description of the role.
	Description *string `json:"description,omitempty"`

	// The actions of the role. Please refer to [IAM roles and
	// actions](https://cloud.ibm.com/docs/account?topic=account-iam-service-roles-actions).
	Actions []string `json:"actions,omitempty"`

	// The role Cloud Resource Name (CRN). Example CRN:
	// 'crn:v1:ibmcloud:public:iam-access-management::a/exampleAccountId::customRole:ExampleRoleName'.
	CRN *string `json:"crn,omitempty"`

	// The name of the role that is used in the CRN. Can only be alphanumeric and has to be capitalized.
	Name *string `json:"name,omitempty"`

	// The account GUID.
	AccountID *string `json:"account_id,omitempty"`

	// The service name.
	ServiceName *string `json:"service_name,omitempty"`

	// The UTC timestamp when the role was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The iam ID of the entity that created the role.
	CreatedByID *string `json:"created_by_id,omitempty"`

	// The UTC timestamp when the role was last modified.
	LastModifiedAt *strfmt.DateTime `json:"last_modified_at,omitempty"`

	// The iam ID of the entity that last modified the policy.
	LastModifiedByID *string `json:"last_modified_by_id,omitempty"`

	// The href link back to the role.
	Href *string `json:"href,omitempty"`
}

// UnmarshalCustomRole unmarshals an instance of CustomRole from the specified map of raw messages.
func UnmarshalCustomRole(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CustomRole)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.CRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_name", &obj.ServiceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by_id", &obj.CreatedByID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_modified_at", &obj.LastModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_modified_by_id", &obj.LastModifiedByID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Policy : The core set of properties associated with a policy.
type Policy struct {
	// The policy ID.
	ID *string `json:"id,omitempty"`

	// The policy type; either 'access' or 'authorization'.
	Type *string `json:"type,omitempty"`

	// Customer-defined description.
	Description *string `json:"description,omitempty"`

	// The subjects associated with a policy.
	Subjects []PolicySubject `json:"subjects,omitempty"`

	// A set of role cloud resource names (CRNs) granted by the policy.
	Roles []PolicyRole `json:"roles,omitempty"`

	// The resources associated with a policy.
	Resources []PolicyResource `json:"resources,omitempty"`

	// The href link back to the policy.
	Href *string `json:"href,omitempty"`

	// The UTC timestamp when the policy was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The iam ID of the entity that created the policy.
	CreatedByID *string `json:"created_by_id,omitempty"`

	// The UTC timestamp when the policy was last modified.
	LastModifiedAt *strfmt.DateTime `json:"last_modified_at,omitempty"`

	// The iam ID of the entity that last modified the policy.
	LastModifiedByID *string `json:"last_modified_by_id,omitempty"`

	// The policy state.
	State *string `json:"state,omitempty"`
}

// Constants associated with the Policy.State property.
// The policy state.
const (
	PolicyStateActiveConst = "active"
	PolicyStateDeletedConst = "deleted"
)

// UnmarshalPolicy unmarshals an instance of Policy from the specified map of raw messages.
func UnmarshalPolicy(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Policy)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "subjects", &obj.Subjects, UnmarshalPolicySubject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "roles", &obj.Roles, UnmarshalPolicyRole)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalPolicyResource)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by_id", &obj.CreatedByID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_modified_at", &obj.LastModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_modified_by_id", &obj.LastModifiedByID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicyList : A collection of policies.
type PolicyList struct {
	// List of policies.
	Policies []Policy `json:"policies,omitempty"`
}

// UnmarshalPolicyList unmarshals an instance of PolicyList from the specified map of raw messages.
func UnmarshalPolicyList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyList)
	err = core.UnmarshalModel(m, "policies", &obj.Policies, UnmarshalPolicy)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicyResource : The attributes of the resource. Note that only one resource is allowed in a policy.
type PolicyResource struct {
	// List of resource attributes.
	Attributes []ResourceAttribute `json:"attributes,omitempty"`

	// List of access management tags.
	Tags []ResourceTag `json:"tags,omitempty"`
}

// UnmarshalPolicyResource unmarshals an instance of PolicyResource from the specified map of raw messages.
func UnmarshalPolicyResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyResource)
	err = core.UnmarshalModel(m, "attributes", &obj.Attributes, UnmarshalResourceAttribute)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tags", &obj.Tags, UnmarshalResourceTag)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicyRole : A role associated with a policy.
type PolicyRole struct {
	// The role Cloud Resource Name (CRN) granted by the policy. Example CRN: 'crn:v1:bluemix:public:iam::::role:Editor'.
	RoleID *string `json:"role_id" validate:"required"`

	// The display name of the role.
	DisplayName *string `json:"display_name,omitempty"`

	// The description of the role.
	Description *string `json:"description,omitempty"`
}

// NewPolicyRole : Instantiate PolicyRole (Generic Model Constructor)
func (*IamPolicyManagementV1) NewPolicyRole(roleID string) (_model *PolicyRole, err error) {
	_model = &PolicyRole{
		RoleID: core.StringPtr(roleID),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalPolicyRole unmarshals an instance of PolicyRole from the specified map of raw messages.
func UnmarshalPolicyRole(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyRole)
	err = core.UnmarshalPrimitive(m, "role_id", &obj.RoleID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicySubject : The subject attribute values that must match in order for this policy to apply in a permission decision.
type PolicySubject struct {
	// List of subject attributes.
	Attributes []SubjectAttribute `json:"attributes,omitempty"`
}

// UnmarshalPolicySubject unmarshals an instance of PolicySubject from the specified map of raw messages.
func UnmarshalPolicySubject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicySubject)
	err = core.UnmarshalModel(m, "attributes", &obj.Attributes, UnmarshalSubjectAttribute)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceAttribute : An attribute associated with a resource.
type ResourceAttribute struct {
	// The name of an attribute.
	Name *string `json:"name" validate:"required"`

	// The value of an attribute.
	Value *string `json:"value" validate:"required"`

	// The operator of an attribute.
	Operator *string `json:"operator,omitempty"`
}

// NewResourceAttribute : Instantiate ResourceAttribute (Generic Model Constructor)
func (*IamPolicyManagementV1) NewResourceAttribute(name string, value string) (_model *ResourceAttribute, err error) {
	_model = &ResourceAttribute{
		Name: core.StringPtr(name),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalResourceAttribute unmarshals an instance of ResourceAttribute from the specified map of raw messages.
func UnmarshalResourceAttribute(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceAttribute)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceTag : A tag associated with a resource.
type ResourceTag struct {
	// The name of an access management tag.
	Name *string `json:"name" validate:"required"`

	// The value of an access management tag.
	Value *string `json:"value" validate:"required"`

	// The operator of an access management tag.
	Operator *string `json:"operator,omitempty"`
}

// NewResourceTag : Instantiate ResourceTag (Generic Model Constructor)
func (*IamPolicyManagementV1) NewResourceTag(name string, value string) (_model *ResourceTag, err error) {
	_model = &ResourceTag{
		Name: core.StringPtr(name),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalResourceTag unmarshals an instance of ResourceTag from the specified map of raw messages.
func UnmarshalResourceTag(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceTag)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Role : A role resource.
type Role struct {
	// The display name of the role that is shown in the console.
	DisplayName *string `json:"display_name,omitempty"`

	// The description of the role.
	Description *string `json:"description,omitempty"`

	// The actions of the role. Please refer to [IAM roles and
	// actions](https://cloud.ibm.com/docs/account?topic=account-iam-service-roles-actions).
	Actions []string `json:"actions,omitempty"`

	// The role Cloud Resource Name (CRN). Example CRN:
	// 'crn:v1:ibmcloud:public:iam-access-management::a/exampleAccountId::customRole:ExampleRoleName'.
	CRN *string `json:"crn,omitempty"`
}

// UnmarshalRole unmarshals an instance of Role from the specified map of raw messages.
func UnmarshalRole(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Role)
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.CRN)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RoleList : A collection of roles returned by the 'list roles' operation.
type RoleList struct {
	// List of custom roles.
	CustomRoles []CustomRole `json:"custom_roles,omitempty"`

	// List of service roles.
	ServiceRoles []Role `json:"service_roles,omitempty"`

	// List of system roles.
	SystemRoles []Role `json:"system_roles,omitempty"`
}

// UnmarshalRoleList unmarshals an instance of RoleList from the specified map of raw messages.
func UnmarshalRoleList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RoleList)
	err = core.UnmarshalModel(m, "custom_roles", &obj.CustomRoles, UnmarshalCustomRole)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "service_roles", &obj.ServiceRoles, UnmarshalRole)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system_roles", &obj.SystemRoles, UnmarshalRole)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SubjectAttribute : An attribute associated with a subject.
type SubjectAttribute struct {
	// The name of an attribute.
	Name *string `json:"name" validate:"required"`

	// The value of an attribute.
	Value *string `json:"value" validate:"required"`
}

// NewSubjectAttribute : Instantiate SubjectAttribute (Generic Model Constructor)
func (*IamPolicyManagementV1) NewSubjectAttribute(name string, value string) (_model *SubjectAttribute, err error) {
	_model = &SubjectAttribute{
		Name: core.StringPtr(name),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalSubjectAttribute unmarshals an instance of SubjectAttribute from the specified map of raw messages.
func UnmarshalSubjectAttribute(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SubjectAttribute)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
