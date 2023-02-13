# SecretStoreProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to **bool** |  | [optional] 
**Vault** | Pointer to **bool** |  | [optional] 

## Methods

### NewSecretStoreProvider

`func NewSecretStoreProvider() *SecretStoreProvider`

NewSecretStoreProvider instantiates a new SecretStoreProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretStoreProviderWithDefaults

`func NewSecretStoreProviderWithDefaults() *SecretStoreProvider`

NewSecretStoreProviderWithDefaults instantiates a new SecretStoreProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *SecretStoreProvider) GetAws() bool`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *SecretStoreProvider) GetAwsOk() (*bool, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *SecretStoreProvider) SetAws(v bool)`

SetAws sets Aws field to given value.

### HasAws

`func (o *SecretStoreProvider) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetVault

`func (o *SecretStoreProvider) GetVault() bool`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *SecretStoreProvider) GetVaultOk() (*bool, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *SecretStoreProvider) SetVault(v bool)`

SetVault sets Vault field to given value.

### HasVault

`func (o *SecretStoreProvider) HasVault() bool`

HasVault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


