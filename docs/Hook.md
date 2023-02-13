# Hook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]ResourceNameAndVersionRef**](ResourceNameAndVersionRef.md) | Specify the resource ref agents | [optional] 
**OnFailure** | Pointer to **string** | Specify the on failure action | [optional] 
**Options** | Pointer to [**HookOptions**](HookOptions.md) |  | [optional] 
**Retry** | Pointer to [**RetryOptions**](RetryOptions.md) |  | [optional] 
**TimeoutSeconds** | Pointer to **int64** | Specify the timeout in seconds | [optional] 
**Type** | Pointer to **string** | Specify the type of hook | [optional] 

## Methods

### NewHook

`func NewHook() *Hook`

NewHook instantiates a new Hook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookWithDefaults

`func NewHookWithDefaults() *Hook`

NewHookWithDefaults instantiates a new Hook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *Hook) GetAgents() []ResourceNameAndVersionRef`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *Hook) GetAgentsOk() (*[]ResourceNameAndVersionRef, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *Hook) SetAgents(v []ResourceNameAndVersionRef)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *Hook) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetOnFailure

`func (o *Hook) GetOnFailure() string`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *Hook) GetOnFailureOk() (*string, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *Hook) SetOnFailure(v string)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *Hook) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOptions

`func (o *Hook) GetOptions() HookOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Hook) GetOptionsOk() (*HookOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Hook) SetOptions(v HookOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Hook) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetRetry

`func (o *Hook) GetRetry() RetryOptions`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *Hook) GetRetryOk() (*RetryOptions, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *Hook) SetRetry(v RetryOptions)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *Hook) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *Hook) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *Hook) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *Hook) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *Hook) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetType

`func (o *Hook) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Hook) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Hook) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Hook) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


