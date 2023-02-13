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

// checks if the AzureCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureCredentials{}

// AzureCredentials Azure Credentials definition
type AzureCredentials struct {
	ClientId string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	SubscriptionId string `json:"subscriptionId"`
	TenantId string `json:"tenantId"`
}

// NewAzureCredentials instantiates a new AzureCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureCredentials(clientId string, clientSecret string, subscriptionId string, tenantId string) *AzureCredentials {
	this := AzureCredentials{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.SubscriptionId = subscriptionId
	this.TenantId = tenantId
	return &this
}

// NewAzureCredentialsWithDefaults instantiates a new AzureCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureCredentialsWithDefaults() *AzureCredentials {
	this := AzureCredentials{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *AzureCredentials) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *AzureCredentials) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *AzureCredentials) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *AzureCredentials) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *AzureCredentials) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *AzureCredentials) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetTenantId returns the TenantId field value
func (o *AzureCredentials) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *AzureCredentials) SetTenantId(v string) {
	o.TenantId = v
}

func (o AzureCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientId"] = o.ClientId
	toSerialize["clientSecret"] = o.ClientSecret
	toSerialize["subscriptionId"] = o.SubscriptionId
	toSerialize["tenantId"] = o.TenantId
	return toSerialize, nil
}

type NullableAzureCredentials struct {
	value *AzureCredentials
	isSet bool
}

func (v NullableAzureCredentials) Get() *AzureCredentials {
	return v.value
}

func (v *NullableAzureCredentials) Set(val *AzureCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureCredentials(val *AzureCredentials) *NullableAzureCredentials {
	return &NullableAzureCredentials{value: val, isSet: true}
}

func (v NullableAzureCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


