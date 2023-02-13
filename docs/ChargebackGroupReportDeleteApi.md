# \ChargebackGroupReportDeleteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSystemK8smgmtIoV3ChargebackgroupreportsNameDelete**](ChargebackGroupReportDeleteApi.md#ApisSystemK8smgmtIoV3ChargebackgroupreportsNameDelete) | **Delete** /apis/system.k8smgmt.io/v3/chargebackgroupreports/{name} | 



## ApisSystemK8smgmtIoV3ChargebackgroupreportsNameDelete

> ApisSystemK8smgmtIoV3ChargebackgroupreportsNameDelete(ctx, name).Execute()





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
    name := "name_example" // string | name of the resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ChargebackGroupReportDeleteApi.ApisSystemK8smgmtIoV3ChargebackgroupreportsNameDelete(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargebackGroupReportDeleteApi.ApisSystemK8smgmtIoV3ChargebackgroupreportsNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSystemK8smgmtIoV3ChargebackgroupreportsNameDeleteRequest struct via the builder pattern


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

