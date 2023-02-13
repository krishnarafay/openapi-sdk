# AWSCredentialsAccessKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | **string** |  | 
**SecretKey** | **string** |  | 
**SessionToken** | Pointer to **string** |  | [optional] 
**Type** | **string** | Type of AWS Credentials Access | [default to "AccessKey"]

## Methods

### NewAWSCredentialsAccessKey

`func NewAWSCredentialsAccessKey(accessId string, secretKey string, type_ string, ) *AWSCredentialsAccessKey`

NewAWSCredentialsAccessKey instantiates a new AWSCredentialsAccessKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSCredentialsAccessKeyWithDefaults

`func NewAWSCredentialsAccessKeyWithDefaults() *AWSCredentialsAccessKey`

NewAWSCredentialsAccessKeyWithDefaults instantiates a new AWSCredentialsAccessKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *AWSCredentialsAccessKey) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *AWSCredentialsAccessKey) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *AWSCredentialsAccessKey) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.


### GetSecretKey

`func (o *AWSCredentialsAccessKey) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AWSCredentialsAccessKey) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AWSCredentialsAccessKey) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetSessionToken

`func (o *AWSCredentialsAccessKey) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *AWSCredentialsAccessKey) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *AWSCredentialsAccessKey) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *AWSCredentialsAccessKey) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetType

`func (o *AWSCredentialsAccessKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWSCredentialsAccessKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWSCredentialsAccessKey) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


