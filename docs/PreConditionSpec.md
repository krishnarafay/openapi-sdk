# PreConditionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**PreConditionSpecConfig**](PreConditionSpecConfig.md) |  | 
**Type** | **string** | type of the stage precondiiton | 

## Methods

### NewPreConditionSpec

`func NewPreConditionSpec(config PreConditionSpecConfig, type_ string, ) *PreConditionSpec`

NewPreConditionSpec instantiates a new PreConditionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreConditionSpecWithDefaults

`func NewPreConditionSpecWithDefaults() *PreConditionSpec`

NewPreConditionSpecWithDefaults instantiates a new PreConditionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *PreConditionSpec) GetConfig() PreConditionSpecConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PreConditionSpec) GetConfigOk() (*PreConditionSpecConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PreConditionSpec) SetConfig(v PreConditionSpecConfig)`

SetConfig sets Config field to given value.


### GetType

`func (o *PreConditionSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PreConditionSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PreConditionSpec) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


