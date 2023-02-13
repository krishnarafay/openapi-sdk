# NamespaceMeshRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the namespace mesh rule resource | [default to "servicemesh.k8smgmt.io/v3"]
**Kind** | **string** | Kind of the namespace mesh rule resource | [default to "NamespaceMeshRule"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**NamespaceMeshRuleSpec**](NamespaceMeshRuleSpec.md) |  | 

## Methods

### NewNamespaceMeshRule

`func NewNamespaceMeshRule(apiVersion string, kind string, metadata Metadata, spec NamespaceMeshRuleSpec, ) *NamespaceMeshRule`

NewNamespaceMeshRule instantiates a new NamespaceMeshRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceMeshRuleWithDefaults

`func NewNamespaceMeshRuleWithDefaults() *NamespaceMeshRule`

NewNamespaceMeshRuleWithDefaults instantiates a new NamespaceMeshRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NamespaceMeshRule) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NamespaceMeshRule) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NamespaceMeshRule) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *NamespaceMeshRule) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NamespaceMeshRule) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NamespaceMeshRule) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NamespaceMeshRule) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NamespaceMeshRule) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NamespaceMeshRule) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *NamespaceMeshRule) GetSpec() NamespaceMeshRuleSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NamespaceMeshRule) GetSpecOk() (*NamespaceMeshRuleSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NamespaceMeshRule) SetSpec(v NamespaceMeshRuleSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


