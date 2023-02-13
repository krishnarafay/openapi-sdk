# Windowsprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminUsername** | Pointer to **string** | Specifies the name of the administrator account, Restriction: Cannot end in &#39;.&#39;, Minimum-length: 1 character, Max-length: 20 characters | [optional] 
**EnableCSIProxy** | Pointer to **bool** | For more details on CSI proxy, see the CSI proxy GitHub repo. | [optional] 
**GmsaProfile** | Pointer to [**Gmsaprofile**](Gmsaprofile.md) |  | [optional] 
**LicenseType** | Pointer to **string** | The license type to use for Windows VMs. See Azure Hybrid User Benefits for more details. Valid values are None, Windows_Server. | [optional] 

## Methods

### NewWindowsprofile

`func NewWindowsprofile() *Windowsprofile`

NewWindowsprofile instantiates a new Windowsprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsprofileWithDefaults

`func NewWindowsprofileWithDefaults() *Windowsprofile`

NewWindowsprofileWithDefaults instantiates a new Windowsprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminUsername

`func (o *Windowsprofile) GetAdminUsername() string`

GetAdminUsername returns the AdminUsername field if non-nil, zero value otherwise.

### GetAdminUsernameOk

`func (o *Windowsprofile) GetAdminUsernameOk() (*string, bool)`

GetAdminUsernameOk returns a tuple with the AdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUsername

`func (o *Windowsprofile) SetAdminUsername(v string)`

SetAdminUsername sets AdminUsername field to given value.

### HasAdminUsername

`func (o *Windowsprofile) HasAdminUsername() bool`

HasAdminUsername returns a boolean if a field has been set.

### GetEnableCSIProxy

`func (o *Windowsprofile) GetEnableCSIProxy() bool`

GetEnableCSIProxy returns the EnableCSIProxy field if non-nil, zero value otherwise.

### GetEnableCSIProxyOk

`func (o *Windowsprofile) GetEnableCSIProxyOk() (*bool, bool)`

GetEnableCSIProxyOk returns a tuple with the EnableCSIProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCSIProxy

`func (o *Windowsprofile) SetEnableCSIProxy(v bool)`

SetEnableCSIProxy sets EnableCSIProxy field to given value.

### HasEnableCSIProxy

`func (o *Windowsprofile) HasEnableCSIProxy() bool`

HasEnableCSIProxy returns a boolean if a field has been set.

### GetGmsaProfile

`func (o *Windowsprofile) GetGmsaProfile() Gmsaprofile`

GetGmsaProfile returns the GmsaProfile field if non-nil, zero value otherwise.

### GetGmsaProfileOk

`func (o *Windowsprofile) GetGmsaProfileOk() (*Gmsaprofile, bool)`

GetGmsaProfileOk returns a tuple with the GmsaProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmsaProfile

`func (o *Windowsprofile) SetGmsaProfile(v Gmsaprofile)`

SetGmsaProfile sets GmsaProfile field to given value.

### HasGmsaProfile

`func (o *Windowsprofile) HasGmsaProfile() bool`

HasGmsaProfile returns a boolean if a field has been set.

### GetLicenseType

`func (o *Windowsprofile) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *Windowsprofile) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *Windowsprofile) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *Windowsprofile) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


