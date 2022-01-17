/*
 * StorageOS API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.5.0
 * Contact: info@storageos.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api
import (
	"time"
)
// Licence A representation of a cluster's licence properties 
type Licence struct {
	// A unique identifier for a cluster. The format of this type is undefined and may change but the defined properties will not change. 
	ClusterID string `json:"clusterID,omitempty"`
	// The time after which a licence will no longer be valid This timestamp is set when the licence is created. String format is RFC3339. 
	ExpiresAt time.Time `json:"expiresAt,omitempty"`
	// The allowed provisioning capacity in bytes This value if for the cluster, if provisioning a volume brings the cluster's total provisioned capacity above it the request will fail 
	ClusterCapacityBytes uint64 `json:"clusterCapacityBytes,omitempty"`
	// Sum of the size of all volumes in the cluster 
	UsedBytes uint64 `json:"usedBytes,omitempty"`
	// Denotes which category the licence belongs to 
	Kind string `json:"kind,omitempty"`
	// A user friendly reference to the customer 
	CustomerName string `json:"customerName,omitempty"`
	// A list of product features which are enabled by the  licence, subject to the installed version. 
	Features *[]string `json:"features,omitempty"`
	// An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change. 
	Version string `json:"version,omitempty"`
}
