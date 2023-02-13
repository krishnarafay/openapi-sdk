# Extra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ExtraData**](ExtraData.md) |  | [optional] 
**Files** | Pointer to [**[]File**](File.md) |  | [optional] 

## Methods

### NewExtra

`func NewExtra() *Extra`

NewExtra instantiates a new Extra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtraWithDefaults

`func NewExtraWithDefaults() *Extra`

NewExtraWithDefaults instantiates a new Extra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Extra) GetData() ExtraData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Extra) GetDataOk() (*ExtraData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Extra) SetData(v ExtraData)`

SetData sets Data field to given value.

### HasData

`func (o *Extra) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFiles

`func (o *Extra) GetFiles() []File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Extra) GetFilesOk() (*[]File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Extra) SetFiles(v []File)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Extra) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


