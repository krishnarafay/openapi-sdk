# \ChargebackShareDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSystemK8smgmtIoV3ChargebacksharesNameGet**](ChargebackShareDetailApi.md#ApisSystemK8smgmtIoV3ChargebacksharesNameGet) | **Get** /apis/system.k8smgmt.io/v3/chargebackshares/{name} | 



## ApisSystemK8smgmtIoV3ChargebacksharesNameGet

> ChargebackShare ApisSystemK8smgmtIoV3ChargebacksharesNameGet(ctx, name).Execute()





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
    resp, r, err := apiClient.ChargebackShareDetailApi.ApisSystemK8smgmtIoV3ChargebacksharesNameGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargebackShareDetailApi.ApisSystemK8smgmtIoV3ChargebacksharesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSystemK8smgmtIoV3ChargebacksharesNameGet`: ChargebackShare
    fmt.Fprintf(os.Stdout, "Response from `ChargebackShareDetailApi.ApisSystemK8smgmtIoV3ChargebacksharesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSystemK8smgmtIoV3ChargebacksharesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChargebackShare**](ChargebackShare.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

