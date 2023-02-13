# SystemSyncConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitterEmail** | Pointer to **string** | value to use as committer email when committing to destination git repository. | [optional] 
**DestinationRepo** | Pointer to [**SystemSyncRepo**](SystemSyncRepo.md) |  | [optional] 
**ExcludedResources** | Pointer to [**[]SystemSyncResource**](SystemSyncResource.md) | resources to exclude from syncing | [optional] 
**GitToSystemSync** | Pointer to **bool** | flag to indicate if git to system sync should be enabled | [optional] 
**IncludedResources** | Pointer to [**[]SystemSyncResource**](SystemSyncResource.md) | resources to include for syncing | [optional] 
**SourceAsDestination** | Pointer to **bool** | flag to indicate if source repository should be used as destination repository | [optional] 
**SourceRepo** | Pointer to [**SystemSyncRepo**](SystemSyncRepo.md) |  | [optional] 
**SystemToGitSync** | Pointer to **bool** | flat to indicate if system to git sync should be enabled | [optional] 

## Methods

### NewSystemSyncConfig

`func NewSystemSyncConfig() *SystemSyncConfig`

NewSystemSyncConfig instantiates a new SystemSyncConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemSyncConfigWithDefaults

`func NewSystemSyncConfigWithDefaults() *SystemSyncConfig`

NewSystemSyncConfigWithDefaults instantiates a new SystemSyncConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitterEmail

`func (o *SystemSyncConfig) GetCommitterEmail() string`

GetCommitterEmail returns the CommitterEmail field if non-nil, zero value otherwise.

### GetCommitterEmailOk

`func (o *SystemSyncConfig) GetCommitterEmailOk() (*string, bool)`

GetCommitterEmailOk returns a tuple with the CommitterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitterEmail

`func (o *SystemSyncConfig) SetCommitterEmail(v string)`

SetCommitterEmail sets CommitterEmail field to given value.

### HasCommitterEmail

`func (o *SystemSyncConfig) HasCommitterEmail() bool`

HasCommitterEmail returns a boolean if a field has been set.

### GetDestinationRepo

`func (o *SystemSyncConfig) GetDestinationRepo() SystemSyncRepo`

GetDestinationRepo returns the DestinationRepo field if non-nil, zero value otherwise.

### GetDestinationRepoOk

`func (o *SystemSyncConfig) GetDestinationRepoOk() (*SystemSyncRepo, bool)`

GetDestinationRepoOk returns a tuple with the DestinationRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationRepo

`func (o *SystemSyncConfig) SetDestinationRepo(v SystemSyncRepo)`

SetDestinationRepo sets DestinationRepo field to given value.

### HasDestinationRepo

`func (o *SystemSyncConfig) HasDestinationRepo() bool`

HasDestinationRepo returns a boolean if a field has been set.

### GetExcludedResources

`func (o *SystemSyncConfig) GetExcludedResources() []SystemSyncResource`

GetExcludedResources returns the ExcludedResources field if non-nil, zero value otherwise.

### GetExcludedResourcesOk

`func (o *SystemSyncConfig) GetExcludedResourcesOk() (*[]SystemSyncResource, bool)`

GetExcludedResourcesOk returns a tuple with the ExcludedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedResources

`func (o *SystemSyncConfig) SetExcludedResources(v []SystemSyncResource)`

SetExcludedResources sets ExcludedResources field to given value.

### HasExcludedResources

`func (o *SystemSyncConfig) HasExcludedResources() bool`

HasExcludedResources returns a boolean if a field has been set.

### GetGitToSystemSync

`func (o *SystemSyncConfig) GetGitToSystemSync() bool`

GetGitToSystemSync returns the GitToSystemSync field if non-nil, zero value otherwise.

### GetGitToSystemSyncOk

`func (o *SystemSyncConfig) GetGitToSystemSyncOk() (*bool, bool)`

GetGitToSystemSyncOk returns a tuple with the GitToSystemSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitToSystemSync

`func (o *SystemSyncConfig) SetGitToSystemSync(v bool)`

SetGitToSystemSync sets GitToSystemSync field to given value.

### HasGitToSystemSync

`func (o *SystemSyncConfig) HasGitToSystemSync() bool`

HasGitToSystemSync returns a boolean if a field has been set.

### GetIncludedResources

`func (o *SystemSyncConfig) GetIncludedResources() []SystemSyncResource`

GetIncludedResources returns the IncludedResources field if non-nil, zero value otherwise.

### GetIncludedResourcesOk

`func (o *SystemSyncConfig) GetIncludedResourcesOk() (*[]SystemSyncResource, bool)`

GetIncludedResourcesOk returns a tuple with the IncludedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedResources

`func (o *SystemSyncConfig) SetIncludedResources(v []SystemSyncResource)`

SetIncludedResources sets IncludedResources field to given value.

### HasIncludedResources

`func (o *SystemSyncConfig) HasIncludedResources() bool`

HasIncludedResources returns a boolean if a field has been set.

### GetSourceAsDestination

`func (o *SystemSyncConfig) GetSourceAsDestination() bool`

GetSourceAsDestination returns the SourceAsDestination field if non-nil, zero value otherwise.

### GetSourceAsDestinationOk

`func (o *SystemSyncConfig) GetSourceAsDestinationOk() (*bool, bool)`

GetSourceAsDestinationOk returns a tuple with the SourceAsDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAsDestination

`func (o *SystemSyncConfig) SetSourceAsDestination(v bool)`

SetSourceAsDestination sets SourceAsDestination field to given value.

### HasSourceAsDestination

`func (o *SystemSyncConfig) HasSourceAsDestination() bool`

HasSourceAsDestination returns a boolean if a field has been set.

### GetSourceRepo

`func (o *SystemSyncConfig) GetSourceRepo() SystemSyncRepo`

GetSourceRepo returns the SourceRepo field if non-nil, zero value otherwise.

### GetSourceRepoOk

`func (o *SystemSyncConfig) GetSourceRepoOk() (*SystemSyncRepo, bool)`

GetSourceRepoOk returns a tuple with the SourceRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepo

`func (o *SystemSyncConfig) SetSourceRepo(v SystemSyncRepo)`

SetSourceRepo sets SourceRepo field to given value.

### HasSourceRepo

`func (o *SystemSyncConfig) HasSourceRepo() bool`

HasSourceRepo returns a boolean if a field has been set.

### GetSystemToGitSync

`func (o *SystemSyncConfig) GetSystemToGitSync() bool`

GetSystemToGitSync returns the SystemToGitSync field if non-nil, zero value otherwise.

### GetSystemToGitSyncOk

`func (o *SystemSyncConfig) GetSystemToGitSyncOk() (*bool, bool)`

GetSystemToGitSyncOk returns a tuple with the SystemToGitSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemToGitSync

`func (o *SystemSyncConfig) SetSystemToGitSync(v bool)`

SetSystemToGitSync sets SystemToGitSync field to given value.

### HasSystemToGitSync

`func (o *SystemSyncConfig) HasSystemToGitSync() bool`

HasSystemToGitSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


