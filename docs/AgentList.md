# AgentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the agent list resource | [optional] [readonly] 
**Items** | Pointer to [**[]Agent**](Agent.md) | List of the agent resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the agent list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewAgentList

`func NewAgentList() *AgentList`

NewAgentList instantiates a new AgentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentListWithDefaults

`func NewAgentListWithDefaults() *AgentList`

NewAgentListWithDefaults instantiates a new AgentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AgentList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AgentList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AgentList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AgentList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *AgentList) GetItems() []Agent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AgentList) GetItemsOk() (*[]Agent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AgentList) SetItems(v []Agent)`

SetItems sets Items field to given value.

### HasItems

`func (o *AgentList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *AgentList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AgentList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AgentList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AgentList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *AgentList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AgentList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AgentList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AgentList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


