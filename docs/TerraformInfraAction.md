# TerraformInfraAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | terraform action | 
**BackendFilePath** | Pointer to [**File**](File.md) |  | [optional] 
**BackendVars** | Pointer to [**[]KeyValue**](KeyValue.md) | backend variables | [optional] 
**Destroy** | Pointer to **bool** | destroy | [optional] 
**EnvVars** | Pointer to [**[]KeyValue**](KeyValue.md) | environment variables | [optional] 
**InputVars** | Pointer to [**[]KeyValue**](KeyValue.md) | input variables for terrafrom | [optional] 
**Refresh** | Pointer to **bool** | refresh | [optional] 
**SecretGroups** | Pointer to **[]string** | Pipeline secrets groups | [optional] 
**Targets** | Pointer to [**[]TerraformTarget**](TerraformTarget.md) | terraform targets | [optional] 
**TfVarsFilePath** | Pointer to [**File**](File.md) |  | [optional] 
**Version** | **string** | terraform version | 

## Methods

### NewTerraformInfraAction

`func NewTerraformInfraAction(action string, version string, ) *TerraformInfraAction`

NewTerraformInfraAction instantiates a new TerraformInfraAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformInfraActionWithDefaults

`func NewTerraformInfraActionWithDefaults() *TerraformInfraAction`

NewTerraformInfraActionWithDefaults instantiates a new TerraformInfraAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TerraformInfraAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TerraformInfraAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TerraformInfraAction) SetAction(v string)`

SetAction sets Action field to given value.


### GetBackendFilePath

`func (o *TerraformInfraAction) GetBackendFilePath() File`

GetBackendFilePath returns the BackendFilePath field if non-nil, zero value otherwise.

### GetBackendFilePathOk

`func (o *TerraformInfraAction) GetBackendFilePathOk() (*File, bool)`

GetBackendFilePathOk returns a tuple with the BackendFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendFilePath

`func (o *TerraformInfraAction) SetBackendFilePath(v File)`

SetBackendFilePath sets BackendFilePath field to given value.

### HasBackendFilePath

`func (o *TerraformInfraAction) HasBackendFilePath() bool`

HasBackendFilePath returns a boolean if a field has been set.

### GetBackendVars

`func (o *TerraformInfraAction) GetBackendVars() []KeyValue`

GetBackendVars returns the BackendVars field if non-nil, zero value otherwise.

### GetBackendVarsOk

`func (o *TerraformInfraAction) GetBackendVarsOk() (*[]KeyValue, bool)`

GetBackendVarsOk returns a tuple with the BackendVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendVars

`func (o *TerraformInfraAction) SetBackendVars(v []KeyValue)`

SetBackendVars sets BackendVars field to given value.

### HasBackendVars

`func (o *TerraformInfraAction) HasBackendVars() bool`

HasBackendVars returns a boolean if a field has been set.

### GetDestroy

`func (o *TerraformInfraAction) GetDestroy() bool`

GetDestroy returns the Destroy field if non-nil, zero value otherwise.

### GetDestroyOk

`func (o *TerraformInfraAction) GetDestroyOk() (*bool, bool)`

GetDestroyOk returns a tuple with the Destroy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroy

`func (o *TerraformInfraAction) SetDestroy(v bool)`

SetDestroy sets Destroy field to given value.

### HasDestroy

`func (o *TerraformInfraAction) HasDestroy() bool`

HasDestroy returns a boolean if a field has been set.

### GetEnvVars

`func (o *TerraformInfraAction) GetEnvVars() []KeyValue`

GetEnvVars returns the EnvVars field if non-nil, zero value otherwise.

### GetEnvVarsOk

`func (o *TerraformInfraAction) GetEnvVarsOk() (*[]KeyValue, bool)`

GetEnvVarsOk returns a tuple with the EnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVars

`func (o *TerraformInfraAction) SetEnvVars(v []KeyValue)`

SetEnvVars sets EnvVars field to given value.

### HasEnvVars

`func (o *TerraformInfraAction) HasEnvVars() bool`

HasEnvVars returns a boolean if a field has been set.

### GetInputVars

`func (o *TerraformInfraAction) GetInputVars() []KeyValue`

GetInputVars returns the InputVars field if non-nil, zero value otherwise.

### GetInputVarsOk

`func (o *TerraformInfraAction) GetInputVarsOk() (*[]KeyValue, bool)`

GetInputVarsOk returns a tuple with the InputVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputVars

`func (o *TerraformInfraAction) SetInputVars(v []KeyValue)`

SetInputVars sets InputVars field to given value.

### HasInputVars

`func (o *TerraformInfraAction) HasInputVars() bool`

HasInputVars returns a boolean if a field has been set.

### GetRefresh

`func (o *TerraformInfraAction) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *TerraformInfraAction) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *TerraformInfraAction) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *TerraformInfraAction) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetSecretGroups

`func (o *TerraformInfraAction) GetSecretGroups() []string`

GetSecretGroups returns the SecretGroups field if non-nil, zero value otherwise.

### GetSecretGroupsOk

`func (o *TerraformInfraAction) GetSecretGroupsOk() (*[]string, bool)`

GetSecretGroupsOk returns a tuple with the SecretGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretGroups

`func (o *TerraformInfraAction) SetSecretGroups(v []string)`

SetSecretGroups sets SecretGroups field to given value.

### HasSecretGroups

`func (o *TerraformInfraAction) HasSecretGroups() bool`

HasSecretGroups returns a boolean if a field has been set.

### GetTargets

`func (o *TerraformInfraAction) GetTargets() []TerraformTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *TerraformInfraAction) GetTargetsOk() (*[]TerraformTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *TerraformInfraAction) SetTargets(v []TerraformTarget)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *TerraformInfraAction) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetTfVarsFilePath

`func (o *TerraformInfraAction) GetTfVarsFilePath() File`

GetTfVarsFilePath returns the TfVarsFilePath field if non-nil, zero value otherwise.

### GetTfVarsFilePathOk

`func (o *TerraformInfraAction) GetTfVarsFilePathOk() (*File, bool)`

GetTfVarsFilePathOk returns a tuple with the TfVarsFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfVarsFilePath

`func (o *TerraformInfraAction) SetTfVarsFilePath(v File)`

SetTfVarsFilePath sets TfVarsFilePath field to given value.

### HasTfVarsFilePath

`func (o *TerraformInfraAction) HasTfVarsFilePath() bool`

HasTfVarsFilePath returns a boolean if a field has been set.

### GetVersion

`func (o *TerraformInfraAction) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TerraformInfraAction) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TerraformInfraAction) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


