# \OPAPolicyListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesGet**](OPAPolicyListApi.md#ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesGet) | **Get** /apis/opa.k8smgmt.io/v3/projects/{project}/opapolicies | 
[**ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysGet**](OPAPolicyListApi.md#ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysGet) | **Get** /apis/opa.k8smgmt.io/v3/projects/{project}/opapolicys | 



## ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesGet

> OPAPolicyList ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesGet(ctx, project).Execute()





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
    resp, r, err := apiClient.OPAPolicyListApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OPAPolicyListApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesGet`: OPAPolicyList
    fmt.Fprintf(os.Stdout, "Response from `OPAPolicyListApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OPAPolicyList**](OPAPolicyList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysGet

> OPAPolicyList ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysGet(ctx, project).Execute()





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
    resp, r, err := apiClient.OPAPolicyListApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OPAPolicyListApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysGet`: OPAPolicyList
    fmt.Fprintf(os.Stdout, "Response from `OPAPolicyListApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisOpaK8smgmtIoV3ProjectsProjectOpapolicysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OPAPolicyList**](OPAPolicyList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

