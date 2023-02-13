# OPAConstraintTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the constraint template resource | [default to "opa.k8smgmt.io/v3"]
**Kind** | **string** | Kind of the constraint template resource | [default to "OPAConstraintTemplate"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**OPAConstraintTemplateSpec**](OPAConstraintTemplateSpec.md) |  | 

## Methods

### NewOPAConstraintTemplate

`func NewOPAConstraintTemplate(apiVersion string, kind string, metadata Metadata, spec OPAConstraintTemplateSpec, ) *OPAConstraintTemplate`

NewOPAConstraintTemplate instantiates a new OPAConstraintTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOPAConstraintTemplateWithDefaults

`func NewOPAConstraintTemplateWithDefaults() *OPAConstraintTemplate`

NewOPAConstraintTemplateWithDefaults instantiates a new OPAConstraintTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *OPAConstraintTemplate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *OPAConstraintTemplate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *OPAConstraintTemplate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *OPAConstraintTemplate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OPAConstraintTemplate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OPAConstraintTemplate) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *OPAConstraintTemplate) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OPAConstraintTemplate) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OPAConstraintTemplate) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *OPAConstraintTemplate) GetSpec() OPAConstraintTemplateSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *OPAConstraintTemplate) GetSpecOk() (*OPAConstraintTemplateSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *OPAConstraintTemplate) SetSpec(v OPAConstraintTemplateSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


