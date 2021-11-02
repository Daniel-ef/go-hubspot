/*
 * Timeline events
 *
 * This feature allows an app to create and configure custom events that can show up in the timelines of certain CRM objects like contacts, companies, tickets, or deals. You'll find multiple use cases for this API in the sections below.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package timeline

type ErrorDetail struct {
	// A human readable message describing the error along with remediation steps where appropriate
	Message string `json:"message"`
	// The name of the field or parameter in which the error was found.
	In string `json:"in,omitempty"`
	// The status code associated with the error detail
	Code string `json:"code,omitempty"`
	// A specific category that contains more specific detail about the error
	SubCategory string `json:"subCategory,omitempty"`
	// Context about the error condition
	Context map[string][]string `json:"context,omitempty"`
}
