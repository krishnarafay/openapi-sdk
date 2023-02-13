# Managedcluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalMetadata** | Pointer to [**Additionalmetadata**](Additionalmetadata.md) |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**ExtendedLocation** | Pointer to [**Extendedlocation**](Extendedlocation.md) |  | [optional] 
**Identity** | Pointer to [**Identity**](Identity.md) |  | [optional] 
**Location** | Pointer to **string** | The geo-location where the resource lives | [optional] 
**Properties** | Pointer to [**ManagedClusterProperties**](ManagedClusterProperties.md) |  | [optional] 
**Sku** | Pointer to [**Sku**](Sku.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** | Resource tags. | [optional] 
**Type** | Pointer to **string** | Azure Resource Type | [optional] [default to "Microsoft.ContainerService/managedClusters"]

## Methods

### NewManagedcluster

`func NewManagedcluster() *Managedcluster`

NewManagedcluster instantiates a new Managedcluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedclusterWithDefaults

`func NewManagedclusterWithDefaults() *Managedcluster`

NewManagedclusterWithDefaults instantiates a new Managedcluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalMetadata

`func (o *Managedcluster) GetAdditionalMetadata() Additionalmetadata`

GetAdditionalMetadata returns the AdditionalMetadata field if non-nil, zero value otherwise.

### GetAdditionalMetadataOk

`func (o *Managedcluster) GetAdditionalMetadataOk() (*Additionalmetadata, bool)`

GetAdditionalMetadataOk returns a tuple with the AdditionalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMetadata

`func (o *Managedcluster) SetAdditionalMetadata(v Additionalmetadata)`

SetAdditionalMetadata sets AdditionalMetadata field to given value.

### HasAdditionalMetadata

`func (o *Managedcluster) HasAdditionalMetadata() bool`

HasAdditionalMetadata returns a boolean if a field has been set.

### GetApiVersion

`func (o *Managedcluster) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Managedcluster) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Managedcluster) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *Managedcluster) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetExtendedLocation

`func (o *Managedcluster) GetExtendedLocation() Extendedlocation`

GetExtendedLocation returns the ExtendedLocation field if non-nil, zero value otherwise.

### GetExtendedLocationOk

`func (o *Managedcluster) GetExtendedLocationOk() (*Extendedlocation, bool)`

GetExtendedLocationOk returns a tuple with the ExtendedLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedLocation

`func (o *Managedcluster) SetExtendedLocation(v Extendedlocation)`

SetExtendedLocation sets ExtendedLocation field to given value.

### HasExtendedLocation

`func (o *Managedcluster) HasExtendedLocation() bool`

HasExtendedLocation returns a boolean if a field has been set.

### GetIdentity

`func (o *Managedcluster) GetIdentity() Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *Managedcluster) GetIdentityOk() (*Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *Managedcluster) SetIdentity(v Identity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *Managedcluster) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetLocation

`func (o *Managedcluster) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Managedcluster) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Managedcluster) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Managedcluster) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetProperties

`func (o *Managedcluster) GetProperties() ManagedClusterProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Managedcluster) GetPropertiesOk() (*ManagedClusterProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Managedcluster) SetProperties(v ManagedClusterProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Managedcluster) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSku

`func (o *Managedcluster) GetSku() Sku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Managedcluster) GetSkuOk() (*Sku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Managedcluster) SetSku(v Sku)`

SetSku sets Sku field to given value.

### HasSku

`func (o *Managedcluster) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetTags

`func (o *Managedcluster) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Managedcluster) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Managedcluster) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Managedcluster) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *Managedcluster) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Managedcluster) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Managedcluster) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Managedcluster) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


