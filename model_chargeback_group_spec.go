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

// checks if the ChargebackGroupSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargebackGroupSpec{}

// ChargebackGroupSpec profile specification
type ChargebackGroupSpec struct {
	Aggregate *ChargebackAggregate `json:"aggregate,omitempty"`
	// exclusions
	Exclusions []ChargebackFilter `json:"exclusions,omitempty"`
	// inclusions
	Inclusions []ChargebackFilter `json:"inclusions,omitempty"`
	// type
	Type *string `json:"type,omitempty"`
}

// NewChargebackGroupSpec instantiates a new ChargebackGroupSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargebackGroupSpec() *ChargebackGroupSpec {
	this := ChargebackGroupSpec{}
	return &this
}

// NewChargebackGroupSpecWithDefaults instantiates a new ChargebackGroupSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargebackGroupSpecWithDefaults() *ChargebackGroupSpec {
	this := ChargebackGroupSpec{}
	return &this
}

// GetAggregate returns the Aggregate field value if set, zero value otherwise.
func (o *ChargebackGroupSpec) GetAggregate() ChargebackAggregate {
	if o == nil || isNil(o.Aggregate) {
		var ret ChargebackAggregate
		return ret
	}
	return *o.Aggregate
}

// GetAggregateOk returns a tuple with the Aggregate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargebackGroupSpec) GetAggregateOk() (*ChargebackAggregate, bool) {
	if o == nil || isNil(o.Aggregate) {
		return nil, false
	}
	return o.Aggregate, true
}

// HasAggregate returns a boolean if a field has been set.
func (o *ChargebackGroupSpec) HasAggregate() bool {
	if o != nil && !isNil(o.Aggregate) {
		return true
	}

	return false
}

// SetAggregate gets a reference to the given ChargebackAggregate and assigns it to the Aggregate field.
func (o *ChargebackGroupSpec) SetAggregate(v ChargebackAggregate) {
	o.Aggregate = &v
}

// GetExclusions returns the Exclusions field value if set, zero value otherwise.
func (o *ChargebackGroupSpec) GetExclusions() []ChargebackFilter {
	if o == nil || isNil(o.Exclusions) {
		var ret []ChargebackFilter
		return ret
	}
	return o.Exclusions
}

// GetExclusionsOk returns a tuple with the Exclusions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargebackGroupSpec) GetExclusionsOk() ([]ChargebackFilter, bool) {
	if o == nil || isNil(o.Exclusions) {
		return nil, false
	}
	return o.Exclusions, true
}

// HasExclusions returns a boolean if a field has been set.
func (o *ChargebackGroupSpec) HasExclusions() bool {
	if o != nil && !isNil(o.Exclusions) {
		return true
	}

	return false
}

// SetExclusions gets a reference to the given []ChargebackFilter and assigns it to the Exclusions field.
func (o *ChargebackGroupSpec) SetExclusions(v []ChargebackFilter) {
	o.Exclusions = v
}

// GetInclusions returns the Inclusions field value if set, zero value otherwise.
func (o *ChargebackGroupSpec) GetInclusions() []ChargebackFilter {
	if o == nil || isNil(o.Inclusions) {
		var ret []ChargebackFilter
		return ret
	}
	return o.Inclusions
}

// GetInclusionsOk returns a tuple with the Inclusions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargebackGroupSpec) GetInclusionsOk() ([]ChargebackFilter, bool) {
	if o == nil || isNil(o.Inclusions) {
		return nil, false
	}
	return o.Inclusions, true
}

// HasInclusions returns a boolean if a field has been set.
func (o *ChargebackGroupSpec) HasInclusions() bool {
	if o != nil && !isNil(o.Inclusions) {
		return true
	}

	return false
}

// SetInclusions gets a reference to the given []ChargebackFilter and assigns it to the Inclusions field.
func (o *ChargebackGroupSpec) SetInclusions(v []ChargebackFilter) {
	o.Inclusions = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ChargebackGroupSpec) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargebackGroupSpec) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ChargebackGroupSpec) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ChargebackGroupSpec) SetType(v string) {
	o.Type = &v
}

func (o ChargebackGroupSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargebackGroupSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Aggregate) {
		toSerialize["aggregate"] = o.Aggregate
	}
	if !isNil(o.Exclusions) {
		toSerialize["exclusions"] = o.Exclusions
	}
	if !isNil(o.Inclusions) {
		toSerialize["inclusions"] = o.Inclusions
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableChargebackGroupSpec struct {
	value *ChargebackGroupSpec
	isSet bool
}

func (v NullableChargebackGroupSpec) Get() *ChargebackGroupSpec {
	return v.value
}

func (v *NullableChargebackGroupSpec) Set(val *ChargebackGroupSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableChargebackGroupSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableChargebackGroupSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargebackGroupSpec(val *ChargebackGroupSpec) *NullableChargebackGroupSpec {
	return &NullableChargebackGroupSpec{value: val, isSet: true}
}

func (v NullableChargebackGroupSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargebackGroupSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

