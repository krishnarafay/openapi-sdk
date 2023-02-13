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

// checks if the TriggerHelmRepoConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerHelmRepoConfiguration{}

// TriggerHelmRepoConfiguration helm repo configuration for pipeline trigger
type TriggerHelmRepoConfiguration struct {
	// name of the chart in repo
	ChartName string `json:"chartName"`
	// version of the chart in repo
	ChartVersion string `json:"chartVersion"`
	// name of the helm repository
	Repository string `json:"repository"`
}

// NewTriggerHelmRepoConfiguration instantiates a new TriggerHelmRepoConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerHelmRepoConfiguration(chartName string, chartVersion string, repository string) *TriggerHelmRepoConfiguration {
	this := TriggerHelmRepoConfiguration{}
	this.ChartName = chartName
	this.ChartVersion = chartVersion
	this.Repository = repository
	return &this
}

// NewTriggerHelmRepoConfigurationWithDefaults instantiates a new TriggerHelmRepoConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerHelmRepoConfigurationWithDefaults() *TriggerHelmRepoConfiguration {
	this := TriggerHelmRepoConfiguration{}
	return &this
}

// GetChartName returns the ChartName field value
func (o *TriggerHelmRepoConfiguration) GetChartName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartName
}

// GetChartNameOk returns a tuple with the ChartName field value
// and a boolean to check if the value has been set.
func (o *TriggerHelmRepoConfiguration) GetChartNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartName, true
}

// SetChartName sets field value
func (o *TriggerHelmRepoConfiguration) SetChartName(v string) {
	o.ChartName = v
}

// GetChartVersion returns the ChartVersion field value
func (o *TriggerHelmRepoConfiguration) GetChartVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChartVersion
}

// GetChartVersionOk returns a tuple with the ChartVersion field value
// and a boolean to check if the value has been set.
func (o *TriggerHelmRepoConfiguration) GetChartVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartVersion, true
}

// SetChartVersion sets field value
func (o *TriggerHelmRepoConfiguration) SetChartVersion(v string) {
	o.ChartVersion = v
}

// GetRepository returns the Repository field value
func (o *TriggerHelmRepoConfiguration) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *TriggerHelmRepoConfiguration) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *TriggerHelmRepoConfiguration) SetRepository(v string) {
	o.Repository = v
}

func (o TriggerHelmRepoConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerHelmRepoConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chartName"] = o.ChartName
	toSerialize["chartVersion"] = o.ChartVersion
	toSerialize["repository"] = o.Repository
	return toSerialize, nil
}

type NullableTriggerHelmRepoConfiguration struct {
	value *TriggerHelmRepoConfiguration
	isSet bool
}

func (v NullableTriggerHelmRepoConfiguration) Get() *TriggerHelmRepoConfiguration {
	return v.value
}

func (v *NullableTriggerHelmRepoConfiguration) Set(val *TriggerHelmRepoConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerHelmRepoConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerHelmRepoConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerHelmRepoConfiguration(val *TriggerHelmRepoConfiguration) *NullableTriggerHelmRepoConfiguration {
	return &NullableTriggerHelmRepoConfiguration{value: val, isSet: true}
}

func (v NullableTriggerHelmRepoConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerHelmRepoConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


