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

// checks if the Sku type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sku{}

// Sku struct for Sku
type Sku struct {
	// The name of a managed cluster SKU.
	Name *string `json:"name,omitempty"`
	// If not specified, the default is Free. See uptime SLA for more details. Valid values are Paid, Free.
	Tier *string `json:"tier,omitempty"`
}

// NewSku instantiates a new Sku object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSku() *Sku {
	this := Sku{}
	var name string = "Basic"
	this.Name = &name
	var tier string = "Free"
	this.Tier = &tier
	return &this
}

// NewSkuWithDefaults instantiates a new Sku object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuWithDefaults() *Sku {
	this := Sku{}
	var name string = "Basic"
	this.Name = &name
	var tier string = "Free"
	this.Tier = &tier
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Sku) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sku) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Sku) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Sku) SetName(v string) {
	o.Name = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *Sku) GetTier() string {
	if o == nil || isNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sku) GetTierOk() (*string, bool) {
	if o == nil || isNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *Sku) HasTier() bool {
	if o != nil && !isNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *Sku) SetTier(v string) {
	o.Tier = &v
}

func (o Sku) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sku) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}
	return toSerialize, nil
}

type NullableSku struct {
	value *Sku
	isSet bool
}

func (v NullableSku) Get() *Sku {
	return v.value
}

func (v *NullableSku) Set(val *Sku) {
	v.value = val
	v.isSet = true
}

func (v NullableSku) IsSet() bool {
	return v.isSet
}

func (v *NullableSku) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSku(val *Sku) *NullableSku {
	return &NullableSku{value: val, isSet: true}
}

func (v NullableSku) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSku) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

