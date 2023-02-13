# ConfigContextList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the config context list resource | [optional] [readonly] 
**Items** | Pointer to [**[]ConfigContext**](ConfigContext.md) | List of the config context resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the config context list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewConfigContextList

`func NewConfigContextList() *ConfigContextList`

NewConfigContextList instantiates a new ConfigContextList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigContextListWithDefaults

`func NewConfigContextListWithDefaults() *ConfigContextList`

NewConfigContextListWithDefaults instantiates a new ConfigContextList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ConfigContextList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ConfigContextList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ConfigContextList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ConfigContextList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *ConfigContextList) GetItems() []ConfigContext`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ConfigContextList) GetItemsOk() (*[]ConfigContext, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ConfigContextList) SetItems(v []ConfigContext)`

SetItems sets Items field to given value.

### HasItems

`func (o *ConfigContextList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *ConfigContextList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConfigContextList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConfigContextList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConfigContextList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *ConfigContextList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConfigContextList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConfigContextList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConfigContextList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


