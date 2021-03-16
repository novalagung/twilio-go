/*
 * Twilio - Video
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

// VideoV1RoomRoomParticipant struct for VideoV1RoomRoomParticipant
type VideoV1RoomRoomParticipant struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// Duration of time in seconds the participant was connected
	Duration *int32 `json:"Duration,omitempty"`
	// The time when the participant disconnected from the room in ISO 8601 format
	EndTime *time.Time `json:"EndTime,omitempty"`
	// The string that identifies the resource's User
	Identity *string `json:"Identity,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"Links,omitempty"`
	// The SID of the participant's room
	RoomSid *string `json:"RoomSid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The time of participant connected to the room in ISO 8601 format
	StartTime *time.Time `json:"StartTime,omitempty"`
	// The status of the Participant
	Status *string `json:"Status,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"Url,omitempty"`
}
