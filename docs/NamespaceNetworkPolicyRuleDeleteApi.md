# \NamespaceNetworkPolicyRuleDeleteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameDelete**](NamespaceNetworkPolicyRuleDeleteApi.md#ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameDelete) | **Delete** /apis/security.k8smgmt.io/v3/projects/{project}/namespacenetworkpolicyrules/{name} | 



## ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameDelete

> ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameDelete(ctx, project, name).Execute()





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
    r, err := apiClient.NamespaceNetworkPolicyRuleDeleteApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameDelete(context.Background(), project, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespaceNetworkPolicyRuleDeleteApi.ApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApisSecurityK8smgmtIoV3ProjectsProjectNamespacenetworkpolicyrulesNameDeleteRequest struct via the builder pattern


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

