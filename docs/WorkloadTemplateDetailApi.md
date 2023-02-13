# \WorkloadTemplateDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameArtifactsGet**](WorkloadTemplateDetailApi.md#ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameArtifactsGet) | **Get** /apis/apps.k8smgmt.io/v3/projects/{project}/workloadtemplates/{name}/artifacts | 
[**ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameGet**](WorkloadTemplateDetailApi.md#ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameGet) | **Get** /apis/apps.k8smgmt.io/v3/projects/{project}/workloadtemplates/{name} | 



## ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameArtifactsGet

> WorkloadTemplate ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.WorkloadTemplateDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplateDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameArtifactsGet`: WorkloadTemplate
    fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplateDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkloadTemplate**](WorkloadTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameGet

> WorkloadTemplate ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.WorkloadTemplateDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplateDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameGet`: WorkloadTemplate
    fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplateDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkloadTemplate**](WorkloadTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

