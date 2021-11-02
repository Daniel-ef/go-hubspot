/*
 * URL redirects
 *
 * URL redirect operations
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package url_redirects

type CollectionResponseWithTotalUrlMapping struct {
	// The number of available results.
	Total int32 `json:"total"`
	// Matched URLs.
	Results []UrlMapping `json:"results"`
	Paging  *Paging      `json:"paging,omitempty"`
}
