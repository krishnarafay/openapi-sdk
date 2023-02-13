# LifecycleEventHooks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to [**[]Hook**](Hook.md) | Specify all the after lifecycle hook | [optional] 
**Before** | Pointer to [**[]Hook**](Hook.md) | Specify all the before lifecycle hook | [optional] 

## Methods

### NewLifecycleEventHooks

`func NewLifecycleEventHooks() *LifecycleEventHooks`

NewLifecycleEventHooks instantiates a new LifecycleEventHooks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleEventHooksWithDefaults

`func NewLifecycleEventHooksWithDefaults() *LifecycleEventHooks`

NewLifecycleEventHooksWithDefaults instantiates a new LifecycleEventHooks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *LifecycleEventHooks) GetAfter() []Hook`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *LifecycleEventHooks) GetAfterOk() (*[]Hook, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *LifecycleEventHooks) SetAfter(v []Hook)`

SetAfter sets After field to given value.

### HasAfter

`func (o *LifecycleEventHooks) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetBefore

`func (o *LifecycleEventHooks) GetBefore() []Hook`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *LifecycleEventHooks) GetBeforeOk() (*[]Hook, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *LifecycleEventHooks) SetBefore(v []Hook)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *LifecycleEventHooks) HasBefore() bool`

HasBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


