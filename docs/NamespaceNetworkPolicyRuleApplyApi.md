# \NamespaceNetworkPolicyRuleApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesPost**](NamespaceNetworkPolicyRuleApplyApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesPost) | **Post** /apis/security.k8smgmt.io/v3/projects/{project}/namespacenetworkpolicyrules | 



## ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesPost

> NamespaceNetworkPolicyRule ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesPost(ctx, project).NamespaceNetworkPolicyRule(namespaceNetworkPolicyRule).Execute()





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
    namespaceNetworkPolicyRule := *openapiclient.NewNamespaceNetworkPolicyRule("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewNamespaceNetworkPolicyRuleSpec()) // NamespaceNetworkPolicyRule | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceNetworkPolicyRuleApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesPost(context.Background(), project).NamespaceNetworkPolicyRule(namespaceNetworkPolicyRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceNetworkPolicyRuleApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesPost`: NamespaceNetworkPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `NamespaceNetworkPolicyRuleApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespaceNetworkPolicyRule** | [**NamespaceNetworkPolicyRule**](NamespaceNetworkPolicyRule.md) | schema of the resource to apply | 

### Return type

[**NamespaceNetworkPolicyRule**](NamespaceNetworkPolicyRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

