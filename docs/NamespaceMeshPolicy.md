# NamespaceMeshPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the namespace policy resource | [default to "servicemesh.k8smgmt.io/v3"]
**Kind** | **string** | Kind of the namespace policy resource | [default to "NamespaceMeshPolicy"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**NamespaceMeshPolicySpec**](NamespaceMeshPolicySpec.md) |  | 

## Methods

### NewNamespaceMeshPolicy

`func NewNamespaceMeshPolicy(apiVersion string, kind string, metadata Metadata, spec NamespaceMeshPolicySpec, ) *NamespaceMeshPolicy`

NewNamespaceMeshPolicy instantiates a new NamespaceMeshPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceMeshPolicyWithDefaults

`func NewNamespaceMeshPolicyWithDefaults() *NamespaceMeshPolicy`

NewNamespaceMeshPolicyWithDefaults instantiates a new NamespaceMeshPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NamespaceMeshPolicy) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NamespaceMeshPolicy) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NamespaceMeshPolicy) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *NamespaceMeshPolicy) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NamespaceMeshPolicy) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NamespaceMeshPolicy) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NamespaceMeshPolicy) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NamespaceMeshPolicy) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NamespaceMeshPolicy) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *NamespaceMeshPolicy) GetSpec() NamespaceMeshPolicySpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NamespaceMeshPolicy) GetSpecOk() (*NamespaceMeshPolicySpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NamespaceMeshPolicy) SetSpec(v NamespaceMeshPolicySpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


