# NetworkPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | flag to specify if network policy is enabled for blueprint | [optional] 
**Policies** | Pointer to [**[]ResourceNameAndVersionRef**](ResourceNameAndVersionRef.md) | name and version of cluster network policy | [optional] 
**Profile** | Pointer to [**ResourceNameAndVersionRef**](ResourceNameAndVersionRef.md) |  | [optional] 

## Methods

### NewNetworkPolicy

`func NewNetworkPolicy() *NetworkPolicy`

NewNetworkPolicy instantiates a new NetworkPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPolicyWithDefaults

`func NewNetworkPolicyWithDefaults() *NetworkPolicy`

NewNetworkPolicyWithDefaults instantiates a new NetworkPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NetworkPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPolicies

`func (o *NetworkPolicy) GetPolicies() []ResourceNameAndVersionRef`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *NetworkPolicy) GetPoliciesOk() (*[]ResourceNameAndVersionRef, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *NetworkPolicy) SetPolicies(v []ResourceNameAndVersionRef)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *NetworkPolicy) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetProfile

`func (o *NetworkPolicy) GetProfile() ResourceNameAndVersionRef`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *NetworkPolicy) GetProfileOk() (*ResourceNameAndVersionRef, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *NetworkPolicy) SetProfile(v ResourceNameAndVersionRef)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *NetworkPolicy) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


