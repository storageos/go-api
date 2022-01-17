# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for a node. The format of this type is undefined and may change but the defined properties will not change.  | [optional] [readonly] 
**Name** | Pointer to **string** | The hostname of the node. This value is set by the node each time it joins the StorageOS cluster.  | [optional] [readonly] 
**Health** | Pointer to [**NodeHealth**](NodeHealth.md) |  | [optional] 
**Capacity** | Pointer to [**CapacityStats**](CapacityStats.md) |  | [optional] 
**IoEndpoint** | Pointer to **string** | Endpoint at which we operate our dataplane&#39;s dfs service. (used for IO operations) This value is set on startup by the corresponding environment variable (IO_ADVERTISE_ADDRESS)  | [optional] [readonly] 
**SupervisorEndpoint** | Pointer to **string** | Endpoint at which we operate our dataplane&#39;s supervisor service (used for sync). This value is set on startup by the corresponding environment variable (SUPERVISOR_ADVERTISE_ADDRESS)  | [optional] [readonly] 
**GossipEndpoint** | Pointer to **string** | Endpoint at which we operate our health checking service. This value is set on startup by the corresponding environment variable (GOSSIP_ADVERTISE_ADDRESS)  | [optional] [readonly] 
**ClusteringEndpoint** | Pointer to **string** | Endpoint at which we operate our clustering gRPC API. This value is set on startup by the corresponding environment variable (INTERNAL_API_ADVERTISE_ADDRESS)  | [optional] [readonly] 
**Cordoned** | Pointer to **bool** | Cordoned describes the cordoned state of the node. A cordoned node will not have new volume deployments scheduled on it  | [optional] [readonly] [default to false]
**CordonedAt** | Pointer to **time.Time** | The time the node has been cordoned. This field&#39;s purpose is informative only.  String format is RFC3339.  | [optional] [readonly] 
**Labels** | Pointer to **map[string]string** | A set of arbitrary key value labels to apply to the entity.  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the entity was created. This timestamp is set by the node that created the entity, and may not be correct if the node&#39;s local clock was skewed. This value is for the user&#39;s informative purposes only, and correctness is not required. String format is RFC3339.  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the entity was last updated. This timestamp is set by the node that last updated the entity, and may not be correct if the node&#39;s local clock was skewed. This value is for the user&#39;s informative purposes only, and correctness is not required. String format is RFC3339.  | [optional] [readonly] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewNode

`func NewNode() *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Node) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Node) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Node) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Node) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Node) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Node) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Node) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHealth

`func (o *Node) GetHealth() NodeHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *Node) GetHealthOk() (*NodeHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *Node) SetHealth(v NodeHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *Node) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetCapacity

`func (o *Node) GetCapacity() CapacityStats`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *Node) GetCapacityOk() (*CapacityStats, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *Node) SetCapacity(v CapacityStats)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *Node) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetIoEndpoint

`func (o *Node) GetIoEndpoint() string`

GetIoEndpoint returns the IoEndpoint field if non-nil, zero value otherwise.

### GetIoEndpointOk

`func (o *Node) GetIoEndpointOk() (*string, bool)`

GetIoEndpointOk returns a tuple with the IoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoEndpoint

`func (o *Node) SetIoEndpoint(v string)`

SetIoEndpoint sets IoEndpoint field to given value.

### HasIoEndpoint

`func (o *Node) HasIoEndpoint() bool`

HasIoEndpoint returns a boolean if a field has been set.

### GetSupervisorEndpoint

`func (o *Node) GetSupervisorEndpoint() string`

GetSupervisorEndpoint returns the SupervisorEndpoint field if non-nil, zero value otherwise.

### GetSupervisorEndpointOk

`func (o *Node) GetSupervisorEndpointOk() (*string, bool)`

GetSupervisorEndpointOk returns a tuple with the SupervisorEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorEndpoint

`func (o *Node) SetSupervisorEndpoint(v string)`

SetSupervisorEndpoint sets SupervisorEndpoint field to given value.

### HasSupervisorEndpoint

`func (o *Node) HasSupervisorEndpoint() bool`

HasSupervisorEndpoint returns a boolean if a field has been set.

### GetGossipEndpoint

`func (o *Node) GetGossipEndpoint() string`

GetGossipEndpoint returns the GossipEndpoint field if non-nil, zero value otherwise.

### GetGossipEndpointOk

`func (o *Node) GetGossipEndpointOk() (*string, bool)`

GetGossipEndpointOk returns a tuple with the GossipEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGossipEndpoint

`func (o *Node) SetGossipEndpoint(v string)`

SetGossipEndpoint sets GossipEndpoint field to given value.

### HasGossipEndpoint

`func (o *Node) HasGossipEndpoint() bool`

HasGossipEndpoint returns a boolean if a field has been set.

### GetClusteringEndpoint

`func (o *Node) GetClusteringEndpoint() string`

GetClusteringEndpoint returns the ClusteringEndpoint field if non-nil, zero value otherwise.

### GetClusteringEndpointOk

`func (o *Node) GetClusteringEndpointOk() (*string, bool)`

GetClusteringEndpointOk returns a tuple with the ClusteringEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusteringEndpoint

`func (o *Node) SetClusteringEndpoint(v string)`

SetClusteringEndpoint sets ClusteringEndpoint field to given value.

### HasClusteringEndpoint

`func (o *Node) HasClusteringEndpoint() bool`

HasClusteringEndpoint returns a boolean if a field has been set.

### GetCordoned

`func (o *Node) GetCordoned() bool`

GetCordoned returns the Cordoned field if non-nil, zero value otherwise.

### GetCordonedOk

`func (o *Node) GetCordonedOk() (*bool, bool)`

GetCordonedOk returns a tuple with the Cordoned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCordoned

`func (o *Node) SetCordoned(v bool)`

SetCordoned sets Cordoned field to given value.

### HasCordoned

`func (o *Node) HasCordoned() bool`

HasCordoned returns a boolean if a field has been set.

### GetCordonedAt

`func (o *Node) GetCordonedAt() time.Time`

GetCordonedAt returns the CordonedAt field if non-nil, zero value otherwise.

### GetCordonedAtOk

`func (o *Node) GetCordonedAtOk() (*time.Time, bool)`

GetCordonedAtOk returns a tuple with the CordonedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCordonedAt

`func (o *Node) SetCordonedAt(v time.Time)`

SetCordonedAt sets CordonedAt field to given value.

### HasCordonedAt

`func (o *Node) HasCordonedAt() bool`

HasCordonedAt returns a boolean if a field has been set.

### GetLabels

`func (o *Node) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Node) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Node) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Node) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Node) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Node) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Node) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Node) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Node) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Node) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Node) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Node) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *Node) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Node) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Node) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Node) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


