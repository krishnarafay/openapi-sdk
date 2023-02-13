# BlueprintPSP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | flag to specify if PSP is enabled for blueprint | 
**Names** | **[]string** | names of the PSPs to be added for the blueprint | 
**Scope** | **string** | scope of the PSP for the blueprint | 

## Methods

### NewBlueprintPSP

`func NewBlueprintPSP(enabled bool, names []string, scope string, ) *BlueprintPSP`

NewBlueprintPSP instantiates a new BlueprintPSP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintPSPWithDefaults

`func NewBlueprintPSPWithDefaults() *BlueprintPSP`

NewBlueprintPSPWithDefaults instantiates a new BlueprintPSP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BlueprintPSP) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BlueprintPSP) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BlueprintPSP) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetNames

`func (o *BlueprintPSP) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *BlueprintPSP) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *BlueprintPSP) SetNames(v []string)`

SetNames sets Names field to given value.


### GetScope

`func (o *BlueprintPSP) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *BlueprintPSP) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *BlueprintPSP) SetScope(v string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


