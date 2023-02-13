# GitRepoConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | [**[]File**](File.md) | paths in the git repo to watch for changes | 
**Provider** | **string** | provider for the git repo | 
**Repository** | **string** | name of the git repository | 
**Revision** | **string** | branch or tag in the git repository to watch for changes | 

## Methods

### NewGitRepoConfig

`func NewGitRepoConfig(paths []File, provider string, repository string, revision string, ) *GitRepoConfig`

NewGitRepoConfig instantiates a new GitRepoConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRepoConfigWithDefaults

`func NewGitRepoConfigWithDefaults() *GitRepoConfig`

NewGitRepoConfigWithDefaults instantiates a new GitRepoConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *GitRepoConfig) GetPaths() []File`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *GitRepoConfig) GetPathsOk() (*[]File, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *GitRepoConfig) SetPaths(v []File)`

SetPaths sets Paths field to given value.


### GetProvider

`func (o *GitRepoConfig) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GitRepoConfig) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GitRepoConfig) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRepository

`func (o *GitRepoConfig) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *GitRepoConfig) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *GitRepoConfig) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetRevision

`func (o *GitRepoConfig) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GitRepoConfig) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GitRepoConfig) SetRevision(v string)`

SetRevision sets Revision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


