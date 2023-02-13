# Networkprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsServiceIP** | Pointer to **string** | An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr. | [optional] 
**DockerBridgeCidr** | Pointer to **string** | A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range. | [optional] 
**IpFamilies** | Pointer to **[]string** | IP families are used to determine single-stack or dual-stack clusters. For single-stack, the expected value is IPv4. For dual-stack, the expected values are IPv4 and IPv6. | [optional] 
**LoadBalancerProfile** | Pointer to [**Loadbalancerprofile**](Loadbalancerprofile.md) |  | [optional] 
**LoadBalancerSku** | Pointer to **string** | The default is standard. See Azure Load Balancer SKUs for more information about the differences between load balancer SKUs. Valid values are standard, basic. | [optional] 
**NatGatewayProfile** | Pointer to [**Natgatewayprofile**](Natgatewayprofile.md) |  | [optional] 
**NetworkMode** | Pointer to **string** | This cannot be specified if networkPlugin is anything other than azure. Valid values are transparent, bridge. | [optional] 
**NetworkPlugin** | Pointer to **string** | This cannot be specified if networkPlugin is anything other than azure. Valid values are transparent, bridge. | [optional] 
**NetworkPolicy** | Pointer to **string** | Network policy used for building the Kubernetes network. Valid values are calico, azure. | [optional] 
**OutboundType** | Pointer to **string** | This can only be set at cluster creation time and cannot be changed later. For more information see egress outbound type. Valid values are loadBalancer, userDefinedRouting, managedNATGateway, userAssignedNATGateway. | [optional] 
**PodCidr** | Pointer to **string** | A CIDR notation IP range from which to assign pod IPs when kubenet is used. | [optional] 
**PodCidrs** | Pointer to **[]string** | One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is expected for dual-stack networking. | [optional] 
**ServiceCidr** | Pointer to **string** | A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges. | [optional] 
**ServiceCidrs** | Pointer to **[]string** | One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is expected for dual-stack networking. They must not overlap with any Subnet IP ranges. | [optional] 

## Methods

### NewNetworkprofile

`func NewNetworkprofile() *Networkprofile`

NewNetworkprofile instantiates a new Networkprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkprofileWithDefaults

`func NewNetworkprofileWithDefaults() *Networkprofile`

NewNetworkprofileWithDefaults instantiates a new Networkprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsServiceIP

`func (o *Networkprofile) GetDnsServiceIP() string`

GetDnsServiceIP returns the DnsServiceIP field if non-nil, zero value otherwise.

### GetDnsServiceIPOk

`func (o *Networkprofile) GetDnsServiceIPOk() (*string, bool)`

GetDnsServiceIPOk returns a tuple with the DnsServiceIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServiceIP

`func (o *Networkprofile) SetDnsServiceIP(v string)`

SetDnsServiceIP sets DnsServiceIP field to given value.

### HasDnsServiceIP

`func (o *Networkprofile) HasDnsServiceIP() bool`

HasDnsServiceIP returns a boolean if a field has been set.

### GetDockerBridgeCidr

`func (o *Networkprofile) GetDockerBridgeCidr() string`

GetDockerBridgeCidr returns the DockerBridgeCidr field if non-nil, zero value otherwise.

### GetDockerBridgeCidrOk

`func (o *Networkprofile) GetDockerBridgeCidrOk() (*string, bool)`

GetDockerBridgeCidrOk returns a tuple with the DockerBridgeCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerBridgeCidr

`func (o *Networkprofile) SetDockerBridgeCidr(v string)`

SetDockerBridgeCidr sets DockerBridgeCidr field to given value.

### HasDockerBridgeCidr

`func (o *Networkprofile) HasDockerBridgeCidr() bool`

HasDockerBridgeCidr returns a boolean if a field has been set.

### GetIpFamilies

`func (o *Networkprofile) GetIpFamilies() []string`

GetIpFamilies returns the IpFamilies field if non-nil, zero value otherwise.

### GetIpFamiliesOk

`func (o *Networkprofile) GetIpFamiliesOk() (*[]string, bool)`

GetIpFamiliesOk returns a tuple with the IpFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFamilies

`func (o *Networkprofile) SetIpFamilies(v []string)`

SetIpFamilies sets IpFamilies field to given value.

### HasIpFamilies

`func (o *Networkprofile) HasIpFamilies() bool`

HasIpFamilies returns a boolean if a field has been set.

### GetLoadBalancerProfile

`func (o *Networkprofile) GetLoadBalancerProfile() Loadbalancerprofile`

GetLoadBalancerProfile returns the LoadBalancerProfile field if non-nil, zero value otherwise.

### GetLoadBalancerProfileOk

`func (o *Networkprofile) GetLoadBalancerProfileOk() (*Loadbalancerprofile, bool)`

GetLoadBalancerProfileOk returns a tuple with the LoadBalancerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerProfile

`func (o *Networkprofile) SetLoadBalancerProfile(v Loadbalancerprofile)`

SetLoadBalancerProfile sets LoadBalancerProfile field to given value.

### HasLoadBalancerProfile

`func (o *Networkprofile) HasLoadBalancerProfile() bool`

HasLoadBalancerProfile returns a boolean if a field has been set.

### GetLoadBalancerSku

`func (o *Networkprofile) GetLoadBalancerSku() string`

GetLoadBalancerSku returns the LoadBalancerSku field if non-nil, zero value otherwise.

### GetLoadBalancerSkuOk

`func (o *Networkprofile) GetLoadBalancerSkuOk() (*string, bool)`

GetLoadBalancerSkuOk returns a tuple with the LoadBalancerSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerSku

`func (o *Networkprofile) SetLoadBalancerSku(v string)`

SetLoadBalancerSku sets LoadBalancerSku field to given value.

### HasLoadBalancerSku

`func (o *Networkprofile) HasLoadBalancerSku() bool`

HasLoadBalancerSku returns a boolean if a field has been set.

### GetNatGatewayProfile

`func (o *Networkprofile) GetNatGatewayProfile() Natgatewayprofile`

GetNatGatewayProfile returns the NatGatewayProfile field if non-nil, zero value otherwise.

### GetNatGatewayProfileOk

`func (o *Networkprofile) GetNatGatewayProfileOk() (*Natgatewayprofile, bool)`

GetNatGatewayProfileOk returns a tuple with the NatGatewayProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatGatewayProfile

`func (o *Networkprofile) SetNatGatewayProfile(v Natgatewayprofile)`

SetNatGatewayProfile sets NatGatewayProfile field to given value.

### HasNatGatewayProfile

`func (o *Networkprofile) HasNatGatewayProfile() bool`

HasNatGatewayProfile returns a boolean if a field has been set.

### GetNetworkMode

`func (o *Networkprofile) GetNetworkMode() string`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *Networkprofile) GetNetworkModeOk() (*string, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *Networkprofile) SetNetworkMode(v string)`

SetNetworkMode sets NetworkMode field to given value.

### HasNetworkMode

`func (o *Networkprofile) HasNetworkMode() bool`

HasNetworkMode returns a boolean if a field has been set.

### GetNetworkPlugin

`func (o *Networkprofile) GetNetworkPlugin() string`

GetNetworkPlugin returns the NetworkPlugin field if non-nil, zero value otherwise.

### GetNetworkPluginOk

`func (o *Networkprofile) GetNetworkPluginOk() (*string, bool)`

GetNetworkPluginOk returns a tuple with the NetworkPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPlugin

`func (o *Networkprofile) SetNetworkPlugin(v string)`

SetNetworkPlugin sets NetworkPlugin field to given value.

### HasNetworkPlugin

`func (o *Networkprofile) HasNetworkPlugin() bool`

HasNetworkPlugin returns a boolean if a field has been set.

### GetNetworkPolicy

`func (o *Networkprofile) GetNetworkPolicy() string`

GetNetworkPolicy returns the NetworkPolicy field if non-nil, zero value otherwise.

### GetNetworkPolicyOk

`func (o *Networkprofile) GetNetworkPolicyOk() (*string, bool)`

GetNetworkPolicyOk returns a tuple with the NetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPolicy

`func (o *Networkprofile) SetNetworkPolicy(v string)`

SetNetworkPolicy sets NetworkPolicy field to given value.

### HasNetworkPolicy

`func (o *Networkprofile) HasNetworkPolicy() bool`

HasNetworkPolicy returns a boolean if a field has been set.

### GetOutboundType

`func (o *Networkprofile) GetOutboundType() string`

GetOutboundType returns the OutboundType field if non-nil, zero value otherwise.

### GetOutboundTypeOk

`func (o *Networkprofile) GetOutboundTypeOk() (*string, bool)`

GetOutboundTypeOk returns a tuple with the OutboundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundType

`func (o *Networkprofile) SetOutboundType(v string)`

SetOutboundType sets OutboundType field to given value.

### HasOutboundType

`func (o *Networkprofile) HasOutboundType() bool`

HasOutboundType returns a boolean if a field has been set.

### GetPodCidr

`func (o *Networkprofile) GetPodCidr() string`

GetPodCidr returns the PodCidr field if non-nil, zero value otherwise.

### GetPodCidrOk

`func (o *Networkprofile) GetPodCidrOk() (*string, bool)`

GetPodCidrOk returns a tuple with the PodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidr

`func (o *Networkprofile) SetPodCidr(v string)`

SetPodCidr sets PodCidr field to given value.

### HasPodCidr

`func (o *Networkprofile) HasPodCidr() bool`

HasPodCidr returns a boolean if a field has been set.

### GetPodCidrs

`func (o *Networkprofile) GetPodCidrs() []string`

GetPodCidrs returns the PodCidrs field if non-nil, zero value otherwise.

### GetPodCidrsOk

`func (o *Networkprofile) GetPodCidrsOk() (*[]string, bool)`

GetPodCidrsOk returns a tuple with the PodCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidrs

`func (o *Networkprofile) SetPodCidrs(v []string)`

SetPodCidrs sets PodCidrs field to given value.

### HasPodCidrs

`func (o *Networkprofile) HasPodCidrs() bool`

HasPodCidrs returns a boolean if a field has been set.

### GetServiceCidr

`func (o *Networkprofile) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *Networkprofile) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *Networkprofile) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.

### HasServiceCidr

`func (o *Networkprofile) HasServiceCidr() bool`

HasServiceCidr returns a boolean if a field has been set.

### GetServiceCidrs

`func (o *Networkprofile) GetServiceCidrs() []string`

GetServiceCidrs returns the ServiceCidrs field if non-nil, zero value otherwise.

### GetServiceCidrsOk

`func (o *Networkprofile) GetServiceCidrsOk() (*[]string, bool)`

GetServiceCidrsOk returns a tuple with the ServiceCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidrs

`func (o *Networkprofile) SetServiceCidrs(v []string)`

SetServiceCidrs sets ServiceCidrs field to given value.

### HasServiceCidrs

`func (o *Networkprofile) HasServiceCidrs() bool`

HasServiceCidrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


