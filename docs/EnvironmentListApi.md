# \EnvironmentListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsGet**](EnvironmentListApi.md#ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsGet) | **Get** /apis/eaas.k8smgmt.io/v1/projects/{project}/environments | 



## ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsGet

> EnvironmentList ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsGet(ctx, project).Execute()





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
    resp, r, err := apiClient.EnvironmentListApi.ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentListApi.ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsGet`: EnvironmentList
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentListApi.ApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisEaasK8smgmtIoV1ProjectsProjectEnvironmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentList**](EnvironmentList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

