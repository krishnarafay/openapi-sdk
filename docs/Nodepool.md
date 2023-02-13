# Nodepool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | Azure Api Version | [optional] 
**Location** | Pointer to **string** | The geo-location where the resource lives | [optional] 
**Name** | Pointer to **string** | The name of the agent pool. | [optional] 
**Properties** | Pointer to [**NodePoolProperties**](NodePoolProperties.md) |  | [optional] 
**Type** | Pointer to **string** | Nodepool azure resource type | [optional] [default to "Microsoft.ContainerService/managedClusters/agentPools"]

## Methods

### NewNodepool

`func NewNodepool() *Nodepool`

NewNodepool instantiates a new Nodepool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodepoolWithDefaults

`func NewNodepoolWithDefaults() *Nodepool`

NewNodepoolWithDefaults instantiates a new Nodepool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *Nodepool) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Nodepool) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Nodepool) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *Nodepool) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetLocation

`func (o *Nodepool) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Nodepool) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Nodepool) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Nodepool) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *Nodepool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Nodepool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Nodepool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Nodepool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *Nodepool) GetProperties() NodePoolProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Nodepool) GetPropertiesOk() (*NodePoolProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Nodepool) SetProperties(v NodePoolProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Nodepool) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetType

`func (o *Nodepool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Nodepool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Nodepool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Nodepool) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


