# Publickeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyData** | Pointer to **string** | Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers. | [optional] 

## Methods

### NewPublickeys

`func NewPublickeys() *Publickeys`

NewPublickeys instantiates a new Publickeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublickeysWithDefaults

`func NewPublickeysWithDefaults() *Publickeys`

NewPublickeysWithDefaults instantiates a new Publickeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyData

`func (o *Publickeys) GetKeyData() string`

GetKeyData returns the KeyData field if non-nil, zero value otherwise.

### GetKeyDataOk

`func (o *Publickeys) GetKeyDataOk() (*string, bool)`

GetKeyDataOk returns a tuple with the KeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyData

`func (o *Publickeys) SetKeyData(v string)`

SetKeyData sets KeyData field to given value.

### HasKeyData

`func (o *Publickeys) HasKeyData() bool`

HasKeyData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


