# HelmInGitRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartPath** | [**File**](File.md) |  | 
**Repository** | **string** | name of the git repository | 
**Revision** | **string** | branch or tag in the git repository | 
**ValuesPaths** | Pointer to [**[]File**](File.md) | relative paths to value file in the git repository | [optional] 
**ValuesRef** | Pointer to [**OverrideRepoReference**](OverrideRepoReference.md) |  | [optional] 

## Methods

### NewHelmInGitRepo

`func NewHelmInGitRepo(chartPath File, repository string, revision string, ) *HelmInGitRepo`

NewHelmInGitRepo instantiates a new HelmInGitRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmInGitRepoWithDefaults

`func NewHelmInGitRepoWithDefaults() *HelmInGitRepo`

NewHelmInGitRepoWithDefaults instantiates a new HelmInGitRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartPath

`func (o *HelmInGitRepo) GetChartPath() File`

GetChartPath returns the ChartPath field if non-nil, zero value otherwise.

### GetChartPathOk

`func (o *HelmInGitRepo) GetChartPathOk() (*File, bool)`

GetChartPathOk returns a tuple with the ChartPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartPath

`func (o *HelmInGitRepo) SetChartPath(v File)`

SetChartPath sets ChartPath field to given value.


### GetRepository

`func (o *HelmInGitRepo) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *HelmInGitRepo) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *HelmInGitRepo) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetRevision

`func (o *HelmInGitRepo) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *HelmInGitRepo) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *HelmInGitRepo) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetValuesPaths

`func (o *HelmInGitRepo) GetValuesPaths() []File`

GetValuesPaths returns the ValuesPaths field if non-nil, zero value otherwise.

### GetValuesPathsOk

`func (o *HelmInGitRepo) GetValuesPathsOk() (*[]File, bool)`

GetValuesPathsOk returns a tuple with the ValuesPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesPaths

`func (o *HelmInGitRepo) SetValuesPaths(v []File)`

SetValuesPaths sets ValuesPaths field to given value.

### HasValuesPaths

`func (o *HelmInGitRepo) HasValuesPaths() bool`

HasValuesPaths returns a boolean if a field has been set.

### GetValuesRef

`func (o *HelmInGitRepo) GetValuesRef() OverrideRepoReference`

GetValuesRef returns the ValuesRef field if non-nil, zero value otherwise.

### GetValuesRefOk

`func (o *HelmInGitRepo) GetValuesRefOk() (*OverrideRepoReference, bool)`

GetValuesRefOk returns a tuple with the ValuesRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesRef

`func (o *HelmInGitRepo) SetValuesRef(v OverrideRepoReference)`

SetValuesRef sets ValuesRef field to given value.

### HasValuesRef

`func (o *HelmInGitRepo) HasValuesRef() bool`

HasValuesRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


