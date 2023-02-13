# \SecretGroupApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsPost**](SecretGroupApplyApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsPost) | **Post** /apis/gitops.k8smgmt.io/v3/projects/{project}/secretgroups | 



## ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsPost

> SecretGroup ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsPost(ctx, project).SecretGroup(secretGroup).Execute()





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
    secretGroup := *openapiclient.NewSecretGroup("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewSecretGroupSpec([]openapiclient.Secret{*openapiclient.NewSecret("FilePath_example", "Secret_example")})) // SecretGroup | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretGroupApplyApi.ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsPost(context.Background(), project).SecretGroup(secretGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretGroupApplyApi.ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsPost`: SecretGroup
    fmt.Fprintf(os.Stdout, "Response from `SecretGroupApplyApi.ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secretGroup** | [**SecretGroup**](SecretGroup.md) | schema of the resource to apply | 

### Return type

[**SecretGroup**](SecretGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

