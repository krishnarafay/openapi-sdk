# OverrideResourceSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **map[string]string** | map of key value labels | [optional] 
**Name** | Pointer to **string** | select resource by name | [optional] 
**Selector** | Pointer to **string** | kubernetes style resource selector | [optional] 

## Methods

### NewOverrideResourceSelector

`func NewOverrideResourceSelector() *OverrideResourceSelector`

NewOverrideResourceSelector instantiates a new OverrideResourceSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideResourceSelectorWithDefaults

`func NewOverrideResourceSelectorWithDefaults() *OverrideResourceSelector`

NewOverrideResourceSelectorWithDefaults instantiates a new OverrideResourceSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *OverrideResourceSelector) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *OverrideResourceSelector) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *OverrideResourceSelector) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *OverrideResourceSelector) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *OverrideResourceSelector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OverrideResourceSelector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OverrideResourceSelector) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OverrideResourceSelector) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSelector

`func (o *OverrideResourceSelector) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *OverrideResourceSelector) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *OverrideResourceSelector) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *OverrideResourceSelector) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


