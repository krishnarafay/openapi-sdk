# \BlueprintDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsNameGet**](BlueprintDetailApi.md#ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsNameGet) | **Get** /apis/infra.k8smgmt.io/v3/projects/{project}/blueprints/{name} | 



## ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsNameGet

> Blueprint ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsNameGet(ctx, project, name).Execute()





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
    resp, r, err := apiClient.BlueprintDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsNameGet(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlueprintDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsNameGet`: Blueprint
    fmt.Fprintf(os.Stdout, "Response from `BlueprintDetailApi.ApisInfraK8smgmtIoV3ProjectsProjectBlueprintsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | project of the resource | 
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisInfraK8smgmtIoV3ProjectsProjectBlueprintsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Blueprint**](Blueprint.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

