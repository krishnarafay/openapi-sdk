/*
Kubernetes Operations APIs

Kubernetes Operations APIs

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rafaysdkgov3

import (
	"encoding/json"
)

// checks if the TerraformProviderOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerraformProviderOptions{}

// TerraformProviderOptions struct for TerraformProviderOptions
type TerraformProviderOptions struct {
	// This can be either a path to an HCL file with key/value assignments (same format as terraform.tfvars) or a 'key=value' format. This is merged with what is in the configuration file
	BackendConfigs []string `json:"backendConfigs,omitempty"`
	// Don't hold a state lock during the operation. This is dangerous if others might concurrently run commands against the same workspace.
	Lock *bool `json:"lock,omitempty"`
	// Duration to retry a state lock.
	LockTimeout *string `json:"lockTimeout,omitempty"`
	// Directory containing plugin binaries. This overrides all default search paths for plugins, and prevents the automatic installation of plugins. This flag can be used multiple times
	PluginDirs []string `json:"pluginDirs,omitempty"`
	// Skip checking for external changes to remote objects while creating the plan. This can potentially make planning faster, but at the expense of possibly planning against a stale record of the remote system state.
	Refresh *bool `json:"refresh,omitempty"`
	// Select the refresh only planning mode, which checks whether remote objects still match the outcome of the most recent Terraform apply but does not propose any actions to undo any changes made outside of Terraform.
	RefreshOnly *bool `json:"refreshOnly,omitempty"`
	// Limit the planning operation to only the given module, resource, or resource instance and all of its dependencies
	TargetResources []string `json:"targetResources,omitempty"`
	// Use system state store, by default it is false
	UseSystemStateStore *bool `json:"useSystemStateStore,omitempty"`
	// Load variable values from the given file, in addition to the default files terraform.tfvars and *.auto.tfvars. Use this option more than once to include more than one variables files
	VarFiles []string `json:"varFiles,omitempty"`
	// Terraform version
	Version *string `json:"version,omitempty"`
}

// NewTerraformProviderOptions instantiates a new TerraformProviderOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerraformProviderOptions() *TerraformProviderOptions {
	this := TerraformProviderOptions{}
	return &this
}

// NewTerraformProviderOptionsWithDefaults instantiates a new TerraformProviderOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerraformProviderOptionsWithDefaults() *TerraformProviderOptions {
	this := TerraformProviderOptions{}
	return &this
}

// GetBackendConfigs returns the BackendConfigs field value if set, zero value otherwise.
func (o *TerraformProviderOptions) GetBackendConfigs() []string {
	if o == nil || isNil(o.BackendConfigs) {
		var ret []string
		return ret
	}
	return o.BackendConfigs
}

// GetBackendConfigsOk returns a tuple with the BackendConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformProviderOptions) GetBackendConfigsOk() ([]string, bool) {
	if o == nil || isNil(o.BackendConfigs) {
		return nil, false
	}
	return o.BackendConfigs, true
}

// HasBackendConfigs returns a boolean if a field has been set.
func (o *TerraformProviderOptions) HasBackendConfigs() bool {
	if o != nil && !isNil(o.BackendConfigs) {
		return true
	}

	return false
}

// SetBackendConfigs gets a reference to the given []string and assigns it to the BackendConfigs field.
func (o *TerraformProviderOptions) SetBackendConfigs(v []string) {
	o.BackendConfigs = v
}

// GetLock returns the Lock field value if set, zero value otherwise.
func (o *TerraformProviderOptions) GetLock() bool {
	if o == nil || isNil(o.Lock) {
		var ret bool
		return ret
	}
	return *o.Lock
}

// GetLockOk returns a tuple with the Lock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformProviderOptions) GetLockOk() (*bool, bool) {
	if o == nil || isNil(o.Lock) {
		return nil, false
	}
	return o.Lock, true
}

// HasLock returns a boolean if a field has been set.
func (o *TerraformProviderOptions) HasLock() bool {
	if o != nil && !isNil(o.Lock) {
		return true
	}

	return false
}

// SetLock gets a reference to the given bool and assigns it to the Lock field.
func (o *TerraformProviderOptions) SetLock(v bool) {
	o.Lock = &v
}

// GetLockTimeout returns the LockTimeout field value if set, zero value otherwise.
func (o *TerraformProviderOptions) GetLockTimeout() string {
	if o == nil || isNil(o.LockTimeout) {
		var ret string
		return ret
	}
	return *o.LockTimeout
}

// GetLockTimeoutOk returns a tuple with the LockTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformProviderOptions) GetLockTimeoutOk() (*string, bool) {
	if o == nil || isNil(o.LockTimeout) {
		return nil, false
	}
	return o.LockTimeout, true
}

// HasLockTimeout returns a boolean if a field has been set.
func (o *TerraformProviderOptions) HasLockTimeout() bool {
	if o != nil && !isNil(o.LockTimeout) {
		return true
	}

	return false
}

// SetLockTimeout gets a reference to the given string and assigns it to the LockTimeout field.
func (o *TerraformProviderOptions) SetLockTimeout(v string) {
	o.LockTimeout = &v
}

// GetPluginDirs returns the PluginDirs field value if set, zero value otherwise.
func (o *TerraformProviderOptions) GetPluginDirs() []string {
	if o == nil || isNil(o.PluginDirs) {
		var ret []string
		return ret
	}
	return o.PluginDirs
}

// GetPluginDirsOk returns a tuple with the PluginDirs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformProviderOptions) GetPluginDirsOk() ([]string, bool) {
	if o == nil || isNil(o.PluginDirs) {
		return nil, false
	}
	return o.PluginDirs, true
}

// HasPluginDirs returns a boolean if a field has been set.
func (o *TerraformProviderOptions) HasPluginDirs() bool {
	if o != nil && !isNil(o.PluginDirs) {
		return true
	}

	return false
}

// SetPluginDirs gets a reference to the given []string and assigns it to the PluginDirs field.
func (o *TerraformProviderOptions) SetPluginDirs(v []string) {
	o.PluginDirs = v
}

// GetRefresh returns the Refresh field value if set, zero value otherwise.
func (o *TerraformProviderOptions) GetRefresh() bool {
	if o == nil || isNil(o.Refresh) {
		var ret bool
		return ret
	}
	return *o.Refresh
}

// GetRefreshOk returns a tuple with the Refresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformProviderOptions) GetRefreshOk() (*bool, bool) {
	if o == nil || isNil(o.Refresh) {
		return nil, false
	}
	return o.Refresh, true
}

// HasRefresh returns a boolean if a field has been set.
func (o *TerraformProviderOptions) HasRefresh() bool {
	if o != nil && !isNil(o.Refresh) {
		return true
	}

	return false
}

// SetRefresh gets a reference to the given bool and assigns it to the Refresh field.
func (o *TerraformProviderOptions) SetRefresh(v bool) {
	o.Refresh = &v
}

// GetRefreshOnly returns the RefreshOnly field value if set, zero value otherwise.
func (o *TerraformProviderOptions) GetRefreshOnly() bool {
	if o == nil || isNil(o.RefreshOnly) {
		var ret bool
		return ret
	}
	return *o.RefreshOnly
}

// GetRefreshOnlyOk returns a tuple with the RefreshOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformProviderOptions) GetRefreshOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.RefreshOnly) {
		return nil, false
	}
	return o.RefreshOnly, true
}

// HasRefreshOnly returns a boolean if a field has been set.
func (o *TerraformProviderOptions) HasRefreshOnly() bool {
	if o != nil && !isNil(o.RefreshOnly) {
		return true
	}

	return false
}

// SetRefreshOnly gets a reference to the given bool and assigns it to the RefreshOnly field.
func (o *TerraformProviderOptions) SetRefreshOnly(v bool) {
	o.RefreshOnly = &v
}

// GetTargetResources returns the TargetResources field value if set, zero value otherwise.
func (o *TerraformProviderOptions) GetTargetResources() []string {
	if o == nil || isNil(o.TargetResources) {
		var ret []string
		return ret
	}
	return o.TargetResources
}

// GetTargetResourcesOk returns a tuple with the TargetResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformProviderOptions) GetTargetResourcesOk() ([]string, bool) {
	if o == nil || isNil(o.TargetResources) {
		return nil, false
	}
	return o.TargetResources, true
}

// HasTargetResources returns a boolean if a field has been set.
func (o *TerraformProviderOptions) HasTargetResources() bool {
	if o != nil && !isNil(o.TargetResources) {
		return true
	}

	return false
}

// SetTargetResources gets a reference to the given []string and assigns it to the TargetResources field.
func (o *TerraformProviderOptions) SetTargetResources(v []string) {
	o.TargetResources = v
}

// GetUseSystemStateStore returns the UseSystemStateStore field value if set, zero value otherwise.
func (o *TerraformProviderOptions) GetUseSystemStateStore() bool {
	if o == nil || isNil(o.UseSystemStateStore) {
		var ret bool
		return ret
	}
	return *o.UseSystemStateStore
}

// GetUseSystemStateStoreOk returns a tuple with the UseSystemStateStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformProviderOptions) GetUseSystemStateStoreOk() (*bool, bool) {
	if o == nil || isNil(o.UseSystemStateStore) {
		return nil, false
	}
	return o.UseSystemStateStore, true
}

// HasUseSystemStateStore returns a boolean if a field has been set.
func (o *TerraformProviderOptions) HasUseSystemStateStore() bool {
	if o != nil && !isNil(o.UseSystemStateStore) {
		return true
	}

	return false
}

// SetUseSystemStateStore gets a reference to the given bool and assigns it to the UseSystemStateStore field.
func (o *TerraformProviderOptions) SetUseSystemStateStore(v bool) {
	o.UseSystemStateStore = &v
}

// GetVarFiles returns the VarFiles field value if set, zero value otherwise.
func (o *TerraformProviderOptions) GetVarFiles() []string {
	if o == nil || isNil(o.VarFiles) {
		var ret []string
		return ret
	}
	return o.VarFiles
}

// GetVarFilesOk returns a tuple with the VarFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformProviderOptions) GetVarFilesOk() ([]string, bool) {
	if o == nil || isNil(o.VarFiles) {
		return nil, false
	}
	return o.VarFiles, true
}

// HasVarFiles returns a boolean if a field has been set.
func (o *TerraformProviderOptions) HasVarFiles() bool {
	if o != nil && !isNil(o.VarFiles) {
		return true
	}

	return false
}

// SetVarFiles gets a reference to the given []string and assigns it to the VarFiles field.
func (o *TerraformProviderOptions) SetVarFiles(v []string) {
	o.VarFiles = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TerraformProviderOptions) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformProviderOptions) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TerraformProviderOptions) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *TerraformProviderOptions) SetVersion(v string) {
	o.Version = &v
}

func (o TerraformProviderOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerraformProviderOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BackendConfigs) {
		toSerialize["backendConfigs"] = o.BackendConfigs
	}
	if !isNil(o.Lock) {
		toSerialize["lock"] = o.Lock
	}
	if !isNil(o.LockTimeout) {
		toSerialize["lockTimeout"] = o.LockTimeout
	}
	if !isNil(o.PluginDirs) {
		toSerialize["pluginDirs"] = o.PluginDirs
	}
	if !isNil(o.Refresh) {
		toSerialize["refresh"] = o.Refresh
	}
	if !isNil(o.RefreshOnly) {
		toSerialize["refreshOnly"] = o.RefreshOnly
	}
	if !isNil(o.TargetResources) {
		toSerialize["targetResources"] = o.TargetResources
	}
	if !isNil(o.UseSystemStateStore) {
		toSerialize["useSystemStateStore"] = o.UseSystemStateStore
	}
	if !isNil(o.VarFiles) {
		toSerialize["varFiles"] = o.VarFiles
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableTerraformProviderOptions struct {
	value *TerraformProviderOptions
	isSet bool
}

func (v NullableTerraformProviderOptions) Get() *TerraformProviderOptions {
	return v.value
}

func (v *NullableTerraformProviderOptions) Set(val *TerraformProviderOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTerraformProviderOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTerraformProviderOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerraformProviderOptions(val *TerraformProviderOptions) *NullableTerraformProviderOptions {
	return &NullableTerraformProviderOptions{value: val, isSet: true}
}

func (v NullableTerraformProviderOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerraformProviderOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


