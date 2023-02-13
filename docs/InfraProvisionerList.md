# InfraProvisionerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API Version of the infrastructure provisioner list resource | [optional] [readonly] 
**Items** | Pointer to [**[]InfraProvisioner**](InfraProvisioner.md) | List of the infrastructure provisioner resources | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the infrastructure provisioner list resource | [optional] [readonly] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewInfraProvisionerList

`func NewInfraProvisionerList() *InfraProvisionerList`

NewInfraProvisionerList instantiates a new InfraProvisionerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraProvisionerListWithDefaults

`func NewInfraProvisionerListWithDefaults() *InfraProvisionerList`

NewInfraProvisionerListWithDefaults instantiates a new InfraProvisionerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *InfraProvisionerList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *InfraProvisionerList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *InfraProvisionerList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *InfraProvisionerList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *InfraProvisionerList) GetItems() []InfraProvisioner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InfraProvisionerList) GetItemsOk() (*[]InfraProvisioner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InfraProvisionerList) SetItems(v []InfraProvisioner)`

SetItems sets Items field to given value.

### HasItems

`func (o *InfraProvisionerList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *InfraProvisionerList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InfraProvisionerList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InfraProvisionerList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *InfraProvisionerList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *InfraProvisionerList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InfraProvisionerList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InfraProvisionerList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InfraProvisionerList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


