//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscheduler

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// JobCollectionsClient contains the methods for the JobCollections group.
// Don't use this type directly, use NewJobCollectionsClient() instead.
type JobCollectionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewJobCollectionsClient creates a new instance of JobCollectionsClient with the specified values.
func NewJobCollectionsClient(con *arm.Connection, subscriptionID string) *JobCollectionsClient {
	return &JobCollectionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Provisions a new job collection or updates an existing job collection.
// If the operation fails it returns a generic error.
func (client *JobCollectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition, options *JobCollectionsCreateOrUpdateOptions) (JobCollectionsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, jobCollectionName, jobCollection, options)
	if err != nil {
		return JobCollectionsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobCollectionsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return JobCollectionsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *JobCollectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition, options *JobCollectionsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, runtime.MarshalAsJSON(req, jobCollection)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *JobCollectionsClient) createOrUpdateHandleResponse(resp *http.Response) (JobCollectionsCreateOrUpdateResponse, error) {
	result := JobCollectionsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobCollectionDefinition); err != nil {
		return JobCollectionsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *JobCollectionsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Deletes a job collection.
// If the operation fails it returns a generic error.
func (client *JobCollectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsBeginDeleteOptions) (JobCollectionsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, jobCollectionName, options)
	if err != nil {
		return JobCollectionsDeletePollerResponse{}, err
	}
	result := JobCollectionsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("JobCollectionsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return JobCollectionsDeletePollerResponse{}, err
	}
	result.Poller = &JobCollectionsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a job collection.
// If the operation fails it returns a generic error.
func (client *JobCollectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, jobCollectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *JobCollectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *JobCollectionsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDisable - Disables all of the jobs in the job collection.
// If the operation fails it returns a generic error.
func (client *JobCollectionsClient) BeginDisable(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsBeginDisableOptions) (JobCollectionsDisablePollerResponse, error) {
	resp, err := client.disable(ctx, resourceGroupName, jobCollectionName, options)
	if err != nil {
		return JobCollectionsDisablePollerResponse{}, err
	}
	result := JobCollectionsDisablePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("JobCollectionsClient.Disable", "", resp, client.pl, client.disableHandleError)
	if err != nil {
		return JobCollectionsDisablePollerResponse{}, err
	}
	result.Poller = &JobCollectionsDisablePoller{
		pt: pt,
	}
	return result, nil
}

// Disable - Disables all of the jobs in the job collection.
// If the operation fails it returns a generic error.
func (client *JobCollectionsClient) disable(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsBeginDisableOptions) (*http.Response, error) {
	req, err := client.disableCreateRequest(ctx, resourceGroupName, jobCollectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.disableHandleError(resp)
	}
	return resp, nil
}

// disableCreateRequest creates the Disable request.
func (client *JobCollectionsClient) disableCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsBeginDisableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/disable"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// disableHandleError handles the Disable error response.
func (client *JobCollectionsClient) disableHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginEnable - Enables all of the jobs in the job collection.
// If the operation fails it returns a generic error.
func (client *JobCollectionsClient) BeginEnable(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsBeginEnableOptions) (JobCollectionsEnablePollerResponse, error) {
	resp, err := client.enable(ctx, resourceGroupName, jobCollectionName, options)
	if err != nil {
		return JobCollectionsEnablePollerResponse{}, err
	}
	result := JobCollectionsEnablePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("JobCollectionsClient.Enable", "", resp, client.pl, client.enableHandleError)
	if err != nil {
		return JobCollectionsEnablePollerResponse{}, err
	}
	result.Poller = &JobCollectionsEnablePoller{
		pt: pt,
	}
	return result, nil
}

// Enable - Enables all of the jobs in the job collection.
// If the operation fails it returns a generic error.
func (client *JobCollectionsClient) enable(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsBeginEnableOptions) (*http.Response, error) {
	req, err := client.enableCreateRequest(ctx, resourceGroupName, jobCollectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.enableHandleError(resp)
	}
	return resp, nil
}

// enableCreateRequest creates the Enable request.
func (client *JobCollectionsClient) enableCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsBeginEnableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/enable"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// enableHandleError handles the Enable error response.
func (client *JobCollectionsClient) enableHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a job collection.
// If the operation fails it returns a generic error.
func (client *JobCollectionsClient) Get(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsGetOptions) (JobCollectionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, jobCollectionName, options)
	if err != nil {
		return JobCollectionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobCollectionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobCollectionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobCollectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobCollectionsClient) getHandleResponse(resp *http.Response) (JobCollectionsGetResponse, error) {
	result := JobCollectionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobCollectionDefinition); err != nil {
		return JobCollectionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *JobCollectionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByResourceGroup - Gets all job collections under specified resource group.
// If the operation fails it returns a generic error.
func (client *JobCollectionsClient) ListByResourceGroup(resourceGroupName string, options *JobCollectionsListByResourceGroupOptions) *JobCollectionsListByResourceGroupPager {
	return &JobCollectionsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp JobCollectionsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobCollectionListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *JobCollectionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *JobCollectionsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *JobCollectionsClient) listByResourceGroupHandleResponse(resp *http.Response) (JobCollectionsListByResourceGroupResponse, error) {
	result := JobCollectionsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobCollectionListResult); err != nil {
		return JobCollectionsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *JobCollectionsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListBySubscription - Gets all job collections under specified subscription.
// If the operation fails it returns a generic error.
func (client *JobCollectionsClient) ListBySubscription(options *JobCollectionsListBySubscriptionOptions) *JobCollectionsListBySubscriptionPager {
	return &JobCollectionsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp JobCollectionsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobCollectionListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *JobCollectionsClient) listBySubscriptionCreateRequest(ctx context.Context, options *JobCollectionsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Scheduler/jobCollections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *JobCollectionsClient) listBySubscriptionHandleResponse(resp *http.Response) (JobCollectionsListBySubscriptionResponse, error) {
	result := JobCollectionsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobCollectionListResult); err != nil {
		return JobCollectionsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *JobCollectionsClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Patch - Patches an existing job collection.
// If the operation fails it returns a generic error.
func (client *JobCollectionsClient) Patch(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition, options *JobCollectionsPatchOptions) (JobCollectionsPatchResponse, error) {
	req, err := client.patchCreateRequest(ctx, resourceGroupName, jobCollectionName, jobCollection, options)
	if err != nil {
		return JobCollectionsPatchResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobCollectionsPatchResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobCollectionsPatchResponse{}, client.patchHandleError(resp)
	}
	return client.patchHandleResponse(resp)
}

// patchCreateRequest creates the Patch request.
func (client *JobCollectionsClient) patchCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition, options *JobCollectionsPatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, runtime.MarshalAsJSON(req, jobCollection)
}

// patchHandleResponse handles the Patch response.
func (client *JobCollectionsClient) patchHandleResponse(resp *http.Response) (JobCollectionsPatchResponse, error) {
	result := JobCollectionsPatchResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobCollectionDefinition); err != nil {
		return JobCollectionsPatchResponse{}, err
	}
	return result, nil
}

// patchHandleError handles the Patch error response.
func (client *JobCollectionsClient) patchHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
