# UserSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for a user. The format of this type is undefined and may change but the defined properties will not change..  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**IsAdmin** | Pointer to **bool** | If true, this user is an administrator of the cluster. Administrators bypass the usual authentication checks and are granted access to all resources. Some actions (such as adding a new user) can only be performed by an administrator.  | [optional] [default to false]
**Groups** | Pointer to **[]string** | Defines a set of policy group IDs this user is a member of. Policy groups can be used to logically group users and  apply authorisation policies to all members.  | [optional] [default to []]
**CreatedAt** | Pointer to **time.Time** | The time the entity was created. This timestamp is set by the node that created the entity, and may not be correct if the node&#39;s local clock was skewed. This value is for the user&#39;s informative purposes only, and correctness is not required. String format is RFC3339.  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the entity was last updated. This timestamp is set by the node that last updated the entity, and may not be correct if the node&#39;s local clock was skewed. This value is for the user&#39;s informative purposes only, and correctness is not required. String format is RFC3339.  | [optional] [readonly] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 
**Session** | Pointer to [**UserSessionAllOfSession**](UserSessionAllOfSession.md) |  | [optional] 

## Methods

### NewUserSession

`func NewUserSession() *UserSession`

NewUserSession instantiates a new UserSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionWithDefaults

`func NewUserSessionWithDefaults() *UserSession`

NewUserSessionWithDefaults instantiates a new UserSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *UserSession) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserSession) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserSession) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserSession) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetIsAdmin

`func (o *UserSession) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *UserSession) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *UserSession) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *UserSession) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetGroups

`func (o *UserSession) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserSession) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserSession) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserSession) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *UserSession) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *UserSession) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetCreatedAt

`func (o *UserSession) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserSession) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserSession) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserSession) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserSession) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserSession) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserSession) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserSession) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *UserSession) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UserSession) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UserSession) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UserSession) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSession

`func (o *UserSession) GetSession() UserSessionAllOfSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *UserSession) GetSessionOk() (*UserSessionAllOfSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *UserSession) SetSession(v UserSessionAllOfSession)`

SetSession sets Session field to given value.

### HasSession

`func (o *UserSession) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


