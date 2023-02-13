# \WorkloadApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsPost**](WorkloadApplyApi.md#ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsPost) | **Post** /apis/apps.k8smgmt.io/v3/projects/{project}/workloads | 



## ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsPost

> Workload ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsPost(ctx, project).Workload(workload).Execute()





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
    workload := *openapiclient.NewWorkload("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewWorkloadSpec()) // Workload | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkloadApplyApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsPost(context.Background(), project).Workload(workload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadApplyApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsPost`: Workload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadApplyApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsK8smgmtIoV3ProjectsProjectWorkloadsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workload** | [**Workload**](Workload.md) | schema of the resource to apply | 

### Return type

[**Workload**](Workload.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

