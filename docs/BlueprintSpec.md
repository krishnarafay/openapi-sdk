# BlueprintSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | Pointer to [**BlueprintBase**](BlueprintBase.md) |  | [optional] 
**CostProfile** | Pointer to [**CostProfile**](CostProfile.md) |  | [optional] 
**CustomAddons** | Pointer to [**[]BlueprintAddon**](BlueprintAddon.md) | custom addon configuration | [optional] 
**DefaultAddons** | Pointer to [**DefaultAddons**](DefaultAddons.md) |  | [optional] 
**Drift** | Pointer to [**DriftSpec**](DriftSpec.md) |  | [optional] 
**NamespaceConfig** | Pointer to [**NsConfig**](NsConfig.md) |  | [optional] 
**NetworkPolicy** | Pointer to [**NetworkPolicy**](NetworkPolicy.md) |  | [optional] 
**OpaPolicy** | Pointer to [**OPAPolicy**](OPAPolicy.md) |  | [optional] 
**Placement** | Pointer to [**BlueprintPlacement**](BlueprintPlacement.md) |  | [optional] 
**PrivateKubeAPIProxies** | Pointer to [**[]KubeAPIProxyNetwork**](KubeAPIProxyNetwork.md) | private  kubernetes API proxy networks | [optional] 
**Psp** | Pointer to [**BlueprintPSP**](BlueprintPSP.md) |  | [optional] 
**ServiceMesh** | Pointer to [**ServiceMesh**](ServiceMesh.md) |  | [optional] 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Type** | Pointer to **string** | Blueprint Type | [optional] 
**Version** | Pointer to **string** | version of the blueprint | [optional] 

## Methods

### NewBlueprintSpec

`func NewBlueprintSpec() *BlueprintSpec`

NewBlueprintSpec instantiates a new BlueprintSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintSpecWithDefaults

`func NewBlueprintSpecWithDefaults() *BlueprintSpec`

NewBlueprintSpecWithDefaults instantiates a new BlueprintSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *BlueprintSpec) GetBase() BlueprintBase`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *BlueprintSpec) GetBaseOk() (*BlueprintBase, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *BlueprintSpec) SetBase(v BlueprintBase)`

SetBase sets Base field to given value.

### HasBase

`func (o *BlueprintSpec) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetCostProfile

`func (o *BlueprintSpec) GetCostProfile() CostProfile`

GetCostProfile returns the CostProfile field if non-nil, zero value otherwise.

### GetCostProfileOk

`func (o *BlueprintSpec) GetCostProfileOk() (*CostProfile, bool)`

GetCostProfileOk returns a tuple with the CostProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostProfile

`func (o *BlueprintSpec) SetCostProfile(v CostProfile)`

SetCostProfile sets CostProfile field to given value.

### HasCostProfile

`func (o *BlueprintSpec) HasCostProfile() bool`

HasCostProfile returns a boolean if a field has been set.

### GetCustomAddons

`func (o *BlueprintSpec) GetCustomAddons() []BlueprintAddon`

GetCustomAddons returns the CustomAddons field if non-nil, zero value otherwise.

### GetCustomAddonsOk

`func (o *BlueprintSpec) GetCustomAddonsOk() (*[]BlueprintAddon, bool)`

GetCustomAddonsOk returns a tuple with the CustomAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAddons

`func (o *BlueprintSpec) SetCustomAddons(v []BlueprintAddon)`

SetCustomAddons sets CustomAddons field to given value.

### HasCustomAddons

`func (o *BlueprintSpec) HasCustomAddons() bool`

HasCustomAddons returns a boolean if a field has been set.

### GetDefaultAddons

`func (o *BlueprintSpec) GetDefaultAddons() DefaultAddons`

GetDefaultAddons returns the DefaultAddons field if non-nil, zero value otherwise.

### GetDefaultAddonsOk

`func (o *BlueprintSpec) GetDefaultAddonsOk() (*DefaultAddons, bool)`

GetDefaultAddonsOk returns a tuple with the DefaultAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAddons

`func (o *BlueprintSpec) SetDefaultAddons(v DefaultAddons)`

SetDefaultAddons sets DefaultAddons field to given value.

### HasDefaultAddons

`func (o *BlueprintSpec) HasDefaultAddons() bool`

HasDefaultAddons returns a boolean if a field has been set.

### GetDrift

`func (o *BlueprintSpec) GetDrift() DriftSpec`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *BlueprintSpec) GetDriftOk() (*DriftSpec, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *BlueprintSpec) SetDrift(v DriftSpec)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *BlueprintSpec) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetNamespaceConfig

`func (o *BlueprintSpec) GetNamespaceConfig() NsConfig`

GetNamespaceConfig returns the NamespaceConfig field if non-nil, zero value otherwise.

### GetNamespaceConfigOk

`func (o *BlueprintSpec) GetNamespaceConfigOk() (*NsConfig, bool)`

GetNamespaceConfigOk returns a tuple with the NamespaceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceConfig

`func (o *BlueprintSpec) SetNamespaceConfig(v NsConfig)`

SetNamespaceConfig sets NamespaceConfig field to given value.

### HasNamespaceConfig

`func (o *BlueprintSpec) HasNamespaceConfig() bool`

HasNamespaceConfig returns a boolean if a field has been set.

### GetNetworkPolicy

`func (o *BlueprintSpec) GetNetworkPolicy() NetworkPolicy`

GetNetworkPolicy returns the NetworkPolicy field if non-nil, zero value otherwise.

### GetNetworkPolicyOk

`func (o *BlueprintSpec) GetNetworkPolicyOk() (*NetworkPolicy, bool)`

GetNetworkPolicyOk returns a tuple with the NetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPolicy

`func (o *BlueprintSpec) SetNetworkPolicy(v NetworkPolicy)`

SetNetworkPolicy sets NetworkPolicy field to given value.

### HasNetworkPolicy

`func (o *BlueprintSpec) HasNetworkPolicy() bool`

HasNetworkPolicy returns a boolean if a field has been set.

### GetOpaPolicy

`func (o *BlueprintSpec) GetOpaPolicy() OPAPolicy`

GetOpaPolicy returns the OpaPolicy field if non-nil, zero value otherwise.

### GetOpaPolicyOk

`func (o *BlueprintSpec) GetOpaPolicyOk() (*OPAPolicy, bool)`

GetOpaPolicyOk returns a tuple with the OpaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaPolicy

`func (o *BlueprintSpec) SetOpaPolicy(v OPAPolicy)`

SetOpaPolicy sets OpaPolicy field to given value.

### HasOpaPolicy

`func (o *BlueprintSpec) HasOpaPolicy() bool`

HasOpaPolicy returns a boolean if a field has been set.

### GetPlacement

`func (o *BlueprintSpec) GetPlacement() BlueprintPlacement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *BlueprintSpec) GetPlacementOk() (*BlueprintPlacement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *BlueprintSpec) SetPlacement(v BlueprintPlacement)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *BlueprintSpec) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetPrivateKubeAPIProxies

`func (o *BlueprintSpec) GetPrivateKubeAPIProxies() []KubeAPIProxyNetwork`

GetPrivateKubeAPIProxies returns the PrivateKubeAPIProxies field if non-nil, zero value otherwise.

### GetPrivateKubeAPIProxiesOk

`func (o *BlueprintSpec) GetPrivateKubeAPIProxiesOk() (*[]KubeAPIProxyNetwork, bool)`

GetPrivateKubeAPIProxiesOk returns a tuple with the PrivateKubeAPIProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKubeAPIProxies

`func (o *BlueprintSpec) SetPrivateKubeAPIProxies(v []KubeAPIProxyNetwork)`

SetPrivateKubeAPIProxies sets PrivateKubeAPIProxies field to given value.

### HasPrivateKubeAPIProxies

`func (o *BlueprintSpec) HasPrivateKubeAPIProxies() bool`

HasPrivateKubeAPIProxies returns a boolean if a field has been set.

### GetPsp

`func (o *BlueprintSpec) GetPsp() BlueprintPSP`

GetPsp returns the Psp field if non-nil, zero value otherwise.

### GetPspOk

`func (o *BlueprintSpec) GetPspOk() (*BlueprintPSP, bool)`

GetPspOk returns a tuple with the Psp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsp

`func (o *BlueprintSpec) SetPsp(v BlueprintPSP)`

SetPsp sets Psp field to given value.

### HasPsp

`func (o *BlueprintSpec) HasPsp() bool`

HasPsp returns a boolean if a field has been set.

### GetServiceMesh

`func (o *BlueprintSpec) GetServiceMesh() ServiceMesh`

GetServiceMesh returns the ServiceMesh field if non-nil, zero value otherwise.

### GetServiceMeshOk

`func (o *BlueprintSpec) GetServiceMeshOk() (*ServiceMesh, bool)`

GetServiceMeshOk returns a tuple with the ServiceMesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMesh

`func (o *BlueprintSpec) SetServiceMesh(v ServiceMesh)`

SetServiceMesh sets ServiceMesh field to given value.

### HasServiceMesh

`func (o *BlueprintSpec) HasServiceMesh() bool`

HasServiceMesh returns a boolean if a field has been set.

### GetSharing

`func (o *BlueprintSpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *BlueprintSpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *BlueprintSpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *BlueprintSpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetType

`func (o *BlueprintSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *BlueprintSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlueprintSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlueprintSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BlueprintSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


