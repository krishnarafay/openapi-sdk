# Ssh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKeys** | Pointer to [**[]Publickeys**](Publickeys.md) | The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified. | [optional] 

## Methods

### NewSsh

`func NewSsh() *Ssh`

NewSsh instantiates a new Ssh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshWithDefaults

`func NewSshWithDefaults() *Ssh`

NewSshWithDefaults instantiates a new Ssh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKeys

`func (o *Ssh) GetPublicKeys() []Publickeys`

GetPublicKeys returns the PublicKeys field if non-nil, zero value otherwise.

### GetPublicKeysOk

`func (o *Ssh) GetPublicKeysOk() (*[]Publickeys, bool)`

GetPublicKeysOk returns a tuple with the PublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeys

`func (o *Ssh) SetPublicKeys(v []Publickeys)`

SetPublicKeys sets PublicKeys field to given value.

### HasPublicKeys

`func (o *Ssh) HasPublicKeys() bool`

HasPublicKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


