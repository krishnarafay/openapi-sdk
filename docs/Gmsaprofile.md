# Gmsaprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsServer** | Pointer to **string** | Specifies the DNS server for Windows gMSA. Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster. | [optional] 
**Enabled** | Pointer to **bool** | Specifies whether to enable Windows gMSA in the managed cluster. | [optional] 
**RootDomainName** | Pointer to **string** | Specifies the root domain name for Windows gMSA. Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.  | [optional] 

## Methods

### NewGmsaprofile

`func NewGmsaprofile() *Gmsaprofile`

NewGmsaprofile instantiates a new Gmsaprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGmsaprofileWithDefaults

`func NewGmsaprofileWithDefaults() *Gmsaprofile`

NewGmsaprofileWithDefaults instantiates a new Gmsaprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsServer

`func (o *Gmsaprofile) GetDnsServer() string`

GetDnsServer returns the DnsServer field if non-nil, zero value otherwise.

### GetDnsServerOk

`func (o *Gmsaprofile) GetDnsServerOk() (*string, bool)`

GetDnsServerOk returns a tuple with the DnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServer

`func (o *Gmsaprofile) SetDnsServer(v string)`

SetDnsServer sets DnsServer field to given value.

### HasDnsServer

`func (o *Gmsaprofile) HasDnsServer() bool`

HasDnsServer returns a boolean if a field has been set.

### GetEnabled

`func (o *Gmsaprofile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Gmsaprofile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Gmsaprofile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Gmsaprofile) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRootDomainName

`func (o *Gmsaprofile) GetRootDomainName() string`

GetRootDomainName returns the RootDomainName field if non-nil, zero value otherwise.

### GetRootDomainNameOk

`func (o *Gmsaprofile) GetRootDomainNameOk() (*string, bool)`

GetRootDomainNameOk returns a tuple with the RootDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDomainName

`func (o *Gmsaprofile) SetRootDomainName(v string)`

SetRootDomainName sets RootDomainName field to given value.

### HasRootDomainName

`func (o *Gmsaprofile) HasRootDomainName() bool`

HasRootDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


