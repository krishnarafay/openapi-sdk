# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionStatus** | Pointer to **string** | status of the condition | [optional] [readonly] 
**ConditionType** | Pointer to **string** | type of the status condition | [optional] [readonly] 
**Extra** | Pointer to [**Extra**](Extra.md) |  | [optional] 
**LastUpdated** | Pointer to [**LastUpdated**](LastUpdated.md) |  | [optional] 
**Reason** | Pointer to **string** | reason of the last condition status | [optional] [readonly] 

## Methods

### NewStatus

`func NewStatus() *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionStatus

`func (o *Status) GetConditionStatus() string`

GetConditionStatus returns the ConditionStatus field if non-nil, zero value otherwise.

### GetConditionStatusOk

`func (o *Status) GetConditionStatusOk() (*string, bool)`

GetConditionStatusOk returns a tuple with the ConditionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionStatus

`func (o *Status) SetConditionStatus(v string)`

SetConditionStatus sets ConditionStatus field to given value.

### HasConditionStatus

`func (o *Status) HasConditionStatus() bool`

HasConditionStatus returns a boolean if a field has been set.

### GetConditionType

`func (o *Status) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *Status) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *Status) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.

### HasConditionType

`func (o *Status) HasConditionType() bool`

HasConditionType returns a boolean if a field has been set.

### GetExtra

`func (o *Status) GetExtra() Extra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *Status) GetExtraOk() (*Extra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *Status) SetExtra(v Extra)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *Status) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Status) GetLastUpdated() LastUpdated`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Status) GetLastUpdatedOk() (*LastUpdated, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Status) SetLastUpdated(v LastUpdated)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Status) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetReason

`func (o *Status) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Status) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Status) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Status) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


