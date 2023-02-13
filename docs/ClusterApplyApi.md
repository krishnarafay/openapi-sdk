# \ClusterApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisInfraK8smgmtIoV3ProjectsProjectClustersPost**](ClusterApplyApi.md#ApisInfraK8smgmtIoV3ProjectsProjectClustersPost) | **Post** /apis/infra.k8smgmt.io/v3/projects/{project}/clusters | 



## ApisInfraK8smgmtIoV3ProjectsProjectClustersPost

> Cluster ApisInfraK8smgmtIoV3ProjectsProjectClustersPost(ctx, project).Cluster(cluster).Execute()





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
    cluster := *openapiclient.NewCluster() // Cluster | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterApplyApi.ApisInfraK8smgmtIoV3ProjectsProjectClustersPost(context.Background(), project).Cluster(cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApplyApi.ApisInfraK8smgmtIoV3ProjectsProjectClustersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectClustersPost`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClusterApplyApi.ApisInfraK8smgmtIoV3ProjectsProjectClustersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectClustersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cluster** | [**Cluster**](Cluster.md) | schema of the resource to apply | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

