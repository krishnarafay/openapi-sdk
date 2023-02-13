# Outboundips

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIPs** | Pointer to [**[]Publicips**](Publicips.md) | A list of public IP resources. | [optional] 

## Methods

### NewOutboundips

`func NewOutboundips() *Outboundips`

NewOutboundips instantiates a new Outboundips object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundipsWithDefaults

`func NewOutboundipsWithDefaults() *Outboundips`

NewOutboundipsWithDefaults instantiates a new Outboundips object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIPs

`func (o *Outboundips) GetPublicIPs() []Publicips`

GetPublicIPs returns the PublicIPs field if non-nil, zero value otherwise.

### GetPublicIPsOk

`func (o *Outboundips) GetPublicIPsOk() (*[]Publicips, bool)`

GetPublicIPsOk returns a tuple with the PublicIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIPs

`func (o *Outboundips) SetPublicIPs(v []Publicips)`

SetPublicIPs sets PublicIPs field to given value.

### HasPublicIPs

`func (o *Outboundips) HasPublicIPs() bool`

HasPublicIPs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


