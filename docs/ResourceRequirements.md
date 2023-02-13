# ResourceRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limits** | [**ResourceQuantity**](ResourceQuantity.md) |  | 

## Methods

### NewResourceRequirements

`func NewResourceRequirements(limits ResourceQuantity, ) *ResourceRequirements`

NewResourceRequirements instantiates a new ResourceRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRequirementsWithDefaults

`func NewResourceRequirementsWithDefaults() *ResourceRequirements`

NewResourceRequirementsWithDefaults instantiates a new ResourceRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimits

`func (o *ResourceRequirements) GetLimits() ResourceQuantity`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *ResourceRequirements) GetLimitsOk() (*ResourceQuantity, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *ResourceRequirements) SetLimits(v ResourceQuantity)`

SetLimits sets Limits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


