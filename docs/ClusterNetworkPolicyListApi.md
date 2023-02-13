# \ClusterNetworkPolicyListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesGet**](ClusterNetworkPolicyListApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesGet) | **Get** /apis/security.k8smgmt.io/v3/projects/{project}/clusternetworkpolicies | 
[**ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysGet**](ClusterNetworkPolicyListApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysGet) | **Get** /apis/security.k8smgmt.io/v3/projects/{project}/clusternetworkpolicys | 



## ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesGet

> ClusterNetworkPolicyList ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesGet(ctx, project).Execute()





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
    resp, r, err := apiClient.ClusterNetworkPolicyListApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNetworkPolicyListApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesGet`: ClusterNetworkPolicyList
    fmt.Fprintf(os.Stdout, "Response from `ClusterNetworkPolicyListApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterNetworkPolicyList**](ClusterNetworkPolicyList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysGet

> ClusterNetworkPolicyList ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysGet(ctx, project).Execute()





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
    resp, r, err := apiClient.ClusterNetworkPolicyListApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNetworkPolicyListApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysGet`: ClusterNetworkPolicyList
    fmt.Fprintf(os.Stdout, "Response from `ClusterNetworkPolicyListApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterNetworkPolicyList**](ClusterNetworkPolicyList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

