# Variable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the variable | [optional] 
**Options** | Pointer to [**VariableOptions**](VariableOptions.md) |  | [optional] 
**Value** | Pointer to **string** | Value of the variable in the specified format | [optional] 
**ValueType** | Pointer to **string** | Specify the variable value type | [optional] 

## Methods

### NewVariable

`func NewVariable() *Variable`

NewVariable instantiates a new Variable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableWithDefaults

`func NewVariableWithDefaults() *Variable`

NewVariableWithDefaults instantiates a new Variable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Variable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Variable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Variable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Variable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *Variable) GetOptions() VariableOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Variable) GetOptionsOk() (*VariableOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Variable) SetOptions(v VariableOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Variable) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetValue

`func (o *Variable) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Variable) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Variable) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Variable) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueType

`func (o *Variable) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *Variable) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *Variable) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *Variable) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


