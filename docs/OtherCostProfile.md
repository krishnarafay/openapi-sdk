# OtherCostProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **string** | The cost of cpu per hour  | [optional] 
**Gpu** | Pointer to **string** | The cost of gpu per hour  | [optional] 
**Memory** | Pointer to **string** | The cost of memory per hour  | [optional] 

## Methods

### NewOtherCostProfile

`func NewOtherCostProfile() *OtherCostProfile`

NewOtherCostProfile instantiates a new OtherCostProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtherCostProfileWithDefaults

`func NewOtherCostProfileWithDefaults() *OtherCostProfile`

NewOtherCostProfileWithDefaults instantiates a new OtherCostProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *OtherCostProfile) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *OtherCostProfile) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *OtherCostProfile) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *OtherCostProfile) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetGpu

`func (o *OtherCostProfile) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *OtherCostProfile) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *OtherCostProfile) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *OtherCostProfile) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetMemory

`func (o *OtherCostProfile) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *OtherCostProfile) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *OtherCostProfile) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *OtherCostProfile) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


