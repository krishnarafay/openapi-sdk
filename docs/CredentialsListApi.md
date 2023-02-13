# \CredentialsListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisInfraK8smgmtIoV3ProjectsProjectCredentialsGet**](CredentialsListApi.md#ApisInfraK8smgmtIoV3ProjectsProjectCredentialsGet) | **Get** /apis/infra.k8smgmt.io/v3/projects/{project}/credentials | 
[**ApisInfraK8smgmtIoV3ProjectsProjectCredentialssGet**](CredentialsListApi.md#ApisInfraK8smgmtIoV3ProjectsProjectCredentialssGet) | **Get** /apis/infra.k8smgmt.io/v3/projects/{project}/credentialss | 



## ApisInfraK8smgmtIoV3ProjectsProjectCredentialsGet

> CredentialsList ApisInfraK8smgmtIoV3ProjectsProjectCredentialsGet(ctx, project).Execute()





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
    resp, r, err := apiClient.CredentialsListApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialsGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsListApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectCredentialsGet`: CredentialsList
    fmt.Fprintf(os.Stdout, "Response from `CredentialsListApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectCredentialsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CredentialsList**](CredentialsList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisInfraK8smgmtIoV3ProjectsProjectCredentialssGet

> CredentialsList ApisInfraK8smgmtIoV3ProjectsProjectCredentialssGet(ctx, project).Execute()





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
    resp, r, err := apiClient.CredentialsListApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialssGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsListApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialssGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectCredentialssGet`: CredentialsList
    fmt.Fprintf(os.Stdout, "Response from `CredentialsListApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialssGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectCredentialssGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CredentialsList**](CredentialsList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

