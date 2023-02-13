# InfraProvisioner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the infrastructure provisioner resource | [default to "gitops.k8smgmt.io/v3"]
**Kind** | **string** | Kind of the infrastructure provisioner resource | [default to "Pipeline"]
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**InfraProvisionerSpec**](InfraProvisionerSpec.md) |  | 

## Methods

### NewInfraProvisioner

`func NewInfraProvisioner(apiVersion string, kind string, metadata Metadata, spec InfraProvisionerSpec, ) *InfraProvisioner`

NewInfraProvisioner instantiates a new InfraProvisioner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfraProvisionerWithDefaults

`func NewInfraProvisionerWithDefaults() *InfraProvisioner`

NewInfraProvisionerWithDefaults instantiates a new InfraProvisioner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *InfraProvisioner) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *InfraProvisioner) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *InfraProvisioner) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *InfraProvisioner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InfraProvisioner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InfraProvisioner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *InfraProvisioner) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InfraProvisioner) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InfraProvisioner) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *InfraProvisioner) GetSpec() InfraProvisionerSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *InfraProvisioner) GetSpecOk() (*InfraProvisionerSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *InfraProvisioner) SetSpec(v InfraProvisionerSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


