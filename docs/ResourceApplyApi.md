# \ResourceApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisEaasK8smgmtIoV1ProjectsProjectResourcesPost**](ResourceApplyApi.md#ApisEaasK8smgmtIoV1ProjectsProjectResourcesPost) | **Post** /apis/eaas.k8smgmt.io/v1/projects/{project}/resources | 



## ApisEaasK8smgmtIoV1ProjectsProjectResourcesPost

> Resource ApisEaasK8smgmtIoV1ProjectsProjectResourcesPost(ctx, project).Resource(resource).Execute()





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
    resource := *openapiclient.NewResource("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewResourceSpec([]openapiclient.Variable{*openapiclient.NewVariable()})) // Resource | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceApplyApi.ApisEaasK8smgmtIoV1ProjectsProjectResourcesPost(context.Background(), project).Resource(resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceApplyApi.ApisEaasK8smgmtIoV1ProjectsProjectResourcesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisEaasK8smgmtIoV1ProjectsProjectResourcesPost`: Resource
    fmt.Fprintf(os.Stdout, "Response from `ResourceApplyApi.ApisEaasK8smgmtIoV1ProjectsProjectResourcesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisEaasK8smgmtIoV1ProjectsProjectResourcesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resource** | [**Resource**](Resource.md) | schema of the resource to apply | 

### Return type

[**Resource**](Resource.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

