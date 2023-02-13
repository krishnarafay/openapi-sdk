# Additionalmetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcrProfile** | Pointer to [**AcrProfile**](AcrProfile.md) |  | [optional] 
**OmsWorkspaceLocation** | Pointer to **string** | If not specified, defaults to the resource group of the managed cluster. Valid only if the Log analytics workspace is specified. | [optional] 

## Methods

### NewAdditionalmetadata

`func NewAdditionalmetadata() *Additionalmetadata`

NewAdditionalmetadata instantiates a new Additionalmetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalmetadataWithDefaults

`func NewAdditionalmetadataWithDefaults() *Additionalmetadata`

NewAdditionalmetadataWithDefaults instantiates a new Additionalmetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcrProfile

`func (o *Additionalmetadata) GetAcrProfile() AcrProfile`

GetAcrProfile returns the AcrProfile field if non-nil, zero value otherwise.

### GetAcrProfileOk

`func (o *Additionalmetadata) GetAcrProfileOk() (*AcrProfile, bool)`

GetAcrProfileOk returns a tuple with the AcrProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrProfile

`func (o *Additionalmetadata) SetAcrProfile(v AcrProfile)`

SetAcrProfile sets AcrProfile field to given value.

### HasAcrProfile

`func (o *Additionalmetadata) HasAcrProfile() bool`

HasAcrProfile returns a boolean if a field has been set.

### GetOmsWorkspaceLocation

`func (o *Additionalmetadata) GetOmsWorkspaceLocation() string`

GetOmsWorkspaceLocation returns the OmsWorkspaceLocation field if non-nil, zero value otherwise.

### GetOmsWorkspaceLocationOk

`func (o *Additionalmetadata) GetOmsWorkspaceLocationOk() (*string, bool)`

GetOmsWorkspaceLocationOk returns a tuple with the OmsWorkspaceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmsWorkspaceLocation

`func (o *Additionalmetadata) SetOmsWorkspaceLocation(v string)`

SetOmsWorkspaceLocation sets OmsWorkspaceLocation field to given value.

### HasOmsWorkspaceLocation

`func (o *Additionalmetadata) HasOmsWorkspaceLocation() bool`

HasOmsWorkspaceLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


