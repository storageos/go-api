/*
 * StorageOS API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.4.0
 * Contact: info@storageos.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api
// ResizeVolumeRequest struct for ResizeVolumeRequest
type ResizeVolumeRequest struct {
	// The desired new size for the volume in  bytes. This value cannot be less than  the current size of the volume. 
	SizeBytes uint64 `json:"sizeBytes,omitempty"`
	// An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change. 
	Version string `json:"version,omitempty"`
}
