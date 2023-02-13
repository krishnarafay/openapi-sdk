# RetryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Determines whether to retry or not, by default it is set to false | [optional] 
**MaxCount** | Pointer to **int32** | Maximun number of retries | [optional] 

## Methods

### NewRetryOptions

`func NewRetryOptions() *RetryOptions`

NewRetryOptions instantiates a new RetryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryOptionsWithDefaults

`func NewRetryOptionsWithDefaults() *RetryOptions`

NewRetryOptionsWithDefaults instantiates a new RetryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RetryOptions) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RetryOptions) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RetryOptions) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RetryOptions) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaxCount

`func (o *RetryOptions) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *RetryOptions) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *RetryOptions) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *RetryOptions) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


