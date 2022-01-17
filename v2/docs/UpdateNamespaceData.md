# UpdateNamespaceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **map[string]string** | A set of arbitrary key value labels to apply to the entity.  | [optional] 
**Version** | Pointer to **string** | An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change.  | [optional] 

## Methods

### NewUpdateNamespaceData

`func NewUpdateNamespaceData() *UpdateNamespaceData`

NewUpdateNamespaceData instantiates a new UpdateNamespaceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNamespaceDataWithDefaults

`func NewUpdateNamespaceDataWithDefaults() *UpdateNamespaceData`

NewUpdateNamespaceDataWithDefaults instantiates a new UpdateNamespaceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *UpdateNamespaceData) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateNamespaceData) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateNamespaceData) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateNamespaceData) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateNamespaceData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateNamespaceData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateNamespaceData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateNamespaceData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


