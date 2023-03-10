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

// RepositorySpecCredentials - struct for RepositorySpecCredentials
type RepositorySpecCredentials struct {
	ContainerRegistrySpecCredentialsOneOf *ContainerRegistrySpecCredentialsOneOf
	RepositorySpecCredentialsOneOf *RepositorySpecCredentialsOneOf
}

// ContainerRegistrySpecCredentialsOneOfAsRepositorySpecCredentials is a convenience function that returns ContainerRegistrySpecCredentialsOneOf wrapped in RepositorySpecCredentials
func ContainerRegistrySpecCredentialsOneOfAsRepositorySpecCredentials(v *ContainerRegistrySpecCredentialsOneOf) RepositorySpecCredentials {
	return RepositorySpecCredentials{
		ContainerRegistrySpecCredentialsOneOf: v,
	}
}

// RepositorySpecCredentialsOneOfAsRepositorySpecCredentials is a convenience function that returns RepositorySpecCredentialsOneOf wrapped in RepositorySpecCredentials
func RepositorySpecCredentialsOneOfAsRepositorySpecCredentials(v *RepositorySpecCredentialsOneOf) RepositorySpecCredentials {
	return RepositorySpecCredentials{
		RepositorySpecCredentialsOneOf: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RepositorySpecCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContainerRegistrySpecCredentialsOneOf
	err = newStrictDecoder(data).Decode(&dst.ContainerRegistrySpecCredentialsOneOf)
	if err == nil {
		jsonContainerRegistrySpecCredentialsOneOf, _ := json.Marshal(dst.ContainerRegistrySpecCredentialsOneOf)
		if string(jsonContainerRegistrySpecCredentialsOneOf) == "{}" { // empty struct
			dst.ContainerRegistrySpecCredentialsOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ContainerRegistrySpecCredentialsOneOf = nil
	}

	// try to unmarshal data into RepositorySpecCredentialsOneOf
	err = newStrictDecoder(data).Decode(&dst.RepositorySpecCredentialsOneOf)
	if err == nil {
		jsonRepositorySpecCredentialsOneOf, _ := json.Marshal(dst.RepositorySpecCredentialsOneOf)
		if string(jsonRepositorySpecCredentialsOneOf) == "{}" { // empty struct
			dst.RepositorySpecCredentialsOneOf = nil
		} else {
			match++
		}
	} else {
		dst.RepositorySpecCredentialsOneOf = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ContainerRegistrySpecCredentialsOneOf = nil
		dst.RepositorySpecCredentialsOneOf = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RepositorySpecCredentials)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RepositorySpecCredentials)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RepositorySpecCredentials) MarshalJSON() ([]byte, error) {
	if src.ContainerRegistrySpecCredentialsOneOf != nil {
		return json.Marshal(&src.ContainerRegistrySpecCredentialsOneOf)
	}

	if src.RepositorySpecCredentialsOneOf != nil {
		return json.Marshal(&src.RepositorySpecCredentialsOneOf)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RepositorySpecCredentials) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ContainerRegistrySpecCredentialsOneOf != nil {
		return obj.ContainerRegistrySpecCredentialsOneOf
	}

	if obj.RepositorySpecCredentialsOneOf != nil {
		return obj.RepositorySpecCredentialsOneOf
	}

	// all schemas are nil
	return nil
}

type NullableRepositorySpecCredentials struct {
	value *RepositorySpecCredentials
	isSet bool
}

func (v NullableRepositorySpecCredentials) Get() *RepositorySpecCredentials {
	return v.value
}

func (v *NullableRepositorySpecCredentials) Set(val *RepositorySpecCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositorySpecCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositorySpecCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositorySpecCredentials(val *RepositorySpecCredentials) *NullableRepositorySpecCredentials {
	return &NullableRepositorySpecCredentials{value: val, isSet: true}
}

func (v NullableRepositorySpecCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositorySpecCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


