# AddonList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the addon list resource | [optional] [readonly] 
**Items** | Pointer to [**[]Addon**](Addon.md) | List of the addon resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the addon list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewAddonList

`func NewAddonList() *AddonList`

NewAddonList instantiates a new AddonList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonListWithDefaults

`func NewAddonListWithDefaults() *AddonList`

NewAddonListWithDefaults instantiates a new AddonList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AddonList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AddonList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AddonList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AddonList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *AddonList) GetItems() []Addon`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AddonList) GetItemsOk() (*[]Addon, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AddonList) SetItems(v []Addon)`

SetItems sets Items field to given value.

### HasItems

`func (o *AddonList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *AddonList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AddonList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AddonList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AddonList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *AddonList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddonList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddonList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddonList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


