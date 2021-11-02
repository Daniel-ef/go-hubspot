/*
 * Schemas
 *
 * The CRM uses schemas to define how custom objects should store and represent information in the HubSpot CRM. Schemas define details about an object's type, properties, and associations. The schema can be uniquely identified by its **object type ID**.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package schemas

// Singular and plural labels for the object. Used in CRM display.
type ObjectTypeDefinitionLabels struct {
	// The word for one object. (There’s no way to change this later.)
	Singular string `json:"singular,omitempty"`
	// The word for multiple objects. (There’s no way to change this later.)
	Plural string `json:"plural,omitempty"`
}
