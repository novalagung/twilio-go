/*
 * Twilio - Taskrouter
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

// TaskrouterV1WorkspaceActivity struct for TaskrouterV1WorkspaceActivity
type TaskrouterV1WorkspaceActivity struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// Whether the Worker should be eligible to receive a Task when it occupies the Activity
	Available *bool `json:"Available,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The string that you assigned to describe the Activity resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The absolute URL of the Activity resource
	Url *string `json:"Url,omitempty"`
	// The SID of the Workspace that contains the Activity
	WorkspaceSid *string `json:"WorkspaceSid,omitempty"`
}
