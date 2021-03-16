/*
 * Twilio - Autopilot
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

// AutopilotV1AssistantQuery struct for AutopilotV1AssistantQuery
type AutopilotV1AssistantQuery struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The SID of the Assistant that is the parent of the resource
	AssistantSid *string `json:"AssistantSid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The SID of the [Dialogue](https://www.twilio.com/docs/autopilot/api/dialogue).
	DialogueSid *string `json:"DialogueSid,omitempty"`
	// The ISO language-country string that specifies the language used by the Query
	Language *string `json:"Language,omitempty"`
	// The SID of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) queried
	ModelBuildSid *string `json:"ModelBuildSid,omitempty"`
	// The end-user's natural language input
	Query *string `json:"Query,omitempty"`
	// The natural language analysis results that include the Task recognized and a list of identified Fields
	Results *map[string]interface{} `json:"Results,omitempty"`
	// The SID of an optional reference to the Sample created from the query
	SampleSid *string `json:"SampleSid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The communication channel from where the end-user input came
	SourceChannel *string `json:"SourceChannel,omitempty"`
	// The status of the Query
	Status *string `json:"Status,omitempty"`
	// The absolute URL of the Query resource
	Url *string `json:"Url,omitempty"`
}
