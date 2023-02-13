# SecretSealerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the secret sealer list resource | [optional] [readonly] 
**Items** | Pointer to [**[]SecretSealer**](SecretSealer.md) | List of the secret sealer resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the secret sealer list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewSecretSealerList

`func NewSecretSealerList() *SecretSealerList`

NewSecretSealerList instantiates a new SecretSealerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretSealerListWithDefaults

`func NewSecretSealerListWithDefaults() *SecretSealerList`

NewSecretSealerListWithDefaults instantiates a new SecretSealerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SecretSealerList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SecretSealerList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SecretSealerList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SecretSealerList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *SecretSealerList) GetItems() []SecretSealer`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SecretSealerList) GetItemsOk() (*[]SecretSealer, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SecretSealerList) SetItems(v []SecretSealer)`

SetItems sets Items field to given value.

### HasItems

`func (o *SecretSealerList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *SecretSealerList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SecretSealerList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SecretSealerList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SecretSealerList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *SecretSealerList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecretSealerList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecretSealerList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SecretSealerList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


