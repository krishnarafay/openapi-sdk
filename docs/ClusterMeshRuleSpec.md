# ClusterMeshRuleSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifact** | Pointer to [**ArtifactSpec**](ArtifactSpec.md) |  | [optional] 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Version** | Pointer to **string** | version of the cluster mesh rule | [optional] 

## Methods

### NewClusterMeshRuleSpec

`func NewClusterMeshRuleSpec() *ClusterMeshRuleSpec`

NewClusterMeshRuleSpec instantiates a new ClusterMeshRuleSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterMeshRuleSpecWithDefaults

`func NewClusterMeshRuleSpecWithDefaults() *ClusterMeshRuleSpec`

NewClusterMeshRuleSpecWithDefaults instantiates a new ClusterMeshRuleSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifact

`func (o *ClusterMeshRuleSpec) GetArtifact() ArtifactSpec`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *ClusterMeshRuleSpec) GetArtifactOk() (*ArtifactSpec, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *ClusterMeshRuleSpec) SetArtifact(v ArtifactSpec)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *ClusterMeshRuleSpec) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetSharing

`func (o *ClusterMeshRuleSpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *ClusterMeshRuleSpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *ClusterMeshRuleSpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *ClusterMeshRuleSpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetVersion

`func (o *ClusterMeshRuleSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterMeshRuleSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterMeshRuleSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClusterMeshRuleSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


