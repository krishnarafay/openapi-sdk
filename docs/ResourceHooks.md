# ResourceHooks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnCompletion** | Pointer to [**[]Hook**](Hook.md) | Configure the on completion lifecycle hook | [optional] 
**OnFailure** | Pointer to [**[]Hook**](Hook.md) | Configure the on failure lifecycle hook | [optional] 
**OnInit** | Pointer to [**[]Hook**](Hook.md) | Configure the on initialize lifecycle hook | [optional] 
**OnSuccess** | Pointer to [**[]Hook**](Hook.md) | Configure the on success lifecycle hook | [optional] 
**Provider** | Pointer to [**ResourceTemplateProviderHooks**](ResourceTemplateProviderHooks.md) |  | [optional] 

## Methods

### NewResourceHooks

`func NewResourceHooks() *ResourceHooks`

NewResourceHooks instantiates a new ResourceHooks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceHooksWithDefaults

`func NewResourceHooksWithDefaults() *ResourceHooks`

NewResourceHooksWithDefaults instantiates a new ResourceHooks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnCompletion

`func (o *ResourceHooks) GetOnCompletion() []Hook`

GetOnCompletion returns the OnCompletion field if non-nil, zero value otherwise.

### GetOnCompletionOk

`func (o *ResourceHooks) GetOnCompletionOk() (*[]Hook, bool)`

GetOnCompletionOk returns a tuple with the OnCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCompletion

`func (o *ResourceHooks) SetOnCompletion(v []Hook)`

SetOnCompletion sets OnCompletion field to given value.

### HasOnCompletion

`func (o *ResourceHooks) HasOnCompletion() bool`

HasOnCompletion returns a boolean if a field has been set.

### GetOnFailure

`func (o *ResourceHooks) GetOnFailure() []Hook`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *ResourceHooks) GetOnFailureOk() (*[]Hook, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *ResourceHooks) SetOnFailure(v []Hook)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *ResourceHooks) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOnInit

`func (o *ResourceHooks) GetOnInit() []Hook`

GetOnInit returns the OnInit field if non-nil, zero value otherwise.

### GetOnInitOk

`func (o *ResourceHooks) GetOnInitOk() (*[]Hook, bool)`

GetOnInitOk returns a tuple with the OnInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnInit

`func (o *ResourceHooks) SetOnInit(v []Hook)`

SetOnInit sets OnInit field to given value.

### HasOnInit

`func (o *ResourceHooks) HasOnInit() bool`

HasOnInit returns a boolean if a field has been set.

### GetOnSuccess

`func (o *ResourceHooks) GetOnSuccess() []Hook`

GetOnSuccess returns the OnSuccess field if non-nil, zero value otherwise.

### GetOnSuccessOk

`func (o *ResourceHooks) GetOnSuccessOk() (*[]Hook, bool)`

GetOnSuccessOk returns a tuple with the OnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuccess

`func (o *ResourceHooks) SetOnSuccess(v []Hook)`

SetOnSuccess sets OnSuccess field to given value.

### HasOnSuccess

`func (o *ResourceHooks) HasOnSuccess() bool`

HasOnSuccess returns a boolean if a field has been set.

### GetProvider

`func (o *ResourceHooks) GetProvider() ResourceTemplateProviderHooks`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ResourceHooks) GetProviderOk() (*ResourceTemplateProviderHooks, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ResourceHooks) SetProvider(v ResourceTemplateProviderHooks)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ResourceHooks) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


