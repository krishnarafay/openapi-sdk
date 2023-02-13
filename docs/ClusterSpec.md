# ClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlueprintConfig** | Pointer to [**BlueprintConfig**](BlueprintConfig.md) |  | [optional] 
**CloudCredentials** | Pointer to **string** | The credentials to be used to interact with the cloud infrastructure | [optional] 
**Config** | Pointer to [**ClusterSpecConfig**](ClusterSpecConfig.md) |  | [optional] 
**Drift** | Pointer to [**DriftSpec**](DriftSpec.md) |  | [optional] 
**ProxyConfig** | Pointer to [**ProxyConfig**](ProxyConfig.md) |  | [optional] 
**Sharing** | Pointer to [**Sharing**](Sharing.md) |  | [optional] 
**SystemComponentsPlacement** | Pointer to [**SystemComponentsPlacement**](SystemComponentsPlacement.md) |  | [optional] 
**Type** | Pointer to **string** | The type of the cluster this spec corresponds to | [optional] [default to "aws-eks"]

## Methods

### NewClusterSpec

`func NewClusterSpec() *ClusterSpec`

NewClusterSpec instantiates a new ClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSpecWithDefaults

`func NewClusterSpecWithDefaults() *ClusterSpec`

NewClusterSpecWithDefaults instantiates a new ClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlueprintConfig

`func (o *ClusterSpec) GetBlueprintConfig() BlueprintConfig`

GetBlueprintConfig returns the BlueprintConfig field if non-nil, zero value otherwise.

### GetBlueprintConfigOk

`func (o *ClusterSpec) GetBlueprintConfigOk() (*BlueprintConfig, bool)`

GetBlueprintConfigOk returns a tuple with the BlueprintConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintConfig

`func (o *ClusterSpec) SetBlueprintConfig(v BlueprintConfig)`

SetBlueprintConfig sets BlueprintConfig field to given value.

### HasBlueprintConfig

`func (o *ClusterSpec) HasBlueprintConfig() bool`

HasBlueprintConfig returns a boolean if a field has been set.

### GetCloudCredentials

`func (o *ClusterSpec) GetCloudCredentials() string`

GetCloudCredentials returns the CloudCredentials field if non-nil, zero value otherwise.

### GetCloudCredentialsOk

`func (o *ClusterSpec) GetCloudCredentialsOk() (*string, bool)`

GetCloudCredentialsOk returns a tuple with the CloudCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentials

`func (o *ClusterSpec) SetCloudCredentials(v string)`

SetCloudCredentials sets CloudCredentials field to given value.

### HasCloudCredentials

`func (o *ClusterSpec) HasCloudCredentials() bool`

HasCloudCredentials returns a boolean if a field has been set.

### GetConfig

`func (o *ClusterSpec) GetConfig() ClusterSpecConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ClusterSpec) GetConfigOk() (*ClusterSpecConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ClusterSpec) SetConfig(v ClusterSpecConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ClusterSpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDrift

`func (o *ClusterSpec) GetDrift() DriftSpec`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *ClusterSpec) GetDriftOk() (*DriftSpec, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *ClusterSpec) SetDrift(v DriftSpec)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *ClusterSpec) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetProxyConfig

`func (o *ClusterSpec) GetProxyConfig() ProxyConfig`

GetProxyConfig returns the ProxyConfig field if non-nil, zero value otherwise.

### GetProxyConfigOk

`func (o *ClusterSpec) GetProxyConfigOk() (*ProxyConfig, bool)`

GetProxyConfigOk returns a tuple with the ProxyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyConfig

`func (o *ClusterSpec) SetProxyConfig(v ProxyConfig)`

SetProxyConfig sets ProxyConfig field to given value.

### HasProxyConfig

`func (o *ClusterSpec) HasProxyConfig() bool`

HasProxyConfig returns a boolean if a field has been set.

### GetSharing

`func (o *ClusterSpec) GetSharing() Sharing`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *ClusterSpec) GetSharingOk() (*Sharing, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *ClusterSpec) SetSharing(v Sharing)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *ClusterSpec) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetSystemComponentsPlacement

`func (o *ClusterSpec) GetSystemComponentsPlacement() SystemComponentsPlacement`

GetSystemComponentsPlacement returns the SystemComponentsPlacement field if non-nil, zero value otherwise.

### GetSystemComponentsPlacementOk

`func (o *ClusterSpec) GetSystemComponentsPlacementOk() (*SystemComponentsPlacement, bool)`

GetSystemComponentsPlacementOk returns a tuple with the SystemComponentsPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemComponentsPlacement

`func (o *ClusterSpec) SetSystemComponentsPlacement(v SystemComponentsPlacement)`

SetSystemComponentsPlacement sets SystemComponentsPlacement field to given value.

### HasSystemComponentsPlacement

`func (o *ClusterSpec) HasSystemComponentsPlacement() bool`

HasSystemComponentsPlacement returns a boolean if a field has been set.

### GetType

`func (o *ClusterSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterSpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


