# \ChargebackShareApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSystemK8smgmtIoV3ChargebacksharesPost**](ChargebackShareApplyApi.md#ApisSystemK8smgmtIoV3ChargebacksharesPost) | **Post** /apis/system.k8smgmt.io/v3/chargebackshares | 



## ApisSystemK8smgmtIoV3ChargebacksharesPost

> ChargebackShare ApisSystemK8smgmtIoV3ChargebacksharesPost(ctx).ChargebackShare(chargebackShare).Execute()





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
    chargebackShare := *openapiclient.NewChargebackShare("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewChargebackShareSpec()) // ChargebackShare | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChargebackShareApplyApi.ApisSystemK8smgmtIoV3ChargebacksharesPost(context.Background()).ChargebackShare(chargebackShare).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargebackShareApplyApi.ApisSystemK8smgmtIoV3ChargebacksharesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSystemK8smgmtIoV3ChargebacksharesPost`: ChargebackShare
    fmt.Fprintf(os.Stdout, "Response from `ChargebackShareApplyApi.ApisSystemK8smgmtIoV3ChargebacksharesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisSystemK8smgmtIoV3ChargebacksharesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chargebackShare** | [**ChargebackShare**](ChargebackShare.md) | schema of the resource to apply | 

### Return type

[**ChargebackShare**](ChargebackShare.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

