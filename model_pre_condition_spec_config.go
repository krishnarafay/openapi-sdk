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

// PreConditionSpecConfig - struct for PreConditionSpecConfig
type PreConditionSpecConfig struct {
	String *string
}

// stringAsPreConditionSpecConfig is a convenience function that returns string wrapped in PreConditionSpecConfig
func StringAsPreConditionSpecConfig(v *string) PreConditionSpecConfig {
	return PreConditionSpecConfig{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PreConditionSpecConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PreConditionSpecConfig)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PreConditionSpecConfig)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PreConditionSpecConfig) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PreConditionSpecConfig) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullablePreConditionSpecConfig struct {
	value *PreConditionSpecConfig
	isSet bool
}

func (v NullablePreConditionSpecConfig) Get() *PreConditionSpecConfig {
	return v.value
}

func (v *NullablePreConditionSpecConfig) Set(val *PreConditionSpecConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePreConditionSpecConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePreConditionSpecConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreConditionSpecConfig(val *PreConditionSpecConfig) *NullablePreConditionSpecConfig {
	return &NullablePreConditionSpecConfig{value: val, isSet: true}
}

func (v NullablePreConditionSpecConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreConditionSpecConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

