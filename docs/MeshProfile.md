# MeshProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the profile resource | [default to "servicemesh.k8smgmt.io/v3"]
**Kind** | **string** | Kind of the profile resource | [default to "MeshProfile"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**MeshProfileSpec**](MeshProfileSpec.md) |  | 

## Methods

### NewMeshProfile

`func NewMeshProfile(apiVersion string, kind string, metadata Metadata, spec MeshProfileSpec, ) *MeshProfile`

NewMeshProfile instantiates a new MeshProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeshProfileWithDefaults

`func NewMeshProfileWithDefaults() *MeshProfile`

NewMeshProfileWithDefaults instantiates a new MeshProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MeshProfile) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MeshProfile) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MeshProfile) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *MeshProfile) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MeshProfile) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MeshProfile) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *MeshProfile) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MeshProfile) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MeshProfile) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *MeshProfile) GetSpec() MeshProfileSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MeshProfile) GetSpecOk() (*MeshProfileSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MeshProfile) SetSpec(v MeshProfileSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


