/*
 * Accounting Extension
 *
 * These APIs allow you to interact with HubSpot's Accounting Extension. It allows you to: * Specify the URLs that HubSpot will use when making webhook requests to your external accounting system. * Respond to webhook calls made to your external accounting system by HubSpot
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package accounting

// Representation of payment terms that are defined in the external accounting system. One of 'dueDays' or 'dueDate' is required.
type AccountingExtensionTerm struct {
	// The due date for payment of the invoice, in ISO-8601 date format (yyyy-MM-dd)
	DueDate string `json:"dueDate,omitempty"`
	// The display name of the payment terms.
	Name string `json:"name"`
	// The ID of the payment terms in the external accounting system.
	Id string `json:"id"`
	// The number of days that these payment terms represent.
	DueDays int32 `json:"dueDays,omitempty"`
}
