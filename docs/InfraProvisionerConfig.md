# InfraProvisionerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**InfraProvisionerConfigAction**](InfraProvisionerConfigAction.md) |  | 
**Agents** | [**[]AgentMeta**](AgentMeta.md) | list of agents to be used for provisioning | 
**PersistWorkingDirectory** | Pointer to **bool** | flag to indicate if working directory should be persisted | [optional] 
**Provisioner** | **string** | name of the infrastructure provisioner | 
**Revision** | **string** | branch or tag for the git repository used in infrastructure provisioner | 
**Type** | **string** | type of infraprovisioner | 
**WorkingDirectory** | Pointer to **string** | working directory for the provisioner | [optional] 

## Methods

### NewInfraProvisionerConfig

`func NewInfraProvisionerConfig(action InfraProvisionerConfigAction, agents []AgentMeta, provisioner string, revision string, type_ string, ) *InfraProvisionerConfig`

NewInfraProvisionerConfig instantiates a new InfraProvisionerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraProvisionerConfigWithDefaults

`func NewInfraProvisionerConfigWithDefaults() *InfraProvisionerConfig`

NewInfraProvisionerConfigWithDefaults instantiates a new InfraProvisionerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *InfraProvisionerConfig) GetAction() InfraProvisionerConfigAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InfraProvisionerConfig) GetActionOk() (*InfraProvisionerConfigAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InfraProvisionerConfig) SetAction(v InfraProvisionerConfigAction)`

SetAction sets Action field to given value.


### GetAgents

`func (o *InfraProvisionerConfig) GetAgents() []AgentMeta`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *InfraProvisionerConfig) GetAgentsOk() (*[]AgentMeta, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *InfraProvisionerConfig) SetAgents(v []AgentMeta)`

SetAgents sets Agents field to given value.


### GetPersistWorkingDirectory

`func (o *InfraProvisionerConfig) GetPersistWorkingDirectory() bool`

GetPersistWorkingDirectory returns the PersistWorkingDirectory field if non-nil, zero value otherwise.

### GetPersistWorkingDirectoryOk

`func (o *InfraProvisionerConfig) GetPersistWorkingDirectoryOk() (*bool, bool)`

GetPersistWorkingDirectoryOk returns a tuple with the PersistWorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistWorkingDirectory

`func (o *InfraProvisionerConfig) SetPersistWorkingDirectory(v bool)`

SetPersistWorkingDirectory sets PersistWorkingDirectory field to given value.

### HasPersistWorkingDirectory

`func (o *InfraProvisionerConfig) HasPersistWorkingDirectory() bool`

HasPersistWorkingDirectory returns a boolean if a field has been set.

### GetProvisioner

`func (o *InfraProvisionerConfig) GetProvisioner() string`

GetProvisioner returns the Provisioner field if non-nil, zero value otherwise.

### GetProvisionerOk

`func (o *InfraProvisionerConfig) GetProvisionerOk() (*string, bool)`

GetProvisionerOk returns a tuple with the Provisioner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioner

`func (o *InfraProvisionerConfig) SetProvisioner(v string)`

SetProvisioner sets Provisioner field to given value.


### GetRevision

`func (o *InfraProvisionerConfig) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *InfraProvisionerConfig) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *InfraProvisionerConfig) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetType

`func (o *InfraProvisionerConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InfraProvisionerConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InfraProvisionerConfig) SetType(v string)`

SetType sets Type field to given value.


### GetWorkingDirectory

`func (o *InfraProvisionerConfig) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *InfraProvisionerConfig) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *InfraProvisionerConfig) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *InfraProvisionerConfig) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


