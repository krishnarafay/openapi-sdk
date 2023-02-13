# \ClusterDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisInfraK8smgmtIoV3ProjectsProjectClustersNameGet**](ClusterDetailApi.md#ApisInfraK8smgmtIoV3ProjectsProjectClustersNameGet) | **Get** /apis/infra.k8smgmt.io/v3/projects/{project}/clusters/{name} | 
[**ApisInfraK8smgmtIoV3ProjectsProjectClustersNameStatusGet**](ClusterDetailApi.md#ApisInfraK8smgmtIoV3ProjectsProjectClustersNameStatusGet) | **Get** /apis/infra.k8smgmt.io/v3/projects/{project}/clusters/{name}/status | 



## ApisInfraK8smgmtIoV3ProjectsProjectClustersNameGet

> Cluster ApisInfraK8smgmtIoV3ProjectsProjectClustersNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.ClusterDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectClustersNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectClustersNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectClustersNameGet`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClusterDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectClustersNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectClustersNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Cluster**](Cluster.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisInfraK8smgmtIoV3ProjectsProjectClustersNameStatusGet

> Cluster ApisInfraK8smgmtIoV3ProjectsProjectClustersNameStatusGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.ClusterDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectClustersNameStatusGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectClustersNameStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectClustersNameStatusGet`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClusterDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectClustersNameStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectClustersNameStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Cluster**](Cluster.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

