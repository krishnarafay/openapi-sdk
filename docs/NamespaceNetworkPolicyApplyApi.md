# \NamespaceNetworkPolicyApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPost**](NamespaceNetworkPolicyApplyApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPost) | **Post** /apis/security.k8smgmt.io/v3/projects/{project}/namespacenetworkpolicies | 
[**ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPost**](NamespaceNetworkPolicyApplyApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPost) | **Post** /apis/security.k8smgmt.io/v3/projects/{project}/namespacenetworkpolicys | 



## ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPost

> NamespaceNetworkPolicy ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPost(ctx, project).NamespaceNetworkPolicy(namespaceNetworkPolicy).Execute()





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
    namespaceNetworkPolicy := *openapiclient.NewNamespaceNetworkPolicy("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewNamespaceNetworkPolicySpec()) // NamespaceNetworkPolicy | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceNetworkPolicyApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPost(context.Background(), project).NamespaceNetworkPolicy(namespaceNetworkPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceNetworkPolicyApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPost`: NamespaceNetworkPolicy
    fmt.Fprintf(os.Stdout, "Response from `NamespaceNetworkPolicyApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpoliciesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespaceNetworkPolicy** | [**NamespaceNetworkPolicy**](NamespaceNetworkPolicy.md) | schema of the resource to apply | 

### Return type

[**NamespaceNetworkPolicy**](NamespaceNetworkPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPost

> NamespaceNetworkPolicy ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPost(ctx, project).NamespaceNetworkPolicy(namespaceNetworkPolicy).Execute()





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
    namespaceNetworkPolicy := *openapiclient.NewNamespaceNetworkPolicy("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewNamespaceNetworkPolicySpec()) // NamespaceNetworkPolicy | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceNetworkPolicyApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPost(context.Background(), project).NamespaceNetworkPolicy(namespaceNetworkPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceNetworkPolicyApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPost`: NamespaceNetworkPolicy
    fmt.Fprintf(os.Stdout, "Response from `NamespaceNetworkPolicyApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespaceNetworkPolicy** | [**NamespaceNetworkPolicy**](NamespaceNetworkPolicy.md) | schema of the resource to apply | 

### Return type

[**NamespaceNetworkPolicy**](NamespaceNetworkPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

