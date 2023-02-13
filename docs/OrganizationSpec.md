# OrganizationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**AreClustersShared** | Pointer to **bool** |  | [optional] 
**BillingAddress** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 
**IsTotpEnabled** | Pointer to **bool** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Zipcode** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganizationSpec

`func NewOrganizationSpec() *OrganizationSpec`

NewOrganizationSpec instantiates a new OrganizationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSpecWithDefaults

`func NewOrganizationSpecWithDefaults() *OrganizationSpec`

NewOrganizationSpecWithDefaults instantiates a new OrganizationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *OrganizationSpec) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OrganizationSpec) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OrganizationSpec) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OrganizationSpec) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddressLine1

`func (o *OrganizationSpec) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *OrganizationSpec) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *OrganizationSpec) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *OrganizationSpec) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *OrganizationSpec) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OrganizationSpec) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OrganizationSpec) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *OrganizationSpec) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetApproved

`func (o *OrganizationSpec) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *OrganizationSpec) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *OrganizationSpec) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *OrganizationSpec) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetAreClustersShared

`func (o *OrganizationSpec) GetAreClustersShared() bool`

GetAreClustersShared returns the AreClustersShared field if non-nil, zero value otherwise.

### GetAreClustersSharedOk

`func (o *OrganizationSpec) GetAreClustersSharedOk() (*bool, bool)`

GetAreClustersSharedOk returns a tuple with the AreClustersShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreClustersShared

`func (o *OrganizationSpec) SetAreClustersShared(v bool)`

SetAreClustersShared sets AreClustersShared field to given value.

### HasAreClustersShared

`func (o *OrganizationSpec) HasAreClustersShared() bool`

HasAreClustersShared returns a boolean if a field has been set.

### GetBillingAddress

`func (o *OrganizationSpec) GetBillingAddress() string`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *OrganizationSpec) GetBillingAddressOk() (*string, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *OrganizationSpec) SetBillingAddress(v string)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *OrganizationSpec) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCity

`func (o *OrganizationSpec) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrganizationSpec) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrganizationSpec) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrganizationSpec) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *OrganizationSpec) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrganizationSpec) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrganizationSpec) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OrganizationSpec) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetIsPrivate

`func (o *OrganizationSpec) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *OrganizationSpec) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *OrganizationSpec) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *OrganizationSpec) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetIsTotpEnabled

`func (o *OrganizationSpec) GetIsTotpEnabled() bool`

GetIsTotpEnabled returns the IsTotpEnabled field if non-nil, zero value otherwise.

### GetIsTotpEnabledOk

`func (o *OrganizationSpec) GetIsTotpEnabledOk() (*bool, bool)`

GetIsTotpEnabledOk returns a tuple with the IsTotpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTotpEnabled

`func (o *OrganizationSpec) SetIsTotpEnabled(v bool)`

SetIsTotpEnabled sets IsTotpEnabled field to given value.

### HasIsTotpEnabled

`func (o *OrganizationSpec) HasIsTotpEnabled() bool`

HasIsTotpEnabled returns a boolean if a field has been set.

### GetPhone

`func (o *OrganizationSpec) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *OrganizationSpec) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *OrganizationSpec) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *OrganizationSpec) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetState

`func (o *OrganizationSpec) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrganizationSpec) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrganizationSpec) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrganizationSpec) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *OrganizationSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrganizationSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrganizationSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrganizationSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetZipcode

`func (o *OrganizationSpec) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *OrganizationSpec) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *OrganizationSpec) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *OrganizationSpec) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


