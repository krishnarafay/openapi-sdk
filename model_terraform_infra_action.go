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

// checks if the TerraformInfraAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerraformInfraAction{}

// TerraformInfraAction configuration for terraform infra action
type TerraformInfraAction struct {
	// terraform action
	Action string `json:"action"`
	BackendFilePath *File `json:"backendFilePath,omitempty"`
	// backend variables
	BackendVars []KeyValue `json:"backendVars,omitempty"`
	// destroy
	Destroy *bool `json:"destroy,omitempty"`
	// environment variables
	EnvVars []KeyValue `json:"envVars,omitempty"`
	// input variables for terrafrom
	InputVars []KeyValue `json:"inputVars,omitempty"`
	// refresh
	Refresh *bool `json:"refresh,omitempty"`
	// Pipeline secrets groups
	SecretGroups []string `json:"secretGroups,omitempty"`
	// terraform targets
	Targets []TerraformTarget `json:"targets,omitempty"`
	TfVarsFilePath *File `json:"tfVarsFilePath,omitempty"`
	// terraform version
	Version string `json:"version"`
}

// NewTerraformInfraAction instantiates a new TerraformInfraAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerraformInfraAction(action string, version string) *TerraformInfraAction {
	this := TerraformInfraAction{}
	this.Action = action
	this.Version = version
	return &this
}

// NewTerraformInfraActionWithDefaults instantiates a new TerraformInfraAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerraformInfraActionWithDefaults() *TerraformInfraAction {
	this := TerraformInfraAction{}
	return &this
}

// GetAction returns the Action field value
func (o *TerraformInfraAction) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *TerraformInfraAction) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *TerraformInfraAction) SetAction(v string) {
	o.Action = v
}

// GetBackendFilePath returns the BackendFilePath field value if set, zero value otherwise.
func (o *TerraformInfraAction) GetBackendFilePath() File {
	if o == nil || isNil(o.BackendFilePath) {
		var ret File
		return ret
	}
	return *o.BackendFilePath
}

// GetBackendFilePathOk returns a tuple with the BackendFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformInfraAction) GetBackendFilePathOk() (*File, bool) {
	if o == nil || isNil(o.BackendFilePath) {
		return nil, false
	}
	return o.BackendFilePath, true
}

// HasBackendFilePath returns a boolean if a field has been set.
func (o *TerraformInfraAction) HasBackendFilePath() bool {
	if o != nil && !isNil(o.BackendFilePath) {
		return true
	}

	return false
}

// SetBackendFilePath gets a reference to the given File and assigns it to the BackendFilePath field.
func (o *TerraformInfraAction) SetBackendFilePath(v File) {
	o.BackendFilePath = &v
}

// GetBackendVars returns the BackendVars field value if set, zero value otherwise.
func (o *TerraformInfraAction) GetBackendVars() []KeyValue {
	if o == nil || isNil(o.BackendVars) {
		var ret []KeyValue
		return ret
	}
	return o.BackendVars
}

// GetBackendVarsOk returns a tuple with the BackendVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformInfraAction) GetBackendVarsOk() ([]KeyValue, bool) {
	if o == nil || isNil(o.BackendVars) {
		return nil, false
	}
	return o.BackendVars, true
}

// HasBackendVars returns a boolean if a field has been set.
func (o *TerraformInfraAction) HasBackendVars() bool {
	if o != nil && !isNil(o.BackendVars) {
		return true
	}

	return false
}

// SetBackendVars gets a reference to the given []KeyValue and assigns it to the BackendVars field.
func (o *TerraformInfraAction) SetBackendVars(v []KeyValue) {
	o.BackendVars = v
}

// GetDestroy returns the Destroy field value if set, zero value otherwise.
func (o *TerraformInfraAction) GetDestroy() bool {
	if o == nil || isNil(o.Destroy) {
		var ret bool
		return ret
	}
	return *o.Destroy
}

// GetDestroyOk returns a tuple with the Destroy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformInfraAction) GetDestroyOk() (*bool, bool) {
	if o == nil || isNil(o.Destroy) {
		return nil, false
	}
	return o.Destroy, true
}

// HasDestroy returns a boolean if a field has been set.
func (o *TerraformInfraAction) HasDestroy() bool {
	if o != nil && !isNil(o.Destroy) {
		return true
	}

	return false
}

// SetDestroy gets a reference to the given bool and assigns it to the Destroy field.
func (o *TerraformInfraAction) SetDestroy(v bool) {
	o.Destroy = &v
}

// GetEnvVars returns the EnvVars field value if set, zero value otherwise.
func (o *TerraformInfraAction) GetEnvVars() []KeyValue {
	if o == nil || isNil(o.EnvVars) {
		var ret []KeyValue
		return ret
	}
	return o.EnvVars
}

// GetEnvVarsOk returns a tuple with the EnvVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformInfraAction) GetEnvVarsOk() ([]KeyValue, bool) {
	if o == nil || isNil(o.EnvVars) {
		return nil, false
	}
	return o.EnvVars, true
}

// HasEnvVars returns a boolean if a field has been set.
func (o *TerraformInfraAction) HasEnvVars() bool {
	if o != nil && !isNil(o.EnvVars) {
		return true
	}

	return false
}

// SetEnvVars gets a reference to the given []KeyValue and assigns it to the EnvVars field.
func (o *TerraformInfraAction) SetEnvVars(v []KeyValue) {
	o.EnvVars = v
}

// GetInputVars returns the InputVars field value if set, zero value otherwise.
func (o *TerraformInfraAction) GetInputVars() []KeyValue {
	if o == nil || isNil(o.InputVars) {
		var ret []KeyValue
		return ret
	}
	return o.InputVars
}

// GetInputVarsOk returns a tuple with the InputVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformInfraAction) GetInputVarsOk() ([]KeyValue, bool) {
	if o == nil || isNil(o.InputVars) {
		return nil, false
	}
	return o.InputVars, true
}

// HasInputVars returns a boolean if a field has been set.
func (o *TerraformInfraAction) HasInputVars() bool {
	if o != nil && !isNil(o.InputVars) {
		return true
	}

	return false
}

// SetInputVars gets a reference to the given []KeyValue and assigns it to the InputVars field.
func (o *TerraformInfraAction) SetInputVars(v []KeyValue) {
	o.InputVars = v
}

// GetRefresh returns the Refresh field value if set, zero value otherwise.
func (o *TerraformInfraAction) GetRefresh() bool {
	if o == nil || isNil(o.Refresh) {
		var ret bool
		return ret
	}
	return *o.Refresh
}

// GetRefreshOk returns a tuple with the Refresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformInfraAction) GetRefreshOk() (*bool, bool) {
	if o == nil || isNil(o.Refresh) {
		return nil, false
	}
	return o.Refresh, true
}

// HasRefresh returns a boolean if a field has been set.
func (o *TerraformInfraAction) HasRefresh() bool {
	if o != nil && !isNil(o.Refresh) {
		return true
	}

	return false
}

// SetRefresh gets a reference to the given bool and assigns it to the Refresh field.
func (o *TerraformInfraAction) SetRefresh(v bool) {
	o.Refresh = &v
}

// GetSecretGroups returns the SecretGroups field value if set, zero value otherwise.
func (o *TerraformInfraAction) GetSecretGroups() []string {
	if o == nil || isNil(o.SecretGroups) {
		var ret []string
		return ret
	}
	return o.SecretGroups
}

// GetSecretGroupsOk returns a tuple with the SecretGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformInfraAction) GetSecretGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.SecretGroups) {
		return nil, false
	}
	return o.SecretGroups, true
}

// HasSecretGroups returns a boolean if a field has been set.
func (o *TerraformInfraAction) HasSecretGroups() bool {
	if o != nil && !isNil(o.SecretGroups) {
		return true
	}

	return false
}

// SetSecretGroups gets a reference to the given []string and assigns it to the SecretGroups field.
func (o *TerraformInfraAction) SetSecretGroups(v []string) {
	o.SecretGroups = v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *TerraformInfraAction) GetTargets() []TerraformTarget {
	if o == nil || isNil(o.Targets) {
		var ret []TerraformTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformInfraAction) GetTargetsOk() ([]TerraformTarget, bool) {
	if o == nil || isNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *TerraformInfraAction) HasTargets() bool {
	if o != nil && !isNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []TerraformTarget and assigns it to the Targets field.
func (o *TerraformInfraAction) SetTargets(v []TerraformTarget) {
	o.Targets = v
}

// GetTfVarsFilePath returns the TfVarsFilePath field value if set, zero value otherwise.
func (o *TerraformInfraAction) GetTfVarsFilePath() File {
	if o == nil || isNil(o.TfVarsFilePath) {
		var ret File
		return ret
	}
	return *o.TfVarsFilePath
}

// GetTfVarsFilePathOk returns a tuple with the TfVarsFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformInfraAction) GetTfVarsFilePathOk() (*File, bool) {
	if o == nil || isNil(o.TfVarsFilePath) {
		return nil, false
	}
	return o.TfVarsFilePath, true
}

// HasTfVarsFilePath returns a boolean if a field has been set.
func (o *TerraformInfraAction) HasTfVarsFilePath() bool {
	if o != nil && !isNil(o.TfVarsFilePath) {
		return true
	}

	return false
}

// SetTfVarsFilePath gets a reference to the given File and assigns it to the TfVarsFilePath field.
func (o *TerraformInfraAction) SetTfVarsFilePath(v File) {
	o.TfVarsFilePath = &v
}

// GetVersion returns the Version field value
func (o *TerraformInfraAction) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *TerraformInfraAction) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *TerraformInfraAction) SetVersion(v string) {
	o.Version = v
}

func (o TerraformInfraAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerraformInfraAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !isNil(o.BackendFilePath) {
		toSerialize["backendFilePath"] = o.BackendFilePath
	}
	if !isNil(o.BackendVars) {
		toSerialize["backendVars"] = o.BackendVars
	}
	if !isNil(o.Destroy) {
		toSerialize["destroy"] = o.Destroy
	}
	if !isNil(o.EnvVars) {
		toSerialize["envVars"] = o.EnvVars
	}
	if !isNil(o.InputVars) {
		toSerialize["inputVars"] = o.InputVars
	}
	if !isNil(o.Refresh) {
		toSerialize["refresh"] = o.Refresh
	}
	if !isNil(o.SecretGroups) {
		toSerialize["secretGroups"] = o.SecretGroups
	}
	if !isNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	if !isNil(o.TfVarsFilePath) {
		toSerialize["tfVarsFilePath"] = o.TfVarsFilePath
	}
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableTerraformInfraAction struct {
	value *TerraformInfraAction
	isSet bool
}

func (v NullableTerraformInfraAction) Get() *TerraformInfraAction {
	return v.value
}

func (v *NullableTerraformInfraAction) Set(val *TerraformInfraAction) {
	v.value = val
	v.isSet = true
}

func (v NullableTerraformInfraAction) IsSet() bool {
	return v.isSet
}

func (v *NullableTerraformInfraAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerraformInfraAction(val *TerraformInfraAction) *NullableTerraformInfraAction {
	return &NullableTerraformInfraAction{value: val, isSet: true}
}

func (v NullableTerraformInfraAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerraformInfraAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


