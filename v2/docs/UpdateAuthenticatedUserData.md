# UpdateAuthenticatedUserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | If not present, the existing password is not changed | [optional] [default to "unchanged"]
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewUpdateAuthenticatedUserData

`func NewUpdateAuthenticatedUserData() *UpdateAuthenticatedUserData`

NewUpdateAuthenticatedUserData instantiates a new UpdateAuthenticatedUserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAuthenticatedUserDataWithDefaults

`func NewUpdateAuthenticatedUserDataWithDefaults() *UpdateAuthenticatedUserData`

NewUpdateAuthenticatedUserDataWithDefaults instantiates a new UpdateAuthenticatedUserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UpdateAuthenticatedUserData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateAuthenticatedUserData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateAuthenticatedUserData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateAuthenticatedUserData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateAuthenticatedUserData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateAuthenticatedUserData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateAuthenticatedUserData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateAuthenticatedUserData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


