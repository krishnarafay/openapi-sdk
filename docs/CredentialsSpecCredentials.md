# CredentialsSpecCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Arn** | **string** |  | 
**ExternalId** | **string** |  | 
**Type** | **string** | Type of AWS Credentials Access | [default to "AccessKey"]
**AccessId** | **string** |  | 
**SecretKey** | **string** |  | 
**SessionToken** | Pointer to **string** |  | [optional] 
**File** | **string** | Blob content of GCP Credentials Access JSON file | 
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**SubscriptionId** | **string** |  | 
**TenantId** | **string** |  | 
**GatewayId** | **string** |  | 
**Password** | **string** |  | 
**Username** | **string** |  | 
**VsphereServer** | **string** |  | 

## Methods

### NewCredentialsSpecCredentials

`func NewCredentialsSpecCredentials(accountId string, arn string, externalId string, type_ string, accessId string, secretKey string, file string, clientId string, clientSecret string, subscriptionId string, tenantId string, gatewayId string, password string, username string, vsphereServer string, ) *CredentialsSpecCredentials`

NewCredentialsSpecCredentials instantiates a new CredentialsSpecCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsSpecCredentialsWithDefaults

`func NewCredentialsSpecCredentialsWithDefaults() *CredentialsSpecCredentials`

NewCredentialsSpecCredentialsWithDefaults instantiates a new CredentialsSpecCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CredentialsSpecCredentials) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CredentialsSpecCredentials) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CredentialsSpecCredentials) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetArn

`func (o *CredentialsSpecCredentials) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *CredentialsSpecCredentials) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *CredentialsSpecCredentials) SetArn(v string)`

SetArn sets Arn field to given value.


### GetExternalId

`func (o *CredentialsSpecCredentials) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CredentialsSpecCredentials) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CredentialsSpecCredentials) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetType

`func (o *CredentialsSpecCredentials) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialsSpecCredentials) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialsSpecCredentials) SetType(v string)`

SetType sets Type field to given value.


### GetAccessId

`func (o *CredentialsSpecCredentials) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *CredentialsSpecCredentials) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *CredentialsSpecCredentials) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.


### GetSecretKey

`func (o *CredentialsSpecCredentials) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *CredentialsSpecCredentials) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *CredentialsSpecCredentials) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetSessionToken

`func (o *CredentialsSpecCredentials) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *CredentialsSpecCredentials) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *CredentialsSpecCredentials) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *CredentialsSpecCredentials) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetFile

`func (o *CredentialsSpecCredentials) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CredentialsSpecCredentials) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CredentialsSpecCredentials) SetFile(v string)`

SetFile sets File field to given value.


### GetClientId

`func (o *CredentialsSpecCredentials) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CredentialsSpecCredentials) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CredentialsSpecCredentials) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *CredentialsSpecCredentials) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CredentialsSpecCredentials) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CredentialsSpecCredentials) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetSubscriptionId

`func (o *CredentialsSpecCredentials) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CredentialsSpecCredentials) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CredentialsSpecCredentials) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTenantId

`func (o *CredentialsSpecCredentials) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CredentialsSpecCredentials) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CredentialsSpecCredentials) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetGatewayId

`func (o *CredentialsSpecCredentials) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *CredentialsSpecCredentials) GetGatewayIdOk() (*string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *CredentialsSpecCredentials) SetGatewayId(v string)`

SetGatewayId sets GatewayId field to given value.


### GetPassword

`func (o *CredentialsSpecCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CredentialsSpecCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CredentialsSpecCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *CredentialsSpecCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CredentialsSpecCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CredentialsSpecCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetVsphereServer

`func (o *CredentialsSpecCredentials) GetVsphereServer() string`

GetVsphereServer returns the VsphereServer field if non-nil, zero value otherwise.

### GetVsphereServerOk

`func (o *CredentialsSpecCredentials) GetVsphereServerOk() (*string, bool)`

GetVsphereServerOk returns a tuple with the VsphereServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereServer

`func (o *CredentialsSpecCredentials) SetVsphereServer(v string)`

SetVsphereServer sets VsphereServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


