# \ConfigContextListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsGet**](ConfigContextListApi.md#ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsGet) | **Get** /apis/eaas.k8smgmt.io/v1/projects/{project}/configcontexts | 



## ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsGet

> ConfigContextList ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsGet(ctx, project).Execute()





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
    resp, r, err := apiClient.ConfigContextListApi.ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigContextListApi.ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsGet`: ConfigContextList
    fmt.Fprintf(os.Stdout, "Response from `ConfigContextListApi.ApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisEaasK8smgmtIoV1ProjectsProjectConfigcontextsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigContextList**](ConfigContextList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

