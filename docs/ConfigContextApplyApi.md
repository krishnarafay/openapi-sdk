# \ConfigContextApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsPost**](ConfigContextApplyApi.md#ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsPost) | **Post** /apis/eaas.k8smgmt.io/v1/projects/{project}/configcontexts | 



## ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsPost

> ConfigContext ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsPost(ctx, project).ConfigContext(configContext).Execute()





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
    configContext := *openapiclient.NewConfigContext("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewConfigContextSpec()) // ConfigContext | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigContextApplyApi.ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsPost(context.Background(), project).ConfigContext(configContext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigContextApplyApi.ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsPost`: ConfigContext
    fmt.Fprintf(os.Stdout, "Response from `ConfigContextApplyApi.ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configContext** | [**ConfigContext**](ConfigContext.md) | schema of the resource to apply | 

### Return type

[**ConfigContext**](ConfigContext.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

