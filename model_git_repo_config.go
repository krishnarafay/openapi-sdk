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

// checks if the GitRepoConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitRepoConfig{}

// GitRepoConfig git repo configuration for pipeline trigger
type GitRepoConfig struct {
	// paths in the git repo to watch for changes
	Paths []File `json:"paths"`
	// provider for the git repo
	Provider string `json:"provider"`
	// name of the git repository
	Repository string `json:"repository"`
	// branch or tag in the git repository to watch for changes
	Revision string `json:"revision"`
}

// NewGitRepoConfig instantiates a new GitRepoConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitRepoConfig(paths []File, provider string, repository string, revision string) *GitRepoConfig {
	this := GitRepoConfig{}
	this.Paths = paths
	this.Provider = provider
	this.Repository = repository
	this.Revision = revision
	return &this
}

// NewGitRepoConfigWithDefaults instantiates a new GitRepoConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitRepoConfigWithDefaults() *GitRepoConfig {
	this := GitRepoConfig{}
	return &this
}

// GetPaths returns the Paths field value
func (o *GitRepoConfig) GetPaths() []File {
	if o == nil {
		var ret []File
		return ret
	}

	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value
// and a boolean to check if the value has been set.
func (o *GitRepoConfig) GetPathsOk() ([]File, bool) {
	if o == nil {
		return nil, false
	}
	return o.Paths, true
}

// SetPaths sets field value
func (o *GitRepoConfig) SetPaths(v []File) {
	o.Paths = v
}

// GetProvider returns the Provider field value
func (o *GitRepoConfig) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *GitRepoConfig) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *GitRepoConfig) SetProvider(v string) {
	o.Provider = v
}

// GetRepository returns the Repository field value
func (o *GitRepoConfig) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *GitRepoConfig) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *GitRepoConfig) SetRepository(v string) {
	o.Repository = v
}

// GetRevision returns the Revision field value
func (o *GitRepoConfig) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *GitRepoConfig) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *GitRepoConfig) SetRevision(v string) {
	o.Revision = v
}

func (o GitRepoConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitRepoConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paths"] = o.Paths
	toSerialize["provider"] = o.Provider
	toSerialize["repository"] = o.Repository
	toSerialize["revision"] = o.Revision
	return toSerialize, nil
}

type NullableGitRepoConfig struct {
	value *GitRepoConfig
	isSet bool
}

func (v NullableGitRepoConfig) Get() *GitRepoConfig {
	return v.value
}

func (v *NullableGitRepoConfig) Set(val *GitRepoConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGitRepoConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGitRepoConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitRepoConfig(val *GitRepoConfig) *NullableGitRepoConfig {
	return &NullableGitRepoConfig{value: val, isSet: true}
}

func (v NullableGitRepoConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitRepoConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


