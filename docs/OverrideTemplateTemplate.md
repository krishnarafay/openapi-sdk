# OverrideTemplateTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inline** | **string** | inline override template | 
**Paths** | [**[]File**](File.md) | paths in the repository containing the override template | 
**Repository** | **string** | git repository containing the override template | 
**Revision** | **string** | branch or tag in the repository | 

## Methods

### NewOverrideTemplateTemplate

`func NewOverrideTemplateTemplate(inline string, paths []File, repository string, revision string, ) *OverrideTemplateTemplate`

NewOverrideTemplateTemplate instantiates a new OverrideTemplateTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideTemplateTemplateWithDefaults

`func NewOverrideTemplateTemplateWithDefaults() *OverrideTemplateTemplate`

NewOverrideTemplateTemplateWithDefaults instantiates a new OverrideTemplateTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInline

`func (o *OverrideTemplateTemplate) GetInline() string`

GetInline returns the Inline field if non-nil, zero value otherwise.

### GetInlineOk

`func (o *OverrideTemplateTemplate) GetInlineOk() (*string, bool)`

GetInlineOk returns a tuple with the Inline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInline

`func (o *OverrideTemplateTemplate) SetInline(v string)`

SetInline sets Inline field to given value.


### GetPaths

`func (o *OverrideTemplateTemplate) GetPaths() []File`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *OverrideTemplateTemplate) GetPathsOk() (*[]File, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *OverrideTemplateTemplate) SetPaths(v []File)`

SetPaths sets Paths field to given value.


### GetRepository

`func (o *OverrideTemplateTemplate) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *OverrideTemplateTemplate) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *OverrideTemplateTemplate) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetRevision

`func (o *OverrideTemplateTemplate) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *OverrideTemplateTemplate) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *OverrideTemplateTemplate) SetRevision(v string)`

SetRevision sets Revision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


