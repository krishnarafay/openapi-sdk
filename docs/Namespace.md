# Namespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the namespace resource | [default to "infra.k8smgmt.io/v3"]
**Kind** | **string** | Kind of the namespace resource | [default to "Namespace"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**NamespaceSpec**](NamespaceSpec.md) |  | 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewNamespace

`func NewNamespace(apiVersion string, kind string, metadata Metadata, spec NamespaceSpec, ) *Namespace`

NewNamespace instantiates a new Namespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceWithDefaults

`func NewNamespaceWithDefaults() *Namespace`

NewNamespaceWithDefaults instantiates a new Namespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *Namespace) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Namespace) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Namespace) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *Namespace) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Namespace) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Namespace) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *Namespace) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Namespace) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Namespace) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *Namespace) GetSpec() NamespaceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Namespace) GetSpecOk() (*NamespaceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Namespace) SetSpec(v NamespaceSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *Namespace) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Namespace) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Namespace) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Namespace) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


