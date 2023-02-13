# Httpproxyconfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpProxy** | Pointer to **string** | The HTTP proxy server endpoint to use. | [optional] 
**HttpsProxy** | Pointer to **string** | The HTTPS proxy server endpoint to use. | [optional] 
**NoProxy** | Pointer to **[]string** | The endpoints that should not go through proxy. | [optional] 
**TrustedCa** | Pointer to **string** | Alternative CA cert to use for connecting to proxy servers. | [optional] 

## Methods

### NewHttpproxyconfig

`func NewHttpproxyconfig() *Httpproxyconfig`

NewHttpproxyconfig instantiates a new Httpproxyconfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpproxyconfigWithDefaults

`func NewHttpproxyconfigWithDefaults() *Httpproxyconfig`

NewHttpproxyconfigWithDefaults instantiates a new Httpproxyconfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpProxy

`func (o *Httpproxyconfig) GetHttpProxy() string`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *Httpproxyconfig) GetHttpProxyOk() (*string, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *Httpproxyconfig) SetHttpProxy(v string)`

SetHttpProxy sets HttpProxy field to given value.

### HasHttpProxy

`func (o *Httpproxyconfig) HasHttpProxy() bool`

HasHttpProxy returns a boolean if a field has been set.

### GetHttpsProxy

`func (o *Httpproxyconfig) GetHttpsProxy() string`

GetHttpsProxy returns the HttpsProxy field if non-nil, zero value otherwise.

### GetHttpsProxyOk

`func (o *Httpproxyconfig) GetHttpsProxyOk() (*string, bool)`

GetHttpsProxyOk returns a tuple with the HttpsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsProxy

`func (o *Httpproxyconfig) SetHttpsProxy(v string)`

SetHttpsProxy sets HttpsProxy field to given value.

### HasHttpsProxy

`func (o *Httpproxyconfig) HasHttpsProxy() bool`

HasHttpsProxy returns a boolean if a field has been set.

### GetNoProxy

`func (o *Httpproxyconfig) GetNoProxy() []string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *Httpproxyconfig) GetNoProxyOk() (*[]string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *Httpproxyconfig) SetNoProxy(v []string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *Httpproxyconfig) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### GetTrustedCa

`func (o *Httpproxyconfig) GetTrustedCa() string`

GetTrustedCa returns the TrustedCa field if non-nil, zero value otherwise.

### GetTrustedCaOk

`func (o *Httpproxyconfig) GetTrustedCaOk() (*string, bool)`

GetTrustedCaOk returns a tuple with the TrustedCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCa

`func (o *Httpproxyconfig) SetTrustedCa(v string)`

SetTrustedCa sets TrustedCa field to given value.

### HasTrustedCa

`func (o *Httpproxyconfig) HasTrustedCa() bool`

HasTrustedCa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


