# CsiSecretStoreConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableSecretRotation** | Pointer to **bool** |  | [optional] 
**Providers** | Pointer to [**SecretStoreProvider**](SecretStoreProvider.md) |  | [optional] 
**RotationPollInterval** | Pointer to **string** |  | [optional] 
**SyncSecrets** | Pointer to **bool** |  | [optional] 

## Methods

### NewCsiSecretStoreConfig

`func NewCsiSecretStoreConfig() *CsiSecretStoreConfig`

NewCsiSecretStoreConfig instantiates a new CsiSecretStoreConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCsiSecretStoreConfigWithDefaults

`func NewCsiSecretStoreConfigWithDefaults() *CsiSecretStoreConfig`

NewCsiSecretStoreConfigWithDefaults instantiates a new CsiSecretStoreConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableSecretRotation

`func (o *CsiSecretStoreConfig) GetEnableSecretRotation() bool`

GetEnableSecretRotation returns the EnableSecretRotation field if non-nil, zero value otherwise.

### GetEnableSecretRotationOk

`func (o *CsiSecretStoreConfig) GetEnableSecretRotationOk() (*bool, bool)`

GetEnableSecretRotationOk returns a tuple with the EnableSecretRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSecretRotation

`func (o *CsiSecretStoreConfig) SetEnableSecretRotation(v bool)`

SetEnableSecretRotation sets EnableSecretRotation field to given value.

### HasEnableSecretRotation

`func (o *CsiSecretStoreConfig) HasEnableSecretRotation() bool`

HasEnableSecretRotation returns a boolean if a field has been set.

### GetProviders

`func (o *CsiSecretStoreConfig) GetProviders() SecretStoreProvider`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CsiSecretStoreConfig) GetProvidersOk() (*SecretStoreProvider, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CsiSecretStoreConfig) SetProviders(v SecretStoreProvider)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *CsiSecretStoreConfig) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetRotationPollInterval

`func (o *CsiSecretStoreConfig) GetRotationPollInterval() string`

GetRotationPollInterval returns the RotationPollInterval field if non-nil, zero value otherwise.

### GetRotationPollIntervalOk

`func (o *CsiSecretStoreConfig) GetRotationPollIntervalOk() (*string, bool)`

GetRotationPollIntervalOk returns a tuple with the RotationPollInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPollInterval

`func (o *CsiSecretStoreConfig) SetRotationPollInterval(v string)`

SetRotationPollInterval sets RotationPollInterval field to given value.

### HasRotationPollInterval

`func (o *CsiSecretStoreConfig) HasRotationPollInterval() bool`

HasRotationPollInterval returns a boolean if a field has been set.

### GetSyncSecrets

`func (o *CsiSecretStoreConfig) GetSyncSecrets() bool`

GetSyncSecrets returns the SyncSecrets field if non-nil, zero value otherwise.

### GetSyncSecretsOk

`func (o *CsiSecretStoreConfig) GetSyncSecretsOk() (*bool, bool)`

GetSyncSecretsOk returns a tuple with the SyncSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSecrets

`func (o *CsiSecretStoreConfig) SetSyncSecrets(v bool)`

SetSyncSecrets sets SyncSecrets field to given value.

### HasSyncSecrets

`func (o *CsiSecretStoreConfig) HasSyncSecrets() bool`

HasSyncSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


