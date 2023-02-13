# \RepositoryApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesPost**](RepositoryApplyApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesPost) | **Post** /apis/integrations.k8smgmt.io/v3/projects/{project}/repositories | 
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysPost**](RepositoryApplyApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysPost) | **Post** /apis/integrations.k8smgmt.io/v3/projects/{project}/repositorys | 



## ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesPost

> Repository ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesPost(ctx, project).Repository(repository).Execute()





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
    repository := *openapiclient.NewRepository("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewRepositorySpec()) // Repository | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesPost(context.Background(), project).Repository(repository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesPost`: Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoryApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repository** | [**Repository**](Repository.md) | schema of the resource to apply | 

### Return type

[**Repository**](Repository.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysPost

> Repository ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysPost(ctx, project).Repository(repository).Execute()





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
    repository := *openapiclient.NewRepository("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewRepositorySpec()) // Repository | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysPost(context.Background(), project).Repository(repository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysPost`: Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoryApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repository** | [**Repository**](Repository.md) | schema of the resource to apply | 

### Return type

[**Repository**](Repository.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

