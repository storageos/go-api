# StrategyOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopologyKey** | **string** | Indicates the node label used to decribe the topology used for data placement decisions. If two nodes are  labelled with this key and have identical values  for that label, the scheduler treats both nodes  as being in the same topology domain.  When topology-aware provisioning is enabled,  the scheduler tries to spread a volume&#39;s master  and replica copies across different topology domains.  | [optional] [default to topology.kubernetes.io/zone]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


