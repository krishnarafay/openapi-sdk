# PipelineList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the pipeline list resource | [optional] [readonly] 
**Items** | Pointer to [**[]Pipeline**](Pipeline.md) | List of the pipeline resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the pipeline list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewPipelineList

`func NewPipelineList() *PipelineList`

NewPipelineList instantiates a new PipelineList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineListWithDefaults

`func NewPipelineListWithDefaults() *PipelineList`

NewPipelineListWithDefaults instantiates a new PipelineList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *PipelineList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PipelineList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PipelineList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PipelineList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *PipelineList) GetItems() []Pipeline`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PipelineList) GetItemsOk() (*[]Pipeline, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PipelineList) SetItems(v []Pipeline)`

SetItems sets Items field to given value.

### HasItems

`func (o *PipelineList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *PipelineList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PipelineList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PipelineList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PipelineList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *PipelineList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PipelineList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PipelineList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PipelineList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


