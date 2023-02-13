# ContainerOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to **[]string** | Specify the set of arguments to be passed | [optional] 
**Commands** | Pointer to **[]string** | Specify the set of commands to be executed | [optional] 
**CpuLimitMilli** | Pointer to **string** | Specify the cpu limit in milli | [optional] 
**Envvars** | Pointer to **map[string]string** | Specify the environment variables to be set in key,value pair | [optional] 
**Image** | Pointer to **string** | Specify the container image to be used | [optional] 
**MemoryLimitMB** | Pointer to **string** | Specify the memory limit to be allocated in MB | [optional] 
**SuccessCondition** | Pointer to **string** | Specify the success condition of the container | [optional] 
**WorkingDirPath** | Pointer to **string** | Specify the working directory path | [optional] 

## Methods

### NewContainerOptions

`func NewContainerOptions() *ContainerOptions`

NewContainerOptions instantiates a new ContainerOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerOptionsWithDefaults

`func NewContainerOptionsWithDefaults() *ContainerOptions`

NewContainerOptionsWithDefaults instantiates a new ContainerOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *ContainerOptions) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ContainerOptions) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ContainerOptions) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *ContainerOptions) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetCommands

`func (o *ContainerOptions) GetCommands() []string`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *ContainerOptions) GetCommandsOk() (*[]string, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *ContainerOptions) SetCommands(v []string)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *ContainerOptions) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetCpuLimitMilli

`func (o *ContainerOptions) GetCpuLimitMilli() string`

GetCpuLimitMilli returns the CpuLimitMilli field if non-nil, zero value otherwise.

### GetCpuLimitMilliOk

`func (o *ContainerOptions) GetCpuLimitMilliOk() (*string, bool)`

GetCpuLimitMilliOk returns a tuple with the CpuLimitMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimitMilli

`func (o *ContainerOptions) SetCpuLimitMilli(v string)`

SetCpuLimitMilli sets CpuLimitMilli field to given value.

### HasCpuLimitMilli

`func (o *ContainerOptions) HasCpuLimitMilli() bool`

HasCpuLimitMilli returns a boolean if a field has been set.

### GetEnvvars

`func (o *ContainerOptions) GetEnvvars() map[string]string`

GetEnvvars returns the Envvars field if non-nil, zero value otherwise.

### GetEnvvarsOk

`func (o *ContainerOptions) GetEnvvarsOk() (*map[string]string, bool)`

GetEnvvarsOk returns a tuple with the Envvars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvvars

`func (o *ContainerOptions) SetEnvvars(v map[string]string)`

SetEnvvars sets Envvars field to given value.

### HasEnvvars

`func (o *ContainerOptions) HasEnvvars() bool`

HasEnvvars returns a boolean if a field has been set.

### GetImage

`func (o *ContainerOptions) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ContainerOptions) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ContainerOptions) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ContainerOptions) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMemoryLimitMB

`func (o *ContainerOptions) GetMemoryLimitMB() string`

GetMemoryLimitMB returns the MemoryLimitMB field if non-nil, zero value otherwise.

### GetMemoryLimitMBOk

`func (o *ContainerOptions) GetMemoryLimitMBOk() (*string, bool)`

GetMemoryLimitMBOk returns a tuple with the MemoryLimitMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitMB

`func (o *ContainerOptions) SetMemoryLimitMB(v string)`

SetMemoryLimitMB sets MemoryLimitMB field to given value.

### HasMemoryLimitMB

`func (o *ContainerOptions) HasMemoryLimitMB() bool`

HasMemoryLimitMB returns a boolean if a field has been set.

### GetSuccessCondition

`func (o *ContainerOptions) GetSuccessCondition() string`

GetSuccessCondition returns the SuccessCondition field if non-nil, zero value otherwise.

### GetSuccessConditionOk

`func (o *ContainerOptions) GetSuccessConditionOk() (*string, bool)`

GetSuccessConditionOk returns a tuple with the SuccessCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCondition

`func (o *ContainerOptions) SetSuccessCondition(v string)`

SetSuccessCondition sets SuccessCondition field to given value.

### HasSuccessCondition

`func (o *ContainerOptions) HasSuccessCondition() bool`

HasSuccessCondition returns a boolean if a field has been set.

### GetWorkingDirPath

`func (o *ContainerOptions) GetWorkingDirPath() string`

GetWorkingDirPath returns the WorkingDirPath field if non-nil, zero value otherwise.

### GetWorkingDirPathOk

`func (o *ContainerOptions) GetWorkingDirPathOk() (*string, bool)`

GetWorkingDirPathOk returns a tuple with the WorkingDirPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirPath

`func (o *ContainerOptions) SetWorkingDirPath(v string)`

SetWorkingDirPath sets WorkingDirPath field to given value.

### HasWorkingDirPath

`func (o *ContainerOptions) HasWorkingDirPath() bool`

HasWorkingDirPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


