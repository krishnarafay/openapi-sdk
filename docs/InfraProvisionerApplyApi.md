# \InfraProvisionerApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersPost**](InfraProvisionerApplyApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersPost) | **Post** /apis/gitops.k8smgmt.io/v3/projects/{project}/infraprovisioners | 



## ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersPost

> InfraProvisioner ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersPost(ctx, project).InfraProvisioner(infraProvisioner).Execute()





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
    infraProvisioner := *openapiclient.NewInfraProvisioner("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewInfraProvisionerSpec(openapiclient.InfraProvisionerSpec_config{TerrafromConfig: openapiclient.NewTerrafromConfig("Version_example")}, *openapiclient.NewFile("Name_example"), "Repository_example", "Revision_example", "Type_example")) // InfraProvisioner | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfraProvisionerApplyApi.ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersPost(context.Background(), project).InfraProvisioner(infraProvisioner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfraProvisionerApplyApi.ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersPost`: InfraProvisioner
    fmt.Fprintf(os.Stdout, "Response from `InfraProvisionerApplyApi.ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **infraProvisioner** | [**InfraProvisioner**](InfraProvisioner.md) | schema of the resource to apply | 

### Return type

[**InfraProvisioner**](InfraProvisioner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

