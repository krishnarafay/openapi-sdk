# NamespaceLimitRangeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to [**ResourceQuantity**](ResourceQuantity.md) |  | [optional] 
**DefaultRequest** | Pointer to [**ResourceQuantity**](ResourceQuantity.md) |  | [optional] 
**Max** | Pointer to [**ResourceQuantity**](ResourceQuantity.md) |  | [optional] 
**Min** | Pointer to [**ResourceQuantity**](ResourceQuantity.md) |  | [optional] 
**Ratio** | Pointer to [**ResourceRatio**](ResourceRatio.md) |  | [optional] 

## Methods

### NewNamespaceLimitRangeConfig

`func NewNamespaceLimitRangeConfig() *NamespaceLimitRangeConfig`

NewNamespaceLimitRangeConfig instantiates a new NamespaceLimitRangeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceLimitRangeConfigWithDefaults

`func NewNamespaceLimitRangeConfigWithDefaults() *NamespaceLimitRangeConfig`

NewNamespaceLimitRangeConfigWithDefaults instantiates a new NamespaceLimitRangeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *NamespaceLimitRangeConfig) GetDefault() ResourceQuantity`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *NamespaceLimitRangeConfig) GetDefaultOk() (*ResourceQuantity, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *NamespaceLimitRangeConfig) SetDefault(v ResourceQuantity)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *NamespaceLimitRangeConfig) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDefaultRequest

`func (o *NamespaceLimitRangeConfig) GetDefaultRequest() ResourceQuantity`

GetDefaultRequest returns the DefaultRequest field if non-nil, zero value otherwise.

### GetDefaultRequestOk

`func (o *NamespaceLimitRangeConfig) GetDefaultRequestOk() (*ResourceQuantity, bool)`

GetDefaultRequestOk returns a tuple with the DefaultRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRequest

`func (o *NamespaceLimitRangeConfig) SetDefaultRequest(v ResourceQuantity)`

SetDefaultRequest sets DefaultRequest field to given value.

### HasDefaultRequest

`func (o *NamespaceLimitRangeConfig) HasDefaultRequest() bool`

HasDefaultRequest returns a boolean if a field has been set.

### GetMax

`func (o *NamespaceLimitRangeConfig) GetMax() ResourceQuantity`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *NamespaceLimitRangeConfig) GetMaxOk() (*ResourceQuantity, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *NamespaceLimitRangeConfig) SetMax(v ResourceQuantity)`

SetMax sets Max field to given value.

### HasMax

`func (o *NamespaceLimitRangeConfig) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *NamespaceLimitRangeConfig) GetMin() ResourceQuantity`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *NamespaceLimitRangeConfig) GetMinOk() (*ResourceQuantity, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *NamespaceLimitRangeConfig) SetMin(v ResourceQuantity)`

SetMin sets Min field to given value.

### HasMin

`func (o *NamespaceLimitRangeConfig) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetRatio

`func (o *NamespaceLimitRangeConfig) GetRatio() ResourceRatio`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *NamespaceLimitRangeConfig) GetRatioOk() (*ResourceRatio, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *NamespaceLimitRangeConfig) SetRatio(v ResourceRatio)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *NamespaceLimitRangeConfig) HasRatio() bool`

HasRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


