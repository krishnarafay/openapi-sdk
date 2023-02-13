# OverrideRepoReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | **string** | name of the git repository | 
**Revision** | **string** | branch or tag in the git repository | 
**ValuesPaths** | Pointer to [**[]File**](File.md) | relative path to value file in the git repository | [optional] 

## Methods

### NewOverrideRepoReference

`func NewOverrideRepoReference(repository string, revision string, ) *OverrideRepoReference`

NewOverrideRepoReference instantiates a new OverrideRepoReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideRepoReferenceWithDefaults

`func NewOverrideRepoReferenceWithDefaults() *OverrideRepoReference`

NewOverrideRepoReferenceWithDefaults instantiates a new OverrideRepoReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *OverrideRepoReference) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *OverrideRepoReference) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *OverrideRepoReference) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetRevision

`func (o *OverrideRepoReference) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *OverrideRepoReference) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *OverrideRepoReference) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetValuesPaths

`func (o *OverrideRepoReference) GetValuesPaths() []File`

GetValuesPaths returns the ValuesPaths field if non-nil, zero value otherwise.

### GetValuesPathsOk

`func (o *OverrideRepoReference) GetValuesPathsOk() (*[]File, bool)`

GetValuesPathsOk returns a tuple with the ValuesPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesPaths

`func (o *OverrideRepoReference) SetValuesPaths(v []File)`

SetValuesPaths sets ValuesPaths field to given value.

### HasValuesPaths

`func (o *OverrideRepoReference) HasValuesPaths() bool`

HasValuesPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


