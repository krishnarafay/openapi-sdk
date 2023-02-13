# DefaultAddons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsiSecretStoreConfig** | Pointer to [**CsiSecretStoreConfig**](CsiSecretStoreConfig.md) |  | [optional] 
**EnableCsiSecretStore** | Pointer to **bool** | enable csi secret store | [optional] 
**EnableIngress** | Pointer to **bool** | enable default ingress | [optional] 
**EnableLogging** | Pointer to **bool** | enable default logging | [optional] 
**EnableMonitoring** | Pointer to **bool** | enable default monitoring | [optional] 
**EnableRookCeph** | Pointer to **bool** | enable rook ceph storage | [optional] 
**EnableVM** | Pointer to **bool** | enable virtual machine workloads | [optional] 
**Monitoring** | Pointer to [**MonitoringConfig**](MonitoringConfig.md) |  | [optional] 

## Methods

### NewDefaultAddons

`func NewDefaultAddons() *DefaultAddons`

NewDefaultAddons instantiates a new DefaultAddons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultAddonsWithDefaults

`func NewDefaultAddonsWithDefaults() *DefaultAddons`

NewDefaultAddonsWithDefaults instantiates a new DefaultAddons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsiSecretStoreConfig

`func (o *DefaultAddons) GetCsiSecretStoreConfig() CsiSecretStoreConfig`

GetCsiSecretStoreConfig returns the CsiSecretStoreConfig field if non-nil, zero value otherwise.

### GetCsiSecretStoreConfigOk

`func (o *DefaultAddons) GetCsiSecretStoreConfigOk() (*CsiSecretStoreConfig, bool)`

GetCsiSecretStoreConfigOk returns a tuple with the CsiSecretStoreConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsiSecretStoreConfig

`func (o *DefaultAddons) SetCsiSecretStoreConfig(v CsiSecretStoreConfig)`

SetCsiSecretStoreConfig sets CsiSecretStoreConfig field to given value.

### HasCsiSecretStoreConfig

`func (o *DefaultAddons) HasCsiSecretStoreConfig() bool`

HasCsiSecretStoreConfig returns a boolean if a field has been set.

### GetEnableCsiSecretStore

`func (o *DefaultAddons) GetEnableCsiSecretStore() bool`

GetEnableCsiSecretStore returns the EnableCsiSecretStore field if non-nil, zero value otherwise.

### GetEnableCsiSecretStoreOk

`func (o *DefaultAddons) GetEnableCsiSecretStoreOk() (*bool, bool)`

GetEnableCsiSecretStoreOk returns a tuple with the EnableCsiSecretStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCsiSecretStore

`func (o *DefaultAddons) SetEnableCsiSecretStore(v bool)`

SetEnableCsiSecretStore sets EnableCsiSecretStore field to given value.

### HasEnableCsiSecretStore

`func (o *DefaultAddons) HasEnableCsiSecretStore() bool`

HasEnableCsiSecretStore returns a boolean if a field has been set.

### GetEnableIngress

`func (o *DefaultAddons) GetEnableIngress() bool`

GetEnableIngress returns the EnableIngress field if non-nil, zero value otherwise.

### GetEnableIngressOk

`func (o *DefaultAddons) GetEnableIngressOk() (*bool, bool)`

GetEnableIngressOk returns a tuple with the EnableIngress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIngress

`func (o *DefaultAddons) SetEnableIngress(v bool)`

SetEnableIngress sets EnableIngress field to given value.

### HasEnableIngress

`func (o *DefaultAddons) HasEnableIngress() bool`

HasEnableIngress returns a boolean if a field has been set.

### GetEnableLogging

`func (o *DefaultAddons) GetEnableLogging() bool`

GetEnableLogging returns the EnableLogging field if non-nil, zero value otherwise.

### GetEnableLoggingOk

`func (o *DefaultAddons) GetEnableLoggingOk() (*bool, bool)`

GetEnableLoggingOk returns a tuple with the EnableLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLogging

`func (o *DefaultAddons) SetEnableLogging(v bool)`

SetEnableLogging sets EnableLogging field to given value.

### HasEnableLogging

`func (o *DefaultAddons) HasEnableLogging() bool`

HasEnableLogging returns a boolean if a field has been set.

### GetEnableMonitoring

`func (o *DefaultAddons) GetEnableMonitoring() bool`

GetEnableMonitoring returns the EnableMonitoring field if non-nil, zero value otherwise.

### GetEnableMonitoringOk

`func (o *DefaultAddons) GetEnableMonitoringOk() (*bool, bool)`

GetEnableMonitoringOk returns a tuple with the EnableMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMonitoring

`func (o *DefaultAddons) SetEnableMonitoring(v bool)`

SetEnableMonitoring sets EnableMonitoring field to given value.

### HasEnableMonitoring

`func (o *DefaultAddons) HasEnableMonitoring() bool`

HasEnableMonitoring returns a boolean if a field has been set.

### GetEnableRookCeph

`func (o *DefaultAddons) GetEnableRookCeph() bool`

GetEnableRookCeph returns the EnableRookCeph field if non-nil, zero value otherwise.

### GetEnableRookCephOk

`func (o *DefaultAddons) GetEnableRookCephOk() (*bool, bool)`

GetEnableRookCephOk returns a tuple with the EnableRookCeph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRookCeph

`func (o *DefaultAddons) SetEnableRookCeph(v bool)`

SetEnableRookCeph sets EnableRookCeph field to given value.

### HasEnableRookCeph

`func (o *DefaultAddons) HasEnableRookCeph() bool`

HasEnableRookCeph returns a boolean if a field has been set.

### GetEnableVM

`func (o *DefaultAddons) GetEnableVM() bool`

GetEnableVM returns the EnableVM field if non-nil, zero value otherwise.

### GetEnableVMOk

`func (o *DefaultAddons) GetEnableVMOk() (*bool, bool)`

GetEnableVMOk returns a tuple with the EnableVM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVM

`func (o *DefaultAddons) SetEnableVM(v bool)`

SetEnableVM sets EnableVM field to given value.

### HasEnableVM

`func (o *DefaultAddons) HasEnableVM() bool`

HasEnableVM returns a boolean if a field has been set.

### GetMonitoring

`func (o *DefaultAddons) GetMonitoring() MonitoringConfig`

GetMonitoring returns the Monitoring field if non-nil, zero value otherwise.

### GetMonitoringOk

`func (o *DefaultAddons) GetMonitoringOk() (*MonitoringConfig, bool)`

GetMonitoringOk returns a tuple with the Monitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoring

`func (o *DefaultAddons) SetMonitoring(v MonitoringConfig)`

SetMonitoring sets Monitoring field to given value.

### HasMonitoring

`func (o *DefaultAddons) HasMonitoring() bool`

HasMonitoring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


