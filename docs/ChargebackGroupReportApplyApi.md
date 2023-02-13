# \ChargebackGroupReportApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSystemK8smgmtIoV3ChargebackgroupreportsPost**](ChargebackGroupReportApplyApi.md#ApisSystemK8smgmtIoV3ChargebackgroupreportsPost) | **Post** /apis/system.k8smgmt.io/v3/chargebackgroupreports | 



## ApisSystemK8smgmtIoV3ChargebackgroupreportsPost

> ChargebackGroupReport ApisSystemK8smgmtIoV3ChargebackgroupreportsPost(ctx).ChargebackGroupReport(chargebackGroupReport).Execute()





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
    chargebackGroupReport := *openapiclient.NewChargebackGroupReport("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example", "Project_example"), *openapiclient.NewChargebackGroupReportSpec()) // ChargebackGroupReport | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChargebackGroupReportApplyApi.ApisSystemK8smgmtIoV3ChargebackgroupreportsPost(context.Background()).ChargebackGroupReport(chargebackGroupReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargebackGroupReportApplyApi.ApisSystemK8smgmtIoV3ChargebackgroupreportsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSystemK8smgmtIoV3ChargebackgroupreportsPost`: ChargebackGroupReport
    fmt.Fprintf(os.Stdout, "Response from `ChargebackGroupReportApplyApi.ApisSystemK8smgmtIoV3ChargebackgroupreportsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisSystemK8smgmtIoV3ChargebackgroupreportsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chargebackGroupReport** | [**ChargebackGroupReport**](ChargebackGroupReport.md) | schema of the resource to apply | 

### Return type

[**ChargebackGroupReport**](ChargebackGroupReport.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

