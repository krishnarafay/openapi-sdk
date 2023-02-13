# \CatalogApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisAppsK8smgmtIoV3ProjectsProjectCatalogsPost**](CatalogApplyApi.md#ApisAppsK8smgmtIoV3ProjectsProjectCatalogsPost) | **Post** /apis/apps.k8smgmt.io/v3/projects/{project}/catalogs | 



## ApisAppsK8smgmtIoV3ProjectsProjectCatalogsPost

> Catalog ApisAppsK8smgmtIoV3ProjectsProjectCatalogsPost(ctx, project).Catalog(catalog).Execute()





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
    catalog := *openapiclient.NewCatalog("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewCatalogSpec("Repository_example", "Type_example")) // Catalog | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogApplyApi.ApisAppsK8smgmtIoV3ProjectsProjectCatalogsPost(context.Background(), project).Catalog(catalog).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApplyApi.ApisAppsK8smgmtIoV3ProjectsProjectCatalogsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsK8smgmtIoV3ProjectsProjectCatalogsPost`: Catalog
    fmt.Fprintf(os.Stdout, "Response from `CatalogApplyApi.ApisAppsK8smgmtIoV3ProjectsProjectCatalogsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsK8smgmtIoV3ProjectsProjectCatalogsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **catalog** | [**Catalog**](Catalog.md) | schema of the resource to apply | 

### Return type

[**Catalog**](Catalog.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

