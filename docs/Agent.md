# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the agent resource | [default to "gitops.k8smgmt.io/v3"]
**Kind** | **string** | Kind of the agent resource | [default to "Agent"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**AgentSpec**](AgentSpec.md) |  | 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewAgent

`func NewAgent(apiVersion string, kind string, metadata Metadata, spec AgentSpec, ) *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *Agent) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Agent) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Agent) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *Agent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Agent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Agent) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *Agent) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Agent) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Agent) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *Agent) GetSpec() AgentSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Agent) GetSpecOk() (*AgentSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Agent) SetSpec(v AgentSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *Agent) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Agent) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Agent) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Agent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


