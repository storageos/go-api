/*
 * Ondat API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.8.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api
// NfsConfig struct for NfsConfig
type NfsConfig struct {
	Exports *[]NfsExportConfig `json:"exports,omitempty"`
	// The address to which the NFS server is bound. 
	ServiceEndpoint *string `json:"serviceEndpoint,omitempty"`
}
