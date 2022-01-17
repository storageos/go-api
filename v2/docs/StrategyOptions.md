# StrategyOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopologyKey** | Pointer to **string** | Indicates the node label used to decribe the topology used for data placement decisions. If two nodes are  labelled with this key and have identical values  for that label, the scheduler treats both nodes  as being in the same topology domain.  When topology-aware provisioning is enabled,  the scheduler tries to spread a volume&#39;s master  and replica copies across different topology domains.  | [optional] [default to "topology.kubernetes.io/zone"]

## Methods

### NewStrategyOptions

`func NewStrategyOptions() *StrategyOptions`

NewStrategyOptions instantiates a new StrategyOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategyOptionsWithDefaults

`func NewStrategyOptionsWithDefaults() *StrategyOptions`

NewStrategyOptionsWithDefaults instantiates a new StrategyOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopologyKey

`func (o *StrategyOptions) GetTopologyKey() string`

GetTopologyKey returns the TopologyKey field if non-nil, zero value otherwise.

### GetTopologyKeyOk

`func (o *StrategyOptions) GetTopologyKeyOk() (*string, bool)`

GetTopologyKeyOk returns a tuple with the TopologyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyKey

`func (o *StrategyOptions) SetTopologyKey(v string)`

SetTopologyKey sets TopologyKey field to given value.

### HasTopologyKey

`func (o *StrategyOptions) HasTopologyKey() bool`

HasTopologyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


