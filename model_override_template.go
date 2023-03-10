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

// checks if the OverrideTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideTemplate{}

// OverrideTemplate repo based override template definition
type OverrideTemplate struct {
	Template *OverrideTemplateTemplate `json:"template,omitempty"`
	// type of override template
	Type *string `json:"type,omitempty"`
	// weight of the override, overrides are applied low to high weight
	Weight *int32 `json:"weight,omitempty"`
}

// NewOverrideTemplate instantiates a new OverrideTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideTemplate() *OverrideTemplate {
	this := OverrideTemplate{}
	return &this
}

// NewOverrideTemplateWithDefaults instantiates a new OverrideTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideTemplateWithDefaults() *OverrideTemplate {
	this := OverrideTemplate{}
	return &this
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *OverrideTemplate) GetTemplate() OverrideTemplateTemplate {
	if o == nil || isNil(o.Template) {
		var ret OverrideTemplateTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideTemplate) GetTemplateOk() (*OverrideTemplateTemplate, bool) {
	if o == nil || isNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *OverrideTemplate) HasTemplate() bool {
	if o != nil && !isNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given OverrideTemplateTemplate and assigns it to the Template field.
func (o *OverrideTemplate) SetTemplate(v OverrideTemplateTemplate) {
	o.Template = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OverrideTemplate) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideTemplate) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OverrideTemplate) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OverrideTemplate) SetType(v string) {
	o.Type = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *OverrideTemplate) GetWeight() int32 {
	if o == nil || isNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideTemplate) GetWeightOk() (*int32, bool) {
	if o == nil || isNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *OverrideTemplate) HasWeight() bool {
	if o != nil && !isNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *OverrideTemplate) SetWeight(v int32) {
	o.Weight = &v
}

func (o OverrideTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableOverrideTemplate struct {
	value *OverrideTemplate
	isSet bool
}

func (v NullableOverrideTemplate) Get() *OverrideTemplate {
	return v.value
}

func (v *NullableOverrideTemplate) Set(val *OverrideTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideTemplate(val *OverrideTemplate) *NullableOverrideTemplate {
	return &NullableOverrideTemplate{value: val, isSet: true}
}

func (v NullableOverrideTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


