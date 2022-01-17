# NfsAcl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to [**NfsAclIdentity**](NfsAclIdentity.md) |  | [optional] 
**SquashConfig** | Pointer to [**NfsAclSquashConfig**](NfsAclSquashConfig.md) |  | [optional] 
**AccessLevel** | Pointer to **string** | The access level this ACL grants - read-only, or read-write.  | [optional] 

## Methods

### NewNfsAcl

`func NewNfsAcl() *NfsAcl`

NewNfsAcl instantiates a new NfsAcl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsAclWithDefaults

`func NewNfsAclWithDefaults() *NfsAcl`

NewNfsAclWithDefaults instantiates a new NfsAcl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *NfsAcl) GetIdentity() NfsAclIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *NfsAcl) GetIdentityOk() (*NfsAclIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *NfsAcl) SetIdentity(v NfsAclIdentity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *NfsAcl) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetSquashConfig

`func (o *NfsAcl) GetSquashConfig() NfsAclSquashConfig`

GetSquashConfig returns the SquashConfig field if non-nil, zero value otherwise.

### GetSquashConfigOk

`func (o *NfsAcl) GetSquashConfigOk() (*NfsAclSquashConfig, bool)`

GetSquashConfigOk returns a tuple with the SquashConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquashConfig

`func (o *NfsAcl) SetSquashConfig(v NfsAclSquashConfig)`

SetSquashConfig sets SquashConfig field to given value.

### HasSquashConfig

`func (o *NfsAcl) HasSquashConfig() bool`

HasSquashConfig returns a boolean if a field has been set.

### GetAccessLevel

`func (o *NfsAcl) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *NfsAcl) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *NfsAcl) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *NfsAcl) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


