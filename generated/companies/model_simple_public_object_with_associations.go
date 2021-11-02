/*
 * Companies
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package companies

import (
	"time"
)

type SimplePublicObjectWithAssociations struct {
	Id           string                                    `json:"id"`
	Properties   map[string]string                         `json:"properties"`
	CreatedAt    time.Time                                 `json:"createdAt"`
	UpdatedAt    time.Time                                 `json:"updatedAt"`
	Archived     bool                                      `json:"archived,omitempty"`
	ArchivedAt   time.Time                                 `json:"archivedAt,omitempty"`
	Associations map[string]CollectionResponseAssociatedId `json:"associations,omitempty"`
}
