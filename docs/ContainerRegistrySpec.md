# ContainerRegistrySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**ContainerRegistrySpecCredentials**](ContainerRegistrySpecCredentials.md) |  | [optional] 
**Endpoint** | Pointer to **string** | Container Registry endpoint | [optional] 
**Provider** | Pointer to **string** | container registry provider | [optional] 
**Secret** | Pointer to [**File**](File.md) |  | [optional] 

## Methods

### NewContainerRegistrySpec

`func NewContainerRegistrySpec() *ContainerRegistrySpec`

NewContainerRegistrySpec instantiates a new ContainerRegistrySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRegistrySpecWithDefaults

`func NewContainerRegistrySpecWithDefaults() *ContainerRegistrySpec`

NewContainerRegistrySpecWithDefaults instantiates a new ContainerRegistrySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ContainerRegistrySpec) GetCredentials() ContainerRegistrySpecCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ContainerRegistrySpec) GetCredentialsOk() (*ContainerRegistrySpecCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ContainerRegistrySpec) SetCredentials(v ContainerRegistrySpecCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ContainerRegistrySpec) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEndpoint

`func (o *ContainerRegistrySpec) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ContainerRegistrySpec) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ContainerRegistrySpec) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ContainerRegistrySpec) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetProvider

`func (o *ContainerRegistrySpec) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ContainerRegistrySpec) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ContainerRegistrySpec) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ContainerRegistrySpec) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSecret

`func (o *ContainerRegistrySpec) GetSecret() File`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ContainerRegistrySpec) GetSecretOk() (*File, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ContainerRegistrySpec) SetSecret(v File)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ContainerRegistrySpec) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


