# ClusterNetworkPolicyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the policy list resource | [optional] [readonly] 
**Items** | Pointer to [**[]ClusterNetworkPolicy**](ClusterNetworkPolicy.md) | List of the cluster policy resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the policy list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewClusterNetworkPolicyList

`func NewClusterNetworkPolicyList() *ClusterNetworkPolicyList`

NewClusterNetworkPolicyList instantiates a new ClusterNetworkPolicyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNetworkPolicyListWithDefaults

`func NewClusterNetworkPolicyListWithDefaults() *ClusterNetworkPolicyList`

NewClusterNetworkPolicyListWithDefaults instantiates a new ClusterNetworkPolicyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ClusterNetworkPolicyList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterNetworkPolicyList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterNetworkPolicyList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterNetworkPolicyList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *ClusterNetworkPolicyList) GetItems() []ClusterNetworkPolicy`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ClusterNetworkPolicyList) GetItemsOk() (*[]ClusterNetworkPolicy, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ClusterNetworkPolicyList) SetItems(v []ClusterNetworkPolicy)`

SetItems sets Items field to given value.

### HasItems

`func (o *ClusterNetworkPolicyList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *ClusterNetworkPolicyList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterNetworkPolicyList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterNetworkPolicyList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterNetworkPolicyList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *ClusterNetworkPolicyList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterNetworkPolicyList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterNetworkPolicyList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ClusterNetworkPolicyList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


