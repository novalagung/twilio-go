/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010Member struct for ApiV2010Member
type ApiV2010Member struct {
	// The SID of the Call the resource is associated with
	CallSid *string `json:"call_sid,omitempty"`
	// The date the member was enqueued
	DateEnqueued *string `json:"date_enqueued,omitempty"`
	// This member's current position in the queue.
	Position *int `json:"position,omitempty"`
	// The SID of the Queue the member is in
	QueueSid *string `json:"queue_sid,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
	// The number of seconds the member has been in the queue.
	WaitTime *int `json:"wait_time,omitempty"`
}
