# PulumiProviderHooks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deploy** | Pointer to [**PulumiDeployHooks**](PulumiDeployHooks.md) |  | [optional] 
**Destroy** | Pointer to [**PulumiDestroyHooks**](PulumiDestroyHooks.md) |  | [optional] 

## Methods

### NewPulumiProviderHooks

`func NewPulumiProviderHooks() *PulumiProviderHooks`

NewPulumiProviderHooks instantiates a new PulumiProviderHooks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPulumiProviderHooksWithDefaults

`func NewPulumiProviderHooksWithDefaults() *PulumiProviderHooks`

NewPulumiProviderHooksWithDefaults instantiates a new PulumiProviderHooks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploy

`func (o *PulumiProviderHooks) GetDeploy() PulumiDeployHooks`

GetDeploy returns the Deploy field if non-nil, zero value otherwise.

### GetDeployOk

`func (o *PulumiProviderHooks) GetDeployOk() (*PulumiDeployHooks, bool)`

GetDeployOk returns a tuple with the Deploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploy

`func (o *PulumiProviderHooks) SetDeploy(v PulumiDeployHooks)`

SetDeploy sets Deploy field to given value.

### HasDeploy

`func (o *PulumiProviderHooks) HasDeploy() bool`

HasDeploy returns a boolean if a field has been set.

### GetDestroy

`func (o *PulumiProviderHooks) GetDestroy() PulumiDestroyHooks`

GetDestroy returns the Destroy field if non-nil, zero value otherwise.

### GetDestroyOk

`func (o *PulumiProviderHooks) GetDestroyOk() (*PulumiDestroyHooks, bool)`

GetDestroyOk returns a tuple with the Destroy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroy

`func (o *PulumiProviderHooks) SetDestroy(v PulumiDestroyHooks)`

SetDestroy sets Destroy field to given value.

### HasDestroy

`func (o *PulumiProviderHooks) HasDestroy() bool`

HasDestroy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


