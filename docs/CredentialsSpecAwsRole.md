# CredentialsSpecAwsRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Arn** | **string** |  | 
**ExternalId** | **string** |  | 
**Type** | **string** | Type of AWS Credentials Access | [default to "Role"]

## Methods

### NewCredentialsSpecAwsRole

`func NewCredentialsSpecAwsRole(accountId string, arn string, externalId string, type_ string, ) *CredentialsSpecAwsRole`

NewCredentialsSpecAwsRole instantiates a new CredentialsSpecAwsRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsSpecAwsRoleWithDefaults

`func NewCredentialsSpecAwsRoleWithDefaults() *CredentialsSpecAwsRole`

NewCredentialsSpecAwsRoleWithDefaults instantiates a new CredentialsSpecAwsRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CredentialsSpecAwsRole) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CredentialsSpecAwsRole) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CredentialsSpecAwsRole) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetArn

`func (o *CredentialsSpecAwsRole) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *CredentialsSpecAwsRole) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *CredentialsSpecAwsRole) SetArn(v string)`

SetArn sets Arn field to given value.


### GetExternalId

`func (o *CredentialsSpecAwsRole) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CredentialsSpecAwsRole) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CredentialsSpecAwsRole) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetType

`func (o *CredentialsSpecAwsRole) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialsSpecAwsRole) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialsSpecAwsRole) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


