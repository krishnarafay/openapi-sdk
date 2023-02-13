# \NamespaceNetworkPolicyDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesNameGet**](NamespaceNetworkPolicyDetailApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesNameGet) | **Get** /apis/security.k8smgmt.io/v3/projects/{project}/namespacenetworkpolicies/{name} | 
[**ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysNameGet**](NamespaceNetworkPolicyDetailApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysNameGet) | **Get** /apis/security.k8smgmt.io/v3/projects/{project}/namespacenetworkpolicys/{name} | 



## ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesNameGet

> NamespaceNetworkPolicy ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.NamespaceNetworkPolicyDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceNetworkPolicyDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesNameGet`: NamespaceNetworkPolicy
    fmt.Fprintf(os.Stdout, "Response from `NamespaceNetworkPolicyDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NamespaceNetworkPolicy**](NamespaceNetworkPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysNameGet

> NamespaceNetworkPolicy ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.NamespaceNetworkPolicyDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceNetworkPolicyDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysNameGet`: NamespaceNetworkPolicy
    fmt.Fprintf(os.Stdout, "Response from `NamespaceNetworkPolicyDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NamespaceNetworkPolicy**](NamespaceNetworkPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

