# BlueprintList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the blueprint list resource | [optional] [readonly] 
**Items** | Pointer to [**[]Blueprint**](Blueprint.md) | List of the blueprint resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the blueprint list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewBlueprintList

`func NewBlueprintList() *BlueprintList`

NewBlueprintList instantiates a new BlueprintList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintListWithDefaults

`func NewBlueprintListWithDefaults() *BlueprintList`

NewBlueprintListWithDefaults instantiates a new BlueprintList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *BlueprintList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BlueprintList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BlueprintList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BlueprintList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *BlueprintList) GetItems() []Blueprint`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BlueprintList) GetItemsOk() (*[]Blueprint, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BlueprintList) SetItems(v []Blueprint)`

SetItems sets Items field to given value.

### HasItems

`func (o *BlueprintList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *BlueprintList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BlueprintList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BlueprintList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *BlueprintList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *BlueprintList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BlueprintList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BlueprintList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BlueprintList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


