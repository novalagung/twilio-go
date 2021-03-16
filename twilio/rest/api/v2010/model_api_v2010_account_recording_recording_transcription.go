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

// ApiV2010AccountRecordingRecordingTranscription struct for ApiV2010AccountRecordingRecordingTranscription
type ApiV2010AccountRecordingRecordingTranscription struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The API version used to create the transcription
	ApiVersion *string `json:"ApiVersion,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"DateUpdated,omitempty"`
	// The duration of the transcribed audio in seconds.
	Duration *string `json:"Duration,omitempty"`
	// The charge for the transcription
	Price *float32 `json:"Price,omitempty"`
	// The currency in which price is measured
	PriceUnit *string `json:"PriceUnit,omitempty"`
	// The SID that identifies the transcription's recording
	RecordingSid *string `json:"RecordingSid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The status of the transcription
	Status *string `json:"Status,omitempty"`
	// The text content of the transcription.
	TranscriptionText *string `json:"TranscriptionText,omitempty"`
	// The transcription type
	Type *string `json:"Type,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"Uri,omitempty"`
}
