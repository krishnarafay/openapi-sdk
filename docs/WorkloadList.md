# WorkloadList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the workload list resource | [optional] [readonly] 
**Items** | Pointer to [**[]Workload**](Workload.md) | List of the workload resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the workload list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewWorkloadList

`func NewWorkloadList() *WorkloadList`

NewWorkloadList instantiates a new WorkloadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadListWithDefaults

`func NewWorkloadListWithDefaults() *WorkloadList`

NewWorkloadListWithDefaults instantiates a new WorkloadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *WorkloadList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *WorkloadList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *WorkloadList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *WorkloadList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *WorkloadList) GetItems() []Workload`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WorkloadList) GetItemsOk() (*[]Workload, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WorkloadList) SetItems(v []Workload)`

SetItems sets Items field to given value.

### HasItems

`func (o *WorkloadList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *WorkloadList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WorkloadList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WorkloadList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WorkloadList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *WorkloadList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WorkloadList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WorkloadList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WorkloadList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


