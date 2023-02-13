# SystemComponentsPlacement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaemonSetOverride** | Pointer to [**DaemonSetOverride**](DaemonSetOverride.md) |  | [optional] 
**NodeSelector** | Pointer to **map[string]string** |  | [optional] 
**Tolerations** | Pointer to [**[]DaemonSetOverrideTolerationsInner**](DaemonSetOverrideTolerationsInner.md) |  | [optional] 

## Methods

### NewSystemComponentsPlacement

`func NewSystemComponentsPlacement() *SystemComponentsPlacement`

NewSystemComponentsPlacement instantiates a new SystemComponentsPlacement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemComponentsPlacementWithDefaults

`func NewSystemComponentsPlacementWithDefaults() *SystemComponentsPlacement`

NewSystemComponentsPlacementWithDefaults instantiates a new SystemComponentsPlacement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaemonSetOverride

`func (o *SystemComponentsPlacement) GetDaemonSetOverride() DaemonSetOverride`

GetDaemonSetOverride returns the DaemonSetOverride field if non-nil, zero value otherwise.

### GetDaemonSetOverrideOk

`func (o *SystemComponentsPlacement) GetDaemonSetOverrideOk() (*DaemonSetOverride, bool)`

GetDaemonSetOverrideOk returns a tuple with the DaemonSetOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonSetOverride

`func (o *SystemComponentsPlacement) SetDaemonSetOverride(v DaemonSetOverride)`

SetDaemonSetOverride sets DaemonSetOverride field to given value.

### HasDaemonSetOverride

`func (o *SystemComponentsPlacement) HasDaemonSetOverride() bool`

HasDaemonSetOverride returns a boolean if a field has been set.

### GetNodeSelector

`func (o *SystemComponentsPlacement) GetNodeSelector() map[string]string`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *SystemComponentsPlacement) GetNodeSelectorOk() (*map[string]string, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *SystemComponentsPlacement) SetNodeSelector(v map[string]string)`

SetNodeSelector sets NodeSelector field to given value.

### HasNodeSelector

`func (o *SystemComponentsPlacement) HasNodeSelector() bool`

HasNodeSelector returns a boolean if a field has been set.

### GetTolerations

`func (o *SystemComponentsPlacement) GetTolerations() []DaemonSetOverrideTolerationsInner`

GetTolerations returns the Tolerations field if non-nil, zero value otherwise.

### GetTolerationsOk

`func (o *SystemComponentsPlacement) GetTolerationsOk() (*[]DaemonSetOverrideTolerationsInner, bool)`

GetTolerationsOk returns a tuple with the Tolerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerations

`func (o *SystemComponentsPlacement) SetTolerations(v []DaemonSetOverrideTolerationsInner)`

SetTolerations sets Tolerations field to given value.

### HasTolerations

`func (o *SystemComponentsPlacement) HasTolerations() bool`

HasTolerations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


