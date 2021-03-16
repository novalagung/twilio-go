/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.11.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// EventsV1EventType struct for EventsV1EventType
type EventsV1EventType struct {
	// The date this Event Type was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date this Event Type was updated.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// Event Type description.
	Description *string                 `json:"Description,omitempty"`
	Links       *map[string]interface{} `json:"Links,omitempty"`
	// The Schema identifier for this Event Type.
	SchemaId *string `json:"SchemaId,omitempty"`
	// The Event Type identifier.
	Type *string `json:"Type,omitempty"`
	// The URL of this resource.
	Url *string `json:"Url,omitempty"`
}
