# WorkloadTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the workload template resource | [default to "apps.k8smgmt.io/v3"]
**Kind** | **string** | Kind of the workload template resource | [default to "WorkloadTemplate"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**WorkloadTemplateSpec**](WorkloadTemplateSpec.md) |  | 

## Methods

### NewWorkloadTemplate

`func NewWorkloadTemplate(apiVersion string, kind string, metadata Metadata, spec WorkloadTemplateSpec, ) *WorkloadTemplate`

NewWorkloadTemplate instantiates a new WorkloadTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadTemplateWithDefaults

`func NewWorkloadTemplateWithDefaults() *WorkloadTemplate`

NewWorkloadTemplateWithDefaults instantiates a new WorkloadTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *WorkloadTemplate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *WorkloadTemplate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *WorkloadTemplate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *WorkloadTemplate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WorkloadTemplate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WorkloadTemplate) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *WorkloadTemplate) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WorkloadTemplate) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WorkloadTemplate) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *WorkloadTemplate) GetSpec() WorkloadTemplateSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *WorkloadTemplate) GetSpecOk() (*WorkloadTemplateSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *WorkloadTemplate) SetSpec(v WorkloadTemplateSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


