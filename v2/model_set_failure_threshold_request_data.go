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
)

// SetFailureThresholdRequestData struct for SetFailureThresholdRequestData
type SetFailureThresholdRequestData struct {
	// The minimum number of replicas required to be online and receiving writes in order for the volume to remain read-writable. This value replaces any previously set failure threshold or intent-based failure mode. 
	FailureThreshold *uint64 `json:"failureThreshold,omitempty"`
	// An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change. 
	Version *string `json:"version,omitempty"`
}

// NewSetFailureThresholdRequestData instantiates a new SetFailureThresholdRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetFailureThresholdRequestData() *SetFailureThresholdRequestData {
	this := SetFailureThresholdRequestData{}
	return &this
}

// NewSetFailureThresholdRequestDataWithDefaults instantiates a new SetFailureThresholdRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetFailureThresholdRequestDataWithDefaults() *SetFailureThresholdRequestData {
	this := SetFailureThresholdRequestData{}
	return &this
}

// GetFailureThreshold returns the FailureThreshold field value if set, zero value otherwise.
func (o *SetFailureThresholdRequestData) GetFailureThreshold() uint64 {
	if o == nil || o.FailureThreshold == nil {
		var ret uint64
		return ret
	}
	return *o.FailureThreshold
}

// GetFailureThresholdOk returns a tuple with the FailureThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetFailureThresholdRequestData) GetFailureThresholdOk() (*uint64, bool) {
	if o == nil || o.FailureThreshold == nil {
		return nil, false
	}
	return o.FailureThreshold, true
}

// HasFailureThreshold returns a boolean if a field has been set.
func (o *SetFailureThresholdRequestData) HasFailureThreshold() bool {
	if o != nil && o.FailureThreshold != nil {
		return true
	}

	return false
}

// SetFailureThreshold gets a reference to the given uint64 and assigns it to the FailureThreshold field.
func (o *SetFailureThresholdRequestData) SetFailureThreshold(v uint64) {
	o.FailureThreshold = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SetFailureThresholdRequestData) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetFailureThresholdRequestData) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SetFailureThresholdRequestData) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SetFailureThresholdRequestData) SetVersion(v string) {
	o.Version = &v
}

func (o SetFailureThresholdRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FailureThreshold != nil {
		toSerialize["failureThreshold"] = o.FailureThreshold
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableSetFailureThresholdRequestData struct {
	value *SetFailureThresholdRequestData
	isSet bool
}

func (v NullableSetFailureThresholdRequestData) Get() *SetFailureThresholdRequestData {
	return v.value
}

func (v *NullableSetFailureThresholdRequestData) Set(val *SetFailureThresholdRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSetFailureThresholdRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSetFailureThresholdRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetFailureThresholdRequestData(val *SetFailureThresholdRequestData) *NullableSetFailureThresholdRequestData {
	return &NullableSetFailureThresholdRequestData{value: val, isSet: true}
}

func (v NullableSetFailureThresholdRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetFailureThresholdRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


