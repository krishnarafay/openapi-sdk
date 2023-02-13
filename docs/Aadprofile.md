# Aadprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminGroupObjectIDs** | Pointer to **[]string** | The list of AAD group object IDs that will have admin role of the cluster. | [optional] 
**ClientAppID** | Pointer to **string** | The client AAD application ID. | [optional] 
**EnableAzureRBAC** | Pointer to **bool** | Whether to enable Azure RBAC for Kubernetes authorization. | [optional] 
**Managed** | Pointer to **bool** | Whether to enable managed AAD. | [optional] 
**ServerAppID** | Pointer to **string** | The server AAD application ID. | [optional] 
**ServerAppSecret** | Pointer to **string** | The server AAD application secret. | [optional] 
**TenantID** | Pointer to **string** | The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription. | [optional] 

## Methods

### NewAadprofile

`func NewAadprofile() *Aadprofile`

NewAadprofile instantiates a new Aadprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAadprofileWithDefaults

`func NewAadprofileWithDefaults() *Aadprofile`

NewAadprofileWithDefaults instantiates a new Aadprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminGroupObjectIDs

`func (o *Aadprofile) GetAdminGroupObjectIDs() []string`

GetAdminGroupObjectIDs returns the AdminGroupObjectIDs field if non-nil, zero value otherwise.

### GetAdminGroupObjectIDsOk

`func (o *Aadprofile) GetAdminGroupObjectIDsOk() (*[]string, bool)`

GetAdminGroupObjectIDsOk returns a tuple with the AdminGroupObjectIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminGroupObjectIDs

`func (o *Aadprofile) SetAdminGroupObjectIDs(v []string)`

SetAdminGroupObjectIDs sets AdminGroupObjectIDs field to given value.

### HasAdminGroupObjectIDs

`func (o *Aadprofile) HasAdminGroupObjectIDs() bool`

HasAdminGroupObjectIDs returns a boolean if a field has been set.

### GetClientAppID

`func (o *Aadprofile) GetClientAppID() string`

GetClientAppID returns the ClientAppID field if non-nil, zero value otherwise.

### GetClientAppIDOk

`func (o *Aadprofile) GetClientAppIDOk() (*string, bool)`

GetClientAppIDOk returns a tuple with the ClientAppID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAppID

`func (o *Aadprofile) SetClientAppID(v string)`

SetClientAppID sets ClientAppID field to given value.

### HasClientAppID

`func (o *Aadprofile) HasClientAppID() bool`

HasClientAppID returns a boolean if a field has been set.

### GetEnableAzureRBAC

`func (o *Aadprofile) GetEnableAzureRBAC() bool`

GetEnableAzureRBAC returns the EnableAzureRBAC field if non-nil, zero value otherwise.

### GetEnableAzureRBACOk

`func (o *Aadprofile) GetEnableAzureRBACOk() (*bool, bool)`

GetEnableAzureRBACOk returns a tuple with the EnableAzureRBAC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAzureRBAC

`func (o *Aadprofile) SetEnableAzureRBAC(v bool)`

SetEnableAzureRBAC sets EnableAzureRBAC field to given value.

### HasEnableAzureRBAC

`func (o *Aadprofile) HasEnableAzureRBAC() bool`

HasEnableAzureRBAC returns a boolean if a field has been set.

### GetManaged

`func (o *Aadprofile) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *Aadprofile) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *Aadprofile) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *Aadprofile) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetServerAppID

`func (o *Aadprofile) GetServerAppID() string`

GetServerAppID returns the ServerAppID field if non-nil, zero value otherwise.

### GetServerAppIDOk

`func (o *Aadprofile) GetServerAppIDOk() (*string, bool)`

GetServerAppIDOk returns a tuple with the ServerAppID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAppID

`func (o *Aadprofile) SetServerAppID(v string)`

SetServerAppID sets ServerAppID field to given value.

### HasServerAppID

`func (o *Aadprofile) HasServerAppID() bool`

HasServerAppID returns a boolean if a field has been set.

### GetServerAppSecret

`func (o *Aadprofile) GetServerAppSecret() string`

GetServerAppSecret returns the ServerAppSecret field if non-nil, zero value otherwise.

### GetServerAppSecretOk

`func (o *Aadprofile) GetServerAppSecretOk() (*string, bool)`

GetServerAppSecretOk returns a tuple with the ServerAppSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAppSecret

`func (o *Aadprofile) SetServerAppSecret(v string)`

SetServerAppSecret sets ServerAppSecret field to given value.

### HasServerAppSecret

`func (o *Aadprofile) HasServerAppSecret() bool`

HasServerAppSecret returns a boolean if a field has been set.

### GetTenantID

`func (o *Aadprofile) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *Aadprofile) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *Aadprofile) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.

### HasTenantID

`func (o *Aadprofile) HasTenantID() bool`

HasTenantID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


