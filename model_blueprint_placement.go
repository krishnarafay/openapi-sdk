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

// checks if the BlueprintPlacement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlueprintPlacement{}

// BlueprintPlacement blueprint fleet placement
type BlueprintPlacement struct {
	// Bool value of autoPublish
	AutoPublish *bool `json:"autoPublish,omitempty"`
	// Array of Fleet value
	FleetValues []string `json:"fleetValues,omitempty"`
}

// NewBlueprintPlacement instantiates a new BlueprintPlacement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintPlacement() *BlueprintPlacement {
	this := BlueprintPlacement{}
	return &this
}

// NewBlueprintPlacementWithDefaults instantiates a new BlueprintPlacement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintPlacementWithDefaults() *BlueprintPlacement {
	this := BlueprintPlacement{}
	return &this
}

// GetAutoPublish returns the AutoPublish field value if set, zero value otherwise.
func (o *BlueprintPlacement) GetAutoPublish() bool {
	if o == nil || isNil(o.AutoPublish) {
		var ret bool
		return ret
	}
	return *o.AutoPublish
}

// GetAutoPublishOk returns a tuple with the AutoPublish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintPlacement) GetAutoPublishOk() (*bool, bool) {
	if o == nil || isNil(o.AutoPublish) {
		return nil, false
	}
	return o.AutoPublish, true
}

// HasAutoPublish returns a boolean if a field has been set.
func (o *BlueprintPlacement) HasAutoPublish() bool {
	if o != nil && !isNil(o.AutoPublish) {
		return true
	}

	return false
}

// SetAutoPublish gets a reference to the given bool and assigns it to the AutoPublish field.
func (o *BlueprintPlacement) SetAutoPublish(v bool) {
	o.AutoPublish = &v
}

// GetFleetValues returns the FleetValues field value if set, zero value otherwise.
func (o *BlueprintPlacement) GetFleetValues() []string {
	if o == nil || isNil(o.FleetValues) {
		var ret []string
		return ret
	}
	return o.FleetValues
}

// GetFleetValuesOk returns a tuple with the FleetValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintPlacement) GetFleetValuesOk() ([]string, bool) {
	if o == nil || isNil(o.FleetValues) {
		return nil, false
	}
	return o.FleetValues, true
}

// HasFleetValues returns a boolean if a field has been set.
func (o *BlueprintPlacement) HasFleetValues() bool {
	if o != nil && !isNil(o.FleetValues) {
		return true
	}

	return false
}

// SetFleetValues gets a reference to the given []string and assigns it to the FleetValues field.
func (o *BlueprintPlacement) SetFleetValues(v []string) {
	o.FleetValues = v
}

func (o BlueprintPlacement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlueprintPlacement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AutoPublish) {
		toSerialize["autoPublish"] = o.AutoPublish
	}
	if !isNil(o.FleetValues) {
		toSerialize["fleetValues"] = o.FleetValues
	}
	return toSerialize, nil
}

type NullableBlueprintPlacement struct {
	value *BlueprintPlacement
	isSet bool
}

func (v NullableBlueprintPlacement) Get() *BlueprintPlacement {
	return v.value
}

func (v *NullableBlueprintPlacement) Set(val *BlueprintPlacement) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintPlacement) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintPlacement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintPlacement(val *BlueprintPlacement) *NullableBlueprintPlacement {
	return &NullableBlueprintPlacement{value: val, isSet: true}
}

func (v NullableBlueprintPlacement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintPlacement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


