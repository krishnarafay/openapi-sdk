# ArtifactSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifact** | [**ArtifactSpecArtifact**](ArtifactSpecArtifact.md) |  | 
**Options** | Pointer to [**ArtifactSpecOptions**](ArtifactSpecOptions.md) |  | [optional] 
**Type** | **string** | type of the artifact | 

## Methods

### NewArtifactSpec

`func NewArtifactSpec(artifact ArtifactSpecArtifact, type_ string, ) *ArtifactSpec`

NewArtifactSpec instantiates a new ArtifactSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactSpecWithDefaults

`func NewArtifactSpecWithDefaults() *ArtifactSpec`

NewArtifactSpecWithDefaults instantiates a new ArtifactSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifact

`func (o *ArtifactSpec) GetArtifact() ArtifactSpecArtifact`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *ArtifactSpec) GetArtifactOk() (*ArtifactSpecArtifact, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *ArtifactSpec) SetArtifact(v ArtifactSpecArtifact)`

SetArtifact sets Artifact field to given value.


### GetOptions

`func (o *ArtifactSpec) GetOptions() ArtifactSpecOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ArtifactSpec) GetOptionsOk() (*ArtifactSpecOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ArtifactSpec) SetOptions(v ArtifactSpecOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ArtifactSpec) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetType

`func (o *ArtifactSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArtifactSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArtifactSpec) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


