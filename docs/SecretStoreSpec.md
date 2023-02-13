# SecretStoreSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**SecretStoreSpecConfig**](SecretStoreSpecConfig.md) |  | [optional] 
**Provider** | Pointer to **string** | repository type | [optional] 

## Methods

### NewSecretStoreSpec

`func NewSecretStoreSpec() *SecretStoreSpec`

NewSecretStoreSpec instantiates a new SecretStoreSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretStoreSpecWithDefaults

`func NewSecretStoreSpecWithDefaults() *SecretStoreSpec`

NewSecretStoreSpecWithDefaults instantiates a new SecretStoreSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SecretStoreSpec) GetConfig() SecretStoreSpecConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SecretStoreSpec) GetConfigOk() (*SecretStoreSpecConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SecretStoreSpec) SetConfig(v SecretStoreSpecConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SecretStoreSpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetProvider

`func (o *SecretStoreSpec) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SecretStoreSpec) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SecretStoreSpec) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SecretStoreSpec) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


