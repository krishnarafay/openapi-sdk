# NamespaceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifact** | Pointer to [**NamespaceSpecArtifact**](NamespaceSpecArtifact.md) |  | [optional] 
**Drift** | [**DriftSpec**](DriftSpec.md) |  | 
**LimitRange** | Pointer to [**NamespaceLimitRange**](NamespaceLimitRange.md) |  | [optional] 
**NamespaceMeshPolicyParams** | Pointer to [**NamespaceMeshPolicyParams**](NamespaceMeshPolicyParams.md) |  | [optional] 
**NetworkPolicyParams** | Pointer to [**NetworkPolicyParams**](NetworkPolicyParams.md) |  | [optional] 
**Placement** | [**PlacementSpec**](PlacementSpec.md) |  | 
**Psp** | [**NamespacePSP**](NamespacePSP.md) |  | 
**ResourceQuotas** | Pointer to [**NamespaceResourceQuotas**](NamespaceResourceQuotas.md) |  | [optional] 

## Methods

### NewNamespaceSpec

`func NewNamespaceSpec(drift DriftSpec, placement PlacementSpec, psp NamespacePSP, ) *NamespaceSpec`

NewNamespaceSpec instantiates a new NamespaceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceSpecWithDefaults

`func NewNamespaceSpecWithDefaults() *NamespaceSpec`

NewNamespaceSpecWithDefaults instantiates a new NamespaceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifact

`func (o *NamespaceSpec) GetArtifact() NamespaceSpecArtifact`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *NamespaceSpec) GetArtifactOk() (*NamespaceSpecArtifact, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *NamespaceSpec) SetArtifact(v NamespaceSpecArtifact)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *NamespaceSpec) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetDrift

`func (o *NamespaceSpec) GetDrift() DriftSpec`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *NamespaceSpec) GetDriftOk() (*DriftSpec, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *NamespaceSpec) SetDrift(v DriftSpec)`

SetDrift sets Drift field to given value.


### GetLimitRange

`func (o *NamespaceSpec) GetLimitRange() NamespaceLimitRange`

GetLimitRange returns the LimitRange field if non-nil, zero value otherwise.

### GetLimitRangeOk

`func (o *NamespaceSpec) GetLimitRangeOk() (*NamespaceLimitRange, bool)`

GetLimitRangeOk returns a tuple with the LimitRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitRange

`func (o *NamespaceSpec) SetLimitRange(v NamespaceLimitRange)`

SetLimitRange sets LimitRange field to given value.

### HasLimitRange

`func (o *NamespaceSpec) HasLimitRange() bool`

HasLimitRange returns a boolean if a field has been set.

### GetNamespaceMeshPolicyParams

`func (o *NamespaceSpec) GetNamespaceMeshPolicyParams() NamespaceMeshPolicyParams`

GetNamespaceMeshPolicyParams returns the NamespaceMeshPolicyParams field if non-nil, zero value otherwise.

### GetNamespaceMeshPolicyParamsOk

`func (o *NamespaceSpec) GetNamespaceMeshPolicyParamsOk() (*NamespaceMeshPolicyParams, bool)`

GetNamespaceMeshPolicyParamsOk returns a tuple with the NamespaceMeshPolicyParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceMeshPolicyParams

`func (o *NamespaceSpec) SetNamespaceMeshPolicyParams(v NamespaceMeshPolicyParams)`

SetNamespaceMeshPolicyParams sets NamespaceMeshPolicyParams field to given value.

### HasNamespaceMeshPolicyParams

`func (o *NamespaceSpec) HasNamespaceMeshPolicyParams() bool`

HasNamespaceMeshPolicyParams returns a boolean if a field has been set.

### GetNetworkPolicyParams

`func (o *NamespaceSpec) GetNetworkPolicyParams() NetworkPolicyParams`

GetNetworkPolicyParams returns the NetworkPolicyParams field if non-nil, zero value otherwise.

### GetNetworkPolicyParamsOk

`func (o *NamespaceSpec) GetNetworkPolicyParamsOk() (*NetworkPolicyParams, bool)`

GetNetworkPolicyParamsOk returns a tuple with the NetworkPolicyParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPolicyParams

`func (o *NamespaceSpec) SetNetworkPolicyParams(v NetworkPolicyParams)`

SetNetworkPolicyParams sets NetworkPolicyParams field to given value.

### HasNetworkPolicyParams

`func (o *NamespaceSpec) HasNetworkPolicyParams() bool`

HasNetworkPolicyParams returns a boolean if a field has been set.

### GetPlacement

`func (o *NamespaceSpec) GetPlacement() PlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *NamespaceSpec) GetPlacementOk() (*PlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *NamespaceSpec) SetPlacement(v PlacementSpec)`

SetPlacement sets Placement field to given value.


### GetPsp

`func (o *NamespaceSpec) GetPsp() NamespacePSP`

GetPsp returns the Psp field if non-nil, zero value otherwise.

### GetPspOk

`func (o *NamespaceSpec) GetPspOk() (*NamespacePSP, bool)`

GetPspOk returns a tuple with the Psp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsp

`func (o *NamespaceSpec) SetPsp(v NamespacePSP)`

SetPsp sets Psp field to given value.


### GetResourceQuotas

`func (o *NamespaceSpec) GetResourceQuotas() NamespaceResourceQuotas`

GetResourceQuotas returns the ResourceQuotas field if non-nil, zero value otherwise.

### GetResourceQuotasOk

`func (o *NamespaceSpec) GetResourceQuotasOk() (*NamespaceResourceQuotas, bool)`

GetResourceQuotasOk returns a tuple with the ResourceQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceQuotas

`func (o *NamespaceSpec) SetResourceQuotas(v NamespaceResourceQuotas)`

SetResourceQuotas sets ResourceQuotas field to given value.

### HasResourceQuotas

`func (o *NamespaceSpec) HasResourceQuotas() bool`

HasResourceQuotas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


