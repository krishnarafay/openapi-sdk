# ResourceTemplateProviderOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pulumi** | Pointer to **map[string]interface{}** |  | [optional] 
**RetryOptions** | Pointer to [**RetryOptions**](RetryOptions.md) |  | [optional] 
**System** | Pointer to **map[string]interface{}** |  | [optional] 
**Terraform** | Pointer to [**TerraformProviderOptions**](TerraformProviderOptions.md) |  | [optional] 
**Terragrunt** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResourceTemplateProviderOptions

`func NewResourceTemplateProviderOptions() *ResourceTemplateProviderOptions`

NewResourceTemplateProviderOptions instantiates a new ResourceTemplateProviderOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTemplateProviderOptionsWithDefaults

`func NewResourceTemplateProviderOptionsWithDefaults() *ResourceTemplateProviderOptions`

NewResourceTemplateProviderOptionsWithDefaults instantiates a new ResourceTemplateProviderOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulumi

`func (o *ResourceTemplateProviderOptions) GetPulumi() map[string]interface{}`

GetPulumi returns the Pulumi field if non-nil, zero value otherwise.

### GetPulumiOk

`func (o *ResourceTemplateProviderOptions) GetPulumiOk() (*map[string]interface{}, bool)`

GetPulumiOk returns a tuple with the Pulumi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulumi

`func (o *ResourceTemplateProviderOptions) SetPulumi(v map[string]interface{})`

SetPulumi sets Pulumi field to given value.

### HasPulumi

`func (o *ResourceTemplateProviderOptions) HasPulumi() bool`

HasPulumi returns a boolean if a field has been set.

### GetRetryOptions

`func (o *ResourceTemplateProviderOptions) GetRetryOptions() RetryOptions`

GetRetryOptions returns the RetryOptions field if non-nil, zero value otherwise.

### GetRetryOptionsOk

`func (o *ResourceTemplateProviderOptions) GetRetryOptionsOk() (*RetryOptions, bool)`

GetRetryOptionsOk returns a tuple with the RetryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryOptions

`func (o *ResourceTemplateProviderOptions) SetRetryOptions(v RetryOptions)`

SetRetryOptions sets RetryOptions field to given value.

### HasRetryOptions

`func (o *ResourceTemplateProviderOptions) HasRetryOptions() bool`

HasRetryOptions returns a boolean if a field has been set.

### GetSystem

`func (o *ResourceTemplateProviderOptions) GetSystem() map[string]interface{}`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ResourceTemplateProviderOptions) GetSystemOk() (*map[string]interface{}, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ResourceTemplateProviderOptions) SetSystem(v map[string]interface{})`

SetSystem sets System field to given value.

### HasSystem

`func (o *ResourceTemplateProviderOptions) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTerraform

`func (o *ResourceTemplateProviderOptions) GetTerraform() TerraformProviderOptions`

GetTerraform returns the Terraform field if non-nil, zero value otherwise.

### GetTerraformOk

`func (o *ResourceTemplateProviderOptions) GetTerraformOk() (*TerraformProviderOptions, bool)`

GetTerraformOk returns a tuple with the Terraform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraform

`func (o *ResourceTemplateProviderOptions) SetTerraform(v TerraformProviderOptions)`

SetTerraform sets Terraform field to given value.

### HasTerraform

`func (o *ResourceTemplateProviderOptions) HasTerraform() bool`

HasTerraform returns a boolean if a field has been set.

### GetTerragrunt

`func (o *ResourceTemplateProviderOptions) GetTerragrunt() map[string]interface{}`

GetTerragrunt returns the Terragrunt field if non-nil, zero value otherwise.

### GetTerragruntOk

`func (o *ResourceTemplateProviderOptions) GetTerragruntOk() (*map[string]interface{}, bool)`

GetTerragruntOk returns a tuple with the Terragrunt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerragrunt

`func (o *ResourceTemplateProviderOptions) SetTerragrunt(v map[string]interface{})`

SetTerragrunt sets Terragrunt field to given value.

### HasTerragrunt

`func (o *ResourceTemplateProviderOptions) HasTerragrunt() bool`

HasTerragrunt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


