# CredentialsSpecAwsAccessKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | **string** |  | 
**SecretKey** | **string** |  | 
**SessionToken** | Pointer to **string** |  | [optional] 
**Type** | **string** | Type of AWS Credentials Access | [default to "AccessKey"]

## Methods

### NewCredentialsSpecAwsAccessKey

`func NewCredentialsSpecAwsAccessKey(accessId string, secretKey string, type_ string, ) *CredentialsSpecAwsAccessKey`

NewCredentialsSpecAwsAccessKey instantiates a new CredentialsSpecAwsAccessKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsSpecAwsAccessKeyWithDefaults

`func NewCredentialsSpecAwsAccessKeyWithDefaults() *CredentialsSpecAwsAccessKey`

NewCredentialsSpecAwsAccessKeyWithDefaults instantiates a new CredentialsSpecAwsAccessKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *CredentialsSpecAwsAccessKey) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *CredentialsSpecAwsAccessKey) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *CredentialsSpecAwsAccessKey) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.


### GetSecretKey

`func (o *CredentialsSpecAwsAccessKey) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *CredentialsSpecAwsAccessKey) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *CredentialsSpecAwsAccessKey) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetSessionToken

`func (o *CredentialsSpecAwsAccessKey) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *CredentialsSpecAwsAccessKey) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *CredentialsSpecAwsAccessKey) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *CredentialsSpecAwsAccessKey) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetType

`func (o *CredentialsSpecAwsAccessKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialsSpecAwsAccessKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialsSpecAwsAccessKey) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


