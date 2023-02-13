# ControlPlaneStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to [**EndpointCp**](EndpointCp.md) |  | [optional] 
**KubeVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewControlPlaneStatus

`func NewControlPlaneStatus() *ControlPlaneStatus`

NewControlPlaneStatus instantiates a new ControlPlaneStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlPlaneStatusWithDefaults

`func NewControlPlaneStatusWithDefaults() *ControlPlaneStatus`

NewControlPlaneStatusWithDefaults instantiates a new ControlPlaneStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *ControlPlaneStatus) GetEndpoint() EndpointCp`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ControlPlaneStatus) GetEndpointOk() (*EndpointCp, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ControlPlaneStatus) SetEndpoint(v EndpointCp)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ControlPlaneStatus) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetKubeVersion

`func (o *ControlPlaneStatus) GetKubeVersion() string`

GetKubeVersion returns the KubeVersion field if non-nil, zero value otherwise.

### GetKubeVersionOk

`func (o *ControlPlaneStatus) GetKubeVersionOk() (*string, bool)`

GetKubeVersionOk returns a tuple with the KubeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeVersion

`func (o *ControlPlaneStatus) SetKubeVersion(v string)`

SetKubeVersion sets KubeVersion field to given value.

### HasKubeVersion

`func (o *ControlPlaneStatus) HasKubeVersion() bool`

HasKubeVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


