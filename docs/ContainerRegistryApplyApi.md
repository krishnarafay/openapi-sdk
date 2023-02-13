# \ContainerRegistryApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesPost**](ContainerRegistryApplyApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesPost) | **Post** /apis/integrations.k8smgmt.io/v3/projects/{project}/containerregistries | 
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysPost**](ContainerRegistryApplyApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysPost) | **Post** /apis/integrations.k8smgmt.io/v3/projects/{project}/containerregistrys | 



## ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesPost

> ContainerRegistry ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesPost(ctx, project).ContainerRegistry(containerRegistry).Execute()





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
    containerRegistry := *openapiclient.NewContainerRegistry("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewContainerRegistrySpec()) // ContainerRegistry | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerRegistryApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesPost(context.Background(), project).ContainerRegistry(containerRegistry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerRegistryApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesPost`: ContainerRegistry
    fmt.Fprintf(os.Stdout, "Response from `ContainerRegistryApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerRegistry** | [**ContainerRegistry**](ContainerRegistry.md) | schema of the resource to apply | 

### Return type

[**ContainerRegistry**](ContainerRegistry.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysPost

> ContainerRegistry ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysPost(ctx, project).ContainerRegistry(containerRegistry).Execute()





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
    containerRegistry := *openapiclient.NewContainerRegistry("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewContainerRegistrySpec()) // ContainerRegistry | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerRegistryApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysPost(context.Background(), project).ContainerRegistry(containerRegistry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerRegistryApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysPost`: ContainerRegistry
    fmt.Fprintf(os.Stdout, "Response from `ContainerRegistryApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerRegistry** | [**ContainerRegistry**](ContainerRegistry.md) | schema of the resource to apply | 

### Return type

[**ContainerRegistry**](ContainerRegistry.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

