# \OverrideApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisAppsK8smgmtIoV3ProjectsProjectOverridesPost**](OverrideApplyApi.md#ApisAppsK8smgmtIoV3ProjectsProjectOverridesPost) | **Post** /apis/apps.k8smgmt.io/v3/projects/{project}/overrides | 



## ApisAppsK8smgmtIoV3ProjectsProjectOverridesPost

> Override ApisAppsK8smgmtIoV3ProjectsProjectOverridesPost(ctx, project).Override(override).Execute()





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
    override := *openapiclient.NewOverride("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewOverrideSpec()) // Override | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverrideApplyApi.ApisAppsK8smgmtIoV3ProjectsProjectOverridesPost(context.Background(), project).Override(override).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverrideApplyApi.ApisAppsK8smgmtIoV3ProjectsProjectOverridesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsK8smgmtIoV3ProjectsProjectOverridesPost`: Override
    fmt.Fprintf(os.Stdout, "Response from `OverrideApplyApi.ApisAppsK8smgmtIoV3ProjectsProjectOverridesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsK8smgmtIoV3ProjectsProjectOverridesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **override** | [**Override**](Override.md) | schema of the resource to apply | 

### Return type

[**Override**](Override.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

