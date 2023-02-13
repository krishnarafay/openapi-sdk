# \PartnerApplyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSystemK8smgmtIoV3PartnersPost**](PartnerApplyApi.md#ApisSystemK8smgmtIoV3PartnersPost) | **Post** /apis/system.k8smgmt.io/v3/partners | 



## ApisSystemK8smgmtIoV3PartnersPost

> Partner ApisSystemK8smgmtIoV3PartnersPost(ctx).Partner(partner).Execute()





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
    partner := *openapiclient.NewPartner() // Partner | schema of the resource to apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerApplyApi.ApisSystemK8smgmtIoV3PartnersPost(context.Background()).Partner(partner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerApplyApi.ApisSystemK8smgmtIoV3PartnersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSystemK8smgmtIoV3PartnersPost`: Partner
    fmt.Fprintf(os.Stdout, "Response from `PartnerApplyApi.ApisSystemK8smgmtIoV3PartnersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisSystemK8smgmtIoV3PartnersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partner** | [**Partner**](Partner.md) | schema of the resource to apply | 

### Return type

[**Partner**](Partner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

