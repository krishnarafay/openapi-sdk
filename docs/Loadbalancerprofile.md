# Loadbalancerprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocatedOutboundPorts** | Pointer to **int32** | The desired number of allocated SNAT ports per VM. Allowed values are in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports. | [optional] 
**EffectiveOutboundIPs** | Pointer to [**[]Effectiveoutboundips**](Effectiveoutboundips.md) | The effective outbound IP resources of the cluster load balancer. | [optional] 
**EnableMultipleStandardLoadBalancers** | Pointer to **bool** | Enable multiple standard load balancers per AKS cluster or not. | [optional] 
**IdleTimeoutInMinutes** | Pointer to **int32** | Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120 (inclusive). The default value is 30 minutes. | [optional] 
**ManagedOutboundIPs** | Pointer to [**Managedoutboundips**](Managedoutboundips.md) |  | [optional] 
**OutboundIPPrefixes** | Pointer to [**Outboundipprefixes**](Outboundipprefixes.md) |  | [optional] 
**OutboundIPs** | Pointer to [**Outboundips**](Outboundips.md) |  | [optional] 

## Methods

### NewLoadbalancerprofile

`func NewLoadbalancerprofile() *Loadbalancerprofile`

NewLoadbalancerprofile instantiates a new Loadbalancerprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadbalancerprofileWithDefaults

`func NewLoadbalancerprofileWithDefaults() *Loadbalancerprofile`

NewLoadbalancerprofileWithDefaults instantiates a new Loadbalancerprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatedOutboundPorts

`func (o *Loadbalancerprofile) GetAllocatedOutboundPorts() int32`

GetAllocatedOutboundPorts returns the AllocatedOutboundPorts field if non-nil, zero value otherwise.

### GetAllocatedOutboundPortsOk

`func (o *Loadbalancerprofile) GetAllocatedOutboundPortsOk() (*int32, bool)`

GetAllocatedOutboundPortsOk returns a tuple with the AllocatedOutboundPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedOutboundPorts

`func (o *Loadbalancerprofile) SetAllocatedOutboundPorts(v int32)`

SetAllocatedOutboundPorts sets AllocatedOutboundPorts field to given value.

### HasAllocatedOutboundPorts

`func (o *Loadbalancerprofile) HasAllocatedOutboundPorts() bool`

HasAllocatedOutboundPorts returns a boolean if a field has been set.

### GetEffectiveOutboundIPs

`func (o *Loadbalancerprofile) GetEffectiveOutboundIPs() []Effectiveoutboundips`

GetEffectiveOutboundIPs returns the EffectiveOutboundIPs field if non-nil, zero value otherwise.

### GetEffectiveOutboundIPsOk

`func (o *Loadbalancerprofile) GetEffectiveOutboundIPsOk() (*[]Effectiveoutboundips, bool)`

GetEffectiveOutboundIPsOk returns a tuple with the EffectiveOutboundIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveOutboundIPs

`func (o *Loadbalancerprofile) SetEffectiveOutboundIPs(v []Effectiveoutboundips)`

SetEffectiveOutboundIPs sets EffectiveOutboundIPs field to given value.

### HasEffectiveOutboundIPs

`func (o *Loadbalancerprofile) HasEffectiveOutboundIPs() bool`

HasEffectiveOutboundIPs returns a boolean if a field has been set.

### GetEnableMultipleStandardLoadBalancers

`func (o *Loadbalancerprofile) GetEnableMultipleStandardLoadBalancers() bool`

GetEnableMultipleStandardLoadBalancers returns the EnableMultipleStandardLoadBalancers field if non-nil, zero value otherwise.

### GetEnableMultipleStandardLoadBalancersOk

`func (o *Loadbalancerprofile) GetEnableMultipleStandardLoadBalancersOk() (*bool, bool)`

GetEnableMultipleStandardLoadBalancersOk returns a tuple with the EnableMultipleStandardLoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultipleStandardLoadBalancers

`func (o *Loadbalancerprofile) SetEnableMultipleStandardLoadBalancers(v bool)`

SetEnableMultipleStandardLoadBalancers sets EnableMultipleStandardLoadBalancers field to given value.

### HasEnableMultipleStandardLoadBalancers

`func (o *Loadbalancerprofile) HasEnableMultipleStandardLoadBalancers() bool`

HasEnableMultipleStandardLoadBalancers returns a boolean if a field has been set.

### GetIdleTimeoutInMinutes

`func (o *Loadbalancerprofile) GetIdleTimeoutInMinutes() int32`

GetIdleTimeoutInMinutes returns the IdleTimeoutInMinutes field if non-nil, zero value otherwise.

### GetIdleTimeoutInMinutesOk

`func (o *Loadbalancerprofile) GetIdleTimeoutInMinutesOk() (*int32, bool)`

GetIdleTimeoutInMinutesOk returns a tuple with the IdleTimeoutInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutInMinutes

`func (o *Loadbalancerprofile) SetIdleTimeoutInMinutes(v int32)`

SetIdleTimeoutInMinutes sets IdleTimeoutInMinutes field to given value.

### HasIdleTimeoutInMinutes

`func (o *Loadbalancerprofile) HasIdleTimeoutInMinutes() bool`

HasIdleTimeoutInMinutes returns a boolean if a field has been set.

### GetManagedOutboundIPs

`func (o *Loadbalancerprofile) GetManagedOutboundIPs() Managedoutboundips`

GetManagedOutboundIPs returns the ManagedOutboundIPs field if non-nil, zero value otherwise.

### GetManagedOutboundIPsOk

`func (o *Loadbalancerprofile) GetManagedOutboundIPsOk() (*Managedoutboundips, bool)`

GetManagedOutboundIPsOk returns a tuple with the ManagedOutboundIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedOutboundIPs

`func (o *Loadbalancerprofile) SetManagedOutboundIPs(v Managedoutboundips)`

SetManagedOutboundIPs sets ManagedOutboundIPs field to given value.

### HasManagedOutboundIPs

`func (o *Loadbalancerprofile) HasManagedOutboundIPs() bool`

HasManagedOutboundIPs returns a boolean if a field has been set.

### GetOutboundIPPrefixes

`func (o *Loadbalancerprofile) GetOutboundIPPrefixes() Outboundipprefixes`

GetOutboundIPPrefixes returns the OutboundIPPrefixes field if non-nil, zero value otherwise.

### GetOutboundIPPrefixesOk

`func (o *Loadbalancerprofile) GetOutboundIPPrefixesOk() (*Outboundipprefixes, bool)`

GetOutboundIPPrefixesOk returns a tuple with the OutboundIPPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundIPPrefixes

`func (o *Loadbalancerprofile) SetOutboundIPPrefixes(v Outboundipprefixes)`

SetOutboundIPPrefixes sets OutboundIPPrefixes field to given value.

### HasOutboundIPPrefixes

`func (o *Loadbalancerprofile) HasOutboundIPPrefixes() bool`

HasOutboundIPPrefixes returns a boolean if a field has been set.

### GetOutboundIPs

`func (o *Loadbalancerprofile) GetOutboundIPs() Outboundips`

GetOutboundIPs returns the OutboundIPs field if non-nil, zero value otherwise.

### GetOutboundIPsOk

`func (o *Loadbalancerprofile) GetOutboundIPsOk() (*Outboundips, bool)`

GetOutboundIPsOk returns a tuple with the OutboundIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundIPs

`func (o *Loadbalancerprofile) SetOutboundIPs(v Outboundips)`

SetOutboundIPs sets OutboundIPs field to given value.

### HasOutboundIPs

`func (o *Loadbalancerprofile) HasOutboundIPs() bool`

HasOutboundIPs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


