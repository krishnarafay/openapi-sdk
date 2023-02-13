# Autoscalerprofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceSimilarNodeGroups** | Pointer to **string** | Valid values are true and false | [optional] 
**Expander** | Pointer to **string** | If not specified, the default is random. See https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders for more information. Valid values are least-waste, most-pods, priority, random. | [optional] [default to "least-waste"]
**MaxEmptyBulkDelete** | Pointer to **string** | The default is 10. | [optional] 
**MaxGracefulTerminationSec** | Pointer to **string** | The default is 600. | [optional] 
**MaxNodeProvisionTime** | Pointer to **string** | The default is 15m. Values must be an integer followed by an m. No unit of time other than minutes (m) is supported. | [optional] 
**MaxTotalUnreadyPercentage** | Pointer to **string** | The default is 45. The maximum is 100 and the minimum is 0. | [optional] 
**NewPodScaleUpDelay** | Pointer to **string** | For scenarios like burst/batch scale where you don&#39;t want CA to act before the kubernetes scheduler could schedule all the pods, you can tell CA to ignore unscheduled pods before they&#39;re a certain age. The default is 0s. Values must be an integer followed by a unit (s for seconds, m for minutes, h for hours, etc). | [optional] 
**OkTotalUnreadyCount** | Pointer to **string** | This must be an integer. The default is 3. | [optional] 
**ScaleDownDelayAfterAdd** | Pointer to **string** | The default is 10m. Values must be an integer followed by an m. No unit of time other than minutes (m) is supported. | [optional] 
**ScaleDownDelayAfterDelete** | Pointer to **string** | The default is the scan-interval. Values must be an integer followed by an m. No unit of time other than minutes (m) is supported. | [optional] 
**ScaleDownDelayAfterFailure** | Pointer to **string** | The default is 3m. Values must be an integer followed by an m. No unit of time other than minutes (m) is supported. | [optional] 
**ScaleDownUnneededTime** | Pointer to **string** | The default is 10m. Values must be an integer followed by an m. No unit of time other than minutes (m) is supported. | [optional] 
**ScaleDownUnreadyTime** | Pointer to **string** | The default is 20m. Values must be an integer followed by an m. No unit of time other than minutes (m) is supported. | [optional] 
**ScaleDownUtilizationThreshold** | Pointer to **string** | The default is 0.5. | [optional] 
**ScanInterval** | Pointer to **string** | The default is 10. Values must be an integer number of seconds. | [optional] 
**SkipNodesWithLocalStorage** | Pointer to **string** | The default is true. | [optional] 
**SkipNodesWithSystemPods** | Pointer to **string** | The default is true. | [optional] 

## Methods

### NewAutoscalerprofile

`func NewAutoscalerprofile() *Autoscalerprofile`

NewAutoscalerprofile instantiates a new Autoscalerprofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalerprofileWithDefaults

`func NewAutoscalerprofileWithDefaults() *Autoscalerprofile`

NewAutoscalerprofileWithDefaults instantiates a new Autoscalerprofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceSimilarNodeGroups

`func (o *Autoscalerprofile) GetBalanceSimilarNodeGroups() string`

GetBalanceSimilarNodeGroups returns the BalanceSimilarNodeGroups field if non-nil, zero value otherwise.

### GetBalanceSimilarNodeGroupsOk

`func (o *Autoscalerprofile) GetBalanceSimilarNodeGroupsOk() (*string, bool)`

GetBalanceSimilarNodeGroupsOk returns a tuple with the BalanceSimilarNodeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceSimilarNodeGroups

`func (o *Autoscalerprofile) SetBalanceSimilarNodeGroups(v string)`

SetBalanceSimilarNodeGroups sets BalanceSimilarNodeGroups field to given value.

### HasBalanceSimilarNodeGroups

`func (o *Autoscalerprofile) HasBalanceSimilarNodeGroups() bool`

HasBalanceSimilarNodeGroups returns a boolean if a field has been set.

### GetExpander

`func (o *Autoscalerprofile) GetExpander() string`

GetExpander returns the Expander field if non-nil, zero value otherwise.

### GetExpanderOk

`func (o *Autoscalerprofile) GetExpanderOk() (*string, bool)`

GetExpanderOk returns a tuple with the Expander field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpander

`func (o *Autoscalerprofile) SetExpander(v string)`

SetExpander sets Expander field to given value.

### HasExpander

`func (o *Autoscalerprofile) HasExpander() bool`

HasExpander returns a boolean if a field has been set.

### GetMaxEmptyBulkDelete

`func (o *Autoscalerprofile) GetMaxEmptyBulkDelete() string`

GetMaxEmptyBulkDelete returns the MaxEmptyBulkDelete field if non-nil, zero value otherwise.

### GetMaxEmptyBulkDeleteOk

`func (o *Autoscalerprofile) GetMaxEmptyBulkDeleteOk() (*string, bool)`

GetMaxEmptyBulkDeleteOk returns a tuple with the MaxEmptyBulkDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEmptyBulkDelete

`func (o *Autoscalerprofile) SetMaxEmptyBulkDelete(v string)`

SetMaxEmptyBulkDelete sets MaxEmptyBulkDelete field to given value.

### HasMaxEmptyBulkDelete

`func (o *Autoscalerprofile) HasMaxEmptyBulkDelete() bool`

HasMaxEmptyBulkDelete returns a boolean if a field has been set.

### GetMaxGracefulTerminationSec

`func (o *Autoscalerprofile) GetMaxGracefulTerminationSec() string`

GetMaxGracefulTerminationSec returns the MaxGracefulTerminationSec field if non-nil, zero value otherwise.

### GetMaxGracefulTerminationSecOk

`func (o *Autoscalerprofile) GetMaxGracefulTerminationSecOk() (*string, bool)`

GetMaxGracefulTerminationSecOk returns a tuple with the MaxGracefulTerminationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGracefulTerminationSec

`func (o *Autoscalerprofile) SetMaxGracefulTerminationSec(v string)`

SetMaxGracefulTerminationSec sets MaxGracefulTerminationSec field to given value.

### HasMaxGracefulTerminationSec

`func (o *Autoscalerprofile) HasMaxGracefulTerminationSec() bool`

HasMaxGracefulTerminationSec returns a boolean if a field has been set.

### GetMaxNodeProvisionTime

`func (o *Autoscalerprofile) GetMaxNodeProvisionTime() string`

GetMaxNodeProvisionTime returns the MaxNodeProvisionTime field if non-nil, zero value otherwise.

### GetMaxNodeProvisionTimeOk

`func (o *Autoscalerprofile) GetMaxNodeProvisionTimeOk() (*string, bool)`

GetMaxNodeProvisionTimeOk returns a tuple with the MaxNodeProvisionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeProvisionTime

`func (o *Autoscalerprofile) SetMaxNodeProvisionTime(v string)`

SetMaxNodeProvisionTime sets MaxNodeProvisionTime field to given value.

### HasMaxNodeProvisionTime

`func (o *Autoscalerprofile) HasMaxNodeProvisionTime() bool`

HasMaxNodeProvisionTime returns a boolean if a field has been set.

### GetMaxTotalUnreadyPercentage

`func (o *Autoscalerprofile) GetMaxTotalUnreadyPercentage() string`

GetMaxTotalUnreadyPercentage returns the MaxTotalUnreadyPercentage field if non-nil, zero value otherwise.

### GetMaxTotalUnreadyPercentageOk

`func (o *Autoscalerprofile) GetMaxTotalUnreadyPercentageOk() (*string, bool)`

GetMaxTotalUnreadyPercentageOk returns a tuple with the MaxTotalUnreadyPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalUnreadyPercentage

`func (o *Autoscalerprofile) SetMaxTotalUnreadyPercentage(v string)`

SetMaxTotalUnreadyPercentage sets MaxTotalUnreadyPercentage field to given value.

### HasMaxTotalUnreadyPercentage

`func (o *Autoscalerprofile) HasMaxTotalUnreadyPercentage() bool`

HasMaxTotalUnreadyPercentage returns a boolean if a field has been set.

### GetNewPodScaleUpDelay

`func (o *Autoscalerprofile) GetNewPodScaleUpDelay() string`

GetNewPodScaleUpDelay returns the NewPodScaleUpDelay field if non-nil, zero value otherwise.

### GetNewPodScaleUpDelayOk

`func (o *Autoscalerprofile) GetNewPodScaleUpDelayOk() (*string, bool)`

GetNewPodScaleUpDelayOk returns a tuple with the NewPodScaleUpDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPodScaleUpDelay

`func (o *Autoscalerprofile) SetNewPodScaleUpDelay(v string)`

SetNewPodScaleUpDelay sets NewPodScaleUpDelay field to given value.

### HasNewPodScaleUpDelay

`func (o *Autoscalerprofile) HasNewPodScaleUpDelay() bool`

HasNewPodScaleUpDelay returns a boolean if a field has been set.

### GetOkTotalUnreadyCount

`func (o *Autoscalerprofile) GetOkTotalUnreadyCount() string`

GetOkTotalUnreadyCount returns the OkTotalUnreadyCount field if non-nil, zero value otherwise.

### GetOkTotalUnreadyCountOk

`func (o *Autoscalerprofile) GetOkTotalUnreadyCountOk() (*string, bool)`

GetOkTotalUnreadyCountOk returns a tuple with the OkTotalUnreadyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOkTotalUnreadyCount

`func (o *Autoscalerprofile) SetOkTotalUnreadyCount(v string)`

SetOkTotalUnreadyCount sets OkTotalUnreadyCount field to given value.

### HasOkTotalUnreadyCount

`func (o *Autoscalerprofile) HasOkTotalUnreadyCount() bool`

HasOkTotalUnreadyCount returns a boolean if a field has been set.

### GetScaleDownDelayAfterAdd

`func (o *Autoscalerprofile) GetScaleDownDelayAfterAdd() string`

GetScaleDownDelayAfterAdd returns the ScaleDownDelayAfterAdd field if non-nil, zero value otherwise.

### GetScaleDownDelayAfterAddOk

`func (o *Autoscalerprofile) GetScaleDownDelayAfterAddOk() (*string, bool)`

GetScaleDownDelayAfterAddOk returns a tuple with the ScaleDownDelayAfterAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleDownDelayAfterAdd

`func (o *Autoscalerprofile) SetScaleDownDelayAfterAdd(v string)`

SetScaleDownDelayAfterAdd sets ScaleDownDelayAfterAdd field to given value.

### HasScaleDownDelayAfterAdd

`func (o *Autoscalerprofile) HasScaleDownDelayAfterAdd() bool`

HasScaleDownDelayAfterAdd returns a boolean if a field has been set.

### GetScaleDownDelayAfterDelete

`func (o *Autoscalerprofile) GetScaleDownDelayAfterDelete() string`

GetScaleDownDelayAfterDelete returns the ScaleDownDelayAfterDelete field if non-nil, zero value otherwise.

### GetScaleDownDelayAfterDeleteOk

`func (o *Autoscalerprofile) GetScaleDownDelayAfterDeleteOk() (*string, bool)`

GetScaleDownDelayAfterDeleteOk returns a tuple with the ScaleDownDelayAfterDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleDownDelayAfterDelete

`func (o *Autoscalerprofile) SetScaleDownDelayAfterDelete(v string)`

SetScaleDownDelayAfterDelete sets ScaleDownDelayAfterDelete field to given value.

### HasScaleDownDelayAfterDelete

`func (o *Autoscalerprofile) HasScaleDownDelayAfterDelete() bool`

HasScaleDownDelayAfterDelete returns a boolean if a field has been set.

### GetScaleDownDelayAfterFailure

`func (o *Autoscalerprofile) GetScaleDownDelayAfterFailure() string`

GetScaleDownDelayAfterFailure returns the ScaleDownDelayAfterFailure field if non-nil, zero value otherwise.

### GetScaleDownDelayAfterFailureOk

`func (o *Autoscalerprofile) GetScaleDownDelayAfterFailureOk() (*string, bool)`

GetScaleDownDelayAfterFailureOk returns a tuple with the ScaleDownDelayAfterFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleDownDelayAfterFailure

`func (o *Autoscalerprofile) SetScaleDownDelayAfterFailure(v string)`

SetScaleDownDelayAfterFailure sets ScaleDownDelayAfterFailure field to given value.

### HasScaleDownDelayAfterFailure

`func (o *Autoscalerprofile) HasScaleDownDelayAfterFailure() bool`

HasScaleDownDelayAfterFailure returns a boolean if a field has been set.

### GetScaleDownUnneededTime

`func (o *Autoscalerprofile) GetScaleDownUnneededTime() string`

GetScaleDownUnneededTime returns the ScaleDownUnneededTime field if non-nil, zero value otherwise.

### GetScaleDownUnneededTimeOk

`func (o *Autoscalerprofile) GetScaleDownUnneededTimeOk() (*string, bool)`

GetScaleDownUnneededTimeOk returns a tuple with the ScaleDownUnneededTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleDownUnneededTime

`func (o *Autoscalerprofile) SetScaleDownUnneededTime(v string)`

SetScaleDownUnneededTime sets ScaleDownUnneededTime field to given value.

### HasScaleDownUnneededTime

`func (o *Autoscalerprofile) HasScaleDownUnneededTime() bool`

HasScaleDownUnneededTime returns a boolean if a field has been set.

### GetScaleDownUnreadyTime

`func (o *Autoscalerprofile) GetScaleDownUnreadyTime() string`

GetScaleDownUnreadyTime returns the ScaleDownUnreadyTime field if non-nil, zero value otherwise.

### GetScaleDownUnreadyTimeOk

`func (o *Autoscalerprofile) GetScaleDownUnreadyTimeOk() (*string, bool)`

GetScaleDownUnreadyTimeOk returns a tuple with the ScaleDownUnreadyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleDownUnreadyTime

`func (o *Autoscalerprofile) SetScaleDownUnreadyTime(v string)`

SetScaleDownUnreadyTime sets ScaleDownUnreadyTime field to given value.

### HasScaleDownUnreadyTime

`func (o *Autoscalerprofile) HasScaleDownUnreadyTime() bool`

HasScaleDownUnreadyTime returns a boolean if a field has been set.

### GetScaleDownUtilizationThreshold

`func (o *Autoscalerprofile) GetScaleDownUtilizationThreshold() string`

GetScaleDownUtilizationThreshold returns the ScaleDownUtilizationThreshold field if non-nil, zero value otherwise.

### GetScaleDownUtilizationThresholdOk

`func (o *Autoscalerprofile) GetScaleDownUtilizationThresholdOk() (*string, bool)`

GetScaleDownUtilizationThresholdOk returns a tuple with the ScaleDownUtilizationThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleDownUtilizationThreshold

`func (o *Autoscalerprofile) SetScaleDownUtilizationThreshold(v string)`

SetScaleDownUtilizationThreshold sets ScaleDownUtilizationThreshold field to given value.

### HasScaleDownUtilizationThreshold

`func (o *Autoscalerprofile) HasScaleDownUtilizationThreshold() bool`

HasScaleDownUtilizationThreshold returns a boolean if a field has been set.

### GetScanInterval

`func (o *Autoscalerprofile) GetScanInterval() string`

GetScanInterval returns the ScanInterval field if non-nil, zero value otherwise.

### GetScanIntervalOk

`func (o *Autoscalerprofile) GetScanIntervalOk() (*string, bool)`

GetScanIntervalOk returns a tuple with the ScanInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanInterval

`func (o *Autoscalerprofile) SetScanInterval(v string)`

SetScanInterval sets ScanInterval field to given value.

### HasScanInterval

`func (o *Autoscalerprofile) HasScanInterval() bool`

HasScanInterval returns a boolean if a field has been set.

### GetSkipNodesWithLocalStorage

`func (o *Autoscalerprofile) GetSkipNodesWithLocalStorage() string`

GetSkipNodesWithLocalStorage returns the SkipNodesWithLocalStorage field if non-nil, zero value otherwise.

### GetSkipNodesWithLocalStorageOk

`func (o *Autoscalerprofile) GetSkipNodesWithLocalStorageOk() (*string, bool)`

GetSkipNodesWithLocalStorageOk returns a tuple with the SkipNodesWithLocalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipNodesWithLocalStorage

`func (o *Autoscalerprofile) SetSkipNodesWithLocalStorage(v string)`

SetSkipNodesWithLocalStorage sets SkipNodesWithLocalStorage field to given value.

### HasSkipNodesWithLocalStorage

`func (o *Autoscalerprofile) HasSkipNodesWithLocalStorage() bool`

HasSkipNodesWithLocalStorage returns a boolean if a field has been set.

### GetSkipNodesWithSystemPods

`func (o *Autoscalerprofile) GetSkipNodesWithSystemPods() string`

GetSkipNodesWithSystemPods returns the SkipNodesWithSystemPods field if non-nil, zero value otherwise.

### GetSkipNodesWithSystemPodsOk

`func (o *Autoscalerprofile) GetSkipNodesWithSystemPodsOk() (*string, bool)`

GetSkipNodesWithSystemPodsOk returns a tuple with the SkipNodesWithSystemPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipNodesWithSystemPods

`func (o *Autoscalerprofile) SetSkipNodesWithSystemPods(v string)`

SetSkipNodesWithSystemPods sets SkipNodesWithSystemPods field to given value.

### HasSkipNodesWithSystemPods

`func (o *Autoscalerprofile) HasSkipNodesWithSystemPods() bool`

HasSkipNodesWithSystemPods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


