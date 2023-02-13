# InfraProvisionerSpecConfig

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

### NewInfraProvisionerSpecConfig

`func NewInfraProvisionerSpecConfig(version string, ) *InfraProvisionerSpecConfig`

NewInfraProvisionerSpecConfig instantiates a new InfraProvisionerSpecConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraProvisionerSpecConfigWithDefaults

`func NewInfraProvisionerSpecConfigWithDefaults() *InfraProvisionerSpecConfig`

NewInfraProvisionerSpecConfigWithDefaults instantiates a new InfraProvisionerSpecConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendFilePath

`func (o *InfraProvisionerSpecConfig) GetBackendFilePath() File`

GetBackendFilePath returns the BackendFilePath field if non-nil, zero value otherwise.

### GetBackendFilePathOk

`func (o *InfraProvisionerSpecConfig) GetBackendFilePathOk() (*File, bool)`

GetBackendFilePathOk returns a tuple with the BackendFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendFilePath

`func (o *InfraProvisionerSpecConfig) SetBackendFilePath(v File)`

SetBackendFilePath sets BackendFilePath field to given value.

### HasBackendFilePath

`func (o *InfraProvisionerSpecConfig) HasBackendFilePath() bool`

HasBackendFilePath returns a boolean if a field has been set.

### GetBackendVars

`func (o *InfraProvisionerSpecConfig) GetBackendVars() []KeyValue`

GetBackendVars returns the BackendVars field if non-nil, zero value otherwise.

### GetBackendVarsOk

`func (o *InfraProvisionerSpecConfig) GetBackendVarsOk() (*[]KeyValue, bool)`

GetBackendVarsOk returns a tuple with the BackendVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendVars

`func (o *InfraProvisionerSpecConfig) SetBackendVars(v []KeyValue)`

SetBackendVars sets BackendVars field to given value.

### HasBackendVars

`func (o *InfraProvisionerSpecConfig) HasBackendVars() bool`

HasBackendVars returns a boolean if a field has been set.

### GetEnvVars

`func (o *InfraProvisionerSpecConfig) GetEnvVars() []KeyValue`

GetEnvVars returns the EnvVars field if non-nil, zero value otherwise.

### GetEnvVarsOk

`func (o *InfraProvisionerSpecConfig) GetEnvVarsOk() (*[]KeyValue, bool)`

GetEnvVarsOk returns a tuple with the EnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVars

`func (o *InfraProvisionerSpecConfig) SetEnvVars(v []KeyValue)`

SetEnvVars sets EnvVars field to given value.

### HasEnvVars

`func (o *InfraProvisionerSpecConfig) HasEnvVars() bool`

HasEnvVars returns a boolean if a field has been set.

### GetInputVars

`func (o *InfraProvisionerSpecConfig) GetInputVars() []KeyValue`

GetInputVars returns the InputVars field if non-nil, zero value otherwise.

### GetInputVarsOk

`func (o *InfraProvisionerSpecConfig) GetInputVarsOk() (*[]KeyValue, bool)`

GetInputVarsOk returns a tuple with the InputVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputVars

`func (o *InfraProvisionerSpecConfig) SetInputVars(v []KeyValue)`

SetInputVars sets InputVars field to given value.

### HasInputVars

`func (o *InfraProvisionerSpecConfig) HasInputVars() bool`

HasInputVars returns a boolean if a field has been set.

### GetSecretGroups

`func (o *InfraProvisionerSpecConfig) GetSecretGroups() []string`

GetSecretGroups returns the SecretGroups field if non-nil, zero value otherwise.

### GetSecretGroupsOk

`func (o *InfraProvisionerSpecConfig) GetSecretGroupsOk() (*[]string, bool)`

GetSecretGroupsOk returns a tuple with the SecretGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretGroups

`func (o *InfraProvisionerSpecConfig) SetSecretGroups(v []string)`

SetSecretGroups sets SecretGroups field to given value.

### HasSecretGroups

`func (o *InfraProvisionerSpecConfig) HasSecretGroups() bool`

HasSecretGroups returns a boolean if a field has been set.

### GetTfVarsFilePath

`func (o *InfraProvisionerSpecConfig) GetTfVarsFilePath() File`

GetTfVarsFilePath returns the TfVarsFilePath field if non-nil, zero value otherwise.

### GetTfVarsFilePathOk

`func (o *InfraProvisionerSpecConfig) GetTfVarsFilePathOk() (*File, bool)`

GetTfVarsFilePathOk returns a tuple with the TfVarsFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfVarsFilePath

`func (o *InfraProvisionerSpecConfig) SetTfVarsFilePath(v File)`

SetTfVarsFilePath sets TfVarsFilePath field to given value.

### HasTfVarsFilePath

`func (o *InfraProvisionerSpecConfig) HasTfVarsFilePath() bool`

HasTfVarsFilePath returns a boolean if a field has been set.

### GetVersion

`func (o *InfraProvisionerSpecConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InfraProvisionerSpecConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InfraProvisionerSpecConfig) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


