# ConfigContextSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Envs** | Pointer to [**[]EnvData**](EnvData.md) | Environment variables data | [optional] 
**Files** | Pointer to [**[]File**](File.md) | File path information | [optional] 
**Variables** | Pointer to [**[]Variable**](Variable.md) | Variables data for config context | [optional] 

## Methods

### NewConfigContextSpec

`func NewConfigContextSpec() *ConfigContextSpec`

NewConfigContextSpec instantiates a new ConfigContextSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigContextSpecWithDefaults

`func NewConfigContextSpecWithDefaults() *ConfigContextSpec`

NewConfigContextSpecWithDefaults instantiates a new ConfigContextSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvs

`func (o *ConfigContextSpec) GetEnvs() []EnvData`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *ConfigContextSpec) GetEnvsOk() (*[]EnvData, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *ConfigContextSpec) SetEnvs(v []EnvData)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *ConfigContextSpec) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetFiles

`func (o *ConfigContextSpec) GetFiles() []File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ConfigContextSpec) GetFilesOk() (*[]File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ConfigContextSpec) SetFiles(v []File)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ConfigContextSpec) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetVariables

`func (o *ConfigContextSpec) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ConfigContextSpec) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ConfigContextSpec) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ConfigContextSpec) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


