# AKSStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityProfile** | Pointer to [**IdentityProfile**](IdentityProfile.md) |  | [optional] 
**Nodepools** | Pointer to [**[]NodepoolStatus**](NodepoolStatus.md) |  | [optional] 

## Methods

### NewAKSStatus

`func NewAKSStatus() *AKSStatus`

NewAKSStatus instantiates a new AKSStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAKSStatusWithDefaults

`func NewAKSStatusWithDefaults() *AKSStatus`

NewAKSStatusWithDefaults instantiates a new AKSStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityProfile

`func (o *AKSStatus) GetIdentityProfile() IdentityProfile`

GetIdentityProfile returns the IdentityProfile field if non-nil, zero value otherwise.

### GetIdentityProfileOk

`func (o *AKSStatus) GetIdentityProfileOk() (*IdentityProfile, bool)`

GetIdentityProfileOk returns a tuple with the IdentityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProfile

`func (o *AKSStatus) SetIdentityProfile(v IdentityProfile)`

SetIdentityProfile sets IdentityProfile field to given value.

### HasIdentityProfile

`func (o *AKSStatus) HasIdentityProfile() bool`

HasIdentityProfile returns a boolean if a field has been set.

### GetNodepools

`func (o *AKSStatus) GetNodepools() []NodepoolStatus`

GetNodepools returns the Nodepools field if non-nil, zero value otherwise.

### GetNodepoolsOk

`func (o *AKSStatus) GetNodepoolsOk() (*[]NodepoolStatus, bool)`

GetNodepoolsOk returns a tuple with the Nodepools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodepools

`func (o *AKSStatus) SetNodepools(v []NodepoolStatus)`

SetNodepools sets Nodepools field to given value.

### HasNodepools

`func (o *AKSStatus) HasNodepools() bool`

HasNodepools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


