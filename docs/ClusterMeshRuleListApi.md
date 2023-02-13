# \ClusterMeshRuleListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesGet**](ClusterMeshRuleListApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesGet) | **Get** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/clustermeshrules | 



## ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesGet

> ClusterMeshRuleList ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesGet(ctx, project).Execute()





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
    resp, r, err := apiClient.ClusterMeshRuleListApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterMeshRuleListApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesGet`: ClusterMeshRuleList
    fmt.Fprintf(os.Stdout, "Response from `ClusterMeshRuleListApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterMeshRuleList**](ClusterMeshRuleList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

