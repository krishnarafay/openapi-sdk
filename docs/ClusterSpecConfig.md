# ClusterSpecConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Raw** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Spec** | Pointer to [**AksV3Spec**](AksV3Spec.md) |  | [optional] 
**KubernetesProvider** | Pointer to **string** | kubernetes provider of the cluster | [optional] 
**Location** | Pointer to **string** | location of the cluster | [optional] 
**ProvisionEnvironment** | Pointer to **string** | provisioning environment for the cluster | [optional] 

## Methods

### NewClusterSpecConfig

`func NewClusterSpecConfig() *ClusterSpecConfig`

NewClusterSpecConfig instantiates a new ClusterSpecConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSpecConfigWithDefaults

`func NewClusterSpecConfigWithDefaults() *ClusterSpecConfig`

NewClusterSpecConfigWithDefaults instantiates a new ClusterSpecConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRaw

`func (o *ClusterSpecConfig) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *ClusterSpecConfig) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *ClusterSpecConfig) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *ClusterSpecConfig) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetApiVersion

`func (o *ClusterSpecConfig) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterSpecConfig) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterSpecConfig) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterSpecConfig) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ClusterSpecConfig) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterSpecConfig) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterSpecConfig) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterSpecConfig) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *ClusterSpecConfig) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterSpecConfig) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterSpecConfig) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ClusterSpecConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *ClusterSpecConfig) GetSpec() AksV3Spec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ClusterSpecConfig) GetSpecOk() (*AksV3Spec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ClusterSpecConfig) SetSpec(v AksV3Spec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ClusterSpecConfig) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetKubernetesProvider

`func (o *ClusterSpecConfig) GetKubernetesProvider() string`

GetKubernetesProvider returns the KubernetesProvider field if non-nil, zero value otherwise.

### GetKubernetesProviderOk

`func (o *ClusterSpecConfig) GetKubernetesProviderOk() (*string, bool)`

GetKubernetesProviderOk returns a tuple with the KubernetesProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesProvider

`func (o *ClusterSpecConfig) SetKubernetesProvider(v string)`

SetKubernetesProvider sets KubernetesProvider field to given value.

### HasKubernetesProvider

`func (o *ClusterSpecConfig) HasKubernetesProvider() bool`

HasKubernetesProvider returns a boolean if a field has been set.

### GetLocation

`func (o *ClusterSpecConfig) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ClusterSpecConfig) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ClusterSpecConfig) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ClusterSpecConfig) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetProvisionEnvironment

`func (o *ClusterSpecConfig) GetProvisionEnvironment() string`

GetProvisionEnvironment returns the ProvisionEnvironment field if non-nil, zero value otherwise.

### GetProvisionEnvironmentOk

`func (o *ClusterSpecConfig) GetProvisionEnvironmentOk() (*string, bool)`

GetProvisionEnvironmentOk returns a tuple with the ProvisionEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionEnvironment

`func (o *ClusterSpecConfig) SetProvisionEnvironment(v string)`

SetProvisionEnvironment sets ProvisionEnvironment field to given value.

### HasProvisionEnvironment

`func (o *ClusterSpecConfig) HasProvisionEnvironment() bool`

HasProvisionEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


