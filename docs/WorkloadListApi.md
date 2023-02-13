# \WorkloadListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsGet**](WorkloadListApi.md#ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsGet) | **Get** /apis/apps.k8smgmt.io/v3/projects/{project}/workloads | 



## ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsGet

> WorkloadList ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsGet(ctx, project).Execute()





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
    resp, r, err := apiClient.WorkloadListApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadListApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsGet`: WorkloadList
    fmt.Fprintf(os.Stdout, "Response from `WorkloadListApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsK8smgmtIoV3ProjectsProjectWorkloadsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkloadList**](WorkloadList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

