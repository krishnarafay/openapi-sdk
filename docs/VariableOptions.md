# VariableOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the variable | [optional] 
**Override** | Pointer to [**VariableOverrideOptions**](VariableOverrideOptions.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** | Specify whether variable is read only or not, by default it is writable | [optional] 
**Required** | Pointer to **bool** | Specify whether this variable is required or optional, by default it is optional | [optional] 
**Sensitive** | Pointer to **bool** | Determines whether the value is sensitive or not, accordingly applies encryption on it | [optional] 

## Methods

### NewVariableOptions

`func NewVariableOptions() *VariableOptions`

NewVariableOptions instantiates a new VariableOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableOptionsWithDefaults

`func NewVariableOptionsWithDefaults() *VariableOptions`

NewVariableOptionsWithDefaults instantiates a new VariableOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VariableOptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VariableOptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VariableOptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VariableOptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOverride

`func (o *VariableOptions) GetOverride() VariableOverrideOptions`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *VariableOptions) GetOverrideOk() (*VariableOverrideOptions, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *VariableOptions) SetOverride(v VariableOverrideOptions)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *VariableOptions) HasOverride() bool`

HasOverride returns a boolean if a field has been set.

### GetReadOnly

`func (o *VariableOptions) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *VariableOptions) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *VariableOptions) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *VariableOptions) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRequired

`func (o *VariableOptions) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *VariableOptions) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *VariableOptions) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *VariableOptions) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSensitive

`func (o *VariableOptions) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *VariableOptions) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *VariableOptions) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *VariableOptions) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


