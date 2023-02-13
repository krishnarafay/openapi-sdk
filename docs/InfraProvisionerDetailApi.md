# \InfraProvisionerDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameArtifactsGet**](InfraProvisionerDetailApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameArtifactsGet) | **Get** /apis/gitops.k8smgmt.io/v3/projects/{project}/infraprovisioners/{name}/artifacts | 
[**ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameGet**](InfraProvisionerDetailApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameGet) | **Get** /apis/gitops.k8smgmt.io/v3/projects/{project}/infraprovisioners/{name} | 



## ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameArtifactsGet

> InfraProvisioner ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.InfraProvisionerDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfraProvisionerDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameArtifactsGet`: InfraProvisioner
    fmt.Fprintf(os.Stdout, "Response from `InfraProvisionerDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InfraProvisioner**](InfraProvisioner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameGet

> InfraProvisioner ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.InfraProvisionerDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfraProvisionerDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameGet`: InfraProvisioner
    fmt.Fprintf(os.Stdout, "Response from `InfraProvisionerDetailApi.ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InfraProvisioner**](InfraProvisioner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

