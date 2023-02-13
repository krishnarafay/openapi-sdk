# Sku

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of a managed cluster SKU. | [optional] [default to "Basic"]
**Tier** | Pointer to **string** | If not specified, the default is Free. See uptime SLA for more details. Valid values are Paid, Free. | [optional] [default to "Free"]

## Methods

### NewSku

`func NewSku() *Sku`

NewSku instantiates a new Sku object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuWithDefaults

`func NewSkuWithDefaults() *Sku`

NewSkuWithDefaults instantiates a new Sku object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Sku) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sku) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sku) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Sku) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTier

`func (o *Sku) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *Sku) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *Sku) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *Sku) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


