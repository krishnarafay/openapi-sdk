# \SecretGroupDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameArtifactsGet**](SecretGroupDetailApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameArtifactsGet) | **Get** /apis/gitops.k8smgmt.io/v3/projects/{project}/secretgroups/{name}/artifacts | 
[**ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameGet**](SecretGroupDetailApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameGet) | **Get** /apis/gitops.k8smgmt.io/v3/projects/{project}/secretgroups/{name} | 
[**ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameStatusGet**](SecretGroupDetailApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameStatusGet) | **Get** /apis/gitops.k8smgmt.io/v3/projects/{project}/secretgroups/{name}/status | 



## ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameArtifactsGet

> SecretGroup ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameArtifactsGet(ctx, project, name).Execute()





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
    name := "name_example" // string | name of the resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretGroupDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretGroupDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameArtifactsGet`: SecretGroup
    fmt.Fprintf(os.Stdout, "Response from `SecretGroupDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SecretGroup**](SecretGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameGet

> SecretGroup ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameGet(ctx, project, name).Execute()





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
    name := "name_example" // string | name of the resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretGroupDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretGroupDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameGet`: SecretGroup
    fmt.Fprintf(os.Stdout, "Response from `SecretGroupDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SecretGroup**](SecretGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameStatusGet

> SecretGroup ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameStatusGet(ctx, project, name).Execute()





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
    name := "name_example" // string | name of the resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretGroupDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameStatusGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretGroupDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameStatusGet`: SecretGroup
    fmt.Fprintf(os.Stdout, "Response from `SecretGroupDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectSecretgroupsNameStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SecretGroup**](SecretGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

