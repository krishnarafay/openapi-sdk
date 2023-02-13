# ImportedV3ConfigObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubernetesProvider** | Pointer to **string** | kubernetes provider of the cluster | [optional] 
**Location** | Pointer to **string** | location of the cluster | [optional] 
**ProvisionEnvironment** | Pointer to **string** | provisioning environment for the cluster | [optional] 

## Methods

### NewImportedV3ConfigObject

`func NewImportedV3ConfigObject() *ImportedV3ConfigObject`

NewImportedV3ConfigObject instantiates a new ImportedV3ConfigObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportedV3ConfigObjectWithDefaults

`func NewImportedV3ConfigObjectWithDefaults() *ImportedV3ConfigObject`

NewImportedV3ConfigObjectWithDefaults instantiates a new ImportedV3ConfigObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubernetesProvider

`func (o *ImportedV3ConfigObject) GetKubernetesProvider() string`

GetKubernetesProvider returns the KubernetesProvider field if non-nil, zero value otherwise.

### GetKubernetesProviderOk

`func (o *ImportedV3ConfigObject) GetKubernetesProviderOk() (*string, bool)`

GetKubernetesProviderOk returns a tuple with the KubernetesProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesProvider

`func (o *ImportedV3ConfigObject) SetKubernetesProvider(v string)`

SetKubernetesProvider sets KubernetesProvider field to given value.

### HasKubernetesProvider

`func (o *ImportedV3ConfigObject) HasKubernetesProvider() bool`

HasKubernetesProvider returns a boolean if a field has been set.

### GetLocation

`func (o *ImportedV3ConfigObject) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ImportedV3ConfigObject) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ImportedV3ConfigObject) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ImportedV3ConfigObject) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetProvisionEnvironment

`func (o *ImportedV3ConfigObject) GetProvisionEnvironment() string`

GetProvisionEnvironment returns the ProvisionEnvironment field if non-nil, zero value otherwise.

### GetProvisionEnvironmentOk

`func (o *ImportedV3ConfigObject) GetProvisionEnvironmentOk() (*string, bool)`

GetProvisionEnvironmentOk returns a tuple with the ProvisionEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionEnvironment

`func (o *ImportedV3ConfigObject) SetProvisionEnvironment(v string)`

SetProvisionEnvironment sets ProvisionEnvironment field to given value.

### HasProvisionEnvironment

`func (o *ImportedV3ConfigObject) HasProvisionEnvironment() bool`

HasProvisionEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


