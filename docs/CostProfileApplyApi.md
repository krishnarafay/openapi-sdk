# \CostProfileApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisCostK8smgmtIoV3ProjectsProjectCostprofilesPost**](CostProfileApplyApi.md#ApisCostK8smgmtIoV3ProjectsProjectCostprofilesPost) | **Post** /apis/cost.k8smgmt.io/v3/projects/{project}/costprofiles | 



## ApisCostK8smgmtIoV3ProjectsProjectCostprofilesPost

> CostProfile ApisCostK8smgmtIoV3ProjectsProjectCostprofilesPost(ctx, project).CostProfile(costProfile).Execute()





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
    costProfile := *openapiclient.NewCostProfile() // CostProfile | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CostProfileApplyApi.ApisCostK8smgmtIoV3ProjectsProjectCostprofilesPost(context.Background(), project).CostProfile(costProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CostProfileApplyApi.ApisCostK8smgmtIoV3ProjectsProjectCostprofilesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisCostK8smgmtIoV3ProjectsProjectCostprofilesPost`: CostProfile
    fmt.Fprintf(os.Stdout, "Response from `CostProfileApplyApi.ApisCostK8smgmtIoV3ProjectsProjectCostprofilesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisCostK8smgmtIoV3ProjectsProjectCostprofilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **costProfile** | [**CostProfile**](CostProfile.md) | schema of the resource to apply | 

### Return type

[**CostProfile**](CostProfile.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

