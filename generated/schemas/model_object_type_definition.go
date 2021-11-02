/*
 * Schemas
 *
 * The CRM uses schemas to define how custom objects should store and represent information in the HubSpot CRM. Schemas define details about an object's type, properties, and associations. The schema can be uniquely identified by its **object type ID**.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package schemas

import (
	"time"
)

// Defines an object type.
type ObjectTypeDefinition struct {
	Labels *ObjectTypeDefinitionLabels `json:"labels"`
	// The names of properties that should be **required** when creating an object of this type.
	RequiredProperties []string `json:"requiredProperties"`
	// Names of properties that will be indexed for this object type in by HubSpot's product search.
	SearchableProperties []string `json:"searchableProperties"`
	// The name of the primary property for this object. This will be displayed as primary on the HubSpot record page for this object type.
	PrimaryDisplayProperty string `json:"primaryDisplayProperty,omitempty"`
	// The names of secondary properties for this object. These will be displayed as secondary on the HubSpot record page for this object type.
	SecondaryDisplayProperties []string `json:"secondaryDisplayProperties"`
	Archived                   bool     `json:"archived"`
	// A unique ID for this object type. Will be defined as {meta-type}-{unique ID}.
	Id                 string `json:"id"`
	FullyQualifiedName string `json:"fullyQualifiedName"`
	// When the object type was created.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// When the object type was last updated.
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
	ObjectTypeId string    `json:"objectTypeId"`
	// A unique name for this object. For internal use only.
	Name string `json:"name"`
	// The ID of the account that this object type is specific to.
	PortalId int32 `json:"portalId,omitempty"`
}
