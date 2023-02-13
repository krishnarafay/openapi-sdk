# ChargebackGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the group resource | [default to "system.k8smgmt.io/v3"]
**Kind** | **string** | Kind of the group resource | [default to "ChargebackGroup"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**ChargebackGroupSpec**](ChargebackGroupSpec.md) |  | 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewChargebackGroup

`func NewChargebackGroup(apiVersion string, kind string, metadata Metadata, spec ChargebackGroupSpec, ) *ChargebackGroup`

NewChargebackGroup instantiates a new ChargebackGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargebackGroupWithDefaults

`func NewChargebackGroupWithDefaults() *ChargebackGroup`

NewChargebackGroupWithDefaults instantiates a new ChargebackGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ChargebackGroup) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ChargebackGroup) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ChargebackGroup) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ChargebackGroup) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ChargebackGroup) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ChargebackGroup) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ChargebackGroup) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ChargebackGroup) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ChargebackGroup) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *ChargebackGroup) GetSpec() ChargebackGroupSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ChargebackGroup) GetSpecOk() (*ChargebackGroupSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ChargebackGroup) SetSpec(v ChargebackGroupSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *ChargebackGroup) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChargebackGroup) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChargebackGroup) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChargebackGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


