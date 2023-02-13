# \SecretProviderClassListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesGet**](SecretProviderClassListApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/secretproviderclasses | 
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssGet**](SecretProviderClassListApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/secretproviderclasss | 



## ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesGet

> SecretProviderClassList ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesGet(ctx, project).Execute()





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
    resp, r, err := apiClient.SecretProviderClassListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretProviderClassListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesGet`: SecretProviderClassList
    fmt.Fprintf(os.Stdout, "Response from `SecretProviderClassListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclassesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecretProviderClassList**](SecretProviderClassList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssGet

> SecretProviderClassList ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssGet(ctx, project).Execute()





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
    resp, r, err := apiClient.SecretProviderClassListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretProviderClassListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssGet`: SecretProviderClassList
    fmt.Fprintf(os.Stdout, "Response from `SecretProviderClassListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectSecretproviderclasssGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecretProviderClassList**](SecretProviderClassList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

