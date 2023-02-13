# OverrideList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the override list resource | [optional] [default to "apps.k8smgmt.io/v3"]
**Items** | Pointer to [**[]Override**](Override.md) | List of the override resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the override list resource | [optional] [default to "Workload"]
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewOverrideList

`func NewOverrideList() *OverrideList`

NewOverrideList instantiates a new OverrideList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideListWithDefaults

`func NewOverrideListWithDefaults() *OverrideList`

NewOverrideListWithDefaults instantiates a new OverrideList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *OverrideList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *OverrideList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *OverrideList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *OverrideList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *OverrideList) GetItems() []Override`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OverrideList) GetItemsOk() (*[]Override, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OverrideList) SetItems(v []Override)`

SetItems sets Items field to given value.

### HasItems

`func (o *OverrideList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *OverrideList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OverrideList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OverrideList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *OverrideList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *OverrideList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OverrideList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OverrideList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OverrideList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


