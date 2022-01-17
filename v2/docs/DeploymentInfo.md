# DeploymentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for a volume deployment. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 
**NodeID** | Pointer to **string** | A unique identifier for a node. The format of this type is undefined and may change but the defined properties will not change.  | [optional] [readonly] 
**Hostname** | Pointer to **string** | The hostname of the node that is hosting the deployment  | [optional] 
**Promotable** | Pointer to **bool** | Indicates if the volume deployment is eligible for promotion  | [optional] 

## Methods

### NewDeploymentInfo

`func NewDeploymentInfo() *DeploymentInfo`

NewDeploymentInfo instantiates a new DeploymentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentInfoWithDefaults

`func NewDeploymentInfoWithDefaults() *DeploymentInfo`

NewDeploymentInfoWithDefaults instantiates a new DeploymentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNodeID

`func (o *DeploymentInfo) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *DeploymentInfo) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *DeploymentInfo) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *DeploymentInfo) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetHostname

`func (o *DeploymentInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DeploymentInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DeploymentInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DeploymentInfo) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPromotable

`func (o *DeploymentInfo) GetPromotable() bool`

GetPromotable returns the Promotable field if non-nil, zero value otherwise.

### GetPromotableOk

`func (o *DeploymentInfo) GetPromotableOk() (*bool, bool)`

GetPromotableOk returns a tuple with the Promotable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotable

`func (o *DeploymentInfo) SetPromotable(v bool)`

SetPromotable sets Promotable field to given value.

### HasPromotable

`func (o *DeploymentInfo) HasPromotable() bool`

HasPromotable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


