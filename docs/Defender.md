# Defender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogAnalyticsWorkspaceResourceId** | Pointer to **string** | Resource ID of the Log Analytics workspace to be associated with Microsoft Defender. When Microsoft Defender is enabled, this field is required and must be a valid workspace resource ID. When Microsoft Defender is disabled, leave the field empty. | [optional] 
**SecurityMonitoring** | Pointer to [**Securitymonitoring**](Securitymonitoring.md) |  | [optional] 

## Methods

### NewDefender

`func NewDefender() *Defender`

NewDefender instantiates a new Defender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefenderWithDefaults

`func NewDefenderWithDefaults() *Defender`

NewDefenderWithDefaults instantiates a new Defender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogAnalyticsWorkspaceResourceId

`func (o *Defender) GetLogAnalyticsWorkspaceResourceId() string`

GetLogAnalyticsWorkspaceResourceId returns the LogAnalyticsWorkspaceResourceId field if non-nil, zero value otherwise.

### GetLogAnalyticsWorkspaceResourceIdOk

`func (o *Defender) GetLogAnalyticsWorkspaceResourceIdOk() (*string, bool)`

GetLogAnalyticsWorkspaceResourceIdOk returns a tuple with the LogAnalyticsWorkspaceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAnalyticsWorkspaceResourceId

`func (o *Defender) SetLogAnalyticsWorkspaceResourceId(v string)`

SetLogAnalyticsWorkspaceResourceId sets LogAnalyticsWorkspaceResourceId field to given value.

### HasLogAnalyticsWorkspaceResourceId

`func (o *Defender) HasLogAnalyticsWorkspaceResourceId() bool`

HasLogAnalyticsWorkspaceResourceId returns a boolean if a field has been set.

### GetSecurityMonitoring

`func (o *Defender) GetSecurityMonitoring() Securitymonitoring`

GetSecurityMonitoring returns the SecurityMonitoring field if non-nil, zero value otherwise.

### GetSecurityMonitoringOk

`func (o *Defender) GetSecurityMonitoringOk() (*Securitymonitoring, bool)`

GetSecurityMonitoringOk returns a tuple with the SecurityMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMonitoring

`func (o *Defender) SetSecurityMonitoring(v Securitymonitoring)`

SetSecurityMonitoring sets SecurityMonitoring field to given value.

### HasSecurityMonitoring

`func (o *Defender) HasSecurityMonitoring() bool`

HasSecurityMonitoring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


