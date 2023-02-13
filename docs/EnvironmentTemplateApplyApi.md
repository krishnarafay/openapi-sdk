# \EnvironmentTemplateApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesPost**](EnvironmentTemplateApplyApi.md#ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesPost) | **Post** /apis/eaas.k8smgmt.io/v1/projects/{project}/environmenttemplates | 



## ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesPost

> EnvironmentTemplate ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesPost(ctx, project).EnvironmentTemplate(environmentTemplate).Execute()





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
    environmentTemplate := *openapiclient.NewEnvironmentTemplate("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewEnvironmentTemplateSpec()) // EnvironmentTemplate | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentTemplateApplyApi.ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesPost(context.Background(), project).EnvironmentTemplate(environmentTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentTemplateApplyApi.ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesPost`: EnvironmentTemplate
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentTemplateApplyApi.ApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisEaasK8smgmtIoV1ProjectsProjectEnvironmenttemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentTemplate** | [**EnvironmentTemplate**](EnvironmentTemplate.md) | schema of the resource to apply | 

### Return type

[**EnvironmentTemplate**](EnvironmentTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

