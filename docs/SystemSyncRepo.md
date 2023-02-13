# SystemSyncRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | [**File**](File.md) |  | 
**Repository** | **string** | name of the git repository | 
**Revision** | **string** | branch or tag in the git repository | 

## Methods

### NewSystemSyncRepo

`func NewSystemSyncRepo(path File, repository string, revision string, ) *SystemSyncRepo`

NewSystemSyncRepo instantiates a new SystemSyncRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemSyncRepoWithDefaults

`func NewSystemSyncRepoWithDefaults() *SystemSyncRepo`

NewSystemSyncRepoWithDefaults instantiates a new SystemSyncRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *SystemSyncRepo) GetPath() File`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SystemSyncRepo) GetPathOk() (*File, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SystemSyncRepo) SetPath(v File)`

SetPath sets Path field to given value.


### GetRepository

`func (o *SystemSyncRepo) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *SystemSyncRepo) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *SystemSyncRepo) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetRevision

`func (o *SystemSyncRepo) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SystemSyncRepo) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SystemSyncRepo) SetRevision(v string)`

SetRevision sets Revision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


