# HookOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approval** | Pointer to **map[string]interface{}** |  | [optional] 
**Container** | Pointer to [**ContainerOptions**](ContainerOptions.md) |  | [optional] 
**Http** | Pointer to [**HttpOptions**](HttpOptions.md) |  | [optional] 
**Notification** | Pointer to **map[string]interface{}** |  | [optional] 
**Script** | Pointer to [**ShellScriptOptions**](ShellScriptOptions.md) |  | [optional] 

## Methods

### NewHookOptions

`func NewHookOptions() *HookOptions`

NewHookOptions instantiates a new HookOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookOptionsWithDefaults

`func NewHookOptionsWithDefaults() *HookOptions`

NewHookOptionsWithDefaults instantiates a new HookOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproval

`func (o *HookOptions) GetApproval() map[string]interface{}`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *HookOptions) GetApprovalOk() (*map[string]interface{}, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *HookOptions) SetApproval(v map[string]interface{})`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *HookOptions) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### GetContainer

`func (o *HookOptions) GetContainer() ContainerOptions`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *HookOptions) GetContainerOk() (*ContainerOptions, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *HookOptions) SetContainer(v ContainerOptions)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *HookOptions) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetHttp

`func (o *HookOptions) GetHttp() HttpOptions`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *HookOptions) GetHttpOk() (*HttpOptions, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *HookOptions) SetHttp(v HttpOptions)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *HookOptions) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetNotification

`func (o *HookOptions) GetNotification() map[string]interface{}`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *HookOptions) GetNotificationOk() (*map[string]interface{}, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *HookOptions) SetNotification(v map[string]interface{})`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *HookOptions) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### GetScript

`func (o *HookOptions) GetScript() ShellScriptOptions`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *HookOptions) GetScriptOk() (*ShellScriptOptions, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *HookOptions) SetScript(v ShellScriptOptions)`

SetScript sets Script field to given value.

### HasScript

`func (o *HookOptions) HasScript() bool`

HasScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


