/*
 * Twilio - Proxy
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

// ProxyV1Service struct for ProxyV1Service
type ProxyV1Service struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The URL we call when the interaction status changes
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// The SID of the Chat Service Instance
	ChatInstanceSid *string `json:"ChatInstanceSid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// Default TTL for a Session, in seconds
	DefaultTtl *int32 `json:"DefaultTtl,omitempty"`
	// Where a proxy number must be located relative to the participant identifier
	GeoMatchLevel *string `json:"GeoMatchLevel,omitempty"`
	// The URL we call on each interaction
	InterceptCallbackUrl *string `json:"InterceptCallbackUrl,omitempty"`
	// The URLs of resources related to the Service
	Links *map[string]interface{} `json:"Links,omitempty"`
	// The preference for Proxy Number selection for the Service instance
	NumberSelectionBehavior *string `json:"NumberSelectionBehavior,omitempty"`
	// The URL we call when an inbound call or SMS action occurs on a closed or non-existent Session
	OutOfSessionCallbackUrl *string `json:"OutOfSessionCallbackUrl,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"UniqueName,omitempty"`
	// The absolute URL of the Service resource
	Url *string `json:"Url,omitempty"`
}
