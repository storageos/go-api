/*
 * Ondat API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.8.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api
// CreatePoolData struct for CreatePoolData
type CreatePoolData struct {
	// The name of the pool 
	Name string `json:"name"`
	NodeToDriveMapping map[string][]string `json:"nodeToDriveMapping"`
}
