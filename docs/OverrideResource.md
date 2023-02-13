# OverrideResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Selector** | [**OverrideResourceSelector**](OverrideResourceSelector.md) |  | 
**Type** | **string** | type of the resource | 

## Methods

### NewOverrideResource

`func NewOverrideResource(selector OverrideResourceSelector, type_ string, ) *OverrideResource`

NewOverrideResource instantiates a new OverrideResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideResourceWithDefaults

`func NewOverrideResourceWithDefaults() *OverrideResource`

NewOverrideResourceWithDefaults instantiates a new OverrideResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelector

`func (o *OverrideResource) GetSelector() OverrideResourceSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *OverrideResource) GetSelectorOk() (*OverrideResourceSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *OverrideResource) SetSelector(v OverrideResourceSelector)`

SetSelector sets Selector field to given value.


### GetType

`func (o *OverrideResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OverrideResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OverrideResource) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


