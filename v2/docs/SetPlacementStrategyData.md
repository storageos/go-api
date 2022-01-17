# SetPlacementStrategyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlacementStrategy** | Pointer to [**Strategy**](Strategy.md) |  | [optional] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewSetPlacementStrategyData

`func NewSetPlacementStrategyData() *SetPlacementStrategyData`

NewSetPlacementStrategyData instantiates a new SetPlacementStrategyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPlacementStrategyDataWithDefaults

`func NewSetPlacementStrategyDataWithDefaults() *SetPlacementStrategyData`

NewSetPlacementStrategyDataWithDefaults instantiates a new SetPlacementStrategyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlacementStrategy

`func (o *SetPlacementStrategyData) GetPlacementStrategy() Strategy`

GetPlacementStrategy returns the PlacementStrategy field if non-nil, zero value otherwise.

### GetPlacementStrategyOk

`func (o *SetPlacementStrategyData) GetPlacementStrategyOk() (*Strategy, bool)`

GetPlacementStrategyOk returns a tuple with the PlacementStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementStrategy

`func (o *SetPlacementStrategyData) SetPlacementStrategy(v Strategy)`

SetPlacementStrategy sets PlacementStrategy field to given value.

### HasPlacementStrategy

`func (o *SetPlacementStrategyData) HasPlacementStrategy() bool`

HasPlacementStrategy returns a boolean if a field has been set.

### GetVersion

`func (o *SetPlacementStrategyData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SetPlacementStrategyData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SetPlacementStrategyData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SetPlacementStrategyData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


