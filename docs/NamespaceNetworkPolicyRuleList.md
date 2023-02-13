# NamespaceNetworkPolicyRuleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the network rule list resource | [optional] [readonly] 
**Items** | Pointer to [**[]NamespaceNetworkPolicyRule**](NamespaceNetworkPolicyRule.md) | List of the network rule resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the network rule list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewNamespaceNetworkPolicyRuleList

`func NewNamespaceNetworkPolicyRuleList() *NamespaceNetworkPolicyRuleList`

NewNamespaceNetworkPolicyRuleList instantiates a new NamespaceNetworkPolicyRuleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceNetworkPolicyRuleListWithDefaults

`func NewNamespaceNetworkPolicyRuleListWithDefaults() *NamespaceNetworkPolicyRuleList`

NewNamespaceNetworkPolicyRuleListWithDefaults instantiates a new NamespaceNetworkPolicyRuleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NamespaceNetworkPolicyRuleList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NamespaceNetworkPolicyRuleList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NamespaceNetworkPolicyRuleList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NamespaceNetworkPolicyRuleList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *NamespaceNetworkPolicyRuleList) GetItems() []NamespaceNetworkPolicyRule`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NamespaceNetworkPolicyRuleList) GetItemsOk() (*[]NamespaceNetworkPolicyRule, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NamespaceNetworkPolicyRuleList) SetItems(v []NamespaceNetworkPolicyRule)`

SetItems sets Items field to given value.

### HasItems

`func (o *NamespaceNetworkPolicyRuleList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *NamespaceNetworkPolicyRuleList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NamespaceNetworkPolicyRuleList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NamespaceNetworkPolicyRuleList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NamespaceNetworkPolicyRuleList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *NamespaceNetworkPolicyRuleList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NamespaceNetworkPolicyRuleList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NamespaceNetworkPolicyRuleList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NamespaceNetworkPolicyRuleList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


