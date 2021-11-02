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

// The state of the timeline event.
type TimelineEvent struct {
	// The event template ID.
	EventTemplateId string `json:"eventTemplateId"`
	// The email address used for contact-specific events. This can be used to identify existing contacts, create new ones, or change the email for an existing contact (if paired with the `objectId`).
	Email string `json:"email,omitempty"`
	// The CRM object identifier. This is required for every event other than contacts (where utk or email can be used).
	ObjectId string `json:"objectId,omitempty"`
	// Use the `utk` parameter to associate an event with a contact by `usertoken`. This is recommended if you don't know a user's email, but have an identifying user token in your cookie.
	Utk string `json:"utk,omitempty"`
	// The event domain (often paired with utk).
	Domain string `json:"domain,omitempty"`
	// The time the event occurred. If not passed in, the curren time will be assumed. This is used to determine where an event is shown on a CRM object's timeline.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// A collection of token keys and values associated with the template tokens.
	Tokens map[string]string `json:"tokens"`
	// Additional event-specific data that can be interpreted by the template's markdown.
	ExtraData      *interface{}         `json:"extraData,omitempty"`
	TimelineIFrame *TimelineEventIFrame `json:"timelineIFrame,omitempty"`
	// Identifier for the event. This is optional, and we recommend you do not pass this in. We will create one for you if you omit this. You can also use `{{uuid}}` anywhere in the ID to generate a unique string, guaranteeing uniqueness.
	Id string `json:"id"`
}
