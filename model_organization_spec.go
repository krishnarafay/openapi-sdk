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

// checks if the OrganizationSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationSpec{}

// OrganizationSpec struct for OrganizationSpec
type OrganizationSpec struct {
	Active *bool `json:"active,omitempty"`
	AddressLine1 *string `json:"addressLine1,omitempty"`
	AddressLine2 *string `json:"addressLine2,omitempty"`
	Approved *bool `json:"approved,omitempty"`
	AreClustersShared *bool `json:"areClustersShared,omitempty"`
	BillingAddress *string `json:"billingAddress,omitempty"`
	City *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
	IsPrivate *bool `json:"isPrivate,omitempty"`
	IsTotpEnabled *bool `json:"isTotpEnabled,omitempty"`
	Phone *string `json:"phone,omitempty"`
	State *string `json:"state,omitempty"`
	Type *string `json:"type,omitempty"`
	Zipcode *string `json:"zipcode,omitempty"`
}

// NewOrganizationSpec instantiates a new OrganizationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationSpec() *OrganizationSpec {
	this := OrganizationSpec{}
	return &this
}

// NewOrganizationSpecWithDefaults instantiates a new OrganizationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationSpecWithDefaults() *OrganizationSpec {
	this := OrganizationSpec{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *OrganizationSpec) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSpec) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *OrganizationSpec) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *OrganizationSpec) SetActive(v bool) {
	o.Active = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *OrganizationSpec) GetAddressLine1() string {
	if o == nil || isNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSpec) GetAddressLine1Ok() (*string, bool) {
	if o == nil || isNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *OrganizationSpec) HasAddressLine1() bool {
	if o != nil && !isNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *OrganizationSpec) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *OrganizationSpec) GetAddressLine2() string {
	if o == nil || isNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSpec) GetAddressLine2Ok() (*string, bool) {
	if o == nil || isNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *OrganizationSpec) HasAddressLine2() bool {
	if o != nil && !isNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *OrganizationSpec) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetApproved returns the Approved field value if set, zero value otherwise.
func (o *OrganizationSpec) GetApproved() bool {
	if o == nil || isNil(o.Approved) {
		var ret bool
		return ret
	}
	return *o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSpec) GetApprovedOk() (*bool, bool) {
	if o == nil || isNil(o.Approved) {
		return nil, false
	}
	return o.Approved, true
}

// HasApproved returns a boolean if a field has been set.
func (o *OrganizationSpec) HasApproved() bool {
	if o != nil && !isNil(o.Approved) {
		return true
	}

	return false
}

// SetApproved gets a reference to the given bool and assigns it to the Approved field.
func (o *OrganizationSpec) SetApproved(v bool) {
	o.Approved = &v
}

// GetAreClustersShared returns the AreClustersShared field value if set, zero value otherwise.
func (o *OrganizationSpec) GetAreClustersShared() bool {
	if o == nil || isNil(o.AreClustersShared) {
		var ret bool
		return ret
	}
	return *o.AreClustersShared
}

// GetAreClustersSharedOk returns a tuple with the AreClustersShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSpec) GetAreClustersSharedOk() (*bool, bool) {
	if o == nil || isNil(o.AreClustersShared) {
		return nil, false
	}
	return o.AreClustersShared, true
}

// HasAreClustersShared returns a boolean if a field has been set.
func (o *OrganizationSpec) HasAreClustersShared() bool {
	if o != nil && !isNil(o.AreClustersShared) {
		return true
	}

	return false
}

// SetAreClustersShared gets a reference to the given bool and assigns it to the AreClustersShared field.
func (o *OrganizationSpec) SetAreClustersShared(v bool) {
	o.AreClustersShared = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *OrganizationSpec) GetBillingAddress() string {
	if o == nil || isNil(o.BillingAddress) {
		var ret string
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSpec) GetBillingAddressOk() (*string, bool) {
	if o == nil || isNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *OrganizationSpec) HasBillingAddress() bool {
	if o != nil && !isNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given string and assigns it to the BillingAddress field.
func (o *OrganizationSpec) SetBillingAddress(v string) {
	o.BillingAddress = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *OrganizationSpec) GetCity() string {
	if o == nil || isNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSpec) GetCityOk() (*string, bool) {
	if o == nil || isNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *OrganizationSpec) HasCity() bool {
	if o != nil && !isNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *OrganizationSpec) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *OrganizationSpec) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSpec) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *OrganizationSpec) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *OrganizationSpec) SetCountry(v string) {
	o.Country = &v
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *OrganizationSpec) GetIsPrivate() bool {
	if o == nil || isNil(o.IsPrivate) {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSpec) GetIsPrivateOk() (*bool, bool) {
	if o == nil || isNil(o.IsPrivate) {
		return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *OrganizationSpec) HasIsPrivate() bool {
	if o != nil && !isNil(o.IsPrivate) {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *OrganizationSpec) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}

// GetIsTotpEnabled returns the IsTotpEnabled field value if set, zero value otherwise.
func (o *OrganizationSpec) GetIsTotpEnabled() bool {
	if o == nil || isNil(o.IsTotpEnabled) {
		var ret bool
		return ret
	}
	return *o.IsTotpEnabled
}

// GetIsTotpEnabledOk returns a tuple with the IsTotpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSpec) GetIsTotpEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsTotpEnabled) {
		return nil, false
	}
	return o.IsTotpEnabled, true
}

// HasIsTotpEnabled returns a boolean if a field has been set.
func (o *OrganizationSpec) HasIsTotpEnabled() bool {
	if o != nil && !isNil(o.IsTotpEnabled) {
		return true
	}

	return false
}

// SetIsTotpEnabled gets a reference to the given bool and assigns it to the IsTotpEnabled field.
func (o *OrganizationSpec) SetIsTotpEnabled(v bool) {
	o.IsTotpEnabled = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *OrganizationSpec) GetPhone() string {
	if o == nil || isNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSpec) GetPhoneOk() (*string, bool) {
	if o == nil || isNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *OrganizationSpec) HasPhone() bool {
	if o != nil && !isNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *OrganizationSpec) SetPhone(v string) {
	o.Phone = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OrganizationSpec) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSpec) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OrganizationSpec) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *OrganizationSpec) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrganizationSpec) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSpec) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrganizationSpec) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrganizationSpec) SetType(v string) {
	o.Type = &v
}

// GetZipcode returns the Zipcode field value if set, zero value otherwise.
func (o *OrganizationSpec) GetZipcode() string {
	if o == nil || isNil(o.Zipcode) {
		var ret string
		return ret
	}
	return *o.Zipcode
}

// GetZipcodeOk returns a tuple with the Zipcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSpec) GetZipcodeOk() (*string, bool) {
	if o == nil || isNil(o.Zipcode) {
		return nil, false
	}
	return o.Zipcode, true
}

// HasZipcode returns a boolean if a field has been set.
func (o *OrganizationSpec) HasZipcode() bool {
	if o != nil && !isNil(o.Zipcode) {
		return true
	}

	return false
}

// SetZipcode gets a reference to the given string and assigns it to the Zipcode field.
func (o *OrganizationSpec) SetZipcode(v string) {
	o.Zipcode = &v
}

func (o OrganizationSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.AddressLine1) {
		toSerialize["addressLine1"] = o.AddressLine1
	}
	if !isNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	if !isNil(o.Approved) {
		toSerialize["approved"] = o.Approved
	}
	if !isNil(o.AreClustersShared) {
		toSerialize["areClustersShared"] = o.AreClustersShared
	}
	if !isNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !isNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !isNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !isNil(o.IsPrivate) {
		toSerialize["isPrivate"] = o.IsPrivate
	}
	if !isNil(o.IsTotpEnabled) {
		toSerialize["isTotpEnabled"] = o.IsTotpEnabled
	}
	if !isNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Zipcode) {
		toSerialize["zipcode"] = o.Zipcode
	}
	return toSerialize, nil
}

type NullableOrganizationSpec struct {
	value *OrganizationSpec
	isSet bool
}

func (v NullableOrganizationSpec) Get() *OrganizationSpec {
	return v.value
}

func (v *NullableOrganizationSpec) Set(val *OrganizationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationSpec(val *OrganizationSpec) *NullableOrganizationSpec {
	return &NullableOrganizationSpec{value: val, isSet: true}
}

func (v NullableOrganizationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


