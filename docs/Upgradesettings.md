# Upgradesettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSurge** | Pointer to **string** | This can either be set to an integer (e.g. 5) or a percentage (e.g. 50%). If a percentage is specified, it is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded up. If not specified, the default is 1. For more information, including best practices, see: https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade | [optional] 

## Methods

### NewUpgradesettings

`func NewUpgradesettings() *Upgradesettings`

NewUpgradesettings instantiates a new Upgradesettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradesettingsWithDefaults

`func NewUpgradesettingsWithDefaults() *Upgradesettings`

NewUpgradesettingsWithDefaults instantiates a new Upgradesettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSurge

`func (o *Upgradesettings) GetMaxSurge() string`

GetMaxSurge returns the MaxSurge field if non-nil, zero value otherwise.

### GetMaxSurgeOk

`func (o *Upgradesettings) GetMaxSurgeOk() (*string, bool)`

GetMaxSurgeOk returns a tuple with the MaxSurge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSurge

`func (o *Upgradesettings) SetMaxSurge(v string)`

SetMaxSurge sets MaxSurge field to given value.

### HasMaxSurge

`func (o *Upgradesettings) HasMaxSurge() bool`

HasMaxSurge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


