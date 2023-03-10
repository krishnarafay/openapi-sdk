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

// checks if the Kubeletconfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Kubeletconfig{}

// Kubeletconfig struct for Kubeletconfig
type Kubeletconfig struct {
	// Allowed list of unsafe sysctls or unsafe sysctl patterns (ending in *).
	AllowedUnsafeSysctls []string `json:"allowedUnsafeSysctls,omitempty"`
	// The maximum number of container log files that can be present for a container. The number must be ≥ 2.
	ContainerLogMaxFiles *int32 `json:"containerLogMaxFiles,omitempty"`
	// The maximum size (e.g. 10Mi) of container log file before it is rotated.
	ContainerLogMaxSizeMB *int32 `json:"containerLogMaxSizeMB,omitempty"`
	// The default is true.
	CpuCfsQuota *bool `json:"cpuCfsQuota,omitempty"`
	// The default is 100ms. Valid values are a sequence of decimal numbers with an optional fraction and a unit suffix. For example: 300ms, 2h45m. Supported units are ns, us, ms, s, m, and h.
	CpuCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty"`
	// The default is none. See Kubernetes CPU management policies for more information. Allowed values are none and static.
	CpuManagerPolicy *string `json:"cpuManagerPolicy,omitempty"`
	// If set to true it will make the Kubelet fail to start if swap is enabled on the node.
	FailSwapOn *bool `json:"failSwapOn,omitempty"`
	// To disable image garbage collection, set to 100. The default is 85%
	ImageGcHighThreshold *int32 `json:"imageGcHighThreshold,omitempty"`
	// This cannot be set higher than imageGcHighThreshold. The default is 80%
	ImageGcLowThreshold *int32 `json:"imageGcLowThreshold,omitempty"`
	// The maximum number of processes per pod.
	PodMaxPids *int32 `json:"podMaxPids,omitempty"`
	// For more information see Kubernetes Topology Manager. The default is none. Allowed values are none, best-effort, restricted, and single-numa-node.
	TopologyManagerPolicy *string `json:"topologyManagerPolicy,omitempty"`
}

// NewKubeletconfig instantiates a new Kubeletconfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubeletconfig() *Kubeletconfig {
	this := Kubeletconfig{}
	return &this
}

// NewKubeletconfigWithDefaults instantiates a new Kubeletconfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubeletconfigWithDefaults() *Kubeletconfig {
	this := Kubeletconfig{}
	return &this
}

// GetAllowedUnsafeSysctls returns the AllowedUnsafeSysctls field value if set, zero value otherwise.
func (o *Kubeletconfig) GetAllowedUnsafeSysctls() []string {
	if o == nil || isNil(o.AllowedUnsafeSysctls) {
		var ret []string
		return ret
	}
	return o.AllowedUnsafeSysctls
}

// GetAllowedUnsafeSysctlsOk returns a tuple with the AllowedUnsafeSysctls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kubeletconfig) GetAllowedUnsafeSysctlsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedUnsafeSysctls) {
		return nil, false
	}
	return o.AllowedUnsafeSysctls, true
}

// HasAllowedUnsafeSysctls returns a boolean if a field has been set.
func (o *Kubeletconfig) HasAllowedUnsafeSysctls() bool {
	if o != nil && !isNil(o.AllowedUnsafeSysctls) {
		return true
	}

	return false
}

// SetAllowedUnsafeSysctls gets a reference to the given []string and assigns it to the AllowedUnsafeSysctls field.
func (o *Kubeletconfig) SetAllowedUnsafeSysctls(v []string) {
	o.AllowedUnsafeSysctls = v
}

// GetContainerLogMaxFiles returns the ContainerLogMaxFiles field value if set, zero value otherwise.
func (o *Kubeletconfig) GetContainerLogMaxFiles() int32 {
	if o == nil || isNil(o.ContainerLogMaxFiles) {
		var ret int32
		return ret
	}
	return *o.ContainerLogMaxFiles
}

// GetContainerLogMaxFilesOk returns a tuple with the ContainerLogMaxFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kubeletconfig) GetContainerLogMaxFilesOk() (*int32, bool) {
	if o == nil || isNil(o.ContainerLogMaxFiles) {
		return nil, false
	}
	return o.ContainerLogMaxFiles, true
}

// HasContainerLogMaxFiles returns a boolean if a field has been set.
func (o *Kubeletconfig) HasContainerLogMaxFiles() bool {
	if o != nil && !isNil(o.ContainerLogMaxFiles) {
		return true
	}

	return false
}

// SetContainerLogMaxFiles gets a reference to the given int32 and assigns it to the ContainerLogMaxFiles field.
func (o *Kubeletconfig) SetContainerLogMaxFiles(v int32) {
	o.ContainerLogMaxFiles = &v
}

// GetContainerLogMaxSizeMB returns the ContainerLogMaxSizeMB field value if set, zero value otherwise.
func (o *Kubeletconfig) GetContainerLogMaxSizeMB() int32 {
	if o == nil || isNil(o.ContainerLogMaxSizeMB) {
		var ret int32
		return ret
	}
	return *o.ContainerLogMaxSizeMB
}

// GetContainerLogMaxSizeMBOk returns a tuple with the ContainerLogMaxSizeMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kubeletconfig) GetContainerLogMaxSizeMBOk() (*int32, bool) {
	if o == nil || isNil(o.ContainerLogMaxSizeMB) {
		return nil, false
	}
	return o.ContainerLogMaxSizeMB, true
}

// HasContainerLogMaxSizeMB returns a boolean if a field has been set.
func (o *Kubeletconfig) HasContainerLogMaxSizeMB() bool {
	if o != nil && !isNil(o.ContainerLogMaxSizeMB) {
		return true
	}

	return false
}

// SetContainerLogMaxSizeMB gets a reference to the given int32 and assigns it to the ContainerLogMaxSizeMB field.
func (o *Kubeletconfig) SetContainerLogMaxSizeMB(v int32) {
	o.ContainerLogMaxSizeMB = &v
}

// GetCpuCfsQuota returns the CpuCfsQuota field value if set, zero value otherwise.
func (o *Kubeletconfig) GetCpuCfsQuota() bool {
	if o == nil || isNil(o.CpuCfsQuota) {
		var ret bool
		return ret
	}
	return *o.CpuCfsQuota
}

// GetCpuCfsQuotaOk returns a tuple with the CpuCfsQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kubeletconfig) GetCpuCfsQuotaOk() (*bool, bool) {
	if o == nil || isNil(o.CpuCfsQuota) {
		return nil, false
	}
	return o.CpuCfsQuota, true
}

// HasCpuCfsQuota returns a boolean if a field has been set.
func (o *Kubeletconfig) HasCpuCfsQuota() bool {
	if o != nil && !isNil(o.CpuCfsQuota) {
		return true
	}

	return false
}

// SetCpuCfsQuota gets a reference to the given bool and assigns it to the CpuCfsQuota field.
func (o *Kubeletconfig) SetCpuCfsQuota(v bool) {
	o.CpuCfsQuota = &v
}

// GetCpuCfsQuotaPeriod returns the CpuCfsQuotaPeriod field value if set, zero value otherwise.
func (o *Kubeletconfig) GetCpuCfsQuotaPeriod() string {
	if o == nil || isNil(o.CpuCfsQuotaPeriod) {
		var ret string
		return ret
	}
	return *o.CpuCfsQuotaPeriod
}

// GetCpuCfsQuotaPeriodOk returns a tuple with the CpuCfsQuotaPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kubeletconfig) GetCpuCfsQuotaPeriodOk() (*string, bool) {
	if o == nil || isNil(o.CpuCfsQuotaPeriod) {
		return nil, false
	}
	return o.CpuCfsQuotaPeriod, true
}

// HasCpuCfsQuotaPeriod returns a boolean if a field has been set.
func (o *Kubeletconfig) HasCpuCfsQuotaPeriod() bool {
	if o != nil && !isNil(o.CpuCfsQuotaPeriod) {
		return true
	}

	return false
}

// SetCpuCfsQuotaPeriod gets a reference to the given string and assigns it to the CpuCfsQuotaPeriod field.
func (o *Kubeletconfig) SetCpuCfsQuotaPeriod(v string) {
	o.CpuCfsQuotaPeriod = &v
}

// GetCpuManagerPolicy returns the CpuManagerPolicy field value if set, zero value otherwise.
func (o *Kubeletconfig) GetCpuManagerPolicy() string {
	if o == nil || isNil(o.CpuManagerPolicy) {
		var ret string
		return ret
	}
	return *o.CpuManagerPolicy
}

// GetCpuManagerPolicyOk returns a tuple with the CpuManagerPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kubeletconfig) GetCpuManagerPolicyOk() (*string, bool) {
	if o == nil || isNil(o.CpuManagerPolicy) {
		return nil, false
	}
	return o.CpuManagerPolicy, true
}

// HasCpuManagerPolicy returns a boolean if a field has been set.
func (o *Kubeletconfig) HasCpuManagerPolicy() bool {
	if o != nil && !isNil(o.CpuManagerPolicy) {
		return true
	}

	return false
}

// SetCpuManagerPolicy gets a reference to the given string and assigns it to the CpuManagerPolicy field.
func (o *Kubeletconfig) SetCpuManagerPolicy(v string) {
	o.CpuManagerPolicy = &v
}

// GetFailSwapOn returns the FailSwapOn field value if set, zero value otherwise.
func (o *Kubeletconfig) GetFailSwapOn() bool {
	if o == nil || isNil(o.FailSwapOn) {
		var ret bool
		return ret
	}
	return *o.FailSwapOn
}

// GetFailSwapOnOk returns a tuple with the FailSwapOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kubeletconfig) GetFailSwapOnOk() (*bool, bool) {
	if o == nil || isNil(o.FailSwapOn) {
		return nil, false
	}
	return o.FailSwapOn, true
}

// HasFailSwapOn returns a boolean if a field has been set.
func (o *Kubeletconfig) HasFailSwapOn() bool {
	if o != nil && !isNil(o.FailSwapOn) {
		return true
	}

	return false
}

// SetFailSwapOn gets a reference to the given bool and assigns it to the FailSwapOn field.
func (o *Kubeletconfig) SetFailSwapOn(v bool) {
	o.FailSwapOn = &v
}

// GetImageGcHighThreshold returns the ImageGcHighThreshold field value if set, zero value otherwise.
func (o *Kubeletconfig) GetImageGcHighThreshold() int32 {
	if o == nil || isNil(o.ImageGcHighThreshold) {
		var ret int32
		return ret
	}
	return *o.ImageGcHighThreshold
}

// GetImageGcHighThresholdOk returns a tuple with the ImageGcHighThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kubeletconfig) GetImageGcHighThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.ImageGcHighThreshold) {
		return nil, false
	}
	return o.ImageGcHighThreshold, true
}

// HasImageGcHighThreshold returns a boolean if a field has been set.
func (o *Kubeletconfig) HasImageGcHighThreshold() bool {
	if o != nil && !isNil(o.ImageGcHighThreshold) {
		return true
	}

	return false
}

// SetImageGcHighThreshold gets a reference to the given int32 and assigns it to the ImageGcHighThreshold field.
func (o *Kubeletconfig) SetImageGcHighThreshold(v int32) {
	o.ImageGcHighThreshold = &v
}

// GetImageGcLowThreshold returns the ImageGcLowThreshold field value if set, zero value otherwise.
func (o *Kubeletconfig) GetImageGcLowThreshold() int32 {
	if o == nil || isNil(o.ImageGcLowThreshold) {
		var ret int32
		return ret
	}
	return *o.ImageGcLowThreshold
}

// GetImageGcLowThresholdOk returns a tuple with the ImageGcLowThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kubeletconfig) GetImageGcLowThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.ImageGcLowThreshold) {
		return nil, false
	}
	return o.ImageGcLowThreshold, true
}

// HasImageGcLowThreshold returns a boolean if a field has been set.
func (o *Kubeletconfig) HasImageGcLowThreshold() bool {
	if o != nil && !isNil(o.ImageGcLowThreshold) {
		return true
	}

	return false
}

// SetImageGcLowThreshold gets a reference to the given int32 and assigns it to the ImageGcLowThreshold field.
func (o *Kubeletconfig) SetImageGcLowThreshold(v int32) {
	o.ImageGcLowThreshold = &v
}

// GetPodMaxPids returns the PodMaxPids field value if set, zero value otherwise.
func (o *Kubeletconfig) GetPodMaxPids() int32 {
	if o == nil || isNil(o.PodMaxPids) {
		var ret int32
		return ret
	}
	return *o.PodMaxPids
}

// GetPodMaxPidsOk returns a tuple with the PodMaxPids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kubeletconfig) GetPodMaxPidsOk() (*int32, bool) {
	if o == nil || isNil(o.PodMaxPids) {
		return nil, false
	}
	return o.PodMaxPids, true
}

// HasPodMaxPids returns a boolean if a field has been set.
func (o *Kubeletconfig) HasPodMaxPids() bool {
	if o != nil && !isNil(o.PodMaxPids) {
		return true
	}

	return false
}

// SetPodMaxPids gets a reference to the given int32 and assigns it to the PodMaxPids field.
func (o *Kubeletconfig) SetPodMaxPids(v int32) {
	o.PodMaxPids = &v
}

// GetTopologyManagerPolicy returns the TopologyManagerPolicy field value if set, zero value otherwise.
func (o *Kubeletconfig) GetTopologyManagerPolicy() string {
	if o == nil || isNil(o.TopologyManagerPolicy) {
		var ret string
		return ret
	}
	return *o.TopologyManagerPolicy
}

// GetTopologyManagerPolicyOk returns a tuple with the TopologyManagerPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kubeletconfig) GetTopologyManagerPolicyOk() (*string, bool) {
	if o == nil || isNil(o.TopologyManagerPolicy) {
		return nil, false
	}
	return o.TopologyManagerPolicy, true
}

// HasTopologyManagerPolicy returns a boolean if a field has been set.
func (o *Kubeletconfig) HasTopologyManagerPolicy() bool {
	if o != nil && !isNil(o.TopologyManagerPolicy) {
		return true
	}

	return false
}

// SetTopologyManagerPolicy gets a reference to the given string and assigns it to the TopologyManagerPolicy field.
func (o *Kubeletconfig) SetTopologyManagerPolicy(v string) {
	o.TopologyManagerPolicy = &v
}

func (o Kubeletconfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Kubeletconfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllowedUnsafeSysctls) {
		toSerialize["allowedUnsafeSysctls"] = o.AllowedUnsafeSysctls
	}
	if !isNil(o.ContainerLogMaxFiles) {
		toSerialize["containerLogMaxFiles"] = o.ContainerLogMaxFiles
	}
	if !isNil(o.ContainerLogMaxSizeMB) {
		toSerialize["containerLogMaxSizeMB"] = o.ContainerLogMaxSizeMB
	}
	if !isNil(o.CpuCfsQuota) {
		toSerialize["cpuCfsQuota"] = o.CpuCfsQuota
	}
	if !isNil(o.CpuCfsQuotaPeriod) {
		toSerialize["cpuCfsQuotaPeriod"] = o.CpuCfsQuotaPeriod
	}
	if !isNil(o.CpuManagerPolicy) {
		toSerialize["cpuManagerPolicy"] = o.CpuManagerPolicy
	}
	if !isNil(o.FailSwapOn) {
		toSerialize["failSwapOn"] = o.FailSwapOn
	}
	if !isNil(o.ImageGcHighThreshold) {
		toSerialize["imageGcHighThreshold"] = o.ImageGcHighThreshold
	}
	if !isNil(o.ImageGcLowThreshold) {
		toSerialize["imageGcLowThreshold"] = o.ImageGcLowThreshold
	}
	if !isNil(o.PodMaxPids) {
		toSerialize["podMaxPids"] = o.PodMaxPids
	}
	if !isNil(o.TopologyManagerPolicy) {
		toSerialize["topologyManagerPolicy"] = o.TopologyManagerPolicy
	}
	return toSerialize, nil
}

type NullableKubeletconfig struct {
	value *Kubeletconfig
	isSet bool
}

func (v NullableKubeletconfig) Get() *Kubeletconfig {
	return v.value
}

func (v *NullableKubeletconfig) Set(val *Kubeletconfig) {
	v.value = val
	v.isSet = true
}

func (v NullableKubeletconfig) IsSet() bool {
	return v.isSet
}

func (v *NullableKubeletconfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubeletconfig(val *Kubeletconfig) *NullableKubeletconfig {
	return &NullableKubeletconfig{value: val, isSet: true}
}

func (v NullableKubeletconfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubeletconfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


