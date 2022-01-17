/*
StorageOS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.5.0
Contact: info@storageos.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// PolicyGroup struct for PolicyGroup
type PolicyGroup struct {
	// A unique identifier for a policy group. The format of this type is undefined and may change but the defined properties will not change.. 
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// The list of user IDs which this policy group governs.
	Users *[]PolicyGroupUsers `json:"users,omitempty"`
	// A set of authorisation policies to apply to the policy group.
	Specs []PoliciesIdSpecs `json:"specs,omitempty"`
	// The time the entity was created. This timestamp is set by the node that created the entity, and may not be correct if the node's local clock was skewed. This value is for the user's informative purposes only, and correctness is not required. String format is RFC3339. 
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The time the entity was last updated. This timestamp is set by the node that last updated the entity, and may not be correct if the node's local clock was skewed. This value is for the user's informative purposes only, and correctness is not required. String format is RFC3339. 
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change. 
	Version *string `json:"version,omitempty"`
}

// NewPolicyGroup instantiates a new PolicyGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyGroup() *PolicyGroup {
	this := PolicyGroup{}
	return &this
}

// NewPolicyGroupWithDefaults instantiates a new PolicyGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyGroupWithDefaults() *PolicyGroup {
	this := PolicyGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PolicyGroup) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyGroup) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PolicyGroup) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PolicyGroup) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PolicyGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PolicyGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PolicyGroup) SetName(v string) {
	o.Name = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *PolicyGroup) GetUsers() []PolicyGroupUsers {
	if o == nil || o.Users == nil {
		var ret []PolicyGroupUsers
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyGroup) GetUsersOk() (*[]PolicyGroupUsers, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *PolicyGroup) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []PolicyGroupUsers and assigns it to the Users field.
func (o *PolicyGroup) SetUsers(v []PolicyGroupUsers) {
	o.Users = &v
}

// GetSpecs returns the Specs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyGroup) GetSpecs() []PoliciesIdSpecs {
	if o == nil  {
		var ret []PoliciesIdSpecs
		return ret
	}
	return o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyGroup) GetSpecsOk() (*[]PoliciesIdSpecs, bool) {
	if o == nil || o.Specs == nil {
		return nil, false
	}
	return &o.Specs, true
}

// HasSpecs returns a boolean if a field has been set.
func (o *PolicyGroup) HasSpecs() bool {
	if o != nil && o.Specs != nil {
		return true
	}

	return false
}

// SetSpecs gets a reference to the given []PoliciesIdSpecs and assigns it to the Specs field.
func (o *PolicyGroup) SetSpecs(v []PoliciesIdSpecs) {
	o.Specs = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PolicyGroup) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyGroup) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PolicyGroup) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PolicyGroup) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PolicyGroup) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyGroup) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PolicyGroup) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PolicyGroup) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PolicyGroup) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyGroup) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PolicyGroup) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PolicyGroup) SetVersion(v string) {
	o.Version = &v
}

func (o PolicyGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Specs != nil {
		toSerialize["specs"] = o.Specs
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyGroup struct {
	value *PolicyGroup
	isSet bool
}

func (v NullablePolicyGroup) Get() *PolicyGroup {
	return v.value
}

func (v *NullablePolicyGroup) Set(val *PolicyGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyGroup(val *PolicyGroup) *NullablePolicyGroup {
	return &NullablePolicyGroup{value: val, isSet: true}
}

func (v NullablePolicyGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


