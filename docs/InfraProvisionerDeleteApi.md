# \InfraProvisionerDeleteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameDelete**](InfraProvisionerDeleteApi.md#ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameDelete) | **Delete** /apis/gitops.k8smgmt.io/v3/projects/{project}/infraprovisioners/{name} | 



## ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameDelete

> ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameDelete(ctx, project, name).Execute()





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
    r, err := apiClient.InfraProvisionerDeleteApi.ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameDelete(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfraProvisionerDeleteApi.ApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisGitopsK8smgmtIoV3ProjectsProjectInfraprovisionersNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

