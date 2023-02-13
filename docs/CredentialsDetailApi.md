# \CredentialsDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameArtifactsGet**](CredentialsDetailApi.md#ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameArtifactsGet) | **Get** /apis/infra.k8smgmt.io/v3/projects/{project}/credentials/{name}/artifacts | 
[**ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameGet**](CredentialsDetailApi.md#ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameGet) | **Get** /apis/infra.k8smgmt.io/v3/projects/{project}/credentials/{name} | 
[**ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameArtifactsGet**](CredentialsDetailApi.md#ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameArtifactsGet) | **Get** /apis/infra.k8smgmt.io/v3/projects/{project}/credentialss/{name}/artifacts | 
[**ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameGet**](CredentialsDetailApi.md#ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameGet) | **Get** /apis/infra.k8smgmt.io/v3/projects/{project}/credentialss/{name} | 



## ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameArtifactsGet

> Credentials ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.CredentialsDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameArtifactsGet`: Credentials
    fmt.Fprintf(os.Stdout, "Response from `CredentialsDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Credentials**](Credentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameGet

> Credentials ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.CredentialsDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameGet`: Credentials
    fmt.Fprintf(os.Stdout, "Response from `CredentialsDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectCredentialsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Credentials**](Credentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameArtifactsGet

> Credentials ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.CredentialsDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameArtifactsGet`: Credentials
    fmt.Fprintf(os.Stdout, "Response from `CredentialsDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Credentials**](Credentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameGet

> Credentials ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.CredentialsDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameGet`: Credentials
    fmt.Fprintf(os.Stdout, "Response from `CredentialsDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectCredentialssNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Credentials**](Credentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

