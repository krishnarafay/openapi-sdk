# Linuxprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminUsername** | Pointer to **string** | The administrator username to use for Linux VMs. | [optional] 
**Ssh** | Pointer to [**Ssh**](Ssh.md) |  | [optional] 

## Methods

### NewLinuxprofile

`func NewLinuxprofile() *Linuxprofile`

NewLinuxprofile instantiates a new Linuxprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinuxprofileWithDefaults

`func NewLinuxprofileWithDefaults() *Linuxprofile`

NewLinuxprofileWithDefaults instantiates a new Linuxprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminUsername

`func (o *Linuxprofile) GetAdminUsername() string`

GetAdminUsername returns the AdminUsername field if non-nil, zero value otherwise.

### GetAdminUsernameOk

`func (o *Linuxprofile) GetAdminUsernameOk() (*string, bool)`

GetAdminUsernameOk returns a tuple with the AdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUsername

`func (o *Linuxprofile) SetAdminUsername(v string)`

SetAdminUsername sets AdminUsername field to given value.

### HasAdminUsername

`func (o *Linuxprofile) HasAdminUsername() bool`

HasAdminUsername returns a boolean if a field has been set.

### GetSsh

`func (o *Linuxprofile) GetSsh() Ssh`

GetSsh returns the Ssh field if non-nil, zero value otherwise.

### GetSshOk

`func (o *Linuxprofile) GetSshOk() (*Ssh, bool)`

GetSshOk returns a tuple with the Ssh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsh

`func (o *Linuxprofile) SetSsh(v Ssh)`

SetSsh sets Ssh field to given value.

### HasSsh

`func (o *Linuxprofile) HasSsh() bool`

HasSsh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


