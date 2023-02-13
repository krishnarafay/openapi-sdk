# \ResourceTemplateApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisEaasK8smgmtIoV1ProjectsProjectResourcetemplatesPost**](ResourceTemplateApplyApi.md#ApisEaasK8smgmtIoV1ProjectsProjectResourcetemplatesPost) | **Post** /apis/eaas.k8smgmt.io/v1/projects/{project}/resourcetemplates | 



## ApisEaasK8smgmtIoV1ProjectsProjectResourcetemplatesPost

> ResourceTemplate ApisEaasK8smgmtIoV1ProjectsProjectResourcetemplatesPost(ctx, project).ResourceTemplate(resourceTemplate).Execute()





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
    resourceTemplate := *openapiclient.NewResourceTemplate("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewResourceTemplateSpec()) // ResourceTemplate | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceTemplateApplyApi.ApisEaasK8smgmtIoV1ProjectsProjectResourcetemplatesPost(context.Background(), project).ResourceTemplate(resourceTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceTemplateApplyApi.ApisEaasK8smgmtIoV1ProjectsProjectResourcetemplatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisEaasK8smgmtIoV1ProjectsProjectResourcetemplatesPost`: ResourceTemplate
    fmt.Fprintf(os.Stdout, "Response from `ResourceTemplateApplyApi.ApisEaasK8smgmtIoV1ProjectsProjectResourcetemplatesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisEaasK8smgmtIoV1ProjectsProjectResourcetemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceTemplate** | [**ResourceTemplate**](ResourceTemplate.md) | schema of the resource to apply | 

### Return type

[**ResourceTemplate**](ResourceTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

