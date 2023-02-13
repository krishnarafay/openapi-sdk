# AwsCostProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsCredentials** | Pointer to [**AwsCredsCostProfile**](AwsCredsCostProfile.md) |  | [optional] 
**CurIntegration** | Pointer to [**AwsCurIntegration**](AwsCurIntegration.md) |  | [optional] 
**SpotIntegration** | Pointer to [**AwsSpotIntegration**](AwsSpotIntegration.md) |  | [optional] 

## Methods

### NewAwsCostProfile

`func NewAwsCostProfile() *AwsCostProfile`

NewAwsCostProfile instantiates a new AwsCostProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsCostProfileWithDefaults

`func NewAwsCostProfileWithDefaults() *AwsCostProfile`

NewAwsCostProfileWithDefaults instantiates a new AwsCostProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsCredentials

`func (o *AwsCostProfile) GetAwsCredentials() AwsCredsCostProfile`

GetAwsCredentials returns the AwsCredentials field if non-nil, zero value otherwise.

### GetAwsCredentialsOk

`func (o *AwsCostProfile) GetAwsCredentialsOk() (*AwsCredsCostProfile, bool)`

GetAwsCredentialsOk returns a tuple with the AwsCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCredentials

`func (o *AwsCostProfile) SetAwsCredentials(v AwsCredsCostProfile)`

SetAwsCredentials sets AwsCredentials field to given value.

### HasAwsCredentials

`func (o *AwsCostProfile) HasAwsCredentials() bool`

HasAwsCredentials returns a boolean if a field has been set.

### GetCurIntegration

`func (o *AwsCostProfile) GetCurIntegration() AwsCurIntegration`

GetCurIntegration returns the CurIntegration field if non-nil, zero value otherwise.

### GetCurIntegrationOk

`func (o *AwsCostProfile) GetCurIntegrationOk() (*AwsCurIntegration, bool)`

GetCurIntegrationOk returns a tuple with the CurIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurIntegration

`func (o *AwsCostProfile) SetCurIntegration(v AwsCurIntegration)`

SetCurIntegration sets CurIntegration field to given value.

### HasCurIntegration

`func (o *AwsCostProfile) HasCurIntegration() bool`

HasCurIntegration returns a boolean if a field has been set.

### GetSpotIntegration

`func (o *AwsCostProfile) GetSpotIntegration() AwsSpotIntegration`

GetSpotIntegration returns the SpotIntegration field if non-nil, zero value otherwise.

### GetSpotIntegrationOk

`func (o *AwsCostProfile) GetSpotIntegrationOk() (*AwsSpotIntegration, bool)`

GetSpotIntegrationOk returns a tuple with the SpotIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotIntegration

`func (o *AwsCostProfile) SetSpotIntegration(v AwsSpotIntegration)`

SetSpotIntegration sets SpotIntegration field to given value.

### HasSpotIntegration

`func (o *AwsCostProfile) HasSpotIntegration() bool`

HasSpotIntegration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


