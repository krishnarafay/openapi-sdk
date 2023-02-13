# \PipelineListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesGet**](PipelineListApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesGet) | **Get** /apis/gitops.k8smgmt.io/v3/projects/{project}/pipelines | 



## ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesGet

> PipelineList ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesGet(ctx, project).Execute()





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
    resp, r, err := apiClient.PipelineListApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineListApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesGet`: PipelineList
    fmt.Fprintf(os.Stdout, "Response from `PipelineListApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectPipelinesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PipelineList**](PipelineList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

