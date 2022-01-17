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

// SetFailureModeIntentRequestData struct for SetFailureModeIntentRequestData
type SetFailureModeIntentRequestData struct {
	Mode *FailureModeIntent `json:"mode,omitempty"`
	// An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change. 
	Version *string `json:"version,omitempty"`
}

// NewSetFailureModeIntentRequestData instantiates a new SetFailureModeIntentRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetFailureModeIntentRequestData() *SetFailureModeIntentRequestData {
	this := SetFailureModeIntentRequestData{}
	return &this
}

// NewSetFailureModeIntentRequestDataWithDefaults instantiates a new SetFailureModeIntentRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetFailureModeIntentRequestDataWithDefaults() *SetFailureModeIntentRequestData {
	this := SetFailureModeIntentRequestData{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *SetFailureModeIntentRequestData) GetMode() FailureModeIntent {
	if o == nil || o.Mode == nil {
		var ret FailureModeIntent
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetFailureModeIntentRequestData) GetModeOk() (*FailureModeIntent, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *SetFailureModeIntentRequestData) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given FailureModeIntent and assigns it to the Mode field.
func (o *SetFailureModeIntentRequestData) SetMode(v FailureModeIntent) {
	o.Mode = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SetFailureModeIntentRequestData) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetFailureModeIntentRequestData) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SetFailureModeIntentRequestData) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SetFailureModeIntentRequestData) SetVersion(v string) {
	o.Version = &v
}

func (o SetFailureModeIntentRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableSetFailureModeIntentRequestData struct {
	value *SetFailureModeIntentRequestData
	isSet bool
}

func (v NullableSetFailureModeIntentRequestData) Get() *SetFailureModeIntentRequestData {
	return v.value
}

func (v *NullableSetFailureModeIntentRequestData) Set(val *SetFailureModeIntentRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSetFailureModeIntentRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSetFailureModeIntentRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetFailureModeIntentRequestData(val *SetFailureModeIntentRequestData) *NullableSetFailureModeIntentRequestData {
	return &NullableSetFailureModeIntentRequestData{value: val, isSet: true}
}

func (v NullableSetFailureModeIntentRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetFailureModeIntentRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


