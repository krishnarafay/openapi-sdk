# Azurekeyvaultkms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether to enable Azure Key Vault key management service. The default is false. | [optional] 
**KeyId** | Pointer to **string** | Identifier of Azure Key Vault key. See key identifier format for more details. When Azure Key Vault key management service is enabled, this field is required and must be a valid key identifier. When Azure Key Vault key management service is disabled, leave the field empty. | [optional] 
**KeyVaultNetworkAccess** | Pointer to **string** | Network access of key vault. The possible values are Public and Private. Public means the key vault allows public access from all networks. Private means the key vault disables public access and enables private link. The default value is Public. | [optional] 
**KeyVaultResourceId** | Pointer to **string** | Resource ID of key vault. When keyVaultNetworkAccess is Private, this field is required and must be a valid resource ID. When keyVaultNetworkAccess is Public, leave the field empty. | [optional] 

## Methods

### NewAzurekeyvaultkms

`func NewAzurekeyvaultkms() *Azurekeyvaultkms`

NewAzurekeyvaultkms instantiates a new Azurekeyvaultkms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzurekeyvaultkmsWithDefaults

`func NewAzurekeyvaultkmsWithDefaults() *Azurekeyvaultkms`

NewAzurekeyvaultkmsWithDefaults instantiates a new Azurekeyvaultkms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *Azurekeyvaultkms) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Azurekeyvaultkms) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Azurekeyvaultkms) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Azurekeyvaultkms) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetKeyId

`func (o *Azurekeyvaultkms) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *Azurekeyvaultkms) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *Azurekeyvaultkms) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *Azurekeyvaultkms) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetKeyVaultNetworkAccess

`func (o *Azurekeyvaultkms) GetKeyVaultNetworkAccess() string`

GetKeyVaultNetworkAccess returns the KeyVaultNetworkAccess field if non-nil, zero value otherwise.

### GetKeyVaultNetworkAccessOk

`func (o *Azurekeyvaultkms) GetKeyVaultNetworkAccessOk() (*string, bool)`

GetKeyVaultNetworkAccessOk returns a tuple with the KeyVaultNetworkAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVaultNetworkAccess

`func (o *Azurekeyvaultkms) SetKeyVaultNetworkAccess(v string)`

SetKeyVaultNetworkAccess sets KeyVaultNetworkAccess field to given value.

### HasKeyVaultNetworkAccess

`func (o *Azurekeyvaultkms) HasKeyVaultNetworkAccess() bool`

HasKeyVaultNetworkAccess returns a boolean if a field has been set.

### GetKeyVaultResourceId

`func (o *Azurekeyvaultkms) GetKeyVaultResourceId() string`

GetKeyVaultResourceId returns the KeyVaultResourceId field if non-nil, zero value otherwise.

### GetKeyVaultResourceIdOk

`func (o *Azurekeyvaultkms) GetKeyVaultResourceIdOk() (*string, bool)`

GetKeyVaultResourceIdOk returns a tuple with the KeyVaultResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVaultResourceId

`func (o *Azurekeyvaultkms) SetKeyVaultResourceId(v string)`

SetKeyVaultResourceId sets KeyVaultResourceId field to given value.

### HasKeyVaultResourceId

`func (o *Azurekeyvaultkms) HasKeyVaultResourceId() bool`

HasKeyVaultResourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


