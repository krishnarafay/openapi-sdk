# NamespaceMeshPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]ResourceNameAndVersionRef**](ResourceNameAndVersionRef.md) | Group of namespace mesh rules | [optional] 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Version** | Pointer to **string** | version of the namespace policy | [optional] 

## Methods

### NewNamespaceMeshPolicySpec

`func NewNamespaceMeshPolicySpec() *NamespaceMeshPolicySpec`

NewNamespaceMeshPolicySpec instantiates a new NamespaceMeshPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceMeshPolicySpecWithDefaults

`func NewNamespaceMeshPolicySpecWithDefaults() *NamespaceMeshPolicySpec`

NewNamespaceMeshPolicySpecWithDefaults instantiates a new NamespaceMeshPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *NamespaceMeshPolicySpec) GetRules() []ResourceNameAndVersionRef`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *NamespaceMeshPolicySpec) GetRulesOk() (*[]ResourceNameAndVersionRef, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *NamespaceMeshPolicySpec) SetRules(v []ResourceNameAndVersionRef)`

SetRules sets Rules field to given value.

### HasRules

`func (o *NamespaceMeshPolicySpec) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSharing

`func (o *NamespaceMeshPolicySpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *NamespaceMeshPolicySpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *NamespaceMeshPolicySpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *NamespaceMeshPolicySpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetVersion

`func (o *NamespaceMeshPolicySpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NamespaceMeshPolicySpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NamespaceMeshPolicySpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NamespaceMeshPolicySpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


