# CreateVolumeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamespaceID** | **string** | A unique identifier for a namespace. The format of this type is undefined and may change but the defined properties will not change..  | 
**Labels** | Pointer to **map[string]string** | A set of arbitrary key value labels to apply to the entity.  | [optional] 
**Name** | **string** | The name of the volume shown in the CLI and UI  | 
**FsType** | [**FsType**](FsType.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**SizeBytes** | **uint64** | A volume&#39;s size in bytes  | 
**PlacementStrategy** | Pointer to [**Strategy**](Strategy.md) |  | [optional] 

## Methods

### NewCreateVolumeData

`func NewCreateVolumeData(namespaceID string, name string, fsType FsType, sizeBytes uint64, ) *CreateVolumeData`

NewCreateVolumeData instantiates a new CreateVolumeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVolumeDataWithDefaults

`func NewCreateVolumeDataWithDefaults() *CreateVolumeData`

NewCreateVolumeDataWithDefaults instantiates a new CreateVolumeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespaceID

`func (o *CreateVolumeData) GetNamespaceID() string`

GetNamespaceID returns the NamespaceID field if non-nil, zero value otherwise.

### GetNamespaceIDOk

`func (o *CreateVolumeData) GetNamespaceIDOk() (*string, bool)`

GetNamespaceIDOk returns a tuple with the NamespaceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceID

`func (o *CreateVolumeData) SetNamespaceID(v string)`

SetNamespaceID sets NamespaceID field to given value.


### GetLabels

`func (o *CreateVolumeData) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateVolumeData) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateVolumeData) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateVolumeData) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *CreateVolumeData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVolumeData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVolumeData) SetName(v string)`

SetName sets Name field to given value.


### GetFsType

`func (o *CreateVolumeData) GetFsType() FsType`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *CreateVolumeData) GetFsTypeOk() (*FsType, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *CreateVolumeData) SetFsType(v FsType)`

SetFsType sets FsType field to given value.


### GetDescription

`func (o *CreateVolumeData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVolumeData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVolumeData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVolumeData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSizeBytes

`func (o *CreateVolumeData) GetSizeBytes() uint64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *CreateVolumeData) GetSizeBytesOk() (*uint64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *CreateVolumeData) SetSizeBytes(v uint64)`

SetSizeBytes sets SizeBytes field to given value.


### GetPlacementStrategy

`func (o *CreateVolumeData) GetPlacementStrategy() Strategy`

GetPlacementStrategy returns the PlacementStrategy field if non-nil, zero value otherwise.

### GetPlacementStrategyOk

`func (o *CreateVolumeData) GetPlacementStrategyOk() (*Strategy, bool)`

GetPlacementStrategyOk returns a tuple with the PlacementStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementStrategy

`func (o *CreateVolumeData) SetPlacementStrategy(v Strategy)`

SetPlacementStrategy sets PlacementStrategy field to given value.

### HasPlacementStrategy

`func (o *CreateVolumeData) HasPlacementStrategy() bool`

HasPlacementStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


