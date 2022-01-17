# ResizeVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeBytes** | Pointer to **uint64** | The desired new size for the volume in  bytes. This value cannot be less than  the current size of the volume.  | [optional] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewResizeVolumeRequest

`func NewResizeVolumeRequest() *ResizeVolumeRequest`

NewResizeVolumeRequest instantiates a new ResizeVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResizeVolumeRequestWithDefaults

`func NewResizeVolumeRequestWithDefaults() *ResizeVolumeRequest`

NewResizeVolumeRequestWithDefaults instantiates a new ResizeVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeBytes

`func (o *ResizeVolumeRequest) GetSizeBytes() uint64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *ResizeVolumeRequest) GetSizeBytesOk() (*uint64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *ResizeVolumeRequest) SetSizeBytes(v uint64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *ResizeVolumeRequest) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetVersion

`func (o *ResizeVolumeRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResizeVolumeRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResizeVolumeRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResizeVolumeRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


