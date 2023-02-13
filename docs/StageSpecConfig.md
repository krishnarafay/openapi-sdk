# StageSpecConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvers** | [**[]Approver**](Approver.md) | list of approvers | 
**Timeout** | Pointer to **string** | timeout for approval | [optional] 
**Type** | **string** | type of infraprovisioner | 
**UseRevisionFromWebhookTriggerEvent** | Pointer to **bool** | flag to indicate weather to deploy workload using revision recieved from webhook trigger | [optional] 
**Workload** | **string** | name of the workload | 
**Namespace** | **string** | namespace workload template should be deployed in | 
**Overrides** | [**[]OverrideTemplate**](OverrideTemplate.md) | overrides for workload template | 
**Placement** | [**PlacementSpec**](PlacementSpec.md) |  | 
**WorkloadTemplate** | **string** | name of workload template | 
**Action** | [**InfraProvisionerConfigAction**](InfraProvisionerConfigAction.md) |  | 
**Agents** | [**[]AgentMeta**](AgentMeta.md) | list of agents to be used for provisioning | 
**PersistWorkingDirectory** | Pointer to **bool** | flag to indicate if working directory should be persisted | [optional] 
**Provisioner** | **string** | name of the infrastructure provisioner | 
**Revision** | **string** | branch or tag for the git repository used in infrastructure provisioner | 
**WorkingDirectory** | Pointer to **string** | working directory for the provisioner | [optional] 
**CommitterEmail** | Pointer to **string** | value to use as committer email when committing to destination git repository. | [optional] 
**DestinationRepo** | Pointer to [**SystemSyncRepo**](SystemSyncRepo.md) |  | [optional] 
**ExcludedResources** | Pointer to [**[]SystemSyncResource**](SystemSyncResource.md) | resources to exclude from syncing | [optional] 
**GitToSystemSync** | Pointer to **bool** | flag to indicate if git to system sync should be enabled | [optional] 
**IncludedResources** | Pointer to [**[]SystemSyncResource**](SystemSyncResource.md) | resources to include for syncing | [optional] 
**SourceAsDestination** | Pointer to **bool** | flag to indicate if source repository should be used as destination repository | [optional] 
**SourceRepo** | Pointer to [**SystemSyncRepo**](SystemSyncRepo.md) |  | [optional] 
**SystemToGitSync** | Pointer to **bool** | flat to indicate if system to git sync should be enabled | [optional] 

## Methods

### NewStageSpecConfig

`func NewStageSpecConfig(approvers []Approver, type_ string, workload string, namespace string, overrides []OverrideTemplate, placement PlacementSpec, workloadTemplate string, action InfraProvisionerConfigAction, agents []AgentMeta, provisioner string, revision string, ) *StageSpecConfig`

NewStageSpecConfig instantiates a new StageSpecConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageSpecConfigWithDefaults

`func NewStageSpecConfigWithDefaults() *StageSpecConfig`

NewStageSpecConfigWithDefaults instantiates a new StageSpecConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovers

`func (o *StageSpecConfig) GetApprovers() []Approver`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *StageSpecConfig) GetApproversOk() (*[]Approver, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *StageSpecConfig) SetApprovers(v []Approver)`

SetApprovers sets Approvers field to given value.


### GetTimeout

`func (o *StageSpecConfig) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *StageSpecConfig) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *StageSpecConfig) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *StageSpecConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *StageSpecConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StageSpecConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StageSpecConfig) SetType(v string)`

SetType sets Type field to given value.


### GetUseRevisionFromWebhookTriggerEvent

`func (o *StageSpecConfig) GetUseRevisionFromWebhookTriggerEvent() bool`

GetUseRevisionFromWebhookTriggerEvent returns the UseRevisionFromWebhookTriggerEvent field if non-nil, zero value otherwise.

### GetUseRevisionFromWebhookTriggerEventOk

`func (o *StageSpecConfig) GetUseRevisionFromWebhookTriggerEventOk() (*bool, bool)`

GetUseRevisionFromWebhookTriggerEventOk returns a tuple with the UseRevisionFromWebhookTriggerEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRevisionFromWebhookTriggerEvent

`func (o *StageSpecConfig) SetUseRevisionFromWebhookTriggerEvent(v bool)`

SetUseRevisionFromWebhookTriggerEvent sets UseRevisionFromWebhookTriggerEvent field to given value.

### HasUseRevisionFromWebhookTriggerEvent

`func (o *StageSpecConfig) HasUseRevisionFromWebhookTriggerEvent() bool`

HasUseRevisionFromWebhookTriggerEvent returns a boolean if a field has been set.

### GetWorkload

`func (o *StageSpecConfig) GetWorkload() string`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *StageSpecConfig) GetWorkloadOk() (*string, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *StageSpecConfig) SetWorkload(v string)`

SetWorkload sets Workload field to given value.


### GetNamespace

`func (o *StageSpecConfig) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *StageSpecConfig) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *StageSpecConfig) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetOverrides

`func (o *StageSpecConfig) GetOverrides() []OverrideTemplate`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *StageSpecConfig) GetOverridesOk() (*[]OverrideTemplate, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *StageSpecConfig) SetOverrides(v []OverrideTemplate)`

SetOverrides sets Overrides field to given value.


### GetPlacement

`func (o *StageSpecConfig) GetPlacement() PlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *StageSpecConfig) GetPlacementOk() (*PlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *StageSpecConfig) SetPlacement(v PlacementSpec)`

SetPlacement sets Placement field to given value.


### GetWorkloadTemplate

`func (o *StageSpecConfig) GetWorkloadTemplate() string`

GetWorkloadTemplate returns the WorkloadTemplate field if non-nil, zero value otherwise.

### GetWorkloadTemplateOk

`func (o *StageSpecConfig) GetWorkloadTemplateOk() (*string, bool)`

GetWorkloadTemplateOk returns a tuple with the WorkloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadTemplate

`func (o *StageSpecConfig) SetWorkloadTemplate(v string)`

SetWorkloadTemplate sets WorkloadTemplate field to given value.


### GetAction

`func (o *StageSpecConfig) GetAction() InfraProvisionerConfigAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *StageSpecConfig) GetActionOk() (*InfraProvisionerConfigAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *StageSpecConfig) SetAction(v InfraProvisionerConfigAction)`

SetAction sets Action field to given value.


### GetAgents

`func (o *StageSpecConfig) GetAgents() []AgentMeta`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *StageSpecConfig) GetAgentsOk() (*[]AgentMeta, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *StageSpecConfig) SetAgents(v []AgentMeta)`

SetAgents sets Agents field to given value.


### GetPersistWorkingDirectory

`func (o *StageSpecConfig) GetPersistWorkingDirectory() bool`

GetPersistWorkingDirectory returns the PersistWorkingDirectory field if non-nil, zero value otherwise.

### GetPersistWorkingDirectoryOk

`func (o *StageSpecConfig) GetPersistWorkingDirectoryOk() (*bool, bool)`

GetPersistWorkingDirectoryOk returns a tuple with the PersistWorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistWorkingDirectory

`func (o *StageSpecConfig) SetPersistWorkingDirectory(v bool)`

SetPersistWorkingDirectory sets PersistWorkingDirectory field to given value.

### HasPersistWorkingDirectory

`func (o *StageSpecConfig) HasPersistWorkingDirectory() bool`

HasPersistWorkingDirectory returns a boolean if a field has been set.

### GetProvisioner

`func (o *StageSpecConfig) GetProvisioner() string`

GetProvisioner returns the Provisioner field if non-nil, zero value otherwise.

### GetProvisionerOk

`func (o *StageSpecConfig) GetProvisionerOk() (*string, bool)`

GetProvisionerOk returns a tuple with the Provisioner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioner

`func (o *StageSpecConfig) SetProvisioner(v string)`

SetProvisioner sets Provisioner field to given value.


### GetRevision

`func (o *StageSpecConfig) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StageSpecConfig) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StageSpecConfig) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetWorkingDirectory

`func (o *StageSpecConfig) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *StageSpecConfig) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *StageSpecConfig) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *StageSpecConfig) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetCommitterEmail

`func (o *StageSpecConfig) GetCommitterEmail() string`

GetCommitterEmail returns the CommitterEmail field if non-nil, zero value otherwise.

### GetCommitterEmailOk

`func (o *StageSpecConfig) GetCommitterEmailOk() (*string, bool)`

GetCommitterEmailOk returns a tuple with the CommitterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitterEmail

`func (o *StageSpecConfig) SetCommitterEmail(v string)`

SetCommitterEmail sets CommitterEmail field to given value.

### HasCommitterEmail

`func (o *StageSpecConfig) HasCommitterEmail() bool`

HasCommitterEmail returns a boolean if a field has been set.

### GetDestinationRepo

`func (o *StageSpecConfig) GetDestinationRepo() SystemSyncRepo`

GetDestinationRepo returns the DestinationRepo field if non-nil, zero value otherwise.

### GetDestinationRepoOk

`func (o *StageSpecConfig) GetDestinationRepoOk() (*SystemSyncRepo, bool)`

GetDestinationRepoOk returns a tuple with the DestinationRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationRepo

`func (o *StageSpecConfig) SetDestinationRepo(v SystemSyncRepo)`

SetDestinationRepo sets DestinationRepo field to given value.

### HasDestinationRepo

`func (o *StageSpecConfig) HasDestinationRepo() bool`

HasDestinationRepo returns a boolean if a field has been set.

### GetExcludedResources

`func (o *StageSpecConfig) GetExcludedResources() []SystemSyncResource`

GetExcludedResources returns the ExcludedResources field if non-nil, zero value otherwise.

### GetExcludedResourcesOk

`func (o *StageSpecConfig) GetExcludedResourcesOk() (*[]SystemSyncResource, bool)`

GetExcludedResourcesOk returns a tuple with the ExcludedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedResources

`func (o *StageSpecConfig) SetExcludedResources(v []SystemSyncResource)`

SetExcludedResources sets ExcludedResources field to given value.

### HasExcludedResources

`func (o *StageSpecConfig) HasExcludedResources() bool`

HasExcludedResources returns a boolean if a field has been set.

### GetGitToSystemSync

`func (o *StageSpecConfig) GetGitToSystemSync() bool`

GetGitToSystemSync returns the GitToSystemSync field if non-nil, zero value otherwise.

### GetGitToSystemSyncOk

`func (o *StageSpecConfig) GetGitToSystemSyncOk() (*bool, bool)`

GetGitToSystemSyncOk returns a tuple with the GitToSystemSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitToSystemSync

`func (o *StageSpecConfig) SetGitToSystemSync(v bool)`

SetGitToSystemSync sets GitToSystemSync field to given value.

### HasGitToSystemSync

`func (o *StageSpecConfig) HasGitToSystemSync() bool`

HasGitToSystemSync returns a boolean if a field has been set.

### GetIncludedResources

`func (o *StageSpecConfig) GetIncludedResources() []SystemSyncResource`

GetIncludedResources returns the IncludedResources field if non-nil, zero value otherwise.

### GetIncludedResourcesOk

`func (o *StageSpecConfig) GetIncludedResourcesOk() (*[]SystemSyncResource, bool)`

GetIncludedResourcesOk returns a tuple with the IncludedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedResources

`func (o *StageSpecConfig) SetIncludedResources(v []SystemSyncResource)`

SetIncludedResources sets IncludedResources field to given value.

### HasIncludedResources

`func (o *StageSpecConfig) HasIncludedResources() bool`

HasIncludedResources returns a boolean if a field has been set.

### GetSourceAsDestination

`func (o *StageSpecConfig) GetSourceAsDestination() bool`

GetSourceAsDestination returns the SourceAsDestination field if non-nil, zero value otherwise.

### GetSourceAsDestinationOk

`func (o *StageSpecConfig) GetSourceAsDestinationOk() (*bool, bool)`

GetSourceAsDestinationOk returns a tuple with the SourceAsDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAsDestination

`func (o *StageSpecConfig) SetSourceAsDestination(v bool)`

SetSourceAsDestination sets SourceAsDestination field to given value.

### HasSourceAsDestination

`func (o *StageSpecConfig) HasSourceAsDestination() bool`

HasSourceAsDestination returns a boolean if a field has been set.

### GetSourceRepo

`func (o *StageSpecConfig) GetSourceRepo() SystemSyncRepo`

GetSourceRepo returns the SourceRepo field if non-nil, zero value otherwise.

### GetSourceRepoOk

`func (o *StageSpecConfig) GetSourceRepoOk() (*SystemSyncRepo, bool)`

GetSourceRepoOk returns a tuple with the SourceRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepo

`func (o *StageSpecConfig) SetSourceRepo(v SystemSyncRepo)`

SetSourceRepo sets SourceRepo field to given value.

### HasSourceRepo

`func (o *StageSpecConfig) HasSourceRepo() bool`

HasSourceRepo returns a boolean if a field has been set.

### GetSystemToGitSync

`func (o *StageSpecConfig) GetSystemToGitSync() bool`

GetSystemToGitSync returns the SystemToGitSync field if non-nil, zero value otherwise.

### GetSystemToGitSyncOk

`func (o *StageSpecConfig) GetSystemToGitSyncOk() (*bool, bool)`

GetSystemToGitSyncOk returns a tuple with the SystemToGitSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemToGitSync

`func (o *StageSpecConfig) SetSystemToGitSync(v bool)`

SetSystemToGitSync sets SystemToGitSync field to given value.

### HasSystemToGitSync

`func (o *StageSpecConfig) HasSystemToGitSync() bool`

HasSystemToGitSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


