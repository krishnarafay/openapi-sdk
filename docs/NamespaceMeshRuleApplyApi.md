# \NamespaceMeshRuleApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPost**](NamespaceMeshRuleApplyApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPost) | **Post** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/namespacemeshrules | 



## ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPost

> NamespaceMeshRule ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPost(ctx, project).NamespaceMeshRule(namespaceMeshRule).Execute()





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
    namespaceMeshRule := *openapiclient.NewNamespaceMeshRule("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewNamespaceMeshRuleSpec()) // NamespaceMeshRule | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceMeshRuleApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPost(context.Background(), project).NamespaceMeshRule(namespaceMeshRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMeshRuleApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPost`: NamespaceMeshRule
    fmt.Fprintf(os.Stdout, "Response from `NamespaceMeshRuleApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespaceMeshRule** | [**NamespaceMeshRule**](NamespaceMeshRule.md) | schema of the resource to apply | 

### Return type

[**NamespaceMeshRule**](NamespaceMeshRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

