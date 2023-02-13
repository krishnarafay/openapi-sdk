# Linuxosconfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwapFileSizeMB** | Pointer to **int32** | The size in MB of a swap file that will be created on each node. | [optional] 
**Sysctls** | Pointer to [**Sysctls**](Sysctls.md) |  | [optional] 
**TransparentHugePageDefrag** | Pointer to **string** | Valid values are always, defer, defer+madvise, madvise and never. The default is madvise. For more information see Transparent Hugepages. | [optional] 
**TransparentHugePageEnabled** | Pointer to **string** | Valid values are always, madvise, and never. The default is always. For more information see Transparent Hugepages. | [optional] 

## Methods

### NewLinuxosconfig

`func NewLinuxosconfig() *Linuxosconfig`

NewLinuxosconfig instantiates a new Linuxosconfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinuxosconfigWithDefaults

`func NewLinuxosconfigWithDefaults() *Linuxosconfig`

NewLinuxosconfigWithDefaults instantiates a new Linuxosconfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwapFileSizeMB

`func (o *Linuxosconfig) GetSwapFileSizeMB() int32`

GetSwapFileSizeMB returns the SwapFileSizeMB field if non-nil, zero value otherwise.

### GetSwapFileSizeMBOk

`func (o *Linuxosconfig) GetSwapFileSizeMBOk() (*int32, bool)`

GetSwapFileSizeMBOk returns a tuple with the SwapFileSizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapFileSizeMB

`func (o *Linuxosconfig) SetSwapFileSizeMB(v int32)`

SetSwapFileSizeMB sets SwapFileSizeMB field to given value.

### HasSwapFileSizeMB

`func (o *Linuxosconfig) HasSwapFileSizeMB() bool`

HasSwapFileSizeMB returns a boolean if a field has been set.

### GetSysctls

`func (o *Linuxosconfig) GetSysctls() Sysctls`

GetSysctls returns the Sysctls field if non-nil, zero value otherwise.

### GetSysctlsOk

`func (o *Linuxosconfig) GetSysctlsOk() (*Sysctls, bool)`

GetSysctlsOk returns a tuple with the Sysctls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysctls

`func (o *Linuxosconfig) SetSysctls(v Sysctls)`

SetSysctls sets Sysctls field to given value.

### HasSysctls

`func (o *Linuxosconfig) HasSysctls() bool`

HasSysctls returns a boolean if a field has been set.

### GetTransparentHugePageDefrag

`func (o *Linuxosconfig) GetTransparentHugePageDefrag() string`

GetTransparentHugePageDefrag returns the TransparentHugePageDefrag field if non-nil, zero value otherwise.

### GetTransparentHugePageDefragOk

`func (o *Linuxosconfig) GetTransparentHugePageDefragOk() (*string, bool)`

GetTransparentHugePageDefragOk returns a tuple with the TransparentHugePageDefrag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentHugePageDefrag

`func (o *Linuxosconfig) SetTransparentHugePageDefrag(v string)`

SetTransparentHugePageDefrag sets TransparentHugePageDefrag field to given value.

### HasTransparentHugePageDefrag

`func (o *Linuxosconfig) HasTransparentHugePageDefrag() bool`

HasTransparentHugePageDefrag returns a boolean if a field has been set.

### GetTransparentHugePageEnabled

`func (o *Linuxosconfig) GetTransparentHugePageEnabled() string`

GetTransparentHugePageEnabled returns the TransparentHugePageEnabled field if non-nil, zero value otherwise.

### GetTransparentHugePageEnabledOk

`func (o *Linuxosconfig) GetTransparentHugePageEnabledOk() (*string, bool)`

GetTransparentHugePageEnabledOk returns a tuple with the TransparentHugePageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentHugePageEnabled

`func (o *Linuxosconfig) SetTransparentHugePageEnabled(v string)`

SetTransparentHugePageEnabled sets TransparentHugePageEnabled field to given value.

### HasTransparentHugePageEnabled

`func (o *Linuxosconfig) HasTransparentHugePageEnabled() bool`

HasTransparentHugePageEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


