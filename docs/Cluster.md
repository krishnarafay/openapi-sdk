# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | api version | [optional] [default to "infra.k8smgmt.io/v3"]
**Kind** | Pointer to **string** | kind | [optional] [default to "Cluster"]
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Spec** | Pointer to [**ClusterSpec**](ClusterSpec.md) |  | [optional] 
**Status** | Pointer to [**ClusterStatus**](ClusterStatus.md) |  | [optional] 

## Methods

### NewCluster

`func NewCluster() *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *Cluster) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Cluster) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Cluster) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *Cluster) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *Cluster) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Cluster) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Cluster) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Cluster) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *Cluster) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Cluster) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Cluster) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Cluster) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *Cluster) GetSpec() ClusterSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Cluster) GetSpecOk() (*ClusterSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Cluster) SetSpec(v ClusterSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *Cluster) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *Cluster) GetStatus() ClusterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cluster) GetStatusOk() (*ClusterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cluster) SetStatus(v ClusterStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Cluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


