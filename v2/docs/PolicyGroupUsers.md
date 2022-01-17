# PolicyGroupUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for a user. The format of this type is undefined and may change but the defined properties will not change..  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyGroupUsers

`func NewPolicyGroupUsers() *PolicyGroupUsers`

NewPolicyGroupUsers instantiates a new PolicyGroupUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyGroupUsersWithDefaults

`func NewPolicyGroupUsersWithDefaults() *PolicyGroupUsers`

NewPolicyGroupUsersWithDefaults instantiates a new PolicyGroupUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyGroupUsers) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyGroupUsers) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyGroupUsers) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyGroupUsers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *PolicyGroupUsers) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PolicyGroupUsers) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PolicyGroupUsers) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PolicyGroupUsers) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


