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

// MasterDeploymentInfo struct for MasterDeploymentInfo
type MasterDeploymentInfo struct {
	// A unique identifier for a volume deployment. The format of this type is undefined and may change but the defined properties will not change. 
	Id *string `json:"id,omitempty"`
	// A unique identifier for a node. The format of this type is undefined and may change but the defined properties will not change. 
	NodeID *string `json:"nodeID,omitempty"`
	// The hostname of the node that is hosting the deployment 
	Hostname *string `json:"hostname,omitempty"`
	// Indicates if the volume deployment is eligible for promotion 
	Promotable *bool `json:"promotable,omitempty"`
	Health *MasterHealth `json:"health,omitempty"`
}

// NewMasterDeploymentInfo instantiates a new MasterDeploymentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMasterDeploymentInfo() *MasterDeploymentInfo {
	this := MasterDeploymentInfo{}
	return &this
}

// NewMasterDeploymentInfoWithDefaults instantiates a new MasterDeploymentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMasterDeploymentInfoWithDefaults() *MasterDeploymentInfo {
	this := MasterDeploymentInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MasterDeploymentInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterDeploymentInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MasterDeploymentInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MasterDeploymentInfo) SetId(v string) {
	o.Id = &v
}

// GetNodeID returns the NodeID field value if set, zero value otherwise.
func (o *MasterDeploymentInfo) GetNodeID() string {
	if o == nil || o.NodeID == nil {
		var ret string
		return ret
	}
	return *o.NodeID
}

// GetNodeIDOk returns a tuple with the NodeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterDeploymentInfo) GetNodeIDOk() (*string, bool) {
	if o == nil || o.NodeID == nil {
		return nil, false
	}
	return o.NodeID, true
}

// HasNodeID returns a boolean if a field has been set.
func (o *MasterDeploymentInfo) HasNodeID() bool {
	if o != nil && o.NodeID != nil {
		return true
	}

	return false
}

// SetNodeID gets a reference to the given string and assigns it to the NodeID field.
func (o *MasterDeploymentInfo) SetNodeID(v string) {
	o.NodeID = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *MasterDeploymentInfo) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterDeploymentInfo) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *MasterDeploymentInfo) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *MasterDeploymentInfo) SetHostname(v string) {
	o.Hostname = &v
}

// GetPromotable returns the Promotable field value if set, zero value otherwise.
func (o *MasterDeploymentInfo) GetPromotable() bool {
	if o == nil || o.Promotable == nil {
		var ret bool
		return ret
	}
	return *o.Promotable
}

// GetPromotableOk returns a tuple with the Promotable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterDeploymentInfo) GetPromotableOk() (*bool, bool) {
	if o == nil || o.Promotable == nil {
		return nil, false
	}
	return o.Promotable, true
}

// HasPromotable returns a boolean if a field has been set.
func (o *MasterDeploymentInfo) HasPromotable() bool {
	if o != nil && o.Promotable != nil {
		return true
	}

	return false
}

// SetPromotable gets a reference to the given bool and assigns it to the Promotable field.
func (o *MasterDeploymentInfo) SetPromotable(v bool) {
	o.Promotable = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *MasterDeploymentInfo) GetHealth() MasterHealth {
	if o == nil || o.Health == nil {
		var ret MasterHealth
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterDeploymentInfo) GetHealthOk() (*MasterHealth, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *MasterDeploymentInfo) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given MasterHealth and assigns it to the Health field.
func (o *MasterDeploymentInfo) SetHealth(v MasterHealth) {
	o.Health = &v
}

func (o MasterDeploymentInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NodeID != nil {
		toSerialize["nodeID"] = o.NodeID
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Promotable != nil {
		toSerialize["promotable"] = o.Promotable
	}
	if o.Health != nil {
		toSerialize["health"] = o.Health
	}
	return json.Marshal(toSerialize)
}

type NullableMasterDeploymentInfo struct {
	value *MasterDeploymentInfo
	isSet bool
}

func (v NullableMasterDeploymentInfo) Get() *MasterDeploymentInfo {
	return v.value
}

func (v *NullableMasterDeploymentInfo) Set(val *MasterDeploymentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMasterDeploymentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMasterDeploymentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMasterDeploymentInfo(val *MasterDeploymentInfo) *NullableMasterDeploymentInfo {
	return &NullableMasterDeploymentInfo{value: val, isSet: true}
}

func (v NullableMasterDeploymentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMasterDeploymentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


