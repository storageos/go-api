# Licence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterID** | Pointer to **string** | A unique identifier for a cluster. The format of this type is undefined and may change but the defined properties will not change.  | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | The time after which a licence will no longer be valid This timestamp is set when the licence is created. String format is RFC3339.  | [optional] [readonly] 
**ClusterCapacityBytes** | Pointer to **uint64** | The allowed provisioning capacity in bytes This value if for the cluster, if provisioning a volume brings the cluster&#39;s total provisioned capacity above it the request will fail  | [optional] 
**UsedBytes** | Pointer to **uint64** | Sum of the size of all volumes in the cluster  | [optional] 
**Kind** | Pointer to **string** | Denotes which category the licence belongs to  | [optional] 
**CustomerName** | Pointer to **string** | A user friendly reference to the customer  | [optional] 
**Features** | Pointer to **[]string** | A list of product features which are enabled by the  licence, subject to the installed version.  | [optional] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewLicence

`func NewLicence() *Licence`

NewLicence instantiates a new Licence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenceWithDefaults

`func NewLicenceWithDefaults() *Licence`

NewLicenceWithDefaults instantiates a new Licence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterID

`func (o *Licence) GetClusterID() string`

GetClusterID returns the ClusterID field if non-nil, zero value otherwise.

### GetClusterIDOk

`func (o *Licence) GetClusterIDOk() (*string, bool)`

GetClusterIDOk returns a tuple with the ClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterID

`func (o *Licence) SetClusterID(v string)`

SetClusterID sets ClusterID field to given value.

### HasClusterID

`func (o *Licence) HasClusterID() bool`

HasClusterID returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Licence) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Licence) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Licence) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Licence) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetClusterCapacityBytes

`func (o *Licence) GetClusterCapacityBytes() uint64`

GetClusterCapacityBytes returns the ClusterCapacityBytes field if non-nil, zero value otherwise.

### GetClusterCapacityBytesOk

`func (o *Licence) GetClusterCapacityBytesOk() (*uint64, bool)`

GetClusterCapacityBytesOk returns a tuple with the ClusterCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCapacityBytes

`func (o *Licence) SetClusterCapacityBytes(v uint64)`

SetClusterCapacityBytes sets ClusterCapacityBytes field to given value.

### HasClusterCapacityBytes

`func (o *Licence) HasClusterCapacityBytes() bool`

HasClusterCapacityBytes returns a boolean if a field has been set.

### GetUsedBytes

`func (o *Licence) GetUsedBytes() uint64`

GetUsedBytes returns the UsedBytes field if non-nil, zero value otherwise.

### GetUsedBytesOk

`func (o *Licence) GetUsedBytesOk() (*uint64, bool)`

GetUsedBytesOk returns a tuple with the UsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBytes

`func (o *Licence) SetUsedBytes(v uint64)`

SetUsedBytes sets UsedBytes field to given value.

### HasUsedBytes

`func (o *Licence) HasUsedBytes() bool`

HasUsedBytes returns a boolean if a field has been set.

### GetKind

`func (o *Licence) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Licence) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Licence) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Licence) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCustomerName

`func (o *Licence) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *Licence) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *Licence) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *Licence) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetFeatures

`func (o *Licence) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Licence) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Licence) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Licence) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeaturesNil

`func (o *Licence) SetFeaturesNil(b bool)`

 SetFeaturesNil sets the value for Features to be an explicit nil

### UnsetFeatures
`func (o *Licence) UnsetFeatures()`

UnsetFeatures ensures that no value is present for Features, not even an explicit nil
### GetVersion

`func (o *Licence) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Licence) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Licence) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Licence) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


