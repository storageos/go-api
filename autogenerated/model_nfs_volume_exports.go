/*
 * Ondat API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.8.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api
// NfsVolumeExports struct for NfsVolumeExports
type NfsVolumeExports struct {
	Exports []NfsExportConfig `json:"exports,omitempty"`
	// An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change. 
	Version string `json:"version,omitempty"`
}