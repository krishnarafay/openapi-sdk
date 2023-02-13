# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of resource | [default to "gitops.k8smgmt.io/v3"]
**Kind** | **string** | Kind of resource | [default to "Resource"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**ResourceSpec**](ResourceSpec.md) |  | 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewResource

`func NewResource(apiVersion string, kind string, metadata Metadata, spec ResourceSpec, ) *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *Resource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Resource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Resource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *Resource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Resource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Resource) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *Resource) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Resource) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Resource) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *Resource) GetSpec() ResourceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Resource) GetSpecOk() (*ResourceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Resource) SetSpec(v ResourceSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *Resource) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Resource) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Resource) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Resource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


