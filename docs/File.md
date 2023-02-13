# File

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | data is the base64 encoded contents of the file | [optional] 
**Name** | **string** | relative path of a artifact | 
**Sensitive** | Pointer to **bool** | data is the base64 encoded contents of the file if set to true | [optional] 

## Methods

### NewFile

`func NewFile(name string, ) *File`

NewFile instantiates a new File object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileWithDefaults

`func NewFileWithDefaults() *File`

NewFileWithDefaults instantiates a new File object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *File) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *File) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *File) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *File) HasData() bool`

HasData returns a boolean if a field has been set.

### GetName

`func (o *File) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *File) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *File) SetName(v string)`

SetName sets Name field to given value.


### GetSensitive

`func (o *File) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *File) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *File) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *File) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


