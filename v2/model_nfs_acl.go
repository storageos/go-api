/*
 * StorageOS API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.4.1-alpha
 * Contact: info@storageos.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api
// NfsAcl struct for NfsAcl
type NfsAcl struct {
	Identity NfsAclIdentity `json:"identity,omitempty"`
	SquashConfig NfsAclSquashConfig `json:"squashConfig,omitempty"`
	// The access level this ACL grants - read-only, or read-write. 
	AccessLevel string `json:"accessLevel,omitempty"`
}
