/*
 * Line Items
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package line_items

type CollectionResponseAssociatedIdForwardPaging struct {
	Results []AssociatedId `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}
