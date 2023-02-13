# \NamespaceNetworkPolicyRuleDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameArtifactsGet**](NamespaceNetworkPolicyRuleDetailApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameArtifactsGet) | **Get** /apis/security.k8smgmt.io/v3/projects/{project}/namespacenetworkpolicyrules/{name}/artifacts | 
[**ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameGet**](NamespaceNetworkPolicyRuleDetailApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameGet) | **Get** /apis/security.k8smgmt.io/v3/projects/{project}/namespacenetworkpolicyrules/{name} | 



## ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameArtifactsGet

> NamespaceNetworkPolicyRule ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.NamespaceNetworkPolicyRuleDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceNetworkPolicyRuleDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameArtifactsGet`: NamespaceNetworkPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `NamespaceNetworkPolicyRuleDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NamespaceNetworkPolicyRule**](NamespaceNetworkPolicyRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameGet

> NamespaceNetworkPolicyRule ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.NamespaceNetworkPolicyRuleDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceNetworkPolicyRuleDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameGet`: NamespaceNetworkPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `NamespaceNetworkPolicyRuleDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NamespaceNetworkPolicyRule**](NamespaceNetworkPolicyRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

