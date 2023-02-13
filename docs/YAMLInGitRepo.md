# YAMLInGitRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | [**[]File**](File.md) | relative paths to file in the git repository | 
**Repository** | **string** | name of the git repository | 
**Revision** | **string** | branch or tag in the git repository | 

## Methods

### NewYAMLInGitRepo

`func NewYAMLInGitRepo(paths []File, repository string, revision string, ) *YAMLInGitRepo`

NewYAMLInGitRepo instantiates a new YAMLInGitRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYAMLInGitRepoWithDefaults

`func NewYAMLInGitRepoWithDefaults() *YAMLInGitRepo`

NewYAMLInGitRepoWithDefaults instantiates a new YAMLInGitRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *YAMLInGitRepo) GetPaths() []File`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *YAMLInGitRepo) GetPathsOk() (*[]File, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *YAMLInGitRepo) SetPaths(v []File)`

SetPaths sets Paths field to given value.


### GetRepository

`func (o *YAMLInGitRepo) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *YAMLInGitRepo) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *YAMLInGitRepo) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetRevision

`func (o *YAMLInGitRepo) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *YAMLInGitRepo) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *YAMLInGitRepo) SetRevision(v string)`

SetRevision sets Revision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


