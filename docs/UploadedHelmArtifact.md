# UploadedHelmArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartPath** | [**File**](File.md) |  | 
**ValuesPaths** | Pointer to [**[]File**](File.md) | represents a Helm Values by their relative path | [optional] 

## Methods

### NewUploadedHelmArtifact

`func NewUploadedHelmArtifact(chartPath File, ) *UploadedHelmArtifact`

NewUploadedHelmArtifact instantiates a new UploadedHelmArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadedHelmArtifactWithDefaults

`func NewUploadedHelmArtifactWithDefaults() *UploadedHelmArtifact`

NewUploadedHelmArtifactWithDefaults instantiates a new UploadedHelmArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartPath

`func (o *UploadedHelmArtifact) GetChartPath() File`

GetChartPath returns the ChartPath field if non-nil, zero value otherwise.

### GetChartPathOk

`func (o *UploadedHelmArtifact) GetChartPathOk() (*File, bool)`

GetChartPathOk returns a tuple with the ChartPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartPath

`func (o *UploadedHelmArtifact) SetChartPath(v File)`

SetChartPath sets ChartPath field to given value.


### GetValuesPaths

`func (o *UploadedHelmArtifact) GetValuesPaths() []File`

GetValuesPaths returns the ValuesPaths field if non-nil, zero value otherwise.

### GetValuesPathsOk

`func (o *UploadedHelmArtifact) GetValuesPathsOk() (*[]File, bool)`

GetValuesPathsOk returns a tuple with the ValuesPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesPaths

`func (o *UploadedHelmArtifact) SetValuesPaths(v []File)`

SetValuesPaths sets ValuesPaths field to given value.

### HasValuesPaths

`func (o *UploadedHelmArtifact) HasValuesPaths() bool`

HasValuesPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


