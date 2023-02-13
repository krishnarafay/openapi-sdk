# StageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**StageSpecConfig**](StageSpecConfig.md) |  | 
**Name** | **string** | name of the pipeline stage | 
**Next** | Pointer to [**[]NextStage**](NextStage.md) | list of stages to be executed after this | [optional] 
**PreConditions** | Pointer to [**[]PreConditionSpec**](PreConditionSpec.md) | conditions to be evaluated before executing current stage | [optional] 
**Type** | **string** | type of pipeline stage | 
**Variables** | Pointer to [**[]VariableSpec**](VariableSpec.md) | stage scoped variables | [optional] 

## Methods

### NewStageSpec

`func NewStageSpec(config StageSpecConfig, name string, type_ string, ) *StageSpec`

NewStageSpec instantiates a new StageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageSpecWithDefaults

`func NewStageSpecWithDefaults() *StageSpec`

NewStageSpecWithDefaults instantiates a new StageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *StageSpec) GetConfig() StageSpecConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *StageSpec) GetConfigOk() (*StageSpecConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *StageSpec) SetConfig(v StageSpecConfig)`

SetConfig sets Config field to given value.


### GetName

`func (o *StageSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StageSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StageSpec) SetName(v string)`

SetName sets Name field to given value.


### GetNext

`func (o *StageSpec) GetNext() []NextStage`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *StageSpec) GetNextOk() (*[]NextStage, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *StageSpec) SetNext(v []NextStage)`

SetNext sets Next field to given value.

### HasNext

`func (o *StageSpec) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPreConditions

`func (o *StageSpec) GetPreConditions() []PreConditionSpec`

GetPreConditions returns the PreConditions field if non-nil, zero value otherwise.

### GetPreConditionsOk

`func (o *StageSpec) GetPreConditionsOk() (*[]PreConditionSpec, bool)`

GetPreConditionsOk returns a tuple with the PreConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreConditions

`func (o *StageSpec) SetPreConditions(v []PreConditionSpec)`

SetPreConditions sets PreConditions field to given value.

### HasPreConditions

`func (o *StageSpec) HasPreConditions() bool`

HasPreConditions returns a boolean if a field has been set.

### GetType

`func (o *StageSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StageSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StageSpec) SetType(v string)`

SetType sets Type field to given value.


### GetVariables

`func (o *StageSpec) GetVariables() []VariableSpec`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *StageSpec) GetVariablesOk() (*[]VariableSpec, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *StageSpec) SetVariables(v []VariableSpec)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *StageSpec) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


