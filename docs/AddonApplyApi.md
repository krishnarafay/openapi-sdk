# \AddonApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisInfraK8smgmtIoV3ProjectsProjectAddonsPost**](AddonApplyApi.md#ApisInfraK8smgmtIoV3ProjectsProjectAddonsPost) | **Post** /apis/infra.k8smgmt.io/v3/projects/{project}/addons | 



## ApisInfraK8smgmtIoV3ProjectsProjectAddonsPost

> Addon ApisInfraK8smgmtIoV3ProjectsProjectAddonsPost(ctx, project).Addon(addon).Execute()





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
    addon := *openapiclient.NewAddon("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewAddonSpec(*openapiclient.NewArtifactSpec(openapiclient.ArtifactSpec_artifact{HelmInCatalog: openapiclient.NewHelmInCatalog("ChartName_example", "ChartVersion_example")}, "Type_example"), "Namespace_example", *openapiclient.NewSharingSpec(false), "Version_example")) // Addon | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddonApplyApi.ApisInfraK8smgmtIoV3ProjectsProjectAddonsPost(context.Background(), project).Addon(addon).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonApplyApi.ApisInfraK8smgmtIoV3ProjectsProjectAddonsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectAddonsPost`: Addon
    fmt.Fprintf(os.Stdout, "Response from `AddonApplyApi.ApisInfraK8smgmtIoV3ProjectsProjectAddonsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectAddonsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addon** | [**Addon**](Addon.md) | schema of the resource to apply | 

### Return type

[**Addon**](Addon.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

