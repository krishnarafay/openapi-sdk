# NamespaceNetworkPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]ResourceNameAndVersionRef**](ResourceNameAndVersionRef.md) | Group of namespace policy rules | [optional] 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Version** | Pointer to **string** | version of the namespace policy | [optional] 

## Methods

### NewNamespaceNetworkPolicySpec

`func NewNamespaceNetworkPolicySpec() *NamespaceNetworkPolicySpec`

NewNamespaceNetworkPolicySpec instantiates a new NamespaceNetworkPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceNetworkPolicySpecWithDefaults

`func NewNamespaceNetworkPolicySpecWithDefaults() *NamespaceNetworkPolicySpec`

NewNamespaceNetworkPolicySpecWithDefaults instantiates a new NamespaceNetworkPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *NamespaceNetworkPolicySpec) GetRules() []ResourceNameAndVersionRef`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *NamespaceNetworkPolicySpec) GetRulesOk() (*[]ResourceNameAndVersionRef, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *NamespaceNetworkPolicySpec) SetRules(v []ResourceNameAndVersionRef)`

SetRules sets Rules field to given value.

### HasRules

`func (o *NamespaceNetworkPolicySpec) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSharing

`func (o *NamespaceNetworkPolicySpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *NamespaceNetworkPolicySpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *NamespaceNetworkPolicySpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *NamespaceNetworkPolicySpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetVersion

`func (o *NamespaceNetworkPolicySpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NamespaceNetworkPolicySpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NamespaceNetworkPolicySpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NamespaceNetworkPolicySpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


