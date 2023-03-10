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

// checks if the Httpproxyconfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Httpproxyconfig{}

// Httpproxyconfig struct for Httpproxyconfig
type Httpproxyconfig struct {
	// The HTTP proxy server endpoint to use.
	HttpProxy *string `json:"httpProxy,omitempty"`
	// The HTTPS proxy server endpoint to use.
	HttpsProxy *string `json:"httpsProxy,omitempty"`
	// The endpoints that should not go through proxy.
	NoProxy []string `json:"noProxy,omitempty"`
	// Alternative CA cert to use for connecting to proxy servers.
	TrustedCa *string `json:"trustedCa,omitempty"`
}

// NewHttpproxyconfig instantiates a new Httpproxyconfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpproxyconfig() *Httpproxyconfig {
	this := Httpproxyconfig{}
	return &this
}

// NewHttpproxyconfigWithDefaults instantiates a new Httpproxyconfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpproxyconfigWithDefaults() *Httpproxyconfig {
	this := Httpproxyconfig{}
	return &this
}

// GetHttpProxy returns the HttpProxy field value if set, zero value otherwise.
func (o *Httpproxyconfig) GetHttpProxy() string {
	if o == nil || isNil(o.HttpProxy) {
		var ret string
		return ret
	}
	return *o.HttpProxy
}

// GetHttpProxyOk returns a tuple with the HttpProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Httpproxyconfig) GetHttpProxyOk() (*string, bool) {
	if o == nil || isNil(o.HttpProxy) {
		return nil, false
	}
	return o.HttpProxy, true
}

// HasHttpProxy returns a boolean if a field has been set.
func (o *Httpproxyconfig) HasHttpProxy() bool {
	if o != nil && !isNil(o.HttpProxy) {
		return true
	}

	return false
}

// SetHttpProxy gets a reference to the given string and assigns it to the HttpProxy field.
func (o *Httpproxyconfig) SetHttpProxy(v string) {
	o.HttpProxy = &v
}

// GetHttpsProxy returns the HttpsProxy field value if set, zero value otherwise.
func (o *Httpproxyconfig) GetHttpsProxy() string {
	if o == nil || isNil(o.HttpsProxy) {
		var ret string
		return ret
	}
	return *o.HttpsProxy
}

// GetHttpsProxyOk returns a tuple with the HttpsProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Httpproxyconfig) GetHttpsProxyOk() (*string, bool) {
	if o == nil || isNil(o.HttpsProxy) {
		return nil, false
	}
	return o.HttpsProxy, true
}

// HasHttpsProxy returns a boolean if a field has been set.
func (o *Httpproxyconfig) HasHttpsProxy() bool {
	if o != nil && !isNil(o.HttpsProxy) {
		return true
	}

	return false
}

// SetHttpsProxy gets a reference to the given string and assigns it to the HttpsProxy field.
func (o *Httpproxyconfig) SetHttpsProxy(v string) {
	o.HttpsProxy = &v
}

// GetNoProxy returns the NoProxy field value if set, zero value otherwise.
func (o *Httpproxyconfig) GetNoProxy() []string {
	if o == nil || isNil(o.NoProxy) {
		var ret []string
		return ret
	}
	return o.NoProxy
}

// GetNoProxyOk returns a tuple with the NoProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Httpproxyconfig) GetNoProxyOk() ([]string, bool) {
	if o == nil || isNil(o.NoProxy) {
		return nil, false
	}
	return o.NoProxy, true
}

// HasNoProxy returns a boolean if a field has been set.
func (o *Httpproxyconfig) HasNoProxy() bool {
	if o != nil && !isNil(o.NoProxy) {
		return true
	}

	return false
}

// SetNoProxy gets a reference to the given []string and assigns it to the NoProxy field.
func (o *Httpproxyconfig) SetNoProxy(v []string) {
	o.NoProxy = v
}

// GetTrustedCa returns the TrustedCa field value if set, zero value otherwise.
func (o *Httpproxyconfig) GetTrustedCa() string {
	if o == nil || isNil(o.TrustedCa) {
		var ret string
		return ret
	}
	return *o.TrustedCa
}

// GetTrustedCaOk returns a tuple with the TrustedCa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Httpproxyconfig) GetTrustedCaOk() (*string, bool) {
	if o == nil || isNil(o.TrustedCa) {
		return nil, false
	}
	return o.TrustedCa, true
}

// HasTrustedCa returns a boolean if a field has been set.
func (o *Httpproxyconfig) HasTrustedCa() bool {
	if o != nil && !isNil(o.TrustedCa) {
		return true
	}

	return false
}

// SetTrustedCa gets a reference to the given string and assigns it to the TrustedCa field.
func (o *Httpproxyconfig) SetTrustedCa(v string) {
	o.TrustedCa = &v
}

func (o Httpproxyconfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Httpproxyconfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HttpProxy) {
		toSerialize["httpProxy"] = o.HttpProxy
	}
	if !isNil(o.HttpsProxy) {
		toSerialize["httpsProxy"] = o.HttpsProxy
	}
	if !isNil(o.NoProxy) {
		toSerialize["noProxy"] = o.NoProxy
	}
	if !isNil(o.TrustedCa) {
		toSerialize["trustedCa"] = o.TrustedCa
	}
	return toSerialize, nil
}

type NullableHttpproxyconfig struct {
	value *Httpproxyconfig
	isSet bool
}

func (v NullableHttpproxyconfig) Get() *Httpproxyconfig {
	return v.value
}

func (v *NullableHttpproxyconfig) Set(val *Httpproxyconfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpproxyconfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpproxyconfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpproxyconfig(val *Httpproxyconfig) *NullableHttpproxyconfig {
	return &NullableHttpproxyconfig{value: val, isSet: true}
}

func (v NullableHttpproxyconfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpproxyconfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


