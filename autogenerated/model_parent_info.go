/*
 * Ondat API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.8.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api
// ParentInfo struct for ParentInfo
type ParentInfo struct {
	Id string `json:"id,omitempty"`
	NamespaceID string `json:"namespaceID,omitempty"`
	// A unique identifier for a snapshot. 
	SnapshotID int32 `json:"snapshotID,omitempty"`
}
