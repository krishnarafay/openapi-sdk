# AzureCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**SubscriptionId** | **string** |  | 
**TenantId** | **string** |  | 

## Methods

### NewAzureCredentials

`func NewAzureCredentials(clientId string, clientSecret string, subscriptionId string, tenantId string, ) *AzureCredentials`

NewAzureCredentials instantiates a new AzureCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCredentialsWithDefaults

`func NewAzureCredentialsWithDefaults() *AzureCredentials`

NewAzureCredentialsWithDefaults instantiates a new AzureCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AzureCredentials) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureCredentials) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AzureCredentials) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *AzureCredentials) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AzureCredentials) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AzureCredentials) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetSubscriptionId

`func (o *AzureCredentials) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureCredentials) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureCredentials) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTenantId

`func (o *AzureCredentials) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureCredentials) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureCredentials) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


