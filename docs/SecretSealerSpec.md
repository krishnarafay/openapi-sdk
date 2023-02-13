# SecretSealerSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Type** | Pointer to **string** | secret sealer type | [optional] 
**Version** | Pointer to **string** | secret sealer sharing configuration | [optional] 

## Methods

### NewSecretSealerSpec

`func NewSecretSealerSpec() *SecretSealerSpec`

NewSecretSealerSpec instantiates a new SecretSealerSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretSealerSpecWithDefaults

`func NewSecretSealerSpecWithDefaults() *SecretSealerSpec`

NewSecretSealerSpecWithDefaults instantiates a new SecretSealerSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharing

`func (o *SecretSealerSpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *SecretSealerSpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *SecretSealerSpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *SecretSealerSpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetType

`func (o *SecretSealerSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecretSealerSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecretSealerSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SecretSealerSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *SecretSealerSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SecretSealerSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SecretSealerSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SecretSealerSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


