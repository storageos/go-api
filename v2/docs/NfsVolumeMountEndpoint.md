# NFSVolumeMountEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountEndpoint** | Pointer to **string** | The address to which the NFS server is bound.  | [optional] [default to ""]
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewNFSVolumeMountEndpoint

`func NewNFSVolumeMountEndpoint() *NFSVolumeMountEndpoint`

NewNFSVolumeMountEndpoint instantiates a new NFSVolumeMountEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFSVolumeMountEndpointWithDefaults

`func NewNFSVolumeMountEndpointWithDefaults() *NFSVolumeMountEndpoint`

NewNFSVolumeMountEndpointWithDefaults instantiates a new NFSVolumeMountEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountEndpoint

`func (o *NFSVolumeMountEndpoint) GetMountEndpoint() string`

GetMountEndpoint returns the MountEndpoint field if non-nil, zero value otherwise.

### GetMountEndpointOk

`func (o *NFSVolumeMountEndpoint) GetMountEndpointOk() (*string, bool)`

GetMountEndpointOk returns a tuple with the MountEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountEndpoint

`func (o *NFSVolumeMountEndpoint) SetMountEndpoint(v string)`

SetMountEndpoint sets MountEndpoint field to given value.

### HasMountEndpoint

`func (o *NFSVolumeMountEndpoint) HasMountEndpoint() bool`

HasMountEndpoint returns a boolean if a field has been set.

### GetVersion

`func (o *NFSVolumeMountEndpoint) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NFSVolumeMountEndpoint) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NFSVolumeMountEndpoint) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NFSVolumeMountEndpoint) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


