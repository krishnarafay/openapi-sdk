# HttpOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Specify the body of http request | [optional] 
**Endpoint** | Pointer to **string** | Specify the endpoint to be invoked | [optional] 
**Headers** | Pointer to **map[string]string** | Specify the http headers in key,value pair | [optional] 
**Method** | Pointer to **string** | Specify the http method to be used | [optional] 
**SuccessCondition** | Pointer to **string** | Specify the success condition of the request | [optional] 

## Methods

### NewHttpOptions

`func NewHttpOptions() *HttpOptions`

NewHttpOptions instantiates a new HttpOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpOptionsWithDefaults

`func NewHttpOptionsWithDefaults() *HttpOptions`

NewHttpOptionsWithDefaults instantiates a new HttpOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *HttpOptions) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *HttpOptions) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *HttpOptions) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *HttpOptions) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetEndpoint

`func (o *HttpOptions) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *HttpOptions) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *HttpOptions) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *HttpOptions) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetHeaders

`func (o *HttpOptions) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpOptions) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpOptions) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HttpOptions) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *HttpOptions) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HttpOptions) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HttpOptions) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *HttpOptions) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetSuccessCondition

`func (o *HttpOptions) GetSuccessCondition() string`

GetSuccessCondition returns the SuccessCondition field if non-nil, zero value otherwise.

### GetSuccessConditionOk

`func (o *HttpOptions) GetSuccessConditionOk() (*string, bool)`

GetSuccessConditionOk returns a tuple with the SuccessCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCondition

`func (o *HttpOptions) SetSuccessCondition(v string)`

SetSuccessCondition sets SuccessCondition field to given value.

### HasSuccessCondition

`func (o *HttpOptions) HasSuccessCondition() bool`

HasSuccessCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


