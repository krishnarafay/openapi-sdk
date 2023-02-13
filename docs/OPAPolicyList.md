# OPAPolicyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the policy list resource | [optional] [readonly] 
**Items** | Pointer to [**[]OPAPolicy**](OPAPolicy.md) | List of the policy resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the policy list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewOPAPolicyList

`func NewOPAPolicyList() *OPAPolicyList`

NewOPAPolicyList instantiates a new OPAPolicyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOPAPolicyListWithDefaults

`func NewOPAPolicyListWithDefaults() *OPAPolicyList`

NewOPAPolicyListWithDefaults instantiates a new OPAPolicyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *OPAPolicyList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *OPAPolicyList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *OPAPolicyList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *OPAPolicyList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *OPAPolicyList) GetItems() []OPAPolicy`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OPAPolicyList) GetItemsOk() (*[]OPAPolicy, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OPAPolicyList) SetItems(v []OPAPolicy)`

SetItems sets Items field to given value.

### HasItems

`func (o *OPAPolicyList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *OPAPolicyList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OPAPolicyList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OPAPolicyList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *OPAPolicyList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *OPAPolicyList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OPAPolicyList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OPAPolicyList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OPAPolicyList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


