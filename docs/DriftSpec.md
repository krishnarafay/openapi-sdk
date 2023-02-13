# DriftSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | flag to specify if sharing is enabled for resource | [optional] 
**Enabled** | **bool** | flag to specify if drift reconcillation is enabled for resource | 

## Methods

### NewDriftSpec

`func NewDriftSpec(enabled bool, ) *DriftSpec`

NewDriftSpec instantiates a new DriftSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriftSpecWithDefaults

`func NewDriftSpecWithDefaults() *DriftSpec`

NewDriftSpecWithDefaults instantiates a new DriftSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DriftSpec) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DriftSpec) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DriftSpec) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DriftSpec) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetEnabled

`func (o *DriftSpec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DriftSpec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DriftSpec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


