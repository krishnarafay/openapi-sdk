# \CredentialsApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisInfraK8smgmtIoV3ProjectsProjectCredentialsPost**](CredentialsApplyApi.md#ApisInfraK8smgmtIoV3ProjectsProjectCredentialsPost) | **Post** /apis/infra.k8smgmt.io/v3/projects/{project}/credentials | 
[**ApisInfraK8smgmtIoV3ProjectsProjectCredentialssPost**](CredentialsApplyApi.md#ApisInfraK8smgmtIoV3ProjectsProjectCredentialssPost) | **Post** /apis/infra.k8smgmt.io/v3/projects/{project}/credentialss | 



## ApisInfraK8smgmtIoV3ProjectsProjectCredentialsPost

> Credentials ApisInfraK8smgmtIoV3ProjectsProjectCredentialsPost(ctx, project).Credentials(credentials).Execute()





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
    credentials := *openapiclient.NewCredentials("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewCredentialsSpec()) // Credentials | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApplyApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialsPost(context.Background(), project).Credentials(credentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApplyApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectCredentialsPost`: Credentials
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApplyApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectCredentialsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentials** | [**Credentials**](Credentials.md) | schema of the resource to apply | 

### Return type

[**Credentials**](Credentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisInfraK8smgmtIoV3ProjectsProjectCredentialssPost

> Credentials ApisInfraK8smgmtIoV3ProjectsProjectCredentialssPost(ctx, project).Credentials(credentials).Execute()





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
    credentials := *openapiclient.NewCredentials("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewCredentialsSpec()) // Credentials | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialsApplyApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialssPost(context.Background(), project).Credentials(credentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialsApplyApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialssPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectCredentialssPost`: Credentials
    fmt.Fprintf(os.Stdout, "Response from `CredentialsApplyApi.ApisInfraK8smgmtIoV3ProjectsProjectCredentialssPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectCredentialssPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentials** | [**Credentials**](Credentials.md) | schema of the resource to apply | 

### Return type

[**Credentials**](Credentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

