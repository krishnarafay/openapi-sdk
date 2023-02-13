# ArtifactSpecOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Atomic** | Pointer to **bool** | deploy Helm artifact with atomic flag | [optional] 
**CleanUpOnFail** | Pointer to **bool** | cleanup deployed resources when chart fails to deploy | [optional] 
**Description** | Pointer to **string** | custom description for the release | [optional] 
**DisableOpenAPIValidation** | Pointer to **bool** | disable OpenAPI validation while deploying the YAML | [optional] 
**Force** | Pointer to **bool** | deploy YAML artifact with force flag | [optional] 
**KeepHistory** | Pointer to **bool** | keep release history after uninstalling | [optional] 
**MaxHistory** | Pointer to **int32** | limit Helm artifact history | [optional] 
**NoHooks** | Pointer to **bool** | deploy Helm artifact without hooks | [optional] 
**RenderSubChartNotes** | Pointer to **bool** | render sub chart notes | [optional] 
**ResetValues** | Pointer to **bool** | reset existing helm values | [optional] 
**ReuseValues** | Pointer to **bool** | reuse existing values | [optional] 
**SetString** | Pointer to **[]string** | pass custom helm values as key&#x3D;value | [optional] 
**SkipCRDs** | Pointer to **bool** | skip deploying crds | [optional] 
**Timeout** | Pointer to **string** | timeout for waiting for the resources to become ready | [optional] 
**Wait** | Pointer to **bool** | deploy Helm artifact with wait flag | [optional] 

## Methods

### NewArtifactSpecOptions

`func NewArtifactSpecOptions() *ArtifactSpecOptions`

NewArtifactSpecOptions instantiates a new ArtifactSpecOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactSpecOptionsWithDefaults

`func NewArtifactSpecOptionsWithDefaults() *ArtifactSpecOptions`

NewArtifactSpecOptionsWithDefaults instantiates a new ArtifactSpecOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtomic

`func (o *ArtifactSpecOptions) GetAtomic() bool`

GetAtomic returns the Atomic field if non-nil, zero value otherwise.

### GetAtomicOk

`func (o *ArtifactSpecOptions) GetAtomicOk() (*bool, bool)`

GetAtomicOk returns a tuple with the Atomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtomic

`func (o *ArtifactSpecOptions) SetAtomic(v bool)`

SetAtomic sets Atomic field to given value.

### HasAtomic

`func (o *ArtifactSpecOptions) HasAtomic() bool`

HasAtomic returns a boolean if a field has been set.

### GetCleanUpOnFail

`func (o *ArtifactSpecOptions) GetCleanUpOnFail() bool`

GetCleanUpOnFail returns the CleanUpOnFail field if non-nil, zero value otherwise.

### GetCleanUpOnFailOk

`func (o *ArtifactSpecOptions) GetCleanUpOnFailOk() (*bool, bool)`

GetCleanUpOnFailOk returns a tuple with the CleanUpOnFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanUpOnFail

`func (o *ArtifactSpecOptions) SetCleanUpOnFail(v bool)`

SetCleanUpOnFail sets CleanUpOnFail field to given value.

### HasCleanUpOnFail

`func (o *ArtifactSpecOptions) HasCleanUpOnFail() bool`

HasCleanUpOnFail returns a boolean if a field has been set.

### GetDescription

`func (o *ArtifactSpecOptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArtifactSpecOptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArtifactSpecOptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArtifactSpecOptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableOpenAPIValidation

`func (o *ArtifactSpecOptions) GetDisableOpenAPIValidation() bool`

GetDisableOpenAPIValidation returns the DisableOpenAPIValidation field if non-nil, zero value otherwise.

### GetDisableOpenAPIValidationOk

`func (o *ArtifactSpecOptions) GetDisableOpenAPIValidationOk() (*bool, bool)`

GetDisableOpenAPIValidationOk returns a tuple with the DisableOpenAPIValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableOpenAPIValidation

`func (o *ArtifactSpecOptions) SetDisableOpenAPIValidation(v bool)`

SetDisableOpenAPIValidation sets DisableOpenAPIValidation field to given value.

### HasDisableOpenAPIValidation

`func (o *ArtifactSpecOptions) HasDisableOpenAPIValidation() bool`

HasDisableOpenAPIValidation returns a boolean if a field has been set.

### GetForce

`func (o *ArtifactSpecOptions) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *ArtifactSpecOptions) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *ArtifactSpecOptions) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *ArtifactSpecOptions) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetKeepHistory

`func (o *ArtifactSpecOptions) GetKeepHistory() bool`

GetKeepHistory returns the KeepHistory field if non-nil, zero value otherwise.

### GetKeepHistoryOk

`func (o *ArtifactSpecOptions) GetKeepHistoryOk() (*bool, bool)`

GetKeepHistoryOk returns a tuple with the KeepHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepHistory

`func (o *ArtifactSpecOptions) SetKeepHistory(v bool)`

SetKeepHistory sets KeepHistory field to given value.

### HasKeepHistory

`func (o *ArtifactSpecOptions) HasKeepHistory() bool`

HasKeepHistory returns a boolean if a field has been set.

### GetMaxHistory

`func (o *ArtifactSpecOptions) GetMaxHistory() int32`

GetMaxHistory returns the MaxHistory field if non-nil, zero value otherwise.

### GetMaxHistoryOk

`func (o *ArtifactSpecOptions) GetMaxHistoryOk() (*int32, bool)`

GetMaxHistoryOk returns a tuple with the MaxHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHistory

`func (o *ArtifactSpecOptions) SetMaxHistory(v int32)`

SetMaxHistory sets MaxHistory field to given value.

### HasMaxHistory

`func (o *ArtifactSpecOptions) HasMaxHistory() bool`

HasMaxHistory returns a boolean if a field has been set.

### GetNoHooks

`func (o *ArtifactSpecOptions) GetNoHooks() bool`

GetNoHooks returns the NoHooks field if non-nil, zero value otherwise.

### GetNoHooksOk

`func (o *ArtifactSpecOptions) GetNoHooksOk() (*bool, bool)`

GetNoHooksOk returns a tuple with the NoHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoHooks

`func (o *ArtifactSpecOptions) SetNoHooks(v bool)`

SetNoHooks sets NoHooks field to given value.

### HasNoHooks

`func (o *ArtifactSpecOptions) HasNoHooks() bool`

HasNoHooks returns a boolean if a field has been set.

### GetRenderSubChartNotes

`func (o *ArtifactSpecOptions) GetRenderSubChartNotes() bool`

GetRenderSubChartNotes returns the RenderSubChartNotes field if non-nil, zero value otherwise.

### GetRenderSubChartNotesOk

`func (o *ArtifactSpecOptions) GetRenderSubChartNotesOk() (*bool, bool)`

GetRenderSubChartNotesOk returns a tuple with the RenderSubChartNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderSubChartNotes

`func (o *ArtifactSpecOptions) SetRenderSubChartNotes(v bool)`

SetRenderSubChartNotes sets RenderSubChartNotes field to given value.

### HasRenderSubChartNotes

`func (o *ArtifactSpecOptions) HasRenderSubChartNotes() bool`

HasRenderSubChartNotes returns a boolean if a field has been set.

### GetResetValues

`func (o *ArtifactSpecOptions) GetResetValues() bool`

GetResetValues returns the ResetValues field if non-nil, zero value otherwise.

### GetResetValuesOk

`func (o *ArtifactSpecOptions) GetResetValuesOk() (*bool, bool)`

GetResetValuesOk returns a tuple with the ResetValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetValues

`func (o *ArtifactSpecOptions) SetResetValues(v bool)`

SetResetValues sets ResetValues field to given value.

### HasResetValues

`func (o *ArtifactSpecOptions) HasResetValues() bool`

HasResetValues returns a boolean if a field has been set.

### GetReuseValues

`func (o *ArtifactSpecOptions) GetReuseValues() bool`

GetReuseValues returns the ReuseValues field if non-nil, zero value otherwise.

### GetReuseValuesOk

`func (o *ArtifactSpecOptions) GetReuseValuesOk() (*bool, bool)`

GetReuseValuesOk returns a tuple with the ReuseValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseValues

`func (o *ArtifactSpecOptions) SetReuseValues(v bool)`

SetReuseValues sets ReuseValues field to given value.

### HasReuseValues

`func (o *ArtifactSpecOptions) HasReuseValues() bool`

HasReuseValues returns a boolean if a field has been set.

### GetSetString

`func (o *ArtifactSpecOptions) GetSetString() []string`

GetSetString returns the SetString field if non-nil, zero value otherwise.

### GetSetStringOk

`func (o *ArtifactSpecOptions) GetSetStringOk() (*[]string, bool)`

GetSetStringOk returns a tuple with the SetString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetString

`func (o *ArtifactSpecOptions) SetSetString(v []string)`

SetSetString sets SetString field to given value.

### HasSetString

`func (o *ArtifactSpecOptions) HasSetString() bool`

HasSetString returns a boolean if a field has been set.

### GetSkipCRDs

`func (o *ArtifactSpecOptions) GetSkipCRDs() bool`

GetSkipCRDs returns the SkipCRDs field if non-nil, zero value otherwise.

### GetSkipCRDsOk

`func (o *ArtifactSpecOptions) GetSkipCRDsOk() (*bool, bool)`

GetSkipCRDsOk returns a tuple with the SkipCRDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCRDs

`func (o *ArtifactSpecOptions) SetSkipCRDs(v bool)`

SetSkipCRDs sets SkipCRDs field to given value.

### HasSkipCRDs

`func (o *ArtifactSpecOptions) HasSkipCRDs() bool`

HasSkipCRDs returns a boolean if a field has been set.

### GetTimeout

`func (o *ArtifactSpecOptions) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ArtifactSpecOptions) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ArtifactSpecOptions) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ArtifactSpecOptions) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetWait

`func (o *ArtifactSpecOptions) GetWait() bool`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *ArtifactSpecOptions) GetWaitOk() (*bool, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *ArtifactSpecOptions) SetWait(v bool)`

SetWait sets Wait field to given value.

### HasWait

`func (o *ArtifactSpecOptions) HasWait() bool`

HasWait returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


