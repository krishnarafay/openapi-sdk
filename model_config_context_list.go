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

// checks if the ConfigContextList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigContextList{}

// ConfigContextList List of config contexts
type ConfigContextList struct {
	// API Version of the config context list resource
	ApiVersion *string `json:"apiVersion,omitempty"`
	// List of the config context resources
	Items []ConfigContext `json:"items,omitempty"`
	// Kind of the config context list resource
	Kind *string `json:"kind,omitempty"`
	Metadata *ListMetadata `json:"metadata,omitempty"`
}

// NewConfigContextList instantiates a new ConfigContextList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigContextList() *ConfigContextList {
	this := ConfigContextList{}
	return &this
}

// NewConfigContextListWithDefaults instantiates a new ConfigContextList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigContextListWithDefaults() *ConfigContextList {
	this := ConfigContextList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ConfigContextList) GetApiVersion() string {
	if o == nil || isNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextList) GetApiVersionOk() (*string, bool) {
	if o == nil || isNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ConfigContextList) HasApiVersion() bool {
	if o != nil && !isNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ConfigContextList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ConfigContextList) GetItems() []ConfigContext {
	if o == nil || isNil(o.Items) {
		var ret []ConfigContext
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextList) GetItemsOk() ([]ConfigContext, bool) {
	if o == nil || isNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ConfigContextList) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ConfigContext and assigns it to the Items field.
func (o *ConfigContextList) SetItems(v []ConfigContext) {
	o.Items = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ConfigContextList) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextList) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ConfigContextList) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ConfigContextList) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ConfigContextList) GetMetadata() ListMetadata {
	if o == nil || isNil(o.Metadata) {
		var ret ListMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextList) GetMetadataOk() (*ListMetadata, bool) {
	if o == nil || isNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ConfigContextList) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ListMetadata and assigns it to the Metadata field.
func (o *ConfigContextList) SetMetadata(v ListMetadata) {
	o.Metadata = &v
}

func (o ConfigContextList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigContextList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: apiVersion is readOnly
	// skip: items is readOnly
	// skip: kind is readOnly
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableConfigContextList struct {
	value *ConfigContextList
	isSet bool
}

func (v NullableConfigContextList) Get() *ConfigContextList {
	return v.value
}

func (v *NullableConfigContextList) Set(val *ConfigContextList) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigContextList) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigContextList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigContextList(val *ConfigContextList) *NullableConfigContextList {
	return &NullableConfigContextList{value: val, isSet: true}
}

func (v NullableConfigContextList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigContextList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


