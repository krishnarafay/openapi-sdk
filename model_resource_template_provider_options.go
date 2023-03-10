/*
Kubernetes Operations APIs

Kubernetes Operations APIs

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rafaysdkgov3

import (
	"encoding/json"
)

// checks if the ResourceTemplateProviderOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceTemplateProviderOptions{}

// ResourceTemplateProviderOptions struct for ResourceTemplateProviderOptions
type ResourceTemplateProviderOptions struct {
	Pulumi map[string]interface{} `json:"pulumi,omitempty"`
	RetryOptions *RetryOptions `json:"retry_options,omitempty"`
	System map[string]interface{} `json:"system,omitempty"`
	Terraform *TerraformProviderOptions `json:"terraform,omitempty"`
	Terragrunt map[string]interface{} `json:"terragrunt,omitempty"`
}

// NewResourceTemplateProviderOptions instantiates a new ResourceTemplateProviderOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceTemplateProviderOptions() *ResourceTemplateProviderOptions {
	this := ResourceTemplateProviderOptions{}
	return &this
}

// NewResourceTemplateProviderOptionsWithDefaults instantiates a new ResourceTemplateProviderOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceTemplateProviderOptionsWithDefaults() *ResourceTemplateProviderOptions {
	this := ResourceTemplateProviderOptions{}
	return &this
}

// GetPulumi returns the Pulumi field value if set, zero value otherwise.
func (o *ResourceTemplateProviderOptions) GetPulumi() map[string]interface{} {
	if o == nil || isNil(o.Pulumi) {
		var ret map[string]interface{}
		return ret
	}
	return o.Pulumi
}

// GetPulumiOk returns a tuple with the Pulumi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTemplateProviderOptions) GetPulumiOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Pulumi) {
		return map[string]interface{}{}, false
	}
	return o.Pulumi, true
}

// HasPulumi returns a boolean if a field has been set.
func (o *ResourceTemplateProviderOptions) HasPulumi() bool {
	if o != nil && !isNil(o.Pulumi) {
		return true
	}

	return false
}

// SetPulumi gets a reference to the given map[string]interface{} and assigns it to the Pulumi field.
func (o *ResourceTemplateProviderOptions) SetPulumi(v map[string]interface{}) {
	o.Pulumi = v
}

// GetRetryOptions returns the RetryOptions field value if set, zero value otherwise.
func (o *ResourceTemplateProviderOptions) GetRetryOptions() RetryOptions {
	if o == nil || isNil(o.RetryOptions) {
		var ret RetryOptions
		return ret
	}
	return *o.RetryOptions
}

// GetRetryOptionsOk returns a tuple with the RetryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTemplateProviderOptions) GetRetryOptionsOk() (*RetryOptions, bool) {
	if o == nil || isNil(o.RetryOptions) {
		return nil, false
	}
	return o.RetryOptions, true
}

// HasRetryOptions returns a boolean if a field has been set.
func (o *ResourceTemplateProviderOptions) HasRetryOptions() bool {
	if o != nil && !isNil(o.RetryOptions) {
		return true
	}

	return false
}

// SetRetryOptions gets a reference to the given RetryOptions and assigns it to the RetryOptions field.
func (o *ResourceTemplateProviderOptions) SetRetryOptions(v RetryOptions) {
	o.RetryOptions = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *ResourceTemplateProviderOptions) GetSystem() map[string]interface{} {
	if o == nil || isNil(o.System) {
		var ret map[string]interface{}
		return ret
	}
	return o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTemplateProviderOptions) GetSystemOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.System) {
		return map[string]interface{}{}, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *ResourceTemplateProviderOptions) HasSystem() bool {
	if o != nil && !isNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given map[string]interface{} and assigns it to the System field.
func (o *ResourceTemplateProviderOptions) SetSystem(v map[string]interface{}) {
	o.System = v
}

// GetTerraform returns the Terraform field value if set, zero value otherwise.
func (o *ResourceTemplateProviderOptions) GetTerraform() TerraformProviderOptions {
	if o == nil || isNil(o.Terraform) {
		var ret TerraformProviderOptions
		return ret
	}
	return *o.Terraform
}

// GetTerraformOk returns a tuple with the Terraform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTemplateProviderOptions) GetTerraformOk() (*TerraformProviderOptions, bool) {
	if o == nil || isNil(o.Terraform) {
		return nil, false
	}
	return o.Terraform, true
}

// HasTerraform returns a boolean if a field has been set.
func (o *ResourceTemplateProviderOptions) HasTerraform() bool {
	if o != nil && !isNil(o.Terraform) {
		return true
	}

	return false
}

// SetTerraform gets a reference to the given TerraformProviderOptions and assigns it to the Terraform field.
func (o *ResourceTemplateProviderOptions) SetTerraform(v TerraformProviderOptions) {
	o.Terraform = &v
}

// GetTerragrunt returns the Terragrunt field value if set, zero value otherwise.
func (o *ResourceTemplateProviderOptions) GetTerragrunt() map[string]interface{} {
	if o == nil || isNil(o.Terragrunt) {
		var ret map[string]interface{}
		return ret
	}
	return o.Terragrunt
}

// GetTerragruntOk returns a tuple with the Terragrunt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTemplateProviderOptions) GetTerragruntOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Terragrunt) {
		return map[string]interface{}{}, false
	}
	return o.Terragrunt, true
}

// HasTerragrunt returns a boolean if a field has been set.
func (o *ResourceTemplateProviderOptions) HasTerragrunt() bool {
	if o != nil && !isNil(o.Terragrunt) {
		return true
	}

	return false
}

// SetTerragrunt gets a reference to the given map[string]interface{} and assigns it to the Terragrunt field.
func (o *ResourceTemplateProviderOptions) SetTerragrunt(v map[string]interface{}) {
	o.Terragrunt = v
}

func (o ResourceTemplateProviderOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceTemplateProviderOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Pulumi) {
		toSerialize["pulumi"] = o.Pulumi
	}
	if !isNil(o.RetryOptions) {
		toSerialize["retry_options"] = o.RetryOptions
	}
	if !isNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !isNil(o.Terraform) {
		toSerialize["terraform"] = o.Terraform
	}
	if !isNil(o.Terragrunt) {
		toSerialize["terragrunt"] = o.Terragrunt
	}
	return toSerialize, nil
}

type NullableResourceTemplateProviderOptions struct {
	value *ResourceTemplateProviderOptions
	isSet bool
}

func (v NullableResourceTemplateProviderOptions) Get() *ResourceTemplateProviderOptions {
	return v.value
}

func (v *NullableResourceTemplateProviderOptions) Set(val *ResourceTemplateProviderOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTemplateProviderOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTemplateProviderOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTemplateProviderOptions(val *ResourceTemplateProviderOptions) *NullableResourceTemplateProviderOptions {
	return &NullableResourceTemplateProviderOptions{value: val, isSet: true}
}

func (v NullableResourceTemplateProviderOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTemplateProviderOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


