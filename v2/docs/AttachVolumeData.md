# AttachVolumeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeID** | Pointer to **string** | A unique identifier for a node. The format of this type is undefined and may change but the defined properties will not change.  | [optional] [readonly] 

## Methods

### NewAttachVolumeData

`func NewAttachVolumeData() *AttachVolumeData`

NewAttachVolumeData instantiates a new AttachVolumeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachVolumeDataWithDefaults

`func NewAttachVolumeDataWithDefaults() *AttachVolumeData`

NewAttachVolumeDataWithDefaults instantiates a new AttachVolumeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeID

`func (o *AttachVolumeData) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *AttachVolumeData) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *AttachVolumeData) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *AttachVolumeData) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


