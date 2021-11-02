/*
 * Schemas
 *
 * The CRM uses schemas to define how custom objects should store and represent information in the HubSpot CRM. Schemas define details about an object's type, properties, and associations. The schema can be uniquely identified by its **object type ID**.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package schemas

// The options available when a property is an enumeration
type Option struct {
	// A human-readable option label that will be shown in HubSpot.
	Label string `json:"label"`
	// The internal value of the option, which must be used when setting the property value through the API.
	Value string `json:"value"`
	// A description of the option.
	Description string `json:"description,omitempty"`
	// Options are displayed in order starting with the lowest positive integer value. Values of -1 will cause the option to be displayed after any positive values.
	DisplayOrder int32 `json:"displayOrder,omitempty"`
	// Hidden options will not be displayed in HubSpot.
	Hidden bool `json:"hidden"`
}
