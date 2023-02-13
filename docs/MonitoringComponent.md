# MonitoringComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discovery** | Pointer to [**MonitoringDiscoveryConfig**](MonitoringDiscoveryConfig.md) |  | [optional] 
**Enabled** | **bool** | flag to specify if monitoring component is enabled | 

## Methods

### NewMonitoringComponent

`func NewMonitoringComponent(enabled bool, ) *MonitoringComponent`

NewMonitoringComponent instantiates a new MonitoringComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringComponentWithDefaults

`func NewMonitoringComponentWithDefaults() *MonitoringComponent`

NewMonitoringComponentWithDefaults instantiates a new MonitoringComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscovery

`func (o *MonitoringComponent) GetDiscovery() MonitoringDiscoveryConfig`

GetDiscovery returns the Discovery field if non-nil, zero value otherwise.

### GetDiscoveryOk

`func (o *MonitoringComponent) GetDiscoveryOk() (*MonitoringDiscoveryConfig, bool)`

GetDiscoveryOk returns a tuple with the Discovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovery

`func (o *MonitoringComponent) SetDiscovery(v MonitoringDiscoveryConfig)`

SetDiscovery sets Discovery field to given value.

### HasDiscovery

`func (o *MonitoringComponent) HasDiscovery() bool`

HasDiscovery returns a boolean if a field has been set.

### GetEnabled

`func (o *MonitoringComponent) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MonitoringComponent) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MonitoringComponent) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


