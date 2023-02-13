# \ClusterMeshPolicyDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameGet**](ClusterMeshPolicyDetailApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameGet) | **Get** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/clustermeshpolicies/{name} | 
[**ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameGet**](ClusterMeshPolicyDetailApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameGet) | **Get** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/clustermeshpolicys/{name} | 



## ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameGet

> ClusterMeshPolicy ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.ClusterMeshPolicyDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterMeshPolicyDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameGet`: ClusterMeshPolicy
    fmt.Fprintf(os.Stdout, "Response from `ClusterMeshPolicyDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterMeshPolicy**](ClusterMeshPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameGet

> ClusterMeshPolicy ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.ClusterMeshPolicyDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterMeshPolicyDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameGet`: ClusterMeshPolicy
    fmt.Fprintf(os.Stdout, "Response from `ClusterMeshPolicyDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterMeshPolicy**](ClusterMeshPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

