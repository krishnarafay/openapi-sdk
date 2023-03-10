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

// checks if the BlueprintSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlueprintSpec{}

// BlueprintSpec blueprint specification
type BlueprintSpec struct {
	Base *BlueprintBase `json:"base,omitempty"`
	CostProfile *CostProfile `json:"costProfile,omitempty"`
	// custom addon configuration
	CustomAddons []BlueprintAddon `json:"customAddons,omitempty"`
	DefaultAddons *DefaultAddons `json:"defaultAddons,omitempty"`
	Drift *DriftSpec `json:"drift,omitempty"`
	NamespaceConfig *NsConfig `json:"namespaceConfig,omitempty"`
	NetworkPolicy *NetworkPolicy `json:"networkPolicy,omitempty"`
	OpaPolicy *OPAPolicy `json:"opaPolicy,omitempty"`
	Placement *BlueprintPlacement `json:"placement,omitempty"`
	// private  kubernetes API proxy networks
	PrivateKubeAPIProxies []KubeAPIProxyNetwork `json:"privateKubeAPIProxies,omitempty"`
	Psp *BlueprintPSP `json:"psp,omitempty"`
	ServiceMesh *ServiceMesh `json:"serviceMesh,omitempty"`
	Sharing *SharingSpec `json:"sharing,omitempty"`
	// Blueprint Type
	Type *string `json:"type,omitempty"`
	// version of the blueprint
	Version *string `json:"version,omitempty"`
}

// NewBlueprintSpec instantiates a new BlueprintSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintSpec() *BlueprintSpec {
	this := BlueprintSpec{}
	return &this
}

// NewBlueprintSpecWithDefaults instantiates a new BlueprintSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintSpecWithDefaults() *BlueprintSpec {
	this := BlueprintSpec{}
	return &this
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *BlueprintSpec) GetBase() BlueprintBase {
	if o == nil || isNil(o.Base) {
		var ret BlueprintBase
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintSpec) GetBaseOk() (*BlueprintBase, bool) {
	if o == nil || isNil(o.Base) {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *BlueprintSpec) HasBase() bool {
	if o != nil && !isNil(o.Base) {
		return true
	}

	return false
}

// SetBase gets a reference to the given BlueprintBase and assigns it to the Base field.
func (o *BlueprintSpec) SetBase(v BlueprintBase) {
	o.Base = &v
}

// GetCostProfile returns the CostProfile field value if set, zero value otherwise.
func (o *BlueprintSpec) GetCostProfile() CostProfile {
	if o == nil || isNil(o.CostProfile) {
		var ret CostProfile
		return ret
	}
	return *o.CostProfile
}

// GetCostProfileOk returns a tuple with the CostProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintSpec) GetCostProfileOk() (*CostProfile, bool) {
	if o == nil || isNil(o.CostProfile) {
		return nil, false
	}
	return o.CostProfile, true
}

// HasCostProfile returns a boolean if a field has been set.
func (o *BlueprintSpec) HasCostProfile() bool {
	if o != nil && !isNil(o.CostProfile) {
		return true
	}

	return false
}

// SetCostProfile gets a reference to the given CostProfile and assigns it to the CostProfile field.
func (o *BlueprintSpec) SetCostProfile(v CostProfile) {
	o.CostProfile = &v
}

// GetCustomAddons returns the CustomAddons field value if set, zero value otherwise.
func (o *BlueprintSpec) GetCustomAddons() []BlueprintAddon {
	if o == nil || isNil(o.CustomAddons) {
		var ret []BlueprintAddon
		return ret
	}
	return o.CustomAddons
}

// GetCustomAddonsOk returns a tuple with the CustomAddons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintSpec) GetCustomAddonsOk() ([]BlueprintAddon, bool) {
	if o == nil || isNil(o.CustomAddons) {
		return nil, false
	}
	return o.CustomAddons, true
}

// HasCustomAddons returns a boolean if a field has been set.
func (o *BlueprintSpec) HasCustomAddons() bool {
	if o != nil && !isNil(o.CustomAddons) {
		return true
	}

	return false
}

// SetCustomAddons gets a reference to the given []BlueprintAddon and assigns it to the CustomAddons field.
func (o *BlueprintSpec) SetCustomAddons(v []BlueprintAddon) {
	o.CustomAddons = v
}

// GetDefaultAddons returns the DefaultAddons field value if set, zero value otherwise.
func (o *BlueprintSpec) GetDefaultAddons() DefaultAddons {
	if o == nil || isNil(o.DefaultAddons) {
		var ret DefaultAddons
		return ret
	}
	return *o.DefaultAddons
}

// GetDefaultAddonsOk returns a tuple with the DefaultAddons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintSpec) GetDefaultAddonsOk() (*DefaultAddons, bool) {
	if o == nil || isNil(o.DefaultAddons) {
		return nil, false
	}
	return o.DefaultAddons, true
}

// HasDefaultAddons returns a boolean if a field has been set.
func (o *BlueprintSpec) HasDefaultAddons() bool {
	if o != nil && !isNil(o.DefaultAddons) {
		return true
	}

	return false
}

// SetDefaultAddons gets a reference to the given DefaultAddons and assigns it to the DefaultAddons field.
func (o *BlueprintSpec) SetDefaultAddons(v DefaultAddons) {
	o.DefaultAddons = &v
}

// GetDrift returns the Drift field value if set, zero value otherwise.
func (o *BlueprintSpec) GetDrift() DriftSpec {
	if o == nil || isNil(o.Drift) {
		var ret DriftSpec
		return ret
	}
	return *o.Drift
}

// GetDriftOk returns a tuple with the Drift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintSpec) GetDriftOk() (*DriftSpec, bool) {
	if o == nil || isNil(o.Drift) {
		return nil, false
	}
	return o.Drift, true
}

// HasDrift returns a boolean if a field has been set.
func (o *BlueprintSpec) HasDrift() bool {
	if o != nil && !isNil(o.Drift) {
		return true
	}

	return false
}

// SetDrift gets a reference to the given DriftSpec and assigns it to the Drift field.
func (o *BlueprintSpec) SetDrift(v DriftSpec) {
	o.Drift = &v
}

// GetNamespaceConfig returns the NamespaceConfig field value if set, zero value otherwise.
func (o *BlueprintSpec) GetNamespaceConfig() NsConfig {
	if o == nil || isNil(o.NamespaceConfig) {
		var ret NsConfig
		return ret
	}
	return *o.NamespaceConfig
}

// GetNamespaceConfigOk returns a tuple with the NamespaceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintSpec) GetNamespaceConfigOk() (*NsConfig, bool) {
	if o == nil || isNil(o.NamespaceConfig) {
		return nil, false
	}
	return o.NamespaceConfig, true
}

// HasNamespaceConfig returns a boolean if a field has been set.
func (o *BlueprintSpec) HasNamespaceConfig() bool {
	if o != nil && !isNil(o.NamespaceConfig) {
		return true
	}

	return false
}

// SetNamespaceConfig gets a reference to the given NsConfig and assigns it to the NamespaceConfig field.
func (o *BlueprintSpec) SetNamespaceConfig(v NsConfig) {
	o.NamespaceConfig = &v
}

// GetNetworkPolicy returns the NetworkPolicy field value if set, zero value otherwise.
func (o *BlueprintSpec) GetNetworkPolicy() NetworkPolicy {
	if o == nil || isNil(o.NetworkPolicy) {
		var ret NetworkPolicy
		return ret
	}
	return *o.NetworkPolicy
}

// GetNetworkPolicyOk returns a tuple with the NetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintSpec) GetNetworkPolicyOk() (*NetworkPolicy, bool) {
	if o == nil || isNil(o.NetworkPolicy) {
		return nil, false
	}
	return o.NetworkPolicy, true
}

// HasNetworkPolicy returns a boolean if a field has been set.
func (o *BlueprintSpec) HasNetworkPolicy() bool {
	if o != nil && !isNil(o.NetworkPolicy) {
		return true
	}

	return false
}

// SetNetworkPolicy gets a reference to the given NetworkPolicy and assigns it to the NetworkPolicy field.
func (o *BlueprintSpec) SetNetworkPolicy(v NetworkPolicy) {
	o.NetworkPolicy = &v
}

// GetOpaPolicy returns the OpaPolicy field value if set, zero value otherwise.
func (o *BlueprintSpec) GetOpaPolicy() OPAPolicy {
	if o == nil || isNil(o.OpaPolicy) {
		var ret OPAPolicy
		return ret
	}
	return *o.OpaPolicy
}

// GetOpaPolicyOk returns a tuple with the OpaPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintSpec) GetOpaPolicyOk() (*OPAPolicy, bool) {
	if o == nil || isNil(o.OpaPolicy) {
		return nil, false
	}
	return o.OpaPolicy, true
}

// HasOpaPolicy returns a boolean if a field has been set.
func (o *BlueprintSpec) HasOpaPolicy() bool {
	if o != nil && !isNil(o.OpaPolicy) {
		return true
	}

	return false
}

// SetOpaPolicy gets a reference to the given OPAPolicy and assigns it to the OpaPolicy field.
func (o *BlueprintSpec) SetOpaPolicy(v OPAPolicy) {
	o.OpaPolicy = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise.
func (o *BlueprintSpec) GetPlacement() BlueprintPlacement {
	if o == nil || isNil(o.Placement) {
		var ret BlueprintPlacement
		return ret
	}
	return *o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintSpec) GetPlacementOk() (*BlueprintPlacement, bool) {
	if o == nil || isNil(o.Placement) {
		return nil, false
	}
	return o.Placement, true
}

// HasPlacement returns a boolean if a field has been set.
func (o *BlueprintSpec) HasPlacement() bool {
	if o != nil && !isNil(o.Placement) {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given BlueprintPlacement and assigns it to the Placement field.
func (o *BlueprintSpec) SetPlacement(v BlueprintPlacement) {
	o.Placement = &v
}

// GetPrivateKubeAPIProxies returns the PrivateKubeAPIProxies field value if set, zero value otherwise.
func (o *BlueprintSpec) GetPrivateKubeAPIProxies() []KubeAPIProxyNetwork {
	if o == nil || isNil(o.PrivateKubeAPIProxies) {
		var ret []KubeAPIProxyNetwork
		return ret
	}
	return o.PrivateKubeAPIProxies
}

// GetPrivateKubeAPIProxiesOk returns a tuple with the PrivateKubeAPIProxies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintSpec) GetPrivateKubeAPIProxiesOk() ([]KubeAPIProxyNetwork, bool) {
	if o == nil || isNil(o.PrivateKubeAPIProxies) {
		return nil, false
	}
	return o.PrivateKubeAPIProxies, true
}

// HasPrivateKubeAPIProxies returns a boolean if a field has been set.
func (o *BlueprintSpec) HasPrivateKubeAPIProxies() bool {
	if o != nil && !isNil(o.PrivateKubeAPIProxies) {
		return true
	}

	return false
}

// SetPrivateKubeAPIProxies gets a reference to the given []KubeAPIProxyNetwork and assigns it to the PrivateKubeAPIProxies field.
func (o *BlueprintSpec) SetPrivateKubeAPIProxies(v []KubeAPIProxyNetwork) {
	o.PrivateKubeAPIProxies = v
}

// GetPsp returns the Psp field value if set, zero value otherwise.
func (o *BlueprintSpec) GetPsp() BlueprintPSP {
	if o == nil || isNil(o.Psp) {
		var ret BlueprintPSP
		return ret
	}
	return *o.Psp
}

// GetPspOk returns a tuple with the Psp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintSpec) GetPspOk() (*BlueprintPSP, bool) {
	if o == nil || isNil(o.Psp) {
		return nil, false
	}
	return o.Psp, true
}

// HasPsp returns a boolean if a field has been set.
func (o *BlueprintSpec) HasPsp() bool {
	if o != nil && !isNil(o.Psp) {
		return true
	}

	return false
}

// SetPsp gets a reference to the given BlueprintPSP and assigns it to the Psp field.
func (o *BlueprintSpec) SetPsp(v BlueprintPSP) {
	o.Psp = &v
}

// GetServiceMesh returns the ServiceMesh field value if set, zero value otherwise.
func (o *BlueprintSpec) GetServiceMesh() ServiceMesh {
	if o == nil || isNil(o.ServiceMesh) {
		var ret ServiceMesh
		return ret
	}
	return *o.ServiceMesh
}

// GetServiceMeshOk returns a tuple with the ServiceMesh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintSpec) GetServiceMeshOk() (*ServiceMesh, bool) {
	if o == nil || isNil(o.ServiceMesh) {
		return nil, false
	}
	return o.ServiceMesh, true
}

// HasServiceMesh returns a boolean if a field has been set.
func (o *BlueprintSpec) HasServiceMesh() bool {
	if o != nil && !isNil(o.ServiceMesh) {
		return true
	}

	return false
}

// SetServiceMesh gets a reference to the given ServiceMesh and assigns it to the ServiceMesh field.
func (o *BlueprintSpec) SetServiceMesh(v ServiceMesh) {
	o.ServiceMesh = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *BlueprintSpec) GetSharing() SharingSpec {
	if o == nil || isNil(o.Sharing) {
		var ret SharingSpec
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintSpec) GetSharingOk() (*SharingSpec, bool) {
	if o == nil || isNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *BlueprintSpec) HasSharing() bool {
	if o != nil && !isNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given SharingSpec and assigns it to the Sharing field.
func (o *BlueprintSpec) SetSharing(v SharingSpec) {
	o.Sharing = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BlueprintSpec) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintSpec) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BlueprintSpec) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BlueprintSpec) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BlueprintSpec) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintSpec) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BlueprintSpec) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *BlueprintSpec) SetVersion(v string) {
	o.Version = &v
}

func (o BlueprintSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlueprintSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Base) {
		toSerialize["base"] = o.Base
	}
	if !isNil(o.CostProfile) {
		toSerialize["costProfile"] = o.CostProfile
	}
	if !isNil(o.CustomAddons) {
		toSerialize["customAddons"] = o.CustomAddons
	}
	if !isNil(o.DefaultAddons) {
		toSerialize["defaultAddons"] = o.DefaultAddons
	}
	if !isNil(o.Drift) {
		toSerialize["drift"] = o.Drift
	}
	if !isNil(o.NamespaceConfig) {
		toSerialize["namespaceConfig"] = o.NamespaceConfig
	}
	if !isNil(o.NetworkPolicy) {
		toSerialize["networkPolicy"] = o.NetworkPolicy
	}
	if !isNil(o.OpaPolicy) {
		toSerialize["opaPolicy"] = o.OpaPolicy
	}
	if !isNil(o.Placement) {
		toSerialize["placement"] = o.Placement
	}
	if !isNil(o.PrivateKubeAPIProxies) {
		toSerialize["privateKubeAPIProxies"] = o.PrivateKubeAPIProxies
	}
	if !isNil(o.Psp) {
		toSerialize["psp"] = o.Psp
	}
	if !isNil(o.ServiceMesh) {
		toSerialize["serviceMesh"] = o.ServiceMesh
	}
	if !isNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableBlueprintSpec struct {
	value *BlueprintSpec
	isSet bool
}

func (v NullableBlueprintSpec) Get() *BlueprintSpec {
	return v.value
}

func (v *NullableBlueprintSpec) Set(val *BlueprintSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintSpec(val *BlueprintSpec) *NullableBlueprintSpec {
	return &NullableBlueprintSpec{value: val, isSet: true}
}

func (v NullableBlueprintSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


