# NfsAclSquashConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **int64** |  | [optional] 
**Gid** | Pointer to **int64** |  | [optional] 
**Squash** | Pointer to **string** | SquashConfig defines the root squashing behaviour.  When a client creates a file, it sends the user UID from the client. If the client is running as root, this sends uid&#x3D;0. Root squashing allows the NFS administrator to prevent the client from writing as \&quot;root\&quot; to the NFS share, instead mapping the client to a new UID/GID (usually nfsnobody, -2). \&quot;none\&quot; performs no UID/GID alterations, using the values sent by the client. \&quot;root\&quot; maps UID &amp; GID 0 to the values specified. \&quot;rootuid\&quot; maps UID 0 and a GID of any value to the value specified. \&quot;all\&quot; maps changes all UID and GID values to those specified.  | [optional] 

## Methods

### NewNfsAclSquashConfig

`func NewNfsAclSquashConfig() *NfsAclSquashConfig`

NewNfsAclSquashConfig instantiates a new NfsAclSquashConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsAclSquashConfigWithDefaults

`func NewNfsAclSquashConfigWithDefaults() *NfsAclSquashConfig`

NewNfsAclSquashConfigWithDefaults instantiates a new NfsAclSquashConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *NfsAclSquashConfig) GetUid() int64`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *NfsAclSquashConfig) GetUidOk() (*int64, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *NfsAclSquashConfig) SetUid(v int64)`

SetUid sets Uid field to given value.

### HasUid

`func (o *NfsAclSquashConfig) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetGid

`func (o *NfsAclSquashConfig) GetGid() int64`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *NfsAclSquashConfig) GetGidOk() (*int64, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *NfsAclSquashConfig) SetGid(v int64)`

SetGid sets Gid field to given value.

### HasGid

`func (o *NfsAclSquashConfig) HasGid() bool`

HasGid returns a boolean if a field has been set.

### GetSquash

`func (o *NfsAclSquashConfig) GetSquash() string`

GetSquash returns the Squash field if non-nil, zero value otherwise.

### GetSquashOk

`func (o *NfsAclSquashConfig) GetSquashOk() (*string, bool)`

GetSquashOk returns a tuple with the Squash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquash

`func (o *NfsAclSquashConfig) SetSquash(v string)`

SetSquash sets Squash field to given value.

### HasSquash

`func (o *NfsAclSquashConfig) HasSquash() bool`

HasSquash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


