# \ProjectApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSystemK8smgmtIoV3ProjectsPost**](ProjectApplyApi.md#ApisSystemK8smgmtIoV3ProjectsPost) | **Post** /apis/system.k8smgmt.io/v3/projects | 



## ApisSystemK8smgmtIoV3ProjectsPost

> Project ApisSystemK8smgmtIoV3ProjectsPost(ctx).Project(project).Execute()





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
    project := *openapiclient.NewProject("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewProjectSpec()) // Project | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApplyApi.ApisSystemK8smgmtIoV3ProjectsPost(context.Background()).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApplyApi.ApisSystemK8smgmtIoV3ProjectsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSystemK8smgmtIoV3ProjectsPost`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectApplyApi.ApisSystemK8smgmtIoV3ProjectsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisSystemK8smgmtIoV3ProjectsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**Project**](Project.md) | schema of the resource to apply | 

### Return type

[**Project**](Project.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

