# EnvironmentResourceOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dedicated** | Pointer to **bool** | Specify if the resource is dedicated to workloads/apps | [optional] 

## Methods

### NewEnvironmentResourceOptions

`func NewEnvironmentResourceOptions() *EnvironmentResourceOptions`

NewEnvironmentResourceOptions instantiates a new EnvironmentResourceOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentResourceOptionsWithDefaults

`func NewEnvironmentResourceOptionsWithDefaults() *EnvironmentResourceOptions`

NewEnvironmentResourceOptionsWithDefaults instantiates a new EnvironmentResourceOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicated

`func (o *EnvironmentResourceOptions) GetDedicated() bool`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### GetDedicatedOk

`func (o *EnvironmentResourceOptions) GetDedicatedOk() (*bool, bool)`

GetDedicatedOk returns a tuple with the Dedicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicated

`func (o *EnvironmentResourceOptions) SetDedicated(v bool)`

SetDedicated sets Dedicated field to given value.

### HasDedicated

`func (o *EnvironmentResourceOptions) HasDedicated() bool`

HasDedicated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


