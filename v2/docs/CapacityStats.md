# CapacityStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **uint64** | Total bytes in the filesystem  | [optional] 
**Free** | Pointer to **uint64** | Free bytes in the filesystem available to root user  | [optional] 
**Available** | Pointer to **uint64** | Byte value available to an unprivileged user  | [optional] 

## Methods

### NewCapacityStats

`func NewCapacityStats() *CapacityStats`

NewCapacityStats instantiates a new CapacityStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityStatsWithDefaults

`func NewCapacityStatsWithDefaults() *CapacityStats`

NewCapacityStatsWithDefaults instantiates a new CapacityStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CapacityStats) GetTotal() uint64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CapacityStats) GetTotalOk() (*uint64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CapacityStats) SetTotal(v uint64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CapacityStats) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetFree

`func (o *CapacityStats) GetFree() uint64`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *CapacityStats) GetFreeOk() (*uint64, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *CapacityStats) SetFree(v uint64)`

SetFree sets Free field to given value.

### HasFree

`func (o *CapacityStats) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetAvailable

`func (o *CapacityStats) GetAvailable() uint64`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CapacityStats) GetAvailableOk() (*uint64, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CapacityStats) SetAvailable(v uint64)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CapacityStats) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


