# IdentityProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIdentity** | Pointer to [**ClusterIdentityStatus**](ClusterIdentityStatus.md) |  | [optional] 
**KubeletIdentity** | Pointer to [**KubeletIdentityStatus**](KubeletIdentityStatus.md) |  | [optional] 

## Methods

### NewIdentityProfile

`func NewIdentityProfile() *IdentityProfile`

NewIdentityProfile instantiates a new IdentityProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProfileWithDefaults

`func NewIdentityProfileWithDefaults() *IdentityProfile`

NewIdentityProfileWithDefaults instantiates a new IdentityProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIdentity

`func (o *IdentityProfile) GetClusterIdentity() ClusterIdentityStatus`

GetClusterIdentity returns the ClusterIdentity field if non-nil, zero value otherwise.

### GetClusterIdentityOk

`func (o *IdentityProfile) GetClusterIdentityOk() (*ClusterIdentityStatus, bool)`

GetClusterIdentityOk returns a tuple with the ClusterIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentity

`func (o *IdentityProfile) SetClusterIdentity(v ClusterIdentityStatus)`

SetClusterIdentity sets ClusterIdentity field to given value.

### HasClusterIdentity

`func (o *IdentityProfile) HasClusterIdentity() bool`

HasClusterIdentity returns a boolean if a field has been set.

### GetKubeletIdentity

`func (o *IdentityProfile) GetKubeletIdentity() KubeletIdentityStatus`

GetKubeletIdentity returns the KubeletIdentity field if non-nil, zero value otherwise.

### GetKubeletIdentityOk

`func (o *IdentityProfile) GetKubeletIdentityOk() (*KubeletIdentityStatus, bool)`

GetKubeletIdentityOk returns a tuple with the KubeletIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletIdentity

`func (o *IdentityProfile) SetKubeletIdentity(v KubeletIdentityStatus)`

SetKubeletIdentity sets KubeletIdentity field to given value.

### HasKubeletIdentity

`func (o *IdentityProfile) HasKubeletIdentity() bool`

HasKubeletIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


