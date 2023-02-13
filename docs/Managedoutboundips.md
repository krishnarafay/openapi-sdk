# Managedoutboundips

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The desired number of IPv4 outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. | [optional] 
**CountIPv6** | Pointer to **int32** | The desired number of IPv6 outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 0 for single-stack and 1 for dual-stack. | [optional] 

## Methods

### NewManagedoutboundips

`func NewManagedoutboundips() *Managedoutboundips`

NewManagedoutboundips instantiates a new Managedoutboundips object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedoutboundipsWithDefaults

`func NewManagedoutboundipsWithDefaults() *Managedoutboundips`

NewManagedoutboundipsWithDefaults instantiates a new Managedoutboundips object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *Managedoutboundips) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Managedoutboundips) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Managedoutboundips) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Managedoutboundips) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCountIPv6

`func (o *Managedoutboundips) GetCountIPv6() int32`

GetCountIPv6 returns the CountIPv6 field if non-nil, zero value otherwise.

### GetCountIPv6Ok

`func (o *Managedoutboundips) GetCountIPv6Ok() (*int32, bool)`

GetCountIPv6Ok returns a tuple with the CountIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountIPv6

`func (o *Managedoutboundips) SetCountIPv6(v int32)`

SetCountIPv6 sets CountIPv6 field to given value.

### HasCountIPv6

`func (o *Managedoutboundips) HasCountIPv6() bool`

HasCountIPv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


