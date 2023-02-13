# Securityprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureKeyVaultKms** | Pointer to [**Azurekeyvaultkms**](Azurekeyvaultkms.md) |  | [optional] 
**Defender** | Pointer to [**Defender**](Defender.md) |  | [optional] 

## Methods

### NewSecurityprofile

`func NewSecurityprofile() *Securityprofile`

NewSecurityprofile instantiates a new Securityprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityprofileWithDefaults

`func NewSecurityprofileWithDefaults() *Securityprofile`

NewSecurityprofileWithDefaults instantiates a new Securityprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureKeyVaultKms

`func (o *Securityprofile) GetAzureKeyVaultKms() Azurekeyvaultkms`

GetAzureKeyVaultKms returns the AzureKeyVaultKms field if non-nil, zero value otherwise.

### GetAzureKeyVaultKmsOk

`func (o *Securityprofile) GetAzureKeyVaultKmsOk() (*Azurekeyvaultkms, bool)`

GetAzureKeyVaultKmsOk returns a tuple with the AzureKeyVaultKms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureKeyVaultKms

`func (o *Securityprofile) SetAzureKeyVaultKms(v Azurekeyvaultkms)`

SetAzureKeyVaultKms sets AzureKeyVaultKms field to given value.

### HasAzureKeyVaultKms

`func (o *Securityprofile) HasAzureKeyVaultKms() bool`

HasAzureKeyVaultKms returns a boolean if a field has been set.

### GetDefender

`func (o *Securityprofile) GetDefender() Defender`

GetDefender returns the Defender field if non-nil, zero value otherwise.

### GetDefenderOk

`func (o *Securityprofile) GetDefenderOk() (*Defender, bool)`

GetDefenderOk returns a tuple with the Defender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefender

`func (o *Securityprofile) SetDefender(v Defender)`

SetDefender sets Defender field to given value.

### HasDefender

`func (o *Securityprofile) HasDefender() bool`

HasDefender returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


