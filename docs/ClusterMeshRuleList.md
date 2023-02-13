# ClusterMeshRuleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the mesh rule list resource | [optional] [readonly] 
**Items** | Pointer to [**[]ClusterMeshRule**](ClusterMeshRule.md) | List of the mesh rule resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the mesh rule list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewClusterMeshRuleList

`func NewClusterMeshRuleList() *ClusterMeshRuleList`

NewClusterMeshRuleList instantiates a new ClusterMeshRuleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterMeshRuleListWithDefaults

`func NewClusterMeshRuleListWithDefaults() *ClusterMeshRuleList`

NewClusterMeshRuleListWithDefaults instantiates a new ClusterMeshRuleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ClusterMeshRuleList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterMeshRuleList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterMeshRuleList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterMeshRuleList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *ClusterMeshRuleList) GetItems() []ClusterMeshRule`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ClusterMeshRuleList) GetItemsOk() (*[]ClusterMeshRule, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ClusterMeshRuleList) SetItems(v []ClusterMeshRule)`

SetItems sets Items field to given value.

### HasItems

`func (o *ClusterMeshRuleList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *ClusterMeshRuleList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterMeshRuleList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterMeshRuleList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterMeshRuleList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *ClusterMeshRuleList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterMeshRuleList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterMeshRuleList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ClusterMeshRuleList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


