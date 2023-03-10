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

// checks if the OPAProfileSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OPAProfileSpec{}

// OPAProfileSpec profile specification
type OPAProfileSpec struct {
	// list of namespaces excluded from gatekeeper processing
	ExcludedNamespaces []ExcludedNamespaces `json:"excludedNamespaces,omitempty"`
	InstallationParams *InstallationParams `json:"installationParams,omitempty"`
	// version of the OPA
	OpaVersion *string `json:"opaVersion,omitempty"`
	Sharing *SharingSpec `json:"sharing,omitempty"`
	// list of k8s objects to be synced
	SyncObjects []SyncObject `json:"syncObjects,omitempty"`
	// version of the profile
	Version *string `json:"version,omitempty"`
}

// NewOPAProfileSpec instantiates a new OPAProfileSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOPAProfileSpec() *OPAProfileSpec {
	this := OPAProfileSpec{}
	return &this
}

// NewOPAProfileSpecWithDefaults instantiates a new OPAProfileSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOPAProfileSpecWithDefaults() *OPAProfileSpec {
	this := OPAProfileSpec{}
	return &this
}

// GetExcludedNamespaces returns the ExcludedNamespaces field value if set, zero value otherwise.
func (o *OPAProfileSpec) GetExcludedNamespaces() []ExcludedNamespaces {
	if o == nil || isNil(o.ExcludedNamespaces) {
		var ret []ExcludedNamespaces
		return ret
	}
	return o.ExcludedNamespaces
}

// GetExcludedNamespacesOk returns a tuple with the ExcludedNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OPAProfileSpec) GetExcludedNamespacesOk() ([]ExcludedNamespaces, bool) {
	if o == nil || isNil(o.ExcludedNamespaces) {
		return nil, false
	}
	return o.ExcludedNamespaces, true
}

// HasExcludedNamespaces returns a boolean if a field has been set.
func (o *OPAProfileSpec) HasExcludedNamespaces() bool {
	if o != nil && !isNil(o.ExcludedNamespaces) {
		return true
	}

	return false
}

// SetExcludedNamespaces gets a reference to the given []ExcludedNamespaces and assigns it to the ExcludedNamespaces field.
func (o *OPAProfileSpec) SetExcludedNamespaces(v []ExcludedNamespaces) {
	o.ExcludedNamespaces = v
}

// GetInstallationParams returns the InstallationParams field value if set, zero value otherwise.
func (o *OPAProfileSpec) GetInstallationParams() InstallationParams {
	if o == nil || isNil(o.InstallationParams) {
		var ret InstallationParams
		return ret
	}
	return *o.InstallationParams
}

// GetInstallationParamsOk returns a tuple with the InstallationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OPAProfileSpec) GetInstallationParamsOk() (*InstallationParams, bool) {
	if o == nil || isNil(o.InstallationParams) {
		return nil, false
	}
	return o.InstallationParams, true
}

// HasInstallationParams returns a boolean if a field has been set.
func (o *OPAProfileSpec) HasInstallationParams() bool {
	if o != nil && !isNil(o.InstallationParams) {
		return true
	}

	return false
}

// SetInstallationParams gets a reference to the given InstallationParams and assigns it to the InstallationParams field.
func (o *OPAProfileSpec) SetInstallationParams(v InstallationParams) {
	o.InstallationParams = &v
}

// GetOpaVersion returns the OpaVersion field value if set, zero value otherwise.
func (o *OPAProfileSpec) GetOpaVersion() string {
	if o == nil || isNil(o.OpaVersion) {
		var ret string
		return ret
	}
	return *o.OpaVersion
}

// GetOpaVersionOk returns a tuple with the OpaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OPAProfileSpec) GetOpaVersionOk() (*string, bool) {
	if o == nil || isNil(o.OpaVersion) {
		return nil, false
	}
	return o.OpaVersion, true
}

// HasOpaVersion returns a boolean if a field has been set.
func (o *OPAProfileSpec) HasOpaVersion() bool {
	if o != nil && !isNil(o.OpaVersion) {
		return true
	}

	return false
}

// SetOpaVersion gets a reference to the given string and assigns it to the OpaVersion field.
func (o *OPAProfileSpec) SetOpaVersion(v string) {
	o.OpaVersion = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *OPAProfileSpec) GetSharing() SharingSpec {
	if o == nil || isNil(o.Sharing) {
		var ret SharingSpec
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OPAProfileSpec) GetSharingOk() (*SharingSpec, bool) {
	if o == nil || isNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *OPAProfileSpec) HasSharing() bool {
	if o != nil && !isNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given SharingSpec and assigns it to the Sharing field.
func (o *OPAProfileSpec) SetSharing(v SharingSpec) {
	o.Sharing = &v
}

// GetSyncObjects returns the SyncObjects field value if set, zero value otherwise.
func (o *OPAProfileSpec) GetSyncObjects() []SyncObject {
	if o == nil || isNil(o.SyncObjects) {
		var ret []SyncObject
		return ret
	}
	return o.SyncObjects
}

// GetSyncObjectsOk returns a tuple with the SyncObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OPAProfileSpec) GetSyncObjectsOk() ([]SyncObject, bool) {
	if o == nil || isNil(o.SyncObjects) {
		return nil, false
	}
	return o.SyncObjects, true
}

// HasSyncObjects returns a boolean if a field has been set.
func (o *OPAProfileSpec) HasSyncObjects() bool {
	if o != nil && !isNil(o.SyncObjects) {
		return true
	}

	return false
}

// SetSyncObjects gets a reference to the given []SyncObject and assigns it to the SyncObjects field.
func (o *OPAProfileSpec) SetSyncObjects(v []SyncObject) {
	o.SyncObjects = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OPAProfileSpec) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OPAProfileSpec) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OPAProfileSpec) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *OPAProfileSpec) SetVersion(v string) {
	o.Version = &v
}

func (o OPAProfileSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OPAProfileSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExcludedNamespaces) {
		toSerialize["excludedNamespaces"] = o.ExcludedNamespaces
	}
	if !isNil(o.InstallationParams) {
		toSerialize["installationParams"] = o.InstallationParams
	}
	if !isNil(o.OpaVersion) {
		toSerialize["opaVersion"] = o.OpaVersion
	}
	if !isNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	if !isNil(o.SyncObjects) {
		toSerialize["syncObjects"] = o.SyncObjects
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableOPAProfileSpec struct {
	value *OPAProfileSpec
	isSet bool
}

func (v NullableOPAProfileSpec) Get() *OPAProfileSpec {
	return v.value
}

func (v *NullableOPAProfileSpec) Set(val *OPAProfileSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableOPAProfileSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableOPAProfileSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOPAProfileSpec(val *OPAProfileSpec) *NullableOPAProfileSpec {
	return &NullableOPAProfileSpec{value: val, isSet: true}
}

func (v NullableOPAProfileSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOPAProfileSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


