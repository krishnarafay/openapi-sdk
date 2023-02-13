# EnvironmentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Template** | [**ResourceNameAndVersionRef**](ResourceNameAndVersionRef.md) |  | 
**Variables** | Pointer to [**[]Variable**](Variable.md) | Variables data for config context | [optional] 

## Methods

### NewEnvironmentSpec

`func NewEnvironmentSpec(template ResourceNameAndVersionRef, ) *EnvironmentSpec`

NewEnvironmentSpec instantiates a new EnvironmentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentSpecWithDefaults

`func NewEnvironmentSpecWithDefaults() *EnvironmentSpec`

NewEnvironmentSpecWithDefaults instantiates a new EnvironmentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharing

`func (o *EnvironmentSpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *EnvironmentSpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *EnvironmentSpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *EnvironmentSpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetTemplate

`func (o *EnvironmentSpec) GetTemplate() ResourceNameAndVersionRef`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EnvironmentSpec) GetTemplateOk() (*ResourceNameAndVersionRef, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EnvironmentSpec) SetTemplate(v ResourceNameAndVersionRef)`

SetTemplate sets Template field to given value.


### GetVariables

`func (o *EnvironmentSpec) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *EnvironmentSpec) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *EnvironmentSpec) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *EnvironmentSpec) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


