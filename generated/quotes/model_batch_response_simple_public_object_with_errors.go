/*
 * Quotes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package quotes

import (
	"time"
)

type BatchResponseSimplePublicObjectWithErrors struct {
	Status      string               `json:"status"`
	Results     []SimplePublicObject `json:"results"`
	NumErrors   int32                `json:"numErrors,omitempty"`
	Errors      []StandardError      `json:"errors,omitempty"`
	RequestedAt time.Time            `json:"requestedAt,omitempty"`
	StartedAt   time.Time            `json:"startedAt"`
	CompletedAt time.Time            `json:"completedAt"`
	Links       map[string]string    `json:"links,omitempty"`
}
