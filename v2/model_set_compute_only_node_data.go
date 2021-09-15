/*
 * StorageOS API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v2.5.0-beta1
 * Contact: info@storageos.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api
// SetComputeOnlyNodeData struct for SetComputeOnlyNodeData
type SetComputeOnlyNodeData struct {
	// Marks the node's desired configuration  state as compute-only. This will result in the node being avoided for volume placement 
	ComputeOnly bool `json:"computeOnly,omitempty"`
	// An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change. 
	Version string `json:"version,omitempty"`
}
