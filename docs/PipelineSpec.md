# PipelineSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | flag to indicate if pipeline is active | [optional] 
**Secret** | Pointer to [**File**](File.md) |  | [optional] 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Stages** | [**[]StageSpec**](StageSpec.md) | stages in the pipeline | 
**Triggers** | Pointer to [**[]TriggerSpec**](TriggerSpec.md) | triggers for the pipeline | [optional] 
**Variables** | Pointer to [**[]VariableSpec**](VariableSpec.md) | pipeline scoped variables | [optional] 

## Methods

### NewPipelineSpec

`func NewPipelineSpec(stages []StageSpec, ) *PipelineSpec`

NewPipelineSpec instantiates a new PipelineSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineSpecWithDefaults

`func NewPipelineSpecWithDefaults() *PipelineSpec`

NewPipelineSpecWithDefaults instantiates a new PipelineSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *PipelineSpec) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PipelineSpec) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PipelineSpec) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PipelineSpec) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSecret

`func (o *PipelineSpec) GetSecret() File`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PipelineSpec) GetSecretOk() (*File, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PipelineSpec) SetSecret(v File)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PipelineSpec) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSharing

`func (o *PipelineSpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *PipelineSpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *PipelineSpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *PipelineSpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetStages

`func (o *PipelineSpec) GetStages() []StageSpec`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *PipelineSpec) GetStagesOk() (*[]StageSpec, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *PipelineSpec) SetStages(v []StageSpec)`

SetStages sets Stages field to given value.


### GetTriggers

`func (o *PipelineSpec) GetTriggers() []TriggerSpec`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *PipelineSpec) GetTriggersOk() (*[]TriggerSpec, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *PipelineSpec) SetTriggers(v []TriggerSpec)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *PipelineSpec) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetVariables

`func (o *PipelineSpec) GetVariables() []VariableSpec`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PipelineSpec) GetVariablesOk() (*[]VariableSpec, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PipelineSpec) SetVariables(v []VariableSpec)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *PipelineSpec) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


