# CredentialsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**CredentialsSpecCredentials**](CredentialsSpecCredentials.md) |  | [optional] 
**Provider** | Pointer to **string** | Provider of Credential Access | [optional] [default to "aws"]
**Secret** | Pointer to [**File**](File.md) |  | [optional] 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Type** | Pointer to **string** | Type of Credentials Access | [optional] [default to "ClusterProvisioning"]

## Methods

### NewCredentialsSpec

`func NewCredentialsSpec() *CredentialsSpec`

NewCredentialsSpec instantiates a new CredentialsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsSpecWithDefaults

`func NewCredentialsSpecWithDefaults() *CredentialsSpec`

NewCredentialsSpecWithDefaults instantiates a new CredentialsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *CredentialsSpec) GetCredentials() CredentialsSpecCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CredentialsSpec) GetCredentialsOk() (*CredentialsSpecCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CredentialsSpec) SetCredentials(v CredentialsSpecCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CredentialsSpec) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetProvider

`func (o *CredentialsSpec) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CredentialsSpec) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CredentialsSpec) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CredentialsSpec) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSecret

`func (o *CredentialsSpec) GetSecret() File`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CredentialsSpec) GetSecretOk() (*File, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CredentialsSpec) SetSecret(v File)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *CredentialsSpec) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSharing

`func (o *CredentialsSpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *CredentialsSpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *CredentialsSpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *CredentialsSpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetType

`func (o *CredentialsSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialsSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialsSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CredentialsSpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


