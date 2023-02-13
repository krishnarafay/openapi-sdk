# ResourceTemplateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]ResourceNameAndVersionRef**](ResourceNameAndVersionRef.md) | Reference to agents to process resource template | [optional] 
**Contexts** | Pointer to [**[]ResourceNameAndVersionRef**](ResourceNameAndVersionRef.md) | Reference to config context data associated with resource templates | [optional] 
**Hooks** | Pointer to [**ResourceHooks**](ResourceHooks.md) |  | [optional] 
**Provider** | Pointer to **string** | Specify the resource template provider | [optional] 
**ProviderOptions** | Pointer to [**ResourceTemplateProviderOptions**](ResourceTemplateProviderOptions.md) |  | [optional] 
**RepositoryOptions** | Pointer to [**ResourceTemplateRepositoryOptions**](ResourceTemplateRepositoryOptions.md) |  | [optional] 
**SharingOptions** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Variables** | Pointer to [**[]Variable**](Variable.md) | Variables associated with resource template | [optional] 
**Version** | Pointer to **string** | Version of the resource template | [optional] 

## Methods

### NewResourceTemplateSpec

`func NewResourceTemplateSpec() *ResourceTemplateSpec`

NewResourceTemplateSpec instantiates a new ResourceTemplateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTemplateSpecWithDefaults

`func NewResourceTemplateSpecWithDefaults() *ResourceTemplateSpec`

NewResourceTemplateSpecWithDefaults instantiates a new ResourceTemplateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *ResourceTemplateSpec) GetAgents() []ResourceNameAndVersionRef`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *ResourceTemplateSpec) GetAgentsOk() (*[]ResourceNameAndVersionRef, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *ResourceTemplateSpec) SetAgents(v []ResourceNameAndVersionRef)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *ResourceTemplateSpec) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetContexts

`func (o *ResourceTemplateSpec) GetContexts() []ResourceNameAndVersionRef`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *ResourceTemplateSpec) GetContextsOk() (*[]ResourceNameAndVersionRef, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *ResourceTemplateSpec) SetContexts(v []ResourceNameAndVersionRef)`

SetContexts sets Contexts field to given value.

### HasContexts

`func (o *ResourceTemplateSpec) HasContexts() bool`

HasContexts returns a boolean if a field has been set.

### GetHooks

`func (o *ResourceTemplateSpec) GetHooks() ResourceHooks`

GetHooks returns the Hooks field if non-nil, zero value otherwise.

### GetHooksOk

`func (o *ResourceTemplateSpec) GetHooksOk() (*ResourceHooks, bool)`

GetHooksOk returns a tuple with the Hooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooks

`func (o *ResourceTemplateSpec) SetHooks(v ResourceHooks)`

SetHooks sets Hooks field to given value.

### HasHooks

`func (o *ResourceTemplateSpec) HasHooks() bool`

HasHooks returns a boolean if a field has been set.

### GetProvider

`func (o *ResourceTemplateSpec) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ResourceTemplateSpec) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ResourceTemplateSpec) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ResourceTemplateSpec) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderOptions

`func (o *ResourceTemplateSpec) GetProviderOptions() ResourceTemplateProviderOptions`

GetProviderOptions returns the ProviderOptions field if non-nil, zero value otherwise.

### GetProviderOptionsOk

`func (o *ResourceTemplateSpec) GetProviderOptionsOk() (*ResourceTemplateProviderOptions, bool)`

GetProviderOptionsOk returns a tuple with the ProviderOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderOptions

`func (o *ResourceTemplateSpec) SetProviderOptions(v ResourceTemplateProviderOptions)`

SetProviderOptions sets ProviderOptions field to given value.

### HasProviderOptions

`func (o *ResourceTemplateSpec) HasProviderOptions() bool`

HasProviderOptions returns a boolean if a field has been set.

### GetRepositoryOptions

`func (o *ResourceTemplateSpec) GetRepositoryOptions() ResourceTemplateRepositoryOptions`

GetRepositoryOptions returns the RepositoryOptions field if non-nil, zero value otherwise.

### GetRepositoryOptionsOk

`func (o *ResourceTemplateSpec) GetRepositoryOptionsOk() (*ResourceTemplateRepositoryOptions, bool)`

GetRepositoryOptionsOk returns a tuple with the RepositoryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryOptions

`func (o *ResourceTemplateSpec) SetRepositoryOptions(v ResourceTemplateRepositoryOptions)`

SetRepositoryOptions sets RepositoryOptions field to given value.

### HasRepositoryOptions

`func (o *ResourceTemplateSpec) HasRepositoryOptions() bool`

HasRepositoryOptions returns a boolean if a field has been set.

### GetSharingOptions

`func (o *ResourceTemplateSpec) GetSharingOptions() SharingSpec`

GetSharingOptions returns the SharingOptions field if non-nil, zero value otherwise.

### GetSharingOptionsOk

`func (o *ResourceTemplateSpec) GetSharingOptionsOk() (*SharingSpec, bool)`

GetSharingOptionsOk returns a tuple with the SharingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingOptions

`func (o *ResourceTemplateSpec) SetSharingOptions(v SharingSpec)`

SetSharingOptions sets SharingOptions field to given value.

### HasSharingOptions

`func (o *ResourceTemplateSpec) HasSharingOptions() bool`

HasSharingOptions returns a boolean if a field has been set.

### GetVariables

`func (o *ResourceTemplateSpec) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ResourceTemplateSpec) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ResourceTemplateSpec) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ResourceTemplateSpec) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetVersion

`func (o *ResourceTemplateSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResourceTemplateSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResourceTemplateSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResourceTemplateSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


