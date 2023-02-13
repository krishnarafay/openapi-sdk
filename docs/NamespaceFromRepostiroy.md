# NamespaceFromRepostiroy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to [**File**](File.md) |  | [optional] 
**Repository** | Pointer to **string** | name of the git repository | [optional] 
**Revision** | Pointer to **string** | branch or tag in the git repository | [optional] 

## Methods

### NewNamespaceFromRepostiroy

`func NewNamespaceFromRepostiroy() *NamespaceFromRepostiroy`

NewNamespaceFromRepostiroy instantiates a new NamespaceFromRepostiroy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceFromRepostiroyWithDefaults

`func NewNamespaceFromRepostiroyWithDefaults() *NamespaceFromRepostiroy`

NewNamespaceFromRepostiroyWithDefaults instantiates a new NamespaceFromRepostiroy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *NamespaceFromRepostiroy) GetPath() File`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *NamespaceFromRepostiroy) GetPathOk() (*File, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *NamespaceFromRepostiroy) SetPath(v File)`

SetPath sets Path field to given value.

### HasPath

`func (o *NamespaceFromRepostiroy) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRepository

`func (o *NamespaceFromRepostiroy) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *NamespaceFromRepostiroy) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *NamespaceFromRepostiroy) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *NamespaceFromRepostiroy) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRevision

`func (o *NamespaceFromRepostiroy) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NamespaceFromRepostiroy) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NamespaceFromRepostiroy) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *NamespaceFromRepostiroy) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


