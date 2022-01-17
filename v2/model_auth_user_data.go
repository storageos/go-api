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

// AuthUserData struct for AuthUserData
type AuthUserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewAuthUserData instantiates a new AuthUserData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthUserData(username string, password string) *AuthUserData {
	this := AuthUserData{}
	this.Username = username
	this.Password = password
	return &this
}

// NewAuthUserDataWithDefaults instantiates a new AuthUserData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthUserDataWithDefaults() *AuthUserData {
	this := AuthUserData{}
	return &this
}

// GetUsername returns the Username field value
func (o *AuthUserData) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AuthUserData) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AuthUserData) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *AuthUserData) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AuthUserData) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AuthUserData) SetPassword(v string) {
	o.Password = v
}

func (o AuthUserData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableAuthUserData struct {
	value *AuthUserData
	isSet bool
}

func (v NullableAuthUserData) Get() *AuthUserData {
	return v.value
}

func (v *NullableAuthUserData) Set(val *AuthUserData) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthUserData) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthUserData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthUserData(val *AuthUserData) *NullableAuthUserData {
	return &NullableAuthUserData{value: val, isSet: true}
}

func (v NullableAuthUserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthUserData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


