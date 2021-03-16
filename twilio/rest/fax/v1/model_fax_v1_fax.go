/*
 * Twilio - Fax
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

// FaxV1Fax struct for FaxV1Fax
type FaxV1Fax struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The API version used to transmit the fax
	ApiVersion *string `json:"ApiVersion,omitempty"`
	// The ISO 8601 formatted date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 formatted date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The direction of the fax
	Direction *string `json:"Direction,omitempty"`
	// The time it took to transmit the fax
	Duration *int32 `json:"Duration,omitempty"`
	// The number the fax was sent from
	From *string `json:"From,omitempty"`
	// The URLs of the fax's related resources
	Links *map[string]interface{} `json:"Links,omitempty"`
	// The SID of the FaxMedia resource that is associated with the Fax
	MediaSid *string `json:"MediaSid,omitempty"`
	// The Twilio-hosted URL that can be used to download fax media
	MediaUrl *string `json:"MediaUrl,omitempty"`
	// The number of pages contained in the fax document
	NumPages *int32 `json:"NumPages,omitempty"`
	// The fax transmission price
	Price *float32 `json:"Price,omitempty"`
	// The ISO 4217 currency used for billing
	PriceUnit *string `json:"PriceUnit,omitempty"`
	// The quality of the fax
	Quality *string `json:"Quality,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The status of the fax
	Status *string `json:"Status,omitempty"`
	// The phone number that received the fax
	To *string `json:"To,omitempty"`
	// The absolute URL of the fax resource
	Url *string `json:"Url,omitempty"`
}
