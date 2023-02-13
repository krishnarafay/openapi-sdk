# \ClusterNetworkPolicyApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesPost**](ClusterNetworkPolicyApplyApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesPost) | **Post** /apis/security.k8smgmt.io/v3/projects/{project}/clusternetworkpolicies | 
[**ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysPost**](ClusterNetworkPolicyApplyApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysPost) | **Post** /apis/security.k8smgmt.io/v3/projects/{project}/clusternetworkpolicys | 



## ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesPost

> ClusterNetworkPolicy ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesPost(ctx, project).ClusterNetworkPolicy(clusterNetworkPolicy).Execute()





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
    clusterNetworkPolicy := *openapiclient.NewClusterNetworkPolicy("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewClusterNetworkPolicySpec()) // ClusterNetworkPolicy | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterNetworkPolicyApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesPost(context.Background(), project).ClusterNetworkPolicy(clusterNetworkPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNetworkPolicyApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesPost`: ClusterNetworkPolicy
    fmt.Fprintf(os.Stdout, "Response from `ClusterNetworkPolicyApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNetworkPolicy** | [**ClusterNetworkPolicy**](ClusterNetworkPolicy.md) | schema of the resource to apply | 

### Return type

[**ClusterNetworkPolicy**](ClusterNetworkPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysPost

> ClusterNetworkPolicy ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysPost(ctx, project).ClusterNetworkPolicy(clusterNetworkPolicy).Execute()





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
    clusterNetworkPolicy := *openapiclient.NewClusterNetworkPolicy("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewClusterNetworkPolicySpec()) // ClusterNetworkPolicy | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterNetworkPolicyApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysPost(context.Background(), project).ClusterNetworkPolicy(clusterNetworkPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNetworkPolicyApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysPost`: ClusterNetworkPolicy
    fmt.Fprintf(os.Stdout, "Response from `ClusterNetworkPolicyApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNetworkPolicy** | [**ClusterNetworkPolicy**](ClusterNetworkPolicy.md) | schema of the resource to apply | 

### Return type

[**ClusterNetworkPolicy**](ClusterNetworkPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

