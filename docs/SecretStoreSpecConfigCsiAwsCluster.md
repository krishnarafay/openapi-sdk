# SecretStoreSpecConfigCsiAwsCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** | Cluster name for cluster linked to secret store resource | [optional] 
**Irsa** | Pointer to [**[]Irsa**](Irsa.md) | List of service account with role to be created | [optional] 
**SecretProviderClasses** | Pointer to [**[]SecretProviderClassMeta**](SecretProviderClassMeta.md) | List of service provider class to be created in cluster | [optional] 

## Methods

### NewSecretStoreSpecConfigCsiAwsCluster

`func NewSecretStoreSpecConfigCsiAwsCluster() *SecretStoreSpecConfigCsiAwsCluster`

NewSecretStoreSpecConfigCsiAwsCluster instantiates a new SecretStoreSpecConfigCsiAwsCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretStoreSpecConfigCsiAwsClusterWithDefaults

`func NewSecretStoreSpecConfigCsiAwsClusterWithDefaults() *SecretStoreSpecConfigCsiAwsCluster`

NewSecretStoreSpecConfigCsiAwsClusterWithDefaults instantiates a new SecretStoreSpecConfigCsiAwsCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *SecretStoreSpecConfigCsiAwsCluster) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *SecretStoreSpecConfigCsiAwsCluster) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *SecretStoreSpecConfigCsiAwsCluster) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *SecretStoreSpecConfigCsiAwsCluster) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetIrsa

`func (o *SecretStoreSpecConfigCsiAwsCluster) GetIrsa() []Irsa`

GetIrsa returns the Irsa field if non-nil, zero value otherwise.

### GetIrsaOk

`func (o *SecretStoreSpecConfigCsiAwsCluster) GetIrsaOk() (*[]Irsa, bool)`

GetIrsaOk returns a tuple with the Irsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrsa

`func (o *SecretStoreSpecConfigCsiAwsCluster) SetIrsa(v []Irsa)`

SetIrsa sets Irsa field to given value.

### HasIrsa

`func (o *SecretStoreSpecConfigCsiAwsCluster) HasIrsa() bool`

HasIrsa returns a boolean if a field has been set.

### GetSecretProviderClasses

`func (o *SecretStoreSpecConfigCsiAwsCluster) GetSecretProviderClasses() []SecretProviderClassMeta`

GetSecretProviderClasses returns the SecretProviderClasses field if non-nil, zero value otherwise.

### GetSecretProviderClassesOk

`func (o *SecretStoreSpecConfigCsiAwsCluster) GetSecretProviderClassesOk() (*[]SecretProviderClassMeta, bool)`

GetSecretProviderClassesOk returns a tuple with the SecretProviderClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretProviderClasses

`func (o *SecretStoreSpecConfigCsiAwsCluster) SetSecretProviderClasses(v []SecretProviderClassMeta)`

SetSecretProviderClasses sets SecretProviderClasses field to given value.

### HasSecretProviderClasses

`func (o *SecretStoreSpecConfigCsiAwsCluster) HasSecretProviderClasses() bool`

HasSecretProviderClasses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


