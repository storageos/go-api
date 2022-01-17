# NfsExportConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportID** | Pointer to **uint64** | ID for this export  | [optional] 
**Path** | Pointer to **string** | The path relative to the volume root to serve as the export root  | [optional] [default to ""]
**PseudoPath** | Pointer to **string** | The configured pseudo path in the NFS virtual filesystem. This is the path clients will see when traversing to this export on the NFS share.  | [optional] [default to ""]
**Acls** | Pointer to [**[]NfsAcl**](NfsAcl.md) |  | [optional] 

## Methods

### NewNfsExportConfig

`func NewNfsExportConfig() *NfsExportConfig`

NewNfsExportConfig instantiates a new NfsExportConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsExportConfigWithDefaults

`func NewNfsExportConfigWithDefaults() *NfsExportConfig`

NewNfsExportConfigWithDefaults instantiates a new NfsExportConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportID

`func (o *NfsExportConfig) GetExportID() uint64`

GetExportID returns the ExportID field if non-nil, zero value otherwise.

### GetExportIDOk

`func (o *NfsExportConfig) GetExportIDOk() (*uint64, bool)`

GetExportIDOk returns a tuple with the ExportID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportID

`func (o *NfsExportConfig) SetExportID(v uint64)`

SetExportID sets ExportID field to given value.

### HasExportID

`func (o *NfsExportConfig) HasExportID() bool`

HasExportID returns a boolean if a field has been set.

### GetPath

`func (o *NfsExportConfig) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *NfsExportConfig) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *NfsExportConfig) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *NfsExportConfig) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPseudoPath

`func (o *NfsExportConfig) GetPseudoPath() string`

GetPseudoPath returns the PseudoPath field if non-nil, zero value otherwise.

### GetPseudoPathOk

`func (o *NfsExportConfig) GetPseudoPathOk() (*string, bool)`

GetPseudoPathOk returns a tuple with the PseudoPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudoPath

`func (o *NfsExportConfig) SetPseudoPath(v string)`

SetPseudoPath sets PseudoPath field to given value.

### HasPseudoPath

`func (o *NfsExportConfig) HasPseudoPath() bool`

HasPseudoPath returns a boolean if a field has been set.

### GetAcls

`func (o *NfsExportConfig) GetAcls() []NfsAcl`

GetAcls returns the Acls field if non-nil, zero value otherwise.

### GetAclsOk

`func (o *NfsExportConfig) GetAclsOk() (*[]NfsAcl, bool)`

GetAclsOk returns a tuple with the Acls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcls

`func (o *NfsExportConfig) SetAcls(v []NfsAcl)`

SetAcls sets Acls field to given value.

### HasAcls

`func (o *NfsExportConfig) HasAcls() bool`

HasAcls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


