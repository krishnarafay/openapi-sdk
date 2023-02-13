# Identity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | For more information see use managed identities in AKS. Valid values are SystemAssigned, UserAssigned, None. | [optional] 
**UserAssignedIdentities** | Pointer to **map[string]string** | The keys must be ARM resource IDs in the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. | [optional] 

## Methods

### NewIdentity

`func NewIdentity() *Identity`

NewIdentity instantiates a new Identity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityWithDefaults

`func NewIdentityWithDefaults() *Identity`

NewIdentityWithDefaults instantiates a new Identity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Identity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Identity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Identity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Identity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserAssignedIdentities

`func (o *Identity) GetUserAssignedIdentities() map[string]string`

GetUserAssignedIdentities returns the UserAssignedIdentities field if non-nil, zero value otherwise.

### GetUserAssignedIdentitiesOk

`func (o *Identity) GetUserAssignedIdentitiesOk() (*map[string]string, bool)`

GetUserAssignedIdentitiesOk returns a tuple with the UserAssignedIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssignedIdentities

`func (o *Identity) SetUserAssignedIdentities(v map[string]string)`

SetUserAssignedIdentities sets UserAssignedIdentities field to given value.

### HasUserAssignedIdentities

`func (o *Identity) HasUserAssignedIdentities() bool`

HasUserAssignedIdentities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


