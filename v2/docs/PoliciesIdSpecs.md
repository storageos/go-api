# PoliciesIdSpecs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamespaceID** | Pointer to **string** | A unique identifier for a namespace. The format of this type is undefined and may change but the defined properties will not change..  | [optional] 
**ResourceType** | Pointer to **string** | The resource type this policy grants access to.  | [optional] 
**ReadOnly** | Pointer to **bool** | If true, disallows requests that attempt to mutate the resource.  | [optional] [default to false]

## Methods

### NewPoliciesIdSpecs

`func NewPoliciesIdSpecs() *PoliciesIdSpecs`

NewPoliciesIdSpecs instantiates a new PoliciesIdSpecs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoliciesIdSpecsWithDefaults

`func NewPoliciesIdSpecsWithDefaults() *PoliciesIdSpecs`

NewPoliciesIdSpecsWithDefaults instantiates a new PoliciesIdSpecs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespaceID

`func (o *PoliciesIdSpecs) GetNamespaceID() string`

GetNamespaceID returns the NamespaceID field if non-nil, zero value otherwise.

### GetNamespaceIDOk

`func (o *PoliciesIdSpecs) GetNamespaceIDOk() (*string, bool)`

GetNamespaceIDOk returns a tuple with the NamespaceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceID

`func (o *PoliciesIdSpecs) SetNamespaceID(v string)`

SetNamespaceID sets NamespaceID field to given value.

### HasNamespaceID

`func (o *PoliciesIdSpecs) HasNamespaceID() bool`

HasNamespaceID returns a boolean if a field has been set.

### GetResourceType

`func (o *PoliciesIdSpecs) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *PoliciesIdSpecs) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *PoliciesIdSpecs) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *PoliciesIdSpecs) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetReadOnly

`func (o *PoliciesIdSpecs) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *PoliciesIdSpecs) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *PoliciesIdSpecs) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *PoliciesIdSpecs) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


