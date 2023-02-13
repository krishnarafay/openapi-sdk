# Diskcsidriver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether to enable AzureDisk CSI Driver. The default value is true. | [optional] 

## Methods

### NewDiskcsidriver

`func NewDiskcsidriver() *Diskcsidriver`

NewDiskcsidriver instantiates a new Diskcsidriver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskcsidriverWithDefaults

`func NewDiskcsidriverWithDefaults() *Diskcsidriver`

NewDiskcsidriverWithDefaults instantiates a new Diskcsidriver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *Diskcsidriver) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Diskcsidriver) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Diskcsidriver) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Diskcsidriver) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


