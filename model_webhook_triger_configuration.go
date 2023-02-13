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

// checks if the WebhookTrigerConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookTrigerConfiguration{}

// WebhookTrigerConfiguration webhook trigger configuration
type WebhookTrigerConfiguration struct {
	Repo CronTriggerConfigRepo `json:"repo"`
}

// NewWebhookTrigerConfiguration instantiates a new WebhookTrigerConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookTrigerConfiguration(repo CronTriggerConfigRepo) *WebhookTrigerConfiguration {
	this := WebhookTrigerConfiguration{}
	this.Repo = repo
	return &this
}

// NewWebhookTrigerConfigurationWithDefaults instantiates a new WebhookTrigerConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookTrigerConfigurationWithDefaults() *WebhookTrigerConfiguration {
	this := WebhookTrigerConfiguration{}
	return &this
}

// GetRepo returns the Repo field value
func (o *WebhookTrigerConfiguration) GetRepo() CronTriggerConfigRepo {
	if o == nil {
		var ret CronTriggerConfigRepo
		return ret
	}

	return o.Repo
}

// GetRepoOk returns a tuple with the Repo field value
// and a boolean to check if the value has been set.
func (o *WebhookTrigerConfiguration) GetRepoOk() (*CronTriggerConfigRepo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repo, true
}

// SetRepo sets field value
func (o *WebhookTrigerConfiguration) SetRepo(v CronTriggerConfigRepo) {
	o.Repo = v
}

func (o WebhookTrigerConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookTrigerConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["repo"] = o.Repo
	return toSerialize, nil
}

type NullableWebhookTrigerConfiguration struct {
	value *WebhookTrigerConfiguration
	isSet bool
}

func (v NullableWebhookTrigerConfiguration) Get() *WebhookTrigerConfiguration {
	return v.value
}

func (v *NullableWebhookTrigerConfiguration) Set(val *WebhookTrigerConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookTrigerConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookTrigerConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookTrigerConfiguration(val *WebhookTrigerConfiguration) *NullableWebhookTrigerConfiguration {
	return &NullableWebhookTrigerConfiguration{value: val, isSet: true}
}

func (v NullableWebhookTrigerConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookTrigerConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

