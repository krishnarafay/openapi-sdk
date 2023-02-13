# RepositorySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]AgentMeta**](AgentMeta.md) | repository agents | [optional] 
**Credentials** | Pointer to [**RepositorySpecCredentials**](RepositorySpecCredentials.md) |  | [optional] 
**Endpoint** | Pointer to **string** | repository endpoint | [optional] 
**Options** | Pointer to [**RepositoryOptions**](RepositoryOptions.md) |  | [optional] 
**Secret** | Pointer to [**File**](File.md) |  | [optional] 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Type** | Pointer to **string** | repository type | [optional] 

## Methods

### NewRepositorySpec

`func NewRepositorySpec() *RepositorySpec`

NewRepositorySpec instantiates a new RepositorySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositorySpecWithDefaults

`func NewRepositorySpecWithDefaults() *RepositorySpec`

NewRepositorySpecWithDefaults instantiates a new RepositorySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *RepositorySpec) GetAgents() []AgentMeta`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *RepositorySpec) GetAgentsOk() (*[]AgentMeta, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *RepositorySpec) SetAgents(v []AgentMeta)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *RepositorySpec) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetCredentials

`func (o *RepositorySpec) GetCredentials() RepositorySpecCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *RepositorySpec) GetCredentialsOk() (*RepositorySpecCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *RepositorySpec) SetCredentials(v RepositorySpecCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *RepositorySpec) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEndpoint

`func (o *RepositorySpec) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *RepositorySpec) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *RepositorySpec) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *RepositorySpec) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetOptions

`func (o *RepositorySpec) GetOptions() RepositoryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *RepositorySpec) GetOptionsOk() (*RepositoryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *RepositorySpec) SetOptions(v RepositoryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *RepositorySpec) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetSecret

`func (o *RepositorySpec) GetSecret() File`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *RepositorySpec) GetSecretOk() (*File, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *RepositorySpec) SetSecret(v File)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *RepositorySpec) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSharing

`func (o *RepositorySpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *RepositorySpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *RepositorySpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *RepositorySpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetType

`func (o *RepositorySpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RepositorySpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RepositorySpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RepositorySpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


