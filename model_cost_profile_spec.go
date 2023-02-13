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

// checks if the CostProfileSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CostProfileSpec{}

// CostProfileSpec profile specification
type CostProfileSpec struct {
	InstallationParams *InstallationParams `json:"installationParams,omitempty"`
	// provider type
	ProviderType *string `json:"providerType,omitempty"`
	Sharing *SharingSpec `json:"sharing,omitempty"`
	// version of the profile
	Version *string `json:"version,omitempty"`
}

// NewCostProfileSpec instantiates a new CostProfileSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCostProfileSpec() *CostProfileSpec {
	this := CostProfileSpec{}
	return &this
}

// NewCostProfileSpecWithDefaults instantiates a new CostProfileSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCostProfileSpecWithDefaults() *CostProfileSpec {
	this := CostProfileSpec{}
	return &this
}

// GetInstallationParams returns the InstallationParams field value if set, zero value otherwise.
func (o *CostProfileSpec) GetInstallationParams() InstallationParams {
	if o == nil || isNil(o.InstallationParams) {
		var ret InstallationParams
		return ret
	}
	return *o.InstallationParams
}

// GetInstallationParamsOk returns a tuple with the InstallationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostProfileSpec) GetInstallationParamsOk() (*InstallationParams, bool) {
	if o == nil || isNil(o.InstallationParams) {
		return nil, false
	}
	return o.InstallationParams, true
}

// HasInstallationParams returns a boolean if a field has been set.
func (o *CostProfileSpec) HasInstallationParams() bool {
	if o != nil && !isNil(o.InstallationParams) {
		return true
	}

	return false
}

// SetInstallationParams gets a reference to the given InstallationParams and assigns it to the InstallationParams field.
func (o *CostProfileSpec) SetInstallationParams(v InstallationParams) {
	o.InstallationParams = &v
}

// GetProviderType returns the ProviderType field value if set, zero value otherwise.
func (o *CostProfileSpec) GetProviderType() string {
	if o == nil || isNil(o.ProviderType) {
		var ret string
		return ret
	}
	return *o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostProfileSpec) GetProviderTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProviderType) {
		return nil, false
	}
	return o.ProviderType, true
}

// HasProviderType returns a boolean if a field has been set.
func (o *CostProfileSpec) HasProviderType() bool {
	if o != nil && !isNil(o.ProviderType) {
		return true
	}

	return false
}

// SetProviderType gets a reference to the given string and assigns it to the ProviderType field.
func (o *CostProfileSpec) SetProviderType(v string) {
	o.ProviderType = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *CostProfileSpec) GetSharing() SharingSpec {
	if o == nil || isNil(o.Sharing) {
		var ret SharingSpec
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostProfileSpec) GetSharingOk() (*SharingSpec, bool) {
	if o == nil || isNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *CostProfileSpec) HasSharing() bool {
	if o != nil && !isNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given SharingSpec and assigns it to the Sharing field.
func (o *CostProfileSpec) SetSharing(v SharingSpec) {
	o.Sharing = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CostProfileSpec) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostProfileSpec) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CostProfileSpec) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CostProfileSpec) SetVersion(v string) {
	o.Version = &v
}

func (o CostProfileSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CostProfileSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.InstallationParams) {
		toSerialize["installationParams"] = o.InstallationParams
	}
	if !isNil(o.ProviderType) {
		toSerialize["providerType"] = o.ProviderType
	}
	if !isNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableCostProfileSpec struct {
	value *CostProfileSpec
	isSet bool
}

func (v NullableCostProfileSpec) Get() *CostProfileSpec {
	return v.value
}

func (v *NullableCostProfileSpec) Set(val *CostProfileSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCostProfileSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCostProfileSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCostProfileSpec(val *CostProfileSpec) *NullableCostProfileSpec {
	return &NullableCostProfileSpec{value: val, isSet: true}
}

func (v NullableCostProfileSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCostProfileSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

