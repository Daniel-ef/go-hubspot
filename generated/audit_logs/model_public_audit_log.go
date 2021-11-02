/*
 * CMS Audit Logs
 *
 * Use this endpoint to query audit logs of CMS changes that occurred on your HubSpot account.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package audit_logs

import (
	"time"
)

type PublicAuditLog struct {
	// The ID of the object.
	ObjectId string `json:"objectId"`
	// The ID of the user who caused the event.
	UserId string `json:"userId"`
	// The timestamp at which the event occurred.
	Timestamp time.Time `json:"timestamp"`
	// The internal name of the object in HubSpot.
	ObjectName string `json:"objectName"`
	// The name of the user who caused the event.
	FullName string `json:"fullName"`
	// The type of event that took place (CREATED, UPDATED, PUBLISHED, DELETED, UNPUBLISHED).
	Event string `json:"event"`
	// The type of the object (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.)
	ObjectType string `json:"objectType"`
}
