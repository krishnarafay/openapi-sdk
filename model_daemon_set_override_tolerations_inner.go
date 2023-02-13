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

// checks if the DaemonSetOverrideTolerationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DaemonSetOverrideTolerationsInner{}

// DaemonSetOverrideTolerationsInner struct for DaemonSetOverrideTolerationsInner
type DaemonSetOverrideTolerationsInner struct {
	Effect *string `json:"effect,omitempty"`
	Key *string `json:"key,omitempty"`
	Operator *string `json:"operator,omitempty"`
	TolerationSeconds *int64 `json:"tolerationSeconds,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewDaemonSetOverrideTolerationsInner instantiates a new DaemonSetOverrideTolerationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaemonSetOverrideTolerationsInner() *DaemonSetOverrideTolerationsInner {
	this := DaemonSetOverrideTolerationsInner{}
	return &this
}

// NewDaemonSetOverrideTolerationsInnerWithDefaults instantiates a new DaemonSetOverrideTolerationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaemonSetOverrideTolerationsInnerWithDefaults() *DaemonSetOverrideTolerationsInner {
	this := DaemonSetOverrideTolerationsInner{}
	return &this
}

// GetEffect returns the Effect field value if set, zero value otherwise.
func (o *DaemonSetOverrideTolerationsInner) GetEffect() string {
	if o == nil || isNil(o.Effect) {
		var ret string
		return ret
	}
	return *o.Effect
}

// GetEffectOk returns a tuple with the Effect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonSetOverrideTolerationsInner) GetEffectOk() (*string, bool) {
	if o == nil || isNil(o.Effect) {
		return nil, false
	}
	return o.Effect, true
}

// HasEffect returns a boolean if a field has been set.
func (o *DaemonSetOverrideTolerationsInner) HasEffect() bool {
	if o != nil && !isNil(o.Effect) {
		return true
	}

	return false
}

// SetEffect gets a reference to the given string and assigns it to the Effect field.
func (o *DaemonSetOverrideTolerationsInner) SetEffect(v string) {
	o.Effect = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *DaemonSetOverrideTolerationsInner) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonSetOverrideTolerationsInner) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *DaemonSetOverrideTolerationsInner) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *DaemonSetOverrideTolerationsInner) SetKey(v string) {
	o.Key = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *DaemonSetOverrideTolerationsInner) GetOperator() string {
	if o == nil || isNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonSetOverrideTolerationsInner) GetOperatorOk() (*string, bool) {
	if o == nil || isNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *DaemonSetOverrideTolerationsInner) HasOperator() bool {
	if o != nil && !isNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *DaemonSetOverrideTolerationsInner) SetOperator(v string) {
	o.Operator = &v
}

// GetTolerationSeconds returns the TolerationSeconds field value if set, zero value otherwise.
func (o *DaemonSetOverrideTolerationsInner) GetTolerationSeconds() int64 {
	if o == nil || isNil(o.TolerationSeconds) {
		var ret int64
		return ret
	}
	return *o.TolerationSeconds
}

// GetTolerationSecondsOk returns a tuple with the TolerationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonSetOverrideTolerationsInner) GetTolerationSecondsOk() (*int64, bool) {
	if o == nil || isNil(o.TolerationSeconds) {
		return nil, false
	}
	return o.TolerationSeconds, true
}

// HasTolerationSeconds returns a boolean if a field has been set.
func (o *DaemonSetOverrideTolerationsInner) HasTolerationSeconds() bool {
	if o != nil && !isNil(o.TolerationSeconds) {
		return true
	}

	return false
}

// SetTolerationSeconds gets a reference to the given int64 and assigns it to the TolerationSeconds field.
func (o *DaemonSetOverrideTolerationsInner) SetTolerationSeconds(v int64) {
	o.TolerationSeconds = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DaemonSetOverrideTolerationsInner) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaemonSetOverrideTolerationsInner) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DaemonSetOverrideTolerationsInner) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DaemonSetOverrideTolerationsInner) SetValue(v string) {
	o.Value = &v
}

func (o DaemonSetOverrideTolerationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DaemonSetOverrideTolerationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Effect) {
		toSerialize["effect"] = o.Effect
	}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !isNil(o.TolerationSeconds) {
		toSerialize["tolerationSeconds"] = o.TolerationSeconds
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableDaemonSetOverrideTolerationsInner struct {
	value *DaemonSetOverrideTolerationsInner
	isSet bool
}

func (v NullableDaemonSetOverrideTolerationsInner) Get() *DaemonSetOverrideTolerationsInner {
	return v.value
}

func (v *NullableDaemonSetOverrideTolerationsInner) Set(val *DaemonSetOverrideTolerationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDaemonSetOverrideTolerationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDaemonSetOverrideTolerationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaemonSetOverrideTolerationsInner(val *DaemonSetOverrideTolerationsInner) *NullableDaemonSetOverrideTolerationsInner {
	return &NullableDaemonSetOverrideTolerationsInner{value: val, isSet: true}
}

func (v NullableDaemonSetOverrideTolerationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaemonSetOverrideTolerationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

