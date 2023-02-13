# SecretStoreList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the secret store list resource | [optional] [readonly] 
**Items** | Pointer to [**[]SecretStore**](SecretStore.md) | List of the secret store resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the secret store list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewSecretStoreList

`func NewSecretStoreList() *SecretStoreList`

NewSecretStoreList instantiates a new SecretStoreList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretStoreListWithDefaults

`func NewSecretStoreListWithDefaults() *SecretStoreList`

NewSecretStoreListWithDefaults instantiates a new SecretStoreList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SecretStoreList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SecretStoreList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SecretStoreList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SecretStoreList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *SecretStoreList) GetItems() []SecretStore`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SecretStoreList) GetItemsOk() (*[]SecretStore, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SecretStoreList) SetItems(v []SecretStore)`

SetItems sets Items field to given value.

### HasItems

`func (o *SecretStoreList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *SecretStoreList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SecretStoreList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SecretStoreList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SecretStoreList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *SecretStoreList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecretStoreList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecretStoreList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SecretStoreList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


