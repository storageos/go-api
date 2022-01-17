# NfsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exports** | Pointer to [**[]NfsExportConfig**](NfsExportConfig.md) |  | [optional] 
**ServiceEndpoint** | Pointer to **NullableString** | The address to which the NFS server is bound.  | [optional] [readonly] [default to ""]

## Methods

### NewNfsConfig

`func NewNfsConfig() *NfsConfig`

NewNfsConfig instantiates a new NfsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsConfigWithDefaults

`func NewNfsConfigWithDefaults() *NfsConfig`

NewNfsConfigWithDefaults instantiates a new NfsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExports

`func (o *NfsConfig) GetExports() []NfsExportConfig`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *NfsConfig) GetExportsOk() (*[]NfsExportConfig, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *NfsConfig) SetExports(v []NfsExportConfig)`

SetExports sets Exports field to given value.

### HasExports

`func (o *NfsConfig) HasExports() bool`

HasExports returns a boolean if a field has been set.

### SetExportsNil

`func (o *NfsConfig) SetExportsNil(b bool)`

 SetExportsNil sets the value for Exports to be an explicit nil

### UnsetExports
`func (o *NfsConfig) UnsetExports()`

UnsetExports ensures that no value is present for Exports, not even an explicit nil
### GetServiceEndpoint

`func (o *NfsConfig) GetServiceEndpoint() string`

GetServiceEndpoint returns the ServiceEndpoint field if non-nil, zero value otherwise.

### GetServiceEndpointOk

`func (o *NfsConfig) GetServiceEndpointOk() (*string, bool)`

GetServiceEndpointOk returns a tuple with the ServiceEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndpoint

`func (o *NfsConfig) SetServiceEndpoint(v string)`

SetServiceEndpoint sets ServiceEndpoint field to given value.

### HasServiceEndpoint

`func (o *NfsConfig) HasServiceEndpoint() bool`

HasServiceEndpoint returns a boolean if a field has been set.

### SetServiceEndpointNil

`func (o *NfsConfig) SetServiceEndpointNil(b bool)`

 SetServiceEndpointNil sets the value for ServiceEndpoint to be an explicit nil

### UnsetServiceEndpoint
`func (o *NfsConfig) UnsetServiceEndpoint()`

UnsetServiceEndpoint ensures that no value is present for ServiceEndpoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


