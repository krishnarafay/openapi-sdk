# RepositoryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCert** | Pointer to [**File**](File.md) |  | [optional] 
**EnableLFS** | Pointer to **bool** | enable git large file support | [optional] 
**EnableSubmodules** | Pointer to **bool** | enable git submodules | [optional] 
**Insecure** | Pointer to **bool** | insecure | [optional] 
**MaxRetires** | Pointer to **int32** | max retries | [optional] 

## Methods

### NewRepositoryOptions

`func NewRepositoryOptions() *RepositoryOptions`

NewRepositoryOptions instantiates a new RepositoryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryOptionsWithDefaults

`func NewRepositoryOptionsWithDefaults() *RepositoryOptions`

NewRepositoryOptionsWithDefaults instantiates a new RepositoryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCert

`func (o *RepositoryOptions) GetCaCert() File`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *RepositoryOptions) GetCaCertOk() (*File, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *RepositoryOptions) SetCaCert(v File)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *RepositoryOptions) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetEnableLFS

`func (o *RepositoryOptions) GetEnableLFS() bool`

GetEnableLFS returns the EnableLFS field if non-nil, zero value otherwise.

### GetEnableLFSOk

`func (o *RepositoryOptions) GetEnableLFSOk() (*bool, bool)`

GetEnableLFSOk returns a tuple with the EnableLFS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLFS

`func (o *RepositoryOptions) SetEnableLFS(v bool)`

SetEnableLFS sets EnableLFS field to given value.

### HasEnableLFS

`func (o *RepositoryOptions) HasEnableLFS() bool`

HasEnableLFS returns a boolean if a field has been set.

### GetEnableSubmodules

`func (o *RepositoryOptions) GetEnableSubmodules() bool`

GetEnableSubmodules returns the EnableSubmodules field if non-nil, zero value otherwise.

### GetEnableSubmodulesOk

`func (o *RepositoryOptions) GetEnableSubmodulesOk() (*bool, bool)`

GetEnableSubmodulesOk returns a tuple with the EnableSubmodules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubmodules

`func (o *RepositoryOptions) SetEnableSubmodules(v bool)`

SetEnableSubmodules sets EnableSubmodules field to given value.

### HasEnableSubmodules

`func (o *RepositoryOptions) HasEnableSubmodules() bool`

HasEnableSubmodules returns a boolean if a field has been set.

### GetInsecure

`func (o *RepositoryOptions) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *RepositoryOptions) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *RepositoryOptions) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *RepositoryOptions) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetMaxRetires

`func (o *RepositoryOptions) GetMaxRetires() int32`

GetMaxRetires returns the MaxRetires field if non-nil, zero value otherwise.

### GetMaxRetiresOk

`func (o *RepositoryOptions) GetMaxRetiresOk() (*int32, bool)`

GetMaxRetiresOk returns a tuple with the MaxRetires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetires

`func (o *RepositoryOptions) SetMaxRetires(v int32)`

SetMaxRetires sets MaxRetires field to given value.

### HasMaxRetires

`func (o *RepositoryOptions) HasMaxRetires() bool`

HasMaxRetires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


