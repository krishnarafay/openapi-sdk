# InstallationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertType** | Pointer to **string** | Certificate Management Type | [optional] 
**EnableIngress** | Pointer to **bool** | Enable Ingress | [optional] 
**EnableNamespacesByDefault** | Pointer to **bool** | Enable Sidecar Injection Globally | [optional] 
**ResourceQuota** | Pointer to [**MeshResourceQuotas**](MeshResourceQuotas.md) |  | [optional] 

## Methods

### NewInstallationParams

`func NewInstallationParams() *InstallationParams`

NewInstallationParams instantiates a new InstallationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallationParamsWithDefaults

`func NewInstallationParamsWithDefaults() *InstallationParams`

NewInstallationParamsWithDefaults instantiates a new InstallationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertType

`func (o *InstallationParams) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *InstallationParams) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *InstallationParams) SetCertType(v string)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *InstallationParams) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### GetEnableIngress

`func (o *InstallationParams) GetEnableIngress() bool`

GetEnableIngress returns the EnableIngress field if non-nil, zero value otherwise.

### GetEnableIngressOk

`func (o *InstallationParams) GetEnableIngressOk() (*bool, bool)`

GetEnableIngressOk returns a tuple with the EnableIngress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIngress

`func (o *InstallationParams) SetEnableIngress(v bool)`

SetEnableIngress sets EnableIngress field to given value.

### HasEnableIngress

`func (o *InstallationParams) HasEnableIngress() bool`

HasEnableIngress returns a boolean if a field has been set.

### GetEnableNamespacesByDefault

`func (o *InstallationParams) GetEnableNamespacesByDefault() bool`

GetEnableNamespacesByDefault returns the EnableNamespacesByDefault field if non-nil, zero value otherwise.

### GetEnableNamespacesByDefaultOk

`func (o *InstallationParams) GetEnableNamespacesByDefaultOk() (*bool, bool)`

GetEnableNamespacesByDefaultOk returns a tuple with the EnableNamespacesByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNamespacesByDefault

`func (o *InstallationParams) SetEnableNamespacesByDefault(v bool)`

SetEnableNamespacesByDefault sets EnableNamespacesByDefault field to given value.

### HasEnableNamespacesByDefault

`func (o *InstallationParams) HasEnableNamespacesByDefault() bool`

HasEnableNamespacesByDefault returns a boolean if a field has been set.

### GetResourceQuota

`func (o *InstallationParams) GetResourceQuota() MeshResourceQuotas`

GetResourceQuota returns the ResourceQuota field if non-nil, zero value otherwise.

### GetResourceQuotaOk

`func (o *InstallationParams) GetResourceQuotaOk() (*MeshResourceQuotas, bool)`

GetResourceQuotaOk returns a tuple with the ResourceQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceQuota

`func (o *InstallationParams) SetResourceQuota(v MeshResourceQuotas)`

SetResourceQuota sets ResourceQuota field to given value.

### HasResourceQuota

`func (o *InstallationParams) HasResourceQuota() bool`

HasResourceQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


