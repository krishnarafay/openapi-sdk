# \SecretSealerApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersPost**](SecretSealerApplyApi.md#ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersPost) | **Post** /apis/integrations.k8smgmt.io/v3/projects/{project}/secretsealers | 



## ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersPost

> SecretSealer ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersPost(ctx, project).SecretSealer(secretSealer).Execute()





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
    secretSealer := *openapiclient.NewSecretSealer("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewSecretSealerSpec()) // SecretSealer | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretSealerApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersPost(context.Background(), project).SecretSealer(secretSealer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretSealerApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersPost`: SecretSealer
    fmt.Fprintf(os.Stdout, "Response from `SecretSealerApplyApi.ApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisIntegrationsK8smgmtIoV3ProjectsProjectSecretsealersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secretSealer** | [**SecretSealer**](SecretSealer.md) | schema of the resource to apply | 

### Return type

[**SecretSealer**](SecretSealer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

