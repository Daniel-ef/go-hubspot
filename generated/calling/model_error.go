/*
 * Calling Extensions API
 *
 * Provides a way for apps to add custom calling options to a contact record. This works in conjunction with the [Calling SDK](#), which is used to build your phone/calling UI. The endpoints here allow your service to appear as an option to HubSpot users when they access the *Call* action on a contact record. Once accessed, your custom phone/calling UI will be displayed in an iframe at the specified URL with the specified dimensions on that record.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package calling

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
