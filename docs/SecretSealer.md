# SecretSealer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the secret sealer resource | [default to "integrations.k8smgmt.io/v3"]
**Kind** | **string** | Kind of the secret sealer resource | [default to "SecretSealer"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**SecretSealerSpec**](SecretSealerSpec.md) |  | 

## Methods

### NewSecretSealer

`func NewSecretSealer(apiVersion string, kind string, metadata Metadata, spec SecretSealerSpec, ) *SecretSealer`

NewSecretSealer instantiates a new SecretSealer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretSealerWithDefaults

`func NewSecretSealerWithDefaults() *SecretSealer`

NewSecretSealerWithDefaults instantiates a new SecretSealer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SecretSealer) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SecretSealer) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SecretSealer) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *SecretSealer) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SecretSealer) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SecretSealer) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *SecretSealer) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecretSealer) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecretSealer) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *SecretSealer) GetSpec() SecretSealerSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SecretSealer) GetSpecOk() (*SecretSealerSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SecretSealer) SetSpec(v SecretSealerSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


