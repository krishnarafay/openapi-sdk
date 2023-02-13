# APIError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | Error Code | [optional] 
**External** | Pointer to **string** | External Error | [optional] 
**Internal** | Pointer to **string** | Internal Error string | [optional] 

## Methods

### NewAPIError

`func NewAPIError() *APIError`

NewAPIError instantiates a new APIError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIErrorWithDefaults

`func NewAPIErrorWithDefaults() *APIError`

NewAPIErrorWithDefaults instantiates a new APIError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *APIError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *APIError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *APIError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *APIError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetExternal

`func (o *APIError) GetExternal() string`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *APIError) GetExternalOk() (*string, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *APIError) SetExternal(v string)`

SetExternal sets External field to given value.

### HasExternal

`func (o *APIError) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetInternal

`func (o *APIError) GetInternal() string`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *APIError) GetInternalOk() (*string, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *APIError) SetInternal(v string)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *APIError) HasInternal() bool`

HasInternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


