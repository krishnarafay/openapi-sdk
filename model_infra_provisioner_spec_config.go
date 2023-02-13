/*
Kubernetes Operations APIs

Kubernetes Operations APIs

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rafaysdkgov3

import (
	"encoding/json"
	"fmt"
)

// InfraProvisionerSpecConfig - struct for InfraProvisionerSpecConfig
type InfraProvisionerSpecConfig struct {
	TerrafromConfig *TerrafromConfig
}

// TerrafromConfigAsInfraProvisionerSpecConfig is a convenience function that returns TerrafromConfig wrapped in InfraProvisionerSpecConfig
func TerrafromConfigAsInfraProvisionerSpecConfig(v *TerrafromConfig) InfraProvisionerSpecConfig {
	return InfraProvisionerSpecConfig{
		TerrafromConfig: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InfraProvisionerSpecConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TerrafromConfig
	err = newStrictDecoder(data).Decode(&dst.TerrafromConfig)
	if err == nil {
		jsonTerrafromConfig, _ := json.Marshal(dst.TerrafromConfig)
		if string(jsonTerrafromConfig) == "{}" { // empty struct
			dst.TerrafromConfig = nil
		} else {
			match++
		}
	} else {
		dst.TerrafromConfig = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TerrafromConfig = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InfraProvisionerSpecConfig)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InfraProvisionerSpecConfig)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InfraProvisionerSpecConfig) MarshalJSON() ([]byte, error) {
	if src.TerrafromConfig != nil {
		return json.Marshal(&src.TerrafromConfig)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InfraProvisionerSpecConfig) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.TerrafromConfig != nil {
		return obj.TerrafromConfig
	}

	// all schemas are nil
	return nil
}

type NullableInfraProvisionerSpecConfig struct {
	value *InfraProvisionerSpecConfig
	isSet bool
}

func (v NullableInfraProvisionerSpecConfig) Get() *InfraProvisionerSpecConfig {
	return v.value
}

func (v *NullableInfraProvisionerSpecConfig) Set(val *InfraProvisionerSpecConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraProvisionerSpecConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraProvisionerSpecConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraProvisionerSpecConfig(val *InfraProvisionerSpecConfig) *NullableInfraProvisionerSpecConfig {
	return &NullableInfraProvisionerSpecConfig{value: val, isSet: true}
}

func (v NullableInfraProvisionerSpecConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraProvisionerSpecConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

