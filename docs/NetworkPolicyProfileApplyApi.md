# \NetworkPolicyProfileApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSecurityK8smgmtIoV3ProjectsProjectNetworkpolicyprofilesPost**](NetworkPolicyProfileApplyApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectNetworkpolicyprofilesPost) | **Post** /apis/security.k8smgmt.io/v3/projects/{project}/networkpolicyprofiles | 



## ApisSecurityK8smgmtIoV3ProjectsProjectNetworkpolicyprofilesPost

> NetworkPolicyProfile ApisSecurityK8smgmtIoV3ProjectsProjectNetworkpolicyprofilesPost(ctx, project).NetworkPolicyProfile(networkPolicyProfile).Execute()





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
    networkPolicyProfile := *openapiclient.NewNetworkPolicyProfile("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewNetworkPolicyProfileSpec()) // NetworkPolicyProfile | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkPolicyProfileApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectNetworkpolicyprofilesPost(context.Background(), project).NetworkPolicyProfile(networkPolicyProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPolicyProfileApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectNetworkpolicyprofilesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSecurityK8smgmtIoV3ProjectsProjectNetworkpolicyprofilesPost`: NetworkPolicyProfile
    fmt.Fprintf(os.Stdout, "Response from `NetworkPolicyProfileApplyApi.ApisSecurityK8smgmtIoV3ProjectsProjectNetworkpolicyprofilesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectNetworkpolicyprofilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkPolicyProfile** | [**NetworkPolicyProfile**](NetworkPolicyProfile.md) | schema of the resource to apply | 

### Return type

[**NetworkPolicyProfile**](NetworkPolicyProfile.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

