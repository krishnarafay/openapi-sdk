# \ClusterMeshPolicyDeleteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameDelete**](ClusterMeshPolicyDeleteApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameDelete) | **Delete** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/clustermeshpolicies/{name} | 
[**ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameDelete**](ClusterMeshPolicyDeleteApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameDelete) | **Delete** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/clustermeshpolicys/{name} | 



## ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameDelete

> ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameDelete(ctx, project, name).Execute()





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
    r, err := apiClient.ClusterMeshPolicyDeleteApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameDelete(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterMeshPolicyDeleteApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpoliciesNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameDelete

> ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameDelete(ctx, project, name).Execute()





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
    r, err := apiClient.ClusterMeshPolicyDeleteApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameDelete(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterMeshPolicyDeleteApi.ApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectClustermeshpolicysNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

