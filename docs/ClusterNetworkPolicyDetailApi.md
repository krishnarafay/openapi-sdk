# \ClusterNetworkPolicyDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesNameGet**](ClusterNetworkPolicyDetailApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesNameGet) | **Get** /apis/security.k8smgmt.io/v3/projects/{project}/clusternetworkpolicies/{name} | 
[**ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysNameGet**](ClusterNetworkPolicyDetailApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysNameGet) | **Get** /apis/security.k8smgmt.io/v3/projects/{project}/clusternetworkpolicys/{name} | 



## ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesNameGet

> ClusterNetworkPolicy ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.ClusterNetworkPolicyDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNetworkPolicyDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesNameGet`: ClusterNetworkPolicy
    fmt.Fprintf(os.Stdout, "Response from `ClusterNetworkPolicyDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpoliciesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterNetworkPolicy**](ClusterNetworkPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysNameGet

> ClusterNetworkPolicy ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.ClusterNetworkPolicyDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterNetworkPolicyDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysNameGet`: ClusterNetworkPolicy
    fmt.Fprintf(os.Stdout, "Response from `ClusterNetworkPolicyDetailApi.ApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectClusternetworkpolicysNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterNetworkPolicy**](ClusterNetworkPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

