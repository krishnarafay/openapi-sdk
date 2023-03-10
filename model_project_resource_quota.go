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

// checks if the ProjectResourceQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectResourceQuota{}

// ProjectResourceQuota Resource quota definition for a cluster/ namespace in a project
type ProjectResourceQuota struct {
	// The maximum number of ConfigMaps that can exist in the project/namespace
	ConfigMaps *string `json:"configMaps,omitempty"`
	// The maximum amount of CPU (in millicores) allocated to the project/namespace
	CpuLimits *string `json:"cpuLimits,omitempty"`
	// The minimum amount of CPU (in millicores) guaranteed to the project/namespace
	CpuRequests *string `json:"cpuRequests,omitempty"`
	// The maximum amount of memory (in bytes) allocated to the project/namespace
	MemoryLimits *string `json:"memoryLimits,omitempty"`
	// The minimum amount of memory (in bytes) guaranteed to the project/namespace
	MemoryRequests *string `json:"memoryRequests,omitempty"`
	// The maximum number of persistent volume claims that can exist in the project/namespace
	PersistentVolumeClaims *string `json:"persistentVolumeClaims,omitempty"`
	// The maximum number of pods that can exist in the project/namespace in a non-terminal state(i.e., pods with a state of .status.phase in (Failed, Succeeded) equal to true)
	Pods *string `json:"pods,omitempty"`
	// The maximum number of replication controllers that can exist in the project/namespace
	ReplicationControllers *string `json:"replicationControllers,omitempty"`
	// The maximum number of secrets that can exist in the project/namespace
	Secrets *string `json:"secrets,omitempty"`
	// The maximum number of services that can exist in the project/namespace
	Services *string `json:"services,omitempty"`
	// The maximum number of load balancers services that can exist in the project/namespace
	ServicesLoadBalancers *string `json:"servicesLoadBalancers,omitempty"`
	// The maximum number of node port services that can exist in the project/namespace
	ServicesNodePorts *string `json:"servicesNodePorts,omitempty"`
	// The minimum amount of storage (in gigabytes) guaranteed to the project/namespace
	StorageRequests *string `json:"storageRequests,omitempty"`
}

// NewProjectResourceQuota instantiates a new ProjectResourceQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectResourceQuota() *ProjectResourceQuota {
	this := ProjectResourceQuota{}
	return &this
}

// NewProjectResourceQuotaWithDefaults instantiates a new ProjectResourceQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectResourceQuotaWithDefaults() *ProjectResourceQuota {
	this := ProjectResourceQuota{}
	return &this
}

// GetConfigMaps returns the ConfigMaps field value if set, zero value otherwise.
func (o *ProjectResourceQuota) GetConfigMaps() string {
	if o == nil || isNil(o.ConfigMaps) {
		var ret string
		return ret
	}
	return *o.ConfigMaps
}

// GetConfigMapsOk returns a tuple with the ConfigMaps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResourceQuota) GetConfigMapsOk() (*string, bool) {
	if o == nil || isNil(o.ConfigMaps) {
		return nil, false
	}
	return o.ConfigMaps, true
}

// HasConfigMaps returns a boolean if a field has been set.
func (o *ProjectResourceQuota) HasConfigMaps() bool {
	if o != nil && !isNil(o.ConfigMaps) {
		return true
	}

	return false
}

// SetConfigMaps gets a reference to the given string and assigns it to the ConfigMaps field.
func (o *ProjectResourceQuota) SetConfigMaps(v string) {
	o.ConfigMaps = &v
}

// GetCpuLimits returns the CpuLimits field value if set, zero value otherwise.
func (o *ProjectResourceQuota) GetCpuLimits() string {
	if o == nil || isNil(o.CpuLimits) {
		var ret string
		return ret
	}
	return *o.CpuLimits
}

// GetCpuLimitsOk returns a tuple with the CpuLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResourceQuota) GetCpuLimitsOk() (*string, bool) {
	if o == nil || isNil(o.CpuLimits) {
		return nil, false
	}
	return o.CpuLimits, true
}

// HasCpuLimits returns a boolean if a field has been set.
func (o *ProjectResourceQuota) HasCpuLimits() bool {
	if o != nil && !isNil(o.CpuLimits) {
		return true
	}

	return false
}

// SetCpuLimits gets a reference to the given string and assigns it to the CpuLimits field.
func (o *ProjectResourceQuota) SetCpuLimits(v string) {
	o.CpuLimits = &v
}

// GetCpuRequests returns the CpuRequests field value if set, zero value otherwise.
func (o *ProjectResourceQuota) GetCpuRequests() string {
	if o == nil || isNil(o.CpuRequests) {
		var ret string
		return ret
	}
	return *o.CpuRequests
}

// GetCpuRequestsOk returns a tuple with the CpuRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResourceQuota) GetCpuRequestsOk() (*string, bool) {
	if o == nil || isNil(o.CpuRequests) {
		return nil, false
	}
	return o.CpuRequests, true
}

// HasCpuRequests returns a boolean if a field has been set.
func (o *ProjectResourceQuota) HasCpuRequests() bool {
	if o != nil && !isNil(o.CpuRequests) {
		return true
	}

	return false
}

// SetCpuRequests gets a reference to the given string and assigns it to the CpuRequests field.
func (o *ProjectResourceQuota) SetCpuRequests(v string) {
	o.CpuRequests = &v
}

// GetMemoryLimits returns the MemoryLimits field value if set, zero value otherwise.
func (o *ProjectResourceQuota) GetMemoryLimits() string {
	if o == nil || isNil(o.MemoryLimits) {
		var ret string
		return ret
	}
	return *o.MemoryLimits
}

// GetMemoryLimitsOk returns a tuple with the MemoryLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResourceQuota) GetMemoryLimitsOk() (*string, bool) {
	if o == nil || isNil(o.MemoryLimits) {
		return nil, false
	}
	return o.MemoryLimits, true
}

// HasMemoryLimits returns a boolean if a field has been set.
func (o *ProjectResourceQuota) HasMemoryLimits() bool {
	if o != nil && !isNil(o.MemoryLimits) {
		return true
	}

	return false
}

// SetMemoryLimits gets a reference to the given string and assigns it to the MemoryLimits field.
func (o *ProjectResourceQuota) SetMemoryLimits(v string) {
	o.MemoryLimits = &v
}

// GetMemoryRequests returns the MemoryRequests field value if set, zero value otherwise.
func (o *ProjectResourceQuota) GetMemoryRequests() string {
	if o == nil || isNil(o.MemoryRequests) {
		var ret string
		return ret
	}
	return *o.MemoryRequests
}

// GetMemoryRequestsOk returns a tuple with the MemoryRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResourceQuota) GetMemoryRequestsOk() (*string, bool) {
	if o == nil || isNil(o.MemoryRequests) {
		return nil, false
	}
	return o.MemoryRequests, true
}

// HasMemoryRequests returns a boolean if a field has been set.
func (o *ProjectResourceQuota) HasMemoryRequests() bool {
	if o != nil && !isNil(o.MemoryRequests) {
		return true
	}

	return false
}

// SetMemoryRequests gets a reference to the given string and assigns it to the MemoryRequests field.
func (o *ProjectResourceQuota) SetMemoryRequests(v string) {
	o.MemoryRequests = &v
}

// GetPersistentVolumeClaims returns the PersistentVolumeClaims field value if set, zero value otherwise.
func (o *ProjectResourceQuota) GetPersistentVolumeClaims() string {
	if o == nil || isNil(o.PersistentVolumeClaims) {
		var ret string
		return ret
	}
	return *o.PersistentVolumeClaims
}

// GetPersistentVolumeClaimsOk returns a tuple with the PersistentVolumeClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResourceQuota) GetPersistentVolumeClaimsOk() (*string, bool) {
	if o == nil || isNil(o.PersistentVolumeClaims) {
		return nil, false
	}
	return o.PersistentVolumeClaims, true
}

// HasPersistentVolumeClaims returns a boolean if a field has been set.
func (o *ProjectResourceQuota) HasPersistentVolumeClaims() bool {
	if o != nil && !isNil(o.PersistentVolumeClaims) {
		return true
	}

	return false
}

// SetPersistentVolumeClaims gets a reference to the given string and assigns it to the PersistentVolumeClaims field.
func (o *ProjectResourceQuota) SetPersistentVolumeClaims(v string) {
	o.PersistentVolumeClaims = &v
}

// GetPods returns the Pods field value if set, zero value otherwise.
func (o *ProjectResourceQuota) GetPods() string {
	if o == nil || isNil(o.Pods) {
		var ret string
		return ret
	}
	return *o.Pods
}

// GetPodsOk returns a tuple with the Pods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResourceQuota) GetPodsOk() (*string, bool) {
	if o == nil || isNil(o.Pods) {
		return nil, false
	}
	return o.Pods, true
}

// HasPods returns a boolean if a field has been set.
func (o *ProjectResourceQuota) HasPods() bool {
	if o != nil && !isNil(o.Pods) {
		return true
	}

	return false
}

// SetPods gets a reference to the given string and assigns it to the Pods field.
func (o *ProjectResourceQuota) SetPods(v string) {
	o.Pods = &v
}

// GetReplicationControllers returns the ReplicationControllers field value if set, zero value otherwise.
func (o *ProjectResourceQuota) GetReplicationControllers() string {
	if o == nil || isNil(o.ReplicationControllers) {
		var ret string
		return ret
	}
	return *o.ReplicationControllers
}

// GetReplicationControllersOk returns a tuple with the ReplicationControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResourceQuota) GetReplicationControllersOk() (*string, bool) {
	if o == nil || isNil(o.ReplicationControllers) {
		return nil, false
	}
	return o.ReplicationControllers, true
}

// HasReplicationControllers returns a boolean if a field has been set.
func (o *ProjectResourceQuota) HasReplicationControllers() bool {
	if o != nil && !isNil(o.ReplicationControllers) {
		return true
	}

	return false
}

// SetReplicationControllers gets a reference to the given string and assigns it to the ReplicationControllers field.
func (o *ProjectResourceQuota) SetReplicationControllers(v string) {
	o.ReplicationControllers = &v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *ProjectResourceQuota) GetSecrets() string {
	if o == nil || isNil(o.Secrets) {
		var ret string
		return ret
	}
	return *o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResourceQuota) GetSecretsOk() (*string, bool) {
	if o == nil || isNil(o.Secrets) {
		return nil, false
	}
	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *ProjectResourceQuota) HasSecrets() bool {
	if o != nil && !isNil(o.Secrets) {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given string and assigns it to the Secrets field.
func (o *ProjectResourceQuota) SetSecrets(v string) {
	o.Secrets = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *ProjectResourceQuota) GetServices() string {
	if o == nil || isNil(o.Services) {
		var ret string
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResourceQuota) GetServicesOk() (*string, bool) {
	if o == nil || isNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *ProjectResourceQuota) HasServices() bool {
	if o != nil && !isNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given string and assigns it to the Services field.
func (o *ProjectResourceQuota) SetServices(v string) {
	o.Services = &v
}

// GetServicesLoadBalancers returns the ServicesLoadBalancers field value if set, zero value otherwise.
func (o *ProjectResourceQuota) GetServicesLoadBalancers() string {
	if o == nil || isNil(o.ServicesLoadBalancers) {
		var ret string
		return ret
	}
	return *o.ServicesLoadBalancers
}

// GetServicesLoadBalancersOk returns a tuple with the ServicesLoadBalancers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResourceQuota) GetServicesLoadBalancersOk() (*string, bool) {
	if o == nil || isNil(o.ServicesLoadBalancers) {
		return nil, false
	}
	return o.ServicesLoadBalancers, true
}

// HasServicesLoadBalancers returns a boolean if a field has been set.
func (o *ProjectResourceQuota) HasServicesLoadBalancers() bool {
	if o != nil && !isNil(o.ServicesLoadBalancers) {
		return true
	}

	return false
}

// SetServicesLoadBalancers gets a reference to the given string and assigns it to the ServicesLoadBalancers field.
func (o *ProjectResourceQuota) SetServicesLoadBalancers(v string) {
	o.ServicesLoadBalancers = &v
}

// GetServicesNodePorts returns the ServicesNodePorts field value if set, zero value otherwise.
func (o *ProjectResourceQuota) GetServicesNodePorts() string {
	if o == nil || isNil(o.ServicesNodePorts) {
		var ret string
		return ret
	}
	return *o.ServicesNodePorts
}

// GetServicesNodePortsOk returns a tuple with the ServicesNodePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResourceQuota) GetServicesNodePortsOk() (*string, bool) {
	if o == nil || isNil(o.ServicesNodePorts) {
		return nil, false
	}
	return o.ServicesNodePorts, true
}

// HasServicesNodePorts returns a boolean if a field has been set.
func (o *ProjectResourceQuota) HasServicesNodePorts() bool {
	if o != nil && !isNil(o.ServicesNodePorts) {
		return true
	}

	return false
}

// SetServicesNodePorts gets a reference to the given string and assigns it to the ServicesNodePorts field.
func (o *ProjectResourceQuota) SetServicesNodePorts(v string) {
	o.ServicesNodePorts = &v
}

// GetStorageRequests returns the StorageRequests field value if set, zero value otherwise.
func (o *ProjectResourceQuota) GetStorageRequests() string {
	if o == nil || isNil(o.StorageRequests) {
		var ret string
		return ret
	}
	return *o.StorageRequests
}

// GetStorageRequestsOk returns a tuple with the StorageRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResourceQuota) GetStorageRequestsOk() (*string, bool) {
	if o == nil || isNil(o.StorageRequests) {
		return nil, false
	}
	return o.StorageRequests, true
}

// HasStorageRequests returns a boolean if a field has been set.
func (o *ProjectResourceQuota) HasStorageRequests() bool {
	if o != nil && !isNil(o.StorageRequests) {
		return true
	}

	return false
}

// SetStorageRequests gets a reference to the given string and assigns it to the StorageRequests field.
func (o *ProjectResourceQuota) SetStorageRequests(v string) {
	o.StorageRequests = &v
}

func (o ProjectResourceQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectResourceQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ConfigMaps) {
		toSerialize["configMaps"] = o.ConfigMaps
	}
	if !isNil(o.CpuLimits) {
		toSerialize["cpuLimits"] = o.CpuLimits
	}
	if !isNil(o.CpuRequests) {
		toSerialize["cpuRequests"] = o.CpuRequests
	}
	if !isNil(o.MemoryLimits) {
		toSerialize["memoryLimits"] = o.MemoryLimits
	}
	if !isNil(o.MemoryRequests) {
		toSerialize["memoryRequests"] = o.MemoryRequests
	}
	if !isNil(o.PersistentVolumeClaims) {
		toSerialize["persistentVolumeClaims"] = o.PersistentVolumeClaims
	}
	if !isNil(o.Pods) {
		toSerialize["pods"] = o.Pods
	}
	if !isNil(o.ReplicationControllers) {
		toSerialize["replicationControllers"] = o.ReplicationControllers
	}
	if !isNil(o.Secrets) {
		toSerialize["secrets"] = o.Secrets
	}
	if !isNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	if !isNil(o.ServicesLoadBalancers) {
		toSerialize["servicesLoadBalancers"] = o.ServicesLoadBalancers
	}
	if !isNil(o.ServicesNodePorts) {
		toSerialize["servicesNodePorts"] = o.ServicesNodePorts
	}
	if !isNil(o.StorageRequests) {
		toSerialize["storageRequests"] = o.StorageRequests
	}
	return toSerialize, nil
}

type NullableProjectResourceQuota struct {
	value *ProjectResourceQuota
	isSet bool
}

func (v NullableProjectResourceQuota) Get() *ProjectResourceQuota {
	return v.value
}

func (v *NullableProjectResourceQuota) Set(val *ProjectResourceQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectResourceQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectResourceQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectResourceQuota(val *ProjectResourceQuota) *NullableProjectResourceQuota {
	return &NullableProjectResourceQuota{value: val, isSet: true}
}

func (v NullableProjectResourceQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectResourceQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


