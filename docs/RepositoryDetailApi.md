# \RepositoryDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameArtifactsGet**](RepositoryDetailApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameArtifactsGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/repositories/{name}/artifacts | 
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameGet**](RepositoryDetailApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/repositories/{name} | 
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameArtifactsGet**](RepositoryDetailApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameArtifactsGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/repositorys/{name}/artifacts | 
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameGet**](RepositoryDetailApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/repositorys/{name} | 



## ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameArtifactsGet

> Repository ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameArtifactsGet`: Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Repository**](Repository.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameGet

> Repository ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameGet`: Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectRepositoriesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Repository**](Repository.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameArtifactsGet

> Repository ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameArtifactsGet`: Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Repository**](Repository.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameGet

> Repository ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameGet`: Repository
    fmt.Fprintf(os.Stdout, "Response from `RepositoryDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectRepositorysNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Repository**](Repository.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

