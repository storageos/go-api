/*
 * StorageOS API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.1.0
 * Contact: info@storageos.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api
// SyncProgress The progress report for an ongoing sync.  
type SyncProgress struct {
	// Number of bytes left remaining to complete the sync. 
	BytesRemaining uint64 `json:"bytesRemaining,omitempty"`
	// The average throughput of the sync given as bytes per  second. 
	ThroughputBytes uint64 `json:"throughputBytes,omitempty"`
	// The estimated time left for the sync to complete, given in seconds. When this field has a value of 0 either the  sync is complete or no duration estimate could be made. The values reported for bytesRemaining and  throughputBytes provide the client with the information needed to choose what to display. 
	EstimatedSecondsRemaining uint64 `json:"estimatedSecondsRemaining,omitempty"`
}
