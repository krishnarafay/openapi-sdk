# DeployWorkloadTemplateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** | namespace workload template should be deployed in | 
**Overrides** | [**[]OverrideTemplate**](OverrideTemplate.md) | overrides for workload template | 
**Placement** | [**PlacementSpec**](PlacementSpec.md) |  | 
**UseRevisionFromWebhookTriggerEvent** | Pointer to **bool** | flag to indicate weather to deploy workload using revision recieved from webhook trigger | [optional] 
**WorkloadTemplate** | **string** | name of workload template | 

## Methods

### NewDeployWorkloadTemplateConfig

`func NewDeployWorkloadTemplateConfig(namespace string, overrides []OverrideTemplate, placement PlacementSpec, workloadTemplate string, ) *DeployWorkloadTemplateConfig`

NewDeployWorkloadTemplateConfig instantiates a new DeployWorkloadTemplateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployWorkloadTemplateConfigWithDefaults

`func NewDeployWorkloadTemplateConfigWithDefaults() *DeployWorkloadTemplateConfig`

NewDeployWorkloadTemplateConfigWithDefaults instantiates a new DeployWorkloadTemplateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *DeployWorkloadTemplateConfig) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DeployWorkloadTemplateConfig) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DeployWorkloadTemplateConfig) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetOverrides

`func (o *DeployWorkloadTemplateConfig) GetOverrides() []OverrideTemplate`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *DeployWorkloadTemplateConfig) GetOverridesOk() (*[]OverrideTemplate, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *DeployWorkloadTemplateConfig) SetOverrides(v []OverrideTemplate)`

SetOverrides sets Overrides field to given value.


### GetPlacement

`func (o *DeployWorkloadTemplateConfig) GetPlacement() PlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *DeployWorkloadTemplateConfig) GetPlacementOk() (*PlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *DeployWorkloadTemplateConfig) SetPlacement(v PlacementSpec)`

SetPlacement sets Placement field to given value.


### GetUseRevisionFromWebhookTriggerEvent

`func (o *DeployWorkloadTemplateConfig) GetUseRevisionFromWebhookTriggerEvent() bool`

GetUseRevisionFromWebhookTriggerEvent returns the UseRevisionFromWebhookTriggerEvent field if non-nil, zero value otherwise.

### GetUseRevisionFromWebhookTriggerEventOk

`func (o *DeployWorkloadTemplateConfig) GetUseRevisionFromWebhookTriggerEventOk() (*bool, bool)`

GetUseRevisionFromWebhookTriggerEventOk returns a tuple with the UseRevisionFromWebhookTriggerEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRevisionFromWebhookTriggerEvent

`func (o *DeployWorkloadTemplateConfig) SetUseRevisionFromWebhookTriggerEvent(v bool)`

SetUseRevisionFromWebhookTriggerEvent sets UseRevisionFromWebhookTriggerEvent field to given value.

### HasUseRevisionFromWebhookTriggerEvent

`func (o *DeployWorkloadTemplateConfig) HasUseRevisionFromWebhookTriggerEvent() bool`

HasUseRevisionFromWebhookTriggerEvent returns a boolean if a field has been set.

### GetWorkloadTemplate

`func (o *DeployWorkloadTemplateConfig) GetWorkloadTemplate() string`

GetWorkloadTemplate returns the WorkloadTemplate field if non-nil, zero value otherwise.

### GetWorkloadTemplateOk

`func (o *DeployWorkloadTemplateConfig) GetWorkloadTemplateOk() (*string, bool)`

GetWorkloadTemplateOk returns a tuple with the WorkloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadTemplate

`func (o *DeployWorkloadTemplateConfig) SetWorkloadTemplate(v string)`

SetWorkloadTemplate sets WorkloadTemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


