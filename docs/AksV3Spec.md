# AksV3Spec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagedCluster** | Pointer to [**Managedcluster**](Managedcluster.md) |  | [optional] 
**NodePools** | Pointer to [**[]Nodepool**](Nodepool.md) |  | [optional] 
**ResourceGroupName** | Pointer to **string** |  | [optional] 
**SubscriptionID** | Pointer to **string** |  | [optional] 

## Methods

### NewAksV3Spec

`func NewAksV3Spec() *AksV3Spec`

NewAksV3Spec instantiates a new AksV3Spec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAksV3SpecWithDefaults

`func NewAksV3SpecWithDefaults() *AksV3Spec`

NewAksV3SpecWithDefaults instantiates a new AksV3Spec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagedCluster

`func (o *AksV3Spec) GetManagedCluster() Managedcluster`

GetManagedCluster returns the ManagedCluster field if non-nil, zero value otherwise.

### GetManagedClusterOk

`func (o *AksV3Spec) GetManagedClusterOk() (*Managedcluster, bool)`

GetManagedClusterOk returns a tuple with the ManagedCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedCluster

`func (o *AksV3Spec) SetManagedCluster(v Managedcluster)`

SetManagedCluster sets ManagedCluster field to given value.

### HasManagedCluster

`func (o *AksV3Spec) HasManagedCluster() bool`

HasManagedCluster returns a boolean if a field has been set.

### GetNodePools

`func (o *AksV3Spec) GetNodePools() []Nodepool`

GetNodePools returns the NodePools field if non-nil, zero value otherwise.

### GetNodePoolsOk

`func (o *AksV3Spec) GetNodePoolsOk() (*[]Nodepool, bool)`

GetNodePoolsOk returns a tuple with the NodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePools

`func (o *AksV3Spec) SetNodePools(v []Nodepool)`

SetNodePools sets NodePools field to given value.

### HasNodePools

`func (o *AksV3Spec) HasNodePools() bool`

HasNodePools returns a boolean if a field has been set.

### GetResourceGroupName

`func (o *AksV3Spec) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AksV3Spec) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AksV3Spec) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AksV3Spec) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### GetSubscriptionID

`func (o *AksV3Spec) GetSubscriptionID() string`

GetSubscriptionID returns the SubscriptionID field if non-nil, zero value otherwise.

### GetSubscriptionIDOk

`func (o *AksV3Spec) GetSubscriptionIDOk() (*string, bool)`

GetSubscriptionIDOk returns a tuple with the SubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionID

`func (o *AksV3Spec) SetSubscriptionID(v string)`

SetSubscriptionID sets SubscriptionID field to given value.

### HasSubscriptionID

`func (o *AksV3Spec) HasSubscriptionID() bool`

HasSubscriptionID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


