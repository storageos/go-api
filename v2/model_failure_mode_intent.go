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
// FailureModeIntent The intent-based failure mode behaviour of a volume. The default behaviour for a volume is \"hard\", in the absence of a directly  configured intent or numerical failure threshold.  
type FailureModeIntent string

// List of FailureModeIntent
const (
	FAILUREMODEINTENT_HARD FailureModeIntent = "hard"
	FAILUREMODEINTENT_SOFT FailureModeIntent = "soft"
	FAILUREMODEINTENT_ALWAYSON FailureModeIntent = "alwayson"
)
