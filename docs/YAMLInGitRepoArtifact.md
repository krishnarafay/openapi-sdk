# YAMLInGitRepoArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | [**[]File**](File.md) | relative paths to file in the git repository | 
**Repository** | **string** | name of the git repository | 
**Revision** | **string** | branch or tag in the git repository | 

## Methods

### NewYAMLInGitRepoArtifact

`func NewYAMLInGitRepoArtifact(paths []File, repository string, revision string, ) *YAMLInGitRepoArtifact`

NewYAMLInGitRepoArtifact instantiates a new YAMLInGitRepoArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYAMLInGitRepoArtifactWithDefaults

`func NewYAMLInGitRepoArtifactWithDefaults() *YAMLInGitRepoArtifact`

NewYAMLInGitRepoArtifactWithDefaults instantiates a new YAMLInGitRepoArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *YAMLInGitRepoArtifact) GetPaths() []File`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *YAMLInGitRepoArtifact) GetPathsOk() (*[]File, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *YAMLInGitRepoArtifact) SetPaths(v []File)`

SetPaths sets Paths field to given value.


### GetRepository

`func (o *YAMLInGitRepoArtifact) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *YAMLInGitRepoArtifact) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *YAMLInGitRepoArtifact) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetRevision

`func (o *YAMLInGitRepoArtifact) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *YAMLInGitRepoArtifact) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *YAMLInGitRepoArtifact) SetRevision(v string)`

SetRevision sets Revision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


