/*
 * CRM Objects
 *
 * CRM objects such as companies, contacts, deals, line items, products, tickets, and quotes are standard objects in HubSpot’s CRM. These core building blocks support custom properties, store critical information, and play a central role in the HubSpot application.  ## Supported Object Types  This API provides access to collections of CRM objects, which return a map of property names to values. Each object type has its own set of default properties, which can be found by exploring the [CRM Object Properties API](https://developers.hubspot.com/docs/methods/crm-properties/crm-properties-overview).  |Object Type |Properties returned by default | |--|--| | `companies` | `name`, `domain` | | `contacts` | `firstname`, `lastname`, `email` | | `deals` | `dealname`, `amount`, `closedate`, `pipeline`, `dealstage` | | `products` | `name`, `description`, `price` | | `tickets` | `content`, `hs_pipeline`, `hs_pipeline_stage`, `hs_ticket_category`, `hs_ticket_priority`, `subject` |  Find a list of all properties for an object type using the [CRM Object Properties](https://developers.hubspot.com/docs/methods/crm-properties/get-properties) API. e.g. `GET https://api.hubapi.com/properties/v2/companies/properties`. Change the properties returned in the response using the `properties` array in the request body.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package objects

type ModelError struct {
	// A human readable message describing the error along with remediation steps where appropriate
	Message string `json:"message"`
	// A unique identifier for the request. Include this value with any error reports or support tickets
	CorrelationId string `json:"correlationId"`
	// The error category
	Category string `json:"category"`
	// A specific category that contains more specific detail about the error
	SubCategory string `json:"subCategory,omitempty"`
	// further information about the error
	Errors []ErrorDetail `json:"errors,omitempty"`
	// Context about the error condition
	Context map[string][]string `json:"context,omitempty"`
	// A map of link names to associated URIs containing documentation about the error or recommended remediation steps
	Links map[string]string `json:"links,omitempty"`
}
