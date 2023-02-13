# SharingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | flag to specify if sharing is enabled for resource | 
**Projects** | Pointer to [**[]ProjectMeta**](ProjectMeta.md) | list of projects this resource is shared to | [optional] 

## Methods

### NewSharingSpec

`func NewSharingSpec(enabled bool, ) *SharingSpec`

NewSharingSpec instantiates a new SharingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharingSpecWithDefaults

`func NewSharingSpecWithDefaults() *SharingSpec`

NewSharingSpecWithDefaults instantiates a new SharingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SharingSpec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SharingSpec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SharingSpec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProjects

`func (o *SharingSpec) GetProjects() []ProjectMeta`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *SharingSpec) GetProjectsOk() (*[]ProjectMeta, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *SharingSpec) SetProjects(v []ProjectMeta)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *SharingSpec) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


