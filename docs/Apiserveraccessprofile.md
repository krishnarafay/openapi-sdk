# Apiserveraccessprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizedIPRanges** | Pointer to **[]string** | IP ranges are specified in CIDR format, e.g. 137.117.106.88/29. This feature is not compatible with clusters that use Public IP Per Node, or clusters that are using a Basic Load Balancer. For more information see API server authorized IP ranges. | [optional] 
**DisableRunCommand** | Pointer to **bool** | Whether to disable run command for the cluster or not. | [optional] 
**EnablePrivateCluster** | Pointer to **bool** | For more details, see https://learn.microsoft.com/en-us/azure/aks/private-clusters | [optional] 
**EnablePrivateClusterPublicFQDN** | Pointer to **bool** | Whether to create additional public FQDN for private cluster or not. | [optional] 
**PrivateDNSZone** | Pointer to **string** | The default is System. For more details see configure private DNS zone. Allowed values are system and none. | [optional] 

## Methods

### NewApiserveraccessprofile

`func NewApiserveraccessprofile() *Apiserveraccessprofile`

NewApiserveraccessprofile instantiates a new Apiserveraccessprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiserveraccessprofileWithDefaults

`func NewApiserveraccessprofileWithDefaults() *Apiserveraccessprofile`

NewApiserveraccessprofileWithDefaults instantiates a new Apiserveraccessprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizedIPRanges

`func (o *Apiserveraccessprofile) GetAuthorizedIPRanges() []string`

GetAuthorizedIPRanges returns the AuthorizedIPRanges field if non-nil, zero value otherwise.

### GetAuthorizedIPRangesOk

`func (o *Apiserveraccessprofile) GetAuthorizedIPRangesOk() (*[]string, bool)`

GetAuthorizedIPRangesOk returns a tuple with the AuthorizedIPRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedIPRanges

`func (o *Apiserveraccessprofile) SetAuthorizedIPRanges(v []string)`

SetAuthorizedIPRanges sets AuthorizedIPRanges field to given value.

### HasAuthorizedIPRanges

`func (o *Apiserveraccessprofile) HasAuthorizedIPRanges() bool`

HasAuthorizedIPRanges returns a boolean if a field has been set.

### GetDisableRunCommand

`func (o *Apiserveraccessprofile) GetDisableRunCommand() bool`

GetDisableRunCommand returns the DisableRunCommand field if non-nil, zero value otherwise.

### GetDisableRunCommandOk

`func (o *Apiserveraccessprofile) GetDisableRunCommandOk() (*bool, bool)`

GetDisableRunCommandOk returns a tuple with the DisableRunCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRunCommand

`func (o *Apiserveraccessprofile) SetDisableRunCommand(v bool)`

SetDisableRunCommand sets DisableRunCommand field to given value.

### HasDisableRunCommand

`func (o *Apiserveraccessprofile) HasDisableRunCommand() bool`

HasDisableRunCommand returns a boolean if a field has been set.

### GetEnablePrivateCluster

`func (o *Apiserveraccessprofile) GetEnablePrivateCluster() bool`

GetEnablePrivateCluster returns the EnablePrivateCluster field if non-nil, zero value otherwise.

### GetEnablePrivateClusterOk

`func (o *Apiserveraccessprofile) GetEnablePrivateClusterOk() (*bool, bool)`

GetEnablePrivateClusterOk returns a tuple with the EnablePrivateCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePrivateCluster

`func (o *Apiserveraccessprofile) SetEnablePrivateCluster(v bool)`

SetEnablePrivateCluster sets EnablePrivateCluster field to given value.

### HasEnablePrivateCluster

`func (o *Apiserveraccessprofile) HasEnablePrivateCluster() bool`

HasEnablePrivateCluster returns a boolean if a field has been set.

### GetEnablePrivateClusterPublicFQDN

`func (o *Apiserveraccessprofile) GetEnablePrivateClusterPublicFQDN() bool`

GetEnablePrivateClusterPublicFQDN returns the EnablePrivateClusterPublicFQDN field if non-nil, zero value otherwise.

### GetEnablePrivateClusterPublicFQDNOk

`func (o *Apiserveraccessprofile) GetEnablePrivateClusterPublicFQDNOk() (*bool, bool)`

GetEnablePrivateClusterPublicFQDNOk returns a tuple with the EnablePrivateClusterPublicFQDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePrivateClusterPublicFQDN

`func (o *Apiserveraccessprofile) SetEnablePrivateClusterPublicFQDN(v bool)`

SetEnablePrivateClusterPublicFQDN sets EnablePrivateClusterPublicFQDN field to given value.

### HasEnablePrivateClusterPublicFQDN

`func (o *Apiserveraccessprofile) HasEnablePrivateClusterPublicFQDN() bool`

HasEnablePrivateClusterPublicFQDN returns a boolean if a field has been set.

### GetPrivateDNSZone

`func (o *Apiserveraccessprofile) GetPrivateDNSZone() string`

GetPrivateDNSZone returns the PrivateDNSZone field if non-nil, zero value otherwise.

### GetPrivateDNSZoneOk

`func (o *Apiserveraccessprofile) GetPrivateDNSZoneOk() (*string, bool)`

GetPrivateDNSZoneOk returns a tuple with the PrivateDNSZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDNSZone

`func (o *Apiserveraccessprofile) SetPrivateDNSZone(v string)`

SetPrivateDNSZone sets PrivateDNSZone field to given value.

### HasPrivateDNSZone

`func (o *Apiserveraccessprofile) HasPrivateDNSZone() bool`

HasPrivateDNSZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


