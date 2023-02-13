# PartnerSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**FavIconLink** | Pointer to **string** |  | [optional] 
**HelpdeskEmail** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**IsTOTPEnabled** | Pointer to **bool** |  | [optional] 
**LogoLink** | Pointer to **string** |  | [optional] 
**NotificationEmail** | Pointer to **string** |  | [optional] 
**OpsHost** | Pointer to **string** |  | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**ExtraData**](ExtraData.md) |  | [optional] 
**SupportTeamName** | Pointer to **string** |  | [optional] 
**TosLink** | Pointer to **string** |  | [optional] 

## Methods

### NewPartnerSpec

`func NewPartnerSpec() *PartnerSpec`

NewPartnerSpec instantiates a new PartnerSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerSpecWithDefaults

`func NewPartnerSpecWithDefaults() *PartnerSpec`

NewPartnerSpecWithDefaults instantiates a new PartnerSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *PartnerSpec) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PartnerSpec) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PartnerSpec) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PartnerSpec) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetFavIconLink

`func (o *PartnerSpec) GetFavIconLink() string`

GetFavIconLink returns the FavIconLink field if non-nil, zero value otherwise.

### GetFavIconLinkOk

`func (o *PartnerSpec) GetFavIconLinkOk() (*string, bool)`

GetFavIconLinkOk returns a tuple with the FavIconLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavIconLink

`func (o *PartnerSpec) SetFavIconLink(v string)`

SetFavIconLink sets FavIconLink field to given value.

### HasFavIconLink

`func (o *PartnerSpec) HasFavIconLink() bool`

HasFavIconLink returns a boolean if a field has been set.

### GetHelpdeskEmail

`func (o *PartnerSpec) GetHelpdeskEmail() string`

GetHelpdeskEmail returns the HelpdeskEmail field if non-nil, zero value otherwise.

### GetHelpdeskEmailOk

`func (o *PartnerSpec) GetHelpdeskEmailOk() (*string, bool)`

GetHelpdeskEmailOk returns a tuple with the HelpdeskEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpdeskEmail

`func (o *PartnerSpec) SetHelpdeskEmail(v string)`

SetHelpdeskEmail sets HelpdeskEmail field to given value.

### HasHelpdeskEmail

`func (o *PartnerSpec) HasHelpdeskEmail() bool`

HasHelpdeskEmail returns a boolean if a field has been set.

### GetHost

`func (o *PartnerSpec) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PartnerSpec) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PartnerSpec) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PartnerSpec) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetIsTOTPEnabled

`func (o *PartnerSpec) GetIsTOTPEnabled() bool`

GetIsTOTPEnabled returns the IsTOTPEnabled field if non-nil, zero value otherwise.

### GetIsTOTPEnabledOk

`func (o *PartnerSpec) GetIsTOTPEnabledOk() (*bool, bool)`

GetIsTOTPEnabledOk returns a tuple with the IsTOTPEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTOTPEnabled

`func (o *PartnerSpec) SetIsTOTPEnabled(v bool)`

SetIsTOTPEnabled sets IsTOTPEnabled field to given value.

### HasIsTOTPEnabled

`func (o *PartnerSpec) HasIsTOTPEnabled() bool`

HasIsTOTPEnabled returns a boolean if a field has been set.

### GetLogoLink

`func (o *PartnerSpec) GetLogoLink() string`

GetLogoLink returns the LogoLink field if non-nil, zero value otherwise.

### GetLogoLinkOk

`func (o *PartnerSpec) GetLogoLinkOk() (*string, bool)`

GetLogoLinkOk returns a tuple with the LogoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoLink

`func (o *PartnerSpec) SetLogoLink(v string)`

SetLogoLink sets LogoLink field to given value.

### HasLogoLink

`func (o *PartnerSpec) HasLogoLink() bool`

HasLogoLink returns a boolean if a field has been set.

### GetNotificationEmail

`func (o *PartnerSpec) GetNotificationEmail() string`

GetNotificationEmail returns the NotificationEmail field if non-nil, zero value otherwise.

### GetNotificationEmailOk

`func (o *PartnerSpec) GetNotificationEmailOk() (*string, bool)`

GetNotificationEmailOk returns a tuple with the NotificationEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationEmail

`func (o *PartnerSpec) SetNotificationEmail(v string)`

SetNotificationEmail sets NotificationEmail field to given value.

### HasNotificationEmail

`func (o *PartnerSpec) HasNotificationEmail() bool`

HasNotificationEmail returns a boolean if a field has been set.

### GetOpsHost

`func (o *PartnerSpec) GetOpsHost() string`

GetOpsHost returns the OpsHost field if non-nil, zero value otherwise.

### GetOpsHostOk

`func (o *PartnerSpec) GetOpsHostOk() (*string, bool)`

GetOpsHostOk returns a tuple with the OpsHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsHost

`func (o *PartnerSpec) SetOpsHost(v string)`

SetOpsHost sets OpsHost field to given value.

### HasOpsHost

`func (o *PartnerSpec) HasOpsHost() bool`

HasOpsHost returns a boolean if a field has been set.

### GetProductName

`func (o *PartnerSpec) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *PartnerSpec) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *PartnerSpec) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *PartnerSpec) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetSettings

`func (o *PartnerSpec) GetSettings() ExtraData`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PartnerSpec) GetSettingsOk() (*ExtraData, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PartnerSpec) SetSettings(v ExtraData)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PartnerSpec) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSupportTeamName

`func (o *PartnerSpec) GetSupportTeamName() string`

GetSupportTeamName returns the SupportTeamName field if non-nil, zero value otherwise.

### GetSupportTeamNameOk

`func (o *PartnerSpec) GetSupportTeamNameOk() (*string, bool)`

GetSupportTeamNameOk returns a tuple with the SupportTeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTeamName

`func (o *PartnerSpec) SetSupportTeamName(v string)`

SetSupportTeamName sets SupportTeamName field to given value.

### HasSupportTeamName

`func (o *PartnerSpec) HasSupportTeamName() bool`

HasSupportTeamName returns a boolean if a field has been set.

### GetTosLink

`func (o *PartnerSpec) GetTosLink() string`

GetTosLink returns the TosLink field if non-nil, zero value otherwise.

### GetTosLinkOk

`func (o *PartnerSpec) GetTosLinkOk() (*string, bool)`

GetTosLinkOk returns a tuple with the TosLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosLink

`func (o *PartnerSpec) SetTosLink(v string)`

SetTosLink sets TosLink field to given value.

### HasTosLink

`func (o *PartnerSpec) HasTosLink() bool`

HasTosLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


