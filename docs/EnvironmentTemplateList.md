# EnvironmentTemplateList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the environment template list resource | [optional] [readonly] 
**Items** | Pointer to [**[]EnvironmentTemplate**](EnvironmentTemplate.md) | List of the environment template resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the environment template list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewEnvironmentTemplateList

`func NewEnvironmentTemplateList() *EnvironmentTemplateList`

NewEnvironmentTemplateList instantiates a new EnvironmentTemplateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentTemplateListWithDefaults

`func NewEnvironmentTemplateListWithDefaults() *EnvironmentTemplateList`

NewEnvironmentTemplateListWithDefaults instantiates a new EnvironmentTemplateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *EnvironmentTemplateList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *EnvironmentTemplateList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *EnvironmentTemplateList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *EnvironmentTemplateList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *EnvironmentTemplateList) GetItems() []EnvironmentTemplate`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EnvironmentTemplateList) GetItemsOk() (*[]EnvironmentTemplate, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EnvironmentTemplateList) SetItems(v []EnvironmentTemplate)`

SetItems sets Items field to given value.

### HasItems

`func (o *EnvironmentTemplateList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *EnvironmentTemplateList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EnvironmentTemplateList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EnvironmentTemplateList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EnvironmentTemplateList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *EnvironmentTemplateList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnvironmentTemplateList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnvironmentTemplateList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EnvironmentTemplateList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


