# ResourceTemplateProviderHooks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pulumi** | Pointer to [**PulumiProviderHooks**](PulumiProviderHooks.md) |  | [optional] 
**Terraform** | Pointer to [**TerraformProviderHooks**](TerraformProviderHooks.md) |  | [optional] 

## Methods

### NewResourceTemplateProviderHooks

`func NewResourceTemplateProviderHooks() *ResourceTemplateProviderHooks`

NewResourceTemplateProviderHooks instantiates a new ResourceTemplateProviderHooks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTemplateProviderHooksWithDefaults

`func NewResourceTemplateProviderHooksWithDefaults() *ResourceTemplateProviderHooks`

NewResourceTemplateProviderHooksWithDefaults instantiates a new ResourceTemplateProviderHooks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulumi

`func (o *ResourceTemplateProviderHooks) GetPulumi() PulumiProviderHooks`

GetPulumi returns the Pulumi field if non-nil, zero value otherwise.

### GetPulumiOk

`func (o *ResourceTemplateProviderHooks) GetPulumiOk() (*PulumiProviderHooks, bool)`

GetPulumiOk returns a tuple with the Pulumi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulumi

`func (o *ResourceTemplateProviderHooks) SetPulumi(v PulumiProviderHooks)`

SetPulumi sets Pulumi field to given value.

### HasPulumi

`func (o *ResourceTemplateProviderHooks) HasPulumi() bool`

HasPulumi returns a boolean if a field has been set.

### GetTerraform

`func (o *ResourceTemplateProviderHooks) GetTerraform() TerraformProviderHooks`

GetTerraform returns the Terraform field if non-nil, zero value otherwise.

### GetTerraformOk

`func (o *ResourceTemplateProviderHooks) GetTerraformOk() (*TerraformProviderHooks, bool)`

GetTerraformOk returns a tuple with the Terraform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraform

`func (o *ResourceTemplateProviderHooks) SetTerraform(v TerraformProviderHooks)`

SetTerraform sets Terraform field to given value.

### HasTerraform

`func (o *ResourceTemplateProviderHooks) HasTerraform() bool`

HasTerraform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


