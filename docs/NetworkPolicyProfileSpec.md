# NetworkPolicyProfileSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallationParams** | Pointer to [**InstallationParams**](InstallationParams.md) |  | [optional] 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Version** | Pointer to **string** | version of the profile | [optional] 

## Methods

### NewNetworkPolicyProfileSpec

`func NewNetworkPolicyProfileSpec() *NetworkPolicyProfileSpec`

NewNetworkPolicyProfileSpec instantiates a new NetworkPolicyProfileSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPolicyProfileSpecWithDefaults

`func NewNetworkPolicyProfileSpecWithDefaults() *NetworkPolicyProfileSpec`

NewNetworkPolicyProfileSpecWithDefaults instantiates a new NetworkPolicyProfileSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallationParams

`func (o *NetworkPolicyProfileSpec) GetInstallationParams() InstallationParams`

GetInstallationParams returns the InstallationParams field if non-nil, zero value otherwise.

### GetInstallationParamsOk

`func (o *NetworkPolicyProfileSpec) GetInstallationParamsOk() (*InstallationParams, bool)`

GetInstallationParamsOk returns a tuple with the InstallationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationParams

`func (o *NetworkPolicyProfileSpec) SetInstallationParams(v InstallationParams)`

SetInstallationParams sets InstallationParams field to given value.

### HasInstallationParams

`func (o *NetworkPolicyProfileSpec) HasInstallationParams() bool`

HasInstallationParams returns a boolean if a field has been set.

### GetSharing

`func (o *NetworkPolicyProfileSpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *NetworkPolicyProfileSpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *NetworkPolicyProfileSpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *NetworkPolicyProfileSpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetVersion

`func (o *NetworkPolicyProfileSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NetworkPolicyProfileSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NetworkPolicyProfileSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NetworkPolicyProfileSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


