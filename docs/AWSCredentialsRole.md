# AWSCredentialsRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Arn** | **string** |  | 
**ExternalId** | **string** |  | 
**Type** | **string** | Type of AWS Credentials Access | [default to "Role"]

## Methods

### NewAWSCredentialsRole

`func NewAWSCredentialsRole(accountId string, arn string, externalId string, type_ string, ) *AWSCredentialsRole`

NewAWSCredentialsRole instantiates a new AWSCredentialsRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSCredentialsRoleWithDefaults

`func NewAWSCredentialsRoleWithDefaults() *AWSCredentialsRole`

NewAWSCredentialsRoleWithDefaults instantiates a new AWSCredentialsRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AWSCredentialsRole) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AWSCredentialsRole) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AWSCredentialsRole) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetArn

`func (o *AWSCredentialsRole) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *AWSCredentialsRole) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *AWSCredentialsRole) SetArn(v string)`

SetArn sets Arn field to given value.


### GetExternalId

`func (o *AWSCredentialsRole) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AWSCredentialsRole) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AWSCredentialsRole) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetType

`func (o *AWSCredentialsRole) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWSCredentialsRole) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWSCredentialsRole) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


