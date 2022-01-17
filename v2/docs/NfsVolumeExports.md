# NFSVolumeExports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exports** | Pointer to [**[]NfsExportConfig**](NfsExportConfig.md) |  | [optional] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewNFSVolumeExports

`func NewNFSVolumeExports() *NFSVolumeExports`

NewNFSVolumeExports instantiates a new NFSVolumeExports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFSVolumeExportsWithDefaults

`func NewNFSVolumeExportsWithDefaults() *NFSVolumeExports`

NewNFSVolumeExportsWithDefaults instantiates a new NFSVolumeExports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExports

`func (o *NFSVolumeExports) GetExports() []NfsExportConfig`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *NFSVolumeExports) GetExportsOk() (*[]NfsExportConfig, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *NFSVolumeExports) SetExports(v []NfsExportConfig)`

SetExports sets Exports field to given value.

### HasExports

`func (o *NFSVolumeExports) HasExports() bool`

HasExports returns a boolean if a field has been set.

### GetVersion

`func (o *NFSVolumeExports) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NFSVolumeExports) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NFSVolumeExports) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NFSVolumeExports) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


