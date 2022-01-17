# Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for a volume. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AttachedOn** | Pointer to **string** |  | [optional] [readonly] 
**AttachedOnHost** | Pointer to **string** | The hostname of the node the volume is attached on | [optional] [readonly] 
**Nfs** | Pointer to [**NfsConfig**](NfsConfig.md) |  | [optional] 
**NamespaceID** | Pointer to **string** |  | [optional] [readonly] 
**Labels** | Pointer to **map[string]string** | A set of arbitrary key value labels to apply to the entity.  | [optional] 
**FsType** | Pointer to [**FsType**](FsType.md) |  | [optional] 
**AttachmentType** | Pointer to [**AttachType**](AttachType.md) |  | [optional] 
**Master** | Pointer to [**MasterDeploymentInfo**](MasterDeploymentInfo.md) |  | [optional] [readonly] 
**Replicas** | Pointer to [**[]ReplicaDeploymentInfo**](ReplicaDeploymentInfo.md) |  | [optional] [readonly] [default to []]
**SizeBytes** | Pointer to **uint64** | A volume&#39;s size in bytes  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the entity was created. This timestamp is set by the node that created the entity, and may not be correct if the node&#39;s local clock was skewed. This value is for the user&#39;s informative purposes only, and correctness is not required. String format is RFC3339.  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the entity was last updated. This timestamp is set by the node that last updated the entity, and may not be correct if the node&#39;s local clock was skewed. This value is for the user&#39;s informative purposes only, and correctness is not required. String format is RFC3339.  | [optional] [readonly] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewVolume

`func NewVolume() *Volume`

NewVolume instantiates a new Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeWithDefaults

`func NewVolumeWithDefaults() *Volume`

NewVolumeWithDefaults instantiates a new Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Volume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Volume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Volume) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Volume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Volume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Volume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Volume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Volume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Volume) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Volume) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Volume) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Volume) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttachedOn

`func (o *Volume) GetAttachedOn() string`

GetAttachedOn returns the AttachedOn field if non-nil, zero value otherwise.

### GetAttachedOnOk

`func (o *Volume) GetAttachedOnOk() (*string, bool)`

GetAttachedOnOk returns a tuple with the AttachedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedOn

`func (o *Volume) SetAttachedOn(v string)`

SetAttachedOn sets AttachedOn field to given value.

### HasAttachedOn

`func (o *Volume) HasAttachedOn() bool`

HasAttachedOn returns a boolean if a field has been set.

### SetAttachedOnNil

`func (o *Volume) SetAttachedOnNil(b bool)`

 SetAttachedOnNil sets the value for AttachedOn to be an explicit nil

### UnsetAttachedOn
`func (o *Volume) UnsetAttachedOn()`

UnsetAttachedOn ensures that no value is present for AttachedOn, not even an explicit nil
### GetAttachedOnHost

`func (o *Volume) GetAttachedOnHost() string`

GetAttachedOnHost returns the AttachedOnHost field if non-nil, zero value otherwise.

### GetAttachedOnHostOk

`func (o *Volume) GetAttachedOnHostOk() (*string, bool)`

GetAttachedOnHostOk returns a tuple with the AttachedOnHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedOnHost

`func (o *Volume) SetAttachedOnHost(v string)`

SetAttachedOnHost sets AttachedOnHost field to given value.

### HasAttachedOnHost

`func (o *Volume) HasAttachedOnHost() bool`

HasAttachedOnHost returns a boolean if a field has been set.

### GetNfs

`func (o *Volume) GetNfs() NfsConfig`

GetNfs returns the Nfs field if non-nil, zero value otherwise.

### GetNfsOk

`func (o *Volume) GetNfsOk() (*NfsConfig, bool)`

GetNfsOk returns a tuple with the Nfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfs

`func (o *Volume) SetNfs(v NfsConfig)`

SetNfs sets Nfs field to given value.

### HasNfs

`func (o *Volume) HasNfs() bool`

HasNfs returns a boolean if a field has been set.

### GetNamespaceID

`func (o *Volume) GetNamespaceID() string`

GetNamespaceID returns the NamespaceID field if non-nil, zero value otherwise.

### GetNamespaceIDOk

`func (o *Volume) GetNamespaceIDOk() (*string, bool)`

GetNamespaceIDOk returns a tuple with the NamespaceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceID

`func (o *Volume) SetNamespaceID(v string)`

SetNamespaceID sets NamespaceID field to given value.

### HasNamespaceID

`func (o *Volume) HasNamespaceID() bool`

HasNamespaceID returns a boolean if a field has been set.

### SetNamespaceIDNil

`func (o *Volume) SetNamespaceIDNil(b bool)`

 SetNamespaceIDNil sets the value for NamespaceID to be an explicit nil

### UnsetNamespaceID
`func (o *Volume) UnsetNamespaceID()`

UnsetNamespaceID ensures that no value is present for NamespaceID, not even an explicit nil
### GetLabels

`func (o *Volume) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Volume) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Volume) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Volume) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetFsType

`func (o *Volume) GetFsType() FsType`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *Volume) GetFsTypeOk() (*FsType, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *Volume) SetFsType(v FsType)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *Volume) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetAttachmentType

`func (o *Volume) GetAttachmentType() AttachType`

GetAttachmentType returns the AttachmentType field if non-nil, zero value otherwise.

### GetAttachmentTypeOk

`func (o *Volume) GetAttachmentTypeOk() (*AttachType, bool)`

GetAttachmentTypeOk returns a tuple with the AttachmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentType

`func (o *Volume) SetAttachmentType(v AttachType)`

SetAttachmentType sets AttachmentType field to given value.

### HasAttachmentType

`func (o *Volume) HasAttachmentType() bool`

HasAttachmentType returns a boolean if a field has been set.

### GetMaster

`func (o *Volume) GetMaster() MasterDeploymentInfo`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *Volume) GetMasterOk() (*MasterDeploymentInfo, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *Volume) SetMaster(v MasterDeploymentInfo)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *Volume) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### SetMasterNil

`func (o *Volume) SetMasterNil(b bool)`

 SetMasterNil sets the value for Master to be an explicit nil

### UnsetMaster
`func (o *Volume) UnsetMaster()`

UnsetMaster ensures that no value is present for Master, not even an explicit nil
### GetReplicas

`func (o *Volume) GetReplicas() []ReplicaDeploymentInfo`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *Volume) GetReplicasOk() (*[]ReplicaDeploymentInfo, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *Volume) SetReplicas(v []ReplicaDeploymentInfo)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *Volume) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### SetReplicasNil

`func (o *Volume) SetReplicasNil(b bool)`

 SetReplicasNil sets the value for Replicas to be an explicit nil

### UnsetReplicas
`func (o *Volume) UnsetReplicas()`

UnsetReplicas ensures that no value is present for Replicas, not even an explicit nil
### GetSizeBytes

`func (o *Volume) GetSizeBytes() uint64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *Volume) GetSizeBytesOk() (*uint64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *Volume) SetSizeBytes(v uint64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *Volume) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Volume) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Volume) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Volume) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Volume) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Volume) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Volume) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Volume) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Volume) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *Volume) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Volume) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Volume) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Volume) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


