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

// checks if the NamespaceFromRepostiroy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamespaceFromRepostiroy{}

// NamespaceFromRepostiroy namespace from git repository
type NamespaceFromRepostiroy struct {
	Path *File `json:"path,omitempty"`
	// name of the git repository
	Repository *string `json:"repository,omitempty"`
	// branch or tag in the git repository
	Revision *string `json:"revision,omitempty"`
}

// NewNamespaceFromRepostiroy instantiates a new NamespaceFromRepostiroy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamespaceFromRepostiroy() *NamespaceFromRepostiroy {
	this := NamespaceFromRepostiroy{}
	return &this
}

// NewNamespaceFromRepostiroyWithDefaults instantiates a new NamespaceFromRepostiroy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamespaceFromRepostiroyWithDefaults() *NamespaceFromRepostiroy {
	this := NamespaceFromRepostiroy{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *NamespaceFromRepostiroy) GetPath() File {
	if o == nil || isNil(o.Path) {
		var ret File
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceFromRepostiroy) GetPathOk() (*File, bool) {
	if o == nil || isNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *NamespaceFromRepostiroy) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given File and assigns it to the Path field.
func (o *NamespaceFromRepostiroy) SetPath(v File) {
	o.Path = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *NamespaceFromRepostiroy) GetRepository() string {
	if o == nil || isNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceFromRepostiroy) GetRepositoryOk() (*string, bool) {
	if o == nil || isNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *NamespaceFromRepostiroy) HasRepository() bool {
	if o != nil && !isNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *NamespaceFromRepostiroy) SetRepository(v string) {
	o.Repository = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *NamespaceFromRepostiroy) GetRevision() string {
	if o == nil || isNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceFromRepostiroy) GetRevisionOk() (*string, bool) {
	if o == nil || isNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *NamespaceFromRepostiroy) HasRevision() bool {
	if o != nil && !isNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *NamespaceFromRepostiroy) SetRevision(v string) {
	o.Revision = &v
}

func (o NamespaceFromRepostiroy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamespaceFromRepostiroy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !isNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	return toSerialize, nil
}

type NullableNamespaceFromRepostiroy struct {
	value *NamespaceFromRepostiroy
	isSet bool
}

func (v NullableNamespaceFromRepostiroy) Get() *NamespaceFromRepostiroy {
	return v.value
}

func (v *NullableNamespaceFromRepostiroy) Set(val *NamespaceFromRepostiroy) {
	v.value = val
	v.isSet = true
}

func (v NullableNamespaceFromRepostiroy) IsSet() bool {
	return v.isSet
}

func (v *NullableNamespaceFromRepostiroy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamespaceFromRepostiroy(val *NamespaceFromRepostiroy) *NullableNamespaceFromRepostiroy {
	return &NullableNamespaceFromRepostiroy{value: val, isSet: true}
}

func (v NullableNamespaceFromRepostiroy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamespaceFromRepostiroy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


