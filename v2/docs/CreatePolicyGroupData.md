# CreatePolicyGroupData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Specs** | Pointer to [**[]PoliciesSpecs**](PoliciesSpecs.md) | A set of authorisation policies to apply to the policy group. | [optional] [default to []]

## Methods

### NewCreatePolicyGroupData

`func NewCreatePolicyGroupData() *CreatePolicyGroupData`

NewCreatePolicyGroupData instantiates a new CreatePolicyGroupData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePolicyGroupDataWithDefaults

`func NewCreatePolicyGroupDataWithDefaults() *CreatePolicyGroupData`

NewCreatePolicyGroupDataWithDefaults instantiates a new CreatePolicyGroupData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePolicyGroupData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePolicyGroupData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePolicyGroupData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatePolicyGroupData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpecs

`func (o *CreatePolicyGroupData) GetSpecs() []PoliciesSpecs`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *CreatePolicyGroupData) GetSpecsOk() (*[]PoliciesSpecs, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *CreatePolicyGroupData) SetSpecs(v []PoliciesSpecs)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *CreatePolicyGroupData) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### SetSpecsNil

`func (o *CreatePolicyGroupData) SetSpecsNil(b bool)`

 SetSpecsNil sets the value for Specs to be an explicit nil

### UnsetSpecs
`func (o *CreatePolicyGroupData) UnsetSpecs()`

UnsetSpecs ensures that no value is present for Specs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


