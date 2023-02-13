# \OrganizationApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSystemK8smgmtIoV3OrganizationsPost**](OrganizationApplyApi.md#ApisSystemK8smgmtIoV3OrganizationsPost) | **Post** /apis/system.k8smgmt.io/v3/organizations | 



## ApisSystemK8smgmtIoV3OrganizationsPost

> Organization ApisSystemK8smgmtIoV3OrganizationsPost(ctx).Organization(organization).Execute()





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
    organization := *openapiclient.NewOrganization() // Organization | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApplyApi.ApisSystemK8smgmtIoV3OrganizationsPost(context.Background()).Organization(organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApplyApi.ApisSystemK8smgmtIoV3OrganizationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSystemK8smgmtIoV3OrganizationsPost`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApplyApi.ApisSystemK8smgmtIoV3OrganizationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisSystemK8smgmtIoV3OrganizationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organization** | [**Organization**](Organization.md) | schema of the resource to apply | 

### Return type

[**Organization**](Organization.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

