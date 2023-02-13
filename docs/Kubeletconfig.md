# Kubeletconfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedUnsafeSysctls** | Pointer to **[]string** | Allowed list of unsafe sysctls or unsafe sysctl patterns (ending in *). | [optional] 
**ContainerLogMaxFiles** | Pointer to **int32** | The maximum number of container log files that can be present for a container. The number must be ≥ 2. | [optional] 
**ContainerLogMaxSizeMB** | Pointer to **int32** | The maximum size (e.g. 10Mi) of container log file before it is rotated. | [optional] 
**CpuCfsQuota** | Pointer to **bool** | The default is true. | [optional] 
**CpuCfsQuotaPeriod** | Pointer to **string** | The default is 100ms. Valid values are a sequence of decimal numbers with an optional fraction and a unit suffix. For example: 300ms, 2h45m. Supported units are ns, us, ms, s, m, and h. | [optional] 
**CpuManagerPolicy** | Pointer to **string** | The default is none. See Kubernetes CPU management policies for more information. Allowed values are none and static. | [optional] 
**FailSwapOn** | Pointer to **bool** | If set to true it will make the Kubelet fail to start if swap is enabled on the node. | [optional] 
**ImageGcHighThreshold** | Pointer to **int32** | To disable image garbage collection, set to 100. The default is 85% | [optional] 
**ImageGcLowThreshold** | Pointer to **int32** | This cannot be set higher than imageGcHighThreshold. The default is 80% | [optional] 
**PodMaxPids** | Pointer to **int32** | The maximum number of processes per pod. | [optional] 
**TopologyManagerPolicy** | Pointer to **string** | For more information see Kubernetes Topology Manager. The default is none. Allowed values are none, best-effort, restricted, and single-numa-node. | [optional] 

## Methods

### NewKubeletconfig

`func NewKubeletconfig() *Kubeletconfig`

NewKubeletconfig instantiates a new Kubeletconfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeletconfigWithDefaults

`func NewKubeletconfigWithDefaults() *Kubeletconfig`

NewKubeletconfigWithDefaults instantiates a new Kubeletconfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedUnsafeSysctls

`func (o *Kubeletconfig) GetAllowedUnsafeSysctls() []string`

GetAllowedUnsafeSysctls returns the AllowedUnsafeSysctls field if non-nil, zero value otherwise.

### GetAllowedUnsafeSysctlsOk

`func (o *Kubeletconfig) GetAllowedUnsafeSysctlsOk() (*[]string, bool)`

GetAllowedUnsafeSysctlsOk returns a tuple with the AllowedUnsafeSysctls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUnsafeSysctls

`func (o *Kubeletconfig) SetAllowedUnsafeSysctls(v []string)`

SetAllowedUnsafeSysctls sets AllowedUnsafeSysctls field to given value.

### HasAllowedUnsafeSysctls

`func (o *Kubeletconfig) HasAllowedUnsafeSysctls() bool`

HasAllowedUnsafeSysctls returns a boolean if a field has been set.

### GetContainerLogMaxFiles

`func (o *Kubeletconfig) GetContainerLogMaxFiles() int32`

GetContainerLogMaxFiles returns the ContainerLogMaxFiles field if non-nil, zero value otherwise.

### GetContainerLogMaxFilesOk

`func (o *Kubeletconfig) GetContainerLogMaxFilesOk() (*int32, bool)`

GetContainerLogMaxFilesOk returns a tuple with the ContainerLogMaxFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerLogMaxFiles

`func (o *Kubeletconfig) SetContainerLogMaxFiles(v int32)`

SetContainerLogMaxFiles sets ContainerLogMaxFiles field to given value.

### HasContainerLogMaxFiles

`func (o *Kubeletconfig) HasContainerLogMaxFiles() bool`

HasContainerLogMaxFiles returns a boolean if a field has been set.

### GetContainerLogMaxSizeMB

`func (o *Kubeletconfig) GetContainerLogMaxSizeMB() int32`

GetContainerLogMaxSizeMB returns the ContainerLogMaxSizeMB field if non-nil, zero value otherwise.

### GetContainerLogMaxSizeMBOk

`func (o *Kubeletconfig) GetContainerLogMaxSizeMBOk() (*int32, bool)`

GetContainerLogMaxSizeMBOk returns a tuple with the ContainerLogMaxSizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerLogMaxSizeMB

`func (o *Kubeletconfig) SetContainerLogMaxSizeMB(v int32)`

SetContainerLogMaxSizeMB sets ContainerLogMaxSizeMB field to given value.

### HasContainerLogMaxSizeMB

`func (o *Kubeletconfig) HasContainerLogMaxSizeMB() bool`

HasContainerLogMaxSizeMB returns a boolean if a field has been set.

### GetCpuCfsQuota

`func (o *Kubeletconfig) GetCpuCfsQuota() bool`

GetCpuCfsQuota returns the CpuCfsQuota field if non-nil, zero value otherwise.

### GetCpuCfsQuotaOk

`func (o *Kubeletconfig) GetCpuCfsQuotaOk() (*bool, bool)`

GetCpuCfsQuotaOk returns a tuple with the CpuCfsQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCfsQuota

`func (o *Kubeletconfig) SetCpuCfsQuota(v bool)`

SetCpuCfsQuota sets CpuCfsQuota field to given value.

### HasCpuCfsQuota

`func (o *Kubeletconfig) HasCpuCfsQuota() bool`

HasCpuCfsQuota returns a boolean if a field has been set.

### GetCpuCfsQuotaPeriod

`func (o *Kubeletconfig) GetCpuCfsQuotaPeriod() string`

GetCpuCfsQuotaPeriod returns the CpuCfsQuotaPeriod field if non-nil, zero value otherwise.

### GetCpuCfsQuotaPeriodOk

`func (o *Kubeletconfig) GetCpuCfsQuotaPeriodOk() (*string, bool)`

GetCpuCfsQuotaPeriodOk returns a tuple with the CpuCfsQuotaPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCfsQuotaPeriod

`func (o *Kubeletconfig) SetCpuCfsQuotaPeriod(v string)`

SetCpuCfsQuotaPeriod sets CpuCfsQuotaPeriod field to given value.

### HasCpuCfsQuotaPeriod

`func (o *Kubeletconfig) HasCpuCfsQuotaPeriod() bool`

HasCpuCfsQuotaPeriod returns a boolean if a field has been set.

### GetCpuManagerPolicy

`func (o *Kubeletconfig) GetCpuManagerPolicy() string`

GetCpuManagerPolicy returns the CpuManagerPolicy field if non-nil, zero value otherwise.

### GetCpuManagerPolicyOk

`func (o *Kubeletconfig) GetCpuManagerPolicyOk() (*string, bool)`

GetCpuManagerPolicyOk returns a tuple with the CpuManagerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuManagerPolicy

`func (o *Kubeletconfig) SetCpuManagerPolicy(v string)`

SetCpuManagerPolicy sets CpuManagerPolicy field to given value.

### HasCpuManagerPolicy

`func (o *Kubeletconfig) HasCpuManagerPolicy() bool`

HasCpuManagerPolicy returns a boolean if a field has been set.

### GetFailSwapOn

`func (o *Kubeletconfig) GetFailSwapOn() bool`

GetFailSwapOn returns the FailSwapOn field if non-nil, zero value otherwise.

### GetFailSwapOnOk

`func (o *Kubeletconfig) GetFailSwapOnOk() (*bool, bool)`

GetFailSwapOnOk returns a tuple with the FailSwapOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailSwapOn

`func (o *Kubeletconfig) SetFailSwapOn(v bool)`

SetFailSwapOn sets FailSwapOn field to given value.

### HasFailSwapOn

`func (o *Kubeletconfig) HasFailSwapOn() bool`

HasFailSwapOn returns a boolean if a field has been set.

### GetImageGcHighThreshold

`func (o *Kubeletconfig) GetImageGcHighThreshold() int32`

GetImageGcHighThreshold returns the ImageGcHighThreshold field if non-nil, zero value otherwise.

### GetImageGcHighThresholdOk

`func (o *Kubeletconfig) GetImageGcHighThresholdOk() (*int32, bool)`

GetImageGcHighThresholdOk returns a tuple with the ImageGcHighThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageGcHighThreshold

`func (o *Kubeletconfig) SetImageGcHighThreshold(v int32)`

SetImageGcHighThreshold sets ImageGcHighThreshold field to given value.

### HasImageGcHighThreshold

`func (o *Kubeletconfig) HasImageGcHighThreshold() bool`

HasImageGcHighThreshold returns a boolean if a field has been set.

### GetImageGcLowThreshold

`func (o *Kubeletconfig) GetImageGcLowThreshold() int32`

GetImageGcLowThreshold returns the ImageGcLowThreshold field if non-nil, zero value otherwise.

### GetImageGcLowThresholdOk

`func (o *Kubeletconfig) GetImageGcLowThresholdOk() (*int32, bool)`

GetImageGcLowThresholdOk returns a tuple with the ImageGcLowThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageGcLowThreshold

`func (o *Kubeletconfig) SetImageGcLowThreshold(v int32)`

SetImageGcLowThreshold sets ImageGcLowThreshold field to given value.

### HasImageGcLowThreshold

`func (o *Kubeletconfig) HasImageGcLowThreshold() bool`

HasImageGcLowThreshold returns a boolean if a field has been set.

### GetPodMaxPids

`func (o *Kubeletconfig) GetPodMaxPids() int32`

GetPodMaxPids returns the PodMaxPids field if non-nil, zero value otherwise.

### GetPodMaxPidsOk

`func (o *Kubeletconfig) GetPodMaxPidsOk() (*int32, bool)`

GetPodMaxPidsOk returns a tuple with the PodMaxPids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodMaxPids

`func (o *Kubeletconfig) SetPodMaxPids(v int32)`

SetPodMaxPids sets PodMaxPids field to given value.

### HasPodMaxPids

`func (o *Kubeletconfig) HasPodMaxPids() bool`

HasPodMaxPids returns a boolean if a field has been set.

### GetTopologyManagerPolicy

`func (o *Kubeletconfig) GetTopologyManagerPolicy() string`

GetTopologyManagerPolicy returns the TopologyManagerPolicy field if non-nil, zero value otherwise.

### GetTopologyManagerPolicyOk

`func (o *Kubeletconfig) GetTopologyManagerPolicyOk() (*string, bool)`

GetTopologyManagerPolicyOk returns a tuple with the TopologyManagerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyManagerPolicy

`func (o *Kubeletconfig) SetTopologyManagerPolicy(v string)`

SetTopologyManagerPolicy sets TopologyManagerPolicy field to given value.

### HasTopologyManagerPolicy

`func (o *Kubeletconfig) HasTopologyManagerPolicy() bool`

HasTopologyManagerPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


