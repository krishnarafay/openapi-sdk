# ChargebackGroupSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregate** | Pointer to [**ChargebackAggregate**](ChargebackAggregate.md) |  | [optional] 
**Exclusions** | Pointer to [**[]ChargebackFilter**](ChargebackFilter.md) | exclusions | [optional] 
**Inclusions** | Pointer to [**[]ChargebackFilter**](ChargebackFilter.md) | inclusions | [optional] 
**Type** | Pointer to **string** | type | [optional] 

## Methods

### NewChargebackGroupSpec

`func NewChargebackGroupSpec() *ChargebackGroupSpec`

NewChargebackGroupSpec instantiates a new ChargebackGroupSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargebackGroupSpecWithDefaults

`func NewChargebackGroupSpecWithDefaults() *ChargebackGroupSpec`

NewChargebackGroupSpecWithDefaults instantiates a new ChargebackGroupSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregate

`func (o *ChargebackGroupSpec) GetAggregate() ChargebackAggregate`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *ChargebackGroupSpec) GetAggregateOk() (*ChargebackAggregate, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *ChargebackGroupSpec) SetAggregate(v ChargebackAggregate)`

SetAggregate sets Aggregate field to given value.

### HasAggregate

`func (o *ChargebackGroupSpec) HasAggregate() bool`

HasAggregate returns a boolean if a field has been set.

### GetExclusions

`func (o *ChargebackGroupSpec) GetExclusions() []ChargebackFilter`

GetExclusions returns the Exclusions field if non-nil, zero value otherwise.

### GetExclusionsOk

`func (o *ChargebackGroupSpec) GetExclusionsOk() (*[]ChargebackFilter, bool)`

GetExclusionsOk returns a tuple with the Exclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusions

`func (o *ChargebackGroupSpec) SetExclusions(v []ChargebackFilter)`

SetExclusions sets Exclusions field to given value.

### HasExclusions

`func (o *ChargebackGroupSpec) HasExclusions() bool`

HasExclusions returns a boolean if a field has been set.

### GetInclusions

`func (o *ChargebackGroupSpec) GetInclusions() []ChargebackFilter`

GetInclusions returns the Inclusions field if non-nil, zero value otherwise.

### GetInclusionsOk

`func (o *ChargebackGroupSpec) GetInclusionsOk() (*[]ChargebackFilter, bool)`

GetInclusionsOk returns a tuple with the Inclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusions

`func (o *ChargebackGroupSpec) SetInclusions(v []ChargebackFilter)`

SetInclusions sets Inclusions field to given value.

### HasInclusions

`func (o *ChargebackGroupSpec) HasInclusions() bool`

HasInclusions returns a boolean if a field has been set.

### GetType

`func (o *ChargebackGroupSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChargebackGroupSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChargebackGroupSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChargebackGroupSpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


