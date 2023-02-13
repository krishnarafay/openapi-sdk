# Approver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsoUser** | Pointer to **bool** | flag to specify if the user is a SSO User | [optional] 
**UserName** | **string** | user name of the approver | 

## Methods

### NewApprover

`func NewApprover(userName string, ) *Approver`

NewApprover instantiates a new Approver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproverWithDefaults

`func NewApproverWithDefaults() *Approver`

NewApproverWithDefaults instantiates a new Approver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsoUser

`func (o *Approver) GetSsoUser() bool`

GetSsoUser returns the SsoUser field if non-nil, zero value otherwise.

### GetSsoUserOk

`func (o *Approver) GetSsoUserOk() (*bool, bool)`

GetSsoUserOk returns a tuple with the SsoUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUser

`func (o *Approver) SetSsoUser(v bool)`

SetSsoUser sets SsoUser field to given value.

### HasSsoUser

`func (o *Approver) HasSsoUser() bool`

HasSsoUser returns a boolean if a field has been set.

### GetUserName

`func (o *Approver) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *Approver) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *Approver) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


