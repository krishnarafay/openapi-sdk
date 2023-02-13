# \RepositoryListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesGet**](RepositoryListApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/repositories | 
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysGet**](RepositoryListApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/repositorys | 



## ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesGet

> RepositoryList ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesGet(ctx, project).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesGet`: RepositoryList
    fmt.Fprintf(os.Stdout, "Response from `RepositoryListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RepositoryList**](RepositoryList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysGet

> RepositoryList ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysGet(ctx, project).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysGet`: RepositoryList
    fmt.Fprintf(os.Stdout, "Response from `RepositoryListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RepositoryList**](RepositoryList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

