# ContainerRegistryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the container registry list resource | [optional] [readonly] 
**Items** | Pointer to [**[]ContainerRegistry**](ContainerRegistry.md) | List of the container registry resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the container registry list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewContainerRegistryList

`func NewContainerRegistryList() *ContainerRegistryList`

NewContainerRegistryList instantiates a new ContainerRegistryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRegistryListWithDefaults

`func NewContainerRegistryListWithDefaults() *ContainerRegistryList`

NewContainerRegistryListWithDefaults instantiates a new ContainerRegistryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ContainerRegistryList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ContainerRegistryList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ContainerRegistryList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ContainerRegistryList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *ContainerRegistryList) GetItems() []ContainerRegistry`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContainerRegistryList) GetItemsOk() (*[]ContainerRegistry, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContainerRegistryList) SetItems(v []ContainerRegistry)`

SetItems sets Items field to given value.

### HasItems

`func (o *ContainerRegistryList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *ContainerRegistryList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ContainerRegistryList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ContainerRegistryList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ContainerRegistryList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *ContainerRegistryList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ContainerRegistryList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ContainerRegistryList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ContainerRegistryList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


