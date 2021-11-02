/*
 * CMS Site Search
 *
 * Use these endpoints for searching content on your HubSpot hosted CMS website(s).
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package site_search

type SearchHitField struct {
	MetadataField bool          `json:"metadataField"`
	Values        []interface{} `json:"values"`
	Name          string        `json:"name"`
	Value         *interface{}  `json:"value"`
}
