# \NamespaceMeshPolicyApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesPost**](NamespaceMeshPolicyApplyApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesPost) | **Post** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/namespacemeshpolicies | 
[**ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysPost**](NamespaceMeshPolicyApplyApi.md#ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysPost) | **Post** /apis/servicemesh.k8smgmt.io/v3/projects/{project}/namespacemeshpolicys | 



## ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesPost

> NamespaceMeshPolicy ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesPost(ctx, project).NamespaceMeshPolicy(namespaceMeshPolicy).Execute()





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
    namespaceMeshPolicy := *openapiclient.NewNamespaceMeshPolicy("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewNamespaceMeshPolicySpec()) // NamespaceMeshPolicy | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceMeshPolicyApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesPost(context.Background(), project).NamespaceMeshPolicy(namespaceMeshPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMeshPolicyApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesPost`: NamespaceMeshPolicy
    fmt.Fprintf(os.Stdout, "Response from `NamespaceMeshPolicyApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpoliciesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespaceMeshPolicy** | [**NamespaceMeshPolicy**](NamespaceMeshPolicy.md) | schema of the resource to apply | 

### Return type

[**NamespaceMeshPolicy**](NamespaceMeshPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysPost

> NamespaceMeshPolicy ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysPost(ctx, project).NamespaceMeshPolicy(namespaceMeshPolicy).Execute()





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
    namespaceMeshPolicy := *openapiclient.NewNamespaceMeshPolicy("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewNamespaceMeshPolicySpec()) // NamespaceMeshPolicy | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespaceMeshPolicyApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysPost(context.Background(), project).NamespaceMeshPolicy(namespaceMeshPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMeshPolicyApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysPost`: NamespaceMeshPolicy
    fmt.Fprintf(os.Stdout, "Response from `NamespaceMeshPolicyApplyApi.ApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisServicemeshK8smgmtIoV3ProjectsProjectNamespacemeshpolicysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespaceMeshPolicy** | [**NamespaceMeshPolicy**](NamespaceMeshPolicy.md) | schema of the resource to apply | 

### Return type

[**NamespaceMeshPolicy**](NamespaceMeshPolicy.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

