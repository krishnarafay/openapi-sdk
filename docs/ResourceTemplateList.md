# ResourceTemplateList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the resource templates list | [optional] [readonly] 
**Items** | Pointer to [**[]ResourceTemplate**](ResourceTemplate.md) | List of the resource templates | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the resource template list | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewResourceTemplateList

`func NewResourceTemplateList() *ResourceTemplateList`

NewResourceTemplateList instantiates a new ResourceTemplateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTemplateListWithDefaults

`func NewResourceTemplateListWithDefaults() *ResourceTemplateList`

NewResourceTemplateListWithDefaults instantiates a new ResourceTemplateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ResourceTemplateList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ResourceTemplateList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ResourceTemplateList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ResourceTemplateList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *ResourceTemplateList) GetItems() []ResourceTemplate`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ResourceTemplateList) GetItemsOk() (*[]ResourceTemplate, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ResourceTemplateList) SetItems(v []ResourceTemplate)`

SetItems sets Items field to given value.

### HasItems

`func (o *ResourceTemplateList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *ResourceTemplateList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ResourceTemplateList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ResourceTemplateList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ResourceTemplateList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *ResourceTemplateList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceTemplateList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceTemplateList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResourceTemplateList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


