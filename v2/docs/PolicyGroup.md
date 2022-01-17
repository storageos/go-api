# PolicyGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for a policy group. The format of this type is undefined and may change but the defined properties will not change..  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**[]PolicyGroupUsers**](PolicyGroupUsers.md) | The list of user IDs which this policy group governs. | [optional] [readonly] 
**Specs** | Pointer to [**[]PoliciesIdSpecs**](PoliciesIdSpecs.md) | A set of authorisation policies to apply to the policy group. | [optional] [default to []]
**CreatedAt** | Pointer to **time.Time** | The time the entity was created. This timestamp is set by the node that created the entity, and may not be correct if the node&#39;s local clock was skewed. This value is for the user&#39;s informative purposes only, and correctness is not required. String format is RFC3339.  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the entity was last updated. This timestamp is set by the node that last updated the entity, and may not be correct if the node&#39;s local clock was skewed. This value is for the user&#39;s informative purposes only, and correctness is not required. String format is RFC3339.  | [optional] [readonly] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewPolicyGroup

`func NewPolicyGroup() *PolicyGroup`

NewPolicyGroup instantiates a new PolicyGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyGroupWithDefaults

`func NewPolicyGroupWithDefaults() *PolicyGroup`

NewPolicyGroupWithDefaults instantiates a new PolicyGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PolicyGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsers

`func (o *PolicyGroup) GetUsers() []PolicyGroupUsers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PolicyGroup) GetUsersOk() (*[]PolicyGroupUsers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PolicyGroup) SetUsers(v []PolicyGroupUsers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PolicyGroup) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetSpecs

`func (o *PolicyGroup) GetSpecs() []PoliciesIdSpecs`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *PolicyGroup) GetSpecsOk() (*[]PoliciesIdSpecs, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *PolicyGroup) SetSpecs(v []PoliciesIdSpecs)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *PolicyGroup) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### SetSpecsNil

`func (o *PolicyGroup) SetSpecsNil(b bool)`

 SetSpecsNil sets the value for Specs to be an explicit nil

### UnsetSpecs
`func (o *PolicyGroup) UnsetSpecs()`

UnsetSpecs ensures that no value is present for Specs, not even an explicit nil
### GetCreatedAt

`func (o *PolicyGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PolicyGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PolicyGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PolicyGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PolicyGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PolicyGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PolicyGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PolicyGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *PolicyGroup) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PolicyGroup) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PolicyGroup) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PolicyGroup) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


