/*
 * Subscriptions
 *
 * Subscriptions allow contacts to control what forms of communications they receive. Contacts can decide whether they want to receive communication pertaining to a specific topic, brand, or an entire HubSpot account.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package communications_status

// The status of a subscription for a contact.
type PublicSubscriptionStatus struct {
	// The ID for the subscription.
	Id string `json:"id"`
	// The name of the subscription.
	Name string `json:"name"`
	// A description of the subscription.
	Description string `json:"description"`
	// Whether the contact is subscribed.
	Status string `json:"status"`
	// Where the status is determined from e.g. PORTAL_WIDE_STATUS if the contact opted out from the portal.
	SourceOfStatus string `json:"sourceOfStatus"`
	// The ID of the brand that the subscription is associated with, if there is one.
	BrandId int64 `json:"brandId,omitempty"`
	// The name of the preferences group that the subscription is associated with.
	PreferenceGroupName string `json:"preferenceGroupName,omitempty"`
	// The legal reason for the current status of the subscription.
	LegalBasis string `json:"legalBasis,omitempty"`
	// A more detailed explanation to go with the legal basis.
	LegalBasisExplanation string `json:"legalBasisExplanation,omitempty"`
}
