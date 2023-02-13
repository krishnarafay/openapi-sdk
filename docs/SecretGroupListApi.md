# \SecretGroupListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsGet**](SecretGroupListApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsGet) | **Get** /apis/gitops.k8smgmt.io/v3/projects/{project}/secretgroups | 



## ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsGet

> SecretGroupList ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsGet(ctx, project).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretGroupListApi.ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsGet(context.Background(), project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretGroupListApi.ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsGet`: SecretGroupList
    fmt.Fprintf(os.Stdout, "Response from `SecretGroupListApi.ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecretGroupList**](SecretGroupList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

