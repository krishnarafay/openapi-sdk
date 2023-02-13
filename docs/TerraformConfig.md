# TerraformConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendFilePath** | Pointer to [**File**](File.md) |  | [optional] 
**BackendVars** | Pointer to [**[]KeyValue**](KeyValue.md) | terraform state store backend variables | [optional] 
**EnvVars** | Pointer to [**[]KeyValue**](KeyValue.md) | environment variables | [optional] 
**InputVars** | Pointer to [**[]KeyValue**](KeyValue.md) | terraform input variables | [optional] 
**SecretGroups** | Pointer to **[]string** | Pipeline secrets groups | [optional] 
**TfVarsFilePath** | Pointer to [**File**](File.md) |  | [optional] 
**Version** | **string** | terraform version | 

## Methods

### NewTerraformConfig

`func NewTerraformConfig(version string, ) *TerraformConfig`

NewTerraformConfig instantiates a new TerraformConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformConfigWithDefaults

`func NewTerraformConfigWithDefaults() *TerraformConfig`

NewTerraformConfigWithDefaults instantiates a new TerraformConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendFilePath

`func (o *TerraformConfig) GetBackendFilePath() File`

GetBackendFilePath returns the BackendFilePath field if non-nil, zero value otherwise.

### GetBackendFilePathOk

`func (o *TerraformConfig) GetBackendFilePathOk() (*File, bool)`

GetBackendFilePathOk returns a tuple with the BackendFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendFilePath

`func (o *TerraformConfig) SetBackendFilePath(v File)`

SetBackendFilePath sets BackendFilePath field to given value.

### HasBackendFilePath

`func (o *TerraformConfig) HasBackendFilePath() bool`

HasBackendFilePath returns a boolean if a field has been set.

### GetBackendVars

`func (o *TerraformConfig) GetBackendVars() []KeyValue`

GetBackendVars returns the BackendVars field if non-nil, zero value otherwise.

### GetBackendVarsOk

`func (o *TerraformConfig) GetBackendVarsOk() (*[]KeyValue, bool)`

GetBackendVarsOk returns a tuple with the BackendVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendVars

`func (o *TerraformConfig) SetBackendVars(v []KeyValue)`

SetBackendVars sets BackendVars field to given value.

### HasBackendVars

`func (o *TerraformConfig) HasBackendVars() bool`

HasBackendVars returns a boolean if a field has been set.

### GetEnvVars

`func (o *TerraformConfig) GetEnvVars() []KeyValue`

GetEnvVars returns the EnvVars field if non-nil, zero value otherwise.

### GetEnvVarsOk

`func (o *TerraformConfig) GetEnvVarsOk() (*[]KeyValue, bool)`

GetEnvVarsOk returns a tuple with the EnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVars

`func (o *TerraformConfig) SetEnvVars(v []KeyValue)`

SetEnvVars sets EnvVars field to given value.

### HasEnvVars

`func (o *TerraformConfig) HasEnvVars() bool`

HasEnvVars returns a boolean if a field has been set.

### GetInputVars

`func (o *TerraformConfig) GetInputVars() []KeyValue`

GetInputVars returns the InputVars field if non-nil, zero value otherwise.

### GetInputVarsOk

`func (o *TerraformConfig) GetInputVarsOk() (*[]KeyValue, bool)`

GetInputVarsOk returns a tuple with the InputVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputVars

`func (o *TerraformConfig) SetInputVars(v []KeyValue)`

SetInputVars sets InputVars field to given value.

### HasInputVars

`func (o *TerraformConfig) HasInputVars() bool`

HasInputVars returns a boolean if a field has been set.

### GetSecretGroups

`func (o *TerraformConfig) GetSecretGroups() []string`

GetSecretGroups returns the SecretGroups field if non-nil, zero value otherwise.

### GetSecretGroupsOk

`func (o *TerraformConfig) GetSecretGroupsOk() (*[]string, bool)`

GetSecretGroupsOk returns a tuple with the SecretGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretGroups

`func (o *TerraformConfig) SetSecretGroups(v []string)`

SetSecretGroups sets SecretGroups field to given value.

### HasSecretGroups

`func (o *TerraformConfig) HasSecretGroups() bool`

HasSecretGroups returns a boolean if a field has been set.

### GetTfVarsFilePath

`func (o *TerraformConfig) GetTfVarsFilePath() File`

GetTfVarsFilePath returns the TfVarsFilePath field if non-nil, zero value otherwise.

### GetTfVarsFilePathOk

`func (o *TerraformConfig) GetTfVarsFilePathOk() (*File, bool)`

GetTfVarsFilePathOk returns a tuple with the TfVarsFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfVarsFilePath

`func (o *TerraformConfig) SetTfVarsFilePath(v File)`

SetTfVarsFilePath sets TfVarsFilePath field to given value.

### HasTfVarsFilePath

`func (o *TerraformConfig) HasTfVarsFilePath() bool`

HasTfVarsFilePath returns a boolean if a field has been set.

### GetVersion

`func (o *TerraformConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TerraformConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TerraformConfig) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


