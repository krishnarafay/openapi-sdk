# TerrafromConfig

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

### NewTerrafromConfig

`func NewTerrafromConfig(version string, ) *TerrafromConfig`

NewTerrafromConfig instantiates a new TerrafromConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerrafromConfigWithDefaults

`func NewTerrafromConfigWithDefaults() *TerrafromConfig`

NewTerrafromConfigWithDefaults instantiates a new TerrafromConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendFilePath

`func (o *TerrafromConfig) GetBackendFilePath() File`

GetBackendFilePath returns the BackendFilePath field if non-nil, zero value otherwise.

### GetBackendFilePathOk

`func (o *TerrafromConfig) GetBackendFilePathOk() (*File, bool)`

GetBackendFilePathOk returns a tuple with the BackendFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendFilePath

`func (o *TerrafromConfig) SetBackendFilePath(v File)`

SetBackendFilePath sets BackendFilePath field to given value.

### HasBackendFilePath

`func (o *TerrafromConfig) HasBackendFilePath() bool`

HasBackendFilePath returns a boolean if a field has been set.

### GetBackendVars

`func (o *TerrafromConfig) GetBackendVars() []KeyValue`

GetBackendVars returns the BackendVars field if non-nil, zero value otherwise.

### GetBackendVarsOk

`func (o *TerrafromConfig) GetBackendVarsOk() (*[]KeyValue, bool)`

GetBackendVarsOk returns a tuple with the BackendVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendVars

`func (o *TerrafromConfig) SetBackendVars(v []KeyValue)`

SetBackendVars sets BackendVars field to given value.

### HasBackendVars

`func (o *TerrafromConfig) HasBackendVars() bool`

HasBackendVars returns a boolean if a field has been set.

### GetEnvVars

`func (o *TerrafromConfig) GetEnvVars() []KeyValue`

GetEnvVars returns the EnvVars field if non-nil, zero value otherwise.

### GetEnvVarsOk

`func (o *TerrafromConfig) GetEnvVarsOk() (*[]KeyValue, bool)`

GetEnvVarsOk returns a tuple with the EnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVars

`func (o *TerrafromConfig) SetEnvVars(v []KeyValue)`

SetEnvVars sets EnvVars field to given value.

### HasEnvVars

`func (o *TerrafromConfig) HasEnvVars() bool`

HasEnvVars returns a boolean if a field has been set.

### GetInputVars

`func (o *TerrafromConfig) GetInputVars() []KeyValue`

GetInputVars returns the InputVars field if non-nil, zero value otherwise.

### GetInputVarsOk

`func (o *TerrafromConfig) GetInputVarsOk() (*[]KeyValue, bool)`

GetInputVarsOk returns a tuple with the InputVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputVars

`func (o *TerrafromConfig) SetInputVars(v []KeyValue)`

SetInputVars sets InputVars field to given value.

### HasInputVars

`func (o *TerrafromConfig) HasInputVars() bool`

HasInputVars returns a boolean if a field has been set.

### GetSecretGroups

`func (o *TerrafromConfig) GetSecretGroups() []string`

GetSecretGroups returns the SecretGroups field if non-nil, zero value otherwise.

### GetSecretGroupsOk

`func (o *TerrafromConfig) GetSecretGroupsOk() (*[]string, bool)`

GetSecretGroupsOk returns a tuple with the SecretGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretGroups

`func (o *TerrafromConfig) SetSecretGroups(v []string)`

SetSecretGroups sets SecretGroups field to given value.

### HasSecretGroups

`func (o *TerrafromConfig) HasSecretGroups() bool`

HasSecretGroups returns a boolean if a field has been set.

### GetTfVarsFilePath

`func (o *TerrafromConfig) GetTfVarsFilePath() File`

GetTfVarsFilePath returns the TfVarsFilePath field if non-nil, zero value otherwise.

### GetTfVarsFilePathOk

`func (o *TerrafromConfig) GetTfVarsFilePathOk() (*File, bool)`

GetTfVarsFilePathOk returns a tuple with the TfVarsFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfVarsFilePath

`func (o *TerrafromConfig) SetTfVarsFilePath(v File)`

SetTfVarsFilePath sets TfVarsFilePath field to given value.

### HasTfVarsFilePath

`func (o *TerrafromConfig) HasTfVarsFilePath() bool`

HasTfVarsFilePath returns a boolean if a field has been set.

### GetVersion

`func (o *TerrafromConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TerrafromConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TerrafromConfig) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


