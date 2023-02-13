# AddonProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **string** | Addon Profile Config | [optional] 
**Enabled** | Pointer to **bool** | Whether the addon profile is enabled | [optional] 

## Methods

### NewAddonProfile

`func NewAddonProfile() *AddonProfile`

NewAddonProfile instantiates a new AddonProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonProfileWithDefaults

`func NewAddonProfileWithDefaults() *AddonProfile`

NewAddonProfileWithDefaults instantiates a new AddonProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *AddonProfile) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddonProfile) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddonProfile) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddonProfile) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *AddonProfile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddonProfile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddonProfile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddonProfile) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


