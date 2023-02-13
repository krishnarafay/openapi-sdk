# \OverrideDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameArtifactsGet**](OverrideDetailApi.md#ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameArtifactsGet) | **Get** /apis/apps.k8smgmt.io/v3/projects/{project}/overrides/{name}/artifacts | 
[**ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameGet**](OverrideDetailApi.md#ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameGet) | **Get** /apis/apps.k8smgmt.io/v3/projects/{project}/overrides/{name} | 



## ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameArtifactsGet

> Override ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.OverrideDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverrideDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameArtifactsGet`: Override
    fmt.Fprintf(os.Stdout, "Response from `OverrideDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsK8smgmtIoV3ProjectsProjectOverridesNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Override**](Override.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameGet

> Override ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.OverrideDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverrideDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameGet`: Override
    fmt.Fprintf(os.Stdout, "Response from `OverrideDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectOverridesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsK8smgmtIoV3ProjectsProjectOverridesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Override**](Override.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

