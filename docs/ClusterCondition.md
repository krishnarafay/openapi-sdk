# ClusterCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdated** | Pointer to [**StatusTime**](StatusTime.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterCondition

`func NewClusterCondition() *ClusterCondition`

NewClusterCondition instantiates a new ClusterCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConditionWithDefaults

`func NewClusterConditionWithDefaults() *ClusterCondition`

NewClusterConditionWithDefaults instantiates a new ClusterCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdated

`func (o *ClusterCondition) GetLastUpdated() StatusTime`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ClusterCondition) GetLastUpdatedOk() (*StatusTime, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ClusterCondition) SetLastUpdated(v StatusTime)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ClusterCondition) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetReason

`func (o *ClusterCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ClusterCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterCondition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterCondition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ClusterCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterCondition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterCondition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


