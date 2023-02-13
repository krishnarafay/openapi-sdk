# \OPAConstraintTemplateApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesPost**](OPAConstraintTemplateApplyApi.md#ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesPost) | **Post** /apis/opa.k8smgmt.io/v3/projects/{project}/opaconstrainttemplates | 



## ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesPost

> OPAConstraintTemplate ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesPost(ctx, project).OPAConstraintTemplate(oPAConstraintTemplate).Execute()





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
    oPAConstraintTemplate := *openapiclient.NewOPAConstraintTemplate("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewOPAConstraintTemplateSpec()) // OPAConstraintTemplate | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OPAConstraintTemplateApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesPost(context.Background(), project).OPAConstraintTemplate(oPAConstraintTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OPAConstraintTemplateApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesPost`: OPAConstraintTemplate
    fmt.Fprintf(os.Stdout, "Response from `OPAConstraintTemplateApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oPAConstraintTemplate** | [**OPAConstraintTemplate**](OPAConstraintTemplate.md) | schema of the resource to apply | 

### Return type

[**OPAConstraintTemplate**](OPAConstraintTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

