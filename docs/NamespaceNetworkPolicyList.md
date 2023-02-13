# NamespaceNetworkPolicyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the policy list resource | [optional] [readonly] 
**Items** | Pointer to [**[]NamespaceNetworkPolicy**](NamespaceNetworkPolicy.md) | List of the namespace policy resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the policy list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewNamespaceNetworkPolicyList

`func NewNamespaceNetworkPolicyList() *NamespaceNetworkPolicyList`

NewNamespaceNetworkPolicyList instantiates a new NamespaceNetworkPolicyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceNetworkPolicyListWithDefaults

`func NewNamespaceNetworkPolicyListWithDefaults() *NamespaceNetworkPolicyList`

NewNamespaceNetworkPolicyListWithDefaults instantiates a new NamespaceNetworkPolicyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NamespaceNetworkPolicyList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NamespaceNetworkPolicyList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NamespaceNetworkPolicyList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NamespaceNetworkPolicyList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *NamespaceNetworkPolicyList) GetItems() []NamespaceNetworkPolicy`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NamespaceNetworkPolicyList) GetItemsOk() (*[]NamespaceNetworkPolicy, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NamespaceNetworkPolicyList) SetItems(v []NamespaceNetworkPolicy)`

SetItems sets Items field to given value.

### HasItems

`func (o *NamespaceNetworkPolicyList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *NamespaceNetworkPolicyList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NamespaceNetworkPolicyList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NamespaceNetworkPolicyList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NamespaceNetworkPolicyList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *NamespaceNetworkPolicyList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NamespaceNetworkPolicyList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NamespaceNetworkPolicyList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NamespaceNetworkPolicyList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


