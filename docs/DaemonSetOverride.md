# DaemonSetOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeSelectionEnabled** | Pointer to **bool** |  | [optional] 
**Tolerations** | Pointer to [**[]DaemonSetOverrideTolerationsInner**](DaemonSetOverrideTolerationsInner.md) |  | [optional] 

## Methods

### NewDaemonSetOverride

`func NewDaemonSetOverride() *DaemonSetOverride`

NewDaemonSetOverride instantiates a new DaemonSetOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaemonSetOverrideWithDefaults

`func NewDaemonSetOverrideWithDefaults() *DaemonSetOverride`

NewDaemonSetOverrideWithDefaults instantiates a new DaemonSetOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeSelectionEnabled

`func (o *DaemonSetOverride) GetNodeSelectionEnabled() bool`

GetNodeSelectionEnabled returns the NodeSelectionEnabled field if non-nil, zero value otherwise.

### GetNodeSelectionEnabledOk

`func (o *DaemonSetOverride) GetNodeSelectionEnabledOk() (*bool, bool)`

GetNodeSelectionEnabledOk returns a tuple with the NodeSelectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelectionEnabled

`func (o *DaemonSetOverride) SetNodeSelectionEnabled(v bool)`

SetNodeSelectionEnabled sets NodeSelectionEnabled field to given value.

### HasNodeSelectionEnabled

`func (o *DaemonSetOverride) HasNodeSelectionEnabled() bool`

HasNodeSelectionEnabled returns a boolean if a field has been set.

### GetTolerations

`func (o *DaemonSetOverride) GetTolerations() []DaemonSetOverrideTolerationsInner`

GetTolerations returns the Tolerations field if non-nil, zero value otherwise.

### GetTolerationsOk

`func (o *DaemonSetOverride) GetTolerationsOk() (*[]DaemonSetOverrideTolerationsInner, bool)`

GetTolerationsOk returns a tuple with the Tolerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerations

`func (o *DaemonSetOverride) SetTolerations(v []DaemonSetOverrideTolerationsInner)`

SetTolerations sets Tolerations field to given value.

### HasTolerations

`func (o *DaemonSetOverride) HasTolerations() bool`

HasTolerations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


