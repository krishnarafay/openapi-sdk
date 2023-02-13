# \ClusterNetworkPolicyRuleApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesPost**](ClusterNetworkPolicyRuleApplyApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesPost) | **Post** /apis/security.k8smgmt.io/v3/projects/{project}/clusternetworkpolicyrules | 



## ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesPost

> ClusterNetworkPolicyRule ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesPost(ctx, project).ClusterNetworkPolicyRule(clusterNetworkPolicyRule).Execute()





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
    clusterNetworkPolicyRule := *openapiclient.NewClusterNetworkPolicyRule("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewClusterNetworkPolicyRuleSpec()) // ClusterNetworkPolicyRule | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterNetworkPolicyRuleApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesPost(context.Background(), project).ClusterNetworkPolicyRule(clusterNetworkPolicyRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNetworkPolicyRuleApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesPost`: ClusterNetworkPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `ClusterNetworkPolicyRuleApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNetworkPolicyRule** | [**ClusterNetworkPolicyRule**](ClusterNetworkPolicyRule.md) | schema of the resource to apply | 

### Return type

[**ClusterNetworkPolicyRule**](ClusterNetworkPolicyRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

