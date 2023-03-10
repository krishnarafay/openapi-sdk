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

// checks if the BlueprintList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlueprintList{}

// BlueprintList Blueprint list
type BlueprintList struct {
	// API Version of the blueprint list resource
	ApiVersion *string `json:"apiVersion,omitempty"`
	// List of the blueprint resources
	Items []Blueprint `json:"items,omitempty"`
	// Kind of the blueprint list resource
	Kind *string `json:"kind,omitempty"`
	Metadata *ListMetadata `json:"metadata,omitempty"`
}

// NewBlueprintList instantiates a new BlueprintList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintList() *BlueprintList {
	this := BlueprintList{}
	return &this
}

// NewBlueprintListWithDefaults instantiates a new BlueprintList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintListWithDefaults() *BlueprintList {
	this := BlueprintList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *BlueprintList) GetApiVersion() string {
	if o == nil || isNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintList) GetApiVersionOk() (*string, bool) {
	if o == nil || isNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *BlueprintList) HasApiVersion() bool {
	if o != nil && !isNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *BlueprintList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BlueprintList) GetItems() []Blueprint {
	if o == nil || isNil(o.Items) {
		var ret []Blueprint
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintList) GetItemsOk() ([]Blueprint, bool) {
	if o == nil || isNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BlueprintList) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Blueprint and assigns it to the Items field.
func (o *BlueprintList) SetItems(v []Blueprint) {
	o.Items = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *BlueprintList) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintList) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *BlueprintList) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *BlueprintList) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *BlueprintList) GetMetadata() ListMetadata {
	if o == nil || isNil(o.Metadata) {
		var ret ListMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintList) GetMetadataOk() (*ListMetadata, bool) {
	if o == nil || isNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BlueprintList) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ListMetadata and assigns it to the Metadata field.
func (o *BlueprintList) SetMetadata(v ListMetadata) {
	o.Metadata = &v
}

func (o BlueprintList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlueprintList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: apiVersion is readOnly
	// skip: items is readOnly
	// skip: kind is readOnly
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableBlueprintList struct {
	value *BlueprintList
	isSet bool
}

func (v NullableBlueprintList) Get() *BlueprintList {
	return v.value
}

func (v *NullableBlueprintList) Set(val *BlueprintList) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintList) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintList(val *BlueprintList) *NullableBlueprintList {
	return &NullableBlueprintList{value: val, isSet: true}
}

func (v NullableBlueprintList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


