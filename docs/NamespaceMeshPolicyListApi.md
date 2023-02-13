# \NamespaceMeshPolicyListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesGet**](NamespaceMeshPolicyListApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesGet) | **Get** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/namespacemeshpolicies | 
[**ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysGet**](NamespaceMeshPolicyListApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysGet) | **Get** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/namespacemeshpolicys | 



## ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesGet

> NamespaceMeshPolicyList ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesGet(ctx, project).Execute()





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
    resp, r, err := apiClient.NamespaceMeshPolicyListApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMeshPolicyListApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesGet`: NamespaceMeshPolicyList
    fmt.Fprintf(os.Stdout, "Response from `NamespaceMeshPolicyListApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NamespaceMeshPolicyList**](NamespaceMeshPolicyList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysGet

> NamespaceMeshPolicyList ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysGet(ctx, project).Execute()





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
    resp, r, err := apiClient.NamespaceMeshPolicyListApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMeshPolicyListApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysGet`: NamespaceMeshPolicyList
    fmt.Fprintf(os.Stdout, "Response from `NamespaceMeshPolicyListApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NamespaceMeshPolicyList**](NamespaceMeshPolicyList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

