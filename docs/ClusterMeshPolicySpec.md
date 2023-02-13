# ClusterMeshPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]ResourceNameAndVersionRef**](ResourceNameAndVersionRef.md) | Group of cluster policy rules | [optional] 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Version** | Pointer to **string** | version of the cluster policy | [optional] 

## Methods

### NewClusterMeshPolicySpec

`func NewClusterMeshPolicySpec() *ClusterMeshPolicySpec`

NewClusterMeshPolicySpec instantiates a new ClusterMeshPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterMeshPolicySpecWithDefaults

`func NewClusterMeshPolicySpecWithDefaults() *ClusterMeshPolicySpec`

NewClusterMeshPolicySpecWithDefaults instantiates a new ClusterMeshPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *ClusterMeshPolicySpec) GetRules() []ResourceNameAndVersionRef`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ClusterMeshPolicySpec) GetRulesOk() (*[]ResourceNameAndVersionRef, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ClusterMeshPolicySpec) SetRules(v []ResourceNameAndVersionRef)`

SetRules sets Rules field to given value.

### HasRules

`func (o *ClusterMeshPolicySpec) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSharing

`func (o *ClusterMeshPolicySpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *ClusterMeshPolicySpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *ClusterMeshPolicySpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *ClusterMeshPolicySpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetVersion

`func (o *ClusterMeshPolicySpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterMeshPolicySpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterMeshPolicySpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClusterMeshPolicySpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


