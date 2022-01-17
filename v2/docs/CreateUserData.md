# CreateUserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Password** | **string** | If not present, the existing password is not changed | [default to "unchanged"]
**IsAdmin** | Pointer to **bool** | If true, this user is an administrator of the cluster. Administrators bypass the usual authentication checks and are granted access to all resources. Some actions (such as adding a new user) can only be performed by an administrator.  | [optional] [default to false]
**Groups** | Pointer to **[]string** | Defines a set of policy group IDs this user is a member of. Policy groups can be used to logically group users and apply authorisation  policies to all members.  | [optional] [default to []]

## Methods

### NewCreateUserData

`func NewCreateUserData(username string, password string, ) *CreateUserData`

NewCreateUserData instantiates a new CreateUserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserDataWithDefaults

`func NewCreateUserDataWithDefaults() *CreateUserData`

NewCreateUserDataWithDefaults instantiates a new CreateUserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateUserData) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateUserData) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateUserData) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CreateUserData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUserData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUserData) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetIsAdmin

`func (o *CreateUserData) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *CreateUserData) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *CreateUserData) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *CreateUserData) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetGroups

`func (o *CreateUserData) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CreateUserData) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CreateUserData) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *CreateUserData) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *CreateUserData) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *CreateUserData) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


