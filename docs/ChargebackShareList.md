# ChargebackShareList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the ChargebackShare list resource | [optional] [readonly] 
**Items** | Pointer to [**[]ChargebackShare**](ChargebackShare.md) | List of the ChargebackShare resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the ChargebackShare list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewChargebackShareList

`func NewChargebackShareList() *ChargebackShareList`

NewChargebackShareList instantiates a new ChargebackShareList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargebackShareListWithDefaults

`func NewChargebackShareListWithDefaults() *ChargebackShareList`

NewChargebackShareListWithDefaults instantiates a new ChargebackShareList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ChargebackShareList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ChargebackShareList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ChargebackShareList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ChargebackShareList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *ChargebackShareList) GetItems() []ChargebackShare`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ChargebackShareList) GetItemsOk() (*[]ChargebackShare, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ChargebackShareList) SetItems(v []ChargebackShare)`

SetItems sets Items field to given value.

### HasItems

`func (o *ChargebackShareList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *ChargebackShareList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ChargebackShareList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ChargebackShareList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ChargebackShareList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *ChargebackShareList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ChargebackShareList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ChargebackShareList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ChargebackShareList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


