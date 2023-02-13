# OPAPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConstraintList** | Pointer to [**[]OPAPolicyConstraint**](OPAPolicyConstraint.md) | version of the policy | [optional] 
**Sharing** | Pointer to [**SharingSpec**](SharingSpec.md) |  | [optional] 
**Version** | Pointer to **string** | version of the policy | [optional] 

## Methods

### NewOPAPolicySpec

`func NewOPAPolicySpec() *OPAPolicySpec`

NewOPAPolicySpec instantiates a new OPAPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOPAPolicySpecWithDefaults

`func NewOPAPolicySpecWithDefaults() *OPAPolicySpec`

NewOPAPolicySpecWithDefaults instantiates a new OPAPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraintList

`func (o *OPAPolicySpec) GetConstraintList() []OPAPolicyConstraint`

GetConstraintList returns the ConstraintList field if non-nil, zero value otherwise.

### GetConstraintListOk

`func (o *OPAPolicySpec) GetConstraintListOk() (*[]OPAPolicyConstraint, bool)`

GetConstraintListOk returns a tuple with the ConstraintList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintList

`func (o *OPAPolicySpec) SetConstraintList(v []OPAPolicyConstraint)`

SetConstraintList sets ConstraintList field to given value.

### HasConstraintList

`func (o *OPAPolicySpec) HasConstraintList() bool`

HasConstraintList returns a boolean if a field has been set.

### GetSharing

`func (o *OPAPolicySpec) GetSharing() SharingSpec`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *OPAPolicySpec) GetSharingOk() (*SharingSpec, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *OPAPolicySpec) SetSharing(v SharingSpec)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *OPAPolicySpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetVersion

`func (o *OPAPolicySpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OPAPolicySpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OPAPolicySpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OPAPolicySpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


