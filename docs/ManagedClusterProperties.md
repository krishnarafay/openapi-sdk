# ManagedClusterProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AadProfile** | Pointer to [**Aadprofile**](Aadprofile.md) |  | [optional] 
**AddonProfiles** | Pointer to **map[string]string** | The profile of managed cluster add-on. | [optional] 
**ApiServerAccessProfile** | Pointer to [**Apiserveraccessprofile**](Apiserveraccessprofile.md) |  | [optional] 
**AutoScalerProfile** | Pointer to [**Autoscalerprofile**](Autoscalerprofile.md) |  | [optional] 
**AutoUpgradeProfile** | Pointer to [**Autoupgradeprofile**](Autoupgradeprofile.md) |  | [optional] 
**DisableLocalAccounts** | Pointer to **bool** | If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see disable local accounts. | [optional] 
**DiskEncryptionSetID** | Pointer to **string** | This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{encryptionSetName} | [optional] 
**DnsPrefix** | Pointer to **string** | This cannot be updated once the Managed Cluster has been created. | [optional] 
**EnablePodSecurityPolicy** | Pointer to **bool** | (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy. | [optional] 
**EnableRBAC** | Pointer to **bool** | Whether to enable Kubernetes Role-Based Access Control. | [optional] 
**FqdnSubdomain** | Pointer to **string** | This cannot be updated once the Managed Cluster has been created. | [optional] 
**HttpProxyConfig** | Pointer to [**Httpproxyconfig**](Httpproxyconfig.md) |  | [optional] 
**IdentityProfile** | Pointer to **map[string]string** | Identities associated with the cluster. | [optional] 
**KubernetesVersion** | Pointer to **string** | Both patch version (e.g. 1.20.13) and (e.g. 1.20) are supported. When is specified, the latest supported GA patch version is chosen automatically. Updating the cluster with the same once it has been created (e.g. 1.14.x -&gt; 1.14) will not trigger an upgrade, even if a newer patch version is available. When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -&gt; 1.15.x or 1.15.x -&gt; 1.16.x are allowed, however 1.14.x -&gt; 1.16.x is not allowed. See upgrading an AKS cluster for more details. | [optional] 
**LinuxProfile** | Pointer to [**Linuxprofile**](Linuxprofile.md) |  | [optional] 
**NetworkProfile** | Pointer to [**Networkprofile**](Networkprofile.md) |  | [optional] 
**NodeResourceGroup** | Pointer to **string** | The name of the resource group containing agent pool nodes. | [optional] 
**PodIdentityProfile** | Pointer to [**Podidentityprofile**](Podidentityprofile.md) |  | [optional] 
**PrivateLinkResources** | Pointer to [**[]Privatelinkresources**](Privatelinkresources.md) | Private link resources associated with the cluster. | [optional] 
**PublicNetworkAccess** | Pointer to **string** | Allow or deny public network access for AKS. Valid values are Enabled, Disabled. | [optional] 
**SecurityProfile** | Pointer to [**Securityprofile**](Securityprofile.md) |  | [optional] 
**ServicePrincipalProfile** | Pointer to [**Serviceprincipalprofile**](Serviceprincipalprofile.md) |  | [optional] 
**StorageProfile** | Pointer to [**Storageprofile**](Storageprofile.md) |  | [optional] 
**WindowsProfile** | Pointer to [**Windowsprofile**](Windowsprofile.md) |  | [optional] 

## Methods

### NewManagedClusterProperties

`func NewManagedClusterProperties() *ManagedClusterProperties`

NewManagedClusterProperties instantiates a new ManagedClusterProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedClusterPropertiesWithDefaults

`func NewManagedClusterPropertiesWithDefaults() *ManagedClusterProperties`

NewManagedClusterPropertiesWithDefaults instantiates a new ManagedClusterProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAadProfile

`func (o *ManagedClusterProperties) GetAadProfile() Aadprofile`

GetAadProfile returns the AadProfile field if non-nil, zero value otherwise.

### GetAadProfileOk

`func (o *ManagedClusterProperties) GetAadProfileOk() (*Aadprofile, bool)`

GetAadProfileOk returns a tuple with the AadProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAadProfile

`func (o *ManagedClusterProperties) SetAadProfile(v Aadprofile)`

SetAadProfile sets AadProfile field to given value.

### HasAadProfile

`func (o *ManagedClusterProperties) HasAadProfile() bool`

HasAadProfile returns a boolean if a field has been set.

### GetAddonProfiles

`func (o *ManagedClusterProperties) GetAddonProfiles() map[string]string`

GetAddonProfiles returns the AddonProfiles field if non-nil, zero value otherwise.

### GetAddonProfilesOk

`func (o *ManagedClusterProperties) GetAddonProfilesOk() (*map[string]string, bool)`

GetAddonProfilesOk returns a tuple with the AddonProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonProfiles

`func (o *ManagedClusterProperties) SetAddonProfiles(v map[string]string)`

SetAddonProfiles sets AddonProfiles field to given value.

### HasAddonProfiles

`func (o *ManagedClusterProperties) HasAddonProfiles() bool`

HasAddonProfiles returns a boolean if a field has been set.

### GetApiServerAccessProfile

`func (o *ManagedClusterProperties) GetApiServerAccessProfile() Apiserveraccessprofile`

GetApiServerAccessProfile returns the ApiServerAccessProfile field if non-nil, zero value otherwise.

### GetApiServerAccessProfileOk

`func (o *ManagedClusterProperties) GetApiServerAccessProfileOk() (*Apiserveraccessprofile, bool)`

GetApiServerAccessProfileOk returns a tuple with the ApiServerAccessProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerAccessProfile

`func (o *ManagedClusterProperties) SetApiServerAccessProfile(v Apiserveraccessprofile)`

SetApiServerAccessProfile sets ApiServerAccessProfile field to given value.

### HasApiServerAccessProfile

`func (o *ManagedClusterProperties) HasApiServerAccessProfile() bool`

HasApiServerAccessProfile returns a boolean if a field has been set.

### GetAutoScalerProfile

`func (o *ManagedClusterProperties) GetAutoScalerProfile() Autoscalerprofile`

GetAutoScalerProfile returns the AutoScalerProfile field if non-nil, zero value otherwise.

### GetAutoScalerProfileOk

`func (o *ManagedClusterProperties) GetAutoScalerProfileOk() (*Autoscalerprofile, bool)`

GetAutoScalerProfileOk returns a tuple with the AutoScalerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalerProfile

`func (o *ManagedClusterProperties) SetAutoScalerProfile(v Autoscalerprofile)`

SetAutoScalerProfile sets AutoScalerProfile field to given value.

### HasAutoScalerProfile

`func (o *ManagedClusterProperties) HasAutoScalerProfile() bool`

HasAutoScalerProfile returns a boolean if a field has been set.

### GetAutoUpgradeProfile

`func (o *ManagedClusterProperties) GetAutoUpgradeProfile() Autoupgradeprofile`

GetAutoUpgradeProfile returns the AutoUpgradeProfile field if non-nil, zero value otherwise.

### GetAutoUpgradeProfileOk

`func (o *ManagedClusterProperties) GetAutoUpgradeProfileOk() (*Autoupgradeprofile, bool)`

GetAutoUpgradeProfileOk returns a tuple with the AutoUpgradeProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpgradeProfile

`func (o *ManagedClusterProperties) SetAutoUpgradeProfile(v Autoupgradeprofile)`

SetAutoUpgradeProfile sets AutoUpgradeProfile field to given value.

### HasAutoUpgradeProfile

`func (o *ManagedClusterProperties) HasAutoUpgradeProfile() bool`

HasAutoUpgradeProfile returns a boolean if a field has been set.

### GetDisableLocalAccounts

`func (o *ManagedClusterProperties) GetDisableLocalAccounts() bool`

GetDisableLocalAccounts returns the DisableLocalAccounts field if non-nil, zero value otherwise.

### GetDisableLocalAccountsOk

`func (o *ManagedClusterProperties) GetDisableLocalAccountsOk() (*bool, bool)`

GetDisableLocalAccountsOk returns a tuple with the DisableLocalAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableLocalAccounts

`func (o *ManagedClusterProperties) SetDisableLocalAccounts(v bool)`

SetDisableLocalAccounts sets DisableLocalAccounts field to given value.

### HasDisableLocalAccounts

`func (o *ManagedClusterProperties) HasDisableLocalAccounts() bool`

HasDisableLocalAccounts returns a boolean if a field has been set.

### GetDiskEncryptionSetID

`func (o *ManagedClusterProperties) GetDiskEncryptionSetID() string`

GetDiskEncryptionSetID returns the DiskEncryptionSetID field if non-nil, zero value otherwise.

### GetDiskEncryptionSetIDOk

`func (o *ManagedClusterProperties) GetDiskEncryptionSetIDOk() (*string, bool)`

GetDiskEncryptionSetIDOk returns a tuple with the DiskEncryptionSetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryptionSetID

`func (o *ManagedClusterProperties) SetDiskEncryptionSetID(v string)`

SetDiskEncryptionSetID sets DiskEncryptionSetID field to given value.

### HasDiskEncryptionSetID

`func (o *ManagedClusterProperties) HasDiskEncryptionSetID() bool`

HasDiskEncryptionSetID returns a boolean if a field has been set.

### GetDnsPrefix

`func (o *ManagedClusterProperties) GetDnsPrefix() string`

GetDnsPrefix returns the DnsPrefix field if non-nil, zero value otherwise.

### GetDnsPrefixOk

`func (o *ManagedClusterProperties) GetDnsPrefixOk() (*string, bool)`

GetDnsPrefixOk returns a tuple with the DnsPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrefix

`func (o *ManagedClusterProperties) SetDnsPrefix(v string)`

SetDnsPrefix sets DnsPrefix field to given value.

### HasDnsPrefix

`func (o *ManagedClusterProperties) HasDnsPrefix() bool`

HasDnsPrefix returns a boolean if a field has been set.

### GetEnablePodSecurityPolicy

`func (o *ManagedClusterProperties) GetEnablePodSecurityPolicy() bool`

GetEnablePodSecurityPolicy returns the EnablePodSecurityPolicy field if non-nil, zero value otherwise.

### GetEnablePodSecurityPolicyOk

`func (o *ManagedClusterProperties) GetEnablePodSecurityPolicyOk() (*bool, bool)`

GetEnablePodSecurityPolicyOk returns a tuple with the EnablePodSecurityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePodSecurityPolicy

`func (o *ManagedClusterProperties) SetEnablePodSecurityPolicy(v bool)`

SetEnablePodSecurityPolicy sets EnablePodSecurityPolicy field to given value.

### HasEnablePodSecurityPolicy

`func (o *ManagedClusterProperties) HasEnablePodSecurityPolicy() bool`

HasEnablePodSecurityPolicy returns a boolean if a field has been set.

### GetEnableRBAC

`func (o *ManagedClusterProperties) GetEnableRBAC() bool`

GetEnableRBAC returns the EnableRBAC field if non-nil, zero value otherwise.

### GetEnableRBACOk

`func (o *ManagedClusterProperties) GetEnableRBACOk() (*bool, bool)`

GetEnableRBACOk returns a tuple with the EnableRBAC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRBAC

`func (o *ManagedClusterProperties) SetEnableRBAC(v bool)`

SetEnableRBAC sets EnableRBAC field to given value.

### HasEnableRBAC

`func (o *ManagedClusterProperties) HasEnableRBAC() bool`

HasEnableRBAC returns a boolean if a field has been set.

### GetFqdnSubdomain

`func (o *ManagedClusterProperties) GetFqdnSubdomain() string`

GetFqdnSubdomain returns the FqdnSubdomain field if non-nil, zero value otherwise.

### GetFqdnSubdomainOk

`func (o *ManagedClusterProperties) GetFqdnSubdomainOk() (*string, bool)`

GetFqdnSubdomainOk returns a tuple with the FqdnSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdnSubdomain

`func (o *ManagedClusterProperties) SetFqdnSubdomain(v string)`

SetFqdnSubdomain sets FqdnSubdomain field to given value.

### HasFqdnSubdomain

`func (o *ManagedClusterProperties) HasFqdnSubdomain() bool`

HasFqdnSubdomain returns a boolean if a field has been set.

### GetHttpProxyConfig

`func (o *ManagedClusterProperties) GetHttpProxyConfig() Httpproxyconfig`

GetHttpProxyConfig returns the HttpProxyConfig field if non-nil, zero value otherwise.

### GetHttpProxyConfigOk

`func (o *ManagedClusterProperties) GetHttpProxyConfigOk() (*Httpproxyconfig, bool)`

GetHttpProxyConfigOk returns a tuple with the HttpProxyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyConfig

`func (o *ManagedClusterProperties) SetHttpProxyConfig(v Httpproxyconfig)`

SetHttpProxyConfig sets HttpProxyConfig field to given value.

### HasHttpProxyConfig

`func (o *ManagedClusterProperties) HasHttpProxyConfig() bool`

HasHttpProxyConfig returns a boolean if a field has been set.

### GetIdentityProfile

`func (o *ManagedClusterProperties) GetIdentityProfile() map[string]string`

GetIdentityProfile returns the IdentityProfile field if non-nil, zero value otherwise.

### GetIdentityProfileOk

`func (o *ManagedClusterProperties) GetIdentityProfileOk() (*map[string]string, bool)`

GetIdentityProfileOk returns a tuple with the IdentityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProfile

`func (o *ManagedClusterProperties) SetIdentityProfile(v map[string]string)`

SetIdentityProfile sets IdentityProfile field to given value.

### HasIdentityProfile

`func (o *ManagedClusterProperties) HasIdentityProfile() bool`

HasIdentityProfile returns a boolean if a field has been set.

### GetKubernetesVersion

`func (o *ManagedClusterProperties) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *ManagedClusterProperties) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *ManagedClusterProperties) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.

### HasKubernetesVersion

`func (o *ManagedClusterProperties) HasKubernetesVersion() bool`

HasKubernetesVersion returns a boolean if a field has been set.

### GetLinuxProfile

`func (o *ManagedClusterProperties) GetLinuxProfile() Linuxprofile`

GetLinuxProfile returns the LinuxProfile field if non-nil, zero value otherwise.

### GetLinuxProfileOk

`func (o *ManagedClusterProperties) GetLinuxProfileOk() (*Linuxprofile, bool)`

GetLinuxProfileOk returns a tuple with the LinuxProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxProfile

`func (o *ManagedClusterProperties) SetLinuxProfile(v Linuxprofile)`

SetLinuxProfile sets LinuxProfile field to given value.

### HasLinuxProfile

`func (o *ManagedClusterProperties) HasLinuxProfile() bool`

HasLinuxProfile returns a boolean if a field has been set.

### GetNetworkProfile

`func (o *ManagedClusterProperties) GetNetworkProfile() Networkprofile`

GetNetworkProfile returns the NetworkProfile field if non-nil, zero value otherwise.

### GetNetworkProfileOk

`func (o *ManagedClusterProperties) GetNetworkProfileOk() (*Networkprofile, bool)`

GetNetworkProfileOk returns a tuple with the NetworkProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProfile

`func (o *ManagedClusterProperties) SetNetworkProfile(v Networkprofile)`

SetNetworkProfile sets NetworkProfile field to given value.

### HasNetworkProfile

`func (o *ManagedClusterProperties) HasNetworkProfile() bool`

HasNetworkProfile returns a boolean if a field has been set.

### GetNodeResourceGroup

`func (o *ManagedClusterProperties) GetNodeResourceGroup() string`

GetNodeResourceGroup returns the NodeResourceGroup field if non-nil, zero value otherwise.

### GetNodeResourceGroupOk

`func (o *ManagedClusterProperties) GetNodeResourceGroupOk() (*string, bool)`

GetNodeResourceGroupOk returns a tuple with the NodeResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeResourceGroup

`func (o *ManagedClusterProperties) SetNodeResourceGroup(v string)`

SetNodeResourceGroup sets NodeResourceGroup field to given value.

### HasNodeResourceGroup

`func (o *ManagedClusterProperties) HasNodeResourceGroup() bool`

HasNodeResourceGroup returns a boolean if a field has been set.

### GetPodIdentityProfile

`func (o *ManagedClusterProperties) GetPodIdentityProfile() Podidentityprofile`

GetPodIdentityProfile returns the PodIdentityProfile field if non-nil, zero value otherwise.

### GetPodIdentityProfileOk

`func (o *ManagedClusterProperties) GetPodIdentityProfileOk() (*Podidentityprofile, bool)`

GetPodIdentityProfileOk returns a tuple with the PodIdentityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodIdentityProfile

`func (o *ManagedClusterProperties) SetPodIdentityProfile(v Podidentityprofile)`

SetPodIdentityProfile sets PodIdentityProfile field to given value.

### HasPodIdentityProfile

`func (o *ManagedClusterProperties) HasPodIdentityProfile() bool`

HasPodIdentityProfile returns a boolean if a field has been set.

### GetPrivateLinkResources

`func (o *ManagedClusterProperties) GetPrivateLinkResources() []Privatelinkresources`

GetPrivateLinkResources returns the PrivateLinkResources field if non-nil, zero value otherwise.

### GetPrivateLinkResourcesOk

`func (o *ManagedClusterProperties) GetPrivateLinkResourcesOk() (*[]Privatelinkresources, bool)`

GetPrivateLinkResourcesOk returns a tuple with the PrivateLinkResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkResources

`func (o *ManagedClusterProperties) SetPrivateLinkResources(v []Privatelinkresources)`

SetPrivateLinkResources sets PrivateLinkResources field to given value.

### HasPrivateLinkResources

`func (o *ManagedClusterProperties) HasPrivateLinkResources() bool`

HasPrivateLinkResources returns a boolean if a field has been set.

### GetPublicNetworkAccess

`func (o *ManagedClusterProperties) GetPublicNetworkAccess() string`

GetPublicNetworkAccess returns the PublicNetworkAccess field if non-nil, zero value otherwise.

### GetPublicNetworkAccessOk

`func (o *ManagedClusterProperties) GetPublicNetworkAccessOk() (*string, bool)`

GetPublicNetworkAccessOk returns a tuple with the PublicNetworkAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkAccess

`func (o *ManagedClusterProperties) SetPublicNetworkAccess(v string)`

SetPublicNetworkAccess sets PublicNetworkAccess field to given value.

### HasPublicNetworkAccess

`func (o *ManagedClusterProperties) HasPublicNetworkAccess() bool`

HasPublicNetworkAccess returns a boolean if a field has been set.

### GetSecurityProfile

`func (o *ManagedClusterProperties) GetSecurityProfile() Securityprofile`

GetSecurityProfile returns the SecurityProfile field if non-nil, zero value otherwise.

### GetSecurityProfileOk

`func (o *ManagedClusterProperties) GetSecurityProfileOk() (*Securityprofile, bool)`

GetSecurityProfileOk returns a tuple with the SecurityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProfile

`func (o *ManagedClusterProperties) SetSecurityProfile(v Securityprofile)`

SetSecurityProfile sets SecurityProfile field to given value.

### HasSecurityProfile

`func (o *ManagedClusterProperties) HasSecurityProfile() bool`

HasSecurityProfile returns a boolean if a field has been set.

### GetServicePrincipalProfile

`func (o *ManagedClusterProperties) GetServicePrincipalProfile() Serviceprincipalprofile`

GetServicePrincipalProfile returns the ServicePrincipalProfile field if non-nil, zero value otherwise.

### GetServicePrincipalProfileOk

`func (o *ManagedClusterProperties) GetServicePrincipalProfileOk() (*Serviceprincipalprofile, bool)`

GetServicePrincipalProfileOk returns a tuple with the ServicePrincipalProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalProfile

`func (o *ManagedClusterProperties) SetServicePrincipalProfile(v Serviceprincipalprofile)`

SetServicePrincipalProfile sets ServicePrincipalProfile field to given value.

### HasServicePrincipalProfile

`func (o *ManagedClusterProperties) HasServicePrincipalProfile() bool`

HasServicePrincipalProfile returns a boolean if a field has been set.

### GetStorageProfile

`func (o *ManagedClusterProperties) GetStorageProfile() Storageprofile`

GetStorageProfile returns the StorageProfile field if non-nil, zero value otherwise.

### GetStorageProfileOk

`func (o *ManagedClusterProperties) GetStorageProfileOk() (*Storageprofile, bool)`

GetStorageProfileOk returns a tuple with the StorageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfile

`func (o *ManagedClusterProperties) SetStorageProfile(v Storageprofile)`

SetStorageProfile sets StorageProfile field to given value.

### HasStorageProfile

`func (o *ManagedClusterProperties) HasStorageProfile() bool`

HasStorageProfile returns a boolean if a field has been set.

### GetWindowsProfile

`func (o *ManagedClusterProperties) GetWindowsProfile() Windowsprofile`

GetWindowsProfile returns the WindowsProfile field if non-nil, zero value otherwise.

### GetWindowsProfileOk

`func (o *ManagedClusterProperties) GetWindowsProfileOk() (*Windowsprofile, bool)`

GetWindowsProfileOk returns a tuple with the WindowsProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsProfile

`func (o *ManagedClusterProperties) SetWindowsProfile(v Windowsprofile)`

SetWindowsProfile sets WindowsProfile field to given value.

### HasWindowsProfile

`func (o *ManagedClusterProperties) HasWindowsProfile() bool`

HasWindowsProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


