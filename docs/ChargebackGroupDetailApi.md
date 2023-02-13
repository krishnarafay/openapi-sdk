# \ChargebackGroupDetailApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisSystemK8smgmtIoV3ChargebackgroupsNameGet**](ChargebackGroupDetailApi.md#ApisSystemK8smgmtIoV3ChargebackgroupsNameGet) | **Get** /apis/system.k8smgmt.io/v3/chargebackgroups/{name} | 
[**ApisSystemK8smgmtIoV3ChargebackgroupsNameStatusGet**](ChargebackGroupDetailApi.md#ApisSystemK8smgmtIoV3ChargebackgroupsNameStatusGet) | **Get** /apis/system.k8smgmt.io/v3/chargebackgroups/{name}/status | 



## ApisSystemK8smgmtIoV3ChargebackgroupsNameGet

> ChargebackGroup ApisSystemK8smgmtIoV3ChargebackgroupsNameGet(ctx, name).Execute()





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
    resp, r, err := apiClient.ChargebackGroupDetailApi.ApisSystemK8smgmtIoV3ChargebackgroupsNameGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargebackGroupDetailApi.ApisSystemK8smgmtIoV3ChargebackgroupsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSystemK8smgmtIoV3ChargebackgroupsNameGet`: ChargebackGroup
    fmt.Fprintf(os.Stdout, "Response from `ChargebackGroupDetailApi.ApisSystemK8smgmtIoV3ChargebackgroupsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSystemK8smgmtIoV3ChargebackgroupsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChargebackGroup**](ChargebackGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisSystemK8smgmtIoV3ChargebackgroupsNameStatusGet

> ChargebackGroup ApisSystemK8smgmtIoV3ChargebackgroupsNameStatusGet(ctx, name).Execute()





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
    resp, r, err := apiClient.ChargebackGroupDetailApi.ApisSystemK8smgmtIoV3ChargebackgroupsNameStatusGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargebackGroupDetailApi.ApisSystemK8smgmtIoV3ChargebackgroupsNameStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisSystemK8smgmtIoV3ChargebackgroupsNameStatusGet`: ChargebackGroup
    fmt.Fprintf(os.Stdout, "Response from `ChargebackGroupDetailApi.ApisSystemK8smgmtIoV3ChargebackgroupsNameStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisSystemK8smgmtIoV3ChargebackgroupsNameStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChargebackGroup**](ChargebackGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

