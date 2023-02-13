# NodePoolProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZones** | Pointer to **[]string** | The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is VirtualMachineScaleSets. | [optional] 
**Count** | Pointer to **int32** | Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1. | [optional] 
**CreationData** | Pointer to [**CreationData**](CreationData.md) |  | [optional] 
**EnableAutoScaling** | Pointer to **bool** | Whether to enable auto-scaler | [optional] 
**EnableEncryptionAtHost** | Pointer to **bool** | This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption | [optional] 
**EnableFIPS** | Pointer to **bool** | See Add a FIPS-enabled node pool for more details. | [optional] 
**EnableNodePublicIP** | Pointer to **bool** | Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see assigning a public IP per node. The default is false. | [optional] 
**EnableUltraSSD** | Pointer to **bool** | Whether to enable UltraSSD | [optional] 
**GpuInstanceProfile** | Pointer to **string** | GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU. Valid values are MIG1g, MIG2g, MIG3g, MIG4g, MIG7g. | [optional] 
**HostGroupID** | Pointer to **string** | This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}. For more information see Azure dedicated hosts. | [optional] 
**KubeletConfig** | Pointer to [**Kubeletconfig**](Kubeletconfig.md) |  | [optional] 
**KubeletDiskType** | Pointer to **string** | Valid values are OS, Temporary. | [optional] 
**LinuxOSConfig** | Pointer to [**Linuxosconfig**](Linuxosconfig.md) |  | [optional] 
**MaxCount** | Pointer to **int32** | The maximum number of nodes for auto-scaling | [optional] 
**MaxPods** | Pointer to **int32** |  The maximum number of pods that can run on a node. | [optional] 
**MinCount** | Pointer to **int32** | The minimum number of nodes for auto-scaling | [optional] 
**Mode** | Pointer to **string** | Valid values are System, User. | [optional] 
**NodeLabels** | Pointer to **map[string]string** | The node labels to be persisted across all nodes in agent pool. | [optional] 
**NodePublicIPPrefixID** | Pointer to **string** | This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName} | [optional] 
**NodeTaints** | Pointer to **[]string** | The taints added to new nodes during node pool create and scale. For example, key&#x3D;value:NoSchedule. | [optional] 
**OrchestratorVersion** | Pointer to **string** | Both patch version (e.g. 1.20.13) and (e.g. 1.20) are supported. When is specified, the latest supported GA patch version is chosen automatically. Updating the cluster with the same once it has been created (e.g. 1.14.x -&gt; 1.14) will not trigger an upgrade, even if a newer patch version is available. As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see upgrading a node pool. | [optional] 
**OsDiskSizeGB** | Pointer to **int32** | OS Disk Size in GB to be used to specify the disk size for every machine in the master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified. | [optional] 
**OsDiskType** | Pointer to **string** | Valid values are Managed, Ephemeral. | [optional] 
**OsSKU** | Pointer to **string** | Valid values are Ubuntu, CBLMariner, Windows2019, Windows2022. | [optional] 
**OsType** | Pointer to **string** | Valid values are Linux, Windows. | [optional] 
**PodSubnetID** | Pointer to **string** | If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName} | [optional] 
**PowerState** | Pointer to [**Powerstate**](Powerstate.md) |  | [optional] 
**ProximityPlacementGroupID** | Pointer to **string** | The ID for Proximity Placement Group. | [optional] 
**ScaleDownMode** | Pointer to **string** | This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete. Valid values are Delete, Deallocate. | [optional] 
**ScaleSetEvictionPolicy** | Pointer to **string** | This cannot be specified unless the scaleSetPriority is Spot. If not specified, the default is Delete. Valid values are Delete, Deallocate. | [optional] 
**ScaleSetPriority** | Pointer to **string** | The Virtual Machine Scale Set priority. If not specified, the default is Regular. Valid values are Spot, Regular. | [optional] 
**SpotMaxPrice** | Pointer to **int32** | Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see spot VMs pricing | [optional] 
**Tags** | Pointer to **map[string]string** | The tags to be persisted on the agent pool virtual machine scale set. | [optional] 
**Type** | Pointer to **string** | Valid values are VirtualMachineScaleSets, AvailabilitySet. | [optional] 
**UpgradeSettings** | Pointer to [**Upgradesettings**](Upgradesettings.md) |  | [optional] 
**VmSize** | Pointer to **string** | VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions | [optional] 
**VnetSubnetID** | Pointer to **string** | If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName} | [optional] 
**WorkloadRuntime** | Pointer to **string** | Valid values are OCIContainer, WasmWasi. | [optional] 

## Methods

### NewNodePoolProperties

`func NewNodePoolProperties() *NodePoolProperties`

NewNodePoolProperties instantiates a new NodePoolProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePoolPropertiesWithDefaults

`func NewNodePoolPropertiesWithDefaults() *NodePoolProperties`

NewNodePoolPropertiesWithDefaults instantiates a new NodePoolProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZones

`func (o *NodePoolProperties) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *NodePoolProperties) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *NodePoolProperties) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *NodePoolProperties) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.

### GetCount

`func (o *NodePoolProperties) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NodePoolProperties) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NodePoolProperties) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *NodePoolProperties) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreationData

`func (o *NodePoolProperties) GetCreationData() CreationData`

GetCreationData returns the CreationData field if non-nil, zero value otherwise.

### GetCreationDataOk

`func (o *NodePoolProperties) GetCreationDataOk() (*CreationData, bool)`

GetCreationDataOk returns a tuple with the CreationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationData

`func (o *NodePoolProperties) SetCreationData(v CreationData)`

SetCreationData sets CreationData field to given value.

### HasCreationData

`func (o *NodePoolProperties) HasCreationData() bool`

HasCreationData returns a boolean if a field has been set.

### GetEnableAutoScaling

`func (o *NodePoolProperties) GetEnableAutoScaling() bool`

GetEnableAutoScaling returns the EnableAutoScaling field if non-nil, zero value otherwise.

### GetEnableAutoScalingOk

`func (o *NodePoolProperties) GetEnableAutoScalingOk() (*bool, bool)`

GetEnableAutoScalingOk returns a tuple with the EnableAutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoScaling

`func (o *NodePoolProperties) SetEnableAutoScaling(v bool)`

SetEnableAutoScaling sets EnableAutoScaling field to given value.

### HasEnableAutoScaling

`func (o *NodePoolProperties) HasEnableAutoScaling() bool`

HasEnableAutoScaling returns a boolean if a field has been set.

### GetEnableEncryptionAtHost

`func (o *NodePoolProperties) GetEnableEncryptionAtHost() bool`

GetEnableEncryptionAtHost returns the EnableEncryptionAtHost field if non-nil, zero value otherwise.

### GetEnableEncryptionAtHostOk

`func (o *NodePoolProperties) GetEnableEncryptionAtHostOk() (*bool, bool)`

GetEnableEncryptionAtHostOk returns a tuple with the EnableEncryptionAtHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEncryptionAtHost

`func (o *NodePoolProperties) SetEnableEncryptionAtHost(v bool)`

SetEnableEncryptionAtHost sets EnableEncryptionAtHost field to given value.

### HasEnableEncryptionAtHost

`func (o *NodePoolProperties) HasEnableEncryptionAtHost() bool`

HasEnableEncryptionAtHost returns a boolean if a field has been set.

### GetEnableFIPS

`func (o *NodePoolProperties) GetEnableFIPS() bool`

GetEnableFIPS returns the EnableFIPS field if non-nil, zero value otherwise.

### GetEnableFIPSOk

`func (o *NodePoolProperties) GetEnableFIPSOk() (*bool, bool)`

GetEnableFIPSOk returns a tuple with the EnableFIPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFIPS

`func (o *NodePoolProperties) SetEnableFIPS(v bool)`

SetEnableFIPS sets EnableFIPS field to given value.

### HasEnableFIPS

`func (o *NodePoolProperties) HasEnableFIPS() bool`

HasEnableFIPS returns a boolean if a field has been set.

### GetEnableNodePublicIP

`func (o *NodePoolProperties) GetEnableNodePublicIP() bool`

GetEnableNodePublicIP returns the EnableNodePublicIP field if non-nil, zero value otherwise.

### GetEnableNodePublicIPOk

`func (o *NodePoolProperties) GetEnableNodePublicIPOk() (*bool, bool)`

GetEnableNodePublicIPOk returns a tuple with the EnableNodePublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNodePublicIP

`func (o *NodePoolProperties) SetEnableNodePublicIP(v bool)`

SetEnableNodePublicIP sets EnableNodePublicIP field to given value.

### HasEnableNodePublicIP

`func (o *NodePoolProperties) HasEnableNodePublicIP() bool`

HasEnableNodePublicIP returns a boolean if a field has been set.

### GetEnableUltraSSD

`func (o *NodePoolProperties) GetEnableUltraSSD() bool`

GetEnableUltraSSD returns the EnableUltraSSD field if non-nil, zero value otherwise.

### GetEnableUltraSSDOk

`func (o *NodePoolProperties) GetEnableUltraSSDOk() (*bool, bool)`

GetEnableUltraSSDOk returns a tuple with the EnableUltraSSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUltraSSD

`func (o *NodePoolProperties) SetEnableUltraSSD(v bool)`

SetEnableUltraSSD sets EnableUltraSSD field to given value.

### HasEnableUltraSSD

`func (o *NodePoolProperties) HasEnableUltraSSD() bool`

HasEnableUltraSSD returns a boolean if a field has been set.

### GetGpuInstanceProfile

`func (o *NodePoolProperties) GetGpuInstanceProfile() string`

GetGpuInstanceProfile returns the GpuInstanceProfile field if non-nil, zero value otherwise.

### GetGpuInstanceProfileOk

`func (o *NodePoolProperties) GetGpuInstanceProfileOk() (*string, bool)`

GetGpuInstanceProfileOk returns a tuple with the GpuInstanceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuInstanceProfile

`func (o *NodePoolProperties) SetGpuInstanceProfile(v string)`

SetGpuInstanceProfile sets GpuInstanceProfile field to given value.

### HasGpuInstanceProfile

`func (o *NodePoolProperties) HasGpuInstanceProfile() bool`

HasGpuInstanceProfile returns a boolean if a field has been set.

### GetHostGroupID

`func (o *NodePoolProperties) GetHostGroupID() string`

GetHostGroupID returns the HostGroupID field if non-nil, zero value otherwise.

### GetHostGroupIDOk

`func (o *NodePoolProperties) GetHostGroupIDOk() (*string, bool)`

GetHostGroupIDOk returns a tuple with the HostGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroupID

`func (o *NodePoolProperties) SetHostGroupID(v string)`

SetHostGroupID sets HostGroupID field to given value.

### HasHostGroupID

`func (o *NodePoolProperties) HasHostGroupID() bool`

HasHostGroupID returns a boolean if a field has been set.

### GetKubeletConfig

`func (o *NodePoolProperties) GetKubeletConfig() Kubeletconfig`

GetKubeletConfig returns the KubeletConfig field if non-nil, zero value otherwise.

### GetKubeletConfigOk

`func (o *NodePoolProperties) GetKubeletConfigOk() (*Kubeletconfig, bool)`

GetKubeletConfigOk returns a tuple with the KubeletConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletConfig

`func (o *NodePoolProperties) SetKubeletConfig(v Kubeletconfig)`

SetKubeletConfig sets KubeletConfig field to given value.

### HasKubeletConfig

`func (o *NodePoolProperties) HasKubeletConfig() bool`

HasKubeletConfig returns a boolean if a field has been set.

### GetKubeletDiskType

`func (o *NodePoolProperties) GetKubeletDiskType() string`

GetKubeletDiskType returns the KubeletDiskType field if non-nil, zero value otherwise.

### GetKubeletDiskTypeOk

`func (o *NodePoolProperties) GetKubeletDiskTypeOk() (*string, bool)`

GetKubeletDiskTypeOk returns a tuple with the KubeletDiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletDiskType

`func (o *NodePoolProperties) SetKubeletDiskType(v string)`

SetKubeletDiskType sets KubeletDiskType field to given value.

### HasKubeletDiskType

`func (o *NodePoolProperties) HasKubeletDiskType() bool`

HasKubeletDiskType returns a boolean if a field has been set.

### GetLinuxOSConfig

`func (o *NodePoolProperties) GetLinuxOSConfig() Linuxosconfig`

GetLinuxOSConfig returns the LinuxOSConfig field if non-nil, zero value otherwise.

### GetLinuxOSConfigOk

`func (o *NodePoolProperties) GetLinuxOSConfigOk() (*Linuxosconfig, bool)`

GetLinuxOSConfigOk returns a tuple with the LinuxOSConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxOSConfig

`func (o *NodePoolProperties) SetLinuxOSConfig(v Linuxosconfig)`

SetLinuxOSConfig sets LinuxOSConfig field to given value.

### HasLinuxOSConfig

`func (o *NodePoolProperties) HasLinuxOSConfig() bool`

HasLinuxOSConfig returns a boolean if a field has been set.

### GetMaxCount

`func (o *NodePoolProperties) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *NodePoolProperties) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *NodePoolProperties) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *NodePoolProperties) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetMaxPods

`func (o *NodePoolProperties) GetMaxPods() int32`

GetMaxPods returns the MaxPods field if non-nil, zero value otherwise.

### GetMaxPodsOk

`func (o *NodePoolProperties) GetMaxPodsOk() (*int32, bool)`

GetMaxPodsOk returns a tuple with the MaxPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPods

`func (o *NodePoolProperties) SetMaxPods(v int32)`

SetMaxPods sets MaxPods field to given value.

### HasMaxPods

`func (o *NodePoolProperties) HasMaxPods() bool`

HasMaxPods returns a boolean if a field has been set.

### GetMinCount

`func (o *NodePoolProperties) GetMinCount() int32`

GetMinCount returns the MinCount field if non-nil, zero value otherwise.

### GetMinCountOk

`func (o *NodePoolProperties) GetMinCountOk() (*int32, bool)`

GetMinCountOk returns a tuple with the MinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCount

`func (o *NodePoolProperties) SetMinCount(v int32)`

SetMinCount sets MinCount field to given value.

### HasMinCount

`func (o *NodePoolProperties) HasMinCount() bool`

HasMinCount returns a boolean if a field has been set.

### GetMode

`func (o *NodePoolProperties) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *NodePoolProperties) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *NodePoolProperties) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *NodePoolProperties) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNodeLabels

`func (o *NodePoolProperties) GetNodeLabels() map[string]string`

GetNodeLabels returns the NodeLabels field if non-nil, zero value otherwise.

### GetNodeLabelsOk

`func (o *NodePoolProperties) GetNodeLabelsOk() (*map[string]string, bool)`

GetNodeLabelsOk returns a tuple with the NodeLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeLabels

`func (o *NodePoolProperties) SetNodeLabels(v map[string]string)`

SetNodeLabels sets NodeLabels field to given value.

### HasNodeLabels

`func (o *NodePoolProperties) HasNodeLabels() bool`

HasNodeLabels returns a boolean if a field has been set.

### GetNodePublicIPPrefixID

`func (o *NodePoolProperties) GetNodePublicIPPrefixID() string`

GetNodePublicIPPrefixID returns the NodePublicIPPrefixID field if non-nil, zero value otherwise.

### GetNodePublicIPPrefixIDOk

`func (o *NodePoolProperties) GetNodePublicIPPrefixIDOk() (*string, bool)`

GetNodePublicIPPrefixIDOk returns a tuple with the NodePublicIPPrefixID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePublicIPPrefixID

`func (o *NodePoolProperties) SetNodePublicIPPrefixID(v string)`

SetNodePublicIPPrefixID sets NodePublicIPPrefixID field to given value.

### HasNodePublicIPPrefixID

`func (o *NodePoolProperties) HasNodePublicIPPrefixID() bool`

HasNodePublicIPPrefixID returns a boolean if a field has been set.

### GetNodeTaints

`func (o *NodePoolProperties) GetNodeTaints() []string`

GetNodeTaints returns the NodeTaints field if non-nil, zero value otherwise.

### GetNodeTaintsOk

`func (o *NodePoolProperties) GetNodeTaintsOk() (*[]string, bool)`

GetNodeTaintsOk returns a tuple with the NodeTaints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeTaints

`func (o *NodePoolProperties) SetNodeTaints(v []string)`

SetNodeTaints sets NodeTaints field to given value.

### HasNodeTaints

`func (o *NodePoolProperties) HasNodeTaints() bool`

HasNodeTaints returns a boolean if a field has been set.

### GetOrchestratorVersion

`func (o *NodePoolProperties) GetOrchestratorVersion() string`

GetOrchestratorVersion returns the OrchestratorVersion field if non-nil, zero value otherwise.

### GetOrchestratorVersionOk

`func (o *NodePoolProperties) GetOrchestratorVersionOk() (*string, bool)`

GetOrchestratorVersionOk returns a tuple with the OrchestratorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestratorVersion

`func (o *NodePoolProperties) SetOrchestratorVersion(v string)`

SetOrchestratorVersion sets OrchestratorVersion field to given value.

### HasOrchestratorVersion

`func (o *NodePoolProperties) HasOrchestratorVersion() bool`

HasOrchestratorVersion returns a boolean if a field has been set.

### GetOsDiskSizeGB

`func (o *NodePoolProperties) GetOsDiskSizeGB() int32`

GetOsDiskSizeGB returns the OsDiskSizeGB field if non-nil, zero value otherwise.

### GetOsDiskSizeGBOk

`func (o *NodePoolProperties) GetOsDiskSizeGBOk() (*int32, bool)`

GetOsDiskSizeGBOk returns a tuple with the OsDiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDiskSizeGB

`func (o *NodePoolProperties) SetOsDiskSizeGB(v int32)`

SetOsDiskSizeGB sets OsDiskSizeGB field to given value.

### HasOsDiskSizeGB

`func (o *NodePoolProperties) HasOsDiskSizeGB() bool`

HasOsDiskSizeGB returns a boolean if a field has been set.

### GetOsDiskType

`func (o *NodePoolProperties) GetOsDiskType() string`

GetOsDiskType returns the OsDiskType field if non-nil, zero value otherwise.

### GetOsDiskTypeOk

`func (o *NodePoolProperties) GetOsDiskTypeOk() (*string, bool)`

GetOsDiskTypeOk returns a tuple with the OsDiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDiskType

`func (o *NodePoolProperties) SetOsDiskType(v string)`

SetOsDiskType sets OsDiskType field to given value.

### HasOsDiskType

`func (o *NodePoolProperties) HasOsDiskType() bool`

HasOsDiskType returns a boolean if a field has been set.

### GetOsSKU

`func (o *NodePoolProperties) GetOsSKU() string`

GetOsSKU returns the OsSKU field if non-nil, zero value otherwise.

### GetOsSKUOk

`func (o *NodePoolProperties) GetOsSKUOk() (*string, bool)`

GetOsSKUOk returns a tuple with the OsSKU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsSKU

`func (o *NodePoolProperties) SetOsSKU(v string)`

SetOsSKU sets OsSKU field to given value.

### HasOsSKU

`func (o *NodePoolProperties) HasOsSKU() bool`

HasOsSKU returns a boolean if a field has been set.

### GetOsType

`func (o *NodePoolProperties) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *NodePoolProperties) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *NodePoolProperties) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *NodePoolProperties) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetPodSubnetID

`func (o *NodePoolProperties) GetPodSubnetID() string`

GetPodSubnetID returns the PodSubnetID field if non-nil, zero value otherwise.

### GetPodSubnetIDOk

`func (o *NodePoolProperties) GetPodSubnetIDOk() (*string, bool)`

GetPodSubnetIDOk returns a tuple with the PodSubnetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodSubnetID

`func (o *NodePoolProperties) SetPodSubnetID(v string)`

SetPodSubnetID sets PodSubnetID field to given value.

### HasPodSubnetID

`func (o *NodePoolProperties) HasPodSubnetID() bool`

HasPodSubnetID returns a boolean if a field has been set.

### GetPowerState

`func (o *NodePoolProperties) GetPowerState() Powerstate`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *NodePoolProperties) GetPowerStateOk() (*Powerstate, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *NodePoolProperties) SetPowerState(v Powerstate)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *NodePoolProperties) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetProximityPlacementGroupID

`func (o *NodePoolProperties) GetProximityPlacementGroupID() string`

GetProximityPlacementGroupID returns the ProximityPlacementGroupID field if non-nil, zero value otherwise.

### GetProximityPlacementGroupIDOk

`func (o *NodePoolProperties) GetProximityPlacementGroupIDOk() (*string, bool)`

GetProximityPlacementGroupIDOk returns a tuple with the ProximityPlacementGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProximityPlacementGroupID

`func (o *NodePoolProperties) SetProximityPlacementGroupID(v string)`

SetProximityPlacementGroupID sets ProximityPlacementGroupID field to given value.

### HasProximityPlacementGroupID

`func (o *NodePoolProperties) HasProximityPlacementGroupID() bool`

HasProximityPlacementGroupID returns a boolean if a field has been set.

### GetScaleDownMode

`func (o *NodePoolProperties) GetScaleDownMode() string`

GetScaleDownMode returns the ScaleDownMode field if non-nil, zero value otherwise.

### GetScaleDownModeOk

`func (o *NodePoolProperties) GetScaleDownModeOk() (*string, bool)`

GetScaleDownModeOk returns a tuple with the ScaleDownMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleDownMode

`func (o *NodePoolProperties) SetScaleDownMode(v string)`

SetScaleDownMode sets ScaleDownMode field to given value.

### HasScaleDownMode

`func (o *NodePoolProperties) HasScaleDownMode() bool`

HasScaleDownMode returns a boolean if a field has been set.

### GetScaleSetEvictionPolicy

`func (o *NodePoolProperties) GetScaleSetEvictionPolicy() string`

GetScaleSetEvictionPolicy returns the ScaleSetEvictionPolicy field if non-nil, zero value otherwise.

### GetScaleSetEvictionPolicyOk

`func (o *NodePoolProperties) GetScaleSetEvictionPolicyOk() (*string, bool)`

GetScaleSetEvictionPolicyOk returns a tuple with the ScaleSetEvictionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleSetEvictionPolicy

`func (o *NodePoolProperties) SetScaleSetEvictionPolicy(v string)`

SetScaleSetEvictionPolicy sets ScaleSetEvictionPolicy field to given value.

### HasScaleSetEvictionPolicy

`func (o *NodePoolProperties) HasScaleSetEvictionPolicy() bool`

HasScaleSetEvictionPolicy returns a boolean if a field has been set.

### GetScaleSetPriority

`func (o *NodePoolProperties) GetScaleSetPriority() string`

GetScaleSetPriority returns the ScaleSetPriority field if non-nil, zero value otherwise.

### GetScaleSetPriorityOk

`func (o *NodePoolProperties) GetScaleSetPriorityOk() (*string, bool)`

GetScaleSetPriorityOk returns a tuple with the ScaleSetPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleSetPriority

`func (o *NodePoolProperties) SetScaleSetPriority(v string)`

SetScaleSetPriority sets ScaleSetPriority field to given value.

### HasScaleSetPriority

`func (o *NodePoolProperties) HasScaleSetPriority() bool`

HasScaleSetPriority returns a boolean if a field has been set.

### GetSpotMaxPrice

`func (o *NodePoolProperties) GetSpotMaxPrice() int32`

GetSpotMaxPrice returns the SpotMaxPrice field if non-nil, zero value otherwise.

### GetSpotMaxPriceOk

`func (o *NodePoolProperties) GetSpotMaxPriceOk() (*int32, bool)`

GetSpotMaxPriceOk returns a tuple with the SpotMaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotMaxPrice

`func (o *NodePoolProperties) SetSpotMaxPrice(v int32)`

SetSpotMaxPrice sets SpotMaxPrice field to given value.

### HasSpotMaxPrice

`func (o *NodePoolProperties) HasSpotMaxPrice() bool`

HasSpotMaxPrice returns a boolean if a field has been set.

### GetTags

`func (o *NodePoolProperties) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NodePoolProperties) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NodePoolProperties) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NodePoolProperties) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *NodePoolProperties) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodePoolProperties) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodePoolProperties) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NodePoolProperties) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpgradeSettings

`func (o *NodePoolProperties) GetUpgradeSettings() Upgradesettings`

GetUpgradeSettings returns the UpgradeSettings field if non-nil, zero value otherwise.

### GetUpgradeSettingsOk

`func (o *NodePoolProperties) GetUpgradeSettingsOk() (*Upgradesettings, bool)`

GetUpgradeSettingsOk returns a tuple with the UpgradeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeSettings

`func (o *NodePoolProperties) SetUpgradeSettings(v Upgradesettings)`

SetUpgradeSettings sets UpgradeSettings field to given value.

### HasUpgradeSettings

`func (o *NodePoolProperties) HasUpgradeSettings() bool`

HasUpgradeSettings returns a boolean if a field has been set.

### GetVmSize

`func (o *NodePoolProperties) GetVmSize() string`

GetVmSize returns the VmSize field if non-nil, zero value otherwise.

### GetVmSizeOk

`func (o *NodePoolProperties) GetVmSizeOk() (*string, bool)`

GetVmSizeOk returns a tuple with the VmSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSize

`func (o *NodePoolProperties) SetVmSize(v string)`

SetVmSize sets VmSize field to given value.

### HasVmSize

`func (o *NodePoolProperties) HasVmSize() bool`

HasVmSize returns a boolean if a field has been set.

### GetVnetSubnetID

`func (o *NodePoolProperties) GetVnetSubnetID() string`

GetVnetSubnetID returns the VnetSubnetID field if non-nil, zero value otherwise.

### GetVnetSubnetIDOk

`func (o *NodePoolProperties) GetVnetSubnetIDOk() (*string, bool)`

GetVnetSubnetIDOk returns a tuple with the VnetSubnetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetSubnetID

`func (o *NodePoolProperties) SetVnetSubnetID(v string)`

SetVnetSubnetID sets VnetSubnetID field to given value.

### HasVnetSubnetID

`func (o *NodePoolProperties) HasVnetSubnetID() bool`

HasVnetSubnetID returns a boolean if a field has been set.

### GetWorkloadRuntime

`func (o *NodePoolProperties) GetWorkloadRuntime() string`

GetWorkloadRuntime returns the WorkloadRuntime field if non-nil, zero value otherwise.

### GetWorkloadRuntimeOk

`func (o *NodePoolProperties) GetWorkloadRuntimeOk() (*string, bool)`

GetWorkloadRuntimeOk returns a tuple with the WorkloadRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadRuntime

`func (o *NodePoolProperties) SetWorkloadRuntime(v string)`

SetWorkloadRuntime sets WorkloadRuntime field to given value.

### HasWorkloadRuntime

`func (o *NodePoolProperties) HasWorkloadRuntime() bool`

HasWorkloadRuntime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


