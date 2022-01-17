# ReplicaDeploymentInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to [**ReplicaHealth**](ReplicaHealth.md) |  | [optional] 
**SyncProgress** | Pointer to [**SyncProgress**](SyncProgress.md) |  | [optional] 

## Methods

### NewReplicaDeploymentInfoAllOf

`func NewReplicaDeploymentInfoAllOf() *ReplicaDeploymentInfoAllOf`

NewReplicaDeploymentInfoAllOf instantiates a new ReplicaDeploymentInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaDeploymentInfoAllOfWithDefaults

`func NewReplicaDeploymentInfoAllOfWithDefaults() *ReplicaDeploymentInfoAllOf`

NewReplicaDeploymentInfoAllOfWithDefaults instantiates a new ReplicaDeploymentInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *ReplicaDeploymentInfoAllOf) GetHealth() ReplicaHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ReplicaDeploymentInfoAllOf) GetHealthOk() (*ReplicaHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ReplicaDeploymentInfoAllOf) SetHealth(v ReplicaHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ReplicaDeploymentInfoAllOf) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetSyncProgress

`func (o *ReplicaDeploymentInfoAllOf) GetSyncProgress() SyncProgress`

GetSyncProgress returns the SyncProgress field if non-nil, zero value otherwise.

### GetSyncProgressOk

`func (o *ReplicaDeploymentInfoAllOf) GetSyncProgressOk() (*SyncProgress, bool)`

GetSyncProgressOk returns a tuple with the SyncProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProgress

`func (o *ReplicaDeploymentInfoAllOf) SetSyncProgress(v SyncProgress)`

SetSyncProgress sets SyncProgress field to given value.

### HasSyncProgress

`func (o *ReplicaDeploymentInfoAllOf) HasSyncProgress() bool`

HasSyncProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


