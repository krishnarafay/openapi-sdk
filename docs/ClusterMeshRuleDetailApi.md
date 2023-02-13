# \ClusterMeshRuleDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameArtifactsGet**](ClusterMeshRuleDetailApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameArtifactsGet) | **Get** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/clustermeshrules/{name}/artifacts | 
[**ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameGet**](ClusterMeshRuleDetailApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameGet) | **Get** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/clustermeshrules/{name} | 



## ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameArtifactsGet

> ClusterMeshRule ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.ClusterMeshRuleDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterMeshRuleDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameArtifactsGet`: ClusterMeshRule
    fmt.Fprintf(os.Stdout, "Response from `ClusterMeshRuleDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterMeshRule**](ClusterMeshRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameGet

> ClusterMeshRule ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.ClusterMeshRuleDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterMeshRuleDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameGet`: ClusterMeshRule
    fmt.Fprintf(os.Stdout, "Response from `ClusterMeshRuleDetailApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshrulesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterMeshRule**](ClusterMeshRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

