# OPAConstraintSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifact** | Pointer to [**ArtifactSpec**](ArtifactSpec.md) |  | [optional] 
**Published** | Pointer to **bool** | publish status of the constraint | [optional] 
**TemplateName** | Pointer to **string** | Name of the opa constraint template | [optional] 
**Version** | Pointer to **string** | version of the constraint | [optional] 

## Methods

### NewOPAConstraintSpec

`func NewOPAConstraintSpec() *OPAConstraintSpec`

NewOPAConstraintSpec instantiates a new OPAConstraintSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOPAConstraintSpecWithDefaults

`func NewOPAConstraintSpecWithDefaults() *OPAConstraintSpec`

NewOPAConstraintSpecWithDefaults instantiates a new OPAConstraintSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifact

`func (o *OPAConstraintSpec) GetArtifact() ArtifactSpec`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *OPAConstraintSpec) GetArtifactOk() (*ArtifactSpec, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *OPAConstraintSpec) SetArtifact(v ArtifactSpec)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *OPAConstraintSpec) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetPublished

`func (o *OPAConstraintSpec) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *OPAConstraintSpec) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *OPAConstraintSpec) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *OPAConstraintSpec) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetTemplateName

`func (o *OPAConstraintSpec) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *OPAConstraintSpec) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *OPAConstraintSpec) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *OPAConstraintSpec) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetVersion

`func (o *OPAConstraintSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OPAConstraintSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OPAConstraintSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OPAConstraintSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


