# \PipelineApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesPost**](PipelineApplyApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesPost) | **Post** /apis/gitops.k8smgmt.io/v3/projects/{project}/pipelines | 



## ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesPost

> Pipeline ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesPost(ctx, project).Pipeline(pipeline).Execute()





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
    pipeline := *openapiclient.NewPipeline("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewPipelineSpec([]openapiclient.StageSpec{*openapiclient.NewStageSpec(openapiclient.StageSpec_config{ApprovalConfiguration: openapiclient.NewApprovalConfiguration([]openapiclient.Approver{*openapiclient.NewApprover("UserName_example")})}, "Name_example", "Type_example")})) // Pipeline | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineApplyApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesPost(context.Background(), project).Pipeline(pipeline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineApplyApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesPost`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelineApplyApi.ApisGitopsK8smgmtIoV3ProjectsProjectPipelinesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectPipelinesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pipeline** | [**Pipeline**](Pipeline.md) | schema of the resource to apply | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

