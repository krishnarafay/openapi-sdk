/*
Kubernetes Operations APIs

Kubernetes Operations APIs

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rafaysdkgov3

import (
	"encoding/json"
)

// checks if the ClusterStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterStatus{}

// ClusterStatus cluster status
type ClusterStatus struct {
	Aks *AKSStatus `json:"aks,omitempty"`
	Blueprint *BlueprintStatus `json:"blueprint,omitempty"`
	CommonStatus *Status `json:"commonStatus,omitempty"`
	Conditions []ClusterCondition `json:"conditions,omitempty"`
	ControPlane *ControlPlaneStatus `json:"controPlane,omitempty"`
	CreatedAt *StatusTime `json:"createdAt,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Extra *ExtraStatus `json:"extra,omitempty"`
	Id *string `json:"id,omitempty"`
	Imported *ImportedStatus `json:"imported,omitempty"`
	Name *string `json:"name,omitempty"`
	ProvisionStatus *string `json:"provisionStatus,omitempty"`
}

// NewClusterStatus instantiates a new ClusterStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterStatus() *ClusterStatus {
	this := ClusterStatus{}
	return &this
}

// NewClusterStatusWithDefaults instantiates a new ClusterStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterStatusWithDefaults() *ClusterStatus {
	this := ClusterStatus{}
	return &this
}

// GetAks returns the Aks field value if set, zero value otherwise.
func (o *ClusterStatus) GetAks() AKSStatus {
	if o == nil || isNil(o.Aks) {
		var ret AKSStatus
		return ret
	}
	return *o.Aks
}

// GetAksOk returns a tuple with the Aks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetAksOk() (*AKSStatus, bool) {
	if o == nil || isNil(o.Aks) {
		return nil, false
	}
	return o.Aks, true
}

// HasAks returns a boolean if a field has been set.
func (o *ClusterStatus) HasAks() bool {
	if o != nil && !isNil(o.Aks) {
		return true
	}

	return false
}

// SetAks gets a reference to the given AKSStatus and assigns it to the Aks field.
func (o *ClusterStatus) SetAks(v AKSStatus) {
	o.Aks = &v
}

// GetBlueprint returns the Blueprint field value if set, zero value otherwise.
func (o *ClusterStatus) GetBlueprint() BlueprintStatus {
	if o == nil || isNil(o.Blueprint) {
		var ret BlueprintStatus
		return ret
	}
	return *o.Blueprint
}

// GetBlueprintOk returns a tuple with the Blueprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetBlueprintOk() (*BlueprintStatus, bool) {
	if o == nil || isNil(o.Blueprint) {
		return nil, false
	}
	return o.Blueprint, true
}

// HasBlueprint returns a boolean if a field has been set.
func (o *ClusterStatus) HasBlueprint() bool {
	if o != nil && !isNil(o.Blueprint) {
		return true
	}

	return false
}

// SetBlueprint gets a reference to the given BlueprintStatus and assigns it to the Blueprint field.
func (o *ClusterStatus) SetBlueprint(v BlueprintStatus) {
	o.Blueprint = &v
}

// GetCommonStatus returns the CommonStatus field value if set, zero value otherwise.
func (o *ClusterStatus) GetCommonStatus() Status {
	if o == nil || isNil(o.CommonStatus) {
		var ret Status
		return ret
	}
	return *o.CommonStatus
}

// GetCommonStatusOk returns a tuple with the CommonStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetCommonStatusOk() (*Status, bool) {
	if o == nil || isNil(o.CommonStatus) {
		return nil, false
	}
	return o.CommonStatus, true
}

// HasCommonStatus returns a boolean if a field has been set.
func (o *ClusterStatus) HasCommonStatus() bool {
	if o != nil && !isNil(o.CommonStatus) {
		return true
	}

	return false
}

// SetCommonStatus gets a reference to the given Status and assigns it to the CommonStatus field.
func (o *ClusterStatus) SetCommonStatus(v Status) {
	o.CommonStatus = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ClusterStatus) GetConditions() []ClusterCondition {
	if o == nil || isNil(o.Conditions) {
		var ret []ClusterCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetConditionsOk() ([]ClusterCondition, bool) {
	if o == nil || isNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ClusterStatus) HasConditions() bool {
	if o != nil && !isNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ClusterCondition and assigns it to the Conditions field.
func (o *ClusterStatus) SetConditions(v []ClusterCondition) {
	o.Conditions = v
}

// GetControPlane returns the ControPlane field value if set, zero value otherwise.
func (o *ClusterStatus) GetControPlane() ControlPlaneStatus {
	if o == nil || isNil(o.ControPlane) {
		var ret ControlPlaneStatus
		return ret
	}
	return *o.ControPlane
}

// GetControPlaneOk returns a tuple with the ControPlane field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetControPlaneOk() (*ControlPlaneStatus, bool) {
	if o == nil || isNil(o.ControPlane) {
		return nil, false
	}
	return o.ControPlane, true
}

// HasControPlane returns a boolean if a field has been set.
func (o *ClusterStatus) HasControPlane() bool {
	if o != nil && !isNil(o.ControPlane) {
		return true
	}

	return false
}

// SetControPlane gets a reference to the given ControlPlaneStatus and assigns it to the ControPlane field.
func (o *ClusterStatus) SetControPlane(v ControlPlaneStatus) {
	o.ControPlane = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ClusterStatus) GetCreatedAt() StatusTime {
	if o == nil || isNil(o.CreatedAt) {
		var ret StatusTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetCreatedAtOk() (*StatusTime, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ClusterStatus) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given StatusTime and assigns it to the CreatedAt field.
func (o *ClusterStatus) SetCreatedAt(v StatusTime) {
	o.CreatedAt = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ClusterStatus) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ClusterStatus) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ClusterStatus) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *ClusterStatus) GetExtra() ExtraStatus {
	if o == nil || isNil(o.Extra) {
		var ret ExtraStatus
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetExtraOk() (*ExtraStatus, bool) {
	if o == nil || isNil(o.Extra) {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *ClusterStatus) HasExtra() bool {
	if o != nil && !isNil(o.Extra) {
		return true
	}

	return false
}

// SetExtra gets a reference to the given ExtraStatus and assigns it to the Extra field.
func (o *ClusterStatus) SetExtra(v ExtraStatus) {
	o.Extra = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterStatus) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterStatus) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClusterStatus) SetId(v string) {
	o.Id = &v
}

// GetImported returns the Imported field value if set, zero value otherwise.
func (o *ClusterStatus) GetImported() ImportedStatus {
	if o == nil || isNil(o.Imported) {
		var ret ImportedStatus
		return ret
	}
	return *o.Imported
}

// GetImportedOk returns a tuple with the Imported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetImportedOk() (*ImportedStatus, bool) {
	if o == nil || isNil(o.Imported) {
		return nil, false
	}
	return o.Imported, true
}

// HasImported returns a boolean if a field has been set.
func (o *ClusterStatus) HasImported() bool {
	if o != nil && !isNil(o.Imported) {
		return true
	}

	return false
}

// SetImported gets a reference to the given ImportedStatus and assigns it to the Imported field.
func (o *ClusterStatus) SetImported(v ImportedStatus) {
	o.Imported = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterStatus) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterStatus) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterStatus) SetName(v string) {
	o.Name = &v
}

// GetProvisionStatus returns the ProvisionStatus field value if set, zero value otherwise.
func (o *ClusterStatus) GetProvisionStatus() string {
	if o == nil || isNil(o.ProvisionStatus) {
		var ret string
		return ret
	}
	return *o.ProvisionStatus
}

// GetProvisionStatusOk returns a tuple with the ProvisionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetProvisionStatusOk() (*string, bool) {
	if o == nil || isNil(o.ProvisionStatus) {
		return nil, false
	}
	return o.ProvisionStatus, true
}

// HasProvisionStatus returns a boolean if a field has been set.
func (o *ClusterStatus) HasProvisionStatus() bool {
	if o != nil && !isNil(o.ProvisionStatus) {
		return true
	}

	return false
}

// SetProvisionStatus gets a reference to the given string and assigns it to the ProvisionStatus field.
func (o *ClusterStatus) SetProvisionStatus(v string) {
	o.ProvisionStatus = &v
}

func (o ClusterStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Aks) {
		toSerialize["aks"] = o.Aks
	}
	if !isNil(o.Blueprint) {
		toSerialize["blueprint"] = o.Blueprint
	}
	if !isNil(o.CommonStatus) {
		toSerialize["commonStatus"] = o.CommonStatus
	}
	if !isNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !isNil(o.ControPlane) {
		toSerialize["controPlane"] = o.ControPlane
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Extra) {
		toSerialize["extra"] = o.Extra
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Imported) {
		toSerialize["imported"] = o.Imported
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ProvisionStatus) {
		toSerialize["provisionStatus"] = o.ProvisionStatus
	}
	return toSerialize, nil
}

type NullableClusterStatus struct {
	value *ClusterStatus
	isSet bool
}

func (v NullableClusterStatus) Get() *ClusterStatus {
	return v.value
}

func (v *NullableClusterStatus) Set(val *ClusterStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStatus(val *ClusterStatus) *NullableClusterStatus {
	return &NullableClusterStatus{value: val, isSet: true}
}

func (v NullableClusterStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

