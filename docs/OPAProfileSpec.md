# OPAProfileSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedNamespaces** | Pointer to [**[]ExcludedNamespaces**](ExcludedNamespaces.md) | list of namespaces excluded from gatekeeper processing | [optional] 
**InstallationParams** | Pointer to [**InstallationParams**](InstallationParams.md) |  | [optional] 
**OpaVersion** | Pointer to **string** | version of the OPA | [optional] 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**SyncObjects** | Pointer to [**[]SyncObject**](SyncObject.md) | list of k8s objects to be synced | [optional] 
**Version** | Pointer to **string** | version of the profile | [optional] 

## Methods

### NewOPAProfileSpec

`func NewOPAProfileSpec() *OPAProfileSpec`

NewOPAProfileSpec instantiates a new OPAProfileSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOPAProfileSpecWithDefaults

`func NewOPAProfileSpecWithDefaults() *OPAProfileSpec`

NewOPAProfileSpecWithDefaults instantiates a new OPAProfileSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedNamespaces

`func (o *OPAProfileSpec) GetExcludedNamespaces() []ExcludedNamespaces`

GetExcludedNamespaces returns the ExcludedNamespaces field if non-nil, zero value otherwise.

### GetExcludedNamespacesOk

`func (o *OPAProfileSpec) GetExcludedNamespacesOk() (*[]ExcludedNamespaces, bool)`

GetExcludedNamespacesOk returns a tuple with the ExcludedNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedNamespaces

`func (o *OPAProfileSpec) SetExcludedNamespaces(v []ExcludedNamespaces)`

SetExcludedNamespaces sets ExcludedNamespaces field to given value.

### HasExcludedNamespaces

`func (o *OPAProfileSpec) HasExcludedNamespaces() bool`

HasExcludedNamespaces returns a boolean if a field has been set.

### GetInstallationParams

`func (o *OPAProfileSpec) GetInstallationParams() InstallationParams`

GetInstallationParams returns the InstallationParams field if non-nil, zero value otherwise.

### GetInstallationParamsOk

`func (o *OPAProfileSpec) GetInstallationParamsOk() (*InstallationParams, bool)`

GetInstallationParamsOk returns a tuple with the InstallationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationParams

`func (o *OPAProfileSpec) SetInstallationParams(v InstallationParams)`

SetInstallationParams sets InstallationParams field to given value.

### HasInstallationParams

`func (o *OPAProfileSpec) HasInstallationParams() bool`

HasInstallationParams returns a boolean if a field has been set.

### GetOpaVersion

`func (o *OPAProfileSpec) GetOpaVersion() string`

GetOpaVersion returns the OpaVersion field if non-nil, zero value otherwise.

### GetOpaVersionOk

`func (o *OPAProfileSpec) GetOpaVersionOk() (*string, bool)`

GetOpaVersionOk returns a tuple with the OpaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaVersion

`func (o *OPAProfileSpec) SetOpaVersion(v string)`

SetOpaVersion sets OpaVersion field to given value.

### HasOpaVersion

`func (o *OPAProfileSpec) HasOpaVersion() bool`

HasOpaVersion returns a boolean if a field has been set.

### GetSharing

`func (o *OPAProfileSpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *OPAProfileSpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *OPAProfileSpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *OPAProfileSpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetSyncObjects

`func (o *OPAProfileSpec) GetSyncObjects() []SyncObject`

GetSyncObjects returns the SyncObjects field if non-nil, zero value otherwise.

### GetSyncObjectsOk

`func (o *OPAProfileSpec) GetSyncObjectsOk() (*[]SyncObject, bool)`

GetSyncObjectsOk returns a tuple with the SyncObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncObjects

`func (o *OPAProfileSpec) SetSyncObjects(v []SyncObject)`

SetSyncObjects sets SyncObjects field to given value.

### HasSyncObjects

`func (o *OPAProfileSpec) HasSyncObjects() bool`

HasSyncObjects returns a boolean if a field has been set.

### GetVersion

`func (o *OPAProfileSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OPAProfileSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OPAProfileSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OPAProfileSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


