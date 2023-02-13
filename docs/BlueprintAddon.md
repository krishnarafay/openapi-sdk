# BlueprintAddon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependsOn** | Pointer to **[]string** | other addons current addons depends on | [optional] 
**Name** | **string** | name of the blueprint addon | 
**Version** | **string** | version of the blueprint addon | 

## Methods

### NewBlueprintAddon

`func NewBlueprintAddon(name string, version string, ) *BlueprintAddon`

NewBlueprintAddon instantiates a new BlueprintAddon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintAddonWithDefaults

`func NewBlueprintAddonWithDefaults() *BlueprintAddon`

NewBlueprintAddonWithDefaults instantiates a new BlueprintAddon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependsOn

`func (o *BlueprintAddon) GetDependsOn() []string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *BlueprintAddon) GetDependsOnOk() (*[]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *BlueprintAddon) SetDependsOn(v []string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *BlueprintAddon) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### GetName

`func (o *BlueprintAddon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintAddon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintAddon) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *BlueprintAddon) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlueprintAddon) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlueprintAddon) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


