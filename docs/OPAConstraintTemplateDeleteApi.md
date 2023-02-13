# \OPAConstraintTemplateDeleteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameDelete**](OPAConstraintTemplateDeleteApi.md#ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameDelete) | **Delete** /apis/opa.k8smgmt.io/v3/projects/{project}/opaconstrainttemplates/{name} | 



## ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameDelete

> ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameDelete(ctx, project, name).Execute()





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
    r, err := apiClient.OPAConstraintTemplateDeleteApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameDelete(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OPAConstraintTemplateDeleteApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisOpaK8smgmtIoV3ProjectsProjectOpaconstrainttemplatesNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

