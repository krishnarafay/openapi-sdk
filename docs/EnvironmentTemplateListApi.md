# \EnvironmentTemplateListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesGet**](EnvironmentTemplateListApi.md#ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesGet) | **Get** /apis/eaas.k8smgmt.io/v1/projects/{project}/environmenttemplates | 



## ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesGet

> EnvironmentTemplateList ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesGet(ctx, project).Execute()





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
    resp, r, err := apiClient.EnvironmentTemplateListApi.ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentTemplateListApi.ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesGet`: EnvironmentTemplateList
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentTemplateListApi.ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentTemplateList**](EnvironmentTemplateList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

