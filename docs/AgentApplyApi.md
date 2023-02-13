# \AgentApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisGitopsK8smgmtIoV3ProjectsProjectAgentsPost**](AgentApplyApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectAgentsPost) | **Post** /apis/gitops.k8smgmt.io/v3/projects/{project}/agents | 



## ApisGitopsK8smgmtIoV3ProjectsProjectAgentsPost

> Agent ApisGitopsK8smgmtIoV3ProjectsProjectAgentsPost(ctx, project).Agent(agent).Execute()





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
    agent := *openapiclient.NewAgent("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewAgentSpec(*openapiclient.NewClusterMeta("Name_example"), "Type_example")) // Agent | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentApplyApi.ApisGitopsK8smgmtIoV3ProjectsProjectAgentsPost(context.Background(), project).Agent(agent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentApplyApi.ApisGitopsK8smgmtIoV3ProjectsProjectAgentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectAgentsPost`: Agent
    fmt.Fprintf(os.Stdout, "Response from `AgentApplyApi.ApisGitopsK8smgmtIoV3ProjectsProjectAgentsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectAgentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agent** | [**Agent**](Agent.md) | schema of the resource to apply | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

