/*
 * Tickets
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tickets

type CollectionResponseAssociatedId struct {
	Results []AssociatedId `json:"results"`
	Paging  *Paging        `json:"paging,omitempty"`
}
