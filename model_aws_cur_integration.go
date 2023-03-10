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

// checks if the AwsCurIntegration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsCurIntegration{}

// AwsCurIntegration Aws custom cost profile params
type AwsCurIntegration struct {
	// Aws Athena Bucket Name
	AthenaBucketName *string `json:"athenaBucketName,omitempty"`
	// Aws Athena Database
	AthenaDatabase *string `json:"athenaDatabase,omitempty"`
	// Aws Athena Region
	AthenaRegion *string `json:"athenaRegion,omitempty"`
	// Aws Athena Table Name
	AthenaTable *string `json:"athenaTable,omitempty"`
	// Aws Account ID
	AwsAccountId *string `json:"awsAccountId,omitempty"`
	// Master Payer Arn
	MasterPayerArn *string `json:"masterPayerArn,omitempty"`
}

// NewAwsCurIntegration instantiates a new AwsCurIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsCurIntegration() *AwsCurIntegration {
	this := AwsCurIntegration{}
	return &this
}

// NewAwsCurIntegrationWithDefaults instantiates a new AwsCurIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsCurIntegrationWithDefaults() *AwsCurIntegration {
	this := AwsCurIntegration{}
	return &this
}

// GetAthenaBucketName returns the AthenaBucketName field value if set, zero value otherwise.
func (o *AwsCurIntegration) GetAthenaBucketName() string {
	if o == nil || isNil(o.AthenaBucketName) {
		var ret string
		return ret
	}
	return *o.AthenaBucketName
}

// GetAthenaBucketNameOk returns a tuple with the AthenaBucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurIntegration) GetAthenaBucketNameOk() (*string, bool) {
	if o == nil || isNil(o.AthenaBucketName) {
		return nil, false
	}
	return o.AthenaBucketName, true
}

// HasAthenaBucketName returns a boolean if a field has been set.
func (o *AwsCurIntegration) HasAthenaBucketName() bool {
	if o != nil && !isNil(o.AthenaBucketName) {
		return true
	}

	return false
}

// SetAthenaBucketName gets a reference to the given string and assigns it to the AthenaBucketName field.
func (o *AwsCurIntegration) SetAthenaBucketName(v string) {
	o.AthenaBucketName = &v
}

// GetAthenaDatabase returns the AthenaDatabase field value if set, zero value otherwise.
func (o *AwsCurIntegration) GetAthenaDatabase() string {
	if o == nil || isNil(o.AthenaDatabase) {
		var ret string
		return ret
	}
	return *o.AthenaDatabase
}

// GetAthenaDatabaseOk returns a tuple with the AthenaDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurIntegration) GetAthenaDatabaseOk() (*string, bool) {
	if o == nil || isNil(o.AthenaDatabase) {
		return nil, false
	}
	return o.AthenaDatabase, true
}

// HasAthenaDatabase returns a boolean if a field has been set.
func (o *AwsCurIntegration) HasAthenaDatabase() bool {
	if o != nil && !isNil(o.AthenaDatabase) {
		return true
	}

	return false
}

// SetAthenaDatabase gets a reference to the given string and assigns it to the AthenaDatabase field.
func (o *AwsCurIntegration) SetAthenaDatabase(v string) {
	o.AthenaDatabase = &v
}

// GetAthenaRegion returns the AthenaRegion field value if set, zero value otherwise.
func (o *AwsCurIntegration) GetAthenaRegion() string {
	if o == nil || isNil(o.AthenaRegion) {
		var ret string
		return ret
	}
	return *o.AthenaRegion
}

// GetAthenaRegionOk returns a tuple with the AthenaRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurIntegration) GetAthenaRegionOk() (*string, bool) {
	if o == nil || isNil(o.AthenaRegion) {
		return nil, false
	}
	return o.AthenaRegion, true
}

// HasAthenaRegion returns a boolean if a field has been set.
func (o *AwsCurIntegration) HasAthenaRegion() bool {
	if o != nil && !isNil(o.AthenaRegion) {
		return true
	}

	return false
}

// SetAthenaRegion gets a reference to the given string and assigns it to the AthenaRegion field.
func (o *AwsCurIntegration) SetAthenaRegion(v string) {
	o.AthenaRegion = &v
}

// GetAthenaTable returns the AthenaTable field value if set, zero value otherwise.
func (o *AwsCurIntegration) GetAthenaTable() string {
	if o == nil || isNil(o.AthenaTable) {
		var ret string
		return ret
	}
	return *o.AthenaTable
}

// GetAthenaTableOk returns a tuple with the AthenaTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurIntegration) GetAthenaTableOk() (*string, bool) {
	if o == nil || isNil(o.AthenaTable) {
		return nil, false
	}
	return o.AthenaTable, true
}

// HasAthenaTable returns a boolean if a field has been set.
func (o *AwsCurIntegration) HasAthenaTable() bool {
	if o != nil && !isNil(o.AthenaTable) {
		return true
	}

	return false
}

// SetAthenaTable gets a reference to the given string and assigns it to the AthenaTable field.
func (o *AwsCurIntegration) SetAthenaTable(v string) {
	o.AthenaTable = &v
}

// GetAwsAccountId returns the AwsAccountId field value if set, zero value otherwise.
func (o *AwsCurIntegration) GetAwsAccountId() string {
	if o == nil || isNil(o.AwsAccountId) {
		var ret string
		return ret
	}
	return *o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurIntegration) GetAwsAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AwsAccountId) {
		return nil, false
	}
	return o.AwsAccountId, true
}

// HasAwsAccountId returns a boolean if a field has been set.
func (o *AwsCurIntegration) HasAwsAccountId() bool {
	if o != nil && !isNil(o.AwsAccountId) {
		return true
	}

	return false
}

// SetAwsAccountId gets a reference to the given string and assigns it to the AwsAccountId field.
func (o *AwsCurIntegration) SetAwsAccountId(v string) {
	o.AwsAccountId = &v
}

// GetMasterPayerArn returns the MasterPayerArn field value if set, zero value otherwise.
func (o *AwsCurIntegration) GetMasterPayerArn() string {
	if o == nil || isNil(o.MasterPayerArn) {
		var ret string
		return ret
	}
	return *o.MasterPayerArn
}

// GetMasterPayerArnOk returns a tuple with the MasterPayerArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurIntegration) GetMasterPayerArnOk() (*string, bool) {
	if o == nil || isNil(o.MasterPayerArn) {
		return nil, false
	}
	return o.MasterPayerArn, true
}

// HasMasterPayerArn returns a boolean if a field has been set.
func (o *AwsCurIntegration) HasMasterPayerArn() bool {
	if o != nil && !isNil(o.MasterPayerArn) {
		return true
	}

	return false
}

// SetMasterPayerArn gets a reference to the given string and assigns it to the MasterPayerArn field.
func (o *AwsCurIntegration) SetMasterPayerArn(v string) {
	o.MasterPayerArn = &v
}

func (o AwsCurIntegration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsCurIntegration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AthenaBucketName) {
		toSerialize["athenaBucketName"] = o.AthenaBucketName
	}
	if !isNil(o.AthenaDatabase) {
		toSerialize["athenaDatabase"] = o.AthenaDatabase
	}
	if !isNil(o.AthenaRegion) {
		toSerialize["athenaRegion"] = o.AthenaRegion
	}
	if !isNil(o.AthenaTable) {
		toSerialize["athenaTable"] = o.AthenaTable
	}
	if !isNil(o.AwsAccountId) {
		toSerialize["awsAccountId"] = o.AwsAccountId
	}
	if !isNil(o.MasterPayerArn) {
		toSerialize["masterPayerArn"] = o.MasterPayerArn
	}
	return toSerialize, nil
}

type NullableAwsCurIntegration struct {
	value *AwsCurIntegration
	isSet bool
}

func (v NullableAwsCurIntegration) Get() *AwsCurIntegration {
	return v.value
}

func (v *NullableAwsCurIntegration) Set(val *AwsCurIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsCurIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsCurIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsCurIntegration(val *AwsCurIntegration) *NullableAwsCurIntegration {
	return &NullableAwsCurIntegration{value: val, isSet: true}
}

func (v NullableAwsCurIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsCurIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


