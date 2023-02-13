# WorkloadSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifact** | Pointer to [**ArtifactSpec**](ArtifactSpec.md) |  | [optional] 
**Drift** | Pointer to [**DriftSpec**](DriftSpec.md) |  | [optional] 
**Namespace** | Pointer to **string** | namespace of the workload resource | [optional] 
**Placement** | Pointer to [**PlacementSpec**](PlacementSpec.md) |  | [optional] 
**Version** | Pointer to **string** | version of the workload resource | [optional] 

## Methods

### NewWorkloadSpec

`func NewWorkloadSpec() *WorkloadSpec`

NewWorkloadSpec instantiates a new WorkloadSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadSpecWithDefaults

`func NewWorkloadSpecWithDefaults() *WorkloadSpec`

NewWorkloadSpecWithDefaults instantiates a new WorkloadSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifact

`func (o *WorkloadSpec) GetArtifact() ArtifactSpec`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *WorkloadSpec) GetArtifactOk() (*ArtifactSpec, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *WorkloadSpec) SetArtifact(v ArtifactSpec)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *WorkloadSpec) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetDrift

`func (o *WorkloadSpec) GetDrift() DriftSpec`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *WorkloadSpec) GetDriftOk() (*DriftSpec, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *WorkloadSpec) SetDrift(v DriftSpec)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *WorkloadSpec) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetNamespace

`func (o *WorkloadSpec) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *WorkloadSpec) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *WorkloadSpec) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *WorkloadSpec) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPlacement

`func (o *WorkloadSpec) GetPlacement() PlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *WorkloadSpec) GetPlacementOk() (*PlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *WorkloadSpec) SetPlacement(v PlacementSpec)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *WorkloadSpec) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetVersion

`func (o *WorkloadSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkloadSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkloadSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkloadSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


