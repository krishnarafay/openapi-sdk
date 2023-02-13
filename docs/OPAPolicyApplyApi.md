# \OPAPolicyApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesPost**](OPAPolicyApplyApi.md#ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesPost) | **Post** /apis/opa.k8smgmt.io/v3/projects/{project}/opapolicies | 
[**ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysPost**](OPAPolicyApplyApi.md#ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysPost) | **Post** /apis/opa.k8smgmt.io/v3/projects/{project}/opapolicys | 



## ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesPost

> OPAPolicy ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesPost(ctx, project).OPAPolicy(oPAPolicy).Execute()





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
    oPAPolicy := *openapiclient.NewOPAPolicy("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewOPAPolicySpec()) // OPAPolicy | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OPAPolicyApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesPost(context.Background(), project).OPAPolicy(oPAPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OPAPolicyApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesPost`: OPAPolicy
    fmt.Fprintf(os.Stdout, "Response from `OPAPolicyApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oPAPolicy** | [**OPAPolicy**](OPAPolicy.md) | schema of the resource to apply | 

### Return type

[**OPAPolicy**](OPAPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysPost

> OPAPolicy ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysPost(ctx, project).OPAPolicy(oPAPolicy).Execute()





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
    oPAPolicy := *openapiclient.NewOPAPolicy("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewOPAPolicySpec()) // OPAPolicy | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OPAPolicyApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysPost(context.Background(), project).OPAPolicy(oPAPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OPAPolicyApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysPost`: OPAPolicy
    fmt.Fprintf(os.Stdout, "Response from `OPAPolicyApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisOpaK8smgmtIoV3ProjectsProjectOpapolicysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oPAPolicy** | [**OPAPolicy**](OPAPolicy.md) | schema of the resource to apply | 

### Return type

[**OPAPolicy**](OPAPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

