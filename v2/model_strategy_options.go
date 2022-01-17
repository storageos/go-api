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

// StrategyOptions Used together with topology-aware strategy to further specify how the placement should be done. 
type StrategyOptions struct {
	// Indicates the node label used to decribe the topology used for data placement decisions. If two nodes are  labelled with this key and have identical values  for that label, the scheduler treats both nodes  as being in the same topology domain.  When topology-aware provisioning is enabled,  the scheduler tries to spread a volume's master  and replica copies across different topology domains. 
	TopologyKey *string `json:"topologyKey,omitempty"`
}

// NewStrategyOptions instantiates a new StrategyOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStrategyOptions() *StrategyOptions {
	this := StrategyOptions{}
	var topologyKey string = "topology.kubernetes.io/zone"
	this.TopologyKey = &topologyKey
	return &this
}

// NewStrategyOptionsWithDefaults instantiates a new StrategyOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStrategyOptionsWithDefaults() *StrategyOptions {
	this := StrategyOptions{}
	var topologyKey string = "topology.kubernetes.io/zone"
	this.TopologyKey = &topologyKey
	return &this
}

// GetTopologyKey returns the TopologyKey field value if set, zero value otherwise.
func (o *StrategyOptions) GetTopologyKey() string {
	if o == nil || o.TopologyKey == nil {
		var ret string
		return ret
	}
	return *o.TopologyKey
}

// GetTopologyKeyOk returns a tuple with the TopologyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StrategyOptions) GetTopologyKeyOk() (*string, bool) {
	if o == nil || o.TopologyKey == nil {
		return nil, false
	}
	return o.TopologyKey, true
}

// HasTopologyKey returns a boolean if a field has been set.
func (o *StrategyOptions) HasTopologyKey() bool {
	if o != nil && o.TopologyKey != nil {
		return true
	}

	return false
}

// SetTopologyKey gets a reference to the given string and assigns it to the TopologyKey field.
func (o *StrategyOptions) SetTopologyKey(v string) {
	o.TopologyKey = &v
}

func (o StrategyOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TopologyKey != nil {
		toSerialize["topologyKey"] = o.TopologyKey
	}
	return json.Marshal(toSerialize)
}

type NullableStrategyOptions struct {
	value *StrategyOptions
	isSet bool
}

func (v NullableStrategyOptions) Get() *StrategyOptions {
	return v.value
}

func (v *NullableStrategyOptions) Set(val *StrategyOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableStrategyOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableStrategyOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStrategyOptions(val *StrategyOptions) *NullableStrategyOptions {
	return &NullableStrategyOptions{value: val, isSet: true}
}

func (v NullableStrategyOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStrategyOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


