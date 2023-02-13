# OverrideSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Placement** | Pointer to [**PlacementSpec**](PlacementSpec.md) |  | [optional] 
**Repo** | Pointer to [**RepoOverride**](RepoOverride.md) |  | [optional] 
**Resource** | Pointer to [**OverrideResource**](OverrideResource.md) |  | [optional] 
**Type** | Pointer to **string** | Override Type  | [optional] 
**ValuesPath** | Pointer to [**File**](File.md) |  | [optional] 

## Methods

### NewOverrideSpec

`func NewOverrideSpec() *OverrideSpec`

NewOverrideSpec instantiates a new OverrideSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideSpecWithDefaults

`func NewOverrideSpecWithDefaults() *OverrideSpec`

NewOverrideSpecWithDefaults instantiates a new OverrideSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlacement

`func (o *OverrideSpec) GetPlacement() PlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *OverrideSpec) GetPlacementOk() (*PlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *OverrideSpec) SetPlacement(v PlacementSpec)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *OverrideSpec) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetRepo

`func (o *OverrideSpec) GetRepo() RepoOverride`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *OverrideSpec) GetRepoOk() (*RepoOverride, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *OverrideSpec) SetRepo(v RepoOverride)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *OverrideSpec) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetResource

`func (o *OverrideSpec) GetResource() OverrideResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OverrideSpec) GetResourceOk() (*OverrideResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OverrideSpec) SetResource(v OverrideResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OverrideSpec) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetType

`func (o *OverrideSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OverrideSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OverrideSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OverrideSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValuesPath

`func (o *OverrideSpec) GetValuesPath() File`

GetValuesPath returns the ValuesPath field if non-nil, zero value otherwise.

### GetValuesPathOk

`func (o *OverrideSpec) GetValuesPathOk() (*File, bool)`

GetValuesPathOk returns a tuple with the ValuesPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesPath

`func (o *OverrideSpec) SetValuesPath(v File)`

SetValuesPath sets ValuesPath field to given value.

### HasValuesPath

`func (o *OverrideSpec) HasValuesPath() bool`

HasValuesPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


