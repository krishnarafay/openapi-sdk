# EnvironmentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependsOn** | Pointer to [**[]ResourceNameAndVersionRef**](ResourceNameAndVersionRef.md) | Specify the environment resource reference that it depends on | [optional] 
**Kind** | Pointer to **string** | Specify the environment resource kind | [optional] 
**Name** | Pointer to **string** | Specify the environment resource name | [optional] 
**ResourceOptions** | Pointer to [**EnvironmentResourceOptions**](EnvironmentResourceOptions.md) |  | [optional] 
**Type** | Pointer to **string** | Specify the environment resource type | [optional] 

## Methods

### NewEnvironmentResource

`func NewEnvironmentResource() *EnvironmentResource`

NewEnvironmentResource instantiates a new EnvironmentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentResourceWithDefaults

`func NewEnvironmentResourceWithDefaults() *EnvironmentResource`

NewEnvironmentResourceWithDefaults instantiates a new EnvironmentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependsOn

`func (o *EnvironmentResource) GetDependsOn() []ResourceNameAndVersionRef`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *EnvironmentResource) GetDependsOnOk() (*[]ResourceNameAndVersionRef, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *EnvironmentResource) SetDependsOn(v []ResourceNameAndVersionRef)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *EnvironmentResource) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### GetKind

`func (o *EnvironmentResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EnvironmentResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EnvironmentResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EnvironmentResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *EnvironmentResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceOptions

`func (o *EnvironmentResource) GetResourceOptions() EnvironmentResourceOptions`

GetResourceOptions returns the ResourceOptions field if non-nil, zero value otherwise.

### GetResourceOptionsOk

`func (o *EnvironmentResource) GetResourceOptionsOk() (*EnvironmentResourceOptions, bool)`

GetResourceOptionsOk returns a tuple with the ResourceOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOptions

`func (o *EnvironmentResource) SetResourceOptions(v EnvironmentResourceOptions)`

SetResourceOptions sets ResourceOptions field to given value.

### HasResourceOptions

`func (o *EnvironmentResource) HasResourceOptions() bool`

HasResourceOptions returns a boolean if a field has been set.

### GetType

`func (o *EnvironmentResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnvironmentResource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


