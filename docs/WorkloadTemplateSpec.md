# WorkloadTemplateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifact** | Pointer to [**ArtifactSpec**](ArtifactSpec.md) |  | [optional] 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 

## Methods

### NewWorkloadTemplateSpec

`func NewWorkloadTemplateSpec() *WorkloadTemplateSpec`

NewWorkloadTemplateSpec instantiates a new WorkloadTemplateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadTemplateSpecWithDefaults

`func NewWorkloadTemplateSpecWithDefaults() *WorkloadTemplateSpec`

NewWorkloadTemplateSpecWithDefaults instantiates a new WorkloadTemplateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifact

`func (o *WorkloadTemplateSpec) GetArtifact() ArtifactSpec`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *WorkloadTemplateSpec) GetArtifactOk() (*ArtifactSpec, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *WorkloadTemplateSpec) SetArtifact(v ArtifactSpec)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *WorkloadTemplateSpec) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetSharing

`func (o *WorkloadTemplateSpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *WorkloadTemplateSpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *WorkloadTemplateSpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *WorkloadTemplateSpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


