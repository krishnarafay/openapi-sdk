# AzureCostProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomPricing** | Pointer to [**AzureCustomPricing**](AzureCustomPricing.md) |  | [optional] 
**GpuEstimates** | Pointer to [**GpuCostProfile**](GpuCostProfile.md) |  | [optional] 

## Methods

### NewAzureCostProfile

`func NewAzureCostProfile() *AzureCostProfile`

NewAzureCostProfile instantiates a new AzureCostProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCostProfileWithDefaults

`func NewAzureCostProfileWithDefaults() *AzureCostProfile`

NewAzureCostProfileWithDefaults instantiates a new AzureCostProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomPricing

`func (o *AzureCostProfile) GetCustomPricing() AzureCustomPricing`

GetCustomPricing returns the CustomPricing field if non-nil, zero value otherwise.

### GetCustomPricingOk

`func (o *AzureCostProfile) GetCustomPricingOk() (*AzureCustomPricing, bool)`

GetCustomPricingOk returns a tuple with the CustomPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPricing

`func (o *AzureCostProfile) SetCustomPricing(v AzureCustomPricing)`

SetCustomPricing sets CustomPricing field to given value.

### HasCustomPricing

`func (o *AzureCostProfile) HasCustomPricing() bool`

HasCustomPricing returns a boolean if a field has been set.

### GetGpuEstimates

`func (o *AzureCostProfile) GetGpuEstimates() GpuCostProfile`

GetGpuEstimates returns the GpuEstimates field if non-nil, zero value otherwise.

### GetGpuEstimatesOk

`func (o *AzureCostProfile) GetGpuEstimatesOk() (*GpuCostProfile, bool)`

GetGpuEstimatesOk returns a tuple with the GpuEstimates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuEstimates

`func (o *AzureCostProfile) SetGpuEstimates(v GpuCostProfile)`

SetGpuEstimates sets GpuEstimates field to given value.

### HasGpuEstimates

`func (o *AzureCostProfile) HasGpuEstimates() bool`

HasGpuEstimates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


