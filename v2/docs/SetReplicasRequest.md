# SetReplicasRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Replicas** | Pointer to **uint64** | The number of replicas desired.  | [optional] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewSetReplicasRequest

`func NewSetReplicasRequest() *SetReplicasRequest`

NewSetReplicasRequest instantiates a new SetReplicasRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetReplicasRequestWithDefaults

`func NewSetReplicasRequestWithDefaults() *SetReplicasRequest`

NewSetReplicasRequestWithDefaults instantiates a new SetReplicasRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicas

`func (o *SetReplicasRequest) GetReplicas() uint64`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *SetReplicasRequest) GetReplicasOk() (*uint64, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *SetReplicasRequest) SetReplicas(v uint64)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *SetReplicasRequest) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetVersion

`func (o *SetReplicasRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SetReplicasRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SetReplicasRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SetReplicasRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


