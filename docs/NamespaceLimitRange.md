# NamespaceLimitRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to [**NamespaceLimitRangeConfig**](NamespaceLimitRangeConfig.md) |  | [optional] 
**Pod** | Pointer to [**NamespaceLimitRangeConfig**](NamespaceLimitRangeConfig.md) |  | [optional] 

## Methods

### NewNamespaceLimitRange

`func NewNamespaceLimitRange() *NamespaceLimitRange`

NewNamespaceLimitRange instantiates a new NamespaceLimitRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceLimitRangeWithDefaults

`func NewNamespaceLimitRangeWithDefaults() *NamespaceLimitRange`

NewNamespaceLimitRangeWithDefaults instantiates a new NamespaceLimitRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *NamespaceLimitRange) GetContainer() NamespaceLimitRangeConfig`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *NamespaceLimitRange) GetContainerOk() (*NamespaceLimitRangeConfig, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *NamespaceLimitRange) SetContainer(v NamespaceLimitRangeConfig)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *NamespaceLimitRange) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetPod

`func (o *NamespaceLimitRange) GetPod() NamespaceLimitRangeConfig`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *NamespaceLimitRange) GetPodOk() (*NamespaceLimitRangeConfig, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *NamespaceLimitRange) SetPod(v NamespaceLimitRangeConfig)`

SetPod sets Pod field to given value.

### HasPod

`func (o *NamespaceLimitRange) HasPod() bool`

HasPod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


