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

// checks if the HelmInHelmRepo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmInHelmRepo{}

// HelmInHelmRepo represents Helm chart in a Helm repo
type HelmInHelmRepo struct {
	// name of the chart
	ChartName string `json:"chartName"`
	// version of the chart
	ChartVersion string `json:"chartVersion"`
	// name of the helm repository
	Repository string `json:"repository"`
	// relative paths to values files
	ValuesPaths []File `json:"valuesPaths,omitempty"`
	ValuesRef *OverrideRepoReference `json:"valuesRef,omitempty"`
}

// NewHelmInHelmRepo instantiates a new HelmInHelmRepo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmInHelmRepo(chartName string, chartVersion string, repository string) *HelmInHelmRepo {
	this := HelmInHelmRepo{}
	this.ChartName = chartName
	this.ChartVersion = chartVersion
	this.Repository = repository
	return &this
}

// NewHelmInHelmRepoWithDefaults instantiates a new HelmInHelmRepo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmInHelmRepoWithDefaults() *HelmInHelmRepo {
	this := HelmInHelmRepo{}
	return &this
}

// GetChartName returns the ChartName field value
func (o *HelmInHelmRepo) GetChartName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartName
}

// GetChartNameOk returns a tuple with the ChartName field value
// and a boolean to check if the value has been set.
func (o *HelmInHelmRepo) GetChartNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartName, true
}

// SetChartName sets field value
func (o *HelmInHelmRepo) SetChartName(v string) {
	o.ChartName = v
}

// GetChartVersion returns the ChartVersion field value
func (o *HelmInHelmRepo) GetChartVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartVersion
}

// GetChartVersionOk returns a tuple with the ChartVersion field value
// and a boolean to check if the value has been set.
func (o *HelmInHelmRepo) GetChartVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartVersion, true
}

// SetChartVersion sets field value
func (o *HelmInHelmRepo) SetChartVersion(v string) {
	o.ChartVersion = v
}

// GetRepository returns the Repository field value
func (o *HelmInHelmRepo) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *HelmInHelmRepo) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *HelmInHelmRepo) SetRepository(v string) {
	o.Repository = v
}

// GetValuesPaths returns the ValuesPaths field value if set, zero value otherwise.
func (o *HelmInHelmRepo) GetValuesPaths() []File {
	if o == nil || isNil(o.ValuesPaths) {
		var ret []File
		return ret
	}
	return o.ValuesPaths
}

// GetValuesPathsOk returns a tuple with the ValuesPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmInHelmRepo) GetValuesPathsOk() ([]File, bool) {
	if o == nil || isNil(o.ValuesPaths) {
		return nil, false
	}
	return o.ValuesPaths, true
}

// HasValuesPaths returns a boolean if a field has been set.
func (o *HelmInHelmRepo) HasValuesPaths() bool {
	if o != nil && !isNil(o.ValuesPaths) {
		return true
	}

	return false
}

// SetValuesPaths gets a reference to the given []File and assigns it to the ValuesPaths field.
func (o *HelmInHelmRepo) SetValuesPaths(v []File) {
	o.ValuesPaths = v
}

// GetValuesRef returns the ValuesRef field value if set, zero value otherwise.
func (o *HelmInHelmRepo) GetValuesRef() OverrideRepoReference {
	if o == nil || isNil(o.ValuesRef) {
		var ret OverrideRepoReference
		return ret
	}
	return *o.ValuesRef
}

// GetValuesRefOk returns a tuple with the ValuesRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmInHelmRepo) GetValuesRefOk() (*OverrideRepoReference, bool) {
	if o == nil || isNil(o.ValuesRef) {
		return nil, false
	}
	return o.ValuesRef, true
}

// HasValuesRef returns a boolean if a field has been set.
func (o *HelmInHelmRepo) HasValuesRef() bool {
	if o != nil && !isNil(o.ValuesRef) {
		return true
	}

	return false
}

// SetValuesRef gets a reference to the given OverrideRepoReference and assigns it to the ValuesRef field.
func (o *HelmInHelmRepo) SetValuesRef(v OverrideRepoReference) {
	o.ValuesRef = &v
}

func (o HelmInHelmRepo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmInHelmRepo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chartName"] = o.ChartName
	toSerialize["chartVersion"] = o.ChartVersion
	toSerialize["repository"] = o.Repository
	if !isNil(o.ValuesPaths) {
		toSerialize["valuesPaths"] = o.ValuesPaths
	}
	if !isNil(o.ValuesRef) {
		toSerialize["valuesRef"] = o.ValuesRef
	}
	return toSerialize, nil
}

type NullableHelmInHelmRepo struct {
	value *HelmInHelmRepo
	isSet bool
}

func (v NullableHelmInHelmRepo) Get() *HelmInHelmRepo {
	return v.value
}

func (v *NullableHelmInHelmRepo) Set(val *HelmInHelmRepo) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmInHelmRepo) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmInHelmRepo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmInHelmRepo(val *HelmInHelmRepo) *NullableHelmInHelmRepo {
	return &NullableHelmInHelmRepo{value: val, isSet: true}
}

func (v NullableHelmInHelmRepo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmInHelmRepo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


