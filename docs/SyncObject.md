# SyncObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **string** | group of k8s object | 
**Kind** | **string** | kind of k8s object | 
**Version** | **string** | version of k8s object | 

## Methods

### NewSyncObject

`func NewSyncObject(group string, kind string, version string, ) *SyncObject`

NewSyncObject instantiates a new SyncObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncObjectWithDefaults

`func NewSyncObjectWithDefaults() *SyncObject`

NewSyncObjectWithDefaults instantiates a new SyncObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *SyncObject) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SyncObject) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SyncObject) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetKind

`func (o *SyncObject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SyncObject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SyncObject) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVersion

`func (o *SyncObject) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SyncObject) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SyncObject) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


