/*
 * Associations
 *
 * Associations define the relationships between objects in HubSpot. These endpoints allow you to create, read, and remove associations.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package associations

type StandardError struct {
	Status      string              `json:"status"`
	Id          string              `json:"id,omitempty"`
	Category    *ErrorCategory      `json:"category"`
	SubCategory *interface{}        `json:"subCategory,omitempty"`
	Message     string              `json:"message"`
	Errors      []ErrorDetail       `json:"errors"`
	Context     map[string][]string `json:"context"`
	Links       map[string]string   `json:"links"`
}
