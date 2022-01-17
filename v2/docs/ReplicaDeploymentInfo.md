# ReplicaDeploymentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for a volume deployment. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 
**NodeID** | Pointer to **string** | A unique identifier for a node. The format of this type is undefined and may change but the defined properties will not change.  | [optional] [readonly] 
**Hostname** | Pointer to **string** | The hostname of the node that is hosting the deployment  | [optional] 
**Promotable** | Pointer to **bool** | Indicates if the volume deployment is eligible for promotion  | [optional] 
**Health** | Pointer to [**ReplicaHealth**](ReplicaHealth.md) |  | [optional] 
**SyncProgress** | Pointer to [**SyncProgress**](SyncProgress.md) |  | [optional] 

## Methods

### NewReplicaDeploymentInfo

`func NewReplicaDeploymentInfo() *ReplicaDeploymentInfo`

NewReplicaDeploymentInfo instantiates a new ReplicaDeploymentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaDeploymentInfoWithDefaults

`func NewReplicaDeploymentInfoWithDefaults() *ReplicaDeploymentInfo`

NewReplicaDeploymentInfoWithDefaults instantiates a new ReplicaDeploymentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReplicaDeploymentInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReplicaDeploymentInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReplicaDeploymentInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReplicaDeploymentInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNodeID

`func (o *ReplicaDeploymentInfo) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *ReplicaDeploymentInfo) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *ReplicaDeploymentInfo) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *ReplicaDeploymentInfo) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetHostname

`func (o *ReplicaDeploymentInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ReplicaDeploymentInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ReplicaDeploymentInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ReplicaDeploymentInfo) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPromotable

`func (o *ReplicaDeploymentInfo) GetPromotable() bool`

GetPromotable returns the Promotable field if non-nil, zero value otherwise.

### GetPromotableOk

`func (o *ReplicaDeploymentInfo) GetPromotableOk() (*bool, bool)`

GetPromotableOk returns a tuple with the Promotable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotable

`func (o *ReplicaDeploymentInfo) SetPromotable(v bool)`

SetPromotable sets Promotable field to given value.

### HasPromotable

`func (o *ReplicaDeploymentInfo) HasPromotable() bool`

HasPromotable returns a boolean if a field has been set.

### GetHealth

`func (o *ReplicaDeploymentInfo) GetHealth() ReplicaHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ReplicaDeploymentInfo) GetHealthOk() (*ReplicaHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ReplicaDeploymentInfo) SetHealth(v ReplicaHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ReplicaDeploymentInfo) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetSyncProgress

`func (o *ReplicaDeploymentInfo) GetSyncProgress() SyncProgress`

GetSyncProgress returns the SyncProgress field if non-nil, zero value otherwise.

### GetSyncProgressOk

`func (o *ReplicaDeploymentInfo) GetSyncProgressOk() (*SyncProgress, bool)`

GetSyncProgressOk returns a tuple with the SyncProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProgress

`func (o *ReplicaDeploymentInfo) SetSyncProgress(v SyncProgress)`

SetSyncProgress sets SyncProgress field to given value.

### HasSyncProgress

`func (o *ReplicaDeploymentInfo) HasSyncProgress() bool`

HasSyncProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


