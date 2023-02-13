# NodepoolStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**StatusTime**](StatusTime.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**KubeVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeImageVersion** | Pointer to **string** |  | [optional] 
**ProvisionStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewNodepoolStatus

`func NewNodepoolStatus() *NodepoolStatus`

NewNodepoolStatus instantiates a new NodepoolStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodepoolStatusWithDefaults

`func NewNodepoolStatusWithDefaults() *NodepoolStatus`

NewNodepoolStatusWithDefaults instantiates a new NodepoolStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *NodepoolStatus) GetCreatedAt() StatusTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NodepoolStatus) GetCreatedAtOk() (*StatusTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NodepoolStatus) SetCreatedAt(v StatusTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NodepoolStatus) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *NodepoolStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodepoolStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodepoolStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NodepoolStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKubeVersion

`func (o *NodepoolStatus) GetKubeVersion() string`

GetKubeVersion returns the KubeVersion field if non-nil, zero value otherwise.

### GetKubeVersionOk

`func (o *NodepoolStatus) GetKubeVersionOk() (*string, bool)`

GetKubeVersionOk returns a tuple with the KubeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeVersion

`func (o *NodepoolStatus) SetKubeVersion(v string)`

SetKubeVersion sets KubeVersion field to given value.

### HasKubeVersion

`func (o *NodepoolStatus) HasKubeVersion() bool`

HasKubeVersion returns a boolean if a field has been set.

### GetName

`func (o *NodepoolStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodepoolStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodepoolStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodepoolStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeImageVersion

`func (o *NodepoolStatus) GetNodeImageVersion() string`

GetNodeImageVersion returns the NodeImageVersion field if non-nil, zero value otherwise.

### GetNodeImageVersionOk

`func (o *NodepoolStatus) GetNodeImageVersionOk() (*string, bool)`

GetNodeImageVersionOk returns a tuple with the NodeImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeImageVersion

`func (o *NodepoolStatus) SetNodeImageVersion(v string)`

SetNodeImageVersion sets NodeImageVersion field to given value.

### HasNodeImageVersion

`func (o *NodepoolStatus) HasNodeImageVersion() bool`

HasNodeImageVersion returns a boolean if a field has been set.

### GetProvisionStatus

`func (o *NodepoolStatus) GetProvisionStatus() string`

GetProvisionStatus returns the ProvisionStatus field if non-nil, zero value otherwise.

### GetProvisionStatusOk

`func (o *NodepoolStatus) GetProvisionStatusOk() (*string, bool)`

GetProvisionStatusOk returns a tuple with the ProvisionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionStatus

`func (o *NodepoolStatus) SetProvisionStatus(v string)`

SetProvisionStatus sets ProvisionStatus field to given value.

### HasProvisionStatus

`func (o *NodepoolStatus) HasProvisionStatus() bool`

HasProvisionStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


