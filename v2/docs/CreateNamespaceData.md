# CreateNamespaceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the namespace shown in the CLI and UI  | [optional] 
**Labels** | Pointer to **map[string]string** | A set of arbitrary key value labels to apply to the entity.  | [optional] 

## Methods

### NewCreateNamespaceData

`func NewCreateNamespaceData() *CreateNamespaceData`

NewCreateNamespaceData instantiates a new CreateNamespaceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNamespaceDataWithDefaults

`func NewCreateNamespaceDataWithDefaults() *CreateNamespaceData`

NewCreateNamespaceDataWithDefaults instantiates a new CreateNamespaceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNamespaceData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNamespaceData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNamespaceData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNamespaceData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *CreateNamespaceData) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateNamespaceData) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateNamespaceData) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateNamespaceData) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


