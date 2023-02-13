# ContainerRegistrySpecCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**AccessKeyID** | Pointer to **string** |  | [optional] 
**AccessSecretKey** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**JsonKeyData** | Pointer to **string** |  | [optional] 
**ServicePrincipalID** | Pointer to **string** |  | [optional] 
**ServicePrincipalPassword** | Pointer to **string** |  | [optional] 

## Methods

### NewContainerRegistrySpecCredentials

`func NewContainerRegistrySpecCredentials() *ContainerRegistrySpecCredentials`

NewContainerRegistrySpecCredentials instantiates a new ContainerRegistrySpecCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRegistrySpecCredentialsWithDefaults

`func NewContainerRegistrySpecCredentialsWithDefaults() *ContainerRegistrySpecCredentials`

NewContainerRegistrySpecCredentialsWithDefaults instantiates a new ContainerRegistrySpecCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *ContainerRegistrySpecCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ContainerRegistrySpecCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ContainerRegistrySpecCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ContainerRegistrySpecCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *ContainerRegistrySpecCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ContainerRegistrySpecCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ContainerRegistrySpecCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ContainerRegistrySpecCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAccessKeyID

`func (o *ContainerRegistrySpecCredentials) GetAccessKeyID() string`

GetAccessKeyID returns the AccessKeyID field if non-nil, zero value otherwise.

### GetAccessKeyIDOk

`func (o *ContainerRegistrySpecCredentials) GetAccessKeyIDOk() (*string, bool)`

GetAccessKeyIDOk returns a tuple with the AccessKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyID

`func (o *ContainerRegistrySpecCredentials) SetAccessKeyID(v string)`

SetAccessKeyID sets AccessKeyID field to given value.

### HasAccessKeyID

`func (o *ContainerRegistrySpecCredentials) HasAccessKeyID() bool`

HasAccessKeyID returns a boolean if a field has been set.

### GetAccessSecretKey

`func (o *ContainerRegistrySpecCredentials) GetAccessSecretKey() string`

GetAccessSecretKey returns the AccessSecretKey field if non-nil, zero value otherwise.

### GetAccessSecretKeyOk

`func (o *ContainerRegistrySpecCredentials) GetAccessSecretKeyOk() (*string, bool)`

GetAccessSecretKeyOk returns a tuple with the AccessSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSecretKey

`func (o *ContainerRegistrySpecCredentials) SetAccessSecretKey(v string)`

SetAccessSecretKey sets AccessSecretKey field to given value.

### HasAccessSecretKey

`func (o *ContainerRegistrySpecCredentials) HasAccessSecretKey() bool`

HasAccessSecretKey returns a boolean if a field has been set.

### GetRegion

`func (o *ContainerRegistrySpecCredentials) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ContainerRegistrySpecCredentials) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ContainerRegistrySpecCredentials) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ContainerRegistrySpecCredentials) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetJsonKeyData

`func (o *ContainerRegistrySpecCredentials) GetJsonKeyData() string`

GetJsonKeyData returns the JsonKeyData field if non-nil, zero value otherwise.

### GetJsonKeyDataOk

`func (o *ContainerRegistrySpecCredentials) GetJsonKeyDataOk() (*string, bool)`

GetJsonKeyDataOk returns a tuple with the JsonKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonKeyData

`func (o *ContainerRegistrySpecCredentials) SetJsonKeyData(v string)`

SetJsonKeyData sets JsonKeyData field to given value.

### HasJsonKeyData

`func (o *ContainerRegistrySpecCredentials) HasJsonKeyData() bool`

HasJsonKeyData returns a boolean if a field has been set.

### GetServicePrincipalID

`func (o *ContainerRegistrySpecCredentials) GetServicePrincipalID() string`

GetServicePrincipalID returns the ServicePrincipalID field if non-nil, zero value otherwise.

### GetServicePrincipalIDOk

`func (o *ContainerRegistrySpecCredentials) GetServicePrincipalIDOk() (*string, bool)`

GetServicePrincipalIDOk returns a tuple with the ServicePrincipalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalID

`func (o *ContainerRegistrySpecCredentials) SetServicePrincipalID(v string)`

SetServicePrincipalID sets ServicePrincipalID field to given value.

### HasServicePrincipalID

`func (o *ContainerRegistrySpecCredentials) HasServicePrincipalID() bool`

HasServicePrincipalID returns a boolean if a field has been set.

### GetServicePrincipalPassword

`func (o *ContainerRegistrySpecCredentials) GetServicePrincipalPassword() string`

GetServicePrincipalPassword returns the ServicePrincipalPassword field if non-nil, zero value otherwise.

### GetServicePrincipalPasswordOk

`func (o *ContainerRegistrySpecCredentials) GetServicePrincipalPasswordOk() (*string, bool)`

GetServicePrincipalPasswordOk returns a tuple with the ServicePrincipalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalPassword

`func (o *ContainerRegistrySpecCredentials) SetServicePrincipalPassword(v string)`

SetServicePrincipalPassword sets ServicePrincipalPassword field to given value.

### HasServicePrincipalPassword

`func (o *ContainerRegistrySpecCredentials) HasServicePrincipalPassword() bool`

HasServicePrincipalPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


