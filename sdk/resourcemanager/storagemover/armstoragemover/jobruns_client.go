//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstoragemover

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// JobRunsClient contains the methods for the JobRuns group.
// Don't use this type directly, use NewJobRunsClient() instead.
type JobRunsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewJobRunsClient creates a new instance of JobRunsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewJobRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobRunsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &JobRunsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets a Job Run resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - projectName - The name of the Project resource.
//   - jobDefinitionName - The name of the Job Definition resource.
//   - jobRunName - The name of the Job Run resource.
//   - options - JobRunsClientGetOptions contains the optional parameters for the JobRunsClient.Get method.
func (client *JobRunsClient) Get(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, jobRunName string, options *JobRunsClientGetOptions) (JobRunsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, storageMoverName, projectName, jobDefinitionName, jobRunName, options)
	if err != nil {
		return JobRunsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobRunsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobRunsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobRunsClient) getCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, jobRunName string, options *JobRunsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/projects/{projectName}/jobDefinitions/{jobDefinitionName}/jobRuns/{jobRunName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	if jobRunName == "" {
		return nil, errors.New("parameter jobRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobRunName}", url.PathEscape(jobRunName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobRunsClient) getHandleResponse(resp *http.Response) (JobRunsClientGetResponse, error) {
	result := JobRunsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobRun); err != nil {
		return JobRunsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all Job Runs in a Job Definition.
//
// Generated from API version 2022-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - projectName - The name of the Project resource.
//   - jobDefinitionName - The name of the Job Definition resource.
//   - options - JobRunsClientListOptions contains the optional parameters for the JobRunsClient.NewListPager method.
func (client *JobRunsClient) NewListPager(resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, options *JobRunsClientListOptions) *runtime.Pager[JobRunsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobRunsClientListResponse]{
		More: func(page JobRunsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobRunsClientListResponse) (JobRunsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, storageMoverName, projectName, jobDefinitionName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobRunsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return JobRunsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobRunsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *JobRunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, options *JobRunsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/projects/{projectName}/jobDefinitions/{jobDefinitionName}/jobRuns"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *JobRunsClient) listHandleResponse(resp *http.Response) (JobRunsClientListResponse, error) {
	result := JobRunsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobRunList); err != nil {
		return JobRunsClientListResponse{}, err
	}
	return result, nil
}