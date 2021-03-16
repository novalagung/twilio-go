/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.11.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountCallCallRecording struct for ApiV2010AccountCallCallRecording
type ApiV2010AccountCallCallRecording struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The API version used to make the recording
	ApiVersion *string `json:"ApiVersion,omitempty"`
	// The SID of the Call the resource is associated with
	CallSid *string `json:"CallSid,omitempty"`
	// The number of channels in the final recording file
	Channels *int32 `json:"Channels,omitempty"`
	// The Conference SID that identifies the conference associated with the recording
	ConferenceSid *string `json:"ConferenceSid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"DateUpdated,omitempty"`
	// The length of the recording in seconds
	Duration *string `json:"Duration,omitempty"`
	// How to decrypt the recording.
	EncryptionDetails *map[string]interface{} `json:"EncryptionDetails,omitempty"`
	// More information about why the recording is missing, if status is `absent`.
	ErrorCode *int32 `json:"ErrorCode,omitempty"`
	// The one-time cost of creating the recording.
	Price *float32 `json:"Price,omitempty"`
	// The currency used in the price property.
	PriceUnit *string `json:"PriceUnit,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// How the recording was created
	Source *string `json:"Source,omitempty"`
	// The start time of the recording, given in RFC 2822 format
	StartTime *string `json:"StartTime,omitempty"`
	// The status of the recording
	Status *string `json:"Status,omitempty"`
	// The recorded track. Can be: `inbound`, `outbound`, or `both`.
	Track *string `json:"Track,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"Uri,omitempty"`
}
