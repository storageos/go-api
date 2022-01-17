# UpdateUserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | If not present, the existing password is not changed | [optional] [default to "unchanged"]
**IsAdmin** | Pointer to **bool** | If true, this user is an administrator of the cluster. Administrators bypass the usual authentication checks and are granted access to all resources. Some actions (such as adding a new user) can only be performed by an administrator.  | [optional] [default to false]
**Groups** | Pointer to **[]string** | Defines a set of policy group IDs this user is a member of. Policy groups can be used to logically group users and apply authorisation  policies to all members.  | [optional] [default to []]
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewUpdateUserData

`func NewUpdateUserData() *UpdateUserData`

NewUpdateUserData instantiates a new UpdateUserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserDataWithDefaults

`func NewUpdateUserDataWithDefaults() *UpdateUserData`

NewUpdateUserDataWithDefaults instantiates a new UpdateUserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UpdateUserData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateUserData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateUserData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateUserData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetIsAdmin

`func (o *UpdateUserData) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *UpdateUserData) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *UpdateUserData) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *UpdateUserData) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetGroups

`func (o *UpdateUserData) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UpdateUserData) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UpdateUserData) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UpdateUserData) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *UpdateUserData) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *UpdateUserData) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetVersion

`func (o *UpdateUserData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateUserData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateUserData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateUserData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


