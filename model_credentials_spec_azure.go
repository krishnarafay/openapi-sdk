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

// checks if the CredentialsSpecAzure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsSpecAzure{}

// CredentialsSpecAzure Azure Credentials definition
type CredentialsSpecAzure struct {
	ClientId string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	SubscriptionId string `json:"subscriptionId"`
	TenantId string `json:"tenantId"`
}

// NewCredentialsSpecAzure instantiates a new CredentialsSpecAzure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsSpecAzure(clientId string, clientSecret string, subscriptionId string, tenantId string) *CredentialsSpecAzure {
	this := CredentialsSpecAzure{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.SubscriptionId = subscriptionId
	this.TenantId = tenantId
	return &this
}

// NewCredentialsSpecAzureWithDefaults instantiates a new CredentialsSpecAzure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsSpecAzureWithDefaults() *CredentialsSpecAzure {
	this := CredentialsSpecAzure{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *CredentialsSpecAzure) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *CredentialsSpecAzure) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *CredentialsSpecAzure) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *CredentialsSpecAzure) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *CredentialsSpecAzure) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *CredentialsSpecAzure) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *CredentialsSpecAzure) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *CredentialsSpecAzure) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *CredentialsSpecAzure) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetTenantId returns the TenantId field value
func (o *CredentialsSpecAzure) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *CredentialsSpecAzure) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *CredentialsSpecAzure) SetTenantId(v string) {
	o.TenantId = v
}

func (o CredentialsSpecAzure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialsSpecAzure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientId"] = o.ClientId
	toSerialize["clientSecret"] = o.ClientSecret
	toSerialize["subscriptionId"] = o.SubscriptionId
	toSerialize["tenantId"] = o.TenantId
	return toSerialize, nil
}

type NullableCredentialsSpecAzure struct {
	value *CredentialsSpecAzure
	isSet bool
}

func (v NullableCredentialsSpecAzure) Get() *CredentialsSpecAzure {
	return v.value
}

func (v *NullableCredentialsSpecAzure) Set(val *CredentialsSpecAzure) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsSpecAzure) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsSpecAzure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsSpecAzure(val *CredentialsSpecAzure) *NullableCredentialsSpecAzure {
	return &NullableCredentialsSpecAzure{value: val, isSet: true}
}

func (v NullableCredentialsSpecAzure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsSpecAzure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

