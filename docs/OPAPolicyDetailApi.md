# \OPAPolicyDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesNameGet**](OPAPolicyDetailApi.md#ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesNameGet) | **Get** /apis/opa.k8smgmt.io/v3/projects/{project}/opapolicies/{name} | 
[**ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysNameGet**](OPAPolicyDetailApi.md#ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysNameGet) | **Get** /apis/opa.k8smgmt.io/v3/projects/{project}/opapolicys/{name} | 



## ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesNameGet

> OPAPolicy ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.OPAPolicyDetailApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OPAPolicyDetailApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesNameGet`: OPAPolicy
    fmt.Fprintf(os.Stdout, "Response from `OPAPolicyDetailApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisOpaK8smgmtIoV3ProjectsProjectOpapoliciesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OPAPolicy**](OPAPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysNameGet

> OPAPolicy ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.OPAPolicyDetailApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OPAPolicyDetailApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysNameGet`: OPAPolicy
    fmt.Fprintf(os.Stdout, "Response from `OPAPolicyDetailApi.ApisOpaK8smgmtIoV3ProjectsProjectOpapolicysNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisOpaK8smgmtIoV3ProjectsProjectOpapolicysNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OPAPolicy**](OPAPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

