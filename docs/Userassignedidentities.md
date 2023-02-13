# Userassignedidentities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindingSelector** | Pointer to **string** | The binding selector to use for the AzureIdentityBinding resource. | [optional] 
**Identity** | Pointer to [**Identity1**](Identity1.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the pod identity. | [optional] 
**Namespace** | Pointer to **string** | The namespace of the pod identity. | [optional] 

## Methods

### NewUserassignedidentities

`func NewUserassignedidentities() *Userassignedidentities`

NewUserassignedidentities instantiates a new Userassignedidentities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserassignedidentitiesWithDefaults

`func NewUserassignedidentitiesWithDefaults() *Userassignedidentities`

NewUserassignedidentitiesWithDefaults instantiates a new Userassignedidentities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingSelector

`func (o *Userassignedidentities) GetBindingSelector() string`

GetBindingSelector returns the BindingSelector field if non-nil, zero value otherwise.

### GetBindingSelectorOk

`func (o *Userassignedidentities) GetBindingSelectorOk() (*string, bool)`

GetBindingSelectorOk returns a tuple with the BindingSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingSelector

`func (o *Userassignedidentities) SetBindingSelector(v string)`

SetBindingSelector sets BindingSelector field to given value.

### HasBindingSelector

`func (o *Userassignedidentities) HasBindingSelector() bool`

HasBindingSelector returns a boolean if a field has been set.

### GetIdentity

`func (o *Userassignedidentities) GetIdentity() Identity1`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *Userassignedidentities) GetIdentityOk() (*Identity1, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *Userassignedidentities) SetIdentity(v Identity1)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *Userassignedidentities) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *Userassignedidentities) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Userassignedidentities) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Userassignedidentities) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Userassignedidentities) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *Userassignedidentities) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Userassignedidentities) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Userassignedidentities) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Userassignedidentities) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


