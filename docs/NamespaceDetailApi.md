# \NamespaceDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameArtifactsGet**](NamespaceDetailApi.md#ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameArtifactsGet) | **Get** /apis/infra.k8smgmt.io/v3/projects/{project}/namespaces/{name}/artifacts | 
[**ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameGet**](NamespaceDetailApi.md#ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameGet) | **Get** /apis/infra.k8smgmt.io/v3/projects/{project}/namespaces/{name} | 
[**ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameStatusGet**](NamespaceDetailApi.md#ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameStatusGet) | **Get** /apis/infra.k8smgmt.io/v3/projects/{project}/namespaces/{name}/status | 



## ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameArtifactsGet

> Namespace ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.NamespaceDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameArtifactsGet`: Namespace
    fmt.Fprintf(os.Stdout, "Response from `NamespaceDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Namespace**](Namespace.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameGet

> Namespace ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.NamespaceDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameGet`: Namespace
    fmt.Fprintf(os.Stdout, "Response from `NamespaceDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Namespace**](Namespace.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameStatusGet

> Namespace ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameStatusGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.NamespaceDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameStatusGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameStatusGet`: Namespace
    fmt.Fprintf(os.Stdout, "Response from `NamespaceDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectNamespacesNameStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Namespace**](Namespace.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

