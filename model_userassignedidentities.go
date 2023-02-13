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

// checks if the Userassignedidentities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Userassignedidentities{}

// Userassignedidentities struct for Userassignedidentities
type Userassignedidentities struct {
	// The binding selector to use for the AzureIdentityBinding resource.
	BindingSelector *string `json:"bindingSelector,omitempty"`
	Identity *Identity1 `json:"identity,omitempty"`
	// The name of the pod identity.
	Name *string `json:"name,omitempty"`
	// The namespace of the pod identity.
	Namespace *string `json:"namespace,omitempty"`
}

// NewUserassignedidentities instantiates a new Userassignedidentities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserassignedidentities() *Userassignedidentities {
	this := Userassignedidentities{}
	return &this
}

// NewUserassignedidentitiesWithDefaults instantiates a new Userassignedidentities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserassignedidentitiesWithDefaults() *Userassignedidentities {
	this := Userassignedidentities{}
	return &this
}

// GetBindingSelector returns the BindingSelector field value if set, zero value otherwise.
func (o *Userassignedidentities) GetBindingSelector() string {
	if o == nil || isNil(o.BindingSelector) {
		var ret string
		return ret
	}
	return *o.BindingSelector
}

// GetBindingSelectorOk returns a tuple with the BindingSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userassignedidentities) GetBindingSelectorOk() (*string, bool) {
	if o == nil || isNil(o.BindingSelector) {
		return nil, false
	}
	return o.BindingSelector, true
}

// HasBindingSelector returns a boolean if a field has been set.
func (o *Userassignedidentities) HasBindingSelector() bool {
	if o != nil && !isNil(o.BindingSelector) {
		return true
	}

	return false
}

// SetBindingSelector gets a reference to the given string and assigns it to the BindingSelector field.
func (o *Userassignedidentities) SetBindingSelector(v string) {
	o.BindingSelector = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *Userassignedidentities) GetIdentity() Identity1 {
	if o == nil || isNil(o.Identity) {
		var ret Identity1
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userassignedidentities) GetIdentityOk() (*Identity1, bool) {
	if o == nil || isNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *Userassignedidentities) HasIdentity() bool {
	if o != nil && !isNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given Identity1 and assigns it to the Identity field.
func (o *Userassignedidentities) SetIdentity(v Identity1) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Userassignedidentities) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userassignedidentities) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Userassignedidentities) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Userassignedidentities) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *Userassignedidentities) GetNamespace() string {
	if o == nil || isNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userassignedidentities) GetNamespaceOk() (*string, bool) {
	if o == nil || isNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *Userassignedidentities) HasNamespace() bool {
	if o != nil && !isNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *Userassignedidentities) SetNamespace(v string) {
	o.Namespace = &v
}

func (o Userassignedidentities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Userassignedidentities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BindingSelector) {
		toSerialize["bindingSelector"] = o.BindingSelector
	}
	if !isNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return toSerialize, nil
}

type NullableUserassignedidentities struct {
	value *Userassignedidentities
	isSet bool
}

func (v NullableUserassignedidentities) Get() *Userassignedidentities {
	return v.value
}

func (v *NullableUserassignedidentities) Set(val *Userassignedidentities) {
	v.value = val
	v.isSet = true
}

func (v NullableUserassignedidentities) IsSet() bool {
	return v.isSet
}

func (v *NullableUserassignedidentities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserassignedidentities(val *Userassignedidentities) *NullableUserassignedidentities {
	return &NullableUserassignedidentities{value: val, isSet: true}
}

func (v NullableUserassignedidentities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserassignedidentities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


