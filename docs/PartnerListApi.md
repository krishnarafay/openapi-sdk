# \PartnerListApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSystemK8smgmtIoV3PartnersGet**](PartnerListApi.md#ApisSystemK8smgmtIoV3PartnersGet) | **Get** /apis/system.k8smgmt.io/v3/partners | 



## ApisSystemK8smgmtIoV3PartnersGet

> PartnerList ApisSystemK8smgmtIoV3PartnersGet(ctx).Execute()





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
    resp, r, err := apiClient.PartnerListApi.ApisSystemK8smgmtIoV3PartnersGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerListApi.ApisSystemK8smgmtIoV3PartnersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSystemK8smgmtIoV3PartnersGet`: PartnerList
    fmt.Fprintf(os.Stdout, "Response from `PartnerListApi.ApisSystemK8smgmtIoV3PartnersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisSystemK8smgmtIoV3PartnersGetRequest struct via the builder pattern


### Return type

[**PartnerList**](PartnerList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

