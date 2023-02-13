# ClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aks** | Pointer to [**AKSStatus**](AKSStatus.md) |  | [optional] 
**Blueprint** | Pointer to [**BlueprintStatus**](BlueprintStatus.md) |  | [optional] 
**CommonStatus** | Pointer to [**Status**](Status.md) |  | [optional] 
**Conditions** | Pointer to [**[]ClusterCondition**](ClusterCondition.md) |  | [optional] 
**ControPlane** | Pointer to [**ControlPlaneStatus**](ControlPlaneStatus.md) |  | [optional] 
**CreatedAt** | Pointer to [**StatusTime**](StatusTime.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Extra** | Pointer to [**ExtraStatus**](ExtraStatus.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Imported** | Pointer to [**ImportedStatus**](ImportedStatus.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProvisionStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterStatus

`func NewClusterStatus() *ClusterStatus`

NewClusterStatus instantiates a new ClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusWithDefaults

`func NewClusterStatusWithDefaults() *ClusterStatus`

NewClusterStatusWithDefaults instantiates a new ClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAks

`func (o *ClusterStatus) GetAks() AKSStatus`

GetAks returns the Aks field if non-nil, zero value otherwise.

### GetAksOk

`func (o *ClusterStatus) GetAksOk() (*AKSStatus, bool)`

GetAksOk returns a tuple with the Aks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAks

`func (o *ClusterStatus) SetAks(v AKSStatus)`

SetAks sets Aks field to given value.

### HasAks

`func (o *ClusterStatus) HasAks() bool`

HasAks returns a boolean if a field has been set.

### GetBlueprint

`func (o *ClusterStatus) GetBlueprint() BlueprintStatus`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *ClusterStatus) GetBlueprintOk() (*BlueprintStatus, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *ClusterStatus) SetBlueprint(v BlueprintStatus)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *ClusterStatus) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetCommonStatus

`func (o *ClusterStatus) GetCommonStatus() Status`

GetCommonStatus returns the CommonStatus field if non-nil, zero value otherwise.

### GetCommonStatusOk

`func (o *ClusterStatus) GetCommonStatusOk() (*Status, bool)`

GetCommonStatusOk returns a tuple with the CommonStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonStatus

`func (o *ClusterStatus) SetCommonStatus(v Status)`

SetCommonStatus sets CommonStatus field to given value.

### HasCommonStatus

`func (o *ClusterStatus) HasCommonStatus() bool`

HasCommonStatus returns a boolean if a field has been set.

### GetConditions

`func (o *ClusterStatus) GetConditions() []ClusterCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ClusterStatus) GetConditionsOk() (*[]ClusterCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ClusterStatus) SetConditions(v []ClusterCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ClusterStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetControPlane

`func (o *ClusterStatus) GetControPlane() ControlPlaneStatus`

GetControPlane returns the ControPlane field if non-nil, zero value otherwise.

### GetControPlaneOk

`func (o *ClusterStatus) GetControPlaneOk() (*ControlPlaneStatus, bool)`

GetControPlaneOk returns a tuple with the ControPlane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControPlane

`func (o *ClusterStatus) SetControPlane(v ControlPlaneStatus)`

SetControPlane sets ControPlane field to given value.

### HasControPlane

`func (o *ClusterStatus) HasControPlane() bool`

HasControPlane returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ClusterStatus) GetCreatedAt() StatusTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterStatus) GetCreatedAtOk() (*StatusTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterStatus) SetCreatedAt(v StatusTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ClusterStatus) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *ClusterStatus) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ClusterStatus) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ClusterStatus) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ClusterStatus) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtra

`func (o *ClusterStatus) GetExtra() ExtraStatus`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *ClusterStatus) GetExtraOk() (*ExtraStatus, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *ClusterStatus) SetExtra(v ExtraStatus)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *ClusterStatus) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetId

`func (o *ClusterStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImported

`func (o *ClusterStatus) GetImported() ImportedStatus`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *ClusterStatus) GetImportedOk() (*ImportedStatus, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *ClusterStatus) SetImported(v ImportedStatus)`

SetImported sets Imported field to given value.

### HasImported

`func (o *ClusterStatus) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetName

`func (o *ClusterStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvisionStatus

`func (o *ClusterStatus) GetProvisionStatus() string`

GetProvisionStatus returns the ProvisionStatus field if non-nil, zero value otherwise.

### GetProvisionStatusOk

`func (o *ClusterStatus) GetProvisionStatusOk() (*string, bool)`

GetProvisionStatusOk returns a tuple with the ProvisionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionStatus

`func (o *ClusterStatus) SetProvisionStatus(v string)`

SetProvisionStatus sets ProvisionStatus field to given value.

### HasProvisionStatus

`func (o *ClusterStatus) HasProvisionStatus() bool`

HasProvisionStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


