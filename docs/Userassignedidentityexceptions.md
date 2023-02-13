# Userassignedidentityexceptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the pod identity exception. | [optional] 
**Namespace** | Pointer to **string** | The namespace of the pod identity exception. | [optional] 
**PodLabels** | Pointer to **map[string]string** | The pod labels to match. | [optional] 

## Methods

### NewUserassignedidentityexceptions

`func NewUserassignedidentityexceptions() *Userassignedidentityexceptions`

NewUserassignedidentityexceptions instantiates a new Userassignedidentityexceptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserassignedidentityexceptionsWithDefaults

`func NewUserassignedidentityexceptionsWithDefaults() *Userassignedidentityexceptions`

NewUserassignedidentityexceptionsWithDefaults instantiates a new Userassignedidentityexceptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Userassignedidentityexceptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Userassignedidentityexceptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Userassignedidentityexceptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Userassignedidentityexceptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *Userassignedidentityexceptions) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Userassignedidentityexceptions) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Userassignedidentityexceptions) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Userassignedidentityexceptions) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPodLabels

`func (o *Userassignedidentityexceptions) GetPodLabels() map[string]string`

GetPodLabels returns the PodLabels field if non-nil, zero value otherwise.

### GetPodLabelsOk

`func (o *Userassignedidentityexceptions) GetPodLabelsOk() (*map[string]string, bool)`

GetPodLabelsOk returns a tuple with the PodLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodLabels

`func (o *Userassignedidentityexceptions) SetPodLabels(v map[string]string)`

SetPodLabels sets PodLabels field to given value.

### HasPodLabels

`func (o *Userassignedidentityexceptions) HasPodLabels() bool`

HasPodLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


