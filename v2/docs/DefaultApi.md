# \DefaultApi

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachNFSVolume**](DefaultApi.md#AttachNFSVolume) | **Post** /namespaces/{namespaceID}/volumes/{id}/nfs/attach | attach and share the volume using NFS
[**AttachVolume**](DefaultApi.md#AttachVolume) | **Post** /namespaces/{namespaceID}/volumes/{id}/attach | Attach a volume to the given node
[**AuthenticateUser**](DefaultApi.md#AuthenticateUser) | **Post** /auth/login | Authenticate a user
[**CreateNamespace**](DefaultApi.md#CreateNamespace) | **Post** /namespaces | Create a new namespace
[**CreatePolicyGroup**](DefaultApi.md#CreatePolicyGroup) | **Post** /policies | Create a new policy group
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /users | Create a new user
[**CreateVolume**](DefaultApi.md#CreateVolume) | **Post** /namespaces/{namespaceID}/volumes | Create a new Volume in the specified namespace
[**DeleteAuthenticatedUser**](DefaultApi.md#DeleteAuthenticatedUser) | **Delete** /users/self | Delete the authenticated user
[**DeleteAuthenticatedUserSessions**](DefaultApi.md#DeleteAuthenticatedUserSessions) | **Delete** /users/self/sessions | Invalidate the logged in user&#39;s sessions
[**DeleteNamespace**](DefaultApi.md#DeleteNamespace) | **Delete** /namespaces/{id} | Delete a namespace
[**DeleteNode**](DefaultApi.md#DeleteNode) | **Delete** /nodes/{id} | Delete a node
[**DeletePolicyGroup**](DefaultApi.md#DeletePolicyGroup) | **Delete** /policies/{id} | Delete a policy group
[**DeleteSessions**](DefaultApi.md#DeleteSessions) | **Delete** /users/{id}/sessions | Invalidate login sessions
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /users/{id} | Delete a user
[**DeleteVolume**](DefaultApi.md#DeleteVolume) | **Delete** /namespaces/{namespaceID}/volumes/{id} | Delete a volume
[**DetachVolume**](DefaultApi.md#DetachVolume) | **Delete** /namespaces/{namespaceID}/volumes/{id}/attach | Detach the given volume
[**GetAuthenticatedUser**](DefaultApi.md#GetAuthenticatedUser) | **Get** /users/self | Get the currently authenticated user&#39;s information
[**GetCluster**](DefaultApi.md#GetCluster) | **Get** /cluster | Retrieves the cluster&#39;s global configuration settings
[**GetDiagnostics**](DefaultApi.md#GetDiagnostics) | **Get** /diagnostics | Retrieves a diagnostics bundle from the target node
[**GetLicence**](DefaultApi.md#GetLicence) | **Get** /cluster/licence | Retrieves the cluster&#39;s licence information
[**GetNamespace**](DefaultApi.md#GetNamespace) | **Get** /namespaces/{id} | Fetch a namespace
[**GetNode**](DefaultApi.md#GetNode) | **Get** /nodes/{id} | Fetch a node
[**GetPolicyGroup**](DefaultApi.md#GetPolicyGroup) | **Get** /policies/{id} | Fetch a policy group
[**GetSingleNodeDiagnostics**](DefaultApi.md#GetSingleNodeDiagnostics) | **Get** /diagnostics/{id} | Retrieves a single node diagnostics bundle from the target node
[**GetUser**](DefaultApi.md#GetUser) | **Get** /users/{id} | Fetch a user
[**GetVolume**](DefaultApi.md#GetVolume) | **Get** /namespaces/{namespaceID}/volumes/{id} | Fetch a volume
[**ListNamespaces**](DefaultApi.md#ListNamespaces) | **Get** /namespaces | Fetch the list of namespaces
[**ListNodes**](DefaultApi.md#ListNodes) | **Get** /nodes | Fetch the list of nodes
[**ListPolicyGroups**](DefaultApi.md#ListPolicyGroups) | **Get** /policies | Fetch the list of policy groups
[**ListUsers**](DefaultApi.md#ListUsers) | **Get** /users | Fetch the list of users
[**ListVolumes**](DefaultApi.md#ListVolumes) | **Get** /namespaces/{namespaceID}/volumes | Fetch the list of volumes in the given namespace
[**RefreshJwt**](DefaultApi.md#RefreshJwt) | **Post** /auth/refresh | Refresh the JWT
[**ResizeVolume**](DefaultApi.md#ResizeVolume) | **Put** /namespaces/{namespaceID}/volumes/{id}/size | Increase the size of a volume.
[**SetComputeOnly**](DefaultApi.md#SetComputeOnly) | **Put** /nodes/{id}/compute-only | Modify the computeonly behaviour state for a node
[**SetCordoned**](DefaultApi.md#SetCordoned) | **Put** /nodes/{id}/cordon | Modify the cordoned state for a node
[**SetFailureMode**](DefaultApi.md#SetFailureMode) | **Put** /namespaces/{namespaceID}/volumes/{id}/failure-mode | Set the failure mode of the volume.
[**SetPlacementStrategy**](DefaultApi.md#SetPlacementStrategy) | **Put** /namespaces/{namespaceID}/volumes/{id}/placement-strategy | Sets the placement strategy of the volume.
[**SetReplicas**](DefaultApi.md#SetReplicas) | **Put** /namespaces/{namespaceID}/volumes/{id}/replicas | Set the number of replicas to maintain for the volume.
[**Spec**](DefaultApi.md#Spec) | **Get** /openapi | Serves this openapi spec file
[**UpdateAuthenticatedUser**](DefaultApi.md#UpdateAuthenticatedUser) | **Put** /users/self | Update the authenticated user&#39;s information
[**UpdateCluster**](DefaultApi.md#UpdateCluster) | **Put** /cluster | Update the cluster&#39;s global configuration settings
[**UpdateLicence**](DefaultApi.md#UpdateLicence) | **Put** /cluster/licence | Update the licence global configuration settings
[**UpdateNFSVolumeExports**](DefaultApi.md#UpdateNFSVolumeExports) | **Put** /namespaces/{namespaceID}/volumes/{id}/nfs/export-config | Update an nfs volume&#39;s export configuration
[**UpdateNFSVolumeMountEndpoint**](DefaultApi.md#UpdateNFSVolumeMountEndpoint) | **Put** /namespaces/{namespaceID}/volumes/{id}/nfs/mount-endpoint | Update an nfs volume&#39;s mount endpoint
[**UpdateNamespace**](DefaultApi.md#UpdateNamespace) | **Put** /namespaces/{id} | Update a namespace
[**UpdateNode**](DefaultApi.md#UpdateNode) | **Put** /nodes/{id} | Update a node
[**UpdatePolicyGroup**](DefaultApi.md#UpdatePolicyGroup) | **Put** /policies/{id} | Update a policy group
[**UpdateUser**](DefaultApi.md#UpdateUser) | **Put** /users/{id} | Update a user
[**UpdateVolume**](DefaultApi.md#UpdateVolume) | **Put** /namespaces/{namespaceID}/volumes/{id} | Update a volume



## AttachNFSVolume

> AttachNFSVolume(ctx, namespaceID, id).AttachNFSVolumeData(attachNFSVolumeData).IgnoreVersion(ignoreVersion).AsyncMax(asyncMax).Execute()

attach and share the volume using NFS



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespaceID := "namespaceID_example" // string | ID of a Namespace
    id := "id_example" // string | ID of a Volume
    attachNFSVolumeData := *openapiclient.NewAttachNFSVolumeData() // AttachNFSVolumeData | 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)
    asyncMax := "asyncMax_example" // string | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \"async-max\" header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \"300ms\", or \"2h45m\". Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\". We reject negative or nil duration values.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AttachNFSVolume(context.Background(), namespaceID, id).AttachNFSVolumeData(attachNFSVolumeData).IgnoreVersion(ignoreVersion).AsyncMax(asyncMax).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AttachNFSVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceID** | **string** | ID of a Namespace | 
**id** | **string** | ID of a Volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachNFSVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attachNFSVolumeData** | [**AttachNFSVolumeData**](AttachNFSVolumeData.md) |  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]
 **asyncMax** | **string** | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \&quot;async-max\&quot; header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \&quot;300ms\&quot;, or \&quot;2h45m\&quot;. Valid time units are \&quot;ns\&quot;, \&quot;us\&quot; (or \&quot;µs\&quot;), \&quot;ms\&quot;, \&quot;s\&quot;, \&quot;m\&quot;, \&quot;h\&quot;. We reject negative or nil duration values.  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachVolume

> AttachVolume(ctx, namespaceID, id).AttachVolumeData(attachVolumeData).Execute()

Attach a volume to the given node



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespaceID := "namespaceID_example" // string | ID of a Namespace
    id := "id_example" // string | ID of a Volume
    attachVolumeData := *openapiclient.NewAttachVolumeData() // AttachVolumeData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AttachVolume(context.Background(), namespaceID, id).AttachVolumeData(attachVolumeData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AttachVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceID** | **string** | ID of a Namespace | 
**id** | **string** | ID of a Volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attachVolumeData** | [**AttachVolumeData**](AttachVolumeData.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticateUser

> UserSession AuthenticateUser(ctx).AuthUserData(authUserData).Execute()

Authenticate a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authUserData := *openapiclient.NewAuthUserData("Username_example", "Password_example") // AuthUserData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AuthenticateUser(context.Background()).AuthUserData(authUserData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AuthenticateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticateUser`: UserSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AuthenticateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authUserData** | [**AuthUserData**](AuthUserData.md) |  | 

### Return type

[**UserSession**](UserSession.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNamespace

> Namespace CreateNamespace(ctx).CreateNamespaceData(createNamespaceData).Execute()

Create a new namespace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createNamespaceData := *openapiclient.NewCreateNamespaceData() // CreateNamespaceData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateNamespace(context.Background()).CreateNamespaceData(createNamespaceData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNamespace`: Namespace
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNamespace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNamespaceData** | [**CreateNamespaceData**](CreateNamespaceData.md) |  | 

### Return type

[**Namespace**](Namespace.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicyGroup

> PolicyGroup CreatePolicyGroup(ctx).CreatePolicyGroupData(createPolicyGroupData).Execute()

Create a new policy group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createPolicyGroupData := *openapiclient.NewCreatePolicyGroupData() // CreatePolicyGroupData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreatePolicyGroup(context.Background()).CreatePolicyGroupData(createPolicyGroupData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreatePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyGroup`: PolicyGroup
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreatePolicyGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPolicyGroupData** | [**CreatePolicyGroupData**](CreatePolicyGroupData.md) |  | 

### Return type

[**PolicyGroup**](PolicyGroup.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> User CreateUser(ctx).CreateUserData(createUserData).Execute()

Create a new user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createUserData := *openapiclient.NewCreateUserData("admin", "turtlesaregreat") // CreateUserData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateUser(context.Background()).CreateUserData(createUserData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserData** | [**CreateUserData**](CreateUserData.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVolume

> Volume CreateVolume(ctx, namespaceID).CreateVolumeData(createVolumeData).AsyncMax(asyncMax).Execute()

Create a new Volume in the specified namespace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespaceID := "namespaceID_example" // string | ID of a Namespace
    createVolumeData := *openapiclient.NewCreateVolumeData("c5666b58-b805-4215-ab4a-cb094948ccc6", "data", openapiclient.FsType("ext2"), uint64(5000)) // CreateVolumeData | 
    asyncMax := "asyncMax_example" // string | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \"async-max\" header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \"300ms\", or \"2h45m\". Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\". We reject negative or nil duration values.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateVolume(context.Background(), namespaceID).CreateVolumeData(createVolumeData).AsyncMax(asyncMax).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVolume`: Volume
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceID** | **string** | ID of a Namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVolumeData** | [**CreateVolumeData**](CreateVolumeData.md) |  | 
 **asyncMax** | **string** | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \&quot;async-max\&quot; header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \&quot;300ms\&quot;, or \&quot;2h45m\&quot;. Valid time units are \&quot;ns\&quot;, \&quot;us\&quot; (or \&quot;µs\&quot;), \&quot;ms\&quot;, \&quot;s\&quot;, \&quot;m\&quot;, \&quot;h\&quot;. We reject negative or nil duration values.  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthenticatedUser

> DeleteAuthenticatedUser(ctx).Version(version).IgnoreVersion(ignoreVersion).Execute()

Delete the authenticated user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    version := "version_example" // string | This value is used to perform a conditional delete or update of the entity. If the entity has been modified since the version token was obtained, the request will fail with a HTTP 409 Conflict. 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAuthenticatedUser(context.Background()).Version(version).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | This value is used to perform a conditional delete or update of the entity. If the entity has been modified since the version token was obtained, the request will fail with a HTTP 409 Conflict.  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthenticatedUserSessions

> DeleteAuthenticatedUserSessions(ctx).Execute()

Invalidate the logged in user's sessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAuthenticatedUserSessions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAuthenticatedUserSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthenticatedUserSessionsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamespace

> DeleteNamespace(ctx, id).Version(version).IgnoreVersion(ignoreVersion).Execute()

Delete a namespace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a namespace
    version := "version_example" // string | This value is used to perform a conditional delete or update of the entity. If the entity has been modified since the version token was obtained, the request will fail with a HTTP 409 Conflict. 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteNamespace(context.Background(), id).Version(version).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | This value is used to perform a conditional delete or update of the entity. If the entity has been modified since the version token was obtained, the request will fail with a HTTP 409 Conflict.  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNode

> DeleteNode(ctx, id).Version(version).IgnoreVersion(ignoreVersion).AsyncMax(asyncMax).Execute()

Delete a node



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a node
    version := "version_example" // string | This value is used to perform a conditional delete or update of the entity. If the entity has been modified since the version token was obtained, the request will fail with a HTTP 409 Conflict. 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)
    asyncMax := "asyncMax_example" // string | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \"async-max\" header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \"300ms\", or \"2h45m\". Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\". We reject negative or nil duration values.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteNode(context.Background(), id).Version(version).IgnoreVersion(ignoreVersion).AsyncMax(asyncMax).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a node | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | This value is used to perform a conditional delete or update of the entity. If the entity has been modified since the version token was obtained, the request will fail with a HTTP 409 Conflict.  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]
 **asyncMax** | **string** | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \&quot;async-max\&quot; header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \&quot;300ms\&quot;, or \&quot;2h45m\&quot;. Valid time units are \&quot;ns\&quot;, \&quot;us\&quot; (or \&quot;µs\&quot;), \&quot;ms\&quot;, \&quot;s\&quot;, \&quot;m\&quot;, \&quot;h\&quot;. We reject negative or nil duration values.  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicyGroup

> DeletePolicyGroup(ctx, id).Version(version).IgnoreVersion(ignoreVersion).Execute()

Delete a policy group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a policy group
    version := "version_example" // string | This value is used to perform a conditional delete or update of the entity. If the entity has been modified since the version token was obtained, the request will fail with a HTTP 409 Conflict. 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeletePolicyGroup(context.Background(), id).Version(version).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeletePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a policy group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | This value is used to perform a conditional delete or update of the entity. If the entity has been modified since the version token was obtained, the request will fail with a HTTP 409 Conflict.  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSessions

> DeleteSessions(ctx, id).Execute()

Invalidate login sessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a user

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSessions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a user | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, id).Version(version).IgnoreVersion(ignoreVersion).Execute()

Delete a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a user
    version := "version_example" // string | This value is used to perform a conditional delete or update of the entity. If the entity has been modified since the version token was obtained, the request will fail with a HTTP 409 Conflict. 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteUser(context.Background(), id).Version(version).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a user | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | This value is used to perform a conditional delete or update of the entity. If the entity has been modified since the version token was obtained, the request will fail with a HTTP 409 Conflict.  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolume

> DeleteVolume(ctx, namespaceID, id).Version(version).IgnoreVersion(ignoreVersion).AsyncMax(asyncMax).OfflineDelete(offlineDelete).Execute()

Delete a volume



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespaceID := "namespaceID_example" // string | ID of a Namespace
    id := "id_example" // string | ID of a Volume
    version := "version_example" // string | This value is used to perform a conditional delete or update of the entity. If the entity has been modified since the version token was obtained, the request will fail with a HTTP 409 Conflict. 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)
    asyncMax := "asyncMax_example" // string | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \"async-max\" header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \"300ms\", or \"2h45m\". Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\". We reject negative or nil duration values.  (optional)
    offlineDelete := true // bool | If set to true, enables deletion of a volume when all  deployments are offline, bypassing the host nodes which cannot be reached. An offline delete request will be rejected when either a) there are online deployments for the target volume or b) there is evidence that an unreachable node still has the volume master  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteVolume(context.Background(), namespaceID, id).Version(version).IgnoreVersion(ignoreVersion).AsyncMax(asyncMax).OfflineDelete(offlineDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceID** | **string** | ID of a Namespace | 
**id** | **string** | ID of a Volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | This value is used to perform a conditional delete or update of the entity. If the entity has been modified since the version token was obtained, the request will fail with a HTTP 409 Conflict.  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]
 **asyncMax** | **string** | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \&quot;async-max\&quot; header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \&quot;300ms\&quot;, or \&quot;2h45m\&quot;. Valid time units are \&quot;ns\&quot;, \&quot;us\&quot; (or \&quot;µs\&quot;), \&quot;ms\&quot;, \&quot;s\&quot;, \&quot;m\&quot;, \&quot;h\&quot;. We reject negative or nil duration values.  | 
 **offlineDelete** | **bool** | If set to true, enables deletion of a volume when all  deployments are offline, bypassing the host nodes which cannot be reached. An offline delete request will be rejected when either a) there are online deployments for the target volume or b) there is evidence that an unreachable node still has the volume master  | [default to false]

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachVolume

> DetachVolume(ctx, namespaceID, id).Version(version).IgnoreVersion(ignoreVersion).AsyncMax(asyncMax).Execute()

Detach the given volume



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespaceID := "namespaceID_example" // string | ID of a Namespace
    id := "id_example" // string | ID of a Volume
    version := "version_example" // string | This value is used to perform a conditional delete or update of the entity. If the entity has been modified since the version token was obtained, the request will fail with a HTTP 409 Conflict. 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)
    asyncMax := "asyncMax_example" // string | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \"async-max\" header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \"300ms\", or \"2h45m\". Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\". We reject negative or nil duration values.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DetachVolume(context.Background(), namespaceID, id).Version(version).IgnoreVersion(ignoreVersion).AsyncMax(asyncMax).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DetachVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceID** | **string** | ID of a Namespace | 
**id** | **string** | ID of a Volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | This value is used to perform a conditional delete or update of the entity. If the entity has been modified since the version token was obtained, the request will fail with a HTTP 409 Conflict.  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]
 **asyncMax** | **string** | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \&quot;async-max\&quot; header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \&quot;300ms\&quot;, or \&quot;2h45m\&quot;. Valid time units are \&quot;ns\&quot;, \&quot;us\&quot; (or \&quot;µs\&quot;), \&quot;ms\&quot;, \&quot;s\&quot;, \&quot;m\&quot;, \&quot;h\&quot;. We reject negative or nil duration values.  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticatedUser

> User GetAuthenticatedUser(ctx).Execute()

Get the currently authenticated user's information



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAuthenticatedUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticatedUser`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticatedUserRequest struct via the builder pattern


### Return type

[**User**](User.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCluster

> Cluster GetCluster(ctx).Execute()

Retrieves the cluster's global configuration settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetCluster(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCluster`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterRequest struct via the builder pattern


### Return type

[**Cluster**](Cluster.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiagnostics

> *os.File GetDiagnostics(ctx).Execute()

Retrieves a diagnostics bundle from the target node



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetDiagnostics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDiagnostics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiagnostics`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDiagnostics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiagnosticsRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicence

> Licence GetLicence(ctx).Execute()

Retrieves the cluster's licence information



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetLicence(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetLicence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicence`: Licence
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetLicence`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenceRequest struct via the builder pattern


### Return type

[**Licence**](Licence.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespace

> Namespace GetNamespace(ctx, id).Execute()

Fetch a namespace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetNamespace(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespace`: Namespace
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Namespace**](Namespace.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNode

> Node GetNode(ctx, id).Execute()

Fetch a node



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a node

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetNode(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNode`: Node
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a node | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Node**](Node.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyGroup

> PolicyGroup GetPolicyGroup(ctx, id).Execute()

Fetch a policy group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a policy group

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetPolicyGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyGroup`: PolicyGroup
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a policy group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyGroup**](PolicyGroup.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleNodeDiagnostics

> *os.File GetSingleNodeDiagnostics(ctx, id).Execute()

Retrieves a single node diagnostics bundle from the target node



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a node

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSingleNodeDiagnostics(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSingleNodeDiagnostics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleNodeDiagnostics`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSingleNodeDiagnostics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a node | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleNodeDiagnosticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, id).Execute()

Fetch a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a user

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolume

> Volume GetVolume(ctx, namespaceID, id).Execute()

Fetch a volume



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespaceID := "namespaceID_example" // string | ID of a Namespace
    id := "id_example" // string | ID of a Volume

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetVolume(context.Background(), namespaceID, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVolume`: Volume
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceID** | **string** | ID of a Namespace | 
**id** | **string** | ID of a Volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Volume**](Volume.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespaces

> []Namespace ListNamespaces(ctx).Execute()

Fetch the list of namespaces



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListNamespaces(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNamespaces`: []Namespace
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListNamespaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacesRequest struct via the builder pattern


### Return type

[**[]Namespace**](Namespace.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodes

> []Node ListNodes(ctx).Execute()

Fetch the list of nodes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListNodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNodes`: []Node
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListNodesRequest struct via the builder pattern


### Return type

[**[]Node**](Node.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicyGroups

> []PolicyGroup ListPolicyGroups(ctx).Execute()

Fetch the list of policy groups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListPolicyGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListPolicyGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicyGroups`: []PolicyGroup
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListPolicyGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPolicyGroupsRequest struct via the builder pattern


### Return type

[**[]PolicyGroup**](PolicyGroup.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> []User ListUsers(ctx).Execute()

Fetch the list of users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: []User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


### Return type

[**[]User**](User.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumes

> []Volume ListVolumes(ctx, namespaceID).Execute()

Fetch the list of volumes in the given namespace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespaceID := "namespaceID_example" // string | ID of a Namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListVolumes(context.Background(), namespaceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListVolumes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVolumes`: []Volume
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceID** | **string** | ID of a Namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Volume**](Volume.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshJwt

> UserSession RefreshJwt(ctx).Execute()

Refresh the JWT



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RefreshJwt(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RefreshJwt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshJwt`: UserSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RefreshJwt`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshJwtRequest struct via the builder pattern


### Return type

[**UserSession**](UserSession.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResizeVolume

> Volume ResizeVolume(ctx, namespaceID, id).ResizeVolumeRequest(resizeVolumeRequest).AsyncMax(asyncMax).IgnoreVersion(ignoreVersion).Execute()

Increase the size of a volume.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespaceID := "namespaceID_example" // string | ID of a Namespace
    id := "id_example" // string | ID of a Volume
    resizeVolumeRequest := *openapiclient.NewResizeVolumeRequest() // ResizeVolumeRequest | 
    asyncMax := "asyncMax_example" // string | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \"async-max\" header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \"300ms\", or \"2h45m\". Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\". We reject negative or nil duration values.  (optional)
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ResizeVolume(context.Background(), namespaceID, id).ResizeVolumeRequest(resizeVolumeRequest).AsyncMax(asyncMax).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResizeVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResizeVolume`: Volume
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ResizeVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceID** | **string** | ID of a Namespace | 
**id** | **string** | ID of a Volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resizeVolumeRequest** | [**ResizeVolumeRequest**](ResizeVolumeRequest.md) |  | 
 **asyncMax** | **string** | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \&quot;async-max\&quot; header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \&quot;300ms\&quot;, or \&quot;2h45m\&quot;. Valid time units are \&quot;ns\&quot;, \&quot;us\&quot; (or \&quot;µs\&quot;), \&quot;ms\&quot;, \&quot;s\&quot;, \&quot;m\&quot;, \&quot;h\&quot;. We reject negative or nil duration values.  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

[**Volume**](Volume.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetComputeOnly

> Node SetComputeOnly(ctx, id).SetComputeOnlyNodeData(setComputeOnlyNodeData).IgnoreVersion(ignoreVersion).Execute()

Modify the computeonly behaviour state for a node



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a Node
    setComputeOnlyNodeData := *openapiclient.NewSetComputeOnlyNodeData() // SetComputeOnlyNodeData | 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SetComputeOnly(context.Background(), id).SetComputeOnlyNodeData(setComputeOnlyNodeData).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetComputeOnly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetComputeOnly`: Node
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SetComputeOnly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a Node | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetComputeOnlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setComputeOnlyNodeData** | [**SetComputeOnlyNodeData**](SetComputeOnlyNodeData.md) |  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

[**Node**](Node.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCordoned

> Node SetCordoned(ctx, id).SetCordonedNodeData(setCordonedNodeData).IgnoreVersion(ignoreVersion).Execute()

Modify the cordoned state for a node



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a Node
    setCordonedNodeData := *openapiclient.NewSetCordonedNodeData() // SetCordonedNodeData | 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SetCordoned(context.Background(), id).SetCordonedNodeData(setCordonedNodeData).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetCordoned``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetCordoned`: Node
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SetCordoned`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a Node | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCordonedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setCordonedNodeData** | [**SetCordonedNodeData**](SetCordonedNodeData.md) |  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

[**Node**](Node.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFailureMode

> Volume SetFailureMode(ctx, namespaceID, id).SetFailureModeRequest(setFailureModeRequest).IgnoreVersion(ignoreVersion).Execute()

Set the failure mode of the volume.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespaceID := "namespaceID_example" // string | ID of a Namespace
    id := "id_example" // string | ID of a Volume
    setFailureModeRequest := openapiclient.SetFailureModeRequest{SetFailureModeIntentRequestData: openapiclient.NewSetFailureModeIntentRequestData()} // SetFailureModeRequest | Failure mode to use
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SetFailureMode(context.Background(), namespaceID, id).SetFailureModeRequest(setFailureModeRequest).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetFailureMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetFailureMode`: Volume
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SetFailureMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceID** | **string** | ID of a Namespace | 
**id** | **string** | ID of a Volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFailureModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setFailureModeRequest** | [**SetFailureModeRequest**](SetFailureModeRequest.md) | Failure mode to use | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

[**Volume**](Volume.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPlacementStrategy

> AcceptedMessage SetPlacementStrategy(ctx, namespaceID, id).SetPlacementStrategyData(setPlacementStrategyData).IgnoreVersion(ignoreVersion).Execute()

Sets the placement strategy of the volume.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespaceID := "namespaceID_example" // string | ID of a Namespace
    id := "id_example" // string | ID of a Volume
    setPlacementStrategyData := *openapiclient.NewSetPlacementStrategyData() // SetPlacementStrategyData | 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SetPlacementStrategy(context.Background(), namespaceID, id).SetPlacementStrategyData(setPlacementStrategyData).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetPlacementStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPlacementStrategy`: AcceptedMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SetPlacementStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceID** | **string** | ID of a Namespace | 
**id** | **string** | ID of a Volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPlacementStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setPlacementStrategyData** | [**SetPlacementStrategyData**](SetPlacementStrategyData.md) |  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

[**AcceptedMessage**](AcceptedMessage.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetReplicas

> AcceptedMessage SetReplicas(ctx, namespaceID, id).SetReplicasRequest(setReplicasRequest).IgnoreVersion(ignoreVersion).Execute()

Set the number of replicas to maintain for the volume.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespaceID := "namespaceID_example" // string | ID of a Namespace
    id := "id_example" // string | ID of a Volume
    setReplicasRequest := *openapiclient.NewSetReplicasRequest() // SetReplicasRequest | 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SetReplicas(context.Background(), namespaceID, id).SetReplicasRequest(setReplicasRequest).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetReplicas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetReplicas`: AcceptedMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SetReplicas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceID** | **string** | ID of a Namespace | 
**id** | **string** | ID of a Volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetReplicasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setReplicasRequest** | [**SetReplicasRequest**](SetReplicasRequest.md) |  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

[**AcceptedMessage**](AcceptedMessage.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Spec

> string Spec(ctx).Execute()

Serves this openapi spec file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Spec(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Spec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Spec`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Spec`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSpecRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthenticatedUser

> User UpdateAuthenticatedUser(ctx).UpdateAuthenticatedUserData(updateAuthenticatedUserData).IgnoreVersion(ignoreVersion).Execute()

Update the authenticated user's information



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateAuthenticatedUserData := *openapiclient.NewUpdateAuthenticatedUserData() // UpdateAuthenticatedUserData | 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAuthenticatedUser(context.Background()).UpdateAuthenticatedUserData(updateAuthenticatedUserData).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthenticatedUser`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAuthenticatedUserData** | [**UpdateAuthenticatedUserData**](UpdateAuthenticatedUserData.md) |  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

[**User**](User.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCluster

> Cluster UpdateCluster(ctx).UpdateClusterData(updateClusterData).IgnoreVersion(ignoreVersion).Execute()

Update the cluster's global configuration settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateClusterData := *openapiclient.NewUpdateClusterData() // UpdateClusterData | 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateCluster(context.Background()).UpdateClusterData(updateClusterData).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateClusterData** | [**UpdateClusterData**](UpdateClusterData.md) |  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

[**Cluster**](Cluster.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLicence

> Licence UpdateLicence(ctx).UpdateLicence(updateLicence).IgnoreVersion(ignoreVersion).Execute()

Update the licence global configuration settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateLicence := *openapiclient.NewUpdateLicence() // UpdateLicence | 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateLicence(context.Background()).UpdateLicence(updateLicence).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateLicence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLicence`: Licence
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateLicence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLicenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateLicence** | [**UpdateLicence**](UpdateLicence.md) |  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

[**Licence**](Licence.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNFSVolumeExports

> UpdateNFSVolumeExports(ctx, namespaceID, id).NFSVolumeExports(nFSVolumeExports).IgnoreVersion(ignoreVersion).AsyncMax(asyncMax).Execute()

Update an nfs volume's export configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespaceID := "namespaceID_example" // string | ID of a Namespace
    id := "id_example" // string | ID of a Volume
    nFSVolumeExports := *openapiclient.NewNFSVolumeExports() // NFSVolumeExports | 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)
    asyncMax := "asyncMax_example" // string | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \"async-max\" header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \"300ms\", or \"2h45m\". Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\". We reject negative or nil duration values.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateNFSVolumeExports(context.Background(), namespaceID, id).NFSVolumeExports(nFSVolumeExports).IgnoreVersion(ignoreVersion).AsyncMax(asyncMax).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateNFSVolumeExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceID** | **string** | ID of a Namespace | 
**id** | **string** | ID of a Volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNFSVolumeExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nFSVolumeExports** | [**NFSVolumeExports**](NFSVolumeExports.md) |  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]
 **asyncMax** | **string** | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \&quot;async-max\&quot; header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \&quot;300ms\&quot;, or \&quot;2h45m\&quot;. Valid time units are \&quot;ns\&quot;, \&quot;us\&quot; (or \&quot;µs\&quot;), \&quot;ms\&quot;, \&quot;s\&quot;, \&quot;m\&quot;, \&quot;h\&quot;. We reject negative or nil duration values.  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNFSVolumeMountEndpoint

> UpdateNFSVolumeMountEndpoint(ctx, namespaceID, id).NFSVolumeMountEndpoint(nFSVolumeMountEndpoint).IgnoreVersion(ignoreVersion).AsyncMax(asyncMax).Execute()

Update an nfs volume's mount endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespaceID := "namespaceID_example" // string | ID of a Namespace
    id := "id_example" // string | ID of a Volume
    nFSVolumeMountEndpoint := *openapiclient.NewNFSVolumeMountEndpoint() // NFSVolumeMountEndpoint | 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)
    asyncMax := "asyncMax_example" // string | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \"async-max\" header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \"300ms\", or \"2h45m\". Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\". We reject negative or nil duration values.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateNFSVolumeMountEndpoint(context.Background(), namespaceID, id).NFSVolumeMountEndpoint(nFSVolumeMountEndpoint).IgnoreVersion(ignoreVersion).AsyncMax(asyncMax).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateNFSVolumeMountEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceID** | **string** | ID of a Namespace | 
**id** | **string** | ID of a Volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNFSVolumeMountEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nFSVolumeMountEndpoint** | [**NFSVolumeMountEndpoint**](NFSVolumeMountEndpoint.md) |  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]
 **asyncMax** | **string** | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \&quot;async-max\&quot; header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \&quot;300ms\&quot;, or \&quot;2h45m\&quot;. Valid time units are \&quot;ns\&quot;, \&quot;us\&quot; (or \&quot;µs\&quot;), \&quot;ms\&quot;, \&quot;s\&quot;, \&quot;m\&quot;, \&quot;h\&quot;. We reject negative or nil duration values.  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNamespace

> Namespace UpdateNamespace(ctx, id).UpdateNamespaceData(updateNamespaceData).IgnoreVersion(ignoreVersion).Execute()

Update a namespace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a namespace
    updateNamespaceData := *openapiclient.NewUpdateNamespaceData() // UpdateNamespaceData | 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateNamespace(context.Background(), id).UpdateNamespaceData(updateNamespaceData).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNamespace`: Namespace
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNamespaceData** | [**UpdateNamespaceData**](UpdateNamespaceData.md) |  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

[**Namespace**](Namespace.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNode

> Node UpdateNode(ctx, id).UpdateNodeData(updateNodeData).Execute()

Update a node



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a node
    updateNodeData := *openapiclient.NewUpdateNodeData() // UpdateNodeData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateNode(context.Background(), id).UpdateNodeData(updateNodeData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNode`: Node
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a node | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNodeData** | [**UpdateNodeData**](UpdateNodeData.md) |  | 

### Return type

[**Node**](Node.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicyGroup

> PolicyGroup UpdatePolicyGroup(ctx, id).UpdatePolicyGroupData(updatePolicyGroupData).IgnoreVersion(ignoreVersion).Execute()

Update a policy group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a policy group
    updatePolicyGroupData := *openapiclient.NewUpdatePolicyGroupData() // UpdatePolicyGroupData | 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdatePolicyGroup(context.Background(), id).UpdatePolicyGroupData(updatePolicyGroupData).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdatePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicyGroup`: PolicyGroup
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdatePolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a policy group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePolicyGroupData** | [**UpdatePolicyGroupData**](UpdatePolicyGroupData.md) |  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

[**PolicyGroup**](PolicyGroup.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> User UpdateUser(ctx, id).UpdateUserData(updateUserData).IgnoreVersion(ignoreVersion).Execute()

Update a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of a user
    updateUserData := *openapiclient.NewUpdateUserData() // UpdateUserData | 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateUser(context.Background(), id).UpdateUserData(updateUserData).IgnoreVersion(ignoreVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of a user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserData** | [**UpdateUserData**](UpdateUserData.md) |  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]

### Return type

[**User**](User.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVolume

> Volume UpdateVolume(ctx, namespaceID, id).UpdateVolumeData(updateVolumeData).IgnoreVersion(ignoreVersion).AsyncMax(asyncMax).Execute()

Update a volume



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespaceID := "namespaceID_example" // string | ID of a Namespace
    id := "id_example" // string | ID of a Volume
    updateVolumeData := *openapiclient.NewUpdateVolumeData() // UpdateVolumeData | 
    ignoreVersion := true // bool | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \"forcing\" the operation.  (optional) (default to false)
    asyncMax := "asyncMax_example" // string | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \"async-max\" header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \"300ms\", or \"2h45m\". Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\". We reject negative or nil duration values.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateVolume(context.Background(), namespaceID, id).UpdateVolumeData(updateVolumeData).IgnoreVersion(ignoreVersion).AsyncMax(asyncMax).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVolume`: Volume
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceID** | **string** | ID of a Namespace | 
**id** | **string** | ID of a Volume | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVolumeData** | [**UpdateVolumeData**](UpdateVolumeData.md) |  | 
 **ignoreVersion** | **bool** | If set to true this value indicates that the user wants to ignore entity version constraints, thereby \&quot;forcing\&quot; the operation.  | [default to false]
 **asyncMax** | **string** | Optional parameter which will make the api request asynchronous. The operation will not be cancelled even if the client disconnect. The URL parameter value overrides the \&quot;async-max\&quot; header value, if any. The value of this header defines the timeout duration for the request, it must be set to a valid duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \&quot;300ms\&quot;, or \&quot;2h45m\&quot;. Valid time units are \&quot;ns\&quot;, \&quot;us\&quot; (or \&quot;µs\&quot;), \&quot;ms\&quot;, \&quot;s\&quot;, \&quot;m\&quot;, \&quot;h\&quot;. We reject negative or nil duration values.  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

