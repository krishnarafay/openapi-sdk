# CatalogSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSync** | Pointer to **bool** | flag to indicate if the catalog is synced periodically | [optional] 
**IconURL** | Pointer to **string** | icon url of catalog | [optional] 
**Repository** | **string** | catalog helm repository name | 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Type** | **string** | type of catalog | 

## Methods

### NewCatalogSpec

`func NewCatalogSpec(repository string, type_ string, ) *CatalogSpec`

NewCatalogSpec instantiates a new CatalogSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogSpecWithDefaults

`func NewCatalogSpecWithDefaults() *CatalogSpec`

NewCatalogSpecWithDefaults instantiates a new CatalogSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSync

`func (o *CatalogSpec) GetAutoSync() bool`

GetAutoSync returns the AutoSync field if non-nil, zero value otherwise.

### GetAutoSyncOk

`func (o *CatalogSpec) GetAutoSyncOk() (*bool, bool)`

GetAutoSyncOk returns a tuple with the AutoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSync

`func (o *CatalogSpec) SetAutoSync(v bool)`

SetAutoSync sets AutoSync field to given value.

### HasAutoSync

`func (o *CatalogSpec) HasAutoSync() bool`

HasAutoSync returns a boolean if a field has been set.

### GetIconURL

`func (o *CatalogSpec) GetIconURL() string`

GetIconURL returns the IconURL field if non-nil, zero value otherwise.

### GetIconURLOk

`func (o *CatalogSpec) GetIconURLOk() (*string, bool)`

GetIconURLOk returns a tuple with the IconURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconURL

`func (o *CatalogSpec) SetIconURL(v string)`

SetIconURL sets IconURL field to given value.

### HasIconURL

`func (o *CatalogSpec) HasIconURL() bool`

HasIconURL returns a boolean if a field has been set.

### GetRepository

`func (o *CatalogSpec) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CatalogSpec) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CatalogSpec) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetSharing

`func (o *CatalogSpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *CatalogSpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *CatalogSpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *CatalogSpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetType

`func (o *CatalogSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogSpec) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


