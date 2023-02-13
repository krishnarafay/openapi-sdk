# ApprovalConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvers** | [**[]Approver**](Approver.md) | list of approvers | 
**Timeout** | Pointer to **string** | timeout for approval | [optional] 
**Type** | Pointer to **string** | type of approval | [optional] 

## Methods

### NewApprovalConfiguration

`func NewApprovalConfiguration(approvers []Approver, ) *ApprovalConfiguration`

NewApprovalConfiguration instantiates a new ApprovalConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalConfigurationWithDefaults

`func NewApprovalConfigurationWithDefaults() *ApprovalConfiguration`

NewApprovalConfigurationWithDefaults instantiates a new ApprovalConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovers

`func (o *ApprovalConfiguration) GetApprovers() []Approver`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *ApprovalConfiguration) GetApproversOk() (*[]Approver, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *ApprovalConfiguration) SetApprovers(v []Approver)`

SetApprovers sets Approvers field to given value.


### GetTimeout

`func (o *ApprovalConfiguration) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ApprovalConfiguration) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ApprovalConfiguration) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ApprovalConfiguration) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *ApprovalConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApprovalConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApprovalConfiguration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApprovalConfiguration) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


