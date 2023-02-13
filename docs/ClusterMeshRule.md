# ClusterMeshRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the cluster mesh rule resource | [default to "servicemesh.k8smgmt.io/v3"]
**Kind** | **string** | Kind of the cluster mesh rule resource | [default to "ClusterMeshRule"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**ClusterMeshRuleSpec**](ClusterMeshRuleSpec.md) |  | 

## Methods

### NewClusterMeshRule

`func NewClusterMeshRule(apiVersion string, kind string, metadata Metadata, spec ClusterMeshRuleSpec, ) *ClusterMeshRule`

NewClusterMeshRule instantiates a new ClusterMeshRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterMeshRuleWithDefaults

`func NewClusterMeshRuleWithDefaults() *ClusterMeshRule`

NewClusterMeshRuleWithDefaults instantiates a new ClusterMeshRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ClusterMeshRule) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterMeshRule) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterMeshRule) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ClusterMeshRule) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterMeshRule) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterMeshRule) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ClusterMeshRule) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterMeshRule) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterMeshRule) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *ClusterMeshRule) GetSpec() ClusterMeshRuleSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ClusterMeshRule) GetSpecOk() (*ClusterMeshRuleSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ClusterMeshRule) SetSpec(v ClusterMeshRuleSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


