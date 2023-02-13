# ApprovalConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvers** | [**[]Approver**](Approver.md) | list of approvers | 
**Timeout** | Pointer to **string** | timeout for approval | [optional] 
**Type** | Pointer to **string** | type of approval | [optional] 

## Methods

### NewApprovalConfig

`func NewApprovalConfig(approvers []Approver, ) *ApprovalConfig`

NewApprovalConfig instantiates a new ApprovalConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalConfigWithDefaults

`func NewApprovalConfigWithDefaults() *ApprovalConfig`

NewApprovalConfigWithDefaults instantiates a new ApprovalConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovers

`func (o *ApprovalConfig) GetApprovers() []Approver`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *ApprovalConfig) GetApproversOk() (*[]Approver, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *ApprovalConfig) SetApprovers(v []Approver)`

SetApprovers sets Approvers field to given value.


### GetTimeout

`func (o *ApprovalConfig) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ApprovalConfig) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ApprovalConfig) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ApprovalConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *ApprovalConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApprovalConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApprovalConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApprovalConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


