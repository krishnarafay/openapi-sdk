# MeshResourceQuotas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuLimits** | Pointer to **string** | The maximum amount of CPU (in millicores) allocated to the service mesh control plane | [optional] 
**CpuRequests** | Pointer to **string** | The minimum amount of CPU (in millicores) guaranteed to the service mesh control plane | [optional] 
**MemoryLimits** | Pointer to **string** | The maximum amount of memory (in MiB) allocated to the service mesh control plane | [optional] 
**MemoryRequests** | Pointer to **string** | The minimum amount of memory (in MiB) allocated to the service mesh control plane | [optional] 

## Methods

### NewMeshResourceQuotas

`func NewMeshResourceQuotas() *MeshResourceQuotas`

NewMeshResourceQuotas instantiates a new MeshResourceQuotas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeshResourceQuotasWithDefaults

`func NewMeshResourceQuotasWithDefaults() *MeshResourceQuotas`

NewMeshResourceQuotasWithDefaults instantiates a new MeshResourceQuotas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuLimits

`func (o *MeshResourceQuotas) GetCpuLimits() string`

GetCpuLimits returns the CpuLimits field if non-nil, zero value otherwise.

### GetCpuLimitsOk

`func (o *MeshResourceQuotas) GetCpuLimitsOk() (*string, bool)`

GetCpuLimitsOk returns a tuple with the CpuLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimits

`func (o *MeshResourceQuotas) SetCpuLimits(v string)`

SetCpuLimits sets CpuLimits field to given value.

### HasCpuLimits

`func (o *MeshResourceQuotas) HasCpuLimits() bool`

HasCpuLimits returns a boolean if a field has been set.

### GetCpuRequests

`func (o *MeshResourceQuotas) GetCpuRequests() string`

GetCpuRequests returns the CpuRequests field if non-nil, zero value otherwise.

### GetCpuRequestsOk

`func (o *MeshResourceQuotas) GetCpuRequestsOk() (*string, bool)`

GetCpuRequestsOk returns a tuple with the CpuRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuRequests

`func (o *MeshResourceQuotas) SetCpuRequests(v string)`

SetCpuRequests sets CpuRequests field to given value.

### HasCpuRequests

`func (o *MeshResourceQuotas) HasCpuRequests() bool`

HasCpuRequests returns a boolean if a field has been set.

### GetMemoryLimits

`func (o *MeshResourceQuotas) GetMemoryLimits() string`

GetMemoryLimits returns the MemoryLimits field if non-nil, zero value otherwise.

### GetMemoryLimitsOk

`func (o *MeshResourceQuotas) GetMemoryLimitsOk() (*string, bool)`

GetMemoryLimitsOk returns a tuple with the MemoryLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimits

`func (o *MeshResourceQuotas) SetMemoryLimits(v string)`

SetMemoryLimits sets MemoryLimits field to given value.

### HasMemoryLimits

`func (o *MeshResourceQuotas) HasMemoryLimits() bool`

HasMemoryLimits returns a boolean if a field has been set.

### GetMemoryRequests

`func (o *MeshResourceQuotas) GetMemoryRequests() string`

GetMemoryRequests returns the MemoryRequests field if non-nil, zero value otherwise.

### GetMemoryRequestsOk

`func (o *MeshResourceQuotas) GetMemoryRequestsOk() (*string, bool)`

GetMemoryRequestsOk returns a tuple with the MemoryRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequests

`func (o *MeshResourceQuotas) SetMemoryRequests(v string)`

SetMemoryRequests sets MemoryRequests field to given value.

### HasMemoryRequests

`func (o *MeshResourceQuotas) HasMemoryRequests() bool`

HasMemoryRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


