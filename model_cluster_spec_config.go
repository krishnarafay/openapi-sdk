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

// ClusterSpecConfig - struct for ClusterSpecConfig
type ClusterSpecConfig struct {
	ClusterSpecConfigOneOf *ClusterSpecConfigOneOf
	ClusterSpecConfigOneOf1 *ClusterSpecConfigOneOf1
	V3ImportedClusterSpecifications *V3ImportedClusterSpecifications
}

// ClusterSpecConfigOneOfAsClusterSpecConfig is a convenience function that returns ClusterSpecConfigOneOf wrapped in ClusterSpecConfig
func ClusterSpecConfigOneOfAsClusterSpecConfig(v *ClusterSpecConfigOneOf) ClusterSpecConfig {
	return ClusterSpecConfig{
		ClusterSpecConfigOneOf: v,
	}
}

// ClusterSpecConfigOneOf1AsClusterSpecConfig is a convenience function that returns ClusterSpecConfigOneOf1 wrapped in ClusterSpecConfig
func ClusterSpecConfigOneOf1AsClusterSpecConfig(v *ClusterSpecConfigOneOf1) ClusterSpecConfig {
	return ClusterSpecConfig{
		ClusterSpecConfigOneOf1: v,
	}
}

// V3ImportedClusterSpecificationsAsClusterSpecConfig is a convenience function that returns V3ImportedClusterSpecifications wrapped in ClusterSpecConfig
func V3ImportedClusterSpecificationsAsClusterSpecConfig(v *V3ImportedClusterSpecifications) ClusterSpecConfig {
	return ClusterSpecConfig{
		V3ImportedClusterSpecifications: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ClusterSpecConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ClusterSpecConfigOneOf
	err = newStrictDecoder(data).Decode(&dst.ClusterSpecConfigOneOf)
	if err == nil {
		jsonClusterSpecConfigOneOf, _ := json.Marshal(dst.ClusterSpecConfigOneOf)
		if string(jsonClusterSpecConfigOneOf) == "{}" { // empty struct
			dst.ClusterSpecConfigOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ClusterSpecConfigOneOf = nil
	}

	// try to unmarshal data into ClusterSpecConfigOneOf1
	err = newStrictDecoder(data).Decode(&dst.ClusterSpecConfigOneOf1)
	if err == nil {
		jsonClusterSpecConfigOneOf1, _ := json.Marshal(dst.ClusterSpecConfigOneOf1)
		if string(jsonClusterSpecConfigOneOf1) == "{}" { // empty struct
			dst.ClusterSpecConfigOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.ClusterSpecConfigOneOf1 = nil
	}

	// try to unmarshal data into V3ImportedClusterSpecifications
	err = newStrictDecoder(data).Decode(&dst.V3ImportedClusterSpecifications)
	if err == nil {
		jsonV3ImportedClusterSpecifications, _ := json.Marshal(dst.V3ImportedClusterSpecifications)
		if string(jsonV3ImportedClusterSpecifications) == "{}" { // empty struct
			dst.V3ImportedClusterSpecifications = nil
		} else {
			match++
		}
	} else {
		dst.V3ImportedClusterSpecifications = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ClusterSpecConfigOneOf = nil
		dst.ClusterSpecConfigOneOf1 = nil
		dst.V3ImportedClusterSpecifications = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ClusterSpecConfig)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ClusterSpecConfig)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ClusterSpecConfig) MarshalJSON() ([]byte, error) {
	if src.ClusterSpecConfigOneOf != nil {
		return json.Marshal(&src.ClusterSpecConfigOneOf)
	}

	if src.ClusterSpecConfigOneOf1 != nil {
		return json.Marshal(&src.ClusterSpecConfigOneOf1)
	}

	if src.V3ImportedClusterSpecifications != nil {
		return json.Marshal(&src.V3ImportedClusterSpecifications)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ClusterSpecConfig) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ClusterSpecConfigOneOf != nil {
		return obj.ClusterSpecConfigOneOf
	}

	if obj.ClusterSpecConfigOneOf1 != nil {
		return obj.ClusterSpecConfigOneOf1
	}

	if obj.V3ImportedClusterSpecifications != nil {
		return obj.V3ImportedClusterSpecifications
	}

	// all schemas are nil
	return nil
}

type NullableClusterSpecConfig struct {
	value *ClusterSpecConfig
	isSet bool
}

func (v NullableClusterSpecConfig) Get() *ClusterSpecConfig {
	return v.value
}

func (v *NullableClusterSpecConfig) Set(val *ClusterSpecConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterSpecConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterSpecConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterSpecConfig(val *ClusterSpecConfig) *NullableClusterSpecConfig {
	return &NullableClusterSpecConfig{value: val, isSet: true}
}

func (v NullableClusterSpecConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterSpecConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


