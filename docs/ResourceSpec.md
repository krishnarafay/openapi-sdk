# ResourceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variables** | [**[]Variable**](Variable.md) | Variables data for resource | 

## Methods

### NewResourceSpec

`func NewResourceSpec(variables []Variable, ) *ResourceSpec`

NewResourceSpec instantiates a new ResourceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSpecWithDefaults

`func NewResourceSpecWithDefaults() *ResourceSpec`

NewResourceSpecWithDefaults instantiates a new ResourceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariables

`func (o *ResourceSpec) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ResourceSpec) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ResourceSpec) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


