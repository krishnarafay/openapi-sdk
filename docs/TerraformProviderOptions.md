# TerraformProviderOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendConfigs** | Pointer to **[]string** | This can be either a path to an HCL file with key/value assignments (same format as terraform.tfvars) or a &#39;key&#x3D;value&#39; format. This is merged with what is in the configuration file | [optional] 
**Lock** | Pointer to **bool** | Don&#39;t hold a state lock during the operation. This is dangerous if others might concurrently run commands against the same workspace. | [optional] 
**LockTimeout** | Pointer to **string** | Duration to retry a state lock. | [optional] 
**PluginDirs** | Pointer to **[]string** | Directory containing plugin binaries. This overrides all default search paths for plugins, and prevents the automatic installation of plugins. This flag can be used multiple times | [optional] 
**Refresh** | Pointer to **bool** | Skip checking for external changes to remote objects while creating the plan. This can potentially make planning faster, but at the expense of possibly planning against a stale record of the remote system state. | [optional] 
**RefreshOnly** | Pointer to **bool** | Select the refresh only planning mode, which checks whether remote objects still match the outcome of the most recent Terraform apply but does not propose any actions to undo any changes made outside of Terraform. | [optional] 
**TargetResources** | Pointer to **[]string** | Limit the planning operation to only the given module, resource, or resource instance and all of its dependencies | [optional] 
**UseSystemStateStore** | Pointer to **bool** | Use system state store, by default it is false | [optional] 
**VarFiles** | Pointer to **[]string** | Load variable values from the given file, in addition to the default files terraform.tfvars and *.auto.tfvars. Use this option more than once to include more than one variables files | [optional] 
**Version** | Pointer to **string** | Terraform version | [optional] 

## Methods

### NewTerraformProviderOptions

`func NewTerraformProviderOptions() *TerraformProviderOptions`

NewTerraformProviderOptions instantiates a new TerraformProviderOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformProviderOptionsWithDefaults

`func NewTerraformProviderOptionsWithDefaults() *TerraformProviderOptions`

NewTerraformProviderOptionsWithDefaults instantiates a new TerraformProviderOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendConfigs

`func (o *TerraformProviderOptions) GetBackendConfigs() []string`

GetBackendConfigs returns the BackendConfigs field if non-nil, zero value otherwise.

### GetBackendConfigsOk

`func (o *TerraformProviderOptions) GetBackendConfigsOk() (*[]string, bool)`

GetBackendConfigsOk returns a tuple with the BackendConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendConfigs

`func (o *TerraformProviderOptions) SetBackendConfigs(v []string)`

SetBackendConfigs sets BackendConfigs field to given value.

### HasBackendConfigs

`func (o *TerraformProviderOptions) HasBackendConfigs() bool`

HasBackendConfigs returns a boolean if a field has been set.

### GetLock

`func (o *TerraformProviderOptions) GetLock() bool`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *TerraformProviderOptions) GetLockOk() (*bool, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *TerraformProviderOptions) SetLock(v bool)`

SetLock sets Lock field to given value.

### HasLock

`func (o *TerraformProviderOptions) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetLockTimeout

`func (o *TerraformProviderOptions) GetLockTimeout() string`

GetLockTimeout returns the LockTimeout field if non-nil, zero value otherwise.

### GetLockTimeoutOk

`func (o *TerraformProviderOptions) GetLockTimeoutOk() (*string, bool)`

GetLockTimeoutOk returns a tuple with the LockTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTimeout

`func (o *TerraformProviderOptions) SetLockTimeout(v string)`

SetLockTimeout sets LockTimeout field to given value.

### HasLockTimeout

`func (o *TerraformProviderOptions) HasLockTimeout() bool`

HasLockTimeout returns a boolean if a field has been set.

### GetPluginDirs

`func (o *TerraformProviderOptions) GetPluginDirs() []string`

GetPluginDirs returns the PluginDirs field if non-nil, zero value otherwise.

### GetPluginDirsOk

`func (o *TerraformProviderOptions) GetPluginDirsOk() (*[]string, bool)`

GetPluginDirsOk returns a tuple with the PluginDirs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginDirs

`func (o *TerraformProviderOptions) SetPluginDirs(v []string)`

SetPluginDirs sets PluginDirs field to given value.

### HasPluginDirs

`func (o *TerraformProviderOptions) HasPluginDirs() bool`

HasPluginDirs returns a boolean if a field has been set.

### GetRefresh

`func (o *TerraformProviderOptions) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *TerraformProviderOptions) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *TerraformProviderOptions) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *TerraformProviderOptions) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetRefreshOnly

`func (o *TerraformProviderOptions) GetRefreshOnly() bool`

GetRefreshOnly returns the RefreshOnly field if non-nil, zero value otherwise.

### GetRefreshOnlyOk

`func (o *TerraformProviderOptions) GetRefreshOnlyOk() (*bool, bool)`

GetRefreshOnlyOk returns a tuple with the RefreshOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshOnly

`func (o *TerraformProviderOptions) SetRefreshOnly(v bool)`

SetRefreshOnly sets RefreshOnly field to given value.

### HasRefreshOnly

`func (o *TerraformProviderOptions) HasRefreshOnly() bool`

HasRefreshOnly returns a boolean if a field has been set.

### GetTargetResources

`func (o *TerraformProviderOptions) GetTargetResources() []string`

GetTargetResources returns the TargetResources field if non-nil, zero value otherwise.

### GetTargetResourcesOk

`func (o *TerraformProviderOptions) GetTargetResourcesOk() (*[]string, bool)`

GetTargetResourcesOk returns a tuple with the TargetResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResources

`func (o *TerraformProviderOptions) SetTargetResources(v []string)`

SetTargetResources sets TargetResources field to given value.

### HasTargetResources

`func (o *TerraformProviderOptions) HasTargetResources() bool`

HasTargetResources returns a boolean if a field has been set.

### GetUseSystemStateStore

`func (o *TerraformProviderOptions) GetUseSystemStateStore() bool`

GetUseSystemStateStore returns the UseSystemStateStore field if non-nil, zero value otherwise.

### GetUseSystemStateStoreOk

`func (o *TerraformProviderOptions) GetUseSystemStateStoreOk() (*bool, bool)`

GetUseSystemStateStoreOk returns a tuple with the UseSystemStateStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSystemStateStore

`func (o *TerraformProviderOptions) SetUseSystemStateStore(v bool)`

SetUseSystemStateStore sets UseSystemStateStore field to given value.

### HasUseSystemStateStore

`func (o *TerraformProviderOptions) HasUseSystemStateStore() bool`

HasUseSystemStateStore returns a boolean if a field has been set.

### GetVarFiles

`func (o *TerraformProviderOptions) GetVarFiles() []string`

GetVarFiles returns the VarFiles field if non-nil, zero value otherwise.

### GetVarFilesOk

`func (o *TerraformProviderOptions) GetVarFilesOk() (*[]string, bool)`

GetVarFilesOk returns a tuple with the VarFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarFiles

`func (o *TerraformProviderOptions) SetVarFiles(v []string)`

SetVarFiles sets VarFiles field to given value.

### HasVarFiles

`func (o *TerraformProviderOptions) HasVarFiles() bool`

HasVarFiles returns a boolean if a field has been set.

### GetVersion

`func (o *TerraformProviderOptions) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TerraformProviderOptions) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TerraformProviderOptions) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TerraformProviderOptions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


