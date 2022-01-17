# UpdateLicence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | A StorageOS product licence key, used to register a cluster. The format of this type is opaque and may change.  | [optional] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewUpdateLicence

`func NewUpdateLicence() *UpdateLicence`

NewUpdateLicence instantiates a new UpdateLicence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLicenceWithDefaults

`func NewUpdateLicenceWithDefaults() *UpdateLicence`

NewUpdateLicenceWithDefaults instantiates a new UpdateLicence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *UpdateLicence) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateLicence) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateLicence) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateLicence) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateLicence) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateLicence) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateLicence) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateLicence) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


