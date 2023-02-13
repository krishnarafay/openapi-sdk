# \SecretProviderClassDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameArtifactsGet**](SecretProviderClassDetailApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameArtifactsGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/secretproviderclasses/{name}/artifacts | 
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameGet**](SecretProviderClassDetailApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/secretproviderclasses/{name} | 
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameArtifactsGet**](SecretProviderClassDetailApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameArtifactsGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/secretproviderclasss/{name}/artifacts | 
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameGet**](SecretProviderClassDetailApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/secretproviderclasss/{name} | 



## ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameArtifactsGet

> SecretProviderClass ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.SecretProviderClassDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretProviderClassDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameArtifactsGet`: SecretProviderClass
    fmt.Fprintf(os.Stdout, "Response from `SecretProviderClassDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SecretProviderClass**](SecretProviderClass.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameGet

> SecretProviderClass ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.SecretProviderClassDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretProviderClassDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameGet`: SecretProviderClass
    fmt.Fprintf(os.Stdout, "Response from `SecretProviderClassDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SecretProviderClass**](SecretProviderClass.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameArtifactsGet

> SecretProviderClass ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.SecretProviderClassDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretProviderClassDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameArtifactsGet`: SecretProviderClass
    fmt.Fprintf(os.Stdout, "Response from `SecretProviderClassDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SecretProviderClass**](SecretProviderClass.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameGet

> SecretProviderClass ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.SecretProviderClassDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretProviderClassDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameGet`: SecretProviderClass
    fmt.Fprintf(os.Stdout, "Response from `SecretProviderClassDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SecretProviderClass**](SecretProviderClass.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

