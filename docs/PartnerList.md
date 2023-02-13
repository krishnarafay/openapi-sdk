# PartnerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]Partner**](Partner.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewPartnerList

`func NewPartnerList() *PartnerList`

NewPartnerList instantiates a new PartnerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerListWithDefaults

`func NewPartnerListWithDefaults() *PartnerList`

NewPartnerListWithDefaults instantiates a new PartnerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *PartnerList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PartnerList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PartnerList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PartnerList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *PartnerList) GetItems() []Partner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PartnerList) GetItemsOk() (*[]Partner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PartnerList) SetItems(v []Partner)`

SetItems sets Items field to given value.

### HasItems

`func (o *PartnerList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *PartnerList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PartnerList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PartnerList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PartnerList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *PartnerList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PartnerList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PartnerList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PartnerList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


