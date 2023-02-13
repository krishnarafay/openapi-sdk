# ClusterNetworkPolicyRuleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the network rule list resource | [optional] [readonly] 
**Items** | Pointer to [**[]ClusterNetworkPolicyRule**](ClusterNetworkPolicyRule.md) | List of the network rule resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the network rule list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewClusterNetworkPolicyRuleList

`func NewClusterNetworkPolicyRuleList() *ClusterNetworkPolicyRuleList`

NewClusterNetworkPolicyRuleList instantiates a new ClusterNetworkPolicyRuleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNetworkPolicyRuleListWithDefaults

`func NewClusterNetworkPolicyRuleListWithDefaults() *ClusterNetworkPolicyRuleList`

NewClusterNetworkPolicyRuleListWithDefaults instantiates a new ClusterNetworkPolicyRuleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ClusterNetworkPolicyRuleList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterNetworkPolicyRuleList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterNetworkPolicyRuleList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterNetworkPolicyRuleList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *ClusterNetworkPolicyRuleList) GetItems() []ClusterNetworkPolicyRule`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ClusterNetworkPolicyRuleList) GetItemsOk() (*[]ClusterNetworkPolicyRule, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ClusterNetworkPolicyRuleList) SetItems(v []ClusterNetworkPolicyRule)`

SetItems sets Items field to given value.

### HasItems

`func (o *ClusterNetworkPolicyRuleList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *ClusterNetworkPolicyRuleList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterNetworkPolicyRuleList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterNetworkPolicyRuleList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterNetworkPolicyRuleList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *ClusterNetworkPolicyRuleList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterNetworkPolicyRuleList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterNetworkPolicyRuleList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ClusterNetworkPolicyRuleList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


