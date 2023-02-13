# HelmInGitRepoArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartPath** | [**File**](File.md) |  | 
**Repository** | **string** | name of the git repository | 
**Revision** | **string** | branch or tag in the git repository | 
**ValuesPaths** | Pointer to [**[]File**](File.md) | relative paths to value file in the git repository | [optional] 
**ValuesRef** | Pointer to [**OverrideRepoReference**](OverrideRepoReference.md) |  | [optional] 

## Methods

### NewHelmInGitRepoArtifact

`func NewHelmInGitRepoArtifact(chartPath File, repository string, revision string, ) *HelmInGitRepoArtifact`

NewHelmInGitRepoArtifact instantiates a new HelmInGitRepoArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmInGitRepoArtifactWithDefaults

`func NewHelmInGitRepoArtifactWithDefaults() *HelmInGitRepoArtifact`

NewHelmInGitRepoArtifactWithDefaults instantiates a new HelmInGitRepoArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartPath

`func (o *HelmInGitRepoArtifact) GetChartPath() File`

GetChartPath returns the ChartPath field if non-nil, zero value otherwise.

### GetChartPathOk

`func (o *HelmInGitRepoArtifact) GetChartPathOk() (*File, bool)`

GetChartPathOk returns a tuple with the ChartPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartPath

`func (o *HelmInGitRepoArtifact) SetChartPath(v File)`

SetChartPath sets ChartPath field to given value.


### GetRepository

`func (o *HelmInGitRepoArtifact) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *HelmInGitRepoArtifact) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *HelmInGitRepoArtifact) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetRevision

`func (o *HelmInGitRepoArtifact) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *HelmInGitRepoArtifact) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *HelmInGitRepoArtifact) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetValuesPaths

`func (o *HelmInGitRepoArtifact) GetValuesPaths() []File`

GetValuesPaths returns the ValuesPaths field if non-nil, zero value otherwise.

### GetValuesPathsOk

`func (o *HelmInGitRepoArtifact) GetValuesPathsOk() (*[]File, bool)`

GetValuesPathsOk returns a tuple with the ValuesPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesPaths

`func (o *HelmInGitRepoArtifact) SetValuesPaths(v []File)`

SetValuesPaths sets ValuesPaths field to given value.

### HasValuesPaths

`func (o *HelmInGitRepoArtifact) HasValuesPaths() bool`

HasValuesPaths returns a boolean if a field has been set.

### GetValuesRef

`func (o *HelmInGitRepoArtifact) GetValuesRef() OverrideRepoReference`

GetValuesRef returns the ValuesRef field if non-nil, zero value otherwise.

### GetValuesRefOk

`func (o *HelmInGitRepoArtifact) GetValuesRefOk() (*OverrideRepoReference, bool)`

GetValuesRefOk returns a tuple with the ValuesRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesRef

`func (o *HelmInGitRepoArtifact) SetValuesRef(v OverrideRepoReference)`

SetValuesRef sets ValuesRef field to given value.

### HasValuesRef

`func (o *HelmInGitRepoArtifact) HasValuesRef() bool`

HasValuesRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


