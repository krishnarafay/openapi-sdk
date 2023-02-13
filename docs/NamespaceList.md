# NamespaceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the namespace list resource | [optional] [readonly] 
**Items** | Pointer to [**[]Namespace**](Namespace.md) | List of the namespace resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the namespace list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewNamespaceList

`func NewNamespaceList() *NamespaceList`

NewNamespaceList instantiates a new NamespaceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceListWithDefaults

`func NewNamespaceListWithDefaults() *NamespaceList`

NewNamespaceListWithDefaults instantiates a new NamespaceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NamespaceList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NamespaceList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NamespaceList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NamespaceList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *NamespaceList) GetItems() []Namespace`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NamespaceList) GetItemsOk() (*[]Namespace, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NamespaceList) SetItems(v []Namespace)`

SetItems sets Items field to given value.

### HasItems

`func (o *NamespaceList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *NamespaceList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NamespaceList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NamespaceList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NamespaceList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *NamespaceList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NamespaceList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NamespaceList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NamespaceList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


