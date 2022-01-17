# AttachNFSVolumeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewAttachNFSVolumeData

`func NewAttachNFSVolumeData() *AttachNFSVolumeData`

NewAttachNFSVolumeData instantiates a new AttachNFSVolumeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachNFSVolumeDataWithDefaults

`func NewAttachNFSVolumeDataWithDefaults() *AttachNFSVolumeData`

NewAttachNFSVolumeDataWithDefaults instantiates a new AttachNFSVolumeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *AttachNFSVolumeData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AttachNFSVolumeData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AttachNFSVolumeData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AttachNFSVolumeData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


