# AcrProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcrName** | Pointer to **string** | The name of the Azure Container Registry resource. | [optional] 
**ResourceGroupName** | Pointer to **string** | If not specified, defaults to the resource group of the managed cluster | [optional] 

## Methods

### NewAcrProfile

`func NewAcrProfile() *AcrProfile`

NewAcrProfile instantiates a new AcrProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcrProfileWithDefaults

`func NewAcrProfileWithDefaults() *AcrProfile`

NewAcrProfileWithDefaults instantiates a new AcrProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcrName

`func (o *AcrProfile) GetAcrName() string`

GetAcrName returns the AcrName field if non-nil, zero value otherwise.

### GetAcrNameOk

`func (o *AcrProfile) GetAcrNameOk() (*string, bool)`

GetAcrNameOk returns a tuple with the AcrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrName

`func (o *AcrProfile) SetAcrName(v string)`

SetAcrName sets AcrName field to given value.

### HasAcrName

`func (o *AcrProfile) HasAcrName() bool`

HasAcrName returns a boolean if a field has been set.

### GetResourceGroupName

`func (o *AcrProfile) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AcrProfile) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AcrProfile) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AcrProfile) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


