# Workload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the workload resource | [default to "apps.k8smgmt.io/v3"]
**Kind** | **string** | Kind of the workload resource | [default to "Workload"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**WorkloadSpec**](WorkloadSpec.md) |  | 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewWorkload

`func NewWorkload(apiVersion string, kind string, metadata Metadata, spec WorkloadSpec, ) *Workload`

NewWorkload instantiates a new Workload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWithDefaults

`func NewWorkloadWithDefaults() *Workload`

NewWorkloadWithDefaults instantiates a new Workload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *Workload) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Workload) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Workload) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *Workload) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Workload) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Workload) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *Workload) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Workload) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Workload) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *Workload) GetSpec() WorkloadSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Workload) GetSpecOk() (*WorkloadSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Workload) SetSpec(v WorkloadSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *Workload) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Workload) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Workload) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Workload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


