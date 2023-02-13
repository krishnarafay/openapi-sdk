# \WorkloadDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameArtifactsGet**](WorkloadDetailApi.md#ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameArtifactsGet) | **Get** /apis/apps.k8smgmt.io/v3/projects/{project}/workloads/{name}/artifacts | 
[**ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameGet**](WorkloadDetailApi.md#ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameGet) | **Get** /apis/apps.k8smgmt.io/v3/projects/{project}/workloads/{name} | 
[**ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameStatusGet**](WorkloadDetailApi.md#ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameStatusGet) | **Get** /apis/apps.k8smgmt.io/v3/projects/{project}/workloads/{name}/status | 



## ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameArtifactsGet

> Workload ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.WorkloadDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameArtifactsGet`: Workload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Workload**](Workload.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameGet

> Workload ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.WorkloadDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameGet`: Workload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Workload**](Workload.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameStatusGet

> Workload ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameStatusGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.WorkloadDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameStatusGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameStatusGet`: Workload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadDetailApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsK8smgmtIoV3ProjectsProjectWorkloadsNameStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Workload**](Workload.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

