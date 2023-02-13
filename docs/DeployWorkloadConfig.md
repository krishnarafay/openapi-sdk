# DeployWorkloadConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseRevisionFromWebhookTriggerEvent** | Pointer to **bool** | flag to indicate weather to deploy workload using revision recieved from webhook trigger | [optional] 
**Workload** | **string** | name of the workload | 

## Methods

### NewDeployWorkloadConfig

`func NewDeployWorkloadConfig(workload string, ) *DeployWorkloadConfig`

NewDeployWorkloadConfig instantiates a new DeployWorkloadConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployWorkloadConfigWithDefaults

`func NewDeployWorkloadConfigWithDefaults() *DeployWorkloadConfig`

NewDeployWorkloadConfigWithDefaults instantiates a new DeployWorkloadConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseRevisionFromWebhookTriggerEvent

`func (o *DeployWorkloadConfig) GetUseRevisionFromWebhookTriggerEvent() bool`

GetUseRevisionFromWebhookTriggerEvent returns the UseRevisionFromWebhookTriggerEvent field if non-nil, zero value otherwise.

### GetUseRevisionFromWebhookTriggerEventOk

`func (o *DeployWorkloadConfig) GetUseRevisionFromWebhookTriggerEventOk() (*bool, bool)`

GetUseRevisionFromWebhookTriggerEventOk returns a tuple with the UseRevisionFromWebhookTriggerEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRevisionFromWebhookTriggerEvent

`func (o *DeployWorkloadConfig) SetUseRevisionFromWebhookTriggerEvent(v bool)`

SetUseRevisionFromWebhookTriggerEvent sets UseRevisionFromWebhookTriggerEvent field to given value.

### HasUseRevisionFromWebhookTriggerEvent

`func (o *DeployWorkloadConfig) HasUseRevisionFromWebhookTriggerEvent() bool`

HasUseRevisionFromWebhookTriggerEvent returns a boolean if a field has been set.

### GetWorkload

`func (o *DeployWorkloadConfig) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *DeployWorkloadConfig) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *DeployWorkloadConfig) SetWorkload(v string)`

SetWorkload sets Workload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


