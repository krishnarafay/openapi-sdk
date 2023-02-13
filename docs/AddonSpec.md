# AddonSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifact** | [**ArtifactSpec**](ArtifactSpec.md) |  | 
**Namespace** | **string** | namespace of the addon | 
**Sharing** | [**SharingSpec**](SharingSpec.md) |  | 
**Version** | **string** | version of the addon | 

## Methods

### NewAddonSpec

`func NewAddonSpec(artifact ArtifactSpec, namespace string, sharing SharingSpec, version string, ) *AddonSpec`

NewAddonSpec instantiates a new AddonSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonSpecWithDefaults

`func NewAddonSpecWithDefaults() *AddonSpec`

NewAddonSpecWithDefaults instantiates a new AddonSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifact

`func (o *AddonSpec) GetArtifact() ArtifactSpec`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *AddonSpec) GetArtifactOk() (*ArtifactSpec, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *AddonSpec) SetArtifact(v ArtifactSpec)`

SetArtifact sets Artifact field to given value.


### GetNamespace

`func (o *AddonSpec) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AddonSpec) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AddonSpec) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetSharing

`func (o *AddonSpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *AddonSpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *AddonSpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.


### GetVersion

`func (o *AddonSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AddonSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AddonSpec) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


