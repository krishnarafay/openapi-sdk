# \SecretStoreDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameGet**](SecretStoreDetailApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/secretstores/{name} | 
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameStatusGet**](SecretStoreDetailApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameStatusGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/secretstores/{name}/status | 



## ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameGet

> SecretStore ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.SecretStoreDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretStoreDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameGet`: SecretStore
    fmt.Fprintf(os.Stdout, "Response from `SecretStoreDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SecretStore**](SecretStore.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameStatusGet

> SecretStore ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameStatusGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.SecretStoreDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameStatusGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretStoreDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameStatusGet`: SecretStore
    fmt.Fprintf(os.Stdout, "Response from `SecretStoreDetailApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresNameStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SecretStore**](SecretStore.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

