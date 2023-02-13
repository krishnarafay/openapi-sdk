# ArtifactSpecArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | [**[]File**](File.md) | relative paths to file in the git repository | 
**ChartPath** | [**File**](File.md) |  | 
**ValuesPaths** | Pointer to [**[]File**](File.md) | relative paths to values files | [optional] 
**Repository** | **string** | name of the helm repository | 
**Revision** | **string** | branch or tag in the git repository | 
**ValuesRef** | Pointer to [**OverrideRepoReference**](OverrideRepoReference.md) |  | [optional] 
**ChartName** | **string** | name of the chart | 
**ChartVersion** | **string** | version of the chart | 
**Configmap** | [**File**](File.md) |  | 
**Configuration** | [**File**](File.md) |  | 
**Secret** | [**File**](File.md) |  | 
**Statefulset** | [**File**](File.md) |  | 
**Catalog** | Pointer to **string** | name of the helm catalog | [optional] 

## Methods

### NewArtifactSpecArtifact

`func NewArtifactSpecArtifact(paths []File, chartPath File, repository string, revision string, chartName string, chartVersion string, configmap File, configuration File, secret File, statefulset File, ) *ArtifactSpecArtifact`

NewArtifactSpecArtifact instantiates a new ArtifactSpecArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactSpecArtifactWithDefaults

`func NewArtifactSpecArtifactWithDefaults() *ArtifactSpecArtifact`

NewArtifactSpecArtifactWithDefaults instantiates a new ArtifactSpecArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *ArtifactSpecArtifact) GetPaths() []File`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *ArtifactSpecArtifact) GetPathsOk() (*[]File, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *ArtifactSpecArtifact) SetPaths(v []File)`

SetPaths sets Paths field to given value.


### GetChartPath

`func (o *ArtifactSpecArtifact) GetChartPath() File`

GetChartPath returns the ChartPath field if non-nil, zero value otherwise.

### GetChartPathOk

`func (o *ArtifactSpecArtifact) GetChartPathOk() (*File, bool)`

GetChartPathOk returns a tuple with the ChartPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartPath

`func (o *ArtifactSpecArtifact) SetChartPath(v File)`

SetChartPath sets ChartPath field to given value.


### GetValuesPaths

`func (o *ArtifactSpecArtifact) GetValuesPaths() []File`

GetValuesPaths returns the ValuesPaths field if non-nil, zero value otherwise.

### GetValuesPathsOk

`func (o *ArtifactSpecArtifact) GetValuesPathsOk() (*[]File, bool)`

GetValuesPathsOk returns a tuple with the ValuesPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesPaths

`func (o *ArtifactSpecArtifact) SetValuesPaths(v []File)`

SetValuesPaths sets ValuesPaths field to given value.

### HasValuesPaths

`func (o *ArtifactSpecArtifact) HasValuesPaths() bool`

HasValuesPaths returns a boolean if a field has been set.

### GetRepository

`func (o *ArtifactSpecArtifact) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ArtifactSpecArtifact) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ArtifactSpecArtifact) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetRevision

`func (o *ArtifactSpecArtifact) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ArtifactSpecArtifact) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ArtifactSpecArtifact) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetValuesRef

`func (o *ArtifactSpecArtifact) GetValuesRef() OverrideRepoReference`

GetValuesRef returns the ValuesRef field if non-nil, zero value otherwise.

### GetValuesRefOk

`func (o *ArtifactSpecArtifact) GetValuesRefOk() (*OverrideRepoReference, bool)`

GetValuesRefOk returns a tuple with the ValuesRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesRef

`func (o *ArtifactSpecArtifact) SetValuesRef(v OverrideRepoReference)`

SetValuesRef sets ValuesRef field to given value.

### HasValuesRef

`func (o *ArtifactSpecArtifact) HasValuesRef() bool`

HasValuesRef returns a boolean if a field has been set.

### GetChartName

`func (o *ArtifactSpecArtifact) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *ArtifactSpecArtifact) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *ArtifactSpecArtifact) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartVersion

`func (o *ArtifactSpecArtifact) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *ArtifactSpecArtifact) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *ArtifactSpecArtifact) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetConfigmap

`func (o *ArtifactSpecArtifact) GetConfigmap() File`

GetConfigmap returns the Configmap field if non-nil, zero value otherwise.

### GetConfigmapOk

`func (o *ArtifactSpecArtifact) GetConfigmapOk() (*File, bool)`

GetConfigmapOk returns a tuple with the Configmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigmap

`func (o *ArtifactSpecArtifact) SetConfigmap(v File)`

SetConfigmap sets Configmap field to given value.


### GetConfiguration

`func (o *ArtifactSpecArtifact) GetConfiguration() File`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ArtifactSpecArtifact) GetConfigurationOk() (*File, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ArtifactSpecArtifact) SetConfiguration(v File)`

SetConfiguration sets Configuration field to given value.


### GetSecret

`func (o *ArtifactSpecArtifact) GetSecret() File`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ArtifactSpecArtifact) GetSecretOk() (*File, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ArtifactSpecArtifact) SetSecret(v File)`

SetSecret sets Secret field to given value.


### GetStatefulset

`func (o *ArtifactSpecArtifact) GetStatefulset() File`

GetStatefulset returns the Statefulset field if non-nil, zero value otherwise.

### GetStatefulsetOk

`func (o *ArtifactSpecArtifact) GetStatefulsetOk() (*File, bool)`

GetStatefulsetOk returns a tuple with the Statefulset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatefulset

`func (o *ArtifactSpecArtifact) SetStatefulset(v File)`

SetStatefulset sets Statefulset field to given value.


### GetCatalog

`func (o *ArtifactSpecArtifact) GetCatalog() string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *ArtifactSpecArtifact) GetCatalogOk() (*string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *ArtifactSpecArtifact) SetCatalog(v string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *ArtifactSpecArtifact) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


