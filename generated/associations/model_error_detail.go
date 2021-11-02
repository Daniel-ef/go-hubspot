/*
 * Associations
 *
 * Associations define the relationships between objects in HubSpot. These endpoints allow you to create, read, and remove associations.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package associations

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
