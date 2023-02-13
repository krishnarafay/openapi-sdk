# EnvironmentTemplateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]ResourceNameAndVersionRef**](ResourceNameAndVersionRef.md) | Environment variables, file data and other variables | [optional] 
**Hooks** | Pointer to [**EnvironmentHooks**](EnvironmentHooks.md) |  | [optional] 
**Resources** | Pointer to [**[]EnvironmentResource**](EnvironmentResource.md) | Environment variables, file data and other variables | [optional] 
**SharingOptions** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Variables** | Pointer to [**[]Variable**](Variable.md) | Environment variables, file data and other variables | [optional] 
**Version** | Pointer to **string** | Environment template version | [optional] 

## Methods

### NewEnvironmentTemplateSpec

`func NewEnvironmentTemplateSpec() *EnvironmentTemplateSpec`

NewEnvironmentTemplateSpec instantiates a new EnvironmentTemplateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentTemplateSpecWithDefaults

`func NewEnvironmentTemplateSpecWithDefaults() *EnvironmentTemplateSpec`

NewEnvironmentTemplateSpecWithDefaults instantiates a new EnvironmentTemplateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *EnvironmentTemplateSpec) GetAgents() []ResourceNameAndVersionRef`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *EnvironmentTemplateSpec) GetAgentsOk() (*[]ResourceNameAndVersionRef, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *EnvironmentTemplateSpec) SetAgents(v []ResourceNameAndVersionRef)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *EnvironmentTemplateSpec) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetHooks

`func (o *EnvironmentTemplateSpec) GetHooks() EnvironmentHooks`

GetHooks returns the Hooks field if non-nil, zero value otherwise.

### GetHooksOk

`func (o *EnvironmentTemplateSpec) GetHooksOk() (*EnvironmentHooks, bool)`

GetHooksOk returns a tuple with the Hooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooks

`func (o *EnvironmentTemplateSpec) SetHooks(v EnvironmentHooks)`

SetHooks sets Hooks field to given value.

### HasHooks

`func (o *EnvironmentTemplateSpec) HasHooks() bool`

HasHooks returns a boolean if a field has been set.

### GetResources

`func (o *EnvironmentTemplateSpec) GetResources() []EnvironmentResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *EnvironmentTemplateSpec) GetResourcesOk() (*[]EnvironmentResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *EnvironmentTemplateSpec) SetResources(v []EnvironmentResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *EnvironmentTemplateSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSharingOptions

`func (o *EnvironmentTemplateSpec) GetSharingOptions() SharingSpec`

GetSharingOptions returns the SharingOptions field if non-nil, zero value otherwise.

### GetSharingOptionsOk

`func (o *EnvironmentTemplateSpec) GetSharingOptionsOk() (*SharingSpec, bool)`

GetSharingOptionsOk returns a tuple with the SharingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingOptions

`func (o *EnvironmentTemplateSpec) SetSharingOptions(v SharingSpec)`

SetSharingOptions sets SharingOptions field to given value.

### HasSharingOptions

`func (o *EnvironmentTemplateSpec) HasSharingOptions() bool`

HasSharingOptions returns a boolean if a field has been set.

### GetVariables

`func (o *EnvironmentTemplateSpec) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *EnvironmentTemplateSpec) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *EnvironmentTemplateSpec) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *EnvironmentTemplateSpec) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetVersion

`func (o *EnvironmentTemplateSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EnvironmentTemplateSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EnvironmentTemplateSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EnvironmentTemplateSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


