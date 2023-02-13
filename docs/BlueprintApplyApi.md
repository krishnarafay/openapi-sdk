# \BlueprintApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsPost**](BlueprintApplyApi.md#ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsPost) | **Post** /apis/infra.k8smgmt.io/v3/projects/{project}/blueprints | 



## ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsPost

> Blueprint ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsPost(ctx, project).Blueprint(blueprint).Execute()





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
    blueprint := *openapiclient.NewBlueprint("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewBlueprintSpec()) // Blueprint | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlueprintApplyApi.ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsPost(context.Background(), project).Blueprint(blueprint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlueprintApplyApi.ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsPost`: Blueprint
    fmt.Fprintf(os.Stdout, "Response from `BlueprintApplyApi.ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectBlueprintsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blueprint** | [**Blueprint**](Blueprint.md) | schema of the resource to apply | 

### Return type

[**Blueprint**](Blueprint.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

