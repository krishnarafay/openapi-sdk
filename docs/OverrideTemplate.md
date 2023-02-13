# OverrideTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to [**OverrideTemplateTemplate**](OverrideTemplateTemplate.md) |  | [optional] 
**Type** | Pointer to **string** | type of override template | [optional] 
**Weight** | Pointer to **int32** | weight of the override, overrides are applied low to high weight | [optional] 

## Methods

### NewOverrideTemplate

`func NewOverrideTemplate() *OverrideTemplate`

NewOverrideTemplate instantiates a new OverrideTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideTemplateWithDefaults

`func NewOverrideTemplateWithDefaults() *OverrideTemplate`

NewOverrideTemplateWithDefaults instantiates a new OverrideTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *OverrideTemplate) GetTemplate() OverrideTemplateTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *OverrideTemplate) GetTemplateOk() (*OverrideTemplateTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *OverrideTemplate) SetTemplate(v OverrideTemplateTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *OverrideTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetType

`func (o *OverrideTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OverrideTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OverrideTemplate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OverrideTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWeight

`func (o *OverrideTemplate) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *OverrideTemplate) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *OverrideTemplate) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *OverrideTemplate) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


