# SecretGroupSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to [**File**](File.md) |  | [optional] 
**Secrets** | [**[]Secret**](Secret.md) |  | 

## Methods

### NewSecretGroupSpec

`func NewSecretGroupSpec(secrets []Secret, ) *SecretGroupSpec`

NewSecretGroupSpec instantiates a new SecretGroupSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretGroupSpecWithDefaults

`func NewSecretGroupSpecWithDefaults() *SecretGroupSpec`

NewSecretGroupSpecWithDefaults instantiates a new SecretGroupSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *SecretGroupSpec) GetSecret() File`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SecretGroupSpec) GetSecretOk() (*File, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SecretGroupSpec) SetSecret(v File)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SecretGroupSpec) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSecrets

`func (o *SecretGroupSpec) GetSecrets() []Secret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *SecretGroupSpec) GetSecretsOk() (*[]Secret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *SecretGroupSpec) SetSecrets(v []Secret)`

SetSecrets sets Secrets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


