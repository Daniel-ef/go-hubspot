/*
 * Accounting Extension
 *
 * These APIs allow you to interact with HubSpot's Accounting Extension. It allows you to: * Specify the URLs that HubSpot will use when making webhook requests to your external accounting system. * Respond to webhook calls made to your external accounting system by HubSpot
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package accounting

// A response that contains the PDF of an invoice
type InvoicePdfResponse struct {
	// Designates if the response is a success ('OK') or failure ('ERR').
	Result string `json:"@result,omitempty"`
	// The bytes of the invoice PDF.
	Invoice string `json:"invoice"`
}
