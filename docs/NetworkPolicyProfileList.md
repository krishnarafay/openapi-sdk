# NetworkPolicyProfileList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the policy list resource | [optional] [readonly] 
**Items** | Pointer to [**[]NetworkPolicyProfile**](NetworkPolicyProfile.md) | List of the profile resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the policy list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewNetworkPolicyProfileList

`func NewNetworkPolicyProfileList() *NetworkPolicyProfileList`

NewNetworkPolicyProfileList instantiates a new NetworkPolicyProfileList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPolicyProfileListWithDefaults

`func NewNetworkPolicyProfileListWithDefaults() *NetworkPolicyProfileList`

NewNetworkPolicyProfileListWithDefaults instantiates a new NetworkPolicyProfileList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkPolicyProfileList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkPolicyProfileList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkPolicyProfileList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkPolicyProfileList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *NetworkPolicyProfileList) GetItems() []NetworkPolicyProfile`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NetworkPolicyProfileList) GetItemsOk() (*[]NetworkPolicyProfile, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NetworkPolicyProfileList) SetItems(v []NetworkPolicyProfile)`

SetItems sets Items field to given value.

### HasItems

`func (o *NetworkPolicyProfileList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *NetworkPolicyProfileList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkPolicyProfileList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkPolicyProfileList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkPolicyProfileList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkPolicyProfileList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkPolicyProfileList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkPolicyProfileList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NetworkPolicyProfileList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


