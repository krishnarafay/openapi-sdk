# ExtraStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaultIntegration** | Pointer to [**VaultClusterStatus**](VaultClusterStatus.md) |  | [optional] 

## Methods

### NewExtraStatus

`func NewExtraStatus() *ExtraStatus`

NewExtraStatus instantiates a new ExtraStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtraStatusWithDefaults

`func NewExtraStatusWithDefaults() *ExtraStatus`

NewExtraStatusWithDefaults instantiates a new ExtraStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVaultIntegration

`func (o *ExtraStatus) GetVaultIntegration() VaultClusterStatus`

GetVaultIntegration returns the VaultIntegration field if non-nil, zero value otherwise.

### GetVaultIntegrationOk

`func (o *ExtraStatus) GetVaultIntegrationOk() (*VaultClusterStatus, bool)`

GetVaultIntegrationOk returns a tuple with the VaultIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIntegration

`func (o *ExtraStatus) SetVaultIntegration(v VaultClusterStatus)`

SetVaultIntegration sets VaultIntegration field to given value.

### HasVaultIntegration

`func (o *ExtraStatus) HasVaultIntegration() bool`

HasVaultIntegration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


