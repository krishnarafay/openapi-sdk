# RepoOverrideTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | [**[]File**](File.md) | paths in the repository containing the override template | 
**Repository** | **string** | git repository containing the override template | 
**Revision** | **string** | branch or tag in the repository | 

## Methods

### NewRepoOverrideTemplate

`func NewRepoOverrideTemplate(paths []File, repository string, revision string, ) *RepoOverrideTemplate`

NewRepoOverrideTemplate instantiates a new RepoOverrideTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepoOverrideTemplateWithDefaults

`func NewRepoOverrideTemplateWithDefaults() *RepoOverrideTemplate`

NewRepoOverrideTemplateWithDefaults instantiates a new RepoOverrideTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *RepoOverrideTemplate) GetPaths() []File`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *RepoOverrideTemplate) GetPathsOk() (*[]File, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *RepoOverrideTemplate) SetPaths(v []File)`

SetPaths sets Paths field to given value.


### GetRepository

`func (o *RepoOverrideTemplate) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RepoOverrideTemplate) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RepoOverrideTemplate) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetRevision

`func (o *RepoOverrideTemplate) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *RepoOverrideTemplate) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *RepoOverrideTemplate) SetRevision(v string)`

SetRevision sets Revision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


