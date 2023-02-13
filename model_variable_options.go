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

// checks if the VariableOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableOptions{}

// VariableOptions struct for VariableOptions
type VariableOptions struct {
	// Description of the variable
	Description *string `json:"description,omitempty"`
	Override *VariableOverrideOptions `json:"override,omitempty"`
	// Specify whether variable is read only or not, by default it is writable
	ReadOnly *bool `json:"readOnly,omitempty"`
	// Specify whether this variable is required or optional, by default it is optional
	Required *bool `json:"required,omitempty"`
	// Determines whether the value is sensitive or not, accordingly applies encryption on it
	Sensitive *bool `json:"sensitive,omitempty"`
}

// NewVariableOptions instantiates a new VariableOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableOptions() *VariableOptions {
	this := VariableOptions{}
	return &this
}

// NewVariableOptionsWithDefaults instantiates a new VariableOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableOptionsWithDefaults() *VariableOptions {
	this := VariableOptions{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VariableOptions) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableOptions) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VariableOptions) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VariableOptions) SetDescription(v string) {
	o.Description = &v
}

// GetOverride returns the Override field value if set, zero value otherwise.
func (o *VariableOptions) GetOverride() VariableOverrideOptions {
	if o == nil || isNil(o.Override) {
		var ret VariableOverrideOptions
		return ret
	}
	return *o.Override
}

// GetOverrideOk returns a tuple with the Override field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableOptions) GetOverrideOk() (*VariableOverrideOptions, bool) {
	if o == nil || isNil(o.Override) {
		return nil, false
	}
	return o.Override, true
}

// HasOverride returns a boolean if a field has been set.
func (o *VariableOptions) HasOverride() bool {
	if o != nil && !isNil(o.Override) {
		return true
	}

	return false
}

// SetOverride gets a reference to the given VariableOverrideOptions and assigns it to the Override field.
func (o *VariableOptions) SetOverride(v VariableOverrideOptions) {
	o.Override = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *VariableOptions) GetReadOnly() bool {
	if o == nil || isNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableOptions) GetReadOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *VariableOptions) HasReadOnly() bool {
	if o != nil && !isNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *VariableOptions) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *VariableOptions) GetRequired() bool {
	if o == nil || isNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableOptions) GetRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *VariableOptions) HasRequired() bool {
	if o != nil && !isNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *VariableOptions) SetRequired(v bool) {
	o.Required = &v
}

// GetSensitive returns the Sensitive field value if set, zero value otherwise.
func (o *VariableOptions) GetSensitive() bool {
	if o == nil || isNil(o.Sensitive) {
		var ret bool
		return ret
	}
	return *o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableOptions) GetSensitiveOk() (*bool, bool) {
	if o == nil || isNil(o.Sensitive) {
		return nil, false
	}
	return o.Sensitive, true
}

// HasSensitive returns a boolean if a field has been set.
func (o *VariableOptions) HasSensitive() bool {
	if o != nil && !isNil(o.Sensitive) {
		return true
	}

	return false
}

// SetSensitive gets a reference to the given bool and assigns it to the Sensitive field.
func (o *VariableOptions) SetSensitive(v bool) {
	o.Sensitive = &v
}

func (o VariableOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Override) {
		toSerialize["override"] = o.Override
	}
	if !isNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if !isNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !isNil(o.Sensitive) {
		toSerialize["sensitive"] = o.Sensitive
	}
	return toSerialize, nil
}

type NullableVariableOptions struct {
	value *VariableOptions
	isSet bool
}

func (v NullableVariableOptions) Get() *VariableOptions {
	return v.value
}

func (v *NullableVariableOptions) Set(val *VariableOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableOptions(val *VariableOptions) *NullableVariableOptions {
	return &NullableVariableOptions{value: val, isSet: true}
}

func (v NullableVariableOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

