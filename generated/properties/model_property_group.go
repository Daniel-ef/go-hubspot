/*
 * Properties
 *
 * All HubSpot objects store data in default and custom properties. These endpoints provide access to read and modify object properties in HubSpot.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package properties

// An ID for a group of properties
type PropertyGroup struct {
	// The internal property group name, which must be used when referencing the property group via the API.
	Name string `json:"name"`
	// A human-readable label that will be shown in HubSpot.
	Label string `json:"label"`
	// Property groups are displayed in order starting with the lowest positive integer value. Values of -1 will cause the property group to be displayed after any positive values.
	DisplayOrder int32 `json:"displayOrder"`
	Archived     bool  `json:"archived"`
}
