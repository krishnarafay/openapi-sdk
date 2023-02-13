# InfraProvisionerSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**InfraProvisionerSpecConfig**](InfraProvisionerSpecConfig.md) |  | 
**FolderPath** | [**File**](File.md) |  | 
**Repository** | **string** | infrastructure provisioner git repository name | 
**Revision** | **string** | infrastructure provisioner git repository branch or tag | 
**Secret** | Pointer to [**File**](File.md) |  | [optional] 
**Type** | **string** | type of infrastructure provisioner | 

## Methods

### NewInfraProvisionerSpec

`func NewInfraProvisionerSpec(config InfraProvisionerSpecConfig, folderPath File, repository string, revision string, type_ string, ) *InfraProvisionerSpec`

NewInfraProvisionerSpec instantiates a new InfraProvisionerSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraProvisionerSpecWithDefaults

`func NewInfraProvisionerSpecWithDefaults() *InfraProvisionerSpec`

NewInfraProvisionerSpecWithDefaults instantiates a new InfraProvisionerSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *InfraProvisionerSpec) GetConfig() InfraProvisionerSpecConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InfraProvisionerSpec) GetConfigOk() (*InfraProvisionerSpecConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InfraProvisionerSpec) SetConfig(v InfraProvisionerSpecConfig)`

SetConfig sets Config field to given value.


### GetFolderPath

`func (o *InfraProvisionerSpec) GetFolderPath() File`

GetFolderPath returns the FolderPath field if non-nil, zero value otherwise.

### GetFolderPathOk

`func (o *InfraProvisionerSpec) GetFolderPathOk() (*File, bool)`

GetFolderPathOk returns a tuple with the FolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderPath

`func (o *InfraProvisionerSpec) SetFolderPath(v File)`

SetFolderPath sets FolderPath field to given value.


### GetRepository

`func (o *InfraProvisionerSpec) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *InfraProvisionerSpec) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *InfraProvisionerSpec) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetRevision

`func (o *InfraProvisionerSpec) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *InfraProvisionerSpec) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *InfraProvisionerSpec) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetSecret

`func (o *InfraProvisionerSpec) GetSecret() File`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *InfraProvisionerSpec) GetSecretOk() (*File, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *InfraProvisionerSpec) SetSecret(v File)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *InfraProvisionerSpec) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetType

`func (o *InfraProvisionerSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InfraProvisionerSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InfraProvisionerSpec) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


