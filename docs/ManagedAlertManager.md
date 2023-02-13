# ManagedAlertManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configmap** | [**File**](File.md) |  | 
**Configuration** | [**File**](File.md) |  | 
**Secret** | [**File**](File.md) |  | 
**Statefulset** | [**File**](File.md) |  | 

## Methods

### NewManagedAlertManager

`func NewManagedAlertManager(configmap File, configuration File, secret File, statefulset File, ) *ManagedAlertManager`

NewManagedAlertManager instantiates a new ManagedAlertManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAlertManagerWithDefaults

`func NewManagedAlertManagerWithDefaults() *ManagedAlertManager`

NewManagedAlertManagerWithDefaults instantiates a new ManagedAlertManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigmap

`func (o *ManagedAlertManager) GetConfigmap() File`

GetConfigmap returns the Configmap field if non-nil, zero value otherwise.

### GetConfigmapOk

`func (o *ManagedAlertManager) GetConfigmapOk() (*File, bool)`

GetConfigmapOk returns a tuple with the Configmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigmap

`func (o *ManagedAlertManager) SetConfigmap(v File)`

SetConfigmap sets Configmap field to given value.


### GetConfiguration

`func (o *ManagedAlertManager) GetConfiguration() File`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ManagedAlertManager) GetConfigurationOk() (*File, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ManagedAlertManager) SetConfiguration(v File)`

SetConfiguration sets Configuration field to given value.


### GetSecret

`func (o *ManagedAlertManager) GetSecret() File`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ManagedAlertManager) GetSecretOk() (*File, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ManagedAlertManager) SetSecret(v File)`

SetSecret sets Secret field to given value.


### GetStatefulset

`func (o *ManagedAlertManager) GetStatefulset() File`

GetStatefulset returns the Statefulset field if non-nil, zero value otherwise.

### GetStatefulsetOk

`func (o *ManagedAlertManager) GetStatefulsetOk() (*File, bool)`

GetStatefulsetOk returns a tuple with the Statefulset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatefulset

`func (o *ManagedAlertManager) SetStatefulset(v File)`

SetStatefulset sets Statefulset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


