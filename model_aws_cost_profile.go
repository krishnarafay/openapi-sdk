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

// checks if the AwsCostProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsCostProfile{}

// AwsCostProfile Aws custom cost profile params
type AwsCostProfile struct {
	AwsCredentials *AwsCredsCostProfile `json:"awsCredentials,omitempty"`
	CurIntegration *AwsCurIntegration `json:"curIntegration,omitempty"`
	SpotIntegration *AwsSpotIntegration `json:"spotIntegration,omitempty"`
}

// NewAwsCostProfile instantiates a new AwsCostProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsCostProfile() *AwsCostProfile {
	this := AwsCostProfile{}
	return &this
}

// NewAwsCostProfileWithDefaults instantiates a new AwsCostProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsCostProfileWithDefaults() *AwsCostProfile {
	this := AwsCostProfile{}
	return &this
}

// GetAwsCredentials returns the AwsCredentials field value if set, zero value otherwise.
func (o *AwsCostProfile) GetAwsCredentials() AwsCredsCostProfile {
	if o == nil || isNil(o.AwsCredentials) {
		var ret AwsCredsCostProfile
		return ret
	}
	return *o.AwsCredentials
}

// GetAwsCredentialsOk returns a tuple with the AwsCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCostProfile) GetAwsCredentialsOk() (*AwsCredsCostProfile, bool) {
	if o == nil || isNil(o.AwsCredentials) {
		return nil, false
	}
	return o.AwsCredentials, true
}

// HasAwsCredentials returns a boolean if a field has been set.
func (o *AwsCostProfile) HasAwsCredentials() bool {
	if o != nil && !isNil(o.AwsCredentials) {
		return true
	}

	return false
}

// SetAwsCredentials gets a reference to the given AwsCredsCostProfile and assigns it to the AwsCredentials field.
func (o *AwsCostProfile) SetAwsCredentials(v AwsCredsCostProfile) {
	o.AwsCredentials = &v
}

// GetCurIntegration returns the CurIntegration field value if set, zero value otherwise.
func (o *AwsCostProfile) GetCurIntegration() AwsCurIntegration {
	if o == nil || isNil(o.CurIntegration) {
		var ret AwsCurIntegration
		return ret
	}
	return *o.CurIntegration
}

// GetCurIntegrationOk returns a tuple with the CurIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCostProfile) GetCurIntegrationOk() (*AwsCurIntegration, bool) {
	if o == nil || isNil(o.CurIntegration) {
		return nil, false
	}
	return o.CurIntegration, true
}

// HasCurIntegration returns a boolean if a field has been set.
func (o *AwsCostProfile) HasCurIntegration() bool {
	if o != nil && !isNil(o.CurIntegration) {
		return true
	}

	return false
}

// SetCurIntegration gets a reference to the given AwsCurIntegration and assigns it to the CurIntegration field.
func (o *AwsCostProfile) SetCurIntegration(v AwsCurIntegration) {
	o.CurIntegration = &v
}

// GetSpotIntegration returns the SpotIntegration field value if set, zero value otherwise.
func (o *AwsCostProfile) GetSpotIntegration() AwsSpotIntegration {
	if o == nil || isNil(o.SpotIntegration) {
		var ret AwsSpotIntegration
		return ret
	}
	return *o.SpotIntegration
}

// GetSpotIntegrationOk returns a tuple with the SpotIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCostProfile) GetSpotIntegrationOk() (*AwsSpotIntegration, bool) {
	if o == nil || isNil(o.SpotIntegration) {
		return nil, false
	}
	return o.SpotIntegration, true
}

// HasSpotIntegration returns a boolean if a field has been set.
func (o *AwsCostProfile) HasSpotIntegration() bool {
	if o != nil && !isNil(o.SpotIntegration) {
		return true
	}

	return false
}

// SetSpotIntegration gets a reference to the given AwsSpotIntegration and assigns it to the SpotIntegration field.
func (o *AwsCostProfile) SetSpotIntegration(v AwsSpotIntegration) {
	o.SpotIntegration = &v
}

func (o AwsCostProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsCostProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AwsCredentials) {
		toSerialize["awsCredentials"] = o.AwsCredentials
	}
	if !isNil(o.CurIntegration) {
		toSerialize["curIntegration"] = o.CurIntegration
	}
	if !isNil(o.SpotIntegration) {
		toSerialize["spotIntegration"] = o.SpotIntegration
	}
	return toSerialize, nil
}

type NullableAwsCostProfile struct {
	value *AwsCostProfile
	isSet bool
}

func (v NullableAwsCostProfile) Get() *AwsCostProfile {
	return v.value
}

func (v *NullableAwsCostProfile) Set(val *AwsCostProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsCostProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsCostProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsCostProfile(val *AwsCostProfile) *NullableAwsCostProfile {
	return &NullableAwsCostProfile{value: val, isSet: true}
}

func (v NullableAwsCostProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsCostProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

