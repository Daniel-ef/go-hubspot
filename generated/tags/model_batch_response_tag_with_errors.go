/*
 * Blog Post endpoints
 *
 * \"Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags\"
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tags

import (
	"time"
)

type BatchResponseTagWithErrors struct {
	Status      string            `json:"status"`
	Results     []Tag             `json:"results"`
	NumErrors   int32             `json:"numErrors,omitempty"`
	Errors      []StandardError   `json:"errors,omitempty"`
	RequestedAt time.Time         `json:"requestedAt,omitempty"`
	StartedAt   time.Time         `json:"startedAt"`
	CompletedAt time.Time         `json:"completedAt"`
	Links       map[string]string `json:"links,omitempty"`
}
