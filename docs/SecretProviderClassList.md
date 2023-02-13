# SecretProviderClassList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the provider class list resource | [optional] [readonly] 
**Items** | Pointer to [**[]SecretProviderClass**](SecretProviderClass.md) | List of the provider class resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the provider class list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewSecretProviderClassList

`func NewSecretProviderClassList() *SecretProviderClassList`

NewSecretProviderClassList instantiates a new SecretProviderClassList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretProviderClassListWithDefaults

`func NewSecretProviderClassListWithDefaults() *SecretProviderClassList`

NewSecretProviderClassListWithDefaults instantiates a new SecretProviderClassList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SecretProviderClassList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SecretProviderClassList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SecretProviderClassList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SecretProviderClassList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *SecretProviderClassList) GetItems() []SecretProviderClass`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SecretProviderClassList) GetItemsOk() (*[]SecretProviderClass, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SecretProviderClassList) SetItems(v []SecretProviderClass)`

SetItems sets Items field to given value.

### HasItems

`func (o *SecretProviderClassList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *SecretProviderClassList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SecretProviderClassList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SecretProviderClassList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SecretProviderClassList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *SecretProviderClassList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecretProviderClassList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecretProviderClassList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SecretProviderClassList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


