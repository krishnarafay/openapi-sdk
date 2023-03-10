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

// checks if the AksV3ConfigObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AksV3ConfigObject{}

// AksV3ConfigObject struct for AksV3ConfigObject
type AksV3ConfigObject struct {
	ApiVersion *string `json:"apiVersion,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
	Spec *AksV3Spec `json:"spec,omitempty"`
}

// NewAksV3ConfigObject instantiates a new AksV3ConfigObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAksV3ConfigObject() *AksV3ConfigObject {
	this := AksV3ConfigObject{}
	return &this
}

// NewAksV3ConfigObjectWithDefaults instantiates a new AksV3ConfigObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAksV3ConfigObjectWithDefaults() *AksV3ConfigObject {
	this := AksV3ConfigObject{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *AksV3ConfigObject) GetApiVersion() string {
	if o == nil || isNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AksV3ConfigObject) GetApiVersionOk() (*string, bool) {
	if o == nil || isNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *AksV3ConfigObject) HasApiVersion() bool {
	if o != nil && !isNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *AksV3ConfigObject) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *AksV3ConfigObject) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AksV3ConfigObject) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *AksV3ConfigObject) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *AksV3ConfigObject) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AksV3ConfigObject) GetMetadata() Metadata {
	if o == nil || isNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AksV3ConfigObject) GetMetadataOk() (*Metadata, bool) {
	if o == nil || isNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AksV3ConfigObject) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *AksV3ConfigObject) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *AksV3ConfigObject) GetSpec() AksV3Spec {
	if o == nil || isNil(o.Spec) {
		var ret AksV3Spec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AksV3ConfigObject) GetSpecOk() (*AksV3Spec, bool) {
	if o == nil || isNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *AksV3ConfigObject) HasSpec() bool {
	if o != nil && !isNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given AksV3Spec and assigns it to the Spec field.
func (o *AksV3ConfigObject) SetSpec(v AksV3Spec) {
	o.Spec = &v
}

func (o AksV3ConfigObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AksV3ConfigObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	return toSerialize, nil
}

type NullableAksV3ConfigObject struct {
	value *AksV3ConfigObject
	isSet bool
}

func (v NullableAksV3ConfigObject) Get() *AksV3ConfigObject {
	return v.value
}

func (v *NullableAksV3ConfigObject) Set(val *AksV3ConfigObject) {
	v.value = val
	v.isSet = true
}

func (v NullableAksV3ConfigObject) IsSet() bool {
	return v.isSet
}

func (v *NullableAksV3ConfigObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAksV3ConfigObject(val *AksV3ConfigObject) *NullableAksV3ConfigObject {
	return &NullableAksV3ConfigObject{value: val, isSet: true}
}

func (v NullableAksV3ConfigObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAksV3ConfigObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


