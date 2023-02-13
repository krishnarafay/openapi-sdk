# \SecretProviderClassApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesPost**](SecretProviderClassApplyApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesPost) | **Post** /apis/integrations.k8smgmt.io/v3/projects/{project}/secretproviderclasses | 
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssPost**](SecretProviderClassApplyApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssPost) | **Post** /apis/integrations.k8smgmt.io/v3/projects/{project}/secretproviderclasss | 



## ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesPost

> SecretProviderClass ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesPost(ctx, project).SecretProviderClass(secretProviderClass).Execute()





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
    secretProviderClass := *openapiclient.NewSecretProviderClass("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewSecretProviderClassSpec()) // SecretProviderClass | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretProviderClassApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesPost(context.Background(), project).SecretProviderClass(secretProviderClass).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretProviderClassApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesPost`: SecretProviderClass
    fmt.Fprintf(os.Stdout, "Response from `SecretProviderClassApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secretProviderClass** | [**SecretProviderClass**](SecretProviderClass.md) | schema of the resource to apply | 

### Return type

[**SecretProviderClass**](SecretProviderClass.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssPost

> SecretProviderClass ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssPost(ctx, project).SecretProviderClass(secretProviderClass).Execute()





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
    secretProviderClass := *openapiclient.NewSecretProviderClass("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewSecretProviderClassSpec()) // SecretProviderClass | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretProviderClassApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssPost(context.Background(), project).SecretProviderClass(secretProviderClass).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretProviderClassApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssPost`: SecretProviderClass
    fmt.Fprintf(os.Stdout, "Response from `SecretProviderClassApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secretProviderClass** | [**SecretProviderClass**](SecretProviderClass.md) | schema of the resource to apply | 

### Return type

[**SecretProviderClass**](SecretProviderClass.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

