# NetworkPolicyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkPolicyEnabled** | Pointer to **bool** | Network policy enabled flag | [optional] 
**Policies** | Pointer to [**[]ResourceNameAndVersionRef**](ResourceNameAndVersionRef.md) | name and version of namespace network policy | [optional] 

## Methods

### NewNetworkPolicyParams

`func NewNetworkPolicyParams() *NetworkPolicyParams`

NewNetworkPolicyParams instantiates a new NetworkPolicyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPolicyParamsWithDefaults

`func NewNetworkPolicyParamsWithDefaults() *NetworkPolicyParams`

NewNetworkPolicyParamsWithDefaults instantiates a new NetworkPolicyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkPolicyEnabled

`func (o *NetworkPolicyParams) GetNetworkPolicyEnabled() bool`

GetNetworkPolicyEnabled returns the NetworkPolicyEnabled field if non-nil, zero value otherwise.

### GetNetworkPolicyEnabledOk

`func (o *NetworkPolicyParams) GetNetworkPolicyEnabledOk() (*bool, bool)`

GetNetworkPolicyEnabledOk returns a tuple with the NetworkPolicyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPolicyEnabled

`func (o *NetworkPolicyParams) SetNetworkPolicyEnabled(v bool)`

SetNetworkPolicyEnabled sets NetworkPolicyEnabled field to given value.

### HasNetworkPolicyEnabled

`func (o *NetworkPolicyParams) HasNetworkPolicyEnabled() bool`

HasNetworkPolicyEnabled returns a boolean if a field has been set.

### GetPolicies

`func (o *NetworkPolicyParams) GetPolicies() []ResourceNameAndVersionRef`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *NetworkPolicyParams) GetPoliciesOk() (*[]ResourceNameAndVersionRef, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *NetworkPolicyParams) SetPolicies(v []ResourceNameAndVersionRef)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *NetworkPolicyParams) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


