# SyncProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesRemaining** | Pointer to **uint64** | Number of bytes left remaining to complete the sync.  | [optional] 
**ThroughputBytes** | Pointer to **uint64** | The average throughput of the sync given as bytes per  second.  | [optional] 
**EstimatedSecondsRemaining** | Pointer to **uint64** | The estimated time left for the sync to complete, given in seconds. When this field has a value of 0 either the  sync is complete or no duration estimate could be made. The values reported for bytesRemaining and  throughputBytes provide the client with the information needed to choose what to display.  | [optional] 

## Methods

### NewSyncProgress

`func NewSyncProgress() *SyncProgress`

NewSyncProgress instantiates a new SyncProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncProgressWithDefaults

`func NewSyncProgressWithDefaults() *SyncProgress`

NewSyncProgressWithDefaults instantiates a new SyncProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesRemaining

`func (o *SyncProgress) GetBytesRemaining() uint64`

GetBytesRemaining returns the BytesRemaining field if non-nil, zero value otherwise.

### GetBytesRemainingOk

`func (o *SyncProgress) GetBytesRemainingOk() (*uint64, bool)`

GetBytesRemainingOk returns a tuple with the BytesRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesRemaining

`func (o *SyncProgress) SetBytesRemaining(v uint64)`

SetBytesRemaining sets BytesRemaining field to given value.

### HasBytesRemaining

`func (o *SyncProgress) HasBytesRemaining() bool`

HasBytesRemaining returns a boolean if a field has been set.

### GetThroughputBytes

`func (o *SyncProgress) GetThroughputBytes() uint64`

GetThroughputBytes returns the ThroughputBytes field if non-nil, zero value otherwise.

### GetThroughputBytesOk

`func (o *SyncProgress) GetThroughputBytesOk() (*uint64, bool)`

GetThroughputBytesOk returns a tuple with the ThroughputBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputBytes

`func (o *SyncProgress) SetThroughputBytes(v uint64)`

SetThroughputBytes sets ThroughputBytes field to given value.

### HasThroughputBytes

`func (o *SyncProgress) HasThroughputBytes() bool`

HasThroughputBytes returns a boolean if a field has been set.

### GetEstimatedSecondsRemaining

`func (o *SyncProgress) GetEstimatedSecondsRemaining() uint64`

GetEstimatedSecondsRemaining returns the EstimatedSecondsRemaining field if non-nil, zero value otherwise.

### GetEstimatedSecondsRemainingOk

`func (o *SyncProgress) GetEstimatedSecondsRemainingOk() (*uint64, bool)`

GetEstimatedSecondsRemainingOk returns a tuple with the EstimatedSecondsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedSecondsRemaining

`func (o *SyncProgress) SetEstimatedSecondsRemaining(v uint64)`

SetEstimatedSecondsRemaining sets EstimatedSecondsRemaining field to given value.

### HasEstimatedSecondsRemaining

`func (o *SyncProgress) HasEstimatedSecondsRemaining() bool`

HasEstimatedSecondsRemaining returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


