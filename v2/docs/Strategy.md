# Strategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strategy** | Pointer to **string** |  | [optional] [default to "recommended"]
**Options** | Pointer to [**NullableStrategyOptions**](StrategyOptions.md) |  | [optional] 

## Methods

### NewStrategy

`func NewStrategy() *Strategy`

NewStrategy instantiates a new Strategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategyWithDefaults

`func NewStrategyWithDefaults() *Strategy`

NewStrategyWithDefaults instantiates a new Strategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrategy

`func (o *Strategy) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *Strategy) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *Strategy) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *Strategy) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetOptions

`func (o *Strategy) GetOptions() StrategyOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Strategy) GetOptionsOk() (*StrategyOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Strategy) SetOptions(v StrategyOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Strategy) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *Strategy) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *Strategy) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


