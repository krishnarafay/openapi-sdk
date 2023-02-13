# CreationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceResourceId** | Pointer to **string** | This is the ARM ID of the source object to be used to create the target object. | [optional] 

## Methods

### NewCreationData

`func NewCreationData() *CreationData`

NewCreationData instantiates a new CreationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreationDataWithDefaults

`func NewCreationDataWithDefaults() *CreationData`

NewCreationDataWithDefaults instantiates a new CreationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceResourceId

`func (o *CreationData) GetSourceResourceId() string`

GetSourceResourceId returns the SourceResourceId field if non-nil, zero value otherwise.

### GetSourceResourceIdOk

`func (o *CreationData) GetSourceResourceIdOk() (*string, bool)`

GetSourceResourceIdOk returns a tuple with the SourceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceResourceId

`func (o *CreationData) SetSourceResourceId(v string)`

SetSourceResourceId sets SourceResourceId field to given value.

### HasSourceResourceId

`func (o *CreationData) HasSourceResourceId() bool`

HasSourceResourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


