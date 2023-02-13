# \EnvironmentApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsPost**](EnvironmentApplyApi.md#ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsPost) | **Post** /apis/eaas.k8smgmt.io/v1/projects/{project}/environments | 



## ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsPost

> Environment ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsPost(ctx, project).Environment(environment).Execute()





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
    environment := *openapiclient.NewEnvironment("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewEnvironmentSpec(*openapiclient.NewResourceNameAndVersionRef("Name_example", "Version_example"))) // Environment | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApplyApi.ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsPost(context.Background(), project).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApplyApi.ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsPost`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApplyApi.ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environment** | [**Environment**](Environment.md) | schema of the resource to apply | 

### Return type

[**Environment**](Environment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

