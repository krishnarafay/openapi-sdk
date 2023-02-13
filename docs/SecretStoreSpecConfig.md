# SecretStoreSpecConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsiAws** | Pointer to [**SecretStoreSpecConfigCsiAws**](SecretStoreSpecConfigCsiAws.md) |  | [optional] 
**Vault** | Pointer to [**SecretStoreSpecConfigVault**](SecretStoreSpecConfigVault.md) |  | [optional] 

## Methods

### NewSecretStoreSpecConfig

`func NewSecretStoreSpecConfig() *SecretStoreSpecConfig`

NewSecretStoreSpecConfig instantiates a new SecretStoreSpecConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretStoreSpecConfigWithDefaults

`func NewSecretStoreSpecConfigWithDefaults() *SecretStoreSpecConfig`

NewSecretStoreSpecConfigWithDefaults instantiates a new SecretStoreSpecConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsiAws

`func (o *SecretStoreSpecConfig) GetCsiAws() SecretStoreSpecConfigCsiAws`

GetCsiAws returns the CsiAws field if non-nil, zero value otherwise.

### GetCsiAwsOk

`func (o *SecretStoreSpecConfig) GetCsiAwsOk() (*SecretStoreSpecConfigCsiAws, bool)`

GetCsiAwsOk returns a tuple with the CsiAws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsiAws

`func (o *SecretStoreSpecConfig) SetCsiAws(v SecretStoreSpecConfigCsiAws)`

SetCsiAws sets CsiAws field to given value.

### HasCsiAws

`func (o *SecretStoreSpecConfig) HasCsiAws() bool`

HasCsiAws returns a boolean if a field has been set.

### GetVault

`func (o *SecretStoreSpecConfig) GetVault() SecretStoreSpecConfigVault`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *SecretStoreSpecConfig) GetVaultOk() (*SecretStoreSpecConfigVault, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *SecretStoreSpecConfig) SetVault(v SecretStoreSpecConfigVault)`

SetVault sets Vault field to given value.

### HasVault

`func (o *SecretStoreSpecConfig) HasVault() bool`

HasVault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


