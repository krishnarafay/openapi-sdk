# HelmInHelmRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | **string** | name of the chart | 
**ChartVersion** | **string** | version of the chart | 
**Repository** | **string** | name of the helm repository | 
**ValuesPaths** | Pointer to [**[]File**](File.md) | relative paths to values files | [optional] 
**ValuesRef** | Pointer to [**OverrideRepoReference**](OverrideRepoReference.md) |  | [optional] 

## Methods

### NewHelmInHelmRepo

`func NewHelmInHelmRepo(chartName string, chartVersion string, repository string, ) *HelmInHelmRepo`

NewHelmInHelmRepo instantiates a new HelmInHelmRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmInHelmRepoWithDefaults

`func NewHelmInHelmRepoWithDefaults() *HelmInHelmRepo`

NewHelmInHelmRepoWithDefaults instantiates a new HelmInHelmRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartName

`func (o *HelmInHelmRepo) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *HelmInHelmRepo) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *HelmInHelmRepo) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartVersion

`func (o *HelmInHelmRepo) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *HelmInHelmRepo) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *HelmInHelmRepo) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetRepository

`func (o *HelmInHelmRepo) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *HelmInHelmRepo) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *HelmInHelmRepo) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetValuesPaths

`func (o *HelmInHelmRepo) GetValuesPaths() []File`

GetValuesPaths returns the ValuesPaths field if non-nil, zero value otherwise.

### GetValuesPathsOk

`func (o *HelmInHelmRepo) GetValuesPathsOk() (*[]File, bool)`

GetValuesPathsOk returns a tuple with the ValuesPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesPaths

`func (o *HelmInHelmRepo) SetValuesPaths(v []File)`

SetValuesPaths sets ValuesPaths field to given value.

### HasValuesPaths

`func (o *HelmInHelmRepo) HasValuesPaths() bool`

HasValuesPaths returns a boolean if a field has been set.

### GetValuesRef

`func (o *HelmInHelmRepo) GetValuesRef() OverrideRepoReference`

GetValuesRef returns the ValuesRef field if non-nil, zero value otherwise.

### GetValuesRefOk

`func (o *HelmInHelmRepo) GetValuesRefOk() (*OverrideRepoReference, bool)`

GetValuesRefOk returns a tuple with the ValuesRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesRef

`func (o *HelmInHelmRepo) SetValuesRef(v OverrideRepoReference)`

SetValuesRef sets ValuesRef field to given value.

### HasValuesRef

`func (o *HelmInHelmRepo) HasValuesRef() bool`

HasValuesRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


