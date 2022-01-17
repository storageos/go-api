# SetFailureModeIntentRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to [**FailureModeIntent**](FailureModeIntent.md) |  | [optional] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewSetFailureModeIntentRequestData

`func NewSetFailureModeIntentRequestData() *SetFailureModeIntentRequestData`

NewSetFailureModeIntentRequestData instantiates a new SetFailureModeIntentRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetFailureModeIntentRequestDataWithDefaults

`func NewSetFailureModeIntentRequestDataWithDefaults() *SetFailureModeIntentRequestData`

NewSetFailureModeIntentRequestDataWithDefaults instantiates a new SetFailureModeIntentRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *SetFailureModeIntentRequestData) GetMode() FailureModeIntent`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SetFailureModeIntentRequestData) GetModeOk() (*FailureModeIntent, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SetFailureModeIntentRequestData) SetMode(v FailureModeIntent)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SetFailureModeIntentRequestData) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetVersion

`func (o *SetFailureModeIntentRequestData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SetFailureModeIntentRequestData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SetFailureModeIntentRequestData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SetFailureModeIntentRequestData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


