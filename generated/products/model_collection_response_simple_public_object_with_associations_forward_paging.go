/*
 * Products
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package products

type CollectionResponseSimplePublicObjectWithAssociationsForwardPaging struct {
	Results []SimplePublicObjectWithAssociations `json:"results"`
	Paging  *ForwardPaging                       `json:"paging,omitempty"`
}
