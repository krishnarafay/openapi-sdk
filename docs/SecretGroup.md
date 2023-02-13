# SecretGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the SecretGroup resource | [default to "gitops.k8smgmt.io/v3"]
**Kind** | **string** | Kind of the SecretGroup resource | [default to "SecretGroup"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**SecretGroupSpec**](SecretGroupSpec.md) |  | 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewSecretGroup

`func NewSecretGroup(apiVersion string, kind string, metadata Metadata, spec SecretGroupSpec, ) *SecretGroup`

NewSecretGroup instantiates a new SecretGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretGroupWithDefaults

`func NewSecretGroupWithDefaults() *SecretGroup`

NewSecretGroupWithDefaults instantiates a new SecretGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SecretGroup) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SecretGroup) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SecretGroup) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *SecretGroup) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SecretGroup) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SecretGroup) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *SecretGroup) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecretGroup) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecretGroup) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *SecretGroup) GetSpec() SecretGroupSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SecretGroup) GetSpecOk() (*SecretGroupSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SecretGroup) SetSpec(v SecretGroupSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *SecretGroup) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecretGroup) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecretGroup) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecretGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


