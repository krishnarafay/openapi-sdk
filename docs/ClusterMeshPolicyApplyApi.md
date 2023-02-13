# \ClusterMeshPolicyApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesPost**](ClusterMeshPolicyApplyApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesPost) | **Post** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/clustermeshpolicies | 
[**ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysPost**](ClusterMeshPolicyApplyApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysPost) | **Post** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/clustermeshpolicys | 



## ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesPost

> ClusterMeshPolicy ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesPost(ctx, project).ClusterMeshPolicy(clusterMeshPolicy).Execute()





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
    clusterMeshPolicy := *openapiclient.NewClusterMeshPolicy("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewClusterMeshPolicySpec()) // ClusterMeshPolicy | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterMeshPolicyApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesPost(context.Background(), project).ClusterMeshPolicy(clusterMeshPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterMeshPolicyApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesPost`: ClusterMeshPolicy
    fmt.Fprintf(os.Stdout, "Response from `ClusterMeshPolicyApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterMeshPolicy** | [**ClusterMeshPolicy**](ClusterMeshPolicy.md) | schema of the resource to apply | 

### Return type

[**ClusterMeshPolicy**](ClusterMeshPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysPost

> ClusterMeshPolicy ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysPost(ctx, project).ClusterMeshPolicy(clusterMeshPolicy).Execute()





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
    clusterMeshPolicy := *openapiclient.NewClusterMeshPolicy("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewClusterMeshPolicySpec()) // ClusterMeshPolicy | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterMeshPolicyApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysPost(context.Background(), project).ClusterMeshPolicy(clusterMeshPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterMeshPolicyApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysPost`: ClusterMeshPolicy
    fmt.Fprintf(os.Stdout, "Response from `ClusterMeshPolicyApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterMeshPolicy** | [**ClusterMeshPolicy**](ClusterMeshPolicy.md) | schema of the resource to apply | 

### Return type

[**ClusterMeshPolicy**](ClusterMeshPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

