# EnvData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key of the environment variable to be set | [optional] 
**Sensitive** | Pointer to **bool** | Determines whether the value is sensitive or not, accordingly applies encryption on it | [optional] 
**Value** | Pointer to **string** | Value of the environment variable to be set | [optional] 

## Methods

### NewEnvData

`func NewEnvData() *EnvData`

NewEnvData instantiates a new EnvData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvDataWithDefaults

`func NewEnvDataWithDefaults() *EnvData`

NewEnvDataWithDefaults instantiates a new EnvData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *EnvData) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EnvData) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EnvData) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EnvData) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSensitive

`func (o *EnvData) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *EnvData) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *EnvData) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *EnvData) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.

### GetValue

`func (o *EnvData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvData) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EnvData) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


