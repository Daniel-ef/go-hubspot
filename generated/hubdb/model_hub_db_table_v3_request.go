/*
 * HubDB endpoints
 *
 * HubDB is a relational data store that presents data as rows, columns, and cells in a table, much like a spreadsheet. HubDB tables can be added or modified [in the HubSpot CMS](https://knowledge.hubspot.com/cos-general/how-to-edit-hubdb-tables), but you can also use the API endpoints documented here. For more information on HubDB tables and using their data on a HubSpot site, see the [CMS developers site](https://designers.hubspot.com/docs/tools/hubdb). You can also see the [documentation for dynamic pages](https://designers.hubspot.com/docs/tutorials/how-to-build-dynamic-pages-with-hubdb) for more details about the `useForPages` field.  HubDB tables support `draft` and `published` versions. This allows you to update data in the table, either for testing or to allow for a manual approval process, without affecting any live pages using the existing data. Draft data can be reviewed, and published by a user working in HubSpot or published via the API. Draft data can also be discarded, allowing users to go back to the published version of the data without disrupting it. If a table is set to be `allowed for public access`, you can access the published version of the table and rows without any authentication by specifying the portal id via the query parameter `portalId`.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package hubdb

type HubDbTableV3Request struct {
	// Name of the table
	Name string `json:"name"`
	// Label of the table
	Label string `json:"label"`
	// Specifies whether the table can be used for creation of dynamic pages
	UseForPages bool `json:"useForPages,omitempty"`
	// Specifies whether the table can be read by public without authorization
	AllowPublicApiAccess bool `json:"allowPublicApiAccess,omitempty"`
	// Specifies whether child tables can be created
	AllowChildTables bool `json:"allowChildTables,omitempty"`
	// Specifies creation of multi-level dynamic pages using child tables
	EnableChildTablePages bool `json:"enableChildTablePages,omitempty"`
	// List of columns in the table
	Columns []ColumnRequest `json:"columns,omitempty"`
	// Specifies the key value pairs of the metadata fields with the associated column ids
	DynamicMetaTags map[string]int32 `json:"dynamicMetaTags,omitempty"`
}
