# Natgatewayprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveOutboundIPs** | Pointer to [**[]Effectiveoutboundips**](Effectiveoutboundips.md) | The effective outbound IP resources of the cluster NAT gateway. | [optional] 
**IdleTimeoutInMinutes** | Pointer to **int32** | Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120 (inclusive). The default value is 4 minutes. | [optional] 
**ManagedOutboundIPProfile** | Pointer to [**Managedoutboundipprofile**](Managedoutboundipprofile.md) |  | [optional] 

## Methods

### NewNatgatewayprofile

`func NewNatgatewayprofile() *Natgatewayprofile`

NewNatgatewayprofile instantiates a new Natgatewayprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatgatewayprofileWithDefaults

`func NewNatgatewayprofileWithDefaults() *Natgatewayprofile`

NewNatgatewayprofileWithDefaults instantiates a new Natgatewayprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveOutboundIPs

`func (o *Natgatewayprofile) GetEffectiveOutboundIPs() []Effectiveoutboundips`

GetEffectiveOutboundIPs returns the EffectiveOutboundIPs field if non-nil, zero value otherwise.

### GetEffectiveOutboundIPsOk

`func (o *Natgatewayprofile) GetEffectiveOutboundIPsOk() (*[]Effectiveoutboundips, bool)`

GetEffectiveOutboundIPsOk returns a tuple with the EffectiveOutboundIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveOutboundIPs

`func (o *Natgatewayprofile) SetEffectiveOutboundIPs(v []Effectiveoutboundips)`

SetEffectiveOutboundIPs sets EffectiveOutboundIPs field to given value.

### HasEffectiveOutboundIPs

`func (o *Natgatewayprofile) HasEffectiveOutboundIPs() bool`

HasEffectiveOutboundIPs returns a boolean if a field has been set.

### GetIdleTimeoutInMinutes

`func (o *Natgatewayprofile) GetIdleTimeoutInMinutes() int32`

GetIdleTimeoutInMinutes returns the IdleTimeoutInMinutes field if non-nil, zero value otherwise.

### GetIdleTimeoutInMinutesOk

`func (o *Natgatewayprofile) GetIdleTimeoutInMinutesOk() (*int32, bool)`

GetIdleTimeoutInMinutesOk returns a tuple with the IdleTimeoutInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutInMinutes

`func (o *Natgatewayprofile) SetIdleTimeoutInMinutes(v int32)`

SetIdleTimeoutInMinutes sets IdleTimeoutInMinutes field to given value.

### HasIdleTimeoutInMinutes

`func (o *Natgatewayprofile) HasIdleTimeoutInMinutes() bool`

HasIdleTimeoutInMinutes returns a boolean if a field has been set.

### GetManagedOutboundIPProfile

`func (o *Natgatewayprofile) GetManagedOutboundIPProfile() Managedoutboundipprofile`

GetManagedOutboundIPProfile returns the ManagedOutboundIPProfile field if non-nil, zero value otherwise.

### GetManagedOutboundIPProfileOk

`func (o *Natgatewayprofile) GetManagedOutboundIPProfileOk() (*Managedoutboundipprofile, bool)`

GetManagedOutboundIPProfileOk returns a tuple with the ManagedOutboundIPProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedOutboundIPProfile

`func (o *Natgatewayprofile) SetManagedOutboundIPProfile(v Managedoutboundipprofile)`

SetManagedOutboundIPProfile sets ManagedOutboundIPProfile field to given value.

### HasManagedOutboundIPProfile

`func (o *Natgatewayprofile) HasManagedOutboundIPProfile() bool`

HasManagedOutboundIPProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


