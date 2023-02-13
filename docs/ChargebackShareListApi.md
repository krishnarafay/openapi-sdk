# \ChargebackShareListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSystemK8smgmtIoV3ChargebacksharesGet**](ChargebackShareListApi.md#ApisSystemK8smgmtIoV3ChargebacksharesGet) | **Get** /apis/system.k8smgmt.io/v3/chargebackshares | 



## ApisSystemK8smgmtIoV3ChargebacksharesGet

> ChargebackShareList ApisSystemK8smgmtIoV3ChargebacksharesGet(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChargebackShareListApi.ApisSystemK8smgmtIoV3ChargebacksharesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargebackShareListApi.ApisSystemK8smgmtIoV3ChargebacksharesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSystemK8smgmtIoV3ChargebacksharesGet`: ChargebackShareList
    fmt.Fprintf(os.Stdout, "Response from `ChargebackShareListApi.ApisSystemK8smgmtIoV3ChargebacksharesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisSystemK8smgmtIoV3ChargebacksharesGetRequest struct via the builder pattern


### Return type

[**ChargebackShareList**](ChargebackShareList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

