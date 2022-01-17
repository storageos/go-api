# NfsAclIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityType** | Pointer to **string** | The identity type used to identify the nfs client.  | [optional] 
**Matcher** | Pointer to **string** | NFS identity matcher. For \&quot;cidr\&quot;, this should be a valid CIDR block string such as \&quot;10.0.0.0/8\&quot;. For \&quot;hostname\&quot;, this must be the hostname sent by the client, with ? and * wildcard characters. For netgroup, this must be in the form of \&quot;@netgroup\&quot; with ? and * wildcard characters.  | [optional] 

## Methods

### NewNfsAclIdentity

`func NewNfsAclIdentity() *NfsAclIdentity`

NewNfsAclIdentity instantiates a new NfsAclIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsAclIdentityWithDefaults

`func NewNfsAclIdentityWithDefaults() *NfsAclIdentity`

NewNfsAclIdentityWithDefaults instantiates a new NfsAclIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityType

`func (o *NfsAclIdentity) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *NfsAclIdentity) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *NfsAclIdentity) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *NfsAclIdentity) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### GetMatcher

`func (o *NfsAclIdentity) GetMatcher() string`

GetMatcher returns the Matcher field if non-nil, zero value otherwise.

### GetMatcherOk

`func (o *NfsAclIdentity) GetMatcherOk() (*string, bool)`

GetMatcherOk returns a tuple with the Matcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcher

`func (o *NfsAclIdentity) SetMatcher(v string)`

SetMatcher sets Matcher field to given value.

### HasMatcher

`func (o *NfsAclIdentity) HasMatcher() bool`

HasMatcher returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


