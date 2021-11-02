/*
 * Timeline events
 *
 * This feature allows an app to create and configure custom events that can show up in the timelines of certain CRM objects like contacts, companies, tickets, or deals. You'll find multiple use cases for this API in the sections below.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package timeline

import (
	"time"
)

type BatchResponseTimelineEventResponseWithErrors struct {
	Status      string                  `json:"status"`
	Results     []TimelineEventResponse `json:"results"`
	NumErrors   int32                   `json:"numErrors,omitempty"`
	Errors      []StandardError         `json:"errors,omitempty"`
	RequestedAt time.Time               `json:"requestedAt,omitempty"`
	StartedAt   time.Time               `json:"startedAt"`
	CompletedAt time.Time               `json:"completedAt"`
	Links       map[string]string       `json:"links,omitempty"`
}
