# SecretStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the secret store resource | [optional] [default to "integrations.k8smgmt.io/v3"]
**Kind** | Pointer to **string** | Kind of the secret store resource | [optional] [default to "SecretStore"]
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Spec** | Pointer to [**SecretStoreSpec**](SecretStoreSpec.md) |  | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewSecretStore

`func NewSecretStore() *SecretStore`

NewSecretStore instantiates a new SecretStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretStoreWithDefaults

`func NewSecretStoreWithDefaults() *SecretStore`

NewSecretStoreWithDefaults instantiates a new SecretStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SecretStore) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SecretStore) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SecretStore) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SecretStore) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SecretStore) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SecretStore) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SecretStore) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SecretStore) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *SecretStore) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecretStore) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecretStore) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SecretStore) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *SecretStore) GetSpec() SecretStoreSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SecretStore) GetSpecOk() (*SecretStoreSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SecretStore) SetSpec(v SecretStoreSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SecretStore) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *SecretStore) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecretStore) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecretStore) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecretStore) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


