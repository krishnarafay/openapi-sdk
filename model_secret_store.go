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

// checks if the SecretStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretStore{}

// SecretStore secret store specification
type SecretStore struct {
	// API Version of the secret store resource
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind of the secret store resource
	Kind *string `json:"kind,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
	Spec *SecretStoreSpec `json:"spec,omitempty"`
	Status *Status `json:"status,omitempty"`
}

// NewSecretStore instantiates a new SecretStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretStore() *SecretStore {
	this := SecretStore{}
	var apiVersion string = "integrations.k8smgmt.io/v3"
	this.ApiVersion = &apiVersion
	var kind string = "SecretStore"
	this.Kind = &kind
	return &this
}

// NewSecretStoreWithDefaults instantiates a new SecretStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretStoreWithDefaults() *SecretStore {
	this := SecretStore{}
	var apiVersion string = "integrations.k8smgmt.io/v3"
	this.ApiVersion = &apiVersion
	var kind string = "SecretStore"
	this.Kind = &kind
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *SecretStore) GetApiVersion() string {
	if o == nil || isNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretStore) GetApiVersionOk() (*string, bool) {
	if o == nil || isNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *SecretStore) HasApiVersion() bool {
	if o != nil && !isNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *SecretStore) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *SecretStore) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretStore) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *SecretStore) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *SecretStore) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SecretStore) GetMetadata() Metadata {
	if o == nil || isNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretStore) GetMetadataOk() (*Metadata, bool) {
	if o == nil || isNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SecretStore) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *SecretStore) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *SecretStore) GetSpec() SecretStoreSpec {
	if o == nil || isNil(o.Spec) {
		var ret SecretStoreSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretStore) GetSpecOk() (*SecretStoreSpec, bool) {
	if o == nil || isNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *SecretStore) HasSpec() bool {
	if o != nil && !isNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given SecretStoreSpec and assigns it to the Spec field.
func (o *SecretStore) SetSpec(v SecretStoreSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SecretStore) GetStatus() Status {
	if o == nil || isNil(o.Status) {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretStore) GetStatusOk() (*Status, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SecretStore) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *SecretStore) SetStatus(v Status) {
	o.Status = &v
}

func (o SecretStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretStore) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableSecretStore struct {
	value *SecretStore
	isSet bool
}

func (v NullableSecretStore) Get() *SecretStore {
	return v.value
}

func (v *NullableSecretStore) Set(val *SecretStore) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretStore) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretStore(val *SecretStore) *NullableSecretStore {
	return &NullableSecretStore{value: val, isSet: true}
}

func (v NullableSecretStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

