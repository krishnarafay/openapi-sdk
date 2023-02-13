# ResourceTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the resource template | [default to "eaas.k8smgmt.io/v1"]
**Kind** | **string** | Kind of the config context resource | [default to "ResourceTemplate"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**ResourceTemplateSpec**](ResourceTemplateSpec.md) |  | 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewResourceTemplate

`func NewResourceTemplate(apiVersion string, kind string, metadata Metadata, spec ResourceTemplateSpec, ) *ResourceTemplate`

NewResourceTemplate instantiates a new ResourceTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTemplateWithDefaults

`func NewResourceTemplateWithDefaults() *ResourceTemplate`

NewResourceTemplateWithDefaults instantiates a new ResourceTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ResourceTemplate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ResourceTemplate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ResourceTemplate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ResourceTemplate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ResourceTemplate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ResourceTemplate) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ResourceTemplate) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceTemplate) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceTemplate) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *ResourceTemplate) GetSpec() ResourceTemplateSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ResourceTemplate) GetSpecOk() (*ResourceTemplateSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ResourceTemplate) SetSpec(v ResourceTemplateSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *ResourceTemplate) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceTemplate) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceTemplate) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


