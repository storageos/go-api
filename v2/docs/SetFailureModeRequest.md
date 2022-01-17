# SetFailureModeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureThreshold** | Pointer to **uint64** | The minimum number of replicas required to be online and receiving writes in order for the volume to remain read-writable. This value replaces any previously set failure threshold or intent-based failure mode.  | [optional] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 
**Mode** | Pointer to [**FailureModeIntent**](FailureModeIntent.md) |  | [optional] 

## Methods

### NewSetFailureModeRequest

`func NewSetFailureModeRequest() *SetFailureModeRequest`

NewSetFailureModeRequest instantiates a new SetFailureModeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetFailureModeRequestWithDefaults

`func NewSetFailureModeRequestWithDefaults() *SetFailureModeRequest`

NewSetFailureModeRequestWithDefaults instantiates a new SetFailureModeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureThreshold

`func (o *SetFailureModeRequest) GetFailureThreshold() uint64`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *SetFailureModeRequest) GetFailureThresholdOk() (*uint64, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *SetFailureModeRequest) SetFailureThreshold(v uint64)`

SetFailureThreshold sets FailureThreshold field to given value.

### HasFailureThreshold

`func (o *SetFailureModeRequest) HasFailureThreshold() bool`

HasFailureThreshold returns a boolean if a field has been set.

### GetVersion

`func (o *SetFailureModeRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SetFailureModeRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SetFailureModeRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SetFailureModeRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMode

`func (o *SetFailureModeRequest) GetMode() FailureModeIntent`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SetFailureModeRequest) GetModeOk() (*FailureModeIntent, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SetFailureModeRequest) SetMode(v FailureModeIntent)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SetFailureModeRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


