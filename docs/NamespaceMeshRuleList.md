# NamespaceMeshRuleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the mesh rule list resource | [optional] [readonly] 
**Items** | Pointer to [**[]NamespaceMeshRule**](NamespaceMeshRule.md) | List of the mesh rule resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the mesh rule list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewNamespaceMeshRuleList

`func NewNamespaceMeshRuleList() *NamespaceMeshRuleList`

NewNamespaceMeshRuleList instantiates a new NamespaceMeshRuleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceMeshRuleListWithDefaults

`func NewNamespaceMeshRuleListWithDefaults() *NamespaceMeshRuleList`

NewNamespaceMeshRuleListWithDefaults instantiates a new NamespaceMeshRuleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NamespaceMeshRuleList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NamespaceMeshRuleList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NamespaceMeshRuleList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NamespaceMeshRuleList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *NamespaceMeshRuleList) GetItems() []NamespaceMeshRule`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NamespaceMeshRuleList) GetItemsOk() (*[]NamespaceMeshRule, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NamespaceMeshRuleList) SetItems(v []NamespaceMeshRule)`

SetItems sets Items field to given value.

### HasItems

`func (o *NamespaceMeshRuleList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *NamespaceMeshRuleList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NamespaceMeshRuleList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NamespaceMeshRuleList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NamespaceMeshRuleList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *NamespaceMeshRuleList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NamespaceMeshRuleList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NamespaceMeshRuleList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NamespaceMeshRuleList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


