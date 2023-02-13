# HelmInCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | Pointer to **string** | name of the helm catalog | [optional] 
**ChartName** | **string** | name of the chart | 
**ChartVersion** | **string** | version of the chart | 
**ValuesPaths** | Pointer to [**[]File**](File.md) | relative paths to values files | [optional] 
**ValuesRef** | Pointer to [**OverrideRepoReference**](OverrideRepoReference.md) |  | [optional] 

## Methods

### NewHelmInCatalog

`func NewHelmInCatalog(chartName string, chartVersion string, ) *HelmInCatalog`

NewHelmInCatalog instantiates a new HelmInCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmInCatalogWithDefaults

`func NewHelmInCatalogWithDefaults() *HelmInCatalog`

NewHelmInCatalogWithDefaults instantiates a new HelmInCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *HelmInCatalog) GetCatalog() string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *HelmInCatalog) GetCatalogOk() (*string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *HelmInCatalog) SetCatalog(v string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *HelmInCatalog) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetChartName

`func (o *HelmInCatalog) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *HelmInCatalog) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *HelmInCatalog) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartVersion

`func (o *HelmInCatalog) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *HelmInCatalog) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *HelmInCatalog) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetValuesPaths

`func (o *HelmInCatalog) GetValuesPaths() []File`

GetValuesPaths returns the ValuesPaths field if non-nil, zero value otherwise.

### GetValuesPathsOk

`func (o *HelmInCatalog) GetValuesPathsOk() (*[]File, bool)`

GetValuesPathsOk returns a tuple with the ValuesPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesPaths

`func (o *HelmInCatalog) SetValuesPaths(v []File)`

SetValuesPaths sets ValuesPaths field to given value.

### HasValuesPaths

`func (o *HelmInCatalog) HasValuesPaths() bool`

HasValuesPaths returns a boolean if a field has been set.

### GetValuesRef

`func (o *HelmInCatalog) GetValuesRef() OverrideRepoReference`

GetValuesRef returns the ValuesRef field if non-nil, zero value otherwise.

### GetValuesRefOk

`func (o *HelmInCatalog) GetValuesRefOk() (*OverrideRepoReference, bool)`

GetValuesRefOk returns a tuple with the ValuesRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesRef

`func (o *HelmInCatalog) SetValuesRef(v OverrideRepoReference)`

SetValuesRef sets ValuesRef field to given value.

### HasValuesRef

`func (o *HelmInCatalog) HasValuesRef() bool`

HasValuesRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


