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

// checks if the Gmsaprofile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Gmsaprofile{}

// Gmsaprofile struct for Gmsaprofile
type Gmsaprofile struct {
	// Specifies the DNS server for Windows gMSA. Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.
	DnsServer *string `json:"dnsServer,omitempty"`
	// Specifies whether to enable Windows gMSA in the managed cluster.
	Enabled *bool `json:"enabled,omitempty"`
	// Specifies the root domain name for Windows gMSA. Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster. 
	RootDomainName *string `json:"rootDomainName,omitempty"`
}

// NewGmsaprofile instantiates a new Gmsaprofile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGmsaprofile() *Gmsaprofile {
	this := Gmsaprofile{}
	return &this
}

// NewGmsaprofileWithDefaults instantiates a new Gmsaprofile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGmsaprofileWithDefaults() *Gmsaprofile {
	this := Gmsaprofile{}
	return &this
}

// GetDnsServer returns the DnsServer field value if set, zero value otherwise.
func (o *Gmsaprofile) GetDnsServer() string {
	if o == nil || isNil(o.DnsServer) {
		var ret string
		return ret
	}
	return *o.DnsServer
}

// GetDnsServerOk returns a tuple with the DnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gmsaprofile) GetDnsServerOk() (*string, bool) {
	if o == nil || isNil(o.DnsServer) {
		return nil, false
	}
	return o.DnsServer, true
}

// HasDnsServer returns a boolean if a field has been set.
func (o *Gmsaprofile) HasDnsServer() bool {
	if o != nil && !isNil(o.DnsServer) {
		return true
	}

	return false
}

// SetDnsServer gets a reference to the given string and assigns it to the DnsServer field.
func (o *Gmsaprofile) SetDnsServer(v string) {
	o.DnsServer = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Gmsaprofile) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gmsaprofile) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Gmsaprofile) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Gmsaprofile) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRootDomainName returns the RootDomainName field value if set, zero value otherwise.
func (o *Gmsaprofile) GetRootDomainName() string {
	if o == nil || isNil(o.RootDomainName) {
		var ret string
		return ret
	}
	return *o.RootDomainName
}

// GetRootDomainNameOk returns a tuple with the RootDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gmsaprofile) GetRootDomainNameOk() (*string, bool) {
	if o == nil || isNil(o.RootDomainName) {
		return nil, false
	}
	return o.RootDomainName, true
}

// HasRootDomainName returns a boolean if a field has been set.
func (o *Gmsaprofile) HasRootDomainName() bool {
	if o != nil && !isNil(o.RootDomainName) {
		return true
	}

	return false
}

// SetRootDomainName gets a reference to the given string and assigns it to the RootDomainName field.
func (o *Gmsaprofile) SetRootDomainName(v string) {
	o.RootDomainName = &v
}

func (o Gmsaprofile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Gmsaprofile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DnsServer) {
		toSerialize["dnsServer"] = o.DnsServer
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.RootDomainName) {
		toSerialize["rootDomainName"] = o.RootDomainName
	}
	return toSerialize, nil
}

type NullableGmsaprofile struct {
	value *Gmsaprofile
	isSet bool
}

func (v NullableGmsaprofile) Get() *Gmsaprofile {
	return v.value
}

func (v *NullableGmsaprofile) Set(val *Gmsaprofile) {
	v.value = val
	v.isSet = true
}

func (v NullableGmsaprofile) IsSet() bool {
	return v.isSet
}

func (v *NullableGmsaprofile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGmsaprofile(val *Gmsaprofile) *NullableGmsaprofile {
	return &NullableGmsaprofile{value: val, isSet: true}
}

func (v NullableGmsaprofile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGmsaprofile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


