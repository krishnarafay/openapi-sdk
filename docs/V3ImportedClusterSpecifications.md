# V3ImportedClusterSpecifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubernetesProvider** | Pointer to **string** | kubernetes provider of the cluster | [optional] 
**Location** | Pointer to **string** | location of the cluster | [optional] 
**ProvisionEnvironment** | Pointer to **string** | provisioning environment for the cluster | [optional] 

## Methods

### NewV3ImportedClusterSpecifications

`func NewV3ImportedClusterSpecifications() *V3ImportedClusterSpecifications`

NewV3ImportedClusterSpecifications instantiates a new V3ImportedClusterSpecifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ImportedClusterSpecificationsWithDefaults

`func NewV3ImportedClusterSpecificationsWithDefaults() *V3ImportedClusterSpecifications`

NewV3ImportedClusterSpecificationsWithDefaults instantiates a new V3ImportedClusterSpecifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubernetesProvider

`func (o *V3ImportedClusterSpecifications) GetKubernetesProvider() string`

GetKubernetesProvider returns the KubernetesProvider field if non-nil, zero value otherwise.

### GetKubernetesProviderOk

`func (o *V3ImportedClusterSpecifications) GetKubernetesProviderOk() (*string, bool)`

GetKubernetesProviderOk returns a tuple with the KubernetesProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesProvider

`func (o *V3ImportedClusterSpecifications) SetKubernetesProvider(v string)`

SetKubernetesProvider sets KubernetesProvider field to given value.

### HasKubernetesProvider

`func (o *V3ImportedClusterSpecifications) HasKubernetesProvider() bool`

HasKubernetesProvider returns a boolean if a field has been set.

### GetLocation

`func (o *V3ImportedClusterSpecifications) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *V3ImportedClusterSpecifications) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *V3ImportedClusterSpecifications) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *V3ImportedClusterSpecifications) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetProvisionEnvironment

`func (o *V3ImportedClusterSpecifications) GetProvisionEnvironment() string`

GetProvisionEnvironment returns the ProvisionEnvironment field if non-nil, zero value otherwise.

### GetProvisionEnvironmentOk

`func (o *V3ImportedClusterSpecifications) GetProvisionEnvironmentOk() (*string, bool)`

GetProvisionEnvironmentOk returns a tuple with the ProvisionEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionEnvironment

`func (o *V3ImportedClusterSpecifications) SetProvisionEnvironment(v string)`

SetProvisionEnvironment sets ProvisionEnvironment field to given value.

### HasProvisionEnvironment

`func (o *V3ImportedClusterSpecifications) HasProvisionEnvironment() bool`

HasProvisionEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


