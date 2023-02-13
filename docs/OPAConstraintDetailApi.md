# \OPAConstraintDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameArtifactsGet**](OPAConstraintDetailApi.md#ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameArtifactsGet) | **Get** /apis/opa.k8smgmt.io/v3/projects/{project}/opaconstraints/{name}/artifacts | 
[**ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameGet**](OPAConstraintDetailApi.md#ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameGet) | **Get** /apis/opa.k8smgmt.io/v3/projects/{project}/opaconstraints/{name} | 



## ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameArtifactsGet

> OPAConstraint ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameArtifactsGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.OPAConstraintDetailApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameArtifactsGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OPAConstraintDetailApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameArtifactsGet`: OPAConstraint
    fmt.Fprintf(os.Stdout, "Response from `OPAConstraintDetailApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OPAConstraint**](OPAConstraint.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameGet

> OPAConstraint ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.OPAConstraintDetailApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OPAConstraintDetailApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameGet`: OPAConstraint
    fmt.Fprintf(os.Stdout, "Response from `OPAConstraintDetailApi.ApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisOpaK8smgmtIoV3ProjectsProjectOpaconstraintsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OPAConstraint**](OPAConstraint.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

