# AksV3ConfigObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Spec** | Pointer to [**AksV3Spec**](AksV3Spec.md) |  | [optional] 

## Methods

### NewAksV3ConfigObject

`func NewAksV3ConfigObject() *AksV3ConfigObject`

NewAksV3ConfigObject instantiates a new AksV3ConfigObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAksV3ConfigObjectWithDefaults

`func NewAksV3ConfigObjectWithDefaults() *AksV3ConfigObject`

NewAksV3ConfigObjectWithDefaults instantiates a new AksV3ConfigObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AksV3ConfigObject) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AksV3ConfigObject) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AksV3ConfigObject) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AksV3ConfigObject) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *AksV3ConfigObject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AksV3ConfigObject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AksV3ConfigObject) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AksV3ConfigObject) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *AksV3ConfigObject) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AksV3ConfigObject) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AksV3ConfigObject) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AksV3ConfigObject) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *AksV3ConfigObject) GetSpec() AksV3Spec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AksV3ConfigObject) GetSpecOk() (*AksV3Spec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AksV3ConfigObject) SetSpec(v AksV3Spec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AksV3ConfigObject) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


