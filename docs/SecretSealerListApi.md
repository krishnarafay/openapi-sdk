# \SecretSealerListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersGet**](SecretSealerListApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersGet) | **Get** /apis/integrations.k8smgmt.io/v3/projects/{project}/secretsealers | 



## ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersGet

> SecretSealerList ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersGet(ctx, project).Execute()





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
    resp, r, err := apiClient.SecretSealerListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretSealerListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersGet`: SecretSealerList
    fmt.Fprintf(os.Stdout, "Response from `SecretSealerListApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecretSealerList**](SecretSealerList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

