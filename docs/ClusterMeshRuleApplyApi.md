# \ClusterMeshRuleApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesPost**](ClusterMeshRuleApplyApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesPost) | **Post** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/clustermeshrules | 



## ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesPost

> ClusterMeshRule ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesPost(ctx, project).ClusterMeshRule(clusterMeshRule).Execute()





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
    clusterMeshRule := *openapiclient.NewClusterMeshRule("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewClusterMeshRuleSpec()) // ClusterMeshRule | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterMeshRuleApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesPost(context.Background(), project).ClusterMeshRule(clusterMeshRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterMeshRuleApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesPost`: ClusterMeshRule
    fmt.Fprintf(os.Stdout, "Response from `ClusterMeshRuleApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterMeshRule** | [**ClusterMeshRule**](ClusterMeshRule.md) | schema of the resource to apply | 

### Return type

[**ClusterMeshRule**](ClusterMeshRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

