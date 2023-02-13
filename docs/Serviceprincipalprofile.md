# Serviceprincipalprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The ID for the service principal. If specified, must be set to [parameters(&#39;servicePrincipalClientId&#39;)]. This would be set to the cloud credential&#39;s client ID during cluster deployment. | [optional] 
**Secret** | Pointer to **string** | The secret password associated with the service principal. If specified, must be set to [parameters(&#39;servicePrincipalClientSecret&#39;)]. This would be set to the cloud credential&#39;s client secret during cluster deployment. | [optional] 

## Methods

### NewServiceprincipalprofile

`func NewServiceprincipalprofile() *Serviceprincipalprofile`

NewServiceprincipalprofile instantiates a new Serviceprincipalprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceprincipalprofileWithDefaults

`func NewServiceprincipalprofileWithDefaults() *Serviceprincipalprofile`

NewServiceprincipalprofileWithDefaults instantiates a new Serviceprincipalprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *Serviceprincipalprofile) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Serviceprincipalprofile) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Serviceprincipalprofile) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Serviceprincipalprofile) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *Serviceprincipalprofile) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Serviceprincipalprofile) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Serviceprincipalprofile) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *Serviceprincipalprofile) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


