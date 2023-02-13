# ProjectSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterResourceQuota** | Pointer to [**ProjectResourceQuota**](ProjectResourceQuota.md) |  | [optional] 
**Default** | Pointer to **bool** | flag to indicate if this is the default project in the organization | [optional] 
**DefaultClusterNamespaceQuota** | Pointer to [**ProjectResourceQuota**](ProjectResourceQuota.md) |  | [optional] 

## Methods

### NewProjectSpec

`func NewProjectSpec() *ProjectSpec`

NewProjectSpec instantiates a new ProjectSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSpecWithDefaults

`func NewProjectSpecWithDefaults() *ProjectSpec`

NewProjectSpecWithDefaults instantiates a new ProjectSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterResourceQuota

`func (o *ProjectSpec) GetClusterResourceQuota() ProjectResourceQuota`

GetClusterResourceQuota returns the ClusterResourceQuota field if non-nil, zero value otherwise.

### GetClusterResourceQuotaOk

`func (o *ProjectSpec) GetClusterResourceQuotaOk() (*ProjectResourceQuota, bool)`

GetClusterResourceQuotaOk returns a tuple with the ClusterResourceQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterResourceQuota

`func (o *ProjectSpec) SetClusterResourceQuota(v ProjectResourceQuota)`

SetClusterResourceQuota sets ClusterResourceQuota field to given value.

### HasClusterResourceQuota

`func (o *ProjectSpec) HasClusterResourceQuota() bool`

HasClusterResourceQuota returns a boolean if a field has been set.

### GetDefault

`func (o *ProjectSpec) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ProjectSpec) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ProjectSpec) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ProjectSpec) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDefaultClusterNamespaceQuota

`func (o *ProjectSpec) GetDefaultClusterNamespaceQuota() ProjectResourceQuota`

GetDefaultClusterNamespaceQuota returns the DefaultClusterNamespaceQuota field if non-nil, zero value otherwise.

### GetDefaultClusterNamespaceQuotaOk

`func (o *ProjectSpec) GetDefaultClusterNamespaceQuotaOk() (*ProjectResourceQuota, bool)`

GetDefaultClusterNamespaceQuotaOk returns a tuple with the DefaultClusterNamespaceQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClusterNamespaceQuota

`func (o *ProjectSpec) SetDefaultClusterNamespaceQuota(v ProjectResourceQuota)`

SetDefaultClusterNamespaceQuota sets DefaultClusterNamespaceQuota field to given value.

### HasDefaultClusterNamespaceQuota

`func (o *ProjectSpec) HasDefaultClusterNamespaceQuota() bool`

HasDefaultClusterNamespaceQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


