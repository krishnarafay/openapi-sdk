# \OPAProfileApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesPost**](OPAProfileApplyApi.md#ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesPost) | **Post** /apis/opa.k8smgmt.io/v3/projects/{project}/opaprofiles | 



## ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesPost

> OPAProfile ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesPost(ctx, project).OPAProfile(oPAProfile).Execute()





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
    oPAProfile := *openapiclient.NewOPAProfile("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewOPAProfileSpec()) // OPAProfile | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OPAProfileApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesPost(context.Background(), project).OPAProfile(oPAProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OPAProfileApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesPost`: OPAProfile
    fmt.Fprintf(os.Stdout, "Response from `OPAProfileApplyApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisOpaK8smgmtIoV3ProjectsProjectOpaprofilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oPAProfile** | [**OPAProfile**](OPAProfile.md) | schema of the resource to apply | 

### Return type

[**OPAProfile**](OPAProfile.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

