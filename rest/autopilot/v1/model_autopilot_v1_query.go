/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.22.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// AutopilotV1Query struct for AutopilotV1Query
type AutopilotV1Query struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Assistant that is the parent of the resource
	AssistantSid *string `json:"assistant_sid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The SID of the [Dialogue](https://www.twilio.com/docs/autopilot/api/dialogue).
	DialogueSid *string `json:"dialogue_sid,omitempty"`
	// The ISO language-country string that specifies the language used by the Query
	Language *string `json:"language,omitempty"`
	// The SID of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) queried
	ModelBuildSid *string `json:"model_build_sid,omitempty"`
	// The end-user's natural language input
	Query *string `json:"query,omitempty"`
	// The natural language analysis results that include the Task recognized and a list of identified Fields
	Results *map[string]interface{} `json:"results,omitempty"`
	// The SID of an optional reference to the Sample created from the query
	SampleSid *string `json:"sample_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The communication channel from where the end-user input came
	SourceChannel *string `json:"source_channel,omitempty"`
	// The status of the Query
	Status *string `json:"status,omitempty"`
	// The absolute URL of the Query resource
	Url *string `json:"url,omitempty"`
}
