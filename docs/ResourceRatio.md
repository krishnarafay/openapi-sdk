# ResourceRatio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **float32** | ratio of cpu requests to limits | 
**Memory** | **float32** | ratio of memory requests to limits | 

## Methods

### NewResourceRatio

`func NewResourceRatio(cpu float32, memory float32, ) *ResourceRatio`

NewResourceRatio instantiates a new ResourceRatio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRatioWithDefaults

`func NewResourceRatioWithDefaults() *ResourceRatio`

NewResourceRatioWithDefaults instantiates a new ResourceRatio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *ResourceRatio) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ResourceRatio) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ResourceRatio) SetCpu(v float32)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *ResourceRatio) GetMemory() float32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ResourceRatio) GetMemoryOk() (*float32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ResourceRatio) SetMemory(v float32)`

SetMemory sets Memory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


