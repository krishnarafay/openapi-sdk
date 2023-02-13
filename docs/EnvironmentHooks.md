# EnvironmentHooks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnCompletion** | Pointer to [**[]Hook**](Hook.md) | On completion environment hook | [optional] 
**OnFailure** | Pointer to [**[]Hook**](Hook.md) | On failure environment hook | [optional] 
**OnInit** | Pointer to [**[]Hook**](Hook.md) | On initialize environment hook | [optional] 
**OnSuccess** | Pointer to [**[]Hook**](Hook.md) | On success environment hook | [optional] 

## Methods

### NewEnvironmentHooks

`func NewEnvironmentHooks() *EnvironmentHooks`

NewEnvironmentHooks instantiates a new EnvironmentHooks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentHooksWithDefaults

`func NewEnvironmentHooksWithDefaults() *EnvironmentHooks`

NewEnvironmentHooksWithDefaults instantiates a new EnvironmentHooks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnCompletion

`func (o *EnvironmentHooks) GetOnCompletion() []Hook`

GetOnCompletion returns the OnCompletion field if non-nil, zero value otherwise.

### GetOnCompletionOk

`func (o *EnvironmentHooks) GetOnCompletionOk() (*[]Hook, bool)`

GetOnCompletionOk returns a tuple with the OnCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCompletion

`func (o *EnvironmentHooks) SetOnCompletion(v []Hook)`

SetOnCompletion sets OnCompletion field to given value.

### HasOnCompletion

`func (o *EnvironmentHooks) HasOnCompletion() bool`

HasOnCompletion returns a boolean if a field has been set.

### GetOnFailure

`func (o *EnvironmentHooks) GetOnFailure() []Hook`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *EnvironmentHooks) GetOnFailureOk() (*[]Hook, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *EnvironmentHooks) SetOnFailure(v []Hook)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *EnvironmentHooks) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOnInit

`func (o *EnvironmentHooks) GetOnInit() []Hook`

GetOnInit returns the OnInit field if non-nil, zero value otherwise.

### GetOnInitOk

`func (o *EnvironmentHooks) GetOnInitOk() (*[]Hook, bool)`

GetOnInitOk returns a tuple with the OnInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnInit

`func (o *EnvironmentHooks) SetOnInit(v []Hook)`

SetOnInit sets OnInit field to given value.

### HasOnInit

`func (o *EnvironmentHooks) HasOnInit() bool`

HasOnInit returns a boolean if a field has been set.

### GetOnSuccess

`func (o *EnvironmentHooks) GetOnSuccess() []Hook`

GetOnSuccess returns the OnSuccess field if non-nil, zero value otherwise.

### GetOnSuccessOk

`func (o *EnvironmentHooks) GetOnSuccessOk() (*[]Hook, bool)`

GetOnSuccessOk returns a tuple with the OnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuccess

`func (o *EnvironmentHooks) SetOnSuccess(v []Hook)`

SetOnSuccess sets OnSuccess field to given value.

### HasOnSuccess

`func (o *EnvironmentHooks) HasOnSuccess() bool`

HasOnSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


