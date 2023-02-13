# \NamespaceMeshPolicyDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesNameGet**](NamespaceMeshPolicyDetailApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesNameGet) | **Get** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/namespacemeshpolicies/{name} | 
[**ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysNameGet**](NamespaceMeshPolicyDetailApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysNameGet) | **Get** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/namespacemeshpolicys/{name} | 



## ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesNameGet

> NamespaceMeshPolicy ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.NamespaceMeshPolicyDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMeshPolicyDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesNameGet`: NamespaceMeshPolicy
    fmt.Fprintf(os.Stdout, "Response from `NamespaceMeshPolicyDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NamespaceMeshPolicy**](NamespaceMeshPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysNameGet

> NamespaceMeshPolicy ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.NamespaceMeshPolicyDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMeshPolicyDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysNameGet`: NamespaceMeshPolicy
    fmt.Fprintf(os.Stdout, "Response from `NamespaceMeshPolicyDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NamespaceMeshPolicy**](NamespaceMeshPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

