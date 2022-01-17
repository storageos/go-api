# SetComputeOnlyNodeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeOnly** | Pointer to **bool** | Marks the node&#39;s desired configuration  state as compute-only. This will result in the node being avoided for volume placement  | [optional] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewSetComputeOnlyNodeData

`func NewSetComputeOnlyNodeData() *SetComputeOnlyNodeData`

NewSetComputeOnlyNodeData instantiates a new SetComputeOnlyNodeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetComputeOnlyNodeDataWithDefaults

`func NewSetComputeOnlyNodeDataWithDefaults() *SetComputeOnlyNodeData`

NewSetComputeOnlyNodeDataWithDefaults instantiates a new SetComputeOnlyNodeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeOnly

`func (o *SetComputeOnlyNodeData) GetComputeOnly() bool`

GetComputeOnly returns the ComputeOnly field if non-nil, zero value otherwise.

### GetComputeOnlyOk

`func (o *SetComputeOnlyNodeData) GetComputeOnlyOk() (*bool, bool)`

GetComputeOnlyOk returns a tuple with the ComputeOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOnly

`func (o *SetComputeOnlyNodeData) SetComputeOnly(v bool)`

SetComputeOnly sets ComputeOnly field to given value.

### HasComputeOnly

`func (o *SetComputeOnlyNodeData) HasComputeOnly() bool`

HasComputeOnly returns a boolean if a field has been set.

### GetVersion

`func (o *SetComputeOnlyNodeData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SetComputeOnlyNodeData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SetComputeOnlyNodeData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SetComputeOnlyNodeData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


