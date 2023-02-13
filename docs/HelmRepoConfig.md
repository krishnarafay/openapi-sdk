# HelmRepoConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | **string** | name of the chart in repo | 
**ChartVersion** | **string** | version of the chart in repo | 
**Repository** | **string** | name of the helm repository | 

## Methods

### NewHelmRepoConfig

`func NewHelmRepoConfig(chartName string, chartVersion string, repository string, ) *HelmRepoConfig`

NewHelmRepoConfig instantiates a new HelmRepoConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRepoConfigWithDefaults

`func NewHelmRepoConfigWithDefaults() *HelmRepoConfig`

NewHelmRepoConfigWithDefaults instantiates a new HelmRepoConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartName

`func (o *HelmRepoConfig) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *HelmRepoConfig) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *HelmRepoConfig) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartVersion

`func (o *HelmRepoConfig) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *HelmRepoConfig) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *HelmRepoConfig) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetRepository

`func (o *HelmRepoConfig) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *HelmRepoConfig) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *HelmRepoConfig) SetRepository(v string)`

SetRepository sets Repository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


