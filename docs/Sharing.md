# Sharing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Projects** | Pointer to [**[]Projects**](Projects.md) |  | [optional] 

## Methods

### NewSharing

`func NewSharing() *Sharing`

NewSharing instantiates a new Sharing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharingWithDefaults

`func NewSharingWithDefaults() *Sharing`

NewSharingWithDefaults instantiates a new Sharing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *Sharing) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Sharing) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Sharing) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Sharing) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProjects

`func (o *Sharing) GetProjects() []Projects`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *Sharing) GetProjectsOk() (*[]Projects, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *Sharing) SetProjects(v []Projects)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *Sharing) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


