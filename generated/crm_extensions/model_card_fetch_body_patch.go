/*
 * CRM cards
 *
 * Allows an app to extend the CRM UI by surfacing custom cards in the sidebar of record pages. These cards are defined up-front as part of app configuration, then populated by external data fetch requests when the record page is accessed by a user.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package crm_extensions

// Variant of CardFetchBody with fields as optional for patches
type CardFetchBodyPatch struct {
	// URL to a service endpoint that will respond with details for this card. HubSpot will call this endpoint each time a user visits a CRM record page where this card should be displayed.
	TargetUrl string `json:"targetUrl,omitempty"`
	// An array of CRM object types where this card should be displayed. HubSpot will call your target URL whenever a user visits a record page of the types defined here.
	ObjectTypes []CardObjectTypeBody `json:"objectTypes"`
}
