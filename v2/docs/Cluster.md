# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for a cluster. The format of this type is undefined and may change but the defined properties will not change.  | [optional] [readonly] 
**DisableTelemetry** | Pointer to **bool** | Disables collection of telemetry data across the cluster. | [optional] [default to false]
**DisableCrashReporting** | Pointer to **bool** | Disables collection of reports for any fatal crashes across the cluster.  | [optional] [default to false]
**DisableVersionCheck** | Pointer to **bool** | Disables the mechanism responsible for checking if there is an updated version of StorageOS available for installation.  | [optional] [default to false]
**LogLevel** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] [default to LOGLEVEL_INFO]
**LogFormat** | Pointer to [**LogFormat**](LogFormat.md) |  | [optional] [default to LOGFORMAT_DEFAULT]
**CreatedAt** | Pointer to **time.Time** | The time the entity was created. This timestamp is set by the node that created the entity, and may not be correct if the node&#39;s local clock was skewed. This value is for the user&#39;s informative purposes only, and correctness is not required. String format is RFC3339.  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the entity was last updated. This timestamp is set by the node that last updated the entity, and may not be correct if the node&#39;s local clock was skewed. This value is for the user&#39;s informative purposes only, and correctness is not required. String format is RFC3339.  | [optional] [readonly] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewCluster

`func NewCluster() *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Cluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisableTelemetry

`func (o *Cluster) GetDisableTelemetry() bool`

GetDisableTelemetry returns the DisableTelemetry field if non-nil, zero value otherwise.

### GetDisableTelemetryOk

`func (o *Cluster) GetDisableTelemetryOk() (*bool, bool)`

GetDisableTelemetryOk returns a tuple with the DisableTelemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTelemetry

`func (o *Cluster) SetDisableTelemetry(v bool)`

SetDisableTelemetry sets DisableTelemetry field to given value.

### HasDisableTelemetry

`func (o *Cluster) HasDisableTelemetry() bool`

HasDisableTelemetry returns a boolean if a field has been set.

### GetDisableCrashReporting

`func (o *Cluster) GetDisableCrashReporting() bool`

GetDisableCrashReporting returns the DisableCrashReporting field if non-nil, zero value otherwise.

### GetDisableCrashReportingOk

`func (o *Cluster) GetDisableCrashReportingOk() (*bool, bool)`

GetDisableCrashReportingOk returns a tuple with the DisableCrashReporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCrashReporting

`func (o *Cluster) SetDisableCrashReporting(v bool)`

SetDisableCrashReporting sets DisableCrashReporting field to given value.

### HasDisableCrashReporting

`func (o *Cluster) HasDisableCrashReporting() bool`

HasDisableCrashReporting returns a boolean if a field has been set.

### GetDisableVersionCheck

`func (o *Cluster) GetDisableVersionCheck() bool`

GetDisableVersionCheck returns the DisableVersionCheck field if non-nil, zero value otherwise.

### GetDisableVersionCheckOk

`func (o *Cluster) GetDisableVersionCheckOk() (*bool, bool)`

GetDisableVersionCheckOk returns a tuple with the DisableVersionCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableVersionCheck

`func (o *Cluster) SetDisableVersionCheck(v bool)`

SetDisableVersionCheck sets DisableVersionCheck field to given value.

### HasDisableVersionCheck

`func (o *Cluster) HasDisableVersionCheck() bool`

HasDisableVersionCheck returns a boolean if a field has been set.

### GetLogLevel

`func (o *Cluster) GetLogLevel() LogLevel`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *Cluster) GetLogLevelOk() (*LogLevel, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *Cluster) SetLogLevel(v LogLevel)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *Cluster) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetLogFormat

`func (o *Cluster) GetLogFormat() LogFormat`

GetLogFormat returns the LogFormat field if non-nil, zero value otherwise.

### GetLogFormatOk

`func (o *Cluster) GetLogFormatOk() (*LogFormat, bool)`

GetLogFormatOk returns a tuple with the LogFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFormat

`func (o *Cluster) SetLogFormat(v LogFormat)`

SetLogFormat sets LogFormat field to given value.

### HasLogFormat

`func (o *Cluster) HasLogFormat() bool`

HasLogFormat returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Cluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Cluster) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Cluster) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Cluster) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Cluster) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Cluster) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *Cluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Cluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Cluster) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Cluster) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


