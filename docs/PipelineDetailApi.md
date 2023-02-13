# \PipelineDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameArtifactsGet**](PipelineDetailApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameArtifactsGet) | **Get** /apis/gitops.k8smgmt.io/v3/projects/{project}/pipelines/{name}/artifacts | 
[**ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameGet**](PipelineDetailApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameGet) | **Get** /apis/gitops.k8smgmt.io/v3/projects/{project}/pipelines/{name} | 
[**ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameStatusGet**](PipelineDetailApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameStatusGet) | **Get** /apis/gitops.k8smgmt.io/v3/projects/{project}/pipelines/{name}/status | 



## ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameArtifactsGet

> Pipeline ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.PipelineDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameArtifactsGet`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelineDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameGet

> Pipeline ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.PipelineDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameGet`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelineDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameStatusGet

> Pipeline ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameStatusGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.PipelineDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameStatusGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameStatusGet`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelineDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectPipelinesNameStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

