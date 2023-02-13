# ShellScriptOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuLimitMilli** | Pointer to **string** | Specify the cpu limit in milli | [optional] 
**Envvars** | Pointer to **map[string]string** | Specify the environment variables to be set in key,value pair | [optional] 
**MemoryLimitMB** | Pointer to **string** | Specify the memory limit to be allocated in MB | [optional] 
**Script** | Pointer to **string** | Specify the script to be executed | [optional] 
**SuccessCondition** | Pointer to **string** | Specify the success condition of the script | [optional] 

## Methods

### NewShellScriptOptions

`func NewShellScriptOptions() *ShellScriptOptions`

NewShellScriptOptions instantiates a new ShellScriptOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShellScriptOptionsWithDefaults

`func NewShellScriptOptionsWithDefaults() *ShellScriptOptions`

NewShellScriptOptionsWithDefaults instantiates a new ShellScriptOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuLimitMilli

`func (o *ShellScriptOptions) GetCpuLimitMilli() string`

GetCpuLimitMilli returns the CpuLimitMilli field if non-nil, zero value otherwise.

### GetCpuLimitMilliOk

`func (o *ShellScriptOptions) GetCpuLimitMilliOk() (*string, bool)`

GetCpuLimitMilliOk returns a tuple with the CpuLimitMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimitMilli

`func (o *ShellScriptOptions) SetCpuLimitMilli(v string)`

SetCpuLimitMilli sets CpuLimitMilli field to given value.

### HasCpuLimitMilli

`func (o *ShellScriptOptions) HasCpuLimitMilli() bool`

HasCpuLimitMilli returns a boolean if a field has been set.

### GetEnvvars

`func (o *ShellScriptOptions) GetEnvvars() map[string]string`

GetEnvvars returns the Envvars field if non-nil, zero value otherwise.

### GetEnvvarsOk

`func (o *ShellScriptOptions) GetEnvvarsOk() (*map[string]string, bool)`

GetEnvvarsOk returns a tuple with the Envvars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvvars

`func (o *ShellScriptOptions) SetEnvvars(v map[string]string)`

SetEnvvars sets Envvars field to given value.

### HasEnvvars

`func (o *ShellScriptOptions) HasEnvvars() bool`

HasEnvvars returns a boolean if a field has been set.

### GetMemoryLimitMB

`func (o *ShellScriptOptions) GetMemoryLimitMB() string`

GetMemoryLimitMB returns the MemoryLimitMB field if non-nil, zero value otherwise.

### GetMemoryLimitMBOk

`func (o *ShellScriptOptions) GetMemoryLimitMBOk() (*string, bool)`

GetMemoryLimitMBOk returns a tuple with the MemoryLimitMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitMB

`func (o *ShellScriptOptions) SetMemoryLimitMB(v string)`

SetMemoryLimitMB sets MemoryLimitMB field to given value.

### HasMemoryLimitMB

`func (o *ShellScriptOptions) HasMemoryLimitMB() bool`

HasMemoryLimitMB returns a boolean if a field has been set.

### GetScript

`func (o *ShellScriptOptions) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *ShellScriptOptions) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *ShellScriptOptions) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *ShellScriptOptions) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetSuccessCondition

`func (o *ShellScriptOptions) GetSuccessCondition() string`

GetSuccessCondition returns the SuccessCondition field if non-nil, zero value otherwise.

### GetSuccessConditionOk

`func (o *ShellScriptOptions) GetSuccessConditionOk() (*string, bool)`

GetSuccessConditionOk returns a tuple with the SuccessCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCondition

`func (o *ShellScriptOptions) SetSuccessCondition(v string)`

SetSuccessCondition sets SuccessCondition field to given value.

### HasSuccessCondition

`func (o *ShellScriptOptions) HasSuccessCondition() bool`

HasSuccessCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


