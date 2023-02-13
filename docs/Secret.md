# Secret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePath** | **string** | File path where secret will be written | 
**Secret** | **string** | Secret | 

## Methods

### NewSecret

`func NewSecret(filePath string, secret string, ) *Secret`

NewSecret instantiates a new Secret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretWithDefaults

`func NewSecretWithDefaults() *Secret`

NewSecretWithDefaults instantiates a new Secret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePath

`func (o *Secret) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *Secret) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *Secret) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.


### GetSecret

`func (o *Secret) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Secret) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Secret) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


