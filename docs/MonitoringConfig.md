# MonitoringConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HelmExporter** | [**MonitoringComponent**](MonitoringComponent.md) |  | 
**KubeStateMetrics** | [**MonitoringComponent**](MonitoringComponent.md) |  | 
**MetricsServer** | [**MonitoringComponent**](MonitoringComponent.md) |  | 
**NodeExporter** | [**MonitoringComponent**](MonitoringComponent.md) |  | 
**PrometheusAdapter** | [**MonitoringComponent**](MonitoringComponent.md) |  | 
**Resources** | Pointer to [**ResourceRequirements**](ResourceRequirements.md) |  | [optional] 

## Methods

### NewMonitoringConfig

`func NewMonitoringConfig(helmExporter MonitoringComponent, kubeStateMetrics MonitoringComponent, metricsServer MonitoringComponent, nodeExporter MonitoringComponent, prometheusAdapter MonitoringComponent, ) *MonitoringConfig`

NewMonitoringConfig instantiates a new MonitoringConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringConfigWithDefaults

`func NewMonitoringConfigWithDefaults() *MonitoringConfig`

NewMonitoringConfigWithDefaults instantiates a new MonitoringConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHelmExporter

`func (o *MonitoringConfig) GetHelmExporter() MonitoringComponent`

GetHelmExporter returns the HelmExporter field if non-nil, zero value otherwise.

### GetHelmExporterOk

`func (o *MonitoringConfig) GetHelmExporterOk() (*MonitoringComponent, bool)`

GetHelmExporterOk returns a tuple with the HelmExporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmExporter

`func (o *MonitoringConfig) SetHelmExporter(v MonitoringComponent)`

SetHelmExporter sets HelmExporter field to given value.


### GetKubeStateMetrics

`func (o *MonitoringConfig) GetKubeStateMetrics() MonitoringComponent`

GetKubeStateMetrics returns the KubeStateMetrics field if non-nil, zero value otherwise.

### GetKubeStateMetricsOk

`func (o *MonitoringConfig) GetKubeStateMetricsOk() (*MonitoringComponent, bool)`

GetKubeStateMetricsOk returns a tuple with the KubeStateMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeStateMetrics

`func (o *MonitoringConfig) SetKubeStateMetrics(v MonitoringComponent)`

SetKubeStateMetrics sets KubeStateMetrics field to given value.


### GetMetricsServer

`func (o *MonitoringConfig) GetMetricsServer() MonitoringComponent`

GetMetricsServer returns the MetricsServer field if non-nil, zero value otherwise.

### GetMetricsServerOk

`func (o *MonitoringConfig) GetMetricsServerOk() (*MonitoringComponent, bool)`

GetMetricsServerOk returns a tuple with the MetricsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsServer

`func (o *MonitoringConfig) SetMetricsServer(v MonitoringComponent)`

SetMetricsServer sets MetricsServer field to given value.


### GetNodeExporter

`func (o *MonitoringConfig) GetNodeExporter() MonitoringComponent`

GetNodeExporter returns the NodeExporter field if non-nil, zero value otherwise.

### GetNodeExporterOk

`func (o *MonitoringConfig) GetNodeExporterOk() (*MonitoringComponent, bool)`

GetNodeExporterOk returns a tuple with the NodeExporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporter

`func (o *MonitoringConfig) SetNodeExporter(v MonitoringComponent)`

SetNodeExporter sets NodeExporter field to given value.


### GetPrometheusAdapter

`func (o *MonitoringConfig) GetPrometheusAdapter() MonitoringComponent`

GetPrometheusAdapter returns the PrometheusAdapter field if non-nil, zero value otherwise.

### GetPrometheusAdapterOk

`func (o *MonitoringConfig) GetPrometheusAdapterOk() (*MonitoringComponent, bool)`

GetPrometheusAdapterOk returns a tuple with the PrometheusAdapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusAdapter

`func (o *MonitoringConfig) SetPrometheusAdapter(v MonitoringComponent)`

SetPrometheusAdapter sets PrometheusAdapter field to given value.


### GetResources

`func (o *MonitoringConfig) GetResources() ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *MonitoringConfig) GetResourcesOk() (*ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *MonitoringConfig) SetResources(v ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *MonitoringConfig) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


