# Managedoutboundipprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The desired number of outbound IPs created/managed by Azure. Allowed values must be in the range of 1 to 16 (inclusive). The default value is 1. | [optional] 

## Methods

### NewManagedoutboundipprofile

`func NewManagedoutboundipprofile() *Managedoutboundipprofile`

NewManagedoutboundipprofile instantiates a new Managedoutboundipprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedoutboundipprofileWithDefaults

`func NewManagedoutboundipprofileWithDefaults() *Managedoutboundipprofile`

NewManagedoutboundipprofileWithDefaults instantiates a new Managedoutboundipprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *Managedoutboundipprofile) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Managedoutboundipprofile) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Managedoutboundipprofile) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Managedoutboundipprofile) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


