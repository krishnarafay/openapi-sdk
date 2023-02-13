# TerraformProviderHooks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deploy** | Pointer to [**TerraformDeployHooks**](TerraformDeployHooks.md) |  | [optional] 
**Destroy** | Pointer to [**TerraformDestroyHooks**](TerraformDestroyHooks.md) |  | [optional] 

## Methods

### NewTerraformProviderHooks

`func NewTerraformProviderHooks() *TerraformProviderHooks`

NewTerraformProviderHooks instantiates a new TerraformProviderHooks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformProviderHooksWithDefaults

`func NewTerraformProviderHooksWithDefaults() *TerraformProviderHooks`

NewTerraformProviderHooksWithDefaults instantiates a new TerraformProviderHooks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploy

`func (o *TerraformProviderHooks) GetDeploy() TerraformDeployHooks`

GetDeploy returns the Deploy field if non-nil, zero value otherwise.

### GetDeployOk

`func (o *TerraformProviderHooks) GetDeployOk() (*TerraformDeployHooks, bool)`

GetDeployOk returns a tuple with the Deploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploy

`func (o *TerraformProviderHooks) SetDeploy(v TerraformDeployHooks)`

SetDeploy sets Deploy field to given value.

### HasDeploy

`func (o *TerraformProviderHooks) HasDeploy() bool`

HasDeploy returns a boolean if a field has been set.

### GetDestroy

`func (o *TerraformProviderHooks) GetDestroy() TerraformDestroyHooks`

GetDestroy returns the Destroy field if non-nil, zero value otherwise.

### GetDestroyOk

`func (o *TerraformProviderHooks) GetDestroyOk() (*TerraformDestroyHooks, bool)`

GetDestroyOk returns a tuple with the Destroy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroy

`func (o *TerraformProviderHooks) SetDestroy(v TerraformDestroyHooks)`

SetDestroy sets Destroy field to given value.

### HasDestroy

`func (o *TerraformProviderHooks) HasDestroy() bool`

HasDestroy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


