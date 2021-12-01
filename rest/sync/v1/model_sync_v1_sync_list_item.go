/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// SyncV1SyncListItem struct for SyncV1SyncListItem
type SyncV1SyncListItem struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The identity of the List Item's creator
	CreatedBy *string `json:"created_by,omitempty"`
	// An arbitrary, schema-less object that the List Item stores
	Data *map[string]interface{} `json:"data,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the List Item expires
	DateExpires *time.Time `json:"date_expires,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The automatically generated index of the List Item
	Index *int `json:"index,omitempty"`
	// The SID of the Sync List that contains the List Item
	ListSid *string `json:"list_sid,omitempty"`
	// The current revision of the item, represented as a string
	Revision *string `json:"revision,omitempty"`
	// The SID of the Sync Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The absolute URL of the List Item resource
	Url *string `json:"url,omitempty"`
}
