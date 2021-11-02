/*
 * Accounting Extension
 *
 * These APIs allow you to interact with HubSpot's Accounting Extension. It allows you to: * Specify the URLs that HubSpot will use when making webhook requests to your external accounting system. * Respond to webhook calls made to your external accounting system by HubSpot
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package accounting

// Represents a unit price
type UnitPrice struct {
	// The actual unit price amount.
	Amount float64 `json:"amount,omitempty"`
	// Indicates if the unit price amount already includes taxes.
	TaxIncluded bool `json:"taxIncluded"`
}
