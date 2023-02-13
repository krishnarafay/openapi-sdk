# AgentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | flag to indicate if the agent is active | [optional] 
**Cluster** | [**ClusterMeta**](ClusterMeta.md) |  | 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Type** | **string** | type of agent | [default to "Cluster"]
**Version** | Pointer to **string** | version of agent | [optional] 

## Methods

### NewAgentSpec

`func NewAgentSpec(cluster ClusterMeta, type_ string, ) *AgentSpec`

NewAgentSpec instantiates a new AgentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSpecWithDefaults

`func NewAgentSpecWithDefaults() *AgentSpec`

NewAgentSpecWithDefaults instantiates a new AgentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AgentSpec) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AgentSpec) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AgentSpec) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AgentSpec) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCluster

`func (o *AgentSpec) GetCluster() ClusterMeta`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *AgentSpec) GetClusterOk() (*ClusterMeta, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *AgentSpec) SetCluster(v ClusterMeta)`

SetCluster sets Cluster field to given value.


### GetSharing

`func (o *AgentSpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *AgentSpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *AgentSpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *AgentSpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetType

`func (o *AgentSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentSpec) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *AgentSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AgentSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AgentSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AgentSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


