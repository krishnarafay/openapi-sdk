# \ContainerRegistryListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesGet**](ContainerRegistryListApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/containerregistries | 
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysGet**](ContainerRegistryListApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/containerregistrys | 



## ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesGet

> ContainerRegistryList ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesGet(ctx, project).Execute()





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
    resp, r, err := apiClient.ContainerRegistryListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerRegistryListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesGet`: ContainerRegistryList
    fmt.Fprintf(os.Stdout, "Response from `ContainerRegistryListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerRegistryList**](ContainerRegistryList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysGet

> ContainerRegistryList ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysGet(ctx, project).Execute()





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
    resp, r, err := apiClient.ContainerRegistryListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerRegistryListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysGet`: ContainerRegistryList
    fmt.Fprintf(os.Stdout, "Response from `ContainerRegistryListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectContainerregistrysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerRegistryList**](ContainerRegistryList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

