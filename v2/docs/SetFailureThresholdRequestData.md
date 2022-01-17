# SetFailureThresholdRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureThreshold** | Pointer to **uint64** | The minimum number of replicas required to be online and receiving writes in order for the volume to remain read-writable. This value replaces any previously set failure threshold or intent-based failure mode.  | [optional] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewSetFailureThresholdRequestData

`func NewSetFailureThresholdRequestData() *SetFailureThresholdRequestData`

NewSetFailureThresholdRequestData instantiates a new SetFailureThresholdRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetFailureThresholdRequestDataWithDefaults

`func NewSetFailureThresholdRequestDataWithDefaults() *SetFailureThresholdRequestData`

NewSetFailureThresholdRequestDataWithDefaults instantiates a new SetFailureThresholdRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureThreshold

`func (o *SetFailureThresholdRequestData) GetFailureThreshold() uint64`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *SetFailureThresholdRequestData) GetFailureThresholdOk() (*uint64, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *SetFailureThresholdRequestData) SetFailureThreshold(v uint64)`

SetFailureThreshold sets FailureThreshold field to given value.

### HasFailureThreshold

`func (o *SetFailureThresholdRequestData) HasFailureThreshold() bool`

HasFailureThreshold returns a boolean if a field has been set.

### GetVersion

`func (o *SetFailureThresholdRequestData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SetFailureThresholdRequestData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SetFailureThresholdRequestData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SetFailureThresholdRequestData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


