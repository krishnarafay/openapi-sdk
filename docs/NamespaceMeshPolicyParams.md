# NamespaceMeshPolicyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeshEnabled** | Pointer to **bool** | Service Mesh enabled flag | [optional] 
**Policies** | Pointer to [**[]ResourceNameAndVersionRef**](ResourceNameAndVersionRef.md) | name and version of namespace mesh policy | [optional] 

## Methods

### NewNamespaceMeshPolicyParams

`func NewNamespaceMeshPolicyParams() *NamespaceMeshPolicyParams`

NewNamespaceMeshPolicyParams instantiates a new NamespaceMeshPolicyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceMeshPolicyParamsWithDefaults

`func NewNamespaceMeshPolicyParamsWithDefaults() *NamespaceMeshPolicyParams`

NewNamespaceMeshPolicyParamsWithDefaults instantiates a new NamespaceMeshPolicyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeshEnabled

`func (o *NamespaceMeshPolicyParams) GetMeshEnabled() bool`

GetMeshEnabled returns the MeshEnabled field if non-nil, zero value otherwise.

### GetMeshEnabledOk

`func (o *NamespaceMeshPolicyParams) GetMeshEnabledOk() (*bool, bool)`

GetMeshEnabledOk returns a tuple with the MeshEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshEnabled

`func (o *NamespaceMeshPolicyParams) SetMeshEnabled(v bool)`

SetMeshEnabled sets MeshEnabled field to given value.

### HasMeshEnabled

`func (o *NamespaceMeshPolicyParams) HasMeshEnabled() bool`

HasMeshEnabled returns a boolean if a field has been set.

### GetPolicies

`func (o *NamespaceMeshPolicyParams) GetPolicies() []ResourceNameAndVersionRef`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *NamespaceMeshPolicyParams) GetPoliciesOk() (*[]ResourceNameAndVersionRef, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *NamespaceMeshPolicyParams) SetPolicies(v []ResourceNameAndVersionRef)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *NamespaceMeshPolicyParams) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


