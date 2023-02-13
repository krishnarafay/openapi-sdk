# \RoleApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSystemK8smgmtIoV3RolesPost**](RoleApplyApi.md#ApisSystemK8smgmtIoV3RolesPost) | **Post** /apis/system.k8smgmt.io/v3/roles | 



## ApisSystemK8smgmtIoV3RolesPost

> Role ApisSystemK8smgmtIoV3RolesPost(ctx).Role(role).Execute()





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
    role := *openapiclient.NewRole("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewRoleSpec()) // Role | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleApplyApi.ApisSystemK8smgmtIoV3RolesPost(context.Background()).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleApplyApi.ApisSystemK8smgmtIoV3RolesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSystemK8smgmtIoV3RolesPost`: Role
    fmt.Fprintf(os.Stdout, "Response from `RoleApplyApi.ApisSystemK8smgmtIoV3RolesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisSystemK8smgmtIoV3RolesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | [**Role**](Role.md) | schema of the resource to apply | 

### Return type

[**Role**](Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

