# \ChargebackGroupApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSystemK8smgmtIoV3ChargebackgroupsPost**](ChargebackGroupApplyApi.md#ApisSystemK8smgmtIoV3ChargebackgroupsPost) | **Post** /apis/system.k8smgmt.io/v3/chargebackgroups | 



## ApisSystemK8smgmtIoV3ChargebackgroupsPost

> ChargebackGroup ApisSystemK8smgmtIoV3ChargebackgroupsPost(ctx).ChargebackGroup(chargebackGroup).Execute()





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
    chargebackGroup := *openapiclient.NewChargebackGroup("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewChargebackGroupSpec()) // ChargebackGroup | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChargebackGroupApplyApi.ApisSystemK8smgmtIoV3ChargebackgroupsPost(context.Background()).ChargebackGroup(chargebackGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargebackGroupApplyApi.ApisSystemK8smgmtIoV3ChargebackgroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSystemK8smgmtIoV3ChargebackgroupsPost`: ChargebackGroup
    fmt.Fprintf(os.Stdout, "Response from `ChargebackGroupApplyApi.ApisSystemK8smgmtIoV3ChargebackgroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisSystemK8smgmtIoV3ChargebackgroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chargebackGroup** | [**ChargebackGroup**](ChargebackGroup.md) | schema of the resource to apply | 

### Return type

[**ChargebackGroup**](ChargebackGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

