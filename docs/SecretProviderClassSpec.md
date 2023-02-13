# SecretProviderClassSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifact** | Pointer to [**ArtifactSpec**](ArtifactSpec.md) |  | [optional] 
**Parameters** | Pointer to **map[string]string** | secret provider class parameters | [optional] 
**Provider** | Pointer to **string** | Name of the secret provider class | [optional] 
**SecretObject** | Pointer to [**[]SecretObject**](SecretObject.md) |  | [optional] 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 

## Methods

### NewSecretProviderClassSpec

`func NewSecretProviderClassSpec() *SecretProviderClassSpec`

NewSecretProviderClassSpec instantiates a new SecretProviderClassSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretProviderClassSpecWithDefaults

`func NewSecretProviderClassSpecWithDefaults() *SecretProviderClassSpec`

NewSecretProviderClassSpecWithDefaults instantiates a new SecretProviderClassSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifact

`func (o *SecretProviderClassSpec) GetArtifact() ArtifactSpec`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *SecretProviderClassSpec) GetArtifactOk() (*ArtifactSpec, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *SecretProviderClassSpec) SetArtifact(v ArtifactSpec)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *SecretProviderClassSpec) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetParameters

`func (o *SecretProviderClassSpec) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *SecretProviderClassSpec) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *SecretProviderClassSpec) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *SecretProviderClassSpec) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetProvider

`func (o *SecretProviderClassSpec) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SecretProviderClassSpec) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SecretProviderClassSpec) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SecretProviderClassSpec) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSecretObject

`func (o *SecretProviderClassSpec) GetSecretObject() []SecretObject`

GetSecretObject returns the SecretObject field if non-nil, zero value otherwise.

### GetSecretObjectOk

`func (o *SecretProviderClassSpec) GetSecretObjectOk() (*[]SecretObject, bool)`

GetSecretObjectOk returns a tuple with the SecretObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretObject

`func (o *SecretProviderClassSpec) SetSecretObject(v []SecretObject)`

SetSecretObject sets SecretObject field to given value.

### HasSecretObject

`func (o *SecretProviderClassSpec) HasSecretObject() bool`

HasSecretObject returns a boolean if a field has been set.

### GetSharing

`func (o *SecretProviderClassSpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *SecretProviderClassSpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *SecretProviderClassSpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *SecretProviderClassSpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


