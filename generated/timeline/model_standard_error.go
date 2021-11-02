/*
 * Timeline events
 *
 * This feature allows an app to create and configure custom events that can show up in the timelines of certain CRM objects like contacts, companies, tickets, or deals. You'll find multiple use cases for this API in the sections below.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package timeline

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
