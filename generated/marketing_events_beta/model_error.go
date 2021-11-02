/*
 * Marketing Events Extension
 *
 * These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package marketing_events_beta

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
