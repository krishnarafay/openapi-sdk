# ClusterNetworkPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the cluster policy resource | [default to "security.k8smgmt.io/v3"]
**Kind** | **string** | Kind of the cluster policy resource | [default to "ClusterNetworkPolicy"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**ClusterNetworkPolicySpec**](ClusterNetworkPolicySpec.md) |  | 

## Methods

### NewClusterNetworkPolicy

`func NewClusterNetworkPolicy(apiVersion string, kind string, metadata Metadata, spec ClusterNetworkPolicySpec, ) *ClusterNetworkPolicy`

NewClusterNetworkPolicy instantiates a new ClusterNetworkPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNetworkPolicyWithDefaults

`func NewClusterNetworkPolicyWithDefaults() *ClusterNetworkPolicy`

NewClusterNetworkPolicyWithDefaults instantiates a new ClusterNetworkPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ClusterNetworkPolicy) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterNetworkPolicy) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterNetworkPolicy) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ClusterNetworkPolicy) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterNetworkPolicy) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterNetworkPolicy) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ClusterNetworkPolicy) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterNetworkPolicy) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterNetworkPolicy) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *ClusterNetworkPolicy) GetSpec() ClusterNetworkPolicySpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ClusterNetworkPolicy) GetSpecOk() (*ClusterNetworkPolicySpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ClusterNetworkPolicy) SetSpec(v ClusterNetworkPolicySpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


