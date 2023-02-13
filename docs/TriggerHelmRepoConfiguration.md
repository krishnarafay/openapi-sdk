# TriggerHelmRepoConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | **string** | name of the chart in repo | 
**ChartVersion** | **string** | version of the chart in repo | 
**Repository** | **string** | name of the helm repository | 

## Methods

### NewTriggerHelmRepoConfiguration

`func NewTriggerHelmRepoConfiguration(chartName string, chartVersion string, repository string, ) *TriggerHelmRepoConfiguration`

NewTriggerHelmRepoConfiguration instantiates a new TriggerHelmRepoConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerHelmRepoConfigurationWithDefaults

`func NewTriggerHelmRepoConfigurationWithDefaults() *TriggerHelmRepoConfiguration`

NewTriggerHelmRepoConfigurationWithDefaults instantiates a new TriggerHelmRepoConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartName

`func (o *TriggerHelmRepoConfiguration) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *TriggerHelmRepoConfiguration) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *TriggerHelmRepoConfiguration) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartVersion

`func (o *TriggerHelmRepoConfiguration) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *TriggerHelmRepoConfiguration) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *TriggerHelmRepoConfiguration) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetRepository

`func (o *TriggerHelmRepoConfiguration) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *TriggerHelmRepoConfiguration) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *TriggerHelmRepoConfiguration) SetRepository(v string)`

SetRepository sets Repository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


