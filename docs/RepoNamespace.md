# RepoNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to [**File**](File.md) |  | [optional] 
**Repository** | Pointer to **string** | name of the git repository | [optional] 
**Revision** | Pointer to **string** | branch or tag in the git repository | [optional] 

## Methods

### NewRepoNamespace

`func NewRepoNamespace() *RepoNamespace`

NewRepoNamespace instantiates a new RepoNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepoNamespaceWithDefaults

`func NewRepoNamespaceWithDefaults() *RepoNamespace`

NewRepoNamespaceWithDefaults instantiates a new RepoNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *RepoNamespace) GetPath() File`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RepoNamespace) GetPathOk() (*File, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RepoNamespace) SetPath(v File)`

SetPath sets Path field to given value.

### HasPath

`func (o *RepoNamespace) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRepository

`func (o *RepoNamespace) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RepoNamespace) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RepoNamespace) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *RepoNamespace) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRevision

`func (o *RepoNamespace) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *RepoNamespace) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *RepoNamespace) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *RepoNamespace) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


