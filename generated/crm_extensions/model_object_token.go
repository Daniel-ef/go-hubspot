/*
 * CRM cards
 *
 * Allows an app to extend the CRM UI by surfacing custom cards in the sidebar of record pages. These cards are defined up-front as part of app configuration, then populated by external data fetch requests when the record page is accessed by a user.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package crm_extensions

type ObjectToken struct {
	Name     string `json:"name,omitempty"`
	Label    string `json:"label,omitempty"`
	DataType string `json:"dataType,omitempty"`
	Value    string `json:"value"`
}
