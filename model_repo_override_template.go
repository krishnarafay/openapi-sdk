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

// checks if the RepoOverrideTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepoOverrideTemplate{}

// RepoOverrideTemplate repo based override template definition
type RepoOverrideTemplate struct {
	// paths in the repository containing the override template
	Paths []File `json:"paths"`
	// git repository containing the override template
	Repository string `json:"repository"`
	// branch or tag in the repository
	Revision string `json:"revision"`
}

// NewRepoOverrideTemplate instantiates a new RepoOverrideTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepoOverrideTemplate(paths []File, repository string, revision string) *RepoOverrideTemplate {
	this := RepoOverrideTemplate{}
	this.Paths = paths
	this.Repository = repository
	this.Revision = revision
	return &this
}

// NewRepoOverrideTemplateWithDefaults instantiates a new RepoOverrideTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepoOverrideTemplateWithDefaults() *RepoOverrideTemplate {
	this := RepoOverrideTemplate{}
	return &this
}

// GetPaths returns the Paths field value
func (o *RepoOverrideTemplate) GetPaths() []File {
	if o == nil {
		var ret []File
		return ret
	}

	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value
// and a boolean to check if the value has been set.
func (o *RepoOverrideTemplate) GetPathsOk() ([]File, bool) {
	if o == nil {
		return nil, false
	}
	return o.Paths, true
}

// SetPaths sets field value
func (o *RepoOverrideTemplate) SetPaths(v []File) {
	o.Paths = v
}

// GetRepository returns the Repository field value
func (o *RepoOverrideTemplate) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *RepoOverrideTemplate) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *RepoOverrideTemplate) SetRepository(v string) {
	o.Repository = v
}

// GetRevision returns the Revision field value
func (o *RepoOverrideTemplate) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *RepoOverrideTemplate) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *RepoOverrideTemplate) SetRevision(v string) {
	o.Revision = v
}

func (o RepoOverrideTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepoOverrideTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paths"] = o.Paths
	toSerialize["repository"] = o.Repository
	toSerialize["revision"] = o.Revision
	return toSerialize, nil
}

type NullableRepoOverrideTemplate struct {
	value *RepoOverrideTemplate
	isSet bool
}

func (v NullableRepoOverrideTemplate) Get() *RepoOverrideTemplate {
	return v.value
}

func (v *NullableRepoOverrideTemplate) Set(val *RepoOverrideTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableRepoOverrideTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableRepoOverrideTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepoOverrideTemplate(val *RepoOverrideTemplate) *NullableRepoOverrideTemplate {
	return &NullableRepoOverrideTemplate{value: val, isSet: true}
}

func (v NullableRepoOverrideTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepoOverrideTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


