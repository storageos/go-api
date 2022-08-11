/*
 * Ondat API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.8.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api
// UserSessionAllOfSession struct for UserSessionAllOfSession
type UserSessionAllOfSession struct {
	// The maximum duration which the auth session  will remain valid for in seconds. 
	ExpiresInSeconds uint64 `json:"expiresInSeconds,omitempty"`
	// The JWT token for the auth session. 
	Token string `json:"token,omitempty"`
}
