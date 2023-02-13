# NamespaceNetworkPolicyRuleSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifact** | Pointer to [**ArtifactSpec**](ArtifactSpec.md) |  | [optional] 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Version** | Pointer to **string** | version of the namespace policy rule | [optional] 

## Methods

### NewNamespaceNetworkPolicyRuleSpec

`func NewNamespaceNetworkPolicyRuleSpec() *NamespaceNetworkPolicyRuleSpec`

NewNamespaceNetworkPolicyRuleSpec instantiates a new NamespaceNetworkPolicyRuleSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceNetworkPolicyRuleSpecWithDefaults

`func NewNamespaceNetworkPolicyRuleSpecWithDefaults() *NamespaceNetworkPolicyRuleSpec`

NewNamespaceNetworkPolicyRuleSpecWithDefaults instantiates a new NamespaceNetworkPolicyRuleSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifact

`func (o *NamespaceNetworkPolicyRuleSpec) GetArtifact() ArtifactSpec`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *NamespaceNetworkPolicyRuleSpec) GetArtifactOk() (*ArtifactSpec, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *NamespaceNetworkPolicyRuleSpec) SetArtifact(v ArtifactSpec)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *NamespaceNetworkPolicyRuleSpec) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetSharing

`func (o *NamespaceNetworkPolicyRuleSpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *NamespaceNetworkPolicyRuleSpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *NamespaceNetworkPolicyRuleSpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *NamespaceNetworkPolicyRuleSpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetVersion

`func (o *NamespaceNetworkPolicyRuleSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NamespaceNetworkPolicyRuleSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NamespaceNetworkPolicyRuleSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NamespaceNetworkPolicyRuleSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


