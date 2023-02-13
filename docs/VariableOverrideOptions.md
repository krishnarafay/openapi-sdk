# VariableOverrideOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestrictedValues** | Pointer to **[]string** | If the override type is restricted, values it is restricted to | [optional] 
**Type** | Pointer to **string** | Specify the type of ovverride this variable supports | [optional] 

## Methods

### NewVariableOverrideOptions

`func NewVariableOverrideOptions() *VariableOverrideOptions`

NewVariableOverrideOptions instantiates a new VariableOverrideOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableOverrideOptionsWithDefaults

`func NewVariableOverrideOptionsWithDefaults() *VariableOverrideOptions`

NewVariableOverrideOptionsWithDefaults instantiates a new VariableOverrideOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestrictedValues

`func (o *VariableOverrideOptions) GetRestrictedValues() []string`

GetRestrictedValues returns the RestrictedValues field if non-nil, zero value otherwise.

### GetRestrictedValuesOk

`func (o *VariableOverrideOptions) GetRestrictedValuesOk() (*[]string, bool)`

GetRestrictedValuesOk returns a tuple with the RestrictedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedValues

`func (o *VariableOverrideOptions) SetRestrictedValues(v []string)`

SetRestrictedValues sets RestrictedValues field to given value.

### HasRestrictedValues

`func (o *VariableOverrideOptions) HasRestrictedValues() bool`

HasRestrictedValues returns a boolean if a field has been set.

### GetType

`func (o *VariableOverrideOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VariableOverrideOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VariableOverrideOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VariableOverrideOptions) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


