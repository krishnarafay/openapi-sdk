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

// checks if the HelmInGitRepoArtifact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmInGitRepoArtifact{}

// HelmInGitRepoArtifact represents Helm files in a Git repo
type HelmInGitRepoArtifact struct {
	ChartPath File `json:"chartPath"`
	// name of the git repository
	Repository string `json:"repository"`
	// branch or tag in the git repository
	Revision string `json:"revision"`
	// relative paths to value file in the git repository
	ValuesPaths []File `json:"valuesPaths,omitempty"`
	ValuesRef *OverrideRepoReference `json:"valuesRef,omitempty"`
}

// NewHelmInGitRepoArtifact instantiates a new HelmInGitRepoArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmInGitRepoArtifact(chartPath File, repository string, revision string) *HelmInGitRepoArtifact {
	this := HelmInGitRepoArtifact{}
	this.ChartPath = chartPath
	this.Repository = repository
	this.Revision = revision
	return &this
}

// NewHelmInGitRepoArtifactWithDefaults instantiates a new HelmInGitRepoArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmInGitRepoArtifactWithDefaults() *HelmInGitRepoArtifact {
	this := HelmInGitRepoArtifact{}
	return &this
}

// GetChartPath returns the ChartPath field value
func (o *HelmInGitRepoArtifact) GetChartPath() File {
	if o == nil {
		var ret File
		return ret
	}

	return o.ChartPath
}

// GetChartPathOk returns a tuple with the ChartPath field value
// and a boolean to check if the value has been set.
func (o *HelmInGitRepoArtifact) GetChartPathOk() (*File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartPath, true
}

// SetChartPath sets field value
func (o *HelmInGitRepoArtifact) SetChartPath(v File) {
	o.ChartPath = v
}

// GetRepository returns the Repository field value
func (o *HelmInGitRepoArtifact) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *HelmInGitRepoArtifact) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *HelmInGitRepoArtifact) SetRepository(v string) {
	o.Repository = v
}

// GetRevision returns the Revision field value
func (o *HelmInGitRepoArtifact) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *HelmInGitRepoArtifact) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *HelmInGitRepoArtifact) SetRevision(v string) {
	o.Revision = v
}

// GetValuesPaths returns the ValuesPaths field value if set, zero value otherwise.
func (o *HelmInGitRepoArtifact) GetValuesPaths() []File {
	if o == nil || isNil(o.ValuesPaths) {
		var ret []File
		return ret
	}
	return o.ValuesPaths
}

// GetValuesPathsOk returns a tuple with the ValuesPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmInGitRepoArtifact) GetValuesPathsOk() ([]File, bool) {
	if o == nil || isNil(o.ValuesPaths) {
		return nil, false
	}
	return o.ValuesPaths, true
}

// HasValuesPaths returns a boolean if a field has been set.
func (o *HelmInGitRepoArtifact) HasValuesPaths() bool {
	if o != nil && !isNil(o.ValuesPaths) {
		return true
	}

	return false
}

// SetValuesPaths gets a reference to the given []File and assigns it to the ValuesPaths field.
func (o *HelmInGitRepoArtifact) SetValuesPaths(v []File) {
	o.ValuesPaths = v
}

// GetValuesRef returns the ValuesRef field value if set, zero value otherwise.
func (o *HelmInGitRepoArtifact) GetValuesRef() OverrideRepoReference {
	if o == nil || isNil(o.ValuesRef) {
		var ret OverrideRepoReference
		return ret
	}
	return *o.ValuesRef
}

// GetValuesRefOk returns a tuple with the ValuesRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmInGitRepoArtifact) GetValuesRefOk() (*OverrideRepoReference, bool) {
	if o == nil || isNil(o.ValuesRef) {
		return nil, false
	}
	return o.ValuesRef, true
}

// HasValuesRef returns a boolean if a field has been set.
func (o *HelmInGitRepoArtifact) HasValuesRef() bool {
	if o != nil && !isNil(o.ValuesRef) {
		return true
	}

	return false
}

// SetValuesRef gets a reference to the given OverrideRepoReference and assigns it to the ValuesRef field.
func (o *HelmInGitRepoArtifact) SetValuesRef(v OverrideRepoReference) {
	o.ValuesRef = &v
}

func (o HelmInGitRepoArtifact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmInGitRepoArtifact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chartPath"] = o.ChartPath
	toSerialize["repository"] = o.Repository
	toSerialize["revision"] = o.Revision
	if !isNil(o.ValuesPaths) {
		toSerialize["valuesPaths"] = o.ValuesPaths
	}
	if !isNil(o.ValuesRef) {
		toSerialize["valuesRef"] = o.ValuesRef
	}
	return toSerialize, nil
}

type NullableHelmInGitRepoArtifact struct {
	value *HelmInGitRepoArtifact
	isSet bool
}

func (v NullableHelmInGitRepoArtifact) Get() *HelmInGitRepoArtifact {
	return v.value
}

func (v *NullableHelmInGitRepoArtifact) Set(val *HelmInGitRepoArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmInGitRepoArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmInGitRepoArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmInGitRepoArtifact(val *HelmInGitRepoArtifact) *NullableHelmInGitRepoArtifact {
	return &NullableHelmInGitRepoArtifact{value: val, isSet: true}
}

func (v NullableHelmInGitRepoArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmInGitRepoArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

