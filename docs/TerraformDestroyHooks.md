# TerraformDestroyHooks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destroy** | Pointer to [**LifecycleEventHooks**](LifecycleEventHooks.md) |  | [optional] 
**Init** | Pointer to [**LifecycleEventHooks**](LifecycleEventHooks.md) |  | [optional] 
**Plan** | Pointer to [**LifecycleEventHooks**](LifecycleEventHooks.md) |  | [optional] 

## Methods

### NewTerraformDestroyHooks

`func NewTerraformDestroyHooks() *TerraformDestroyHooks`

NewTerraformDestroyHooks instantiates a new TerraformDestroyHooks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformDestroyHooksWithDefaults

`func NewTerraformDestroyHooksWithDefaults() *TerraformDestroyHooks`

NewTerraformDestroyHooksWithDefaults instantiates a new TerraformDestroyHooks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestroy

`func (o *TerraformDestroyHooks) GetDestroy() LifecycleEventHooks`

GetDestroy returns the Destroy field if non-nil, zero value otherwise.

### GetDestroyOk

`func (o *TerraformDestroyHooks) GetDestroyOk() (*LifecycleEventHooks, bool)`

GetDestroyOk returns a tuple with the Destroy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroy

`func (o *TerraformDestroyHooks) SetDestroy(v LifecycleEventHooks)`

SetDestroy sets Destroy field to given value.

### HasDestroy

`func (o *TerraformDestroyHooks) HasDestroy() bool`

HasDestroy returns a boolean if a field has been set.

### GetInit

`func (o *TerraformDestroyHooks) GetInit() LifecycleEventHooks`

GetInit returns the Init field if non-nil, zero value otherwise.

### GetInitOk

`func (o *TerraformDestroyHooks) GetInitOk() (*LifecycleEventHooks, bool)`

GetInitOk returns a tuple with the Init field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInit

`func (o *TerraformDestroyHooks) SetInit(v LifecycleEventHooks)`

SetInit sets Init field to given value.

### HasInit

`func (o *TerraformDestroyHooks) HasInit() bool`

HasInit returns a boolean if a field has been set.

### GetPlan

`func (o *TerraformDestroyHooks) GetPlan() LifecycleEventHooks`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *TerraformDestroyHooks) GetPlanOk() (*LifecycleEventHooks, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *TerraformDestroyHooks) SetPlan(v LifecycleEventHooks)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *TerraformDestroyHooks) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


