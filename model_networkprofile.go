/*
Kubernetes Operations APIs

Kubernetes Operations APIs

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rafaysdkgov3

import (
	"encoding/json"
)

// checks if the Networkprofile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Networkprofile{}

// Networkprofile struct for Networkprofile
type Networkprofile struct {
	// An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.
	DnsServiceIP *string `json:"dnsServiceIP,omitempty"`
	// A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.
	DockerBridgeCidr *string `json:"dockerBridgeCidr,omitempty"`
	// IP families are used to determine single-stack or dual-stack clusters. For single-stack, the expected value is IPv4. For dual-stack, the expected values are IPv4 and IPv6.
	IpFamilies []string `json:"ipFamilies,omitempty"`
	LoadBalancerProfile *Loadbalancerprofile `json:"loadBalancerProfile,omitempty"`
	// The default is standard. See Azure Load Balancer SKUs for more information about the differences between load balancer SKUs. Valid values are standard, basic.
	LoadBalancerSku *string `json:"loadBalancerSku,omitempty"`
	NatGatewayProfile *Natgatewayprofile `json:"natGatewayProfile,omitempty"`
	// This cannot be specified if networkPlugin is anything other than azure. Valid values are transparent, bridge.
	NetworkMode *string `json:"networkMode,omitempty"`
	// This cannot be specified if networkPlugin is anything other than azure. Valid values are transparent, bridge.
	NetworkPlugin *string `json:"networkPlugin,omitempty"`
	// Network policy used for building the Kubernetes network. Valid values are calico, azure.
	NetworkPolicy *string `json:"networkPolicy,omitempty"`
	// This can only be set at cluster creation time and cannot be changed later. For more information see egress outbound type. Valid values are loadBalancer, userDefinedRouting, managedNATGateway, userAssignedNATGateway.
	OutboundType *string `json:"outboundType,omitempty"`
	// A CIDR notation IP range from which to assign pod IPs when kubenet is used.
	PodCidr *string `json:"podCidr,omitempty"`
	// One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is expected for dual-stack networking.
	PodCidrs []string `json:"podCidrs,omitempty"`
	// A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.
	ServiceCidr *string `json:"serviceCidr,omitempty"`
	// One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is expected for dual-stack networking. They must not overlap with any Subnet IP ranges.
	ServiceCidrs []string `json:"serviceCidrs,omitempty"`
}

// NewNetworkprofile instantiates a new Networkprofile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkprofile() *Networkprofile {
	this := Networkprofile{}
	return &this
}

// NewNetworkprofileWithDefaults instantiates a new Networkprofile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkprofileWithDefaults() *Networkprofile {
	this := Networkprofile{}
	return &this
}

// GetDnsServiceIP returns the DnsServiceIP field value if set, zero value otherwise.
func (o *Networkprofile) GetDnsServiceIP() string {
	if o == nil || isNil(o.DnsServiceIP) {
		var ret string
		return ret
	}
	return *o.DnsServiceIP
}

// GetDnsServiceIPOk returns a tuple with the DnsServiceIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Networkprofile) GetDnsServiceIPOk() (*string, bool) {
	if o == nil || isNil(o.DnsServiceIP) {
		return nil, false
	}
	return o.DnsServiceIP, true
}

// HasDnsServiceIP returns a boolean if a field has been set.
func (o *Networkprofile) HasDnsServiceIP() bool {
	if o != nil && !isNil(o.DnsServiceIP) {
		return true
	}

	return false
}

// SetDnsServiceIP gets a reference to the given string and assigns it to the DnsServiceIP field.
func (o *Networkprofile) SetDnsServiceIP(v string) {
	o.DnsServiceIP = &v
}

// GetDockerBridgeCidr returns the DockerBridgeCidr field value if set, zero value otherwise.
func (o *Networkprofile) GetDockerBridgeCidr() string {
	if o == nil || isNil(o.DockerBridgeCidr) {
		var ret string
		return ret
	}
	return *o.DockerBridgeCidr
}

// GetDockerBridgeCidrOk returns a tuple with the DockerBridgeCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Networkprofile) GetDockerBridgeCidrOk() (*string, bool) {
	if o == nil || isNil(o.DockerBridgeCidr) {
		return nil, false
	}
	return o.DockerBridgeCidr, true
}

// HasDockerBridgeCidr returns a boolean if a field has been set.
func (o *Networkprofile) HasDockerBridgeCidr() bool {
	if o != nil && !isNil(o.DockerBridgeCidr) {
		return true
	}

	return false
}

// SetDockerBridgeCidr gets a reference to the given string and assigns it to the DockerBridgeCidr field.
func (o *Networkprofile) SetDockerBridgeCidr(v string) {
	o.DockerBridgeCidr = &v
}

// GetIpFamilies returns the IpFamilies field value if set, zero value otherwise.
func (o *Networkprofile) GetIpFamilies() []string {
	if o == nil || isNil(o.IpFamilies) {
		var ret []string
		return ret
	}
	return o.IpFamilies
}

// GetIpFamiliesOk returns a tuple with the IpFamilies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Networkprofile) GetIpFamiliesOk() ([]string, bool) {
	if o == nil || isNil(o.IpFamilies) {
		return nil, false
	}
	return o.IpFamilies, true
}

// HasIpFamilies returns a boolean if a field has been set.
func (o *Networkprofile) HasIpFamilies() bool {
	if o != nil && !isNil(o.IpFamilies) {
		return true
	}

	return false
}

// SetIpFamilies gets a reference to the given []string and assigns it to the IpFamilies field.
func (o *Networkprofile) SetIpFamilies(v []string) {
	o.IpFamilies = v
}

// GetLoadBalancerProfile returns the LoadBalancerProfile field value if set, zero value otherwise.
func (o *Networkprofile) GetLoadBalancerProfile() Loadbalancerprofile {
	if o == nil || isNil(o.LoadBalancerProfile) {
		var ret Loadbalancerprofile
		return ret
	}
	return *o.LoadBalancerProfile
}

// GetLoadBalancerProfileOk returns a tuple with the LoadBalancerProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Networkprofile) GetLoadBalancerProfileOk() (*Loadbalancerprofile, bool) {
	if o == nil || isNil(o.LoadBalancerProfile) {
		return nil, false
	}
	return o.LoadBalancerProfile, true
}

// HasLoadBalancerProfile returns a boolean if a field has been set.
func (o *Networkprofile) HasLoadBalancerProfile() bool {
	if o != nil && !isNil(o.LoadBalancerProfile) {
		return true
	}

	return false
}

// SetLoadBalancerProfile gets a reference to the given Loadbalancerprofile and assigns it to the LoadBalancerProfile field.
func (o *Networkprofile) SetLoadBalancerProfile(v Loadbalancerprofile) {
	o.LoadBalancerProfile = &v
}

// GetLoadBalancerSku returns the LoadBalancerSku field value if set, zero value otherwise.
func (o *Networkprofile) GetLoadBalancerSku() string {
	if o == nil || isNil(o.LoadBalancerSku) {
		var ret string
		return ret
	}
	return *o.LoadBalancerSku
}

// GetLoadBalancerSkuOk returns a tuple with the LoadBalancerSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Networkprofile) GetLoadBalancerSkuOk() (*string, bool) {
	if o == nil || isNil(o.LoadBalancerSku) {
		return nil, false
	}
	return o.LoadBalancerSku, true
}

// HasLoadBalancerSku returns a boolean if a field has been set.
func (o *Networkprofile) HasLoadBalancerSku() bool {
	if o != nil && !isNil(o.LoadBalancerSku) {
		return true
	}

	return false
}

// SetLoadBalancerSku gets a reference to the given string and assigns it to the LoadBalancerSku field.
func (o *Networkprofile) SetLoadBalancerSku(v string) {
	o.LoadBalancerSku = &v
}

// GetNatGatewayProfile returns the NatGatewayProfile field value if set, zero value otherwise.
func (o *Networkprofile) GetNatGatewayProfile() Natgatewayprofile {
	if o == nil || isNil(o.NatGatewayProfile) {
		var ret Natgatewayprofile
		return ret
	}
	return *o.NatGatewayProfile
}

// GetNatGatewayProfileOk returns a tuple with the NatGatewayProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Networkprofile) GetNatGatewayProfileOk() (*Natgatewayprofile, bool) {
	if o == nil || isNil(o.NatGatewayProfile) {
		return nil, false
	}
	return o.NatGatewayProfile, true
}

// HasNatGatewayProfile returns a boolean if a field has been set.
func (o *Networkprofile) HasNatGatewayProfile() bool {
	if o != nil && !isNil(o.NatGatewayProfile) {
		return true
	}

	return false
}

// SetNatGatewayProfile gets a reference to the given Natgatewayprofile and assigns it to the NatGatewayProfile field.
func (o *Networkprofile) SetNatGatewayProfile(v Natgatewayprofile) {
	o.NatGatewayProfile = &v
}

// GetNetworkMode returns the NetworkMode field value if set, zero value otherwise.
func (o *Networkprofile) GetNetworkMode() string {
	if o == nil || isNil(o.NetworkMode) {
		var ret string
		return ret
	}
	return *o.NetworkMode
}

// GetNetworkModeOk returns a tuple with the NetworkMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Networkprofile) GetNetworkModeOk() (*string, bool) {
	if o == nil || isNil(o.NetworkMode) {
		return nil, false
	}
	return o.NetworkMode, true
}

// HasNetworkMode returns a boolean if a field has been set.
func (o *Networkprofile) HasNetworkMode() bool {
	if o != nil && !isNil(o.NetworkMode) {
		return true
	}

	return false
}

// SetNetworkMode gets a reference to the given string and assigns it to the NetworkMode field.
func (o *Networkprofile) SetNetworkMode(v string) {
	o.NetworkMode = &v
}

// GetNetworkPlugin returns the NetworkPlugin field value if set, zero value otherwise.
func (o *Networkprofile) GetNetworkPlugin() string {
	if o == nil || isNil(o.NetworkPlugin) {
		var ret string
		return ret
	}
	return *o.NetworkPlugin
}

// GetNetworkPluginOk returns a tuple with the NetworkPlugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Networkprofile) GetNetworkPluginOk() (*string, bool) {
	if o == nil || isNil(o.NetworkPlugin) {
		return nil, false
	}
	return o.NetworkPlugin, true
}

// HasNetworkPlugin returns a boolean if a field has been set.
func (o *Networkprofile) HasNetworkPlugin() bool {
	if o != nil && !isNil(o.NetworkPlugin) {
		return true
	}

	return false
}

// SetNetworkPlugin gets a reference to the given string and assigns it to the NetworkPlugin field.
func (o *Networkprofile) SetNetworkPlugin(v string) {
	o.NetworkPlugin = &v
}

// GetNetworkPolicy returns the NetworkPolicy field value if set, zero value otherwise.
func (o *Networkprofile) GetNetworkPolicy() string {
	if o == nil || isNil(o.NetworkPolicy) {
		var ret string
		return ret
	}
	return *o.NetworkPolicy
}

// GetNetworkPolicyOk returns a tuple with the NetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Networkprofile) GetNetworkPolicyOk() (*string, bool) {
	if o == nil || isNil(o.NetworkPolicy) {
		return nil, false
	}
	return o.NetworkPolicy, true
}

// HasNetworkPolicy returns a boolean if a field has been set.
func (o *Networkprofile) HasNetworkPolicy() bool {
	if o != nil && !isNil(o.NetworkPolicy) {
		return true
	}

	return false
}

// SetNetworkPolicy gets a reference to the given string and assigns it to the NetworkPolicy field.
func (o *Networkprofile) SetNetworkPolicy(v string) {
	o.NetworkPolicy = &v
}

// GetOutboundType returns the OutboundType field value if set, zero value otherwise.
func (o *Networkprofile) GetOutboundType() string {
	if o == nil || isNil(o.OutboundType) {
		var ret string
		return ret
	}
	return *o.OutboundType
}

// GetOutboundTypeOk returns a tuple with the OutboundType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Networkprofile) GetOutboundTypeOk() (*string, bool) {
	if o == nil || isNil(o.OutboundType) {
		return nil, false
	}
	return o.OutboundType, true
}

// HasOutboundType returns a boolean if a field has been set.
func (o *Networkprofile) HasOutboundType() bool {
	if o != nil && !isNil(o.OutboundType) {
		return true
	}

	return false
}

// SetOutboundType gets a reference to the given string and assigns it to the OutboundType field.
func (o *Networkprofile) SetOutboundType(v string) {
	o.OutboundType = &v
}

// GetPodCidr returns the PodCidr field value if set, zero value otherwise.
func (o *Networkprofile) GetPodCidr() string {
	if o == nil || isNil(o.PodCidr) {
		var ret string
		return ret
	}
	return *o.PodCidr
}

// GetPodCidrOk returns a tuple with the PodCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Networkprofile) GetPodCidrOk() (*string, bool) {
	if o == nil || isNil(o.PodCidr) {
		return nil, false
	}
	return o.PodCidr, true
}

// HasPodCidr returns a boolean if a field has been set.
func (o *Networkprofile) HasPodCidr() bool {
	if o != nil && !isNil(o.PodCidr) {
		return true
	}

	return false
}

// SetPodCidr gets a reference to the given string and assigns it to the PodCidr field.
func (o *Networkprofile) SetPodCidr(v string) {
	o.PodCidr = &v
}

// GetPodCidrs returns the PodCidrs field value if set, zero value otherwise.
func (o *Networkprofile) GetPodCidrs() []string {
	if o == nil || isNil(o.PodCidrs) {
		var ret []string
		return ret
	}
	return o.PodCidrs
}

// GetPodCidrsOk returns a tuple with the PodCidrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Networkprofile) GetPodCidrsOk() ([]string, bool) {
	if o == nil || isNil(o.PodCidrs) {
		return nil, false
	}
	return o.PodCidrs, true
}

// HasPodCidrs returns a boolean if a field has been set.
func (o *Networkprofile) HasPodCidrs() bool {
	if o != nil && !isNil(o.PodCidrs) {
		return true
	}

	return false
}

// SetPodCidrs gets a reference to the given []string and assigns it to the PodCidrs field.
func (o *Networkprofile) SetPodCidrs(v []string) {
	o.PodCidrs = v
}

// GetServiceCidr returns the ServiceCidr field value if set, zero value otherwise.
func (o *Networkprofile) GetServiceCidr() string {
	if o == nil || isNil(o.ServiceCidr) {
		var ret string
		return ret
	}
	return *o.ServiceCidr
}

// GetServiceCidrOk returns a tuple with the ServiceCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Networkprofile) GetServiceCidrOk() (*string, bool) {
	if o == nil || isNil(o.ServiceCidr) {
		return nil, false
	}
	return o.ServiceCidr, true
}

// HasServiceCidr returns a boolean if a field has been set.
func (o *Networkprofile) HasServiceCidr() bool {
	if o != nil && !isNil(o.ServiceCidr) {
		return true
	}

	return false
}

// SetServiceCidr gets a reference to the given string and assigns it to the ServiceCidr field.
func (o *Networkprofile) SetServiceCidr(v string) {
	o.ServiceCidr = &v
}

// GetServiceCidrs returns the ServiceCidrs field value if set, zero value otherwise.
func (o *Networkprofile) GetServiceCidrs() []string {
	if o == nil || isNil(o.ServiceCidrs) {
		var ret []string
		return ret
	}
	return o.ServiceCidrs
}

// GetServiceCidrsOk returns a tuple with the ServiceCidrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Networkprofile) GetServiceCidrsOk() ([]string, bool) {
	if o == nil || isNil(o.ServiceCidrs) {
		return nil, false
	}
	return o.ServiceCidrs, true
}

// HasServiceCidrs returns a boolean if a field has been set.
func (o *Networkprofile) HasServiceCidrs() bool {
	if o != nil && !isNil(o.ServiceCidrs) {
		return true
	}

	return false
}

// SetServiceCidrs gets a reference to the given []string and assigns it to the ServiceCidrs field.
func (o *Networkprofile) SetServiceCidrs(v []string) {
	o.ServiceCidrs = v
}

func (o Networkprofile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Networkprofile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DnsServiceIP) {
		toSerialize["dnsServiceIP"] = o.DnsServiceIP
	}
	if !isNil(o.DockerBridgeCidr) {
		toSerialize["dockerBridgeCidr"] = o.DockerBridgeCidr
	}
	if !isNil(o.IpFamilies) {
		toSerialize["ipFamilies"] = o.IpFamilies
	}
	if !isNil(o.LoadBalancerProfile) {
		toSerialize["loadBalancerProfile"] = o.LoadBalancerProfile
	}
	if !isNil(o.LoadBalancerSku) {
		toSerialize["loadBalancerSku"] = o.LoadBalancerSku
	}
	if !isNil(o.NatGatewayProfile) {
		toSerialize["natGatewayProfile"] = o.NatGatewayProfile
	}
	if !isNil(o.NetworkMode) {
		toSerialize["networkMode"] = o.NetworkMode
	}
	if !isNil(o.NetworkPlugin) {
		toSerialize["networkPlugin"] = o.NetworkPlugin
	}
	if !isNil(o.NetworkPolicy) {
		toSerialize["networkPolicy"] = o.NetworkPolicy
	}
	if !isNil(o.OutboundType) {
		toSerialize["outboundType"] = o.OutboundType
	}
	if !isNil(o.PodCidr) {
		toSerialize["podCidr"] = o.PodCidr
	}
	if !isNil(o.PodCidrs) {
		toSerialize["podCidrs"] = o.PodCidrs
	}
	if !isNil(o.ServiceCidr) {
		toSerialize["serviceCidr"] = o.ServiceCidr
	}
	if !isNil(o.ServiceCidrs) {
		toSerialize["serviceCidrs"] = o.ServiceCidrs
	}
	return toSerialize, nil
}

type NullableNetworkprofile struct {
	value *Networkprofile
	isSet bool
}

func (v NullableNetworkprofile) Get() *Networkprofile {
	return v.value
}

func (v *NullableNetworkprofile) Set(val *Networkprofile) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkprofile) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkprofile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkprofile(val *Networkprofile) *NullableNetworkprofile {
	return &NullableNetworkprofile{value: val, isSet: true}
}

func (v NullableNetworkprofile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkprofile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


