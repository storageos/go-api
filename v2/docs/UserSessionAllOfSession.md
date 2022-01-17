# UserSessionAllOfSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresInSeconds** | Pointer to **uint64** | The maximum duration which the auth session  will remain valid for in seconds.  | [optional] 
**Token** | Pointer to **string** | The JWT token for the auth session.  | [optional] 

## Methods

### NewUserSessionAllOfSession

`func NewUserSessionAllOfSession() *UserSessionAllOfSession`

NewUserSessionAllOfSession instantiates a new UserSessionAllOfSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionAllOfSessionWithDefaults

`func NewUserSessionAllOfSessionWithDefaults() *UserSessionAllOfSession`

NewUserSessionAllOfSessionWithDefaults instantiates a new UserSessionAllOfSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresInSeconds

`func (o *UserSessionAllOfSession) GetExpiresInSeconds() uint64`

GetExpiresInSeconds returns the ExpiresInSeconds field if non-nil, zero value otherwise.

### GetExpiresInSecondsOk

`func (o *UserSessionAllOfSession) GetExpiresInSecondsOk() (*uint64, bool)`

GetExpiresInSecondsOk returns a tuple with the ExpiresInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInSeconds

`func (o *UserSessionAllOfSession) SetExpiresInSeconds(v uint64)`

SetExpiresInSeconds sets ExpiresInSeconds field to given value.

### HasExpiresInSeconds

`func (o *UserSessionAllOfSession) HasExpiresInSeconds() bool`

HasExpiresInSeconds returns a boolean if a field has been set.

### GetToken

`func (o *UserSessionAllOfSession) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserSessionAllOfSession) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserSessionAllOfSession) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UserSessionAllOfSession) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


