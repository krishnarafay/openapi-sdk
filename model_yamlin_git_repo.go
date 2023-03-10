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

// checks if the YAMLInGitRepo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &YAMLInGitRepo{}

// YAMLInGitRepo represents YAML files in a Git repo
type YAMLInGitRepo struct {
	// relative paths to file in the git repository
	Paths []File `json:"paths"`
	// name of the git repository
	Repository string `json:"repository"`
	// branch or tag in the git repository
	Revision string `json:"revision"`
}

// NewYAMLInGitRepo instantiates a new YAMLInGitRepo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYAMLInGitRepo(paths []File, repository string, revision string) *YAMLInGitRepo {
	this := YAMLInGitRepo{}
	this.Paths = paths
	this.Repository = repository
	this.Revision = revision
	return &this
}

// NewYAMLInGitRepoWithDefaults instantiates a new YAMLInGitRepo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYAMLInGitRepoWithDefaults() *YAMLInGitRepo {
	this := YAMLInGitRepo{}
	return &this
}

// GetPaths returns the Paths field value
func (o *YAMLInGitRepo) GetPaths() []File {
	if o == nil {
		var ret []File
		return ret
	}

	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value
// and a boolean to check if the value has been set.
func (o *YAMLInGitRepo) GetPathsOk() ([]File, bool) {
	if o == nil {
		return nil, false
	}
	return o.Paths, true
}

// SetPaths sets field value
func (o *YAMLInGitRepo) SetPaths(v []File) {
	o.Paths = v
}

// GetRepository returns the Repository field value
func (o *YAMLInGitRepo) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *YAMLInGitRepo) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *YAMLInGitRepo) SetRepository(v string) {
	o.Repository = v
}

// GetRevision returns the Revision field value
func (o *YAMLInGitRepo) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *YAMLInGitRepo) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *YAMLInGitRepo) SetRevision(v string) {
	o.Revision = v
}

func (o YAMLInGitRepo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o YAMLInGitRepo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paths"] = o.Paths
	toSerialize["repository"] = o.Repository
	toSerialize["revision"] = o.Revision
	return toSerialize, nil
}

type NullableYAMLInGitRepo struct {
	value *YAMLInGitRepo
	isSet bool
}

func (v NullableYAMLInGitRepo) Get() *YAMLInGitRepo {
	return v.value
}

func (v *NullableYAMLInGitRepo) Set(val *YAMLInGitRepo) {
	v.value = val
	v.isSet = true
}

func (v NullableYAMLInGitRepo) IsSet() bool {
	return v.isSet
}

func (v *NullableYAMLInGitRepo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYAMLInGitRepo(val *YAMLInGitRepo) *NullableYAMLInGitRepo {
	return &NullableYAMLInGitRepo{value: val, isSet: true}
}

func (v NullableYAMLInGitRepo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYAMLInGitRepo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


