# \WorkloadTemplateApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesPost**](WorkloadTemplateApplyApi.md#ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesPost) | **Post** /apis/apps.k8smgmt.io/v3/projects/{project}/workloadtemplates | 



## ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesPost

> WorkloadTemplate ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesPost(ctx, project).WorkloadTemplate(workloadTemplate).Execute()





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
    workloadTemplate := *openapiclient.NewWorkloadTemplate("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewWorkloadTemplateSpec()) // WorkloadTemplate | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkloadTemplateApplyApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesPost(context.Background(), project).WorkloadTemplate(workloadTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadTemplateApplyApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesPost`: WorkloadTemplate
    fmt.Fprintf(os.Stdout, "Response from `WorkloadTemplateApplyApi.ApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsK8smgmtIoV3ProjectsProjectWorkloadtemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workloadTemplate** | [**WorkloadTemplate**](WorkloadTemplate.md) | schema of the resource to apply | 

### Return type

[**WorkloadTemplate**](WorkloadTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

