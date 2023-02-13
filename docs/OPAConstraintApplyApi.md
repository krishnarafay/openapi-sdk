# \OPAConstraintApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsPost**](OPAConstraintApplyApi.md#ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsPost) | **Post** /apis/opa.k8smgmt.io/v3/projects/{project}/opaconstraints | 



## ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsPost

> OPAConstraint ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsPost(ctx, project).OPAConstraint(oPAConstraint).Execute()





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
    oPAConstraint := *openapiclient.NewOPAConstraint("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewOPAConstraintSpec()) // OPAConstraint | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OPAConstraintApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsPost(context.Background(), project).OPAConstraint(oPAConstraint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OPAConstraintApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsPost`: OPAConstraint
    fmt.Fprintf(os.Stdout, "Response from `OPAConstraintApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oPAConstraint** | [**OPAConstraint**](OPAConstraint.md) | schema of the resource to apply | 

### Return type

[**OPAConstraint**](OPAConstraint.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

