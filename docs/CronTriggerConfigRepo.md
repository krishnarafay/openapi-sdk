# CronTriggerConfigRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | [**[]File**](File.md) | paths in the git repo to watch for changes | 
**Provider** | **string** | provider for the git repo | 
**Repository** | **string** | name of the helm repository | 
**Revision** | **string** | branch or tag in the git repository to watch for changes | 
**ChartName** | **string** | name of the chart in repo | 
**ChartVersion** | **string** | version of the chart in repo | 

## Methods

### NewCronTriggerConfigRepo

`func NewCronTriggerConfigRepo(paths []File, provider string, repository string, revision string, chartName string, chartVersion string, ) *CronTriggerConfigRepo`

NewCronTriggerConfigRepo instantiates a new CronTriggerConfigRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronTriggerConfigRepoWithDefaults

`func NewCronTriggerConfigRepoWithDefaults() *CronTriggerConfigRepo`

NewCronTriggerConfigRepoWithDefaults instantiates a new CronTriggerConfigRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *CronTriggerConfigRepo) GetPaths() []File`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *CronTriggerConfigRepo) GetPathsOk() (*[]File, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *CronTriggerConfigRepo) SetPaths(v []File)`

SetPaths sets Paths field to given value.


### GetProvider

`func (o *CronTriggerConfigRepo) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CronTriggerConfigRepo) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CronTriggerConfigRepo) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRepository

`func (o *CronTriggerConfigRepo) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CronTriggerConfigRepo) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CronTriggerConfigRepo) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetRevision

`func (o *CronTriggerConfigRepo) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *CronTriggerConfigRepo) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *CronTriggerConfigRepo) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetChartName

`func (o *CronTriggerConfigRepo) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *CronTriggerConfigRepo) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *CronTriggerConfigRepo) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartVersion

`func (o *CronTriggerConfigRepo) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *CronTriggerConfigRepo) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *CronTriggerConfigRepo) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


