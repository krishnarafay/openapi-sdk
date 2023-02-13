# \SecretStoreApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresPost**](SecretStoreApplyApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresPost) | **Post** /apis/integrations.k8smgmt.io/v3/projects/{project}/secretstores | 



## ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresPost

> SecretStore ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresPost(ctx, project).SecretStore(secretStore).Execute()





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
    secretStore := *openapiclient.NewSecretStore() // SecretStore | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretStoreApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresPost(context.Background(), project).SecretStore(secretStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretStoreApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresPost`: SecretStore
    fmt.Fprintf(os.Stdout, "Response from `SecretStoreApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectSecretstoresPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secretStore** | [**SecretStore**](SecretStore.md) | schema of the resource to apply | 

### Return type

[**SecretStore**](SecretStore.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

