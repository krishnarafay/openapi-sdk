# MonitoringDiscoveryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **map[string]string** | labels of the monitoring component | [optional] 
**Namespace** | **string** | namespace of the monitoring component | 
**Resource** | **string** | resource name of the monitoring component | 

## Methods

### NewMonitoringDiscoveryConfig

`func NewMonitoringDiscoveryConfig(namespace string, resource string, ) *MonitoringDiscoveryConfig`

NewMonitoringDiscoveryConfig instantiates a new MonitoringDiscoveryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringDiscoveryConfigWithDefaults

`func NewMonitoringDiscoveryConfigWithDefaults() *MonitoringDiscoveryConfig`

NewMonitoringDiscoveryConfigWithDefaults instantiates a new MonitoringDiscoveryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *MonitoringDiscoveryConfig) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MonitoringDiscoveryConfig) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MonitoringDiscoveryConfig) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MonitoringDiscoveryConfig) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetNamespace

`func (o *MonitoringDiscoveryConfig) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *MonitoringDiscoveryConfig) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *MonitoringDiscoveryConfig) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetResource

`func (o *MonitoringDiscoveryConfig) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MonitoringDiscoveryConfig) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MonitoringDiscoveryConfig) SetResource(v string)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


