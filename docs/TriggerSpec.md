# TriggerSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**TriggerSpecConfig**](TriggerSpecConfig.md) |  | 
**Name** | **string** | name of the trigger | 
**Type** | **string** | trigger type | 
**Variables** | Pointer to [**[]VariableSpec**](VariableSpec.md) | trigger scoped variables | [optional] 

## Methods

### NewTriggerSpec

`func NewTriggerSpec(config TriggerSpecConfig, name string, type_ string, ) *TriggerSpec`

NewTriggerSpec instantiates a new TriggerSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerSpecWithDefaults

`func NewTriggerSpecWithDefaults() *TriggerSpec`

NewTriggerSpecWithDefaults instantiates a new TriggerSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *TriggerSpec) GetConfig() TriggerSpecConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TriggerSpec) GetConfigOk() (*TriggerSpecConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TriggerSpec) SetConfig(v TriggerSpecConfig)`

SetConfig sets Config field to given value.


### GetName

`func (o *TriggerSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TriggerSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TriggerSpec) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TriggerSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerSpec) SetType(v string)`

SetType sets Type field to given value.


### GetVariables

`func (o *TriggerSpec) GetVariables() []VariableSpec`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *TriggerSpec) GetVariablesOk() (*[]VariableSpec, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *TriggerSpec) SetVariables(v []VariableSpec)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *TriggerSpec) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


