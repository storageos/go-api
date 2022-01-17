# SetCordonedNodeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cordoned** | Pointer to **bool** | Marks the node&#39;s desired cordoned state state.  A cordoned node will not have new volume deployments scheduled on it.  | [optional] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewSetCordonedNodeData

`func NewSetCordonedNodeData() *SetCordonedNodeData`

NewSetCordonedNodeData instantiates a new SetCordonedNodeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetCordonedNodeDataWithDefaults

`func NewSetCordonedNodeDataWithDefaults() *SetCordonedNodeData`

NewSetCordonedNodeDataWithDefaults instantiates a new SetCordonedNodeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCordoned

`func (o *SetCordonedNodeData) GetCordoned() bool`

GetCordoned returns the Cordoned field if non-nil, zero value otherwise.

### GetCordonedOk

`func (o *SetCordonedNodeData) GetCordonedOk() (*bool, bool)`

GetCordonedOk returns a tuple with the Cordoned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCordoned

`func (o *SetCordonedNodeData) SetCordoned(v bool)`

SetCordoned sets Cordoned field to given value.

### HasCordoned

`func (o *SetCordonedNodeData) HasCordoned() bool`

HasCordoned returns a boolean if a field has been set.

### GetVersion

`func (o *SetCordonedNodeData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SetCordonedNodeData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SetCordonedNodeData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SetCordonedNodeData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


