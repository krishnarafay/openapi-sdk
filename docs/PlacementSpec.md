# PlacementSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to [**[]PlacementLabel**](PlacementLabel.md) | list of labels for placement | [optional] 
**Selector** | Pointer to **string** | Kubernetes style label selector | [optional] 

## Methods

### NewPlacementSpec

`func NewPlacementSpec() *PlacementSpec`

NewPlacementSpec instantiates a new PlacementSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementSpecWithDefaults

`func NewPlacementSpecWithDefaults() *PlacementSpec`

NewPlacementSpecWithDefaults instantiates a new PlacementSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *PlacementSpec) GetLabels() []PlacementLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PlacementSpec) GetLabelsOk() (*[]PlacementLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PlacementSpec) SetLabels(v []PlacementLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PlacementSpec) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetSelector

`func (o *PlacementSpec) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *PlacementSpec) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *PlacementSpec) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *PlacementSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


