# \AgentDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameGet**](AgentDetailApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameGet) | **Get** /apis/gitops.k8smgmt.io/v3/projects/{project}/agents/{name} | 
[**ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameStatusGet**](AgentDetailApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameStatusGet) | **Get** /apis/gitops.k8smgmt.io/v3/projects/{project}/agents/{name}/status | 



## ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameGet

> Agent ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameGet(ctx, project, name).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    project := "project_example" // string | project of the resource
    name := "name_example" // string | name of the resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameGet`: Agent
    fmt.Fprintf(os.Stdout, "Response from `AgentDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameStatusGet

> Agent ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameStatusGet(ctx, project, name).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    project := "project_example" // string | project of the resource
    name := "name_example" // string | name of the resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameStatusGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameStatusGet`: Agent
    fmt.Fprintf(os.Stdout, "Response from `AgentDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectAgentsNameStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

