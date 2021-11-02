/*
 * Schemas
 *
 * The CRM uses schemas to define how custom objects should store and represent information in the HubSpot CRM. Schemas define details about an object's type, properties, and associations. The schema can be uniquely identified by its **object type ID**.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package schemas

// Defines a new object type, its properties, and associations.
type ObjectSchemaEgg struct {
	Labels *ObjectTypeDefinitionLabels `json:"labels"`
	// The names of properties that should be **required** when creating an object of this type.
	RequiredProperties []string `json:"requiredProperties"`
	// Names of properties that will be indexed for this object type in by HubSpot's product search.
	SearchableProperties []string `json:"searchableProperties"`
	// The name of the primary property for this object. This will be displayed as primary on the HubSpot record page for this object type.
	PrimaryDisplayProperty string `json:"primaryDisplayProperty,omitempty"`
	// The names of secondary properties for this object. These will be displayed as secondary on the HubSpot record page for this object type.
	SecondaryDisplayProperties []string `json:"secondaryDisplayProperties"`
	// Properties defined for this object type.
	Properties []ObjectTypePropertyCreate `json:"properties"`
	// Associations defined for this object type.
	AssociatedObjects []string `json:"associatedObjects"`
	// A unique name for this object. For internal use only.
	Name string `json:"name"`
}
