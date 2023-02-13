# ProjectResourceQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMaps** | Pointer to **string** | The maximum number of ConfigMaps that can exist in the project/namespace | [optional] 
**CpuLimits** | Pointer to **string** | The maximum amount of CPU (in millicores) allocated to the project/namespace | [optional] 
**CpuRequests** | Pointer to **string** | The minimum amount of CPU (in millicores) guaranteed to the project/namespace | [optional] 
**MemoryLimits** | Pointer to **string** | The maximum amount of memory (in bytes) allocated to the project/namespace | [optional] 
**MemoryRequests** | Pointer to **string** | The minimum amount of memory (in bytes) guaranteed to the project/namespace | [optional] 
**PersistentVolumeClaims** | Pointer to **string** | The maximum number of persistent volume claims that can exist in the project/namespace | [optional] 
**Pods** | Pointer to **string** | The maximum number of pods that can exist in the project/namespace in a non-terminal state(i.e., pods with a state of .status.phase in (Failed, Succeeded) equal to true) | [optional] 
**ReplicationControllers** | Pointer to **string** | The maximum number of replication controllers that can exist in the project/namespace | [optional] 
**Secrets** | Pointer to **string** | The maximum number of secrets that can exist in the project/namespace | [optional] 
**Services** | Pointer to **string** | The maximum number of services that can exist in the project/namespace | [optional] 
**ServicesLoadBalancers** | Pointer to **string** | The maximum number of load balancers services that can exist in the project/namespace | [optional] 
**ServicesNodePorts** | Pointer to **string** | The maximum number of node port services that can exist in the project/namespace | [optional] 
**StorageRequests** | Pointer to **string** | The minimum amount of storage (in gigabytes) guaranteed to the project/namespace | [optional] 

## Methods

### NewProjectResourceQuota

`func NewProjectResourceQuota() *ProjectResourceQuota`

NewProjectResourceQuota instantiates a new ProjectResourceQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectResourceQuotaWithDefaults

`func NewProjectResourceQuotaWithDefaults() *ProjectResourceQuota`

NewProjectResourceQuotaWithDefaults instantiates a new ProjectResourceQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMaps

`func (o *ProjectResourceQuota) GetConfigMaps() string`

GetConfigMaps returns the ConfigMaps field if non-nil, zero value otherwise.

### GetConfigMapsOk

`func (o *ProjectResourceQuota) GetConfigMapsOk() (*string, bool)`

GetConfigMapsOk returns a tuple with the ConfigMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMaps

`func (o *ProjectResourceQuota) SetConfigMaps(v string)`

SetConfigMaps sets ConfigMaps field to given value.

### HasConfigMaps

`func (o *ProjectResourceQuota) HasConfigMaps() bool`

HasConfigMaps returns a boolean if a field has been set.

### GetCpuLimits

`func (o *ProjectResourceQuota) GetCpuLimits() string`

GetCpuLimits returns the CpuLimits field if non-nil, zero value otherwise.

### GetCpuLimitsOk

`func (o *ProjectResourceQuota) GetCpuLimitsOk() (*string, bool)`

GetCpuLimitsOk returns a tuple with the CpuLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimits

`func (o *ProjectResourceQuota) SetCpuLimits(v string)`

SetCpuLimits sets CpuLimits field to given value.

### HasCpuLimits

`func (o *ProjectResourceQuota) HasCpuLimits() bool`

HasCpuLimits returns a boolean if a field has been set.

### GetCpuRequests

`func (o *ProjectResourceQuota) GetCpuRequests() string`

GetCpuRequests returns the CpuRequests field if non-nil, zero value otherwise.

### GetCpuRequestsOk

`func (o *ProjectResourceQuota) GetCpuRequestsOk() (*string, bool)`

GetCpuRequestsOk returns a tuple with the CpuRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuRequests

`func (o *ProjectResourceQuota) SetCpuRequests(v string)`

SetCpuRequests sets CpuRequests field to given value.

### HasCpuRequests

`func (o *ProjectResourceQuota) HasCpuRequests() bool`

HasCpuRequests returns a boolean if a field has been set.

### GetMemoryLimits

`func (o *ProjectResourceQuota) GetMemoryLimits() string`

GetMemoryLimits returns the MemoryLimits field if non-nil, zero value otherwise.

### GetMemoryLimitsOk

`func (o *ProjectResourceQuota) GetMemoryLimitsOk() (*string, bool)`

GetMemoryLimitsOk returns a tuple with the MemoryLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimits

`func (o *ProjectResourceQuota) SetMemoryLimits(v string)`

SetMemoryLimits sets MemoryLimits field to given value.

### HasMemoryLimits

`func (o *ProjectResourceQuota) HasMemoryLimits() bool`

HasMemoryLimits returns a boolean if a field has been set.

### GetMemoryRequests

`func (o *ProjectResourceQuota) GetMemoryRequests() string`

GetMemoryRequests returns the MemoryRequests field if non-nil, zero value otherwise.

### GetMemoryRequestsOk

`func (o *ProjectResourceQuota) GetMemoryRequestsOk() (*string, bool)`

GetMemoryRequestsOk returns a tuple with the MemoryRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequests

`func (o *ProjectResourceQuota) SetMemoryRequests(v string)`

SetMemoryRequests sets MemoryRequests field to given value.

### HasMemoryRequests

`func (o *ProjectResourceQuota) HasMemoryRequests() bool`

HasMemoryRequests returns a boolean if a field has been set.

### GetPersistentVolumeClaims

`func (o *ProjectResourceQuota) GetPersistentVolumeClaims() string`

GetPersistentVolumeClaims returns the PersistentVolumeClaims field if non-nil, zero value otherwise.

### GetPersistentVolumeClaimsOk

`func (o *ProjectResourceQuota) GetPersistentVolumeClaimsOk() (*string, bool)`

GetPersistentVolumeClaimsOk returns a tuple with the PersistentVolumeClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeClaims

`func (o *ProjectResourceQuota) SetPersistentVolumeClaims(v string)`

SetPersistentVolumeClaims sets PersistentVolumeClaims field to given value.

### HasPersistentVolumeClaims

`func (o *ProjectResourceQuota) HasPersistentVolumeClaims() bool`

HasPersistentVolumeClaims returns a boolean if a field has been set.

### GetPods

`func (o *ProjectResourceQuota) GetPods() string`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *ProjectResourceQuota) GetPodsOk() (*string, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *ProjectResourceQuota) SetPods(v string)`

SetPods sets Pods field to given value.

### HasPods

`func (o *ProjectResourceQuota) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetReplicationControllers

`func (o *ProjectResourceQuota) GetReplicationControllers() string`

GetReplicationControllers returns the ReplicationControllers field if non-nil, zero value otherwise.

### GetReplicationControllersOk

`func (o *ProjectResourceQuota) GetReplicationControllersOk() (*string, bool)`

GetReplicationControllersOk returns a tuple with the ReplicationControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationControllers

`func (o *ProjectResourceQuota) SetReplicationControllers(v string)`

SetReplicationControllers sets ReplicationControllers field to given value.

### HasReplicationControllers

`func (o *ProjectResourceQuota) HasReplicationControllers() bool`

HasReplicationControllers returns a boolean if a field has been set.

### GetSecrets

`func (o *ProjectResourceQuota) GetSecrets() string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ProjectResourceQuota) GetSecretsOk() (*string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ProjectResourceQuota) SetSecrets(v string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *ProjectResourceQuota) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetServices

`func (o *ProjectResourceQuota) GetServices() string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ProjectResourceQuota) GetServicesOk() (*string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ProjectResourceQuota) SetServices(v string)`

SetServices sets Services field to given value.

### HasServices

`func (o *ProjectResourceQuota) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetServicesLoadBalancers

`func (o *ProjectResourceQuota) GetServicesLoadBalancers() string`

GetServicesLoadBalancers returns the ServicesLoadBalancers field if non-nil, zero value otherwise.

### GetServicesLoadBalancersOk

`func (o *ProjectResourceQuota) GetServicesLoadBalancersOk() (*string, bool)`

GetServicesLoadBalancersOk returns a tuple with the ServicesLoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesLoadBalancers

`func (o *ProjectResourceQuota) SetServicesLoadBalancers(v string)`

SetServicesLoadBalancers sets ServicesLoadBalancers field to given value.

### HasServicesLoadBalancers

`func (o *ProjectResourceQuota) HasServicesLoadBalancers() bool`

HasServicesLoadBalancers returns a boolean if a field has been set.

### GetServicesNodePorts

`func (o *ProjectResourceQuota) GetServicesNodePorts() string`

GetServicesNodePorts returns the ServicesNodePorts field if non-nil, zero value otherwise.

### GetServicesNodePortsOk

`func (o *ProjectResourceQuota) GetServicesNodePortsOk() (*string, bool)`

GetServicesNodePortsOk returns a tuple with the ServicesNodePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesNodePorts

`func (o *ProjectResourceQuota) SetServicesNodePorts(v string)`

SetServicesNodePorts sets ServicesNodePorts field to given value.

### HasServicesNodePorts

`func (o *ProjectResourceQuota) HasServicesNodePorts() bool`

HasServicesNodePorts returns a boolean if a field has been set.

### GetStorageRequests

`func (o *ProjectResourceQuota) GetStorageRequests() string`

GetStorageRequests returns the StorageRequests field if non-nil, zero value otherwise.

### GetStorageRequestsOk

`func (o *ProjectResourceQuota) GetStorageRequestsOk() (*string, bool)`

GetStorageRequestsOk returns a tuple with the StorageRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRequests

`func (o *ProjectResourceQuota) SetStorageRequests(v string)`

SetStorageRequests sets StorageRequests field to given value.

### HasStorageRequests

`func (o *ProjectResourceQuota) HasStorageRequests() bool`

HasStorageRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


