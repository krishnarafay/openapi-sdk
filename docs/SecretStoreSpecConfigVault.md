# SecretStoreSpecConfigVault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethod** | Pointer to **string** | Authentication method for secret store resource | [optional] 
**Clusters** | Pointer to [**[]SecretStoreSpecConfigVaultCluster**](SecretStoreSpecConfigVaultCluster.md) | Vault details of the secret store resource | [optional] 
**Host** | Pointer to **string** | Host details | [optional] 

## Methods

### NewSecretStoreSpecConfigVault

`func NewSecretStoreSpecConfigVault() *SecretStoreSpecConfigVault`

NewSecretStoreSpecConfigVault instantiates a new SecretStoreSpecConfigVault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretStoreSpecConfigVaultWithDefaults

`func NewSecretStoreSpecConfigVaultWithDefaults() *SecretStoreSpecConfigVault`

NewSecretStoreSpecConfigVaultWithDefaults instantiates a new SecretStoreSpecConfigVault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMethod

`func (o *SecretStoreSpecConfigVault) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *SecretStoreSpecConfigVault) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *SecretStoreSpecConfigVault) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *SecretStoreSpecConfigVault) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetClusters

`func (o *SecretStoreSpecConfigVault) GetClusters() []SecretStoreSpecConfigVaultCluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *SecretStoreSpecConfigVault) GetClustersOk() (*[]SecretStoreSpecConfigVaultCluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *SecretStoreSpecConfigVault) SetClusters(v []SecretStoreSpecConfigVaultCluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *SecretStoreSpecConfigVault) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetHost

`func (o *SecretStoreSpecConfigVault) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SecretStoreSpecConfigVault) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SecretStoreSpecConfigVault) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SecretStoreSpecConfigVault) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


