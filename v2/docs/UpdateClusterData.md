# UpdateClusterData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableTelemetry** | Pointer to **bool** | Disables collection of telemetry data across the cluster.  | [optional] [default to false]
**DisableCrashReporting** | Pointer to **bool** | Disables collection of reports for any fatal crashes across the cluster.  | [optional] [default to false]
**DisableVersionCheck** | Pointer to **bool** | Disables the mechanism responsible for checking if there is an updated version of StorageOS available for installation.  | [optional] [default to false]
**LogLevel** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] [default to LOGLEVEL_INFO]
**LogFormat** | Pointer to [**LogFormat**](LogFormat.md) |  | [optional] [default to LOGFORMAT_DEFAULT]
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewUpdateClusterData

`func NewUpdateClusterData() *UpdateClusterData`

NewUpdateClusterData instantiates a new UpdateClusterData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterDataWithDefaults

`func NewUpdateClusterDataWithDefaults() *UpdateClusterData`

NewUpdateClusterDataWithDefaults instantiates a new UpdateClusterData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableTelemetry

`func (o *UpdateClusterData) GetDisableTelemetry() bool`

GetDisableTelemetry returns the DisableTelemetry field if non-nil, zero value otherwise.

### GetDisableTelemetryOk

`func (o *UpdateClusterData) GetDisableTelemetryOk() (*bool, bool)`

GetDisableTelemetryOk returns a tuple with the DisableTelemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTelemetry

`func (o *UpdateClusterData) SetDisableTelemetry(v bool)`

SetDisableTelemetry sets DisableTelemetry field to given value.

### HasDisableTelemetry

`func (o *UpdateClusterData) HasDisableTelemetry() bool`

HasDisableTelemetry returns a boolean if a field has been set.

### GetDisableCrashReporting

`func (o *UpdateClusterData) GetDisableCrashReporting() bool`

GetDisableCrashReporting returns the DisableCrashReporting field if non-nil, zero value otherwise.

### GetDisableCrashReportingOk

`func (o *UpdateClusterData) GetDisableCrashReportingOk() (*bool, bool)`

GetDisableCrashReportingOk returns a tuple with the DisableCrashReporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCrashReporting

`func (o *UpdateClusterData) SetDisableCrashReporting(v bool)`

SetDisableCrashReporting sets DisableCrashReporting field to given value.

### HasDisableCrashReporting

`func (o *UpdateClusterData) HasDisableCrashReporting() bool`

HasDisableCrashReporting returns a boolean if a field has been set.

### GetDisableVersionCheck

`func (o *UpdateClusterData) GetDisableVersionCheck() bool`

GetDisableVersionCheck returns the DisableVersionCheck field if non-nil, zero value otherwise.

### GetDisableVersionCheckOk

`func (o *UpdateClusterData) GetDisableVersionCheckOk() (*bool, bool)`

GetDisableVersionCheckOk returns a tuple with the DisableVersionCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableVersionCheck

`func (o *UpdateClusterData) SetDisableVersionCheck(v bool)`

SetDisableVersionCheck sets DisableVersionCheck field to given value.

### HasDisableVersionCheck

`func (o *UpdateClusterData) HasDisableVersionCheck() bool`

HasDisableVersionCheck returns a boolean if a field has been set.

### GetLogLevel

`func (o *UpdateClusterData) GetLogLevel() LogLevel`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *UpdateClusterData) GetLogLevelOk() (*LogLevel, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *UpdateClusterData) SetLogLevel(v LogLevel)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *UpdateClusterData) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetLogFormat

`func (o *UpdateClusterData) GetLogFormat() LogFormat`

GetLogFormat returns the LogFormat field if non-nil, zero value otherwise.

### GetLogFormatOk

`func (o *UpdateClusterData) GetLogFormatOk() (*LogFormat, bool)`

GetLogFormatOk returns a tuple with the LogFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFormat

`func (o *UpdateClusterData) SetLogFormat(v LogFormat)`

SetLogFormat sets LogFormat field to given value.

### HasLogFormat

`func (o *UpdateClusterData) HasLogFormat() bool`

HasLogFormat returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateClusterData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateClusterData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateClusterData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateClusterData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


