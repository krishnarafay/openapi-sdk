# KubeletIdentityStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 

## Methods

### NewKubeletIdentityStatus

`func NewKubeletIdentityStatus() *KubeletIdentityStatus`

NewKubeletIdentityStatus instantiates a new KubeletIdentityStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeletIdentityStatusWithDefaults

`func NewKubeletIdentityStatusWithDefaults() *KubeletIdentityStatus`

NewKubeletIdentityStatusWithDefaults instantiates a new KubeletIdentityStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalId

`func (o *KubeletIdentityStatus) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *KubeletIdentityStatus) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *KubeletIdentityStatus) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.

### HasPrincipalId

`func (o *KubeletIdentityStatus) HasPrincipalId() bool`

HasPrincipalId returns a boolean if a field has been set.

### GetTenantId

`func (o *KubeletIdentityStatus) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *KubeletIdentityStatus) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *KubeletIdentityStatus) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *KubeletIdentityStatus) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


