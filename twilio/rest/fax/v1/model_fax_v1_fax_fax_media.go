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

// FaxV1FaxFaxMedia struct for FaxV1FaxFaxMedia
type FaxV1FaxFaxMedia struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The content type of the stored fax media
	ContentType *string `json:"ContentType,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The SID of the fax the FaxMedia resource is associated with
	FaxSid *string `json:"FaxSid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The absolute URL of the FaxMedia resource
	Url *string `json:"Url,omitempty"`
}
