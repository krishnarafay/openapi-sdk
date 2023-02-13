# Podidentityprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowNetworkPluginKubenet** | Pointer to **bool** | Running in Kubenet is disabled by default due to the security related nature of AAD Pod Identity and the risks of IP spoofing. See https://learn.microsoft.com/en-us/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities for more information. | [optional] 
**Enabled** | Pointer to **bool** | Whether the pod identity addon is enabled. | [optional] 
**UserAssignedIdentities** | Pointer to [**[]Userassignedidentities**](Userassignedidentities.md) | The pod identities to use in the cluster. | [optional] 
**UserAssignedIdentityExceptions** | Pointer to [**[]Userassignedidentityexceptions**](Userassignedidentityexceptions.md) | The pod identity exceptions to allow. | [optional] 

## Methods

### NewPodidentityprofile

`func NewPodidentityprofile() *Podidentityprofile`

NewPodidentityprofile instantiates a new Podidentityprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodidentityprofileWithDefaults

`func NewPodidentityprofileWithDefaults() *Podidentityprofile`

NewPodidentityprofileWithDefaults instantiates a new Podidentityprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowNetworkPluginKubenet

`func (o *Podidentityprofile) GetAllowNetworkPluginKubenet() bool`

GetAllowNetworkPluginKubenet returns the AllowNetworkPluginKubenet field if non-nil, zero value otherwise.

### GetAllowNetworkPluginKubenetOk

`func (o *Podidentityprofile) GetAllowNetworkPluginKubenetOk() (*bool, bool)`

GetAllowNetworkPluginKubenetOk returns a tuple with the AllowNetworkPluginKubenet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNetworkPluginKubenet

`func (o *Podidentityprofile) SetAllowNetworkPluginKubenet(v bool)`

SetAllowNetworkPluginKubenet sets AllowNetworkPluginKubenet field to given value.

### HasAllowNetworkPluginKubenet

`func (o *Podidentityprofile) HasAllowNetworkPluginKubenet() bool`

HasAllowNetworkPluginKubenet returns a boolean if a field has been set.

### GetEnabled

`func (o *Podidentityprofile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Podidentityprofile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Podidentityprofile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Podidentityprofile) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUserAssignedIdentities

`func (o *Podidentityprofile) GetUserAssignedIdentities() []Userassignedidentities`

GetUserAssignedIdentities returns the UserAssignedIdentities field if non-nil, zero value otherwise.

### GetUserAssignedIdentitiesOk

`func (o *Podidentityprofile) GetUserAssignedIdentitiesOk() (*[]Userassignedidentities, bool)`

GetUserAssignedIdentitiesOk returns a tuple with the UserAssignedIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssignedIdentities

`func (o *Podidentityprofile) SetUserAssignedIdentities(v []Userassignedidentities)`

SetUserAssignedIdentities sets UserAssignedIdentities field to given value.

### HasUserAssignedIdentities

`func (o *Podidentityprofile) HasUserAssignedIdentities() bool`

HasUserAssignedIdentities returns a boolean if a field has been set.

### GetUserAssignedIdentityExceptions

`func (o *Podidentityprofile) GetUserAssignedIdentityExceptions() []Userassignedidentityexceptions`

GetUserAssignedIdentityExceptions returns the UserAssignedIdentityExceptions field if non-nil, zero value otherwise.

### GetUserAssignedIdentityExceptionsOk

`func (o *Podidentityprofile) GetUserAssignedIdentityExceptionsOk() (*[]Userassignedidentityexceptions, bool)`

GetUserAssignedIdentityExceptionsOk returns a tuple with the UserAssignedIdentityExceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssignedIdentityExceptions

`func (o *Podidentityprofile) SetUserAssignedIdentityExceptions(v []Userassignedidentityexceptions)`

SetUserAssignedIdentityExceptions sets UserAssignedIdentityExceptions field to given value.

### HasUserAssignedIdentityExceptions

`func (o *Podidentityprofile) HasUserAssignedIdentityExceptions() bool`

HasUserAssignedIdentityExceptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


