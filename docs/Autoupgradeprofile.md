# Autoupgradeprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeChannel** | Pointer to **string** | For more information see https://learn.microsoft.com/en-us/azure/aks/upgrade-cluster?tabs&#x3D;azure-cli#set-auto-upgrade-channel. Valid values are rapid, stable, patch, node-image, none. | [optional] [default to "rapid"]

## Methods

### NewAutoupgradeprofile

`func NewAutoupgradeprofile() *Autoupgradeprofile`

NewAutoupgradeprofile instantiates a new Autoupgradeprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoupgradeprofileWithDefaults

`func NewAutoupgradeprofileWithDefaults() *Autoupgradeprofile`

NewAutoupgradeprofileWithDefaults instantiates a new Autoupgradeprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeChannel

`func (o *Autoupgradeprofile) GetUpgradeChannel() string`

GetUpgradeChannel returns the UpgradeChannel field if non-nil, zero value otherwise.

### GetUpgradeChannelOk

`func (o *Autoupgradeprofile) GetUpgradeChannelOk() (*string, bool)`

GetUpgradeChannelOk returns a tuple with the UpgradeChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeChannel

`func (o *Autoupgradeprofile) SetUpgradeChannel(v string)`

SetUpgradeChannel sets UpgradeChannel field to given value.

### HasUpgradeChannel

`func (o *Autoupgradeprofile) HasUpgradeChannel() bool`

HasUpgradeChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


