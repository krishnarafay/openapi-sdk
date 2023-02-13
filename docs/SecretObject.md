# SecretObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Data** | Pointer to [**[]SecretObjectData**](SecretObjectData.md) |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**SecretName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSecretObject

`func NewSecretObject() *SecretObject`

NewSecretObject instantiates a new SecretObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretObjectWithDefaults

`func NewSecretObjectWithDefaults() *SecretObject`

NewSecretObjectWithDefaults instantiates a new SecretObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *SecretObject) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *SecretObject) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *SecretObject) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *SecretObject) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetData

`func (o *SecretObject) GetData() []SecretObjectData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SecretObject) GetDataOk() (*[]SecretObjectData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SecretObject) SetData(v []SecretObjectData)`

SetData sets Data field to given value.

### HasData

`func (o *SecretObject) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLabels

`func (o *SecretObject) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SecretObject) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SecretObject) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SecretObject) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetSecretName

`func (o *SecretObject) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *SecretObject) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *SecretObject) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.

### HasSecretName

`func (o *SecretObject) HasSecretName() bool`

HasSecretName returns a boolean if a field has been set.

### GetType

`func (o *SecretObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecretObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecretObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SecretObject) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


