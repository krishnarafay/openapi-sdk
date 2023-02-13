# \NamespaceMeshRuleListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesGet**](NamespaceMeshRuleListApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesGet) | **Get** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/namespacemeshrules | 



## ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesGet

> NamespaceMeshRuleList ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesGet(ctx, project).Execute()





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
    resp, r, err := apiClient.NamespaceMeshRuleListApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMeshRuleListApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesGet`: NamespaceMeshRuleList
    fmt.Fprintf(os.Stdout, "Response from `NamespaceMeshRuleListApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshrulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NamespaceMeshRuleList**](NamespaceMeshRuleList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

