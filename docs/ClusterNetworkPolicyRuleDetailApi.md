# \ClusterNetworkPolicyRuleDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameArtifactsGet**](ClusterNetworkPolicyRuleDetailApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameArtifactsGet) | **Get** /apis/security.k8smgmt.io/v3/projects/{project}/clusternetworkpolicyrules/{name}/artifacts | 
[**ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameGet**](ClusterNetworkPolicyRuleDetailApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameGet) | **Get** /apis/security.k8smgmt.io/v3/projects/{project}/clusternetworkpolicyrules/{name} | 



## ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameArtifactsGet

> ClusterNetworkPolicyRule ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.ClusterNetworkPolicyRuleDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNetworkPolicyRuleDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameArtifactsGet`: ClusterNetworkPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `ClusterNetworkPolicyRuleDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterNetworkPolicyRule**](ClusterNetworkPolicyRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameGet

> ClusterNetworkPolicyRule ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.ClusterNetworkPolicyRuleDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNetworkPolicyRuleDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameGet`: ClusterNetworkPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `ClusterNetworkPolicyRuleDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicyrulesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterNetworkPolicyRule**](ClusterNetworkPolicyRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

