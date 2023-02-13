# \OPAProfileListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesGet**](OPAProfileListApi.md#ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesGet) | **Get** /apis/opa.k8smgmt.io/v3/projects/{project}/opaprofiles | 



## ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesGet

> OPAProfileList ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesGet(ctx, project).Execute()





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
    resp, r, err := apiClient.OPAProfileListApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OPAProfileListApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesGet`: OPAProfileList
    fmt.Fprintf(os.Stdout, "Response from `OPAProfileListApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OPAProfileList**](OPAProfileList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

