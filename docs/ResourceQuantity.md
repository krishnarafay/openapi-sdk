# ResourceQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **string** | quantity of CPU in cores | 
**Memory** | **string** | quantity of memory in MB | 

## Methods

### NewResourceQuantity

`func NewResourceQuantity(cpu string, memory string, ) *ResourceQuantity`

NewResourceQuantity instantiates a new ResourceQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceQuantityWithDefaults

`func NewResourceQuantityWithDefaults() *ResourceQuantity`

NewResourceQuantityWithDefaults instantiates a new ResourceQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *ResourceQuantity) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ResourceQuantity) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ResourceQuantity) SetCpu(v string)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *ResourceQuantity) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ResourceQuantity) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ResourceQuantity) SetMemory(v string)`

SetMemory sets Memory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


