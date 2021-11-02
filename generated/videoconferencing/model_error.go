/*
 * Video Conference Extension
 *
 * These APIs allow you to specify URLs that can be used to interact with a video conferencing application, to allow HubSpot to add video conference links to meeting requests with contacts.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package videoconferencing

type ModelError struct {
	// A human readable message describing the error along with remediation steps where appropriate
	Message string `json:"message"`
	// A unique identifier for the request. Include this value with any error reports or support tickets
	CorrelationId string `json:"correlationId"`
	// The error category
	Category string `json:"category"`
	// A specific category that contains more specific detail about the error
	SubCategory string `json:"subCategory,omitempty"`
	// further information about the error
	Errors []ErrorDetail `json:"errors,omitempty"`
	// Context about the error condition
	Context map[string][]string `json:"context,omitempty"`
	// A map of link names to associated URIs containing documentation about the error or recommended remediation steps
	Links map[string]string `json:"links,omitempty"`
}
