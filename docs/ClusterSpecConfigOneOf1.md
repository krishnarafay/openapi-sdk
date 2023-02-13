# ClusterSpecConfigOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Spec** | Pointer to [**AksV3Spec**](AksV3Spec.md) |  | [optional] 

## Methods

### NewClusterSpecConfigOneOf1

`func NewClusterSpecConfigOneOf1() *ClusterSpecConfigOneOf1`

NewClusterSpecConfigOneOf1 instantiates a new ClusterSpecConfigOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSpecConfigOneOf1WithDefaults

`func NewClusterSpecConfigOneOf1WithDefaults() *ClusterSpecConfigOneOf1`

NewClusterSpecConfigOneOf1WithDefaults instantiates a new ClusterSpecConfigOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ClusterSpecConfigOneOf1) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterSpecConfigOneOf1) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterSpecConfigOneOf1) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterSpecConfigOneOf1) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ClusterSpecConfigOneOf1) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterSpecConfigOneOf1) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterSpecConfigOneOf1) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterSpecConfigOneOf1) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *ClusterSpecConfigOneOf1) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterSpecConfigOneOf1) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterSpecConfigOneOf1) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ClusterSpecConfigOneOf1) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *ClusterSpecConfigOneOf1) GetSpec() AksV3Spec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ClusterSpecConfigOneOf1) GetSpecOk() (*AksV3Spec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ClusterSpecConfigOneOf1) SetSpec(v AksV3Spec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ClusterSpecConfigOneOf1) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


