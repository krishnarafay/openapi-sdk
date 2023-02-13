# VaultClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuxiliaryTask** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**SecretStore** | Pointer to **string** |  | [optional] 

## Methods

### NewVaultClusterStatus

`func NewVaultClusterStatus() *VaultClusterStatus`

NewVaultClusterStatus instantiates a new VaultClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultClusterStatusWithDefaults

`func NewVaultClusterStatusWithDefaults() *VaultClusterStatus`

NewVaultClusterStatusWithDefaults instantiates a new VaultClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuxiliaryTask

`func (o *VaultClusterStatus) GetAuxiliaryTask() string`

GetAuxiliaryTask returns the AuxiliaryTask field if non-nil, zero value otherwise.

### GetAuxiliaryTaskOk

`func (o *VaultClusterStatus) GetAuxiliaryTaskOk() (*string, bool)`

GetAuxiliaryTaskOk returns a tuple with the AuxiliaryTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryTask

`func (o *VaultClusterStatus) SetAuxiliaryTask(v string)`

SetAuxiliaryTask sets AuxiliaryTask field to given value.

### HasAuxiliaryTask

`func (o *VaultClusterStatus) HasAuxiliaryTask() bool`

HasAuxiliaryTask returns a boolean if a field has been set.

### GetEnabled

`func (o *VaultClusterStatus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VaultClusterStatus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VaultClusterStatus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VaultClusterStatus) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSecretStore

`func (o *VaultClusterStatus) GetSecretStore() string`

GetSecretStore returns the SecretStore field if non-nil, zero value otherwise.

### GetSecretStoreOk

`func (o *VaultClusterStatus) GetSecretStoreOk() (*string, bool)`

GetSecretStoreOk returns a tuple with the SecretStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretStore

`func (o *VaultClusterStatus) SetSecretStore(v string)`

SetSecretStore sets SecretStore field to given value.

### HasSecretStore

`func (o *VaultClusterStatus) HasSecretStore() bool`

HasSecretStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


