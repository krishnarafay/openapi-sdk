# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | flag to specify if PSP is enabled for blueprint | [optional] 
**Name** | Pointer to **string** | name of the opa policy | [optional] 
**Version** | Pointer to **string** | version of the opa policy | [optional] 

## Methods

### NewPolicy

`func NewPolicy() *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *Policy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Policy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Policy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Policy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *Policy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Policy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Policy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Policy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *Policy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Policy) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Policy) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Policy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


