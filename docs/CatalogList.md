# CatalogList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the catalog list resource | [optional] [readonly] 
**Items** | Pointer to [**[]Catalog**](Catalog.md) | List of the catalog resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the catalog list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewCatalogList

`func NewCatalogList() *CatalogList`

NewCatalogList instantiates a new CatalogList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogListWithDefaults

`func NewCatalogListWithDefaults() *CatalogList`

NewCatalogListWithDefaults instantiates a new CatalogList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CatalogList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CatalogList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CatalogList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CatalogList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *CatalogList) GetItems() []Catalog`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CatalogList) GetItemsOk() (*[]Catalog, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CatalogList) SetItems(v []Catalog)`

SetItems sets Items field to given value.

### HasItems

`func (o *CatalogList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *CatalogList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CatalogList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CatalogList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CatalogList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *CatalogList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CatalogList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CatalogList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CatalogList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


