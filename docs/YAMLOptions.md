# YAMLOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableOpenAPIValidation** | Pointer to **bool** | disable OpenAPI validation while deploying the YAML | [optional] 
**Force** | Pointer to **bool** | deploy YAML artifact with force flag | [optional] 

## Methods

### NewYAMLOptions

`func NewYAMLOptions() *YAMLOptions`

NewYAMLOptions instantiates a new YAMLOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYAMLOptionsWithDefaults

`func NewYAMLOptionsWithDefaults() *YAMLOptions`

NewYAMLOptionsWithDefaults instantiates a new YAMLOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableOpenAPIValidation

`func (o *YAMLOptions) GetDisableOpenAPIValidation() bool`

GetDisableOpenAPIValidation returns the DisableOpenAPIValidation field if non-nil, zero value otherwise.

### GetDisableOpenAPIValidationOk

`func (o *YAMLOptions) GetDisableOpenAPIValidationOk() (*bool, bool)`

GetDisableOpenAPIValidationOk returns a tuple with the DisableOpenAPIValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableOpenAPIValidation

`func (o *YAMLOptions) SetDisableOpenAPIValidation(v bool)`

SetDisableOpenAPIValidation sets DisableOpenAPIValidation field to given value.

### HasDisableOpenAPIValidation

`func (o *YAMLOptions) HasDisableOpenAPIValidation() bool`

HasDisableOpenAPIValidation returns a boolean if a field has been set.

### GetForce

`func (o *YAMLOptions) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *YAMLOptions) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *YAMLOptions) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *YAMLOptions) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


