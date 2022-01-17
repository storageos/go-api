# UpdatePolicyGroupData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Specs** | Pointer to [**[]PoliciesIdSpecs**](PoliciesIdSpecs.md) | A set of authorisation policies to apply to the policy group. | [optional] [default to []]
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewUpdatePolicyGroupData

`func NewUpdatePolicyGroupData() *UpdatePolicyGroupData`

NewUpdatePolicyGroupData instantiates a new UpdatePolicyGroupData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePolicyGroupDataWithDefaults

`func NewUpdatePolicyGroupDataWithDefaults() *UpdatePolicyGroupData`

NewUpdatePolicyGroupDataWithDefaults instantiates a new UpdatePolicyGroupData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpecs

`func (o *UpdatePolicyGroupData) GetSpecs() []PoliciesIdSpecs`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *UpdatePolicyGroupData) GetSpecsOk() (*[]PoliciesIdSpecs, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *UpdatePolicyGroupData) SetSpecs(v []PoliciesIdSpecs)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *UpdatePolicyGroupData) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### SetSpecsNil

`func (o *UpdatePolicyGroupData) SetSpecsNil(b bool)`

 SetSpecsNil sets the value for Specs to be an explicit nil

### UnsetSpecs
`func (o *UpdatePolicyGroupData) UnsetSpecs()`

UnsetSpecs ensures that no value is present for Specs, not even an explicit nil
### GetVersion

`func (o *UpdatePolicyGroupData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdatePolicyGroupData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdatePolicyGroupData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdatePolicyGroupData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


