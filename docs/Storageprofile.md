# Storageprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskCSIDriver** | Pointer to [**Diskcsidriver**](Diskcsidriver.md) |  | [optional] 
**FileCSIDriver** | Pointer to [**Filecsidriver**](Filecsidriver.md) |  | [optional] 
**SnapshotController** | Pointer to [**Snapshotcontroller**](Snapshotcontroller.md) |  | [optional] 

## Methods

### NewStorageprofile

`func NewStorageprofile() *Storageprofile`

NewStorageprofile instantiates a new Storageprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageprofileWithDefaults

`func NewStorageprofileWithDefaults() *Storageprofile`

NewStorageprofileWithDefaults instantiates a new Storageprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskCSIDriver

`func (o *Storageprofile) GetDiskCSIDriver() Diskcsidriver`

GetDiskCSIDriver returns the DiskCSIDriver field if non-nil, zero value otherwise.

### GetDiskCSIDriverOk

`func (o *Storageprofile) GetDiskCSIDriverOk() (*Diskcsidriver, bool)`

GetDiskCSIDriverOk returns a tuple with the DiskCSIDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCSIDriver

`func (o *Storageprofile) SetDiskCSIDriver(v Diskcsidriver)`

SetDiskCSIDriver sets DiskCSIDriver field to given value.

### HasDiskCSIDriver

`func (o *Storageprofile) HasDiskCSIDriver() bool`

HasDiskCSIDriver returns a boolean if a field has been set.

### GetFileCSIDriver

`func (o *Storageprofile) GetFileCSIDriver() Filecsidriver`

GetFileCSIDriver returns the FileCSIDriver field if non-nil, zero value otherwise.

### GetFileCSIDriverOk

`func (o *Storageprofile) GetFileCSIDriverOk() (*Filecsidriver, bool)`

GetFileCSIDriverOk returns a tuple with the FileCSIDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCSIDriver

`func (o *Storageprofile) SetFileCSIDriver(v Filecsidriver)`

SetFileCSIDriver sets FileCSIDriver field to given value.

### HasFileCSIDriver

`func (o *Storageprofile) HasFileCSIDriver() bool`

HasFileCSIDriver returns a boolean if a field has been set.

### GetSnapshotController

`func (o *Storageprofile) GetSnapshotController() Snapshotcontroller`

GetSnapshotController returns the SnapshotController field if non-nil, zero value otherwise.

### GetSnapshotControllerOk

`func (o *Storageprofile) GetSnapshotControllerOk() (*Snapshotcontroller, bool)`

GetSnapshotControllerOk returns a tuple with the SnapshotController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotController

`func (o *Storageprofile) SetSnapshotController(v Snapshotcontroller)`

SetSnapshotController sets SnapshotController field to given value.

### HasSnapshotController

`func (o *Storageprofile) HasSnapshotController() bool`

HasSnapshotController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


