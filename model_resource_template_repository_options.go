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

// checks if the ResourceTemplateRepositoryOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceTemplateRepositoryOptions{}

// ResourceTemplateRepositoryOptions struct for ResourceTemplateRepositoryOptions
type ResourceTemplateRepositoryOptions struct {
	// Specify the branch
	Branch *string `json:"branch,omitempty"`
	// Specify the directory path
	DirectoryPath *string `json:"directory_path,omitempty"`
	// Specify the name of the repository
	Name *string `json:"name,omitempty"`
}

// NewResourceTemplateRepositoryOptions instantiates a new ResourceTemplateRepositoryOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceTemplateRepositoryOptions() *ResourceTemplateRepositoryOptions {
	this := ResourceTemplateRepositoryOptions{}
	return &this
}

// NewResourceTemplateRepositoryOptionsWithDefaults instantiates a new ResourceTemplateRepositoryOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceTemplateRepositoryOptionsWithDefaults() *ResourceTemplateRepositoryOptions {
	this := ResourceTemplateRepositoryOptions{}
	return &this
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *ResourceTemplateRepositoryOptions) GetBranch() string {
	if o == nil || isNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTemplateRepositoryOptions) GetBranchOk() (*string, bool) {
	if o == nil || isNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *ResourceTemplateRepositoryOptions) HasBranch() bool {
	if o != nil && !isNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *ResourceTemplateRepositoryOptions) SetBranch(v string) {
	o.Branch = &v
}

// GetDirectoryPath returns the DirectoryPath field value if set, zero value otherwise.
func (o *ResourceTemplateRepositoryOptions) GetDirectoryPath() string {
	if o == nil || isNil(o.DirectoryPath) {
		var ret string
		return ret
	}
	return *o.DirectoryPath
}

// GetDirectoryPathOk returns a tuple with the DirectoryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTemplateRepositoryOptions) GetDirectoryPathOk() (*string, bool) {
	if o == nil || isNil(o.DirectoryPath) {
		return nil, false
	}
	return o.DirectoryPath, true
}

// HasDirectoryPath returns a boolean if a field has been set.
func (o *ResourceTemplateRepositoryOptions) HasDirectoryPath() bool {
	if o != nil && !isNil(o.DirectoryPath) {
		return true
	}

	return false
}

// SetDirectoryPath gets a reference to the given string and assigns it to the DirectoryPath field.
func (o *ResourceTemplateRepositoryOptions) SetDirectoryPath(v string) {
	o.DirectoryPath = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceTemplateRepositoryOptions) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTemplateRepositoryOptions) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceTemplateRepositoryOptions) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceTemplateRepositoryOptions) SetName(v string) {
	o.Name = &v
}

func (o ResourceTemplateRepositoryOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceTemplateRepositoryOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !isNil(o.DirectoryPath) {
		toSerialize["directory_path"] = o.DirectoryPath
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableResourceTemplateRepositoryOptions struct {
	value *ResourceTemplateRepositoryOptions
	isSet bool
}

func (v NullableResourceTemplateRepositoryOptions) Get() *ResourceTemplateRepositoryOptions {
	return v.value
}

func (v *NullableResourceTemplateRepositoryOptions) Set(val *ResourceTemplateRepositoryOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTemplateRepositoryOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTemplateRepositoryOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTemplateRepositoryOptions(val *ResourceTemplateRepositoryOptions) *NullableResourceTemplateRepositoryOptions {
	return &NullableResourceTemplateRepositoryOptions{value: val, isSet: true}
}

func (v NullableResourceTemplateRepositoryOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTemplateRepositoryOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

