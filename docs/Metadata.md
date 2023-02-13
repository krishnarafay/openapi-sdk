# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** | annotations of the resource | [optional] 
**Description** | Pointer to **string** | description of the resource | [optional] 
**Labels** | Pointer to **map[string]string** | labels of the resource | [optional] 
**Name** | **string** | name of the resource | [default to "example-cluster"]
**Project** | **string** | Project of the resource | [default to "defaultproject"]

## Methods

### NewMetadata

`func NewMetadata(name string, project string, ) *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *Metadata) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Metadata) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Metadata) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Metadata) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetDescription

`func (o *Metadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Metadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Metadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Metadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *Metadata) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Metadata) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Metadata) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Metadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *Metadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Metadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Metadata) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *Metadata) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Metadata) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Metadata) SetProject(v string)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


