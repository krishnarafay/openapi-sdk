# ServiceMesh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | flag to specify if service mesh is enabled for blueprint | [optional] 
**Policies** | Pointer to [**[]ResourceNameAndVersionRef**](ResourceNameAndVersionRef.md) | name and version of cluster mesh policy | [optional] 
**Profile** | Pointer to [**ResourceNameAndVersionRef**](ResourceNameAndVersionRef.md) |  | [optional] 

## Methods

### NewServiceMesh

`func NewServiceMesh() *ServiceMesh`

NewServiceMesh instantiates a new ServiceMesh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMeshWithDefaults

`func NewServiceMeshWithDefaults() *ServiceMesh`

NewServiceMeshWithDefaults instantiates a new ServiceMesh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ServiceMesh) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceMesh) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceMesh) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceMesh) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPolicies

`func (o *ServiceMesh) GetPolicies() []ResourceNameAndVersionRef`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ServiceMesh) GetPoliciesOk() (*[]ResourceNameAndVersionRef, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ServiceMesh) SetPolicies(v []ResourceNameAndVersionRef)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ServiceMesh) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetProfile

`func (o *ServiceMesh) GetProfile() ResourceNameAndVersionRef`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ServiceMesh) GetProfileOk() (*ResourceNameAndVersionRef, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ServiceMesh) SetProfile(v ResourceNameAndVersionRef)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ServiceMesh) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


