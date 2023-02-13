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

// checks if the ClusterNetworkPolicyRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterNetworkPolicyRule{}

// ClusterNetworkPolicyRule ClusterNetworkPolicyRule  definition
type ClusterNetworkPolicyRule struct {
	// API Version of the cluster network policy rule resource
	ApiVersion string `json:"apiVersion"`
	// Kind of the cluster network policy rule resource
	Kind string `json:"kind"`
	Metadata Metadata `json:"metadata"`
	Spec ClusterNetworkPolicyRuleSpec `json:"spec"`
}

// NewClusterNetworkPolicyRule instantiates a new ClusterNetworkPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterNetworkPolicyRule(apiVersion string, kind string, metadata Metadata, spec ClusterNetworkPolicyRuleSpec) *ClusterNetworkPolicyRule {
	this := ClusterNetworkPolicyRule{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewClusterNetworkPolicyRuleWithDefaults instantiates a new ClusterNetworkPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterNetworkPolicyRuleWithDefaults() *ClusterNetworkPolicyRule {
	this := ClusterNetworkPolicyRule{}
	var apiVersion string = "security.k8smgmt.io/v3"
	this.ApiVersion = apiVersion
	var kind string = "ClusterNetworkPolicyRule"
	this.Kind = kind
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *ClusterNetworkPolicyRule) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *ClusterNetworkPolicyRule) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *ClusterNetworkPolicyRule) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *ClusterNetworkPolicyRule) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ClusterNetworkPolicyRule) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ClusterNetworkPolicyRule) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *ClusterNetworkPolicyRule) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ClusterNetworkPolicyRule) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ClusterNetworkPolicyRule) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *ClusterNetworkPolicyRule) GetSpec() ClusterNetworkPolicyRuleSpec {
	if o == nil {
		var ret ClusterNetworkPolicyRuleSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *ClusterNetworkPolicyRule) GetSpecOk() (*ClusterNetworkPolicyRuleSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *ClusterNetworkPolicyRule) SetSpec(v ClusterNetworkPolicyRuleSpec) {
	o.Spec = v
}

func (o ClusterNetworkPolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterNetworkPolicyRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	toSerialize["spec"] = o.Spec
	return toSerialize, nil
}

type NullableClusterNetworkPolicyRule struct {
	value *ClusterNetworkPolicyRule
	isSet bool
}

func (v NullableClusterNetworkPolicyRule) Get() *ClusterNetworkPolicyRule {
	return v.value
}

func (v *NullableClusterNetworkPolicyRule) Set(val *ClusterNetworkPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterNetworkPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterNetworkPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterNetworkPolicyRule(val *ClusterNetworkPolicyRule) *NullableClusterNetworkPolicyRule {
	return &NullableClusterNetworkPolicyRule{value: val, isSet: true}
}

func (v NullableClusterNetworkPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterNetworkPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


