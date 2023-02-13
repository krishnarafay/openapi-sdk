# \NamespaceNetworkPolicyListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesGet**](NamespaceNetworkPolicyListApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesGet) | **Get** /apis/security.k8smgmt.io/v3/projects/{project}/namespacenetworkpolicies | 
[**ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysGet**](NamespaceNetworkPolicyListApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysGet) | **Get** /apis/security.k8smgmt.io/v3/projects/{project}/namespacenetworkpolicys | 



## ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesGet

> NamespaceNetworkPolicyList ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesGet(ctx, project).Execute()





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
    resp, r, err := apiClient.NamespaceNetworkPolicyListApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceNetworkPolicyListApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesGet`: NamespaceNetworkPolicyList
    fmt.Fprintf(os.Stdout, "Response from `NamespaceNetworkPolicyListApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NamespaceNetworkPolicyList**](NamespaceNetworkPolicyList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysGet

> NamespaceNetworkPolicyList ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysGet(ctx, project).Execute()





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
    resp, r, err := apiClient.NamespaceNetworkPolicyListApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceNetworkPolicyListApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysGet`: NamespaceNetworkPolicyList
    fmt.Fprintf(os.Stdout, "Response from `NamespaceNetworkPolicyListApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NamespaceNetworkPolicyList**](NamespaceNetworkPolicyList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

