# TriggerSpecConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repo** | [**CronTriggerConfigRepo**](CronTriggerConfigRepo.md) |  | 
**CronExpression** | **string** | cron expression for trigger | 

## Methods

### NewTriggerSpecConfig

`func NewTriggerSpecConfig(repo CronTriggerConfigRepo, cronExpression string, ) *TriggerSpecConfig`

NewTriggerSpecConfig instantiates a new TriggerSpecConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerSpecConfigWithDefaults

`func NewTriggerSpecConfigWithDefaults() *TriggerSpecConfig`

NewTriggerSpecConfigWithDefaults instantiates a new TriggerSpecConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepo

`func (o *TriggerSpecConfig) GetRepo() CronTriggerConfigRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *TriggerSpecConfig) GetRepoOk() (*CronTriggerConfigRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *TriggerSpecConfig) SetRepo(v CronTriggerConfigRepo)`

SetRepo sets Repo field to given value.


### GetCronExpression

`func (o *TriggerSpecConfig) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *TriggerSpecConfig) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *TriggerSpecConfig) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


