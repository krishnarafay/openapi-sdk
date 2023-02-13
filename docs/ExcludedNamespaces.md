# ExcludedNamespaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespaces** | [**[]ResourceRef**](ResourceRef.md) | List of namespaces to be excluded | 
**Processes** | **[]string** | List of processes to be excluding namespace from | 

## Methods

### NewExcludedNamespaces

`func NewExcludedNamespaces(namespaces []ResourceRef, processes []string, ) *ExcludedNamespaces`

NewExcludedNamespaces instantiates a new ExcludedNamespaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExcludedNamespacesWithDefaults

`func NewExcludedNamespacesWithDefaults() *ExcludedNamespaces`

NewExcludedNamespacesWithDefaults instantiates a new ExcludedNamespaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespaces

`func (o *ExcludedNamespaces) GetNamespaces() []ResourceRef`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *ExcludedNamespaces) GetNamespacesOk() (*[]ResourceRef, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *ExcludedNamespaces) SetNamespaces(v []ResourceRef)`

SetNamespaces sets Namespaces field to given value.


### GetProcesses

`func (o *ExcludedNamespaces) GetProcesses() []string`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *ExcludedNamespaces) GetProcessesOk() (*[]string, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *ExcludedNamespaces) SetProcesses(v []string)`

SetProcesses sets Processes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


